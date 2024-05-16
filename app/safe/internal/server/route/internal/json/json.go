/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package json

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/core/audit/journal"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/v1/reqres/safe"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// UnmarshalSecretUpsertRequest takes a JSON-encoded request body and attempts
// to unmarshal it into a SecretUpsertRequest struct. It handles JSON
// unmarshalling errors by logging, responding with an HTTP error, and returning
// nil. This function is typically used in HTTP server handlers to process
// incoming requests for secret upsert operations.
//
// Parameters:
//   - cid (string): Correlation ID for operation tracing and logging.
//   - body ([]byte): The JSON-encoded request body to be unmarshalled.
//   - j (audit.JournalEntry): An audit journal entry for recording the event.
//   - w (http.ResponseWriter): The HTTP response writer to send back errors in
//     case of failure.
//
// Returns:
//   - *reqres.SecretUpsertRequest: A pointer to the unmarshalled
//     SecretUpsertRequest struct, or nil if unmarshalling fails.
func UnmarshalSecretUpsertRequest(
	cid string, body []byte, j journal.Entry,
	w http.ResponseWriter,
) *reqres.SecretUpsertRequest {
	var sr reqres.SecretUpsertRequest

	if err := json.Unmarshal(body, &sr); err != nil {
		j.Event = event.RequestTypeMismatch
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}

		return nil
	}

	return &sr
}

// UnmarshalKeyInputRequest takes a JSON-encoded request body and attempts to
// unmarshal it into a KeyInputRequest struct. Similar to
// UnmarshalSecretUpsertRequest, it deals with JSON unmarshalling errors by
// logging, issuing an HTTP error response, and returning nil. This function is
// utilized within HTTP server handlers to parse incoming requests for key input
// operations.
//
// Parameters:
//   - cid (string): Correlation ID for operation tracing and logging.
//   - body ([]byte): The JSON-encoded request body to be unmarshalled.
//   - j (audit.JournalEntry): An audit journal entry for recording the event.
//   - w (http.ResponseWriter): The HTTP response writer to send back errors in
//     case of failure.
//
// Returns:
//   - *reqres.KeyInputRequest: A pointer to the unmarshalled KeyInputRequest
//     struct, or nil if unmarshalling fails.
func UnmarshalKeyInputRequest(cid string, body []byte, j journal.Entry,
	w http.ResponseWriter) *reqres.KeyInputRequest {
	var sr reqres.KeyInputRequest

	err := json.Unmarshal(body, &sr)
	if err != nil {
		j.Event = event.RequestTypeMismatch
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")

		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}

		return nil
	}

	return &sr
}
