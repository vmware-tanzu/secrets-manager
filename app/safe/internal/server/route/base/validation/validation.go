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

	ioState "github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/io"
	"github.com/vmware-tanzu/secrets-manager/core/audit/journal"
	"github.com/vmware-tanzu/secrets-manager/core/constants/audit"
	"github.com/vmware-tanzu/secrets-manager/core/constants/val"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
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

// CheckDatabaseReadiness checks if the database is ready for use.
//
// This function verifies the readiness of the database, specifically for
// PostgreSQL mode. If the database is not initialized when PostgreSQL
// mode is enabled, it returns an error response.
//
// Parameters:
//   - cid: A string representing the context or correlation ID for logging.
//   - w: An http.ResponseWriter to write the HTTP response.
//
// Returns:
//   - bool: true if the database is ready, false otherwise.
//
// Side effects:
//   - Writes an HTTP 503 (Service Unavailable) status and response body
//     if the database is not ready.
//   - Logs information about the database status.
func CheckDatabaseReadiness(cid string, w http.ResponseWriter) bool {

	// If postgres mode enabled and db is not initialized, return error.
	if env.BackingStoreForSafe() == data.Postgres && !ioState.PostgresReady() {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, err := io.WriteString(w, val.NotOk)
		if err != nil {
			log.ErrorLn(&cid, "error writing response", err.Error())
		}
		log.InfoLn(&cid, "Secret: Database not initialized")
		return false
	}

	return true
}
