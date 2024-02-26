/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package receive

import (
	"io"
	"net/http"
	"strings"

	httq "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/http"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/journal"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/json"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/validation"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func Keys(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	j := journal.CreateDefaultEntry(cid, spiffeid, r)
	j.Entity = reqres.KeyInputRequest{}
	audit.Log(j)

	if !validation.IsSentinel(j, cid, w, spiffeid) {
		j.Event = event.BadSpiffeId
		audit.Log(j)
		return
	}

	log.DebugLn(&cid, "Keys: sentinel spiffeid:", spiffeid)

	body := httq.ReadBody(cid, r, w, j)
	if body == nil {
		j.Event = event.BadPayload
		audit.Log(j)
		return
	}

	ur := json.UnmarshalKeyInputRequest(cid, body, j, w)
	if ur == nil {
		j.Event = event.BadPayload
		audit.Log(j)
		return
	}

	sr := *ur
	j.Entity = sr

	aesCipherKey := strings.TrimSpace(sr.AesCipherKey)
	agePrivateKey := strings.TrimSpace(sr.AgeSecretKey)
	agePublicKey := strings.TrimSpace(sr.AgePublicKey)

	if aesCipherKey == "" || agePrivateKey == "" || agePublicKey == "" {
		j.Event = event.BadPayload
		audit.Log(j)
		return
	}

	keysCombined := agePrivateKey + "\n" + agePublicKey + "\n" + aesCipherKey
	state.SetRootKey(keysCombined)

	log.DebugLn(&cid, "Keys: before response")

	_, err := io.WriteString(w, "OK")
	if err != nil {
		log.InfoLn(&cid, "Keys: Problem sending response", err.Error())
	}

	log.DebugLn(&cid, "Keys: after response")
}
