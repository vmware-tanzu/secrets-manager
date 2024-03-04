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
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// RootKeySet checks if a root key has been set in the application's state. This
// function is typically used to determine if the application's cryptographic
// foundation is initialized and ready for operation.
//
// Returns:
//   - bool: True if a root key is set, indicating that cryptographic operations
//     can proceed. False indicates that the root key is not set, and cryptographic
//     operations may not be possible.
func RootKeySet() bool {
	return state.RootKeySet()
}

// Upsert handles the insertion or update of a secret in the application's state.
// It supports appending values to existing secrets and logs the completion of
// the operation. If specified, it also sends an HTTP response indicating success.
//
// Parameters:
//   - secretToStore (entity.SecretStored): The secret entity to be inserted or updated.
//   - appendValue (bool): A flag indicating whether to append the value to an existing
//     secret (if true) or overwrite the existing secret (if false).
//   - workloadId (string): The identifier of the workload associated with the
//     secret operation, used for logging purposes.
//   - cid (string): Correlation ID for operation tracing and logging.
//   - j (audit.JournalEntry): An audit journal entry for recording the event.
//   - w (http.ResponseWriter): The HTTP response writer to send back the
//     operation's outcome.
func Upsert(secretToStore entity.SecretStored,
	appendValue bool, workloadId string, cid string,
	j audit.JournalEntry, w http.ResponseWriter,
) {
	state.UpsertSecret(secretToStore, appendValue)
	log.DebugLn(&cid, "Secret:UpsertEnd: workloadId", workloadId)

	j.Event = event.Ok
	audit.Log(j)

	_, err := io.WriteString(w, "OK")
	if err != nil {
		log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
	}
}
