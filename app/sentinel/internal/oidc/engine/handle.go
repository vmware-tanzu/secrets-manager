/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package engine

import (
	"encoding/json"
	"net/http"
)

// HandleSecrets processes incoming HTTP requests related to secrets management.
// This function is specifically designed to handle POST requests. It decodes
// the JSON body to a SecretRequest and delegates the actual business logic to
// HandleCommandSecrets.
//
// Parameters:
//   - w: the http.ResponseWriter used to write the HTTP response.
//   - r: the *http.Request containing all the request data including headers,
//     query parameters, and the body.
//
// Usage:
//
//	http.HandleFunc("/api/secrets", secrets.HandleSecrets)
//
// Errors:
//   - Returns HTTP 405 if the method is not POST.
//   - Returns HTTP 400 if the request body cannot be decoded into
//     SecretRequest.
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
