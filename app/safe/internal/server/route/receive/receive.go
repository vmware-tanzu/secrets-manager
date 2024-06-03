/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package receive

import (
	"io"
	"net/http"
	"strings"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/bootstrap"
	httq "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/http"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/json"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/validation"
	"github.com/vmware-tanzu/secrets-manager/core/audit/journal"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/spiffe"
)

// Keys processes a request to set root cryptographic keys within the application,
// validating the SPIFFE ID of the requester and the payload structure before
// proceeding.
//
// This function is pivotal in scenarios where updating the application's
// cryptographic foundation is required, often performed by a trusted
// VSecM Sentinel entity.
//
// The returned keys need to be protected and kept secret, as they are the
// foundation for the cryptographic operations within the application. The keys
// are used to encrypt and decrypt secrets, and to sign and verify the integrity
// of the data.
//
// Parameters:
//   - cid (string): Correlation ID for operation tracing and logging.
//   - w (http.ResponseWriter): The HTTP response writer to send back responses
//     or errors.
//   - r (*http.Request): The incoming HTTP request containing the payload.
//   - spiffeid (string): The SPIFFE ID associated with the requester, used for
//     authorization validation.
func Keys(cid string, w http.ResponseWriter, r *http.Request) {
	spiffeid := spiffe.IdAsString(cid, r)

	j := journal.CreateDefaultEntry(cid, spiffeid, r)
	journal.Log(j)

	// Only sentinel can set keys.
	if ok, respond := validation.IsSentinel(j, cid, spiffeid); !ok {
		respond(w)

		j.Event = event.BadSpiffeId
		journal.Log(j)

		return
	}

	log.DebugLn(&cid, "Keys: sentinel spiffeid:", spiffeid)

	body, _ := httq.ReadBody(cid, r)
	if body == nil {
		j.Event = event.BrokenBody
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Keys: Problem sending response", err.Error())
		}

		return
	}

	ur, _ := json.UnmarshalKeyInputRequest(body)
	if ur == nil {
		j.Event = event.BadPayload
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Keys: Problem sending response", err.Error())
		}

		return
	}

	sr := *ur

	aesCipherKey := strings.TrimSpace(sr.AesCipherKey)
	agePrivateKey := strings.TrimSpace(sr.AgeSecretKey)
	agePublicKey := strings.TrimSpace(sr.AgePublicKey)

	if aesCipherKey == "" || agePrivateKey == "" || agePublicKey == "" {
		j.Event = event.BadPayload
		journal.Log(j)
		return
	}

	keysCombined := agePrivateKey + "\n" + agePublicKey + "\n" + aesCipherKey
	crypto.SetRootKeyInMemory(keysCombined)

	if err := bootstrap.PersistRootKeysToRootKeyBackingStore(
		crypto.RootKeyCollection{
			PrivateKey: agePrivateKey,
			PublicKey:  agePublicKey,
			AesSeed:    aesCipherKey,
		},
	); err != nil {
		log.ErrorLn(&cid, "Keys: Problem persisting keys", err.Error())
	}

	log.DebugLn(&cid, "Keys: before response")

	_, err := io.WriteString(w, "OK")
	if err != nil {
		log.InfoLn(&cid, "Keys: Problem sending response", err.Error())
	}

	log.DebugLn(&cid, "Keys: after response")
}
