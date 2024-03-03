/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package crypto

import (
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// Encrypt takes a plain text value and encrypts it. If the encryption is successful,
// the encrypted value is written to the HTTP response. In case of an error or if
// the input value is empty, it logs the event, updates the HTTP response status
// accordingly, and may log additional errors encountered when sending
// the HTTP response.
//
// Parameters:
//   - cid (string): Correlation ID for tracing the operation through logs.
//   - value (string): The plain text value to be encrypted.
//   - j (audit.JournalEntry): An audit journal entry for recording the event.
//   - w (http.ResponseWriter): The HTTP response writer to send back the encrypted
//     value or errors.
func Encrypt(cid string, value string, j audit.JournalEntry,
	w http.ResponseWriter) {
	if value == "" {
		j.Event = event.NoValue
		audit.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")

		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}

		return
	}

	encrypted, err := state.EncryptValue(value)
	if err != nil {
		j.Event = event.EncryptionFailed
		audit.Log(j)

		w.WriteHeader(http.StatusInternalServerError)
		_, err2 := io.WriteString(w, "")
		if err2 != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err2.Error())
		}

		return
	}

	_, err = io.WriteString(w, encrypted)
	if err != nil {
		log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
	}
	return
}

// Decrypt attempts to decrypt a given encrypted value. It returns the decrypted
// string and a boolean indicating whether a decryption error occurred. In case of
// a decryption failure, it logs the event, sends an appropriate HTTP response,
// and returns an indication of the failure.
//
// Parameters:
// - cid (string): Correlation ID for operation tracing.
// - value (string): The encrypted value to be decrypted.
// - j (audit.JournalEntry): An audit journal entry for event recording.
// - w (http.ResponseWriter): The HTTP response writer for sending back errors.
//
// Returns:
//   - (string, bool): The first return value is the decrypted string, which will
//     be empty in case of an error. The second return value is a boolean flag that
//     is true if a decryption error occurred, and false otherwise.
func Decrypt(cid string, value string, j audit.JournalEntry,
	w http.ResponseWriter) (string, bool) {
	decrypted, err := state.DecryptValue(value)
	if err != nil {
		j.Event = event.DecryptionFailed
		audit.Log(j)

		w.WriteHeader(http.StatusInternalServerError)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}

		return "", true
	}

	return decrypted, false
}
