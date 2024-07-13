/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package secret

import (
	"io"
	"net/http"
	"time"

	httq "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/base/http"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/base/json"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/base/state"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/base/validation"
	"github.com/vmware-tanzu/secrets-manager/core/audit/journal"
	"github.com/vmware-tanzu/secrets-manager/core/constants/audit"
	"github.com/vmware-tanzu/secrets-manager/core/constants/val"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	data "github.com/vmware-tanzu/secrets-manager/lib/entity"
	s "github.com/vmware-tanzu/secrets-manager/lib/spiffe"
)

// Secret handles the creation, updating, and management of secrets.
// It performs several checks and operations based on the request parameters.
//
// Parameters:
//   - cid: A string representing the correlation ID for the request, used for
//     logging and tracking purposes.
//   - w: An http.ResponseWriter object used to send responses back to the
//     client.
//   - r: An http.Request object containing the details of the client's request.
//   - spiffeid: A string representing the SPIFFE ID of the client making the
//     request.
func Secret(cid string, r *http.Request, w http.ResponseWriter) {
	spiffeid := s.IdAsString(r)
	if spiffeid == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, val.NotOk)
		if err != nil {
			log.ErrorLn(&cid, "error writing response", err.Error())
		}

		return
	}

	if !crypto.RootKeySetInMemory() {
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, val.NotOk)
		if err != nil {
			log.ErrorLn(&cid, "error writing response", err.Error())
		}
		log.InfoLn(&cid, "Secret: Root key not set")

		return
	}

	j := journal.CreateDefaultEntry(cid, spiffeid, r)
	journal.Log(j)

	// Only sentinel can do this.
	if ok, respond := validation.IsSentinel(j, cid, spiffeid); !ok {
		j.Event = audit.NotSentinel
		journal.Log(j)
		respond(w)
		return
	}

	log.DebugLn(&cid, "Secret: sentinel spiffeid:", spiffeid)

	body, _ := httq.ReadBody(cid, r)
	if body == nil {
		j.Event = audit.BadPayload
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}

		return
	}

	log.DebugLn(&cid, "Secret: Parsed request body")

	ur, _ := json.UnmarshalSecretUpsertRequest(body)
	if ur == nil {
		j.Event = audit.BadPayload
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}

		return
	}

	sr := *ur

	workloadIds := sr.WorkloadIds
	value := sr.Value
	namespaces := sr.Namespaces
	template := sr.Template
	format := sr.Format
	encrypt := sr.Encrypt
	appendValue := sr.AppendValue
	notBefore := sr.NotBefore
	expiresAfter := sr.Expires

	if len(workloadIds) == 0 && encrypt {
		httq.SendEncryptedValue(cid, value, j, w)

		return
	}

	if len(namespaces) == 0 {
		namespaces = []string{"default"}
	}

	log.DebugLn(&cid, "Secret:Upsert: ", "workloadIds:", workloadIds,
		"namespaces:", namespaces,
		"template:", template, "format:", format, "encrypt:", encrypt,
		"appendValue:", appendValue,
		"notBefore:", notBefore, "expiresAfter:", expiresAfter)

	if len(workloadIds) == 0 && !encrypt {
		j.Event = audit.NoWorkloadId
		journal.Log(j)

		return
	}

	// `encrypt` means that the value is encrypted, so we need to decrypt it.
	if encrypt {
		log.TraceLn(&cid, "Secret: Value is encrypted")

		decrypted, err := crypto.DecryptValue(value)

		// If decryption failed, return an error response.
		if err != nil {
			log.InfoLn(&cid, "Secret: Decryption failed", err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			_, err := io.WriteString(w, "")
			if err != nil {
				log.InfoLn(&cid,
					"Secret: Problem sending response", err.Error())
			}

			return
		}

		// Update the value of the request to the decoded value.
		sr.Value = decrypted
		value = sr.Value
	} else {
		log.TraceLn(&cid, "Secret: Value is not encrypted")
	}

	nb := data.JsonTime{}
	exp := data.JsonTime{}

	if notBefore == "now" {
		nb = data.JsonTime(time.Now())
	} else {
		nbTime, err := time.Parse(time.RFC3339, notBefore)
		if err != nil {
			nb = data.JsonTime(time.Now())
		} else {
			nb = data.JsonTime(nbTime)
		}
	}

	if expiresAfter == "never" {
		// This is the largest time go stdlib can represent.
		// It is far enough into the future that the author does not care
		// what happens after.
		exp = data.JsonTime(
			time.Date(9999, time.December,
				31, 23, 59, 59, 999999999, time.UTC),
		)
	} else {
		expTime, err := time.Parse(time.RFC3339, expiresAfter)
		if err != nil {
			exp = data.JsonTime(
				time.Date(9999, time.December,
					31, 23, 59, 59, 999999999, time.UTC),
			)
		} else {
			exp = data.JsonTime(expTime)
		}
	}

	for _, workloadId := range workloadIds {
		secretToStore := entity.SecretStored{
			Name: workloadId,
			Meta: entity.SecretMeta{
				Namespaces:    namespaces,
				Template:      template,
				Format:        format,
				CorrelationId: cid,
			},
			Values:       []string{value},
			NotBefore:    time.Time(nb),
			ExpiresAfter: time.Time(exp),
		}

		state.Upsert(secretToStore, appendValue, workloadId, cid, j, w)
	}
}
