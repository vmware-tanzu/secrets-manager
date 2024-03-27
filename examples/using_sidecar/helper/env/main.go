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
	"os"
)

func sidecarSecretsPath() string {
	p := os.Getenv("VSECM_SIDECAR_SECRETS_PATH")
	if p == "" {
		p = "/opt/vsecm/secrets.json"
	}
	return p
}

func main() {
	dat, err := os.ReadFile(sidecarSecretsPath())
	if err != nil {
		fmt.Print("ERR_READ_SECRET")
	} else {
		fmt.Print(string(dat))
	}
}
