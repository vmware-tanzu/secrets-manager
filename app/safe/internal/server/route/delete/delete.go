/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package delete

import (
	"encoding/json"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/validation"
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// Delete handles the deletion of a secret identified by a workload ID.
// It performs a series of checks and logging steps before carrying out the deletion.
//
// Parameters:
//   - cid: A string representing the correlation ID for the request, used for
//     tracking and logging purposes.
//   - w: An http.ResponseWriter object used to send responses back to the client.
//   - r: An http.Request object containing the request details from the client.
//   - spiffeid: A string representing the SPIFFE ID of the client making the request.
func Delete(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	if !state.RootKeySet() {
		log.InfoLn(&cid, "Delete: Root key not set")
		return
	}

	j := audit.JournalEntry{
		CorrelationId: cid,
		Entity:        reqres.SecretDeleteRequest{},
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         event.Enter,
	}

	if !validation.IsSentinel(j, cid, w, spiffeid) {
		j.Event = event.BadSpiffeId
		audit.Log(j)
		return
	}

	log.DebugLn(&cid, "Delete: sentinel spiffeid:", spiffeid)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		j.Event = event.BrokenBody
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
		j.Event = event.RequestTypeMismatch
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
		j.Event = event.NoWorkloadId
		audit.Log(j)
		return
	}

	log.DebugLn(&cid, "Secret:Delete: ", "workloadId:", workloadId)

	if workloadId == "" {
		j.Event = event.NoWorkloadId
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

	j.Event = event.Ok
	audit.Log(j)

	_, err = io.WriteString(w, "OK")
	if err != nil {
		log.InfoLn(&cid, "Delete: Problem sending response", err.Error())
	}
}
