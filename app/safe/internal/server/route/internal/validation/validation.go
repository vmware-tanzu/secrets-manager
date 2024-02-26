/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package validation

import (
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

func IsSentinel(j audit.JournalEntry, cid string, w http.ResponseWriter, spiffeid string) bool {
	audit.Log(j)

	if validation.IsSentinel(spiffeid) {
		return true
	}

	j.Event = event.BadSpiffeId
	audit.Log(j)

	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "")
	if err != nil {
		log.InfoLn(&cid, "Delete: Problem sending response", err.Error())
	}

	return false
}
