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
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func upsert(secretToStore entity.SecretStored,
	appendValue bool, workloadId string, cid string,
	j audit.JournalEntry, w http.ResponseWriter,
) {
	state.UpsertSecret(secretToStore, appendValue)
	log.DebugLn(&cid, "Secret:UpsertEnd: workloadId", workloadId)

	j.Event = event.Ok
	audit.Log(j)

	_, err := io.WriteString(w, "OK")
	if err != nil {
		log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
	}
}
