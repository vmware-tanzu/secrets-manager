/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package keystone

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/base/validation"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/secret/collection"
	"github.com/vmware-tanzu/secrets-manager/core/audit/journal"
	"github.com/vmware-tanzu/secrets-manager/core/constants/audit"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/v1/reqres/safe"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	s "github.com/vmware-tanzu/secrets-manager/lib/spiffe"
)

// Status handles HTTP requests to determine the current status of
// VSecM Keystone. The assumption is, if VSecM Keystone has an associated
// secret, then VSecM Sentinel will have finished its "init commands" flow
// successfully and will not need to re-run the init commands if it crashes
// or gets evicted.
//
// Parameters:
//   - cid: A unique identifier for the correlation of logs and audit entries.
//   - w: The http.ResponseWriter object through which HTTP responses are
//     written.
//   - r: The http.Request received from the client. This contains all the
//     details about the request made by the client.
//   - spiffeid: The SPIFFE ID of the entity making the request, used for
//     authentication and logging.
func Status(
	cid string, r *http.Request, w http.ResponseWriter,
) {
	spiffeid := s.IdAsString(r)

	j := data.JournalEntry{
		CorrelationId: cid,
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         audit.Enter,
	}

	journal.Log(j)

	if spiffeid == "" {
		log.InfoLn(&cid, "Status: Bad SPIFFE ID")
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.ErrorLn(&cid, "Status: Problem sending response", err.Error())
		}

		j.Event = audit.BadSpiffeId
		journal.Log(j)

		return
	}

	if !crypto.RootKeySetInMemory() {
		log.InfoLn(&cid, "Status: Root key not set")
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.ErrorLn(&cid, "Status: Problem sending response", err.Error())
		}

		j.Event = audit.RootKeyNotSet
		journal.Log(j)

		return
	}

	// Only sentinel can get the status.
	if ok, respond := validation.IsSentinel(j, cid, spiffeid); !ok {
		respond(w)

		j.Event = audit.BadSpiffeId
		journal.Log(j)

		return
	}

	log.TraceLn(&cid, "Status: before defer")

	defer func(b io.ReadCloser) {
		err := b.Close()
		if err != nil {
			log.InfoLn(&cid, "Status: Problem closing body")
		}
	}(r.Body)

	log.TraceLn(&cid, "Status: after defer")

	tmp := strings.Replace(spiffeid, env.SpiffeIdPrefixForSentinel(), "", 1)
	parts := strings.Split(tmp, "/")

	if len(parts) == 0 {
		j.Event = audit.BadPeerSvid
		journal.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Status: Problem with spiffeid", spiffeid)
		}

		return
	}

	if collection.KeystoneInitialized(cid) {
		log.TraceLn(&cid, "Status: keystone initialized")

		res := reqres.KeystoneStatusResponse{
			Status: data.Ready,
		}

		j.Event = audit.Ok
		journal.Log(j)

		resp, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, err := io.WriteString(w, "Status: Problem marshalling response")
			if err != nil {
				log.ErrorLn(&cid, "Status: Problem sending response", err.Error())
			}
			return
		}

		_, err = io.WriteString(w, string(resp))
		if err != nil {
			log.ErrorLn(&cid, "Status: Problem sending response", err.Error())
		}

		log.DebugLn(&cid, "Status: after response")
		return
	}

	// Below: not initialized

	log.TraceLn(&cid, "Status: keystone not initialized")

	res := reqres.KeystoneStatusResponse{
		Status: data.Pending,
	}

	j.Event = audit.Ok
	journal.Log(j)

	resp, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := io.WriteString(w, "Status: Problem marshalling response")
		if err != nil {
			log.ErrorLn(&cid, "Status: Problem sending response", err.Error())
		}
		return
	}

	_, err = io.WriteString(w, string(resp))
	if err != nil {
		log.ErrorLn(&cid, "Status: Problem sending response", err.Error())
	}

	log.DebugLn(&cid, "Status: after response")
}
