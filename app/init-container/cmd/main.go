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
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"github.com/vmware-tanzu/secrets-manager/core/system"
	"github.com/vmware-tanzu/secrets-manager/sdk/startup"
)

func main() {
	id := "AEGIICNT"

	log.InfoLn(&id, "Starting VSecM Init Container")
	go startup.Watch()

	//Print the diagnostic information about the environment.
	envVarsToPrint := []string{"APP_VERSION", "VSECM_LOG_LEVEL", "VSECM_SAFE_ENDPOINT_URL"}
	log.PrintEnvironmentInfo(&id, envVarsToPrint)

	// Block the process from exiting, but also be graceful and honor the
	// termination signals that may come from the orchestrator.
	system.KeepAlive()
}
