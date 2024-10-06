package main

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"io"
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

type EncryptedResponse struct {
	EncryptedData string `json:"encryptedData"`
	Signature     string `json:"signature"`
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

func encryptData(data []string, publicKey *rsa.PublicKey) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, jsonData)
	if err != nil {
		return nil, err
	}

	return encryptedData, nil
}

func signData(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, fmt.Errorf("error signing data: %v", err)
	}

	return signature, nil
}

var (
	endpoints Endpoints
	secrets   Secrets
	source    *workloadapi.X509Source
)

func parsePublicKey(publicKeyPEM string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return rsaPub, nil
}

//func generateKeyPair() error {
//	// Generate a new private key
//	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
//	if err != nil {
//		return err
//	}
//
//	// Create a self-signed certificate template
//	template := x509.Certificate{
//		SerialNumber: big.NewInt(1),
//		Subject: pkix.Name{
//			Organization: []string{"Your Organization"},
//		},
//		NotBefore:             time.Now(),
//		NotAfter:              time.Now().Add(365 * 24 * time.Hour), // Valid for 1 year
//		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
//		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
//		BasicConstraintsValid: true,
//	}
//
//	// Create the certificate
//	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
//	if err != nil {
//		return err
//	}
//
//	// Save the private key to a file
//	privateKeyFile, err := os.Create("private.key")
//	if err != nil {
//		return err
//	}
//	defer privateKeyFile.Close()
//
//	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
//	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
//		return err
//	}
//
//	// Save the certificate to a file
//	certFile, err := os.Create("cert.pem")
//	if err != nil {
//		return err
//	}
//	defer certFile.Close()
//
//	if err := pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
//		return err
//	}
//
//	return nil
//}

func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("error generating keypair")
		http.Error(w, "Error generating keypair", http.StatusInternalServerError)
		return
	}

	// Read the client's public key from the request body
	body, err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading request body")
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	clientPublicKey, err = parsePublicKey(string(body))
	if err != nil {
		fmt.Println("error parsing public key")
		http.Error(w, "Invalid client public key", http.StatusBadRequest)
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
		fmt.Println("No secret found for the given SPIFFE ID")
		http.Error(w, "No secret found for the given SPIFFE ID", http.StatusNotFound)
		return
	}

	// Encrypt the secret using the client's public key
	encryptedData, err := encryptData(secretValue, clientPublicKey)
	if err != nil {
		fmt.Println("error encrypting data")
		http.Error(w, "Error encrypting data", http.StatusInternalServerError)
		return
	}

	// Sign the encrypted data using the newly generated private key
	signature, err := signData(encryptedData, privateKey)
	if err != nil {
		fmt.Println("error signing data")
		http.Error(w, "Error signing data", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	response := EncryptedResponse{
		EncryptedData: base64.StdEncoding.EncodeToString(encryptedData),
		Signature:     base64.StdEncoding.EncodeToString(signature),
	}

	// Encode the public key to PEM format
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})

	// Add the public key to the response headers
	w.Header().Set("X-Public-Key", base64.StdEncoding.EncodeToString(publicKeyPEM))

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Starting mTLS Secret Relay Server...")

	// Load endpoints and secrets
	endpoints = loadEndpoints("/vsecm-relay/data/endpoints.json")
	secrets = loadSecrets("/vsecm-relay/data/secrets.json")

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
			handleRequest(w, r)
		}),
	}

	fmt.Println("Starting server on https://0.0.0.0:443")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		panic("Error starting server: " + err.Error())
	}
}
