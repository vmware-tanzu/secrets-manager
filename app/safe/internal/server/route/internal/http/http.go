/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package http

import (
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func ReadBody(cid string, r *http.Request, w http.ResponseWriter,
	j audit.JournalEntry) []byte {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		j.Event = event.BrokenBody
		audit.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err2 := io.WriteString(w, "")
		if err2 != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err2.Error())
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
