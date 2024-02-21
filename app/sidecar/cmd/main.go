/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/system"
	"github.com/vmware-tanzu/secrets-manager/sdk/sentry"
)

func main() {
	id := "AEGSSDCR"
	log.InfoLn(&id, "Starting VSecM Sidecar")

	//Print the diagnostic information about the environment.
	envVarsToPrint := []string{"APP_VERSION", "VSECM_LOG_LEVEL"}
	log.PrintEnvironmentInfo(&id, envVarsToPrint)

	go sentry.Watch()
	// Keep the main routine alive:
	system.KeepAlive()
}
