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
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/base/validation"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/secret/collection"
	"github.com/vmware-tanzu/secrets-manager/core/audit/journal"
	"github.com/vmware-tanzu/secrets-manager/core/constants/audit"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/v1/reqres/safe"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	s "github.com/vmware-tanzu/secrets-manager/lib/spiffe"
)

// Delete handles the deletion of a secret identified by a workload ID.
// It performs a series of checks and logging steps before carrying out the
// deletion.
//
// Parameters:
//   - cid: A string representing the correlation ID for the request, used for
//     tracking and logging purposes.
//   - w: An http.ResponseWriter object used to send responses back to the
//     client.
//   - r: An http.Request object containing the request details from the client.
//   - spiffeid: A string representing the SPIFFE ID of the client making the
//     request.
func Delete(
	cid string, r *http.Request, w http.ResponseWriter,
) {
	spiffeid := s.IdAsString(r)

	if !crypto.RootKeySetInMemory() {
		log.InfoLn(&cid, "Delete: Root key not set")

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(
				&cid, "Delete: Problem sending response",
				err.Error())
		}

		return
	}

	j := entity.JournalEntry{
		CorrelationId: cid,
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         audit.Enter,
	}

	// Only sentinel can execute delete requests.
	if ok, respond := validation.IsSentinel(j, cid, spiffeid); !ok {
		j.Event = audit.NotSentinel
		journal.Log(j)
		respond(w)
		return
	}

	log.DebugLn(&cid, "Delete: sentinel spiffeid:", spiffeid)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		j.Event = audit.BrokenBody
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")

		if err != nil {
			log.InfoLn(
				&cid,
				"Delete: Problem sending response",
				err.Error())
		}

		return
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}

		err := b.Close()
		if err != nil {
			log.InfoLn(
				&cid,
				"Delete: Problem closing body", err.Error())
		}
	}(r.Body)

	log.DebugLn(&cid, "Delete: Parsed request body")

	var sr reqres.SecretDeleteRequest
	err = json.Unmarshal(body, &sr)
	if err != nil {
		log.DebugLn(&cid,
			"Delete: Error unmarshalling request body",
			err.Error())

		j.Event = audit.RequestTypeMismatch
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(
				&cid,
				"Delete: Problem sending response",
				err.Error())
		}

		log.TraceLn(&cid, "Delete: Exiting from error case")
		return
	}

	workloadIds := sr.WorkloadIds

	if len(workloadIds) == 0 {
		log.TraceLn(&cid, "Delete: Empty workload ids")

		j.Event = audit.NoWorkloadId
		journal.Log(j)

		log.TraceLn(
			&cid,
			"Delete: Exiting from empty workload ids case")

		return
	}

	log.DebugLn(&cid, "Secret:Delete: ", "workloadIds:", workloadIds)

	for _, workloadId := range workloadIds {
		collection.DeleteSecret(entity.SecretStored{
			Name: workloadId,
			Meta: entity.SecretMeta{
				CorrelationId: cid,
			},
		})
	}

	log.DebugLn(&cid, "Delete:End: workloadIds:", workloadIds)

	j.Event = audit.Ok
	journal.Log(j)

	_, err = io.WriteString(w, "OK")
	if err != nil {
		log.InfoLn(
			&cid,
			"Delete: Problem sending response", err.Error())
	}
}
