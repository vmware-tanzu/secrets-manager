/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"context"
	"fmt"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var source *workloadapi.X509Source

	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(
			workloadapi.WithAddr("unix:///spire-agent-socket/spire-agent.sock"),
		),
	)

	if err != nil {
		panic("Error acquiring source")
	}

	defer func(s *workloadapi.X509Source) {
		if s == nil {
			return
		}
		err := s.Close()
		if err != nil {
			fmt.Println("error closing source")
		}
	}(source)

	svid, err := source.GetX509SVID()
	if err != nil {
		panic("error getting svid")
	}
	fmt.Println("svid ID: ", svid.ID.String())

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		// In a real-world scenario, you'd implement proper authorization logic here
		return nil
	})

	baseURL := os.Getenv("CONTROL_PLANE_URL")
	if baseURL == "" {
		baseURL = "https://10.211.55.112"
	}

	p, err := url.JoinPath(baseURL, "/")
	if err != nil {
		panic("problem in url")
	}

	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   tlsconfig.MTLSClientConfig(source, source, authorizer),
		},
	}

	r, err := client.Get(p)
	if err != nil {
		fmt.Println(err.Error())
		panic("error getting from client")
	}

	defer func(b io.ReadCloser) {
		_ = b.Close()
	}(r.Body)

	r.Close = true

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic("error reading body")
	}

	fmt.Printf("My secret is: '%s'.", string(body))
}

func main() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			run()
		}
	}
}
