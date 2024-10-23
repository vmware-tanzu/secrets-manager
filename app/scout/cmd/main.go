/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"fmt"
	"github.com/vmware-tanzu/secrets-manager/app/scout/internal/net"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", net.Webhook)

	// Has side effect of initializing jwt token if provided.
	tlsConfig := TlsConfig()

	if env.ScoutTlsEnabled() {
		server := &http.Server{
			Addr:      env.ScoutHttpPort(),
			TLSConfig: tlsConfig,
		}

		fmt.Println("Server is running on :8443 with TLS enabled")
		log.Fatal(server.ListenAndServeTLS("", ""))
		//                                 ^   ^
		// Empty strings because we've already provided the cert and key in TLSConfig.

		return
	}

	fmt.Println("Server is running on :8080")
	server := &http.Server{
		Addr: env.ScoutHttpPort(),
	}

	log.Fatalln(server.ListenAndServe())
}
