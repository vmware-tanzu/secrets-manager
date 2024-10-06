package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func handleRequest(w http.ResponseWriter, r *http.Request, secrets Secrets) {
	fmt.Println("handle request", "method", r.Method)

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
		fmt.Println("Invalid SPIFFE ID")
		http.Error(w, "Invalid SPIFFE ID", http.StatusBadRequest)
		return
	}
	trustDomain := parts[2]

	// Read the client's public key from the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error in request body")
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	clientPublicKey, err := parsePublicKey(string(body))
	if err != nil {
		fmt.Println("error in parsing public key")
		http.Error(w, "Invalid client public key", http.StatusBadRequest)
		return
	}

	// Generate a new keypair for this response
	privateKey, publicKey, err := generateKeyPair()
	if err != nil {
		http.Error(w, "Error generating keypair", http.StatusInternalServerError)
		return
	}

	// Find the corresponding secret
	secretName := fmt.Sprintf("vsecm-relay:%s", trustDomain)
	var secretValue []string
	for _, secret := range secrets.Secrets {
		if secret.Name == secretName {
			secretValue = secret.Value
			break
		}
	}

	if secretValue == nil {
		http.Error(w, "No secret found for the given SPIFFE ID", http.StatusNotFound)
		return
	}

	// Generate AES key
	aesKey, err := generateAESKey()
	if err != nil {
		http.Error(w, "Error generating AES key", http.StatusInternalServerError)
		return
	}

	// Encrypt the AES key with the client's public key
	encryptedAESKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, clientPublicKey, aesKey, nil)
	if err != nil {
		http.Error(w, "Error encrypting AES key", http.StatusInternalServerError)
		return
	}

	// Encrypt the secret with AES
	plaintext, err := json.Marshal(secretValue)
	if err != nil {
		http.Error(w, "Error marshaling secret", http.StatusInternalServerError)
		return
	}
	encryptedData, err := encryptAES(plaintext, aesKey)
	if err != nil {
		http.Error(w, "Error encrypting data", http.StatusInternalServerError)
		return
	}

	// Sign the encrypted data
	signature, err := signData(encryptedData, privateKey)
	if err != nil {
		http.Error(w, "Error signing data", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	response := EncryptedResponse{
		EncryptedAESKey: base64.StdEncoding.EncodeToString(encryptedAESKey),
		EncryptedData:   base64.StdEncoding.EncodeToString(encryptedData),
		Signature:       base64.StdEncoding.EncodeToString(signature),
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		http.Error(w, "Error encoding public key", http.StatusInternalServerError)
		return
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	// Add the public key to the response headers
	w.Header().Set("X-Public-Key", base64.StdEncoding.EncodeToString(publicKeyPEM))

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
