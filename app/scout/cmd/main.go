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
	"net/http"

	nets "github.com/vmware-tanzu/secrets-manager/app/scout/internal/net"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func main() {
	id := crypto.Id()

	http.HandleFunc("/webhook", nets.Webhook)

	// Has side effect of initializing jwt token if provided.
	tlsConfig := nets.TlsConfig()

	if env.ScoutTlsEnabled() {
		server := &http.Server{
			Addr:      env.ScoutHttpPort(),
			TLSConfig: tlsConfig,
		}

		log.InfoLn(&id, "Server is running on",
			env.ScoutHttpPort(), "with TLS enabled")

		fmt.Println("Server is running on :8443 with TLS enabled")
		if err := server.ListenAndServeTLS("", ""); err != nil {
			log.InfoLn(&id, "Failed", err.Error())
		}

		return
	}

	log.InfoLn(&id, "Server is running on", env.ScoutHttpPort())
	server := &http.Server{
		Addr: env.ScoutHttpPort(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.InfoLn(&id, "Failed", err.Error())
	}
}
