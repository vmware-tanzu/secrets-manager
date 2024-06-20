/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package deletion

import (
	"os"
	"path"

	"github.com/vmware-tanzu/secrets-manager/core/constants/file"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// SecretDeleteQueue items are persisted to files. They are buffered, so that
// they can be written in the order they are queued and there are no concurrent
// writes to the same file at a time.
var SecretDeleteQueue = make(
	chan entity.SecretStored,
	env.SecretDeleteBufferSizeForSafe(),
)

// ProcessSecretBackingStoreQueue continuously processes a queue of secrets
// scheduled for deletion, removing each secret from disk. This function plays
// a crucial role in the secure management of secrets by ensuring that outdated
// or unnecessary secrets are not left stored, potentially posing a security risk.
//
// It operates in an endless loop, monitoring a global queue of secrets to be
// deleted.
func ProcessSecretBackingStoreQueue() {
	cid := crypto.Id()

	errChan := make(chan error)

	go func() {
		for e := range errChan {
			// If the `delete` operation spews out an error, log it.
			log.ErrorLn(&cid,
				"processSecretDeleteQueue: error deleting secret:", e.Error())
		}
	}()

	for {
		// Buffer overflow check.
		if len(SecretDeleteQueue) == env.SecretBufferSizeForSafe() {
			log.ErrorLn(
				&cid,
				"processSecretDeleteQueue: "+
					"there are too many k8s secrets queued. "+
					"The goroutine will BLOCK until the queue is cleared.",
			)
		}

		// Get a secret to be removed from the disk.
		secret := <-SecretDeleteQueue

		store := env.BackingStoreForSafe()
		switch store {
		case entity.Memory:
			log.TraceLn(&cid, "ProcessSecretQueue: using in-memory store.")
			return
		case entity.File:
			log.TraceLn(&cid, "ProcessSecretQueue: Will persist to disk.")
		case entity.Kubernetes:
			panic("implement kubernetes store")
		case entity.AwsSecretStore:
			panic("implement aws secret store")
		case entity.AzureSecretStore:
			panic("implement azure secret store")
		case entity.GcpSecretStore:
			panic("implement gcp secret store")
		}

		if secret.Name == "" {
			log.WarnLn(&cid,
				"processSecretDeleteQueue: trying to delete an empty secret. "+
					"Possibly picked a nil secret", len(SecretDeleteQueue))
			return
		}

		log.TraceLn(&cid,
			"processSecretDeleteQueue: picked a secret", len(SecretDeleteQueue))

		// Remove secret from disk.
		dataPath := path.Join(env.DataPathForSafe(), secret.Name+file.AgeExtension)
		log.TraceLn(&cid,
			"processSecretDeleteQueue: removing secret from disk:", dataPath)
		err := os.Remove(dataPath)
		if err != nil && !os.IsNotExist(err) {
			log.WarnLn(&cid,
				"processSecretDeleteQueue: failed to remove secret", err.Error())
		}

		log.TraceLn(&cid,
			"processSecretDeleteQueue: should have deleted the secret.")
	}
}
