package main

import (
	"encoding/json"
	"fmt"
	"github.com/vmware-tanzu/secrets-manager/sdk/sentry"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	fmt.Println("before fetching")
	sfr, err := sentry.Fetch()
	if err != nil {
		fmt.Println("error", err.Error())
		fmt.Println("will sleep for 1 hour")
		time.Sleep(time.Second * 3600)
	}
	fmt.Println("after fetching")

	// TODO: JWT token validation

	if err == nil {
		fmt.Println("data", sfr.Data)
	}

	http.HandleFunc("/webhook", webhookHandler)
	fmt.Println("Server is running on :8443")

	log.Fatal(http.ListenAndServe(":8443", nil))

	// log.Fatal(http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil))
}

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

	if key != "coca-cola.cluster-001" {
		http.Error(w, "Invalid key", http.StatusUnauthorized)
		return
	}

	fmt.Println("path", path)

	if path == "" {
		http.Error(w, "Path is required", http.StatusBadRequest)
		return
	}

	// This is now our arbitrary JSON structure
	jsonData := []byte(`                                                                                                                                                                                                                                     
    {                                                                                                                                                                                                                                                        
        "namespaces": {                                                                                                                                                                                                                                      
            "cokeSystem": {                                                                                                                                                                                                                                  
                "secrets": {                                                                                                                                                                                                                                 
                    "adminCredentials": {                                                                                                                                                                                                                    
                        "type": "k8s",                                                                                                                                                                                                                       
                        "value": "super-secret-secret",                                                                                                                                                                                                      
                        "metadata": {                                                                                                                                                                                                                        
                            "labels": {                                                                                                                                                                                                                      
                                "managedBy": "coke-system"                                                                                                                                                                                                   
                            },                                                                                                                                                                                                                               
                            "annotations": {                                                                                                                                                                                                                 
                                "injectSidecar": "true"                                                                                                                                                                                                      
                            },                                                                                                                                                                                                                               
                            "creationTimestamp": "2024-01-01",                                                                                                                                                                                               
                            "lastUpdated": "2024-01-01"                                                                                                                                                                                                      
                        },                                                                                                                                                                                                                                   
                        "expires": "2024-01-01",                                                                                                                                                                                                             
                        "notBefore": "2024-01-01"                                                                                                                                                                                                            
                    }                                                                                                                                                                                                                                        
                }                                                                                                                                                                                                                                            
            }                                                                                                                                                                                                                                                
        }                                                                                                                                                                                                                                                    
    }                                                                                                                                                                                                                                                        
    `)

	var data interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		http.Error(w, "Error parsing JSON data", http.StatusInternalServerError)
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

func getValueFromPath(data interface{}, path string) (interface{}, error) {
	parts := strings.Split(path, ".")

	var current interface{} = data
	for _, part := range parts {
		switch v := current.(type) {
		case map[string]interface{}:
			if val, ok := v[part]; ok {
				current = val
			} else {
				return nil, fmt.Errorf("key not found: %s", part)
			}
		case []interface{}:
			return nil, fmt.Errorf("arrays are not supported in path queries")
		default:
			return nil, fmt.Errorf("cannot navigate further from %v", current)
		}
	}

	return current, nil
}
