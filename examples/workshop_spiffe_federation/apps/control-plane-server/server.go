package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"net/http"
)

func main() {
	fmt.Println("Starting mTLS Secret Relay Server...")

	// Load endpoints and secrets
	// endpoints = loadEndpoints("/vsecm-relay/data/endpoints.json")
	secrets := loadSecrets("/vsecm-relay/data/secrets.json")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, err := workloadapi.NewX509Source(
		ctx,
		workloadapi.WithClientOptions(
			workloadapi.WithAddr("unix:///spire-agent-socket/spire-agent.sock"),
		),
	)
	if err != nil {
		panic("Error acquiring X.509 source: " + err.Error())
	}
	defer func(source *workloadapi.X509Source) {
		err := source.Close()
		if err != nil {
			fmt.Println("Error closing X.509 source: " + err.Error())
		}
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

	server := &http.Server{
		Addr:      ":443",
		TLSConfig: serverConfig,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handleRequest(w, r, secrets)
		}),
	}

	fmt.Println("Starting server on https://0.0.0.0:443")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		panic("Error starting server: " + err.Error())
	}
}
