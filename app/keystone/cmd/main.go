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
	sys "log"
	"os"

	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/system"
)

func main() {
	id := crypto.Id()

	//Print the diagnostic information about the environment.
	envVarsToPrint := []string{"APP_VERSION"}
	log.PrintEnvironmentInfo(&id, envVarsToPrint)

	sys.Println(
		"VSecM Keystone",
		fmt.Sprintf("v%s", os.Getenv("APP_VERSION")),
	)

	// Run on the main thread to wait forever.
	system.KeepAlive()
}
