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

	// Create the Initializer with all dependencies
	initializer := initialization.NewInitializer(
		&initialization.OSFileOpener{},
		&initialization.EnvConfigReader{},
		&initialization.StandardLogger{},
		&initialization.SafeClient{},
		&initialization.SpiffeClient{},
	)

	// Execute the initialization commands (if any)
	initializer.RunInitCommands(ctx)

	log.InfoLn(&id, "Initialization commands executed successfully")

	if env.SentinelEnableOIDCResourceServer() {
		go server.Serve()
	}

	// Run on the main thread to wait forever.
	system.KeepAlive()
}
