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
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"io"
	"net/http"
	"strings"
)

func doList(cid string, w http.ResponseWriter, r *http.Request,
	spiffeid string, encrypted bool,
) {
	if env.SafeManualKeyInput() && !state.MasterKeySet() {
		log.InfoLn(&cid, "List: Master key not set")
		return
	}

	j := audit.JournalEntry{
		CorrelationId: cid,
		Entity:        reqres.SecretListRequest{},
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         audit.EventEnter,
	}

	audit.Log(j)

	// Only sentinel can list.
	if !isSentinel(j, cid, w, spiffeid) {
		return
	}

	log.TraceLn(&cid, "List: before defer")

	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.InfoLn(&cid, "List: Problem closing body")
		}
	}()

	log.TraceLn(&cid, "List: after defer")

	tmp := strings.Replace(spiffeid, env.SentinelSpiffeIdPrefix(), "", 1)
	parts := strings.Split(tmp, "/")
	if len(parts) == 0 {
		j.Event = audit.EventBadPeerSvid
		audit.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "List: Problem with spiffeid", spiffeid)
		}
		return
	}

	workloadId := parts[0]
	secrets := state.AllSecrets(cid)

	log.DebugLn(&cid, "List: will send. workload id:", workloadId)

	if encrypted {
		algo := "age"
		if env.SafeFipsCompliant() {
			algo = "aes"
		}

		secrets := state.AllSecretsEncrypted(cid)

		sfr := reqres.SecretEncryptedListResponse{
			Secrets:   secrets,
			Algorithm: algo,
		}

		j.Event = audit.EventOk
		j.Entity = sfr
		audit.Log(j)

		resp, err := json.Marshal(sfr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := io.WriteString(w, "List: Problem unmarshalling response")
			if err != nil {
				log.InfoLn(&cid, "List: Problem sending response", err.Error())
			}
			return
		}

		log.DebugLn(&cid, "List: before response")

		_, err = io.WriteString(w, string(resp))
		if err != nil {
			log.InfoLn(&cid, "List: Problem sending response", err.Error())
		}

		log.DebugLn(&cid, "List: after response")
		return
	}

	sfr := reqres.SecretListResponse{
		Secrets: secrets,
	}

	j.Event = audit.EventOk
	j.Entity = sfr
	audit.Log(j)

	resp, err := json.Marshal(sfr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := io.WriteString(w, "List: Problem unmarshalling response")
		if err != nil {
			log.InfoLn(&cid, "List: Problem sending response", err.Error())
		}
		return
	}

	log.DebugLn(&cid, "List: before response")

	_, err = io.WriteString(w, string(resp))
	if err != nil {
		log.InfoLn(&cid, "List: Problem sending response", err.Error())
	}

	log.DebugLn(&cid, "List: after response")
}

// List returns all registered workloads to the system with some metadata
// that is secure to share. For example, it returns secret names but not values.
//
// - cid: A string representing the client identifier.
// - w: An http.ResponseWriter used to write the HTTP response.
// - r: A pointer to an http.Request representing the received HTTP request.
// - spiffeid: spiffe id of the caller.
func List(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	doList(cid, w, r, spiffeid, false)
}

// ListEncrypted returns all registered workloads to the system. Similar to `List`
// it return meta information; however, it also returns encrypted secret values
// where an operator can decrypt if they know the VSecM root key.
//
// - cid: A string representing the client identifier.
// - w: An http.ResponseWriter used to write the HTTP response.
// - r: A pointer to an http.Request representing the received HTTP request.
// - spiffeid: spiffe id of the caller.
func ListEncrypted(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	doList(cid, w, r, spiffeid, true)
}
