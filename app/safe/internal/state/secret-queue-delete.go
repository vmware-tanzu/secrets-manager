/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package state

import (
	"os"
	"path"

	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// These are persisted to files. They are buffered, so that they can
// be written in the order they are queued and there are no concurrent
// writes to the same file at a time. An alternative approach would be
// to have a map of queues of `SecretsStored`s per file name but that
// feels like an overkill.
var secretDeleteQueue = make(chan entity.SecretStored, env.SafeSecretDeleteBufferSize())

func processSecretDeleteQueue() {
	errChan := make(chan error)

	id := "AEGIHSCD"

	go func() {
		for e := range errChan {
			// If the `delete` operation spews out an error, log it.
			log.ErrorLn(&id, "processSecretDeleteQueue: error deleting secret:", e.Error())
		}
	}()

	for {
		// Buffer overflow check.
		if len(secretDeleteQueue) == env.SafeSecretBufferSize() {
			log.ErrorLn(
				&id,
				"processSecretDeleteQueue: there are too many k8s secrets queued. "+
					"The goroutine will BLOCK until the queue is cleared.",
			)
		}

		// Get a secret to be removed from the disk.
		secret := <-secretDeleteQueue

		if secret.Name == "" {
			log.WarnLn(&id, "processSecretDeleteQueue: trying to delete an empty secret. "+
				"Possibly picked a nil secret", len(secretQueue))
			return
		}

		log.TraceLn(&id, "processSecretDeleteQueue: picked a secret", len(secretQueue))

		// Remove secret from disk.
		dataPath := path.Join(env.SafeDataPath(), secret.Name+".age")
		log.TraceLn(&id, "processSecretDeleteQueue: removing secret from disk:", dataPath)
		err := os.Remove(dataPath)
		if err != nil && !os.IsNotExist(err) {
			log.WarnLn(&id, "processSecretDeleteQueue: failed to remove secret", err.Error())
		}

		log.TraceLn(&id, "processSecretDeleteQueue: should have deleted the secret.")
	}
}
