/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package validation

import (
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

// IsSentinel evaluates if a given SPIFFE ID corresponds to a VSecM Sentinel entity.
// It logs the operation and, if the SPIFFE ID is not recognized as a sentinel,
// logs an error event and sends an HTTP bad request response. This function is
// typically used to validate SPIFFE IDs for specific operations or access control.
//
// Parameters:
//   - j (audit.JournalEntry): The initial audit journal entry for the operation.
//   - cid (string): Correlation ID for operation tracing and logging.
//   - w (http.ResponseWriter): The HTTP response writer to send back errors in
//     case of invalid SPIFFE ID.
//   - spiffeid (string): The SPIFFE ID to be evaluated.
//
// Returns:
//   - bool: Returns true if the given SPIFFE ID is recognized as a VSecM Sentinel
//     entity, indicating a special or reserved identity within the system. Returns
//     false if the SPIFFE ID does not correspond to a sentinel entity, signaling
//     a validation failure.
func IsSentinel(j audit.JournalEntry, cid string, w http.ResponseWriter,
	spiffeid string) bool {
	audit.Log(j)

	if validation.IsSentinel(spiffeid) {
		return true
	}

	if env.SafeEnableOIDCResourceServer() {
		return true
	}

	j.Event = event.BadSpiffeId
	audit.Log(j)

	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "NOK!")
	if err != nil {
		log.ErrorLn(&cid, "Problem sending response!", err.Error())
	}

	return false
}
