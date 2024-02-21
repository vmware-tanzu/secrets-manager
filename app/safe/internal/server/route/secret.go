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
	"time"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func createDefaultJournalEntry(cid, spiffeid string,
	r *http.Request) audit.JournalEntry {
	return audit.JournalEntry{
		CorrelationId: cid,
		Entity:        reqres.SecretFetchRequest{},
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         audit.EventEnter,
	}
}

func readBody(cid string, r *http.Request, w http.ResponseWriter,
	j audit.JournalEntry) []byte {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		j.Event = audit.EventBrokenBody
		audit.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err2 := io.WriteString(w, "")
		if err2 != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err2.Error())
		}
		return nil
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

	return body
}

func unmarshalRequest(cid string, body []byte, j audit.JournalEntry,
	w http.ResponseWriter) *reqres.SecretUpsertRequest {
	var sr reqres.SecretUpsertRequest
	err := json.Unmarshal(body, &sr)
	if err != nil {
		j.Event = audit.EventRequestTypeMismatch
		audit.Log(j)
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}
		return nil
	}
	return &sr
}

func unmarshalKeyInputRequest(cid string, body []byte, j audit.JournalEntry,
	w http.ResponseWriter) *reqres.KeyInputRequest {
	var sr reqres.KeyInputRequest
	err := json.Unmarshal(body, &sr)
	if err != nil {
		j.Event = audit.EventRequestTypeMismatch
		audit.Log(j)
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}
		return nil
	}
	return &sr
}

func encryptValue(cid string, value string, j audit.JournalEntry,
	w http.ResponseWriter) {
	if value == "" {
		j.Event = audit.EventNoValue
		audit.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}
		return
	}

	encrypted, err := state.EncryptValue(value)
	if err != nil {
		j.Event = audit.EventEncryptionFailed
		audit.Log(j)

		w.WriteHeader(http.StatusInternalServerError)
		_, err2 := io.WriteString(w, "")
		if err2 != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err2.Error())
		}
		return
	}

	_, err = io.WriteString(w, encrypted)
	if err != nil {
		log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
	}
	return
}

func decryptValue(cid string, value string, j audit.JournalEntry,
	w http.ResponseWriter) (string, bool) {
	decrypted, err := state.DecryptValue(value)
	if err != nil {
		j.Event = audit.EventDecryptionFailed
		audit.Log(j)

		w.WriteHeader(http.StatusInternalServerError)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}
		return "", true
	}

	return decrypted, false
}

func upsert(secretToStore entity.SecretStored,
	appendValue bool, workloadId string, cid string,
	j audit.JournalEntry, w http.ResponseWriter,
) {
	state.UpsertSecret(secretToStore, appendValue)
	log.DebugLn(&cid, "Secret:UpsertEnd: workloadId", workloadId)

	j.Event = audit.EventOk
	audit.Log(j)

	_, err := io.WriteString(w, "OK")
	if err != nil {
		log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
	}
}

func Secret(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	if env.SafeManualKeyInput() && !state.MasterKeySet() {
		log.InfoLn(&cid, "Secret: Master key not set")
		return
	}

	j := createDefaultJournalEntry(cid, spiffeid, r)
	audit.Log(j)

	if !isSentinel(j, cid, w, spiffeid) {
		j.Event = audit.EventBadSvid
		audit.Log(j)
		return
	}

	log.DebugLn(&cid, "Secret: sentinel spiffeid:", spiffeid)

	body := readBody(cid, r, w, j)
	if body == nil {
		j.Event = audit.EventBadPayload
		audit.Log(j)
		return
	}

	log.DebugLn(&cid, "Secret: Parsed request body")

	ur := unmarshalRequest(cid, body, j, w)
	if ur == nil {
		j.Event = audit.EventBadPayload
		audit.Log(j)
		return
	}

	sr := *ur

	j.Entity = sr

	workloadId := sr.WorkloadId
	value := sr.Value
	backingStore := sr.BackingStore
	useK8s := sr.UseKubernetes
	namespaces := sr.Namespaces
	template := sr.Template
	format := sr.Format
	encrypt := sr.Encrypt
	appendValue := sr.AppendValue
	notBefore := sr.NotBefore
	expiresAfter := sr.Expires

	if workloadId == "" && encrypt {
		// has side effect of sending response.
		encryptValue(cid, value, j, w)
		return
	}

	if len(namespaces) == 0 {
		namespaces = []string{"default"}
	}

	log.DebugLn(&cid, "Secret:Upsert: ", "workloadId:", workloadId,
		"namespaces:", namespaces, "backingStore:", backingStore,
		"template:", template, "format:", format, "encrypt:", encrypt,
		"appendValue:", appendValue, "useK8s", useK8s,
		"notBefore:", notBefore, "expiresAfter:", expiresAfter)

	if workloadId == "" && !encrypt {
		j.Event = audit.EventNoWorkloadId
		audit.Log(j)

		return
	}

	// `encrypt` means that the value is encrypted, so we need to decrypt it.
	if encrypt {
		v, failed := decryptValue(cid, value, j, w)
		value = v

		// If decryption failed, we already sent the response.
		if failed {
			return
		}
	}

	nb := entity.JsonTime{}
	exp := entity.JsonTime{}

	if notBefore == "now" {
		nb = entity.JsonTime(time.Now())
	} else {
		nbTime, err := time.Parse(time.RFC3339, notBefore)
		if err != nil {
			nb = entity.JsonTime(time.Now())
		} else {
			nb = entity.JsonTime(nbTime)
		}
	}

	if expiresAfter == "never" {
		// This is the largest time go std. lib can represent.
		// It is far enough into the future that the author does not care
		// what happens after.
		exp = entity.JsonTime(
			time.Date(9999, time.December, 31, 23, 59, 59, 999999999, time.UTC),
		)
	} else {
		expTime, err := time.Parse(time.RFC3339, expiresAfter)
		if err != nil {
			exp = entity.JsonTime(
				time.Date(9999, time.December, 31, 23, 59, 59, 999999999, time.UTC),
			)
		} else {
			exp = entity.JsonTime(expTime)
		}
	}

	secretToStore := entity.SecretStored{
		Name: workloadId,
		Meta: entity.SecretMeta{
			UseKubernetesSecret: useK8s,
			Namespaces:          namespaces,
			BackingStore:        backingStore,
			Template:            template,
			Format:              format,
			CorrelationId:       cid,
		},
		Values:       []string{value},
		NotBefore:    time.Time(nb),
		ExpiresAfter: time.Time(exp),
	}

	upsert(secretToStore, appendValue, workloadId, cid, j, w)
}
