/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package bootstrap

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"filippo.io/age"
	"github.com/pkg/errors"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"github.com/vmware-tanzu/secrets-manager/core/probe"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
	"os"
	"time"
)

// NotifyTimeout waits for the duration specified by env.SafeBootstrapTimeout()
// and then sends a 'true' value to the provided 'timedOut' channel. This function
// can be used to notify other parts of the application when a specific timeout
// has been reached.
func NotifyTimeout(timedOut chan<- bool) {
	time.Sleep(env.SafeBootstrapTimeout())
	timedOut <- true
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
	acquiredSvid <-chan bool,
	updatedSecret <-chan bool,
	serverStarted <-chan bool,
	timedOut <-chan bool,
) {
	counter := 3
	for {
		if counter == 0 {
			break
		}
		select {
		case <-acquiredSvid:
			log.InfoLn(correlationId, "Acquired identity.")
			counter--
			log.InfoLn(correlationId, "remaining:", counter)
			if counter == 0 {
				state.Initialize()
				log.DebugLn(correlationId, "Creating readiness probe.")
				go probe.CreateReadiness()
			}
		case <-updatedSecret:
			log.InfoLn(correlationId, "Updated age key.")
			counter--
			log.InfoLn(correlationId, "remaining:", counter)
			if counter == 0 {
				state.Initialize()
				log.DebugLn(correlationId, "Creating readiness probe.")
				go probe.CreateReadiness()
			}
		case <-serverStarted:
			log.InfoLn(correlationId, "Server ready.")
			counter--
			log.InfoLn(correlationId, "remaining:", counter)
			if counter == 0 {
				state.Initialize()
				log.DebugLn(correlationId, "Creating readiness probe.")
				go probe.CreateReadiness()
			}
		case <-timedOut:
			log.FatalLn(correlationId, "Failed to acquire an identity in a timely manner.")
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
	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(
			workloadapi.WithAddr(env.SpiffeSocketUrl()),
		),
	)

	id := ctx.Value("correlationId").(*string)

	if err != nil {
		log.FatalLn(id, "Unable to fetch X.509 Bundle", err.Error())
	}

	if source == nil {
		log.FatalLn(id, "Could not find source")
	}

	svid, err := source.GetX509SVID()
	if err != nil {
		log.FatalLn(id, "Unable to get X.509 SVID from source bundle", err.Error())
	}

	svidId := svid.ID
	if !validation.IsSafe(svid.ID.String()) {
		log.FatalLn(
			id, "Svid check: I don’t know you, and it’s crazy:", svidId.String(),
		)
	}

	log.TraceLn(id, "Sending: Acquired SVID", len(acquiredSvid))
	acquiredSvid <- true
	log.TraceLn(id, "Sent: Acquired SVID", len(acquiredSvid))

	return source
}

func generateAesSeed() (string, error) {
	// Generate a 256 bit key
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		return "", errors.Wrap(err, "generateAesSeed: failed to generate random key")
	}

	return hex.EncodeToString(key), nil
}

// CreateCryptoKey generates or reuses a cryptographic key pair for the
// application, taking an id for logging purposes and a channel updatedSecret
// to signal when the secret has been updated. If the secret key is not mounted
// at the expected location or there are any errors reading the key file, the
// function logs a fatal message and exits. If the secret has not been set in
// the cluster, the function generates a new key pair, persists them, and
// signals the updatedSecret channel.
func CreateCryptoKey(id *string, updatedSecret chan<- bool) {
	if env.SafeManualKeyInput() {
		log.InfoLn(id, "Manual key input enabled. Skipping automatic key generation.")
		updatedSecret <- true
		return
	}

	// This is a Kubernetes Secret, mounted as a file.
	keyPath := env.SafeAgeKeyPath()

	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		log.FatalLn(id, "CreateCryptoKey: Secret key not mounted at", keyPath)
		return
	}

	data, err := os.ReadFile(keyPath)
	if err != nil {
		log.FatalLn(id, "CreateCryptoKey: Error reading file:", err.Error())
		return
	}

	secret := string(data)

	if secret != state.BlankAgeKeyValue {
		log.InfoLn(id, "Secret has been set in the cluster, will reuse it")
		state.SetMasterKey(secret)
		updatedSecret <- true
		return
	}

	log.InfoLn(id, "Secret has not been set yet. Will compute a secure secret.")

	identity, err := age.GenerateX25519Identity()
	if err != nil {
		log.FatalLn(id, "Failed to generate key pair", err.Error())
	}

	publicKey := identity.Recipient().String()
	privateKey := identity.String()
	aesSeed, err := generateAesSeed()

	if err != nil {
		log.FatalLn(id, "Failed to generate AES seed", err.Error())
	}

	log.TraceLn(id, "Public key: %s...  ", identity.Recipient().String()[:4])
	log.TraceLn(id, "Private key: %s...  ", identity.String()[:16])

	if err = persistKeys(privateKey, publicKey, aesSeed); err != nil {
		log.FatalLn(id, "Failed to persist keys", err.Error())
	}

	updatedSecret <- true
}
