package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cid := crypto.Id()

	source, err := workloadapi.NewX509Source(
		ctx,
		workloadapi.WithClientOptions(
			workloadapi.WithAddr(env.SpiffeSocketUrl()),
		),
	)

	if err != nil {
		log.FatalLn(&cid, "Unable to fetch X.509 Bundle", err.Error())
		return
	}

	if source == nil {
		log.FatalLn(&cid, "Could not find source")
		return
	}

	svid, err := source.GetX509SVID()
	if err != nil {
		log.FatalLn(&cid, "Unable to get X.509 SVID from source bundle", err.Error())
		return
	}

	if svid == nil {
		log.FatalLn(&cid, "Could not find SVID in source bundle")
		return
	}

	svidId := svid.ID
	if !validation.IsRelayClient(svidId.String()) {
		log.FatalLn(
			&cid,
			"SpiffeId check: RelayClient:bootstrap: I don't know you, and it's crazy:",
			svidId.String(),
		)
		return
	}

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsRelayServer(id.String()) {
			return nil
		}

		return fmt.Errorf(
			"TLS Config: I don't know you, and it's crazy '%s'", id.String(),
		)
	})

	tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	serverURL := env.RelayServerUrl()

	for {
		err := sendPostRequest(client, serverURL, cid)
		if err != nil {
			log.WarnLn(&cid, "Failed to send POST request:", err.Error())
		}

		time.Sleep(30 * time.Second) // Wait for 30 seconds before sending the next request
	}
}

func sendPostRequest(client *http.Client, url string, cid string) error {
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.WarnLn(&cid, "Failed to close response body:", err.Error())
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	log.DebugLn(&cid, "Response status:", resp.Status)
	log.DebugLn(&cid, "Response body:", string(body))

	return nil
}
