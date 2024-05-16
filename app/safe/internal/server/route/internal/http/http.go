/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package http

import (
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/core/audit/journal"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// ReadBody reads the HTTP request body and returns it as a byte slice. It
// handles errors by logging and responding with an appropriate HTTP status
// code.
//
// Parameters:
//   - cid (string): Correlation ID for operation tracing and logging.
//   - r (*http.Request): The HTTP request from which the body will be read.
//   - w (http.ResponseWriter): The HTTP response writer to send back errors in
//     case of failure.
//   - j (audit.JournalEntry): An audit journal entry for recording the event.
//
// Returns:
//   - []byte: The request body as a byte slice. Returns nil if an error occurs
//     while reading the body.
func ReadBody(
	cid string, r *http.Request, w http.ResponseWriter,
	j journal.Entry,
) []byte {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		j.Event = event.BrokenBody
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}

		return nil
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem closing body", err.Error())
		}
	}(r.Body)

	return body
}
