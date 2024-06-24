/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package fallback

import (
	"io"
	"net/http"

	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func Fallback(
	cid string, r *http.Request, w http.ResponseWriter,
) {
	log.DebugLn(&cid, "Handler: route mismatch:", r.RequestURI)

	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "")
	if err != nil {
		log.WarnLn(&cid, "Problem writing response:", err.Error())
	}
}
