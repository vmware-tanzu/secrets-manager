package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Get the 'key' query parameter
	encodedKey := r.URL.Query().Get("key")

	// Unescape the key parameter
	decodedKey, err := url.QueryUnescape(encodedKey)
	if err != nil {
		http.Error(w, "Failed to decode key parameter", http.StatusBadRequest)
		return
	}

	// Parse the decoded key as a query string
	values, err := url.ParseQuery(decodedKey)
	if err != nil {
		http.Error(w, "Invalid key format", http.StatusBadRequest)
		return
	}

	key := values.Get("key")
	path := values.Get("path")

	secretValue, exists := secretsToServe[key]
	if !exists {
		http.Error(w, "Invalid key", http.StatusUnauthorized)
		return
	}

	if path == "" {
		http.Error(w, "Path is required", http.StatusBadRequest)
		return
	}

	var data interface{}
	err = json.Unmarshal([]byte(secretValue), &data)
	if err != nil {
		http.Error(w, "Error parsing secret data", http.StatusInternalServerError)
		return
	}

	result, err := getValueFromPath(data, path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}
