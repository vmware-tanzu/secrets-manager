package main

import (
	"crypto/tls"
	"encoding/base64"
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

	fmt.Println("---------")
	fmt.Println("server cert")
	fmt.Println(serverCert)
	fmt.Println("server key")
	fmt.Println(serverKey)
	fmt.Println("--------")

	// Decode base64 encoded certificate and key
	decodedCert, err := base64.StdEncoding.DecodeString(serverCert)
	if err != nil {
		log.Fatalf("Error decoding server certificate: %v", err)
	}

	decodedKey, err := base64.StdEncoding.DecodeString(serverKey)
	if err != nil {
		log.Fatalf("Error decoding server key: %v", err)
	}

	// Configure TLS with decoded certificate and key
	cert, err := tls.X509KeyPair(decodedCert, decodedKey)
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
