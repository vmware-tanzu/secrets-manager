package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Log detailed request information
	fmt.Println("--- Incoming Request Details ---")
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("URL: %s\n", r.URL.String())
	fmt.Printf("Protocol: %s\n", r.Proto)
	fmt.Println("Headers:")
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Printf("  %v: %v\n", name, h)
		}
	}

	// Get the 'key' query parameter
	encodedKey := r.URL.Query().Get("key")
	fmt.Printf("Raw 'key' parameter: %s\n", encodedKey)

	// Unescape the key parameter
	decodedKey, err := url.QueryUnescape(encodedKey)
	if err != nil {
		http.Error(w, "Failed to decode key parameter", http.StatusBadRequest)
		return
	}
	fmt.Printf("Decoded 'key' parameter: %s\n", decodedKey)

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

	fmt.Println("path", path)

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
