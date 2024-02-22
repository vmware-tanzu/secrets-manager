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
	"context"

	"github.com/vmware-tanzu/secrets-manager/app/sentinel/busywait/initialization"
	"github.com/vmware-tanzu/secrets-manager/core/log/rpc"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/probe"
	"github.com/vmware-tanzu/secrets-manager/core/system"
)

func main() {
	id := "VSECMSENTINEL"

	go probe.CreateLiveness()
	go rpc.CreateLogServer()

	//Print the diagnostic information about the environment.
	envVarsToPrint := []string{"APP_VERSION", "VSECM_LOG_LEVEL",
		"VSECM_SENTINEL_SECRET_GENERATION_PREFIX"}
	log.PrintEnvironmentInfo(&id, envVarsToPrint)

	log.InfoLn(&id, "Executing the initialization commands (if any)")

	// Execute the initialization commands (if any)
	// This overloads the functionality of this process.
	// If we end up adding more functionality to this process,
	// we should refactor this and create a new process for the
	// new functionality.
	initialization.RunInitCommands(
		context.WithValue(context.Background(), "correlationId", &id),
	)

	log.InfoLn(&id, "Initialization commands executed successfully")

	// Run on the main thread to wait forever.
	system.KeepAlive()
}
