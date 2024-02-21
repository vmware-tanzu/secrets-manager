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
	"encoding/json"
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

func isSentinel(j audit.JournalEntry, cid string, w http.ResponseWriter, spiffeid string) bool {
	audit.Log(j)

	if validation.IsSentinel(spiffeid) {
		return true
	}

	j.Event = audit.EventBadSvid
	audit.Log(j)

	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "")
	if err != nil {
		log.InfoLn(&cid, "Delete: Problem sending response", err.Error())
	}

	return false
}

func Delete(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	if env.SafeManualKeyInput() && !state.MasterKeySet() {
		log.InfoLn(&cid, "Delete: Master key not set")
		return
	}

	j := audit.JournalEntry{
		CorrelationId: cid,
		Entity:        reqres.SecretDeleteRequest{},
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         audit.EventEnter,
	}

	if !isSentinel(j, cid, w, spiffeid) {
		return
	}

	log.DebugLn(&cid, "Delete: sentinel spiffeid:", spiffeid)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		j.Event = audit.EventBrokenBody
		audit.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Delete: Problem sending response", err.Error())
		}
		return
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

	log.DebugLn(&cid, "Secret: Parsed request body")

	var sr reqres.SecretDeleteRequest
	err = json.Unmarshal(body, &sr)
	if err != nil {
		j.Event = audit.EventRequestTypeMismatch
		audit.Log(j)
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Delete: Problem sending response", err.Error())
		}
		return
	}

	j.Entity = sr

	workloadId := sr.WorkloadId

	if workloadId == "" {
		j.Event = audit.EventNoWorkloadId
		audit.Log(j)
		return
	}

	log.DebugLn(&cid, "Secret:Delete: ", "workloadId:", workloadId)

	if workloadId == "" {
		j.Event = audit.EventNoWorkloadId
		audit.Log(j)

		return
	}

	state.DeleteSecret(entity.SecretStored{
		Name: workloadId,
		Meta: entity.SecretMeta{
			CorrelationId: cid,
		},
	})
	log.DebugLn(&cid, "Delete:End: workloadId:", workloadId)

	j.Event = audit.EventOk
	audit.Log(j)

	_, err = io.WriteString(w, "OK")
	if err != nil {
		log.InfoLn(&cid, "Delete: Problem sending response", err.Error())
	}
}
