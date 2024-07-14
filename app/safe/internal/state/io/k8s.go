/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package io

import (
	s "github.com/vmware-tanzu/secrets-manager/core/constants/secret"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/lib/backoff"
)

// PersistToK8s attempts to save a provided secret entity into a Kubernetes
// cluster. The function is structured to handle potential errors through
// retries and communicates any persistent issues back to the caller via an
// error channel. It employs logging for traceability of the operation's
// progress and outcomes.
//
// Parameters:
//   - secret (entity.SecretStored): A structured entity containing the secret's
//     metadata and values to be persisted. The metadata includes a
//     CorrelationId for tracing the operation.
//   - errChan (chan<- error): A channel through which errors are reported. This
//     channel allows the function to notify the caller of any failures in
//     persisting the secret, enabling asynchronous error handling.
func PersistToK8s(secret entity.SecretStored, errChan chan<- error) {
	cid := secret.Meta.CorrelationId

	log.TraceLn(&cid, "persistK8s: Will persist k8s secret.")

	if len(secret.Values) == 0 {
		secret.Values = append(secret.Values, s.InitialValue)
	}

	// Defensive coding:
	// secret's value is never empty because when the value is set to an
	// empty secret, it is scheduled for deletion and not persisted to the
	// file system or the cluster. However, it that happens, we would at least
	// want an indicator that it happened.
	if secret.Values[0] == "" {
		secret.Values[0] = s.InitialValue
	}

	log.TraceLn(&cid, "persistK8s: Will try saving secret to k8s.")
	err := backoff.RetryExponential("PersistToK8s", func() error {
		return saveSecretToKubernetes(secret)
	})

	if err != nil {
		log.TraceLn(
			&cid, "persistK8s: still error, pushing the error to errChan")
		errChan <- err
	}
}
