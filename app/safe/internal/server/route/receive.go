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
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"io"
	"net/http"
	"strings"
)

func ReceiveKeys(cid string, w http.ResponseWriter, r *http.Request, svid string) {
	j := createDefaultJournalEntry(cid, svid, r)
	j.Entity = reqres.KeyInputRequest{}
	audit.Log(j)

	if !isSentinel(j, cid, w, svid) {
		j.Event = audit.EventBadSvid
		audit.Log(j)
		return
	}

	log.DebugLn(&cid, "ReceiveKeys: sentinel svid:", svid)

	body := readBody(cid, r, w, j)
	if body == nil {
		j.Event = audit.EventBadPayload
		audit.Log(j)
		return
	}

	ur := unmarshalKeyInputRequest(cid, body, j, w)
	if ur == nil {
		j.Event = audit.EventBadPayload
		audit.Log(j)
		return
	}

	sr := *ur
	j.Entity = sr

	aesCipherKey := strings.TrimSpace(sr.AesCipherKey)
	agePrivateKey := strings.TrimSpace(sr.AgeSecretKey)
	agePublicKey := strings.TrimSpace(sr.AgePublicKey)

	if aesCipherKey == "" || agePrivateKey == "" || agePublicKey == "" {
		j.Event = audit.EventBadPayload
		audit.Log(j)
		return
	}

	keysCombined := agePrivateKey + "\n" + agePublicKey + "\n" + aesCipherKey
	state.SetMasterKey(keysCombined)

	log.DebugLn(&cid, "ReceiveKeys: before response")

	_, err := io.WriteString(w, "OK")
	if err != nil {
		log.InfoLn(&cid, "ReceiveKeys: Problem sending response", err.Error())
	}

	log.DebugLn(&cid, "ReceiveKeys: after response")
}
