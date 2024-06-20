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
	"github.com/vmware-tanzu/secrets-manager/core/constants/audit"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// SendEncryptedValue takes a plain text value and encrypts it. If the
// encryption is successful, the encrypted value is written to the HTTP
// response. In case of an error or if the input value is empty, it logs the
// event, updates the HTTP response status accordingly, and may log additional
// errors encountered when sending the HTTP response.
//
// Parameters:
//   - cid (string): Correlation ID for tracing the operation through logs.
//   - value (string): The plain text value to be encrypted.
//   - j (audit.JournalEntry): An audit journal entry for recording the event.
//   - w (http.ResponseWriter): The HTTP response writer to send back the
//     encrypted value or errors.
func SendEncryptedValue(
	cid string, value string, j data.JournalEntry, w http.ResponseWriter,
) {
	if value == "" {
		j.Event = audit.NoValue
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")

		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}

		return
	}

	encrypted, err := crypto.EncryptValue(value)
	if err != nil {
		j.Event = audit.EncryptionFailed
		journal.Log(j)

		w.WriteHeader(http.StatusInternalServerError)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}

		return
	}

	_, err = io.WriteString(w, encrypted)
	if err != nil {
		log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
	}
}
