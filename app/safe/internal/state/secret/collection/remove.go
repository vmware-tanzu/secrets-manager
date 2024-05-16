/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package collection

import (
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/secret/queue/deletion"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/stats"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// DeleteSecret orchestrates the deletion of a specified secret from both the
// application's internal cache and its persisted storage locations, which may
// include local filesystem and Kubernetes secrets. The deletion process is
// contingent upon the secret's metadata, specifically its backing store and
// whether it is used as a Kubernetes secret.
//
// Parameters:
//   - secretToDelete (entity.SecretStored): The secret entity marked for
//     deletion, containing necessary metadata such as the name of the secret,
//     its correlation ID for logging, and metadata specifying where and how
//     the secret is stored.
func DeleteSecret(secretToDelete entity.SecretStored) {
	cid := secretToDelete.Meta.CorrelationId

	_, exists := Secrets.Load(secretToDelete.Name)
	if !exists {
		log.WarnLn(&cid,
			"DeleteSecret: Secret does not exist. Cannot delete.",
			secretToDelete.Name)
		return
	}

	log.TraceLn(
		&cid, "DeleteSecret: Will delete secret. len",
		len(deletion.SecretDeleteQueue),
		"cap", cap(deletion.SecretDeleteQueue))

	// The deletion queue will remove the secret from the backing store.
	// The backing store is determined by the env.BackingStoreForSafe()
	// function.
	deletion.SecretDeleteQueue <- secretToDelete

	log.TraceLn(
		&cid, "DeleteSecret: Pushed secret to delete. len",
		len(deletion.SecretDeleteQueue), "cap",
		cap(deletion.SecretDeleteQueue))

	// Remove the secret from the memory.
	stats.CurrentState.Decrement(secretToDelete.Name, Secrets.Load)
	Secrets.Delete(secretToDelete.Name)
}
