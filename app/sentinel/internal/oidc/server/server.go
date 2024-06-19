/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package server

import (
	"log"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/oidc/engine"
)

// Serve initializes and starts an HTTP server for VSecM Sentinel.
//
// This function sets up an HTTP server with a multiplexer for handling requests.
// It specifically registers the "/secrets" endpoint with the HandleSecrets
// function from the engine package.
//
// Example usage:
//
//	Serve() // This will start the server and listen for incoming requests.
//
// Details:
//   - mux: An HTTP request multiplexer that routes incoming requests to the
//     registered handler functions.
//   - HandleSecrets: A function from the engine package that processes
//     requests to the "/secrets" endpoint.
func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/secrets", engine.HandleSecrets)

	log.Println("VSecM Server started at :8085")
	log.Fatal(http.ListenAndServe(":8085", mux))
}
