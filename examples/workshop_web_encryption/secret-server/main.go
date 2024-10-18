package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/nacl/box"
	"io/ioutil"
	"net/http"
)

type VerificationData struct {
	PublicKey string `json:"publicKey"`
	Nonce     string `json:"nonce"`
	Signature string `json:"signature"`
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("verifyhandler")

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		fmt.Println("options")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data VerificationData

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Cannot read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// https://github.com/vmware-tanzu/secrets-manager

	// Parse JSON data
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Decode Base64 encoded data
	publicKeyBytes, err := base64.StdEncoding.DecodeString(data.PublicKey)
	if err != nil || len(publicKeyBytes) != ed25519.PublicKeySize {
		http.Error(w, "Invalid public key", http.StatusBadRequest)
		return
	}

	nonceBytes, err := base64.StdEncoding.DecodeString(data.Nonce)
	if err != nil {
		http.Error(w, "Invalid nonce", http.StatusBadRequest)
		return
	}

	signatureBytes, err := base64.StdEncoding.DecodeString(data.Signature)
	if err != nil || len(signatureBytes) != ed25519.SignatureSize {
		http.Error(w, "Invalid signature", http.StatusBadRequest)
		return
	}

	// Verify the signature
	valid := ed25519.Verify(publicKeyBytes, nonceBytes, signatureBytes)
	if valid {
		fmt.Fprintln(w, "Signature is valid")
	} else {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
	}
}

var (
	serverSignPublicKey  ed25519.PublicKey
	serverSignPrivateKey ed25519.PrivateKey
	serverBoxPublicKey   [32]byte
	serverBoxPrivateKey  [32]byte

	aesKey []byte // Store the AES key generated in getEncryptedDataHandler

	storedMessage []byte
)

func nonceHandler(w http.ResponseWriter, r *http.Request) {
	// Allow CORS for client at localhost:8081
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Generate a nonce (random 32-byte value)
	nonce := make([]byte, 32)
	_, err := rand.Read(nonce)
	if err != nil {
		http.Error(w, "Failed to generate nonce", http.StatusInternalServerError)
		return
	}

	// Sign the nonce with the server's private key
	signature := ed25519.Sign(serverSignPrivateKey, nonce)

	// Encode the data in Base64
	nonceBase64 := base64.StdEncoding.EncodeToString(nonce)
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)
	publicKeyBase64 := base64.StdEncoding.EncodeToString(serverSignPublicKey)

	// Prepare the JSON response
	response := fmt.Sprintf(`{
        "nonce": "%s",
        "signature": "%s",
        "publicKey": "%s"
    }`, nonceBase64, signatureBase64, publicKeyBase64)

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, response)
}

func init() {
	storedMessage = []byte("This is a secret message")

	//// Generate the server's key pair
	//var err error
	//serverSignPublicKey, serverSignPrivateKey, err = ed25519.GenerateKey(rand.Reader)
	//if err != nil {
	//	panic("Failed to generate server key pair")
	//}

	// Generate Ed25519 key pair for signing
	var err error
	serverSignPublicKey, serverSignPrivateKey, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic("Failed to generate server Ed25519 key pair")
	}

	// Generate Curve25519 key pair for encryption
	var privateKey [32]byte
	_, err = rand.Read(privateKey[:])
	if err != nil {
		panic("Failed to generate server Curve25519 private key")
	}
	serverBoxPrivateKey = privateKey

	var publicKey [32]byte
	curve25519.ScalarBaseMult(&publicKey, &privateKey)
	serverBoxPublicKey = publicKey
}

type ClientKeys struct {
	SignPublicKey ed25519.PublicKey
	BoxPublicKey  [32]byte
}

var clientKeys ClientKeys

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read and parse the request body
	var data struct {
		SignPublicKey string `json:"signPublicKey"`
		BoxPublicKey  string `json:"boxPublicKey"`
		Signature     string `json:"signature"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Decode Base64 encoded keys
	signPublicKeyBytes, err := base64.StdEncoding.DecodeString(data.SignPublicKey)
	if err != nil {
		http.Error(w, "Invalid sign public key", http.StatusBadRequest)
		return
	}

	boxPublicKeyBytes, err := base64.StdEncoding.DecodeString(data.BoxPublicKey)
	if err != nil {
		http.Error(w, "Invalid box public key", http.StatusBadRequest)
		return
	}

	// Store the client's public keys
	clientKeys.SignPublicKey = signPublicKeyBytes
	copy(clientKeys.BoxPublicKey[:], boxPublicKeyBytes)

	// Reconstruct data to verify
	dataToVerify := data.BoxPublicKey

	signatureBytes, err := base64.StdEncoding.DecodeString(data.Signature)

	// Verify the signature using client's signing public key
	isValid := ed25519.Verify(clientKeys.SignPublicKey, []byte(dataToVerify), signatureBytes)
	if !isValid {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	fmt.Println("Client signature verified")

	fmt.Fprintln(w, "Client public keys registered successfully")
}

func getEncryptedDataHandler(w http.ResponseWriter, r *http.Request) {
	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Ensure client's public key is registered
	if clientKeys.BoxPublicKey == [32]byte{} {
		http.Error(w, "Client public key not registered", http.StatusBadRequest)
		return
	}

	// Generate AES key (32 bytes)
	aesKey = make([]byte, 32)
	_, err := rand.Read(aesKey)
	if err != nil {
		http.Error(w, "Failed to generate AES key", http.StatusInternalServerError)
		return
	}

	// Encrypt AES key using client's Curve25519 public key
	var nonceAESKey [24]byte
	_, err = rand.Read(nonceAESKey[:])
	if err != nil {
		http.Error(w, "Failed to generate nonce", http.StatusInternalServerError)
		return
	}

	encryptedAESKey := box.Seal(nil, aesKey, &nonceAESKey, &clientKeys.BoxPublicKey, &serverBoxPrivateKey)

	// Encrypt message using AES key
	message := storedMessage
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		http.Error(w, "Failed to create cipher", http.StatusInternalServerError)
		return
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(w, "Failed to create GCM", http.StatusInternalServerError)
		return
	}

	nonceMessage := make([]byte, aesGCM.NonceSize())
	_, err = rand.Read(nonceMessage)
	if err != nil {
		http.Error(w, "Failed to generate nonce", http.StatusInternalServerError)
		return
	}

	encryptedMessage := aesGCM.Seal(nil, nonceMessage, message, nil)

	// Encode data to Base64
	encryptedAESKeyBase64 := base64.StdEncoding.EncodeToString(encryptedAESKey)
	nonceAESKeyBase64 := base64.StdEncoding.EncodeToString(nonceAESKey[:])
	encryptedMessageBase64 := base64.StdEncoding.EncodeToString(encryptedMessage)
	nonceMessageBase64 := base64.StdEncoding.EncodeToString(nonceMessage)
	serverBoxPublicKeyBase64 := base64.StdEncoding.EncodeToString(serverBoxPublicKey[:])

	// Prepare the data to be signed
	// dataToSign := encryptedAESKeyBase64 + nonceAESKeyBase64 + encryptedMessageBase64 + nonceMessageBase64 //+ serverBoxPublicKeyBase64
	dataToSign := nonceAESKeyBase64 + encryptedMessageBase64 + nonceMessageBase64 //+ serverBoxPublicKeyBase64
	// dataToSign := nonceAESKeyBase64 // + encryptedMessageBase64 + nonceMessageBase64 + serverBoxPublicKeyBase64

	// Sign the data using the server's signing private key
	signature := ed25519.Sign(serverSignPrivateKey, []byte(dataToSign))
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	fmt.Println("################")
	fmt.Println("dataToSign:", dataToSign)
	fmt.Println("################")

	// Include the signature and server's signing public key in the response
	response := map[string]string{
		"encryptedAESKey":     encryptedAESKeyBase64,
		"nonceAESKey":         nonceAESKeyBase64,
		"encryptedMessage":    encryptedMessageBase64,
		"nonceMessage":        nonceMessageBase64,
		"serverBoxPublicKey":  serverBoxPublicKeyBase64,
		"signature":           signatureBase64,
		"serverSignPublicKey": base64.StdEncoding.EncodeToString(serverSignPublicKey),
	}

	fmt.Println("sspk", base64.StdEncoding.EncodeToString(serverSignPublicKey))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	//// Prepare JSON response
	//response := map[string]string{
	//	"encryptedAESKey":    encryptedAESKeyBase64,
	//	"nonceAESKey":        nonceAESKeyBase64,
	//	"encryptedMessage":   encryptedMessageBase64,
	//	"nonceMessage":       nonceMessageBase64,
	//	"serverBoxPublicKey": serverBoxPublicKeyBase64,
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(response)
}

func receiveEncryptedDataHandler(w http.ResponseWriter, r *http.Request) {
	// Allow CORS
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read and parse the request body
	var data struct {
		NonceMessageToServer     string `json:"nonceMessageToServer"`
		EncryptedMessageToServer string `json:"encryptedMessageToServer"`
		Signature                string `json:"signature"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Decode Base64 encoded data
	nonceMessageToServerBytes, err := base64.StdEncoding.DecodeString(data.NonceMessageToServer)
	if err != nil {
		http.Error(w, "Invalid nonce", http.StatusBadRequest)
		return
	}

	encryptedMessageToServerBytes, err := base64.StdEncoding.DecodeString(data.EncryptedMessageToServer)
	if err != nil {
		http.Error(w, "Invalid encrypted message", http.StatusBadRequest)
		return
	}

	signatureBytes, err := base64.StdEncoding.DecodeString(data.Signature)
	if err != nil {
		http.Error(w, "Invalid signature", http.StatusBadRequest)
		return
	}

	// Reconstruct data to verify
	dataToVerify := data.NonceMessageToServer + data.EncryptedMessageToServer

	// Verify the signature using client's signing public key
	isValid := ed25519.Verify(clientKeys.SignPublicKey, []byte(dataToVerify), signatureBytes)
	if !isValid {
		http.Error(w, "Invalid signature", http.StatusUnauthorized)
		return
	}

	// Decrypt the message using AES-GCM
	if aesKey == nil {
		http.Error(w, "AES key not found", http.StatusInternalServerError)
		return
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		http.Error(w, "Failed to create cipher", http.StatusInternalServerError)
		return
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		http.Error(w, "Failed to create GCM", http.StatusInternalServerError)
		return
	}

	plaintext, err := aesGCM.Open(nil, nonceMessageToServerBytes, encryptedMessageToServerBytes, nil)
	if err != nil {
		http.Error(w, "Failed to decrypt message", http.StatusBadRequest)
		return
	}

	storedMessage = plaintext

	// Process the decrypted message
	fmt.Println("Decrypted message from client:", string(plaintext))

	w.Write([]byte("Message received and decrypted successfully"))
}

func main() {
	http.HandleFunc("/verify", verifyHandler)
	http.HandleFunc("/getNonce", nonceHandler)

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/getEncryptedData", getEncryptedDataHandler)
	http.HandleFunc("/receiveEncryptedData", receiveEncryptedDataHandler)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
