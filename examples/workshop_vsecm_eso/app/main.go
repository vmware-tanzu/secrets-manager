package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/vmware-tanzu/secrets-manager/sdk/sentry"
	"log"
	"net/http"
	"strings"
)

var (
	jwtSecret      string
	secretsToServe map[string]string
)

func main() {
	fmt.Println("Fetching secrets...")
	sfr, err := sentry.Fetch()
	if err != nil {
		log.Fatalf("Error fetching secrets: %v", err)
	}

	// fmt.Println("data", sfr.Data)

	var secrets []map[string]interface{}
	err = json.Unmarshal([]byte(sfr.Data), &secrets)
	if err != nil {
		log.Fatalf("Error unmarshalling secrets: %v", err)
	}

	secretsToServe = make(map[string]string)

	var serverCert, serverKey string
	//, caCert string

	for _, secret := range secrets {
		name := secret["name"].(string)
		value := secret["value"].(string)

		switch name {
		case "raw:vsecm-scout-jwt-secret":
			jwtSecret = value
		case "raw:vsecm-scout-crt":
			serverCert = value
		case "raw:vsecm-scout-key":
			serverKey = value
		//case "raw:vsecm-scout-ca-crt":
		//	caCert = value
		default:
			if strings.HasPrefix(name, "raw:") && !strings.HasPrefix(name, "raw:vsecm-scout") {
				secretsToServe[strings.TrimPrefix(name, "raw:")] = value
			}
		}
	}

	// Configure TLS
	cert, err := tls.X509KeyPair([]byte(serverCert), []byte(serverKey))
	if err != nil {
		log.Fatalf("Error loading server certificate and key: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		// You might want to add more TLS configuration here
	}

	http.HandleFunc("/webhook", webhookHandler)

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on :8443 with TLS enabled")
	log.Fatal(server.ListenAndServeTLS("", "")) // Empty strings because we've already provided the cert and key in TLSConfig
}
