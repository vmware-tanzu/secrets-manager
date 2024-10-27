package net

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spiffe/vsecm-sdk-go/sentry"
	"log"
)

var (
	jwtSecret      string
	secretsToServe map[string]string
)

func TlsConfig() *tls.Config {
	fmt.Println("Fetching secrets...")
	sfr, err := sentry.Fetch()
	if err != nil {
		log.Fatalf("Error fetching secrets: %v", err)
	}

	var secrets []map[string]interface{}
	err = json.Unmarshal([]byte(sfr.Data), &secrets)
	if err != nil {
		log.Fatalf("Error unmarshalling secrets: %v", err.Error())
	}

	secretsToServe = make(map[string]string)

	var serverCert, serverKey string

	for _, secret := range secrets {
		name := secret["name"].(string)
		value := secret["value"].(string)

		switch name {
		// do this initialization elsewhere
		// also you might need a lock since jwtsecret is a shared resource.
		//case "raw:vsecm-scout-jwt-secret":
		//	jwtSecret = value
		case "raw:vsecm-scout-crt":
			serverCert = value
		case "raw:vsecm-scout-key":
			serverKey = value

			// This is not related to TLS config. Move it elsewhere.
			// Ideally, update it in a loop. Also, `secretsToServe` is a shared
			// resource; so you might want a thread-safe map for it.
			//default:
			//	if strings.HasPrefix(name, "raw:") &&
			//		!strings.HasPrefix(name, "raw:vsecm-scout") {
			//		secretsToServe[strings.TrimPrefix(name, "raw:")] = value
			//	}
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

	return &tls.Config{
		Certificates: []tls.Certificate{cert},

		// Further configuration options to strengthen the server's security.
		// Can be applied via env configuration.
		//
		// Minimum TLS version
		// MinVersion: tls.VersionTLS12,
		//
		// Preferred cipher suites
		// CipherSuites: []uint16{
		//	tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		//	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		//	tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		//	tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		// },
		//
		// Disable TLS renegotiation
		// Renegotiation: tls.RenegotiateNever,
		//
		// Enable HTTP/2
		// NextProtos: []string{"h2", "http/1.1"},
		//
		// Enable client authentication if needed
		// ClientAuth: tls.RequireAndVerifyClientCert,
		//
		// Curve preferences
		// CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}
}
