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
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

func isSentinel(j audit.JournalEntry, cid string, w http.ResponseWriter, spiffeid string) bool {
	audit.Log(j)

	if validation.IsSentinel(spiffeid) {
		return true
	}

	if env.SafeEnableOIDCResourceServer() {
		return true
	}

	j.Event = event.BadSpiffeId
	audit.Log(j)

	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "NOK!")
	if err != nil {
		log.ErrorLn(&cid, "Problem sending response", err.Error())
	}

	return false
}
