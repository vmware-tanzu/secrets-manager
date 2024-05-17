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
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/io"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// K8sSecretUpsertQueue has the secrets to be synced with their Kubernetes
// `Secret` counterparts.
var K8sSecretUpsertQueue = make(
	chan entity.SecretStored, env.K8sSecretBufferSizeForSafe())

// ProcessK8sPrefixedSecretQueue continuously processes a queue of Kubernetes
// secrets (K8sSecretUpsertQueue), attempting to persist each secret into the
// Kubernetes cluster, specifically into etcd as a Kubernetes Secret. The
// function employs asynchronous error handling and is designed to operate
// continuously within a dedicated goroutine.
//
// This queue is for secrets that are generated via the `k8s:` prefix in their
// workload name. When the workload name starts with a `k8s:` prefix, VSecM will
// create an actual Kubernetes secret along with the secret it stores locally.
// This is especially useful for legacy use cases where a workload cannot
// directly consume vsecm-minted secrets.
func ProcessK8sPrefixedSecretQueue() {
	errChan := make(chan error)

	id := crypto.Id()

	go func() {
		for e := range errChan {
			// If the `persistK8s` operation spews out an error, log it.
			log.ErrorLn(&id,
				"processK8sSecretQueue: error persisting secret:", e.Error())
		}
	}()

	for {
		// Buffer overflow check.
		if len(SecretUpsertQueue) == env.SecretBufferSizeForSafe() {
			log.ErrorLn(
				&id,
				"processK8sSecretQueue:"+
					" there are too many k8s secrets queued. "+
					"The goroutine will BLOCK until the queue is cleared.",
			)
		}

		// Get a secret to be persisted to the disk.
		secret := <-K8sSecretUpsertQueue

		cid := secret.Meta.CorrelationId

		log.TraceLn(&cid, "processK8sSecretQueue: picked k8s secret")

		// Sync up the secret to etcd as a Kubernetes Secret.
		//
		// Each secret is synced one at a time, with the order they
		// come in.
		//
		// Do not call this function elsewhere.
		// It is meant to be called inside this `processK8sSecretQueue`
		// goroutine.
		io.PersistToK8s(secret, errChan)

		log.TraceLn(&cid,
			"processK8sSecretQueue: Should have persisted k8s secret")
	}
}
