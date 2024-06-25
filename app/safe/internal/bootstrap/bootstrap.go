/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package bootstrap

import (
	"context"
	"os"
	"reflect"
	"time"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/core/constants/val"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

// NotifyTimeout waits for the duration specified by
// env.BootstrapTimeoutForSafe() and then sends a 'true' value to the provided
// 'timedOut' channel. This function can be used to notify other parts of the
// application when a specific timeout has been reached.
func NotifyTimeout(timedOut chan<- bool) {
	time.Sleep(env.BootstrapTimeoutForSafe())
	timedOut <- true
}

// ChannelsToMonitor holds channels to monitor various asynchronous signals.
type ChannelsToMonitor struct {
	AcquiredSvid  <-chan bool
	UpdatedSecret <-chan bool
	ServerStarted <-chan bool
}

// Size returns the number of fields in the ChannelsToMonitor struct.
// This is useful to programmatically check the count of different channels
// that need to be monitored.
func (c ChannelsToMonitor) Size() int {
	t := reflect.TypeOf(c)
	return t.NumField()
}

// Monitor listens to various channels to track the progress of acquiring an
// identity, updating the age key, and starting the server. It takes a
// correlationId for logging purposes and four channels: acquiredSvid,
// updatedSecret, serverStarted, and timedOut. When all three of the first
// events (acquiring identity, updating age key, and starting the server) have
// occurred, the function initializes the state and creates a readiness probe.
// If a timeout occurs before all three events happen, the function logs a
// fatal message.
func Monitor(
	correlationId *string,
	channels ChannelsToMonitor,
	timedOut <-chan bool,
) {
	// Number of channels ins `channels` to wait for.
	counter := channels.Size()

	for {
		if counter == 0 {
			break
		}

		select {
		// Acquired SVID for this workload from the SPIRE Agent via
		// workload API:
		case <-channels.AcquiredSvid:
			log.AuditLn(correlationId, "Acquired identity.")
			counter--
			log.InfoLn(
				correlationId,
				"remaining operations before ready:", counter)
			if counter == 0 {
				completeInitialization(correlationId)
			}
		// Updated the root key:
		case <-channels.UpdatedSecret:
			log.DebugLn(correlationId, "Updated age key.")
			counter--
			log.InfoLn(
				correlationId,
				"remaining operations before ready:", counter)
			if counter == 0 {
				completeInitialization(correlationId)
			}
		// VSecM Safe REST API is ready to serve:
		case <-channels.ServerStarted:
			log.DebugLn(correlationId, "Server ready.")
			counter--
			log.InfoLn(
				correlationId,
				"remaining operations before ready:", counter)
			if counter == 0 {
				completeInitialization(correlationId)
			}
		// Things didn't start in a timely manner:
		case <-timedOut:
			log.FatalLn(
				correlationId,
				"Failed to acquire"+
					" an identity in a timely manner.")
		}
	}
}

// AcquireSource establishes a connection to the workload API, fetches the
// X.509 bundle, and returns an X509Source. It takes a context and a channel
// acquiredSvid to signal when the SVID has been acquired. If there are any
// errors during the process, the function logs a fatal message and exits.
func AcquireSource(
	ctx context.Context, acquiredSvid chan<- bool,
) *workloadapi.X509Source {
	cid := ctx.Value(key.CorrelationId).(*string)

	log.InfoLn(cid, "Acquiring source...")

	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(
			workloadapi.WithAddr(env.SpiffeSocketUrl()),
		),
	)

	if err != nil {
		log.FatalLn(cid, "Unable to fetch X.509 Bundle", err.Error())
		return nil
	}

	if source == nil {
		log.FatalLn(cid, "Could not find source")
		return nil
	}

	//goland:noinspection ALL
	svid, err := source.GetX509SVID()
	if err != nil {
		log.FatalLn(cid, "Unable to get X.509 SVID from source bundle", err.Error())
		return nil
	}

	if svid == nil {
		log.FatalLn(cid, "Could not find SVID in source bundle")
		return nil
	}

	svidId := svid.ID
	if !validation.IsSafe(svidId.String()) {
		log.FatalLn(
			cid,
			"SpiffeId check: Safe:bootstrap: I don't know you, and it's crazy:",
			svidId.String(),
		)
		return nil
	}

	log.TraceLn(cid, "Sending: Acquired SVID", len(acquiredSvid))

	acquiredSvid <- true

	log.TraceLn(cid, "Sent: Acquired SVID", len(acquiredSvid))

	return source
}

// CreateRootKey generates or reuses a cryptographic key pair for the
// application, taking an id for logging purposes and a channel updatedSecret
// to signal when the secret has been updated. If the secret key is not mounted
// at the expected location or there are any errors reading the key file, the
// function logs a fatal message and exits. If the secret has not been set in
// the cluster, the function generates a new key pair, persists them, and
// signals the updatedSecret channel.
func CreateRootKey(id *string, updatedSecret chan<- bool) {
	if env.RootKeyInputModeManual() {
		log.InfoLn(id,
			"Manual key input enabled. Skipping automatic key generation.")

		updatedSecret <- true

		return
	}

	// This is a Kubernetes Secret, mounted as a file.
	keyPath := env.RootKeyPathForSafe()

	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		log.FatalLn(id,
			"CreateRootKey: Secret key not mounted at", keyPath)
		return
	}

	data, err := os.ReadFile(keyPath)
	if err != nil {
		log.FatalLn(id,
			"CreateRootKey: Error reading file:", err.Error())
		return
	}

	secret := string(data)

	if secret != val.BlankRootKey {
		log.InfoLn(id,
			"Secret has been set in the cluster, will reuse it")
		crypto.SetRootKeyInMemory(secret)

		updatedSecret <- true

		return
	}

	log.InfoLn(id,
		"Secret has not been set yet. Will compute a secure secret.")

	rkt, err := crypto.NewRootKeyCollection()

	if err != nil {
		log.FatalLn(id, "Failed to generate keys", err.Error())
	}

	log.InfoLn(id, "Generated public key, private key, and aes seed")

	// Save the key to VSecM Safe's memory, and also save them to
	// VSecM Safe's root key Kubernetes Secret.
	// If the root key input mode is not manual, VSecM Safe will
	// use a trusted backing store to retrieve the key in case of
	// Pod crash or eviction. As of Mar,17, 2024 the only trusted
	// backing store is the VSecM root key Kubernetes Secret; however
	// this will change in the future.
	if err = PersistRootKeysToRootKeyBackingStore(rkt); err != nil {
		log.FatalLn(id, "Failed to persist keys", err.Error())
	}

	updatedSecret <- true
}
