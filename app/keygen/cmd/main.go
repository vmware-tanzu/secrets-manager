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
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func main() {
	id := "VSECMKEYGEN"

	//Print the diagnostic information about the environment.
	envVarsToPrint := []string{"APP_VERSION", "VSECM_LOG_LEVEL", "VSECM_KEYGEN_DECRYPT"}
	log.PrintEnvironmentInfo(&id, envVarsToPrint)

	d := env.KeyGenDecrypt()

	if d {
		printDecryptedKeys()
		return
	}

	printGeneratedKeys()
}
