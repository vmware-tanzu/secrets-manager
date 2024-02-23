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
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"io"
	"net/http"

	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func handleBadSvidResponse(cid string, w http.ResponseWriter, spiffeid string,
	j audit.JournalEntry,
) {
	j.Event = event.BadSpiffeId
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
	j.Event = event.BadPeerSvid
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
	j.Event = event.NoSecret
	audit.Log(j)

	w.WriteHeader(http.StatusNotFound)
	_, err2 := io.WriteString(w, "")
	if err2 != nil {
		log.InfoLn(&cid, "Fetch: Problem sending response", err2.Error())
	}
}

func handleInitCompleteSuccessResponse(cid string, w http.ResponseWriter,
	j audit.JournalEntry, sfr reqres.SentinelInitCompleteResponse) {
	j.Event = event.Ok
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
	j.Event = event.Ok
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
