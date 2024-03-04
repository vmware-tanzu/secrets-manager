/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleSecrets(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}

	var req SecretRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	HandleCommandSecrets(w, r, &req)
}

func ok(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "OK")
	if err != nil {
		log.Printf("probe response failure: %s", err.Error())
		return
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", ok)
	mux.HandleFunc("/secrets", HandleSecrets)

	log.Println("Vsecm Rest Server Started on :8085")
	log.Fatal(http.ListenAndServe(":8085", mux))
}
