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
	e "github.com/vmware-tanzu/secrets-manager/core/constants/env"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/lib/system"
	"github.com/vmware-tanzu/secrets-manager/sdk/startup"
)

func main() {
	id := crypto.Id()

	//Print the diagnostic information about the environment.
	log.PrintEnvironmentInfo(&id, []string{
		string(e.AppVersion),
		string(e.VSecMLogLevel),
		string(e.VSecMSafeEndpointUrl),
	})

	log.InfoLn(&id, "Starting VSecM Init Container")

	// Wait for a specified duration before exiting the init container.
	// This can be useful when you want things to reconcile before
	// starting the main container.
	go startup.Watch(env.WaitBeforeExitForInitContainer())

	// Block the process from exiting, but also be graceful and honor the
	// termination signals that may come from the orchestrator.
	system.KeepAlive()
}
