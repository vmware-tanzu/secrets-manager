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

	var secrets []map[string]interface{}
	err = json.Unmarshal([]byte(sfr.Data), &secrets)
	if err != nil {
		log.Fatalf("Error unmarshalling secrets: %v", err)
	}

	secretsToServe = make(map[string]string)

	var serverCert, serverKey string

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
		default:
			if strings.HasPrefix(name, "raw:") &&
				!strings.HasPrefix(name, "raw:vsecm-scout") {
				secretsToServe[strings.TrimPrefix(name, "raw:")] = value
			}
		}
	}

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

		// Minimum TLS version
		MinVersion: tls.VersionTLS12,

		// Preferred cipher suites
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		},

		// Disable TLS renegotiation
		Renegotiation: tls.RenegotiateNever,

		// Enable HTTP/2
		NextProtos: []string{"h2", "http/1.1"},

		// Enable client authentication if needed
		// ClientAuth: tls.RequireAndVerifyClientCert,

		// Curve preferences
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	http.HandleFunc("/webhook", webhookHandler)

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on :8443 with TLS enabled")
	log.Fatal(server.ListenAndServeTLS("", ""))
	//                                 ^   ^
	// Empty strings because we've already provided the cert and key in TLSConfig.
}
