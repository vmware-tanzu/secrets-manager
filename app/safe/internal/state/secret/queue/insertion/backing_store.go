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
	"time"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/io"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// SecretUpsertQueue items are persisted to files. They are buffered, so
// that they can be written in the order they are queued and there are
// no concurrent writes to the same file at a time. An alternative
// approach would be to have a map of queues of `SecretsStored`s per file
// name but that feels like an overkill.
var SecretUpsertQueue = make(
	chan entity.SecretStored,
	env.SecretBufferSizeForSafe(),
)

func pickSecretFromQueue(cid string) entity.SecretStored {
	// Get a secret to be persisted.
	secret := <-SecretUpsertQueue

	log.TraceLn(&cid, "ProcessSecretQueue: Will persist to Postgres.",
		len(SecretUpsertQueue))

	return secret
}

// ProcessSecretBackingStoreQueue manages a continuous loop that processes
// secrets from the SecretUpsertQueue, persisting each secret to disk storage.
// This function is crucial for ensuring that changes to secrets are reliably
// stored, supporting both new secrets and updates to existing ones. The
// operations of this function is critical for maintaining the integrity and
// consistency of secret data within the system.
func ProcessSecretBackingStoreQueue() {
	errChan := make(chan error)

	cid := crypto.Id()

	go func() {
		for e := range errChan {
			// If the `persist` operation spews out an error, log it.
			log.ErrorLn(
				&cid, "processSecretQueue: error persisting secret:", e.Error(),
			)
		}
	}()

	for {
		// Buffer overflow check.
		if len(SecretUpsertQueue) == env.SecretBufferSizeForSafe() {
			log.ErrorLn(
				&cid,
				"processSecretQueue: there are too many k8s secrets queued. "+
					"The goroutine will BLOCK until the queue is cleared.",
			)
		}

		// Persist the secret:
		//
		// Each secret is persisted one at a time, with the order they come in.
		//
		// DO NOT call `PersistToXyz()` functions (i.e.,[1],[2]) elsewhere.
		// They mean to be called inside this goroutine that
		// `ProcessSecretBackingStoreQueue` owns.

		store := env.BackingStoreForSafe()
		switch store {
		case entity.Memory:
			log.TraceLn(&cid, "ProcessSecretQueue: using in-memory store.")
			return
		case entity.Kubernetes:
			panic("implement kubernetes store")
		case entity.AwsSecretStore:
			panic("implement aws secret store")
		case entity.AzureSecretStore:
			panic("implement azure secret store")
		case entity.GcpSecretStore:
			panic("implement gcp secret store")
		case entity.File:
			log.TraceLn(&cid, "ProcessSecretQueue: Will persist to disk.")

			// This is blocking. [1]
			io.PersistToDisk(pickSecretFromQueue(cid), errChan)
		case entity.Postgres:
			if !io.PostgresReady() {
				log.TraceLn(&cid, "ProcessSecretQueue: Postgres is not ready. Waiting...")

				// Give the loop some time to breathe.
				// This is to prevent the loop from spinning too fast.
				time.Sleep(5 * time.Second)

				continue
			}

			log.TraceLn(&cid, "ProcessSecretQueue: Postgres is ready. Persisting...")

			// This is blocking. [2]
			io.PersistToPostgres(pickSecretFromQueue(cid), errChan)
		}

		log.TraceLn(&cid, "processSecretQueue: should have persisted the secret.")
	}
}
