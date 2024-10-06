package main

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
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
	"net/url"
	"os"
	"strings"
	"time"
)

type EncryptedResponse struct {
	EncryptedData string `json:"encryptedData"`
	Signature     string `json:"signature"`
}

func generateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func encodePublicKeyToPEM(publicKey *rsa.PublicKey) (string, error) {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})
	return string(publicKeyPEM), nil
}

func decodePublicKeyFromPEM(pemStr string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
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

func verifySignature(data []byte, signature []byte, publicKey *rsa.PublicKey) error {
	hash := sha256.Sum256(data)
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
}

func run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, err := workloadapi.NewX509Source(
		ctx,
		workloadapi.WithClientOptions(
			workloadapi.WithAddr("unix:///spire-agent-socket/spire-agent.sock"),
		),
	)
	if err != nil {
		panic("Error acquiring source")
	}
	defer func(source *workloadapi.X509Source) {
		err := source.Close()
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

	// Generate a new keypair for this request
	privateKey, publicKey, err := generateKeyPair()
	if err != nil {
		panic("Error generating keypair")
	}

	// Encode the public key to PEM format
	publicKeyPEM, err := encodePublicKeyToPEM(publicKey)
	if err != nil {
		panic("Error encoding public key to PEM")
	}

	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   tlsconfig.MTLSClientConfig(source, source, authorizer),
		},
	}

	// Create a new request with the public key in the body
	req, err := http.NewRequest("POST", p, strings.NewReader(publicKeyPEM))
	if err != nil {
		panic("Error creating request")
	}
	req.Header.Set("Content-Type", "application/x-pem-file")

	r, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		panic("error getting from client")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing body")
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic("error reading body")
	}

	var response EncryptedResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic("error unmarshaling response")
	}

	// Decode the server's public key from the response header
	serverPublicKeyPEM, err := base64.StdEncoding.DecodeString(r.Header.Get("X-Public-Key"))
	if err != nil {
		panic("error decoding server public key")
	}
	serverPublicKey, err := decodePublicKeyFromPEM(string(serverPublicKeyPEM))
	if err != nil {
		panic("error parsing server public key")
	}

	// Decode the encrypted data and signature
	encryptedData, err := base64.StdEncoding.DecodeString(response.EncryptedData)
	if err != nil {
		panic("error decoding encrypted data")
	}
	signature, err := base64.StdEncoding.DecodeString(response.Signature)
	if err != nil {
		panic("error decoding signature")
	}

	// Verify the signature
	err = verifySignature(encryptedData, signature, serverPublicKey)
	if err != nil {
		panic("signature verification failed")
	}

	// Decrypt the data using the client's private key
	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedData)
	if err != nil {
		panic("error decrypting data")
	}

	fmt.Printf("My secret is: '%s'.\n", string(decryptedData))
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
