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
	"encoding/json"
	"fmt"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/secret/collection"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"time"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/bootstrap"
	server "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/engine"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/io"
	"github.com/vmware-tanzu/secrets-manager/core/constants/env"
	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	env2 "github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/probe"
)

type SafeConfig struct {
	Config struct {
		BackingStore   string `json:"backingStore"`
		DataSourceName string `json:"dataSourceName"`
	} `json:"config"`
}

func pollForConfig(ctx context.Context, id string) (*SafeConfig, error) {
	for {
		log.InfoLn(&id, "Polling for VSecM Safe internal configuration")
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			vSecMSafeInternalConfig, err := collection.ReadSecret(id, "vsecm-safe")
			if err != nil {
				log.InfoLn(&id, "Failed to load VSecM Safe internal configuration", err.Error())
			} else if vSecMSafeInternalConfig != nil && len(vSecMSafeInternalConfig.Values) > 0 {
				var safeConfig SafeConfig
				err := json.Unmarshal([]byte(vSecMSafeInternalConfig.Values[0]), &safeConfig)
				if err != nil {
					log.InfoLn(&id, "Failed to parse VSecM Safe internal configuration", err.Error())
				} else {
					return &safeConfig, nil
				}
			}
			time.Sleep(5 * time.Second)
		}
	}
}

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

	go func() {
		log.InfoLn(&id, "Waiting for VSecM Safe internal configuration...")
		safeConfig, err := pollForConfig(ctx, id)
		if err != nil {
			log.FatalLn(&id, "Failed to retrieve VSecM Safe internal configuration", err.Error())
		}

		log.InfoLn(&id, "VSecM Safe internal configuration loaded")
		fmt.Printf("Backing Store: %s\n", safeConfig.Config.BackingStore)
		fmt.Printf("Data Source Name: %s\n", safeConfig.Config.DataSourceName)
	}()

	// TODO: this should be part of initialization counter too.
	if env2.BackingStoreForSafe() == data.Postgres {
		dataSourceName := env2.PostgresDataSourceNameForSafe() // You'll need to implement this function
		err := io.InitDB(dataSourceName)
		if err != nil {
			log.FatalLn(&id, "Failed to initialize database:", err)
		}
	}

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
