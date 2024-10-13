package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type InitRequest struct {
	ClientPublicKey string `json:"clientPublicKey"`
	ClientNonce     string `json:"clientNonce"`
	SignedNonce     string `json:"signedNonce"`
}

type InitResponse struct {
	ServerPublicKey string `json:"serverSignPublicKey"`
	Nonce           string `json:"nonce"`
	SignedNonce     string `json:"signedNonce"`
	EncryptedAESKey string `json:"encryptedAESKey"`
}

var (
	serverRSAKey *rsa.PrivateKey
)

func main() {
	var err error
	serverRSAKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate server RSA key: %v", err)
	}

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081"}
	config.AllowMethods = []string{"POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	r.Use(cors.New(config))

	r.POST("/init", handleInit)
	r.Run(":8080")
}

func handleInit(c *gin.Context) {
	var req InitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received init request: %+v", req)

	if req.ClientPublicKey == "" || req.ClientNonce == "" || req.SignedNonce == "" {
		log.Println("Missing required fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	clientPublicKey, err := parsePublicKey(req.ClientPublicKey)
	if err != nil {
		log.Printf("Invalid client public key: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client public key"})
		return
	}

	log.Printf("Verifying signature - Nonce: %s, SignedNonce: %s, PublicKey: %s", req.ClientNonce, req.SignedNonce, req.ClientPublicKey)

	// Verify client's signed nonce
	if !verifySignature(req.ClientNonce, req.SignedNonce, clientPublicKey) {
		log.Printf("Invalid client signature. Nonce: %s, SignedNonce: %s", req.ClientNonce, req.SignedNonce)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid client signature"})
		return
	}

	// Generate server nonce
	nonce := make([]byte, 32)
	if _, err := rand.Read(nonce); err != nil {
		log.Printf("Failed to generate nonce: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate nonce"})
		return
	}

	// Sign server nonce
	signedNonce, err := signData(nonce, serverRSAKey)
	if err != nil {
		log.Printf("Failed to sign nonce: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign nonce"})
		return
	}

	// Generate AES key
	aesKey := make([]byte, 32)
	if _, err := rand.Read(aesKey); err != nil {
		log.Printf("Failed to generate AES key: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate AES key"})
		return
	}

	// Encrypt AES key with client's public key
	encryptedAESKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, clientPublicKey, aesKey, nil)
	if err != nil {
		log.Printf("Failed to encrypt AES key: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt AES key"})
		return
	}

	response := InitResponse{
		ServerPublicKey: exportPublicKey(&serverRSAKey.PublicKey),
		Nonce:           base64.StdEncoding.EncodeToString(nonce),
		SignedNonce:     base64.StdEncoding.EncodeToString(signedNonce),
		EncryptedAESKey: base64.StdEncoding.EncodeToString(encryptedAESKey),
	}

	log.Printf("Sending response: %+v", response)
	c.JSON(http.StatusOK, response)
}

func verifySignature(nonce, signedNonce string, publicKey *rsa.PublicKey) bool {
	nonceBytes, err := base64.StdEncoding.DecodeString(nonce)
	if err != nil {
		log.Printf("Failed to decode nonce: %v", err)
		return false
	}

	signatureBytes, err := base64.StdEncoding.DecodeString(signedNonce)
	if err != nil {
		log.Printf("Failed to decode signature: %v", err)
		return false
	}

	hashed := sha256.Sum256(nonceBytes)
	err = rsa.VerifyPSS(publicKey, crypto.SHA256, hashed[:], signatureBytes, &rsa.PSSOptions{
		SaltLength: 32, // Ensure this matches the client's salt length
	})
	if err != nil {
		log.Printf("Signature verification failed: %v", err)
		return false
	}
	return true
}

func signData(data []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	hashed := sha256.Sum256(data)
	return rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed[:], &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
	})
}

func parsePublicKey(publicKeyString string) (*rsa.PublicKey, error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyString)
	if err != nil {
		return nil, err
	}

	pub, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	return rsaPub, nil
}

func exportPublicKey(publicKey *rsa.PublicKey) string {
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatalf("Failed to marshal public key: %v", err)
	}
	return base64.StdEncoding.EncodeToString(pubKeyBytes)
}
