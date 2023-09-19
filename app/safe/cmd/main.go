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
)

func main() {
	id := "VSECMSAFE"

	bootstrap.ValidateEnvironment()

	log.InfoLn(&id, "Acquiring identity…")

	timedOut := make(chan bool, 1)
	// These channels mus complete in a timely manner, otherwise
	// the timeOut will be fired and will crash the app.
	acquiredSvid := make(chan bool, 1)
	updatedSecret := make(chan bool, 1)
	serverStarted := make(chan bool, 1)

	go bootstrap.NotifyTimeout(timedOut)
	go bootstrap.CreateCryptoKey(&id, updatedSecret)
	go bootstrap.Monitor(&id, acquiredSvid, updatedSecret, serverStarted, timedOut)

	// App is alive; however, not yet ready to accept connections.
	go probe.CreateLiveness()

	ctx, cancel := context.WithCancel(
		context.WithValue(context.Background(), "correlationId", &id),
	)

	defer cancel()

	source := bootstrap.AcquireSource(ctx, acquiredSvid)
	defer func() {
		if err := source.Close(); err != nil {
			log.InfoLn(&id, "Problem closing SVID Bundle source: %v\n", err)
		}
	}()

	if err := server.Serve(source, serverStarted); err != nil {
		log.FatalLn(&id, "failed to serve", err.Error())
	}
}
