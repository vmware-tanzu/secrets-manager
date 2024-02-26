/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package json

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func UnmarshalSecretUpsertRequest(cid string, body []byte, j audit.JournalEntry,
	w http.ResponseWriter) *reqres.SecretUpsertRequest {
	var sr reqres.SecretUpsertRequest

	err := json.Unmarshal(body, &sr)
	if err != nil {
		j.Event = event.RequestTypeMismatch
		audit.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}

		return nil
	}

	return &sr
}

func UnmarshalKeyInputRequest(cid string, body []byte, j audit.JournalEntry,
	w http.ResponseWriter) *reqres.KeyInputRequest {
	var sr reqres.KeyInputRequest

	err := json.Unmarshal(body, &sr)
	if err != nil {
		j.Event = event.RequestTypeMismatch
		audit.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")

		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}

		return nil
	}

	return &sr
}
