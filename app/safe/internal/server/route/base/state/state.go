/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package state

import (
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"io"
	"net/http"

	net "github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/io"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/secret/collection"
	"github.com/vmware-tanzu/secrets-manager/core/audit/journal"
	"github.com/vmware-tanzu/secrets-manager/core/constants/audit"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// Upsert handles the insertion or update of a secret in the application's
// state. It supports appending values to existing secrets and logs the
// completion of the operation. If specified, it also sends an HTTP response
// indicating success.
//
// Parameters:
//   - secretToStore (entity.SecretStored): The secret entity to be inserted or
//     updated.
//   - appendValue (bool): A flag indicating whether to append the value to an
//     existing secret (if true) or overwrite the existing secret (if false).
//   - workloadId (string): The identifier of the workload associated with the
//     secret operation, used for logging purposes.
//   - cid (string): Correlation ID for operation tracing and logging.
//   - j (audit.JournalEntry): An audit journal entry for recording the event.
//   - w (http.ResponseWriter): The HTTP response writer to send back the
//     operation's outcome.
func Upsert(secretToStore entity.SecretStored,
	workloadId string, cid string,
	j entity.JournalEntry, w http.ResponseWriter,
) {
	// If the secret is not internal VSecM Safe configuration secret and
	// if db persistence is enabled, and the db is not ready,
	// then respond with an error.
	if secretToStore.Name != "vsecm-safe" &&
		env.BackingStoreForSafe() == entity.Postgres &&
		!net.PostgresReady() {
		log.InfoLn(&cid, "Secret: DB not ready. Responding with error.")
		w.WriteHeader(http.StatusInternalServerError)
		_, err := io.WriteString(w, "DB not ready")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}
		return
	}

	collection.UpsertSecret(secretToStore)
	log.DebugLn(&cid, "Secret:UpsertEnd: workloadId", workloadId)

	j.Event = audit.Ok
	journal.Log(j)

	_, err := io.WriteString(w, "OK")
	if err != nil {
		log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
	}
}
