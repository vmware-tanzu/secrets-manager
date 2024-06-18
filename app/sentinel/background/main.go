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
	"context"

	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/initialization"
	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/oidc/server"
	"github.com/vmware-tanzu/secrets-manager/core/constants"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log/rpc"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/probe"
	"github.com/vmware-tanzu/secrets-manager/core/system"
)

func main() {
	id := crypto.Id()

	//Print the diagnostic information about the environment.
	log.PrintEnvironmentInfo(&id, []string{
		string(constants.AppVersion),
		string(constants.VSecMLogLevel),
	})

	go probe.CreateLiveness()
	go rpc.CreateLogServer()

	log.InfoLn(&id, "Executing the initialization commands (if any)")

	ctx := context.WithValue(context.Background(), "correlationId", &id)

	log.TraceLn(&id, "before RunInitCommands")

	// Execute the initialization commands (if any)
	// This overloads the functionality of this process.
	// If we end up adding more functionality to this process,
	// we should refactor this and create a new process for the
	// new functionality.
	initialization.RunInitCommands(ctx)

	log.InfoLn(&id, "Initialization commands executed successfully")

	if env.SentinelEnableOIDCResourceServer() {
		go server.Serve()
	}

	// Run on the main thread to wait forever.
	system.KeepAlive()
}
