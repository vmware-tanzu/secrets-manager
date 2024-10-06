package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"net/http"
	"os"
	"strings"
)

type Endpoint struct {
	Name              string `json:"name"`
	BundleEndpointURL string `json:"bundleEndpointUrl"`
	TrustDomain       string `json:"trustDomain"`
	EndpointSPIFFEID  string `json:"endpointSPIFFEID"`
}

type Endpoints struct {
	Endpoints    map[string]Endpoint `json:""`
	FederateWith []string            `json:"federateWith"`
}

type Secret struct {
	Name         string   `json:"name"`
	Value        []string `json:"value"`
	Created      string   `json:"created"`
	Updated      string   `json:"updated"`
	NotBefore    string   `json:"notBefore"`
	ExpiresAfter string   `json:"expiresAfter"`
}

type Secrets struct {
	Secrets   []Secret `json:"secrets"`
	Algorithm string   `json:"algorithm"`
}

func loadEndpoints(filename string) Endpoints {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading endpoints file: " + err.Error())
	}

	var endpoints Endpoints
	err = json.Unmarshal(data, &endpoints)
	if err != nil {
		panic("Error parsing endpoints JSON: " + err.Error())
	}

	return endpoints
}

func loadSecrets(filename string) Secrets {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading secrets file: " + err.Error())
	}

	var secrets Secrets
	err = json.Unmarshal(data, &secrets)
	if err != nil {
		panic("Error parsing secrets JSON: " + err.Error())
	}

	return secrets
}

func handleRequest(w http.ResponseWriter, r *http.Request, endpoints Endpoints, secrets Secrets) {
	// Extract SPIFFE ID from the client certificate
	if r.TLS == nil || len(r.TLS.PeerCertificates) == 0 {
		http.Error(w, "No client certificate provided", http.StatusUnauthorized)
		return
	}

	spiffeID := r.TLS.PeerCertificates[0].URIs[0].String()
	fmt.Printf("Received request from SPIFFE ID: %s\n", spiffeID)

	// Extract trust domain from SPIFFE ID
	parts := strings.Split(spiffeID, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid SPIFFE ID", http.StatusBadRequest)
		return
	}
	trustDomain := parts[2]

	// Find the corresponding secret
	secretName := fmt.Sprintf("vsecm-relay:%s", trustDomain)
	for _, secret := range secrets.Secrets {
		if secret.Name == secretName {
			// Send the secret value
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(secret.Value)
			return
		}
	}

	http.Error(w, "No secret found for the given SPIFFE ID", http.StatusNotFound)
}

func main() {
	fmt.Println("Starting mTLS Secret Relay Server...")

	// Load endpoints and secrets
	endpoints := loadEndpoints("/vsecm-relay/data/endpoints.json")
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
			handleRequest(w, r, endpoints, secrets)
		}),
	}

	fmt.Println("Starting server on https://0.0.0.0:443")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		panic("Error starting server: " + err.Error())
	}
}
