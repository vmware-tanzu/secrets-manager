/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package route

import (
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func encryptValue(cid string, value string, j audit.JournalEntry,
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

func decryptValue(cid string, value string, j audit.JournalEntry,
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
