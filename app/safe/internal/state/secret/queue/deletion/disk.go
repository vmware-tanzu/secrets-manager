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

	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// SecretDeleteQueue items are persisted to files. They are buffered, so that
// they can be written in the order they are queued and there are no concurrent
// writes to the same file at a time.
var SecretDeleteQueue = make(chan entity.SecretStored, env.SecretDeleteBufferSizeForSafe())

// ProcessSecretDeleteQueue continuously processes a queue of secrets scheduled for
// deletion, removing each secret from disk. This function plays a crucial role in
// the secure management of secrets by ensuring that outdated or unnecessary
// secrets are not left stored, potentially posing a security risk.
//
// It operates in an endless loop, monitoring a global queue of secrets to be
// deleted.
func ProcessSecretDeleteQueue() {
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
		if len(SecretDeleteQueue) == env.SecretBufferSizeForSafe() {
			log.ErrorLn(
				&id,
				"processSecretDeleteQueue: there are too many k8s secrets queued. "+
					"The goroutine will BLOCK until the queue is cleared.",
			)
		}

		// Get a secret to be removed from the disk.
		secret := <-SecretDeleteQueue

		if secret.Name == "" {
			log.WarnLn(&id, "processSecretDeleteQueue: trying to delete an empty secret. "+
				"Possibly picked a nil secret", len(SecretDeleteQueue))
			return
		}

		log.TraceLn(&id, "processSecretDeleteQueue: picked a secret", len(SecretDeleteQueue))

		// Remove secret from disk.
		dataPath := path.Join(env.DataPathForSafe(), secret.Name+".age")
		log.TraceLn(&id, "processSecretDeleteQueue: removing secret from disk:", dataPath)
		err := os.Remove(dataPath)
		if err != nil && !os.IsNotExist(err) {
			log.WarnLn(&id, "processSecretDeleteQueue: failed to remove secret", err.Error())
		}

		log.TraceLn(&id, "processSecretDeleteQueue: should have deleted the secret.")
	}
}
