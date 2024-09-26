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
	"crypto/tls"
	"fmt"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"net/http"
	"sync"
)

func main() {
	fmt.Println("In main...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("Before querying the workload api")

	source, err := workloadapi.NewX509Source(
		ctx,
		workloadapi.WithClientOptions(
			workloadapi.WithAddr("unix:///spire-agent-socket/spire-agent.sock"),
		),
	)

	fmt.Println("After querying the workload api")

	if err != nil {
		panic("Error acquiring X.509 source")
	}
	defer func(source *workloadapi.X509Source) {
		_ = source.Close()
	}(source)

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		// In a real-world scenario, you'd implement proper authorization logic here
		return nil
	})

	serverConfig := &tls.Config{
		ClientAuth: tls.RequireAnyClientCert,
		GetConfigForClient: func(*tls.ClientHelloInfo) (*tls.Config, error) {
			return tlsconfig.MTLSServerConfig(source, source, authorizer), nil
		},
	}

	var counter = 0
	var counterLock sync.Mutex

	server := &http.Server{
		Addr:      ":443",
		TLSConfig: serverConfig,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			counterLock.Lock()
			defer counterLock.Unlock()
			counter = counter + 1
			_, _ = fmt.Fprintf(w, "hello: %d", counter)
		}),
	}

	fmt.Println("Starting server on https://0.0.0.0:443")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		panic("Error starting server: " + err.Error())
	}
	fmt.Println("Server started.")
}
