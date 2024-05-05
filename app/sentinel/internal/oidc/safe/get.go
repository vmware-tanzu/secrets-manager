/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package safe

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

func acquireSource(ctx context.Context) (*workloadapi.X509Source, bool) {
	resultChan := make(chan *workloadapi.X509Source)
	errorChan := make(chan error)

	cid := ctx.Value("correlationId").(*string)

	go func() {
		source, err := workloadapi.NewX509Source(
			ctx, workloadapi.WithClientOptions(
				workloadapi.WithAddr(env.SpiffeSocketUrl()),
			),
		)

		if err != nil {
			errorChan <- err
			return
		}

		if err != nil {
			log.Println(cid, "acquireSource: I am having trouble fetching my identity from SPIRE.", err.Error())
			log.Println(cid,
				"acquireSource: I won't proceed until you put me in a secured container.", err.Error())
			errorChan <- err
			return
		}
		resultChan <- source
	}()

	select {
	case source := <-resultChan:
		return source, true
	case err := <-errorChan:
		log.Println(cid, "acquireSource: I cannot execute command because I cannot talk to SPIRE.", err.Error())
		return nil, false
	case <-ctx.Done():
		log.Println(cid, "acquireSource: Operation was cancelled.")
		return nil, false
	}
}

func Get(ctx context.Context, r *http.Request, showEncryptedSecrets bool) (string, error) {
	cid := ctx.Value("correlationId").(*string)
	log.Println(cid, "Get: start")

	source, proceed := acquireSource(ctx)
	defer func(s *workloadapi.X509Source) {
		if s == nil {
			return
		}
		err := s.Close()
		if err != nil {
			log.Println(cid, "Get: Problem closing the workload source.", err.Error())
		}
	}(source)
	if !proceed {
		return "", fmt.Errorf("could not proceed")
	}

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New("I don't know you, and it's crazy: '" + id.String() + "'")
	})

	safeUrl := "/sentinel/v1/secrets"
	if showEncryptedSecrets {
		safeUrl = "/sentinel/v1/secrets?reveal=true"
	}

	p, err := url.JoinPath(env.EndpointUrlForSafe(), safeUrl)
	if err != nil {
		log.Println(cid, "Get: I am having problem generating VSecM Safe secrets api endpoint URL.", err.Error())
		return "", fmt.Errorf("get: I am having problem generating VSecM Safe secrets api endpoint URL: %v", err.Error())
	}

	tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	req, err := http.NewRequest(http.MethodGet, p, nil)
	selectedHeaders := []string{"Authorization", "Content-Type", "ClientId", "ClientSecret", "UserName"}
	for _, headerName := range selectedHeaders {
		if value := r.Header.Get(headerName); value != "" {
			req.Header.Set(headerName, value)
		}
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(cid, "Get: Problem connecting to VSecM Safe API endpoint URL.", err.Error())
		return "", fmt.Errorf("get: Problem connecting to VSecM Safe API endpoint URL: %v", err.Error())
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.Println(cid, "Get: Problem closing request body.", err.Error())
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(cid, "Get: Unable to read the response body from VSecM Safe.", err.Error())
		return "", fmt.Errorf("get: Unable to read the response body from VSecM Safe: %v", err.Error())
	}

	return string(body), nil
}
