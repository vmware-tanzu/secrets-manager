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
	"github.com/vmware-tanzu/secrets-manager/core/constants/env"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/lib/system"
	"github.com/vmware-tanzu/secrets-manager/sdk/sentry"
)

func main() {
	id := crypto.Id()
	log.InfoLn(&id, "Starting VSecM Sidecar")

	//Print the diagnostic information about the environment.
	envVarsToPrint := []string{
		string(env.AppVersion),
		string(env.VSecMLogLevel),
	}
	log.PrintEnvironmentInfo(&id, envVarsToPrint)

	// Periodically update secret values:
	go sentry.Watch()

	// Keep the main routine alive:
	system.KeepAlive()
}
