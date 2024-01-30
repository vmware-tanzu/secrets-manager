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
	"context"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/bootstrap"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"github.com/vmware-tanzu/secrets-manager/core/probe"
	"github.com/vmware-tanzu/secrets-manager/core/util"
)

func main() {
	id := "VSECMSAFE"

	bootstrap.ValidateEnvironment()

	log.InfoLn(&id, "Acquiring identity…")

	// Channel to notify when the bootstrap timeout has been reached.
	timedOut := make(chan bool, 1)

	// These channels must complete in a timely manner, otherwise
	// the timeOut will be fired and will crash the app.
	acquiredSvid := make(chan bool, 1)
	updatedSecret := make(chan bool, 1)
	serverStarted := make(chan bool, 1)

	// Monitor the progress of acquiring an identity, updating the age key,
	// and starting the server. If the timeout occurs before all three events
	// happen, the function logs a fatal message and the process crashes.
	go bootstrap.Monitor(&id,
		bootstrap.ChannelsToMonitor{
			AcquiredSvid:  acquiredSvid,
			UpdatedSecret: updatedSecret,
			ServerStarted: serverStarted,
		}, timedOut,
	)

	//Print the diagnostic information about the environment.
	envVarsToPrint := []string{"APP_VERSION", "VSECM_LOG_LEVEL", "VSECM_SAFE_FIPS_COMPLIANT", "VSECM_SAFE_SPIFFEID_PREFIX", "VSECM_SAFE_TLS_PORT"}
	go util.PrintEnvironmentInfo(&id, envVarsToPrint)

	// Time out if things take too long.
	go bootstrap.NotifyTimeout(timedOut)

	// Create initial cryptographic seeds off-cycle.
	go bootstrap.CreateCryptoKey(&id, updatedSecret)

	// App is alive; however, not yet ready to accept connections.
	go probe.CreateLiveness()

	ctx, cancel := context.WithCancel(
		context.WithValue(context.Background(), "correlationId", &id),
	)

	defer cancel()

	source := bootstrap.AcquireSource(ctx, acquiredSvid)
	defer func() {
		// Close the source after the server (1) is done serving, likely
		// when the app is shutting down due to an eviction or a panic.
		if err := source.Close(); err != nil {
			log.InfoLn(&id, "Problem closing SVID Bundle source: %v\n", err)
		}
	}()

	// (1)
	if err := server.Serve(source, serverStarted); err != nil {
		log.FatalLn(&id, "failed to serve", err.Error())
	}
}
