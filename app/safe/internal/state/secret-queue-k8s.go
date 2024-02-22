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
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// The secrets put here are synced with their Kubernetes Secret counterparts.
var k8sSecretQueue = make(chan entity.SecretStored, env.SafeK8sSecretBufferSize())

func processK8sSecretQueue() {
	errChan := make(chan error)

	id := "AEGIHK8S"

	go func() {
		for e := range errChan {
			// If the `persistK8s` operation spews out an error, log it.
			log.ErrorLn(&id, "processK8sSecretQueue: error persisting secret:", e.Error())
		}
	}()

	for {
		// Buffer overflow check.
		if len(secretQueue) == env.SafeSecretBufferSize() {
			log.ErrorLn(
				&id,
				"processK8sSecretQueue: there are too many k8s secrets queued. "+
					"The goroutine will BLOCK until the queue is cleared.",
			)
		}

		// Get a secret to be persisted to the disk.
		secret := <-k8sSecretQueue

		cid := secret.Meta.CorrelationId

		log.TraceLn(&cid, "processK8sSecretQueue: picked k8s secret")

		// Sync up the secret to etcd as a Kubernetes Secret.
		//
		// Each secret is synced one at a time, with the order they
		// come in.
		//
		// Do not call this function elsewhere.
		// It is meant to be called inside this `processK8sSecretQueue` goroutine.
		persistK8s(secret, errChan)

		log.TraceLn(&cid, "processK8sSecretQueue: Should have persisted k8s secret")
	}
}
