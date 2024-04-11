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
	"github.com/pkg/errors"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware-tanzu/secrets-manager/app/sentinel/background/initialization"
	"github.com/vmware-tanzu/secrets-manager/app/sentinel/rest"
	"github.com/vmware-tanzu/secrets-manager/core/backoff"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log/rpc"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/probe"
	"github.com/vmware-tanzu/secrets-manager/core/spiffe"
	"github.com/vmware-tanzu/secrets-manager/core/system"
)

func backoffStrategy() backoff.Strategy {
	return backoff.Strategy{
		MaxRetries:  env.BackoffMaxRetries(),
		Delay:       env.BackoffDelay(),
		Exponential: env.BackoffMode() != "linear",
		MaxDuration: env.BackoffMaxDuration(),
	}
}

func main() {
	id := "VSECMSENTINEL"

	go probe.CreateLiveness()
	go rpc.CreateLogServer()

	//Print the diagnostic information about the environment.
	envVarsToPrint := []string{"APP_VERSION", "VSECM_LOG_LEVEL",
		"VSECM_SENTINEL_SECRET_GENERATION_PREFIX"}
	log.PrintEnvironmentInfo(&id, envVarsToPrint)

	log.InfoLn(&id, "Executing the initialization commands (if any)")

	ctx := context.WithValue(context.Background(), "correlationId", &id)

	str := backoffStrategy()
	var src *workloadapi.X509Source = nil
	for {
		log.TraceLn(&id, "Trying to fetch source...")

		err := backoff.Retry("background:fetchSource", func() error {
			log.TraceLn(&id, "Before acquiring source for sentinel")

			s, proceed := spiffe.AcquireSourceForSentinel(ctx)

			if !proceed {
				log.TraceLn(&id, "will retry")
				return errors.New("failed to acquire source")
			}

			log.TraceLn(&id, "acquired source!")
			src = s
			return nil
		}, str)

		if err != nil {
			log.TraceLn(&id, "continue")
			continue
		}

		if src != nil {
			log.TraceLn(&id, "break")
			break
		}
	}

	// Execute the initialization commands (if any)
	// This overloads the functionality of this process.
	// If we end up adding more functionality to this process,
	// we should refactor this and create a new process for the
	// new functionality.
	initialization.RunInitCommands(ctx, src)

	log.InfoLn(&id, "Initialization commands executed successfully")

	if env.SentinelEnableOIDCResourceServer() {
		go rest.RunRestServer()
	}
	// Run on the main thread to wait forever.
	system.KeepAlive()
}
