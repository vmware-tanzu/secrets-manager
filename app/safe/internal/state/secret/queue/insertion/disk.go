/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package insertion

import (
	// "github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/io/persistence"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// SecretUpsertQueue items are persisted to files. They are buffered, so that they
// can be written in the order they are queued and there are no concurrent
// writes to the same file at a time. An alternative approach would be
// to have a map of queues of `SecretsStored`s per file name but that
// feels like an overkill.
var SecretUpsertQueue = make(chan entity.SecretStored, env.SecretBufferSizeForSafe())

// ProcessSecretQueue manages a continuous loop that processes secrets from the
// SecretUpsertQueue, persisting each secret to disk storage. This function is
// crucial for ensuring that changes to secrets are reliably stored, supporting
// both new secrets and updates to existing ones. The operations of this function
// is critical for maintaining the integrity and consistency of secret data
// within the system.
func ProcessSecretQueue(rootKeyTriplet []string) {
	errChan := make(chan error)

	id := "AEGIHSCR"

	go func() {
		for e := range errChan {
			// If the `persist` operation spews out an error, log it.
			log.ErrorLn(&id, "processSecretQueue: error persisting secret:", e.Error())
		}
	}()

	for {
		// Buffer overflow check.
		if len(SecretUpsertQueue) == env.SecretBufferSizeForSafe() {
			log.ErrorLn(
				&id,
				"processSecretQueue: there are too many k8s secrets queued. "+
					"The goroutine will BLOCK until the queue is cleared.",
			)
		}

		// Get a secret to be persisted to the disk.
		secret := <-SecretUpsertQueue

		cid := secret.Meta.CorrelationId

		log.TraceLn(&cid, "processSecretQueue: picked a secret", len(SecretUpsertQueue))

		// Persist the secret to disk.
		//
		// Each secret is persisted one at a time, with the order they
		// come in.
		//
		// Do not call this function elsewhere.
		// It is meant to be called inside this `processSecretQueue` goroutine.
		persistence.PersistToDisk(secret, rootKeyTriplet, errChan)

		log.TraceLn(&cid, "processSecretQueue: should have persisted the secret.")
	}
}
