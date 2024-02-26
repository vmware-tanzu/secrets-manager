/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package fetch

import (
	"fmt"
	"net/http"
	"time"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/extract"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/handle"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

// Fetch handles the retrieval of a secret for a given workload, identified by
// its SPIFFE ID.
// The function performs several checks to ensure the request is valid and then
// fetches the secret.
//
// Parameters:
//   - cid: A string representing the correlation ID for the request, used for
//     tracking and logging purposes.
//   - w: An http.ResponseWriter object used to send responses back to the client.
//   - r: An http.Request object containing the request details from the client.
//   - spiffeid: A string representing the SPIFFE ID of the client making the request.
func Fetch(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	if !state.RootKeySet() {
		log.InfoLn(&cid, "Fetch: Root key not set")
		return
	}

	j := audit.JournalEntry{
		CorrelationId: cid,
		Entity:        reqres.SecretFetchRequest{},
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         event.Enter,
	}

	audit.Log(j)

	// Only workloads can fetch.
	if !validation.IsWorkload(spiffeid) {
		handle.BadSvidResponse(cid, w, spiffeid, j)
		return
	}

	log.DebugLn(&cid, "Fetch: sending response")

	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.InfoLn(&cid, "Fetch: Problem closing body")
		}
	}()

	log.DebugLn(&cid, "Fetch: preparing request")

	workloadId, parts := extract.WorkloadIDAndParts(spiffeid)
	if len(parts) == 0 {
		handle.BadPeerSvidResponse(cid, w, spiffeid, j)
		return
	}

	secret, err := state.ReadSecret(cid, workloadId)
	if err != nil {
		log.WarnLn(&cid, "Fetch: Problem reading secret", err.Error())
	}

	log.TraceLn(&cid, "Fetch: workloadId", workloadId)

	// If secret does not exist, send an empty response.
	if secret == nil {
		handle.NoSecretResponse(cid, w, j)
		return
	}

	log.DebugLn(&cid, "Fetch: will send. workload id:", workloadId)

	value := extract.SecretValue(cid, secret)

	// RFC3339 is what Go uses internally when marshaling dates.
	// Choosing it to be consistent.
	sfr := reqres.SecretFetchResponse{
		Data:    value,
		Created: fmt.Sprintf("\"%s\"", secret.Created.Format(time.RFC3339)),
		Updated: fmt.Sprintf("\"%s\"", secret.Updated.Format(time.RFC3339)),
	}

	handle.SuccessResponse(cid, w, j, sfr)
}
