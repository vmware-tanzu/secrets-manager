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

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/validation"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	data "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// Status handles HTTP requests to determine the current status of
// VSecM Keystone. The assumption is, if VSecM Keystone has an associated
// secret, then VSecM Sentinel will have finished its "init commands" flow
// successfully and will not need to re-run the init commands if it crashes
// or gets evicted.
//
// Parameters:
//   - cid: A unique identifier for the correlation of logs and audit entries.
//   - w: The http.ResponseWriter object through which HTTP responses are written.
//   - r: The http.Request received from the client. This contains all the details
//     about the request made by the client.
//   - spiffeid: The SPIFFE ID of the entity making the request, used for
//     authentication and logging.
//
// Note:
//   - This function is designed to be called by VSecM Sentinel (a trusted entity).
//   - Proper SPIFFE ID format and keystone initialization are crucial for the
//     correct execution of this function.
func Status(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	if !state.RootKeySet() {
		log.InfoLn(&cid, "Status: Root key not set")
		return
	}

	j := audit.JournalEntry{
		CorrelationId: cid,
		Entity:        reqres.KeystoneStatusRequest{},
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         event.Enter,
	}

	audit.Log(j)

	// Only sentinel can get the status.
	if !validation.IsSentinel(j, cid, w, spiffeid) {
		j.Event = event.BadSpiffeId
		audit.Log(j)
		return
	}

	log.TraceLn(&cid, "Status: before defer")

	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.InfoLn(&cid, "Status: Problem closing body")
		}
	}()

	log.TraceLn(&cid, "Status: after defer")

	tmp := strings.Replace(spiffeid, env.SpiffeIdPrefixForSentinel(), "", 1)
	parts := strings.Split(tmp, "/")

	if len(parts) == 0 {
		j.Event = event.BadPeerSvid
		audit.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Status: Problem with spiffeid", spiffeid)
		}

		return
	}

	initialized := state.KeystoneInitialized(cid)

	if initialized {
		log.TraceLn(&cid, "Status: keystone initialized")

		res := reqres.KeystoneStatusResponse{
			Status: data.Ready,
		}

		j.Event = event.Ok
		j.Entity = res
		audit.Log(j)

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

	j.Event = event.Ok
	j.Entity = res
	audit.Log(j)

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
