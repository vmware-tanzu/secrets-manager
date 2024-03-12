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
	"net/http"
	"time"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/crypto"
	httq "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/http"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/journal"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/json"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/state"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/validation"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// Secret handles the creation, updating, and management of secrets.
// It performs several checks and operations based on the request parameters.
//
// Parameters:
//   - cid: A string representing the correlation ID for the request, used for
//     logging and tracking purposes.
//   - w: An http.ResponseWriter object used to send responses back to the client.
//   - r: An http.Request object containing the details of the client's request.
//   - spiffeid: A string representing the SPIFFE ID of the client making the request.
func Secret(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	if !state.RootKeySet() {
		log.InfoLn(&cid, "Secret: Root key not set")
		return
	}

	j := journal.CreateDefaultEntry(cid, spiffeid, r)
	audit.Log(j)

	if !validation.IsSentinel(j, cid, w, spiffeid) {
		return
	}

	log.DebugLn(&cid, "Secret: sentinel spiffeid:", spiffeid)

	body := httq.ReadBody(cid, r, w, j)
	if body == nil {
		j.Event = event.BadPayload
		audit.Log(j)

		return
	}

	log.DebugLn(&cid, "Secret: Parsed request body")

	ur := json.UnmarshalSecretUpsertRequest(cid, body, j, w)
	if ur == nil {
		j.Event = event.BadPayload
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
		// has a side effect of sending response.
		crypto.Encrypt(cid, value, j, w)

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
		j.Event = event.NoWorkloadId
		audit.Log(j)

		return
	}

	// `encrypt` means that the value is encrypted, so we need to decrypt it.
	if encrypt {
		v, failed := crypto.Decrypt(cid, value, j, w)
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

	state.Upsert(secretToStore, appendValue, workloadId, cid, j, w)
}
