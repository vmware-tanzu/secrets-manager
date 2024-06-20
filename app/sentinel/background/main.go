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
	e "github.com/vmware-tanzu/secrets-manager/core/constants/env"
	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log/rpc"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/probe"
	"github.com/vmware-tanzu/secrets-manager/lib/system"
)

func main() {
	id := crypto.Id()

	//Print the diagnostic information about the environment.
	log.PrintEnvironmentInfo(&id, []string{
		string(e.AppVersion),
		string(e.VSecMLogLevel),
	})

	<-probe.CreateLiveness()
	go rpc.CreateLogServer()

	log.InfoLn(&id, "Executing the initialization commands (if any)")

	ctx := context.WithValue(context.Background(),
		key.CorrelationId, &id)

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
