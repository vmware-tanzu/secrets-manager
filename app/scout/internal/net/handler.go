/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package net

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang-jwt/jwt/v4"

	"github.com/vmware-tanzu/secrets-manager/app/scout/internal/filter"
	"github.com/vmware-tanzu/secrets-manager/core/env"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("webhookHandler: Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Default to JWT for now. Eventually we'll support other auth methods.
	if env.ScoutAuthenticationMode() != env.AuthenticationModeNone {
		// Extract the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("webhookHandler: Missing Authorization header")
			http.Error(w, "Missing Authorization header",
				http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and validate the token
		token, err := jwt.Parse(tokenString,
			func(token *jwt.Token) (interface{}, error) {
				// Validate the algorithm
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil,
						fmt.Errorf(
							"unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(jwtSecret), nil
			})

		if err != nil {
			log.Print("webhookHandler: Nil token")
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			log.Println("webhookHandler: Invalid token")
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
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
		log.Println("webhookHandler: Invalid key format")
		http.Error(w, "Invalid key format", http.StatusBadRequest)
		return
	}

	key := values.Get("key")
	path := values.Get("path")

	secretValue, exists := secretsToServe[key]
	if !exists {
		log.Println("webhookHandler: Invalid key")
		http.Error(w, "Invalid key", http.StatusUnauthorized)
		return
	}

	if path == "" {
		log.Println("webhookHandler: Path is required")
		http.Error(w, "Path is required", http.StatusBadRequest)
		return
	}

	var data interface{}
	err = json.Unmarshal([]byte(secretValue), &data)
	if err != nil {
		log.Println("webhookHandler: Error parsing secret data")
		http.Error(w, "Error parsing secret data", http.StatusInternalServerError)
		return
	}

	result, err := filter.ValueFromPath(data, path)
	if err != nil {
		log.Println("webhookHandler: Error getting value from path")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}
