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
	"github.com/vmware-tanzu/secrets-manager/core/probe"
	"github.com/vmware-tanzu/secrets-manager/core/system"
	"github.com/vmware-tanzu/secrets-manager/core/util"
)

func main() {
	id := "VSECMSENTINEL"

	go probe.CreateLiveness()

	//Print the diagnostic information about the environment.
	envVarsToPrint := []string{"APP_VERSION", "VSECM_LOG_LEVEL", "VSECM_SENTINEL_SECRET_GENERATION_PREFIX"}
	go util.PrintEnvironmentInfo(&id, envVarsToPrint)

	// Run on the main thread to wait forever.
	system.KeepAlive()
}
