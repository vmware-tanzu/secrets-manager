/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package list

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/lib/validation"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/secret/collection"
	"github.com/vmware-tanzu/secrets-manager/core/audit/journal"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/v1/reqres/safe"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/spiffe"
)

func doList(
	cid string, w http.ResponseWriter, r *http.Request, encrypted bool,
) {
	spiffeid := spiffe.IdAsString(cid, r)

	if !crypto.RootKeySetInMemory() {
		log.InfoLn(&cid, "Masked: Root key not set")

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Masked: Problem with spiffeid", spiffeid)
		}

		return
	}

	j := journal.Entry{
		CorrelationId: cid,
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         event.Enter,
	}
	journal.Log(j)

	// Only sentinel can list.
	if ok, respond := validation.IsSentinel(j, cid, spiffeid); !ok {
		respond(w)
		return
	}

	log.TraceLn(&cid, "Masked: before defer")

	defer func(b io.ReadCloser) {
		err := b.Close()
		if err != nil {
			log.InfoLn(&cid, "Masked: Problem closing body")
		}
	}(r.Body)

	log.TraceLn(&cid, "Masked: after defer")

	tmp := strings.Replace(spiffeid, env.SpiffeIdPrefixForSentinel(), "", 1)
	parts := strings.Split(tmp, "/")

	if len(parts) == 0 {
		j.Event = event.BadPeerSvid
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Masked: Problem with spiffeid", spiffeid)
		}

		return
	}

	workloadId := parts[0]
	secrets := collection.AllSecrets(cid)

	log.DebugLn(&cid, "Masked: will send. workload id:", workloadId)

	if encrypted {
		algo := crypto.Age
		if env.FipsCompliantModeForSafe() {
			algo = crypto.Aes
		}

		secrets := collection.AllSecretsEncrypted(cid)

		sfr := reqres.SecretEncryptedListResponse{
			Secrets:   secrets,
			Algorithm: algo,
		}

		j.Event = event.Ok
		journal.Log(j)

		resp, err := json.Marshal(sfr)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := io.WriteString(w, "Masked: Problem marshalling response")
			if err != nil {
				log.ErrorLn(&cid,
					"Masked: Problem sending response", err.Error())
			}
			return
		}

		_, err = io.WriteString(w, string(resp))
		if err != nil {
			log.ErrorLn(&cid, "Masked: Problem sending response", err.Error())
		}

		log.DebugLn(&cid, "Masked: after response")
		return
	}

	sfr := reqres.SecretListResponse{
		Secrets: secrets,
	}

	j.Event = event.Ok
	journal.Log(j)

	resp, err := json.Marshal(sfr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := io.WriteString(w, "Masked: Problem marshalling response")
		if err != nil {
			log.ErrorLn(&cid, "Masked: Problem sending response", err.Error())
		}
		return
	}

	_, err = io.WriteString(w, string(resp))
	if err != nil {
		log.ErrorLn(&cid, "Masked: Problem sending response", err.Error())
	}

	log.DebugLn(&cid, "Masked: after response")
}
