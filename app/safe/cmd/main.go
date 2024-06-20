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

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/bootstrap"
	server "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/engine"
	"github.com/vmware-tanzu/secrets-manager/core/constants/env"
	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/probe"
)

func main() {
	id := crypto.Id()

	//Print the diagnostic information about the environment.
	log.PrintEnvironmentInfo(&id, []string{
		string(env.AppVersion),
		string(env.VSecMLogLevel),
	})

	ctx, cancel := context.WithCancel(
		context.WithValue(context.Background(), key.CorrelationId, &id),
	)
	defer cancel()

	log.InfoLn(&id, "Acquiring identity...")

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

	// Time out if things take too long.
	go bootstrap.NotifyTimeout(timedOut)

	// Create initial cryptographic seeds off-cycle.
	go bootstrap.CreateRootKey(&id, updatedSecret)

	// App is alive; however, not yet ready to accept connections.
	<-probe.CreateLiveness()

	log.InfoLn(&id, "before acquiring source...")
	source := bootstrap.AcquireSource(ctx, acquiredSvid)
	defer func(s *workloadapi.X509Source) {
		if s == nil {
			return
		}

		// Close the source after the server (1) is done serving, likely
		// when the app is shutting down due to an eviction or a panic.
		if err := s.Close(); err != nil {
			log.InfoLn(&id, "Problem closing SVID Bundle source: %v\n", err)
		}
	}(source)

	// (1)
	if err := server.Serve(source, serverStarted); err != nil {
		log.FatalLn(&id, "failed to serve", err.Error())
	}
}
