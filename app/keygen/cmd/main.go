/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/util"
)

func main() {
	id := "VSECMKEYGEN"

	//Print the diagnostic information about the environment.
	envVarsToPrint := []string{"APP_VERSION", "VSECM_LOG_LEVEL", "VSECM_KEYGEN_DECRYPT"}
	go util.PrintEnvironmentInfo(&id, envVarsToPrint)

	d := env.KeyGenDecrypt()

	if d {
		printDecryptedKeys()
		return
	}

	printGeneratedKeys()
}
