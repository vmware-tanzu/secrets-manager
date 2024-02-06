/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package route

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	v1 "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

func handleBadSvidResponse(cid string, w http.ResponseWriter, spiffeid string,
	j audit.JournalEntry,
) {
	j.Event = audit.EventBadSvid
	audit.Log(j)

	log.DebugLn(&cid, "Fetch: bad spiffeid", spiffeid)

	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "")
	if err != nil {
		log.InfoLn(&cid, "Fetch: Problem sending response", err.Error())
	}
}

func handleBadPeerSvidResponse(cid string, w http.ResponseWriter,
	spiffeid string, j audit.JournalEntry,
) {
	j.Event = audit.EventBadPeerSvid
	audit.Log(j)

	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "")
	if err != nil {
		log.InfoLn(&cid, "Fetch: Problem with spiffeid", spiffeid)
	}
}

func handleNoSecretResponse(cid string, w http.ResponseWriter,
	j audit.JournalEntry,
) {
	j.Event = audit.EventNoSecret
	audit.Log(j)

	w.WriteHeader(http.StatusNotFound)
	_, err2 := io.WriteString(w, "")
	if err2 != nil {
		log.InfoLn(&cid, "Fetch: Problem sending response", err2.Error())
	}
}

func handleInitCompleteSuccessResponse(cid string, w http.ResponseWriter,
	j audit.JournalEntry, sfr reqres.SentinelInitCompleteResponse) {
	j.Event = audit.EventOk
	j.Entity = sfr
	audit.Log(j)

	resp, err := json.Marshal(sfr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err2 := io.WriteString(w, "Problem unmarshalling response")
		if err2 != nil {
			log.InfoLn(&cid, "Problem sending response", err2.Error())
		}
		return
	}

	log.DebugLn(&cid, "before response")

	_, err = io.WriteString(w, string(resp))
	if err != nil {
		log.InfoLn(&cid, "Problem sending response", err.Error())
	}

	log.DebugLn(&cid, "after response")
}

func handleSuccessResponse(cid string, w http.ResponseWriter,
	j audit.JournalEntry, sfr reqres.SecretFetchResponse) {
	j.Event = audit.EventOk
	j.Entity = sfr
	audit.Log(j)

	resp, err := json.Marshal(sfr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err2 := io.WriteString(w, "Problem unmarshalling response")
		if err2 != nil {
			log.InfoLn(&cid, "Fetch: Problem sending response", err2.Error())
		}
		return
	}

	log.DebugLn(&cid, "Fetch: before response")

	_, err = io.WriteString(w, string(resp))
	if err != nil {
		log.InfoLn(&cid, "Problem sending response", err.Error())
	}

	log.DebugLn(&cid, "Fetch: after response")
}

func getWorkloadIDAndParts(spiffeid string) (string, []string) {
	tmp := strings.Replace(spiffeid, env.WorkloadSpiffeIdPrefix(), "", 1)
	parts := strings.Split(tmp, "/")
	if len(parts) > 0 {
		return parts[0], parts
	}
	return "", nil
}

func getSecretValue(cid string, secret *v1.SecretStored) string {
	if secret.ValueTransformed != "" {
		log.TraceLn(&cid, "Fetch: using transformed value")
		return secret.ValueTransformed
	}

	// This part is for backwards compatibility.
	// It probably won’t execute because `secret.ValueTransformed` will
	// always be set.

	log.TraceLn(&cid, "Fetch: using raw value")

	if len(secret.Values) == 1 {
		return secret.Values[0]
	}

	jsonData, err := json.Marshal(secret.Values)
	if err != nil {
		log.WarnLn(&cid, "Fetch: Problem marshaling values", err.Error())
	} else {
		return string(jsonData)
	}

	return ""
}

func Fetch(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	if env.SafeManualKeyInput() && !state.MasterKeySet() {
		log.InfoLn(&cid, "Fetch: Master key not set")
		return
	}

	j := audit.JournalEntry{
		CorrelationId: cid,
		Entity:        reqres.SecretFetchRequest{},
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         audit.EventEnter,
	}

	audit.Log(j)

	// Only workloads can fetch.
	if !validation.IsWorkload(spiffeid) {
		handleBadSvidResponse(cid, w, spiffeid, j)
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

	workloadId, parts := getWorkloadIDAndParts(spiffeid)
	if len(parts) == 0 {
		handleBadPeerSvidResponse(cid, w, spiffeid, j)
		return
	}

	secret, err := state.ReadSecret(cid, workloadId)
	if err != nil {
		log.WarnLn(&cid, "Fetch: Problem reading secret", err.Error())
	}

	log.TraceLn(&cid, "Fetch: workloadId", workloadId)

	// If secret does not exist, send an empty response.
	if secret == nil {
		handleNoSecretResponse(cid, w, j)
		return
	}

	log.DebugLn(&cid, "Fetch: will send. workload id:", workloadId)

	value := getSecretValue(cid, secret)

	// RFC3339 is what Go uses internally when marshaling dates.
	// Choosing it to be consistent.
	sfr := reqres.SecretFetchResponse{
		Data:    value,
		Created: fmt.Sprintf("\"%s\"", secret.Created.Format(time.RFC3339)),
		Updated: fmt.Sprintf("\"%s\"", secret.Updated.Format(time.RFC3339)),
	}

	handleSuccessResponse(cid, w, j, sfr)
}
