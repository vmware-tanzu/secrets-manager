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

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/secrets", engine.HandleSecrets)
	log.Println("VSecM Server started at :8085")
	log.Fatal(http.ListenAndServe(":8085", mux))
}
