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
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/core/audit/journal"
	"github.com/vmware-tanzu/secrets-manager/core/constants/audit"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

// IsSentinel evaluates if a given SPIFFE ID corresponds to a VSecM Sentinel
// entity. It logs the operation and, if the SPIFFE ID is not recognized as
// VSecM Sentinel, logs an error event and sends an HTTP bad request response.
//
// Parameters:
//   - j: An instance of journal.Entry which is an audit log.
//   - cid: A string representing the correlation ID for the operation, used
//     primarily for logging.
//   - spiffeid: A string representing the SPIFFE ID to be validated against
//     sentinel conditions.
//
// Returns:
//   - bool: Returns true if the SPIFFE ID is a sentinel, otherwise false.
//   - func(http.ResponseWriter): Returns an HTTP handler function. If the
//     SPIFFE ID represents VSecM Sentinel, the handler is a no-op.
//     If the SPIFFE ID is not for VSecM Sentinel, it returns a handler that
//     responds with HTTP 400 Bad Request and logs the error if the response
//     writing fails.
//
// Note:
// This function should be used in scenarios where SPIFFE ID validation is
// critical for further processing steps, and appropriate HTTP response behavior
// needs to be enforced based on the validation results.
func IsSentinel(
	j data.JournalEntry, cid string, spiffeid string,
) (bool, func(http.ResponseWriter)) {
	journal.Log(j)

	if validation.IsSentinel(spiffeid) {
		return true, func(writer http.ResponseWriter) {}
	}

	j.Event = audit.BadSpiffeId
	journal.Log(j)

	var responder = func(w http.ResponseWriter) {
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Fetch: Problem sending response", err.Error())
		}
	}

	return false, responder
}
