/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package handle

import (
	"io"
	"net/http"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	deleteRoute "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/delete"
	fetchRoute "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/fetch"
	keystoneRoute "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/keystone"
	listRoute "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/list"
	receiveRoute "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/receive"
	secretRoute "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/secret"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

// InitializeRoutes initializes the HTTP routes for the web server. It sets up an
// HTTP handler function for the root URL ("/"). The handler uses the given
// X509Source to retrieve X.509 SVIDs for validating incoming connections.
//
// Parameters:
//   - source: A pointer to a `workloadapi.X509Source`, used to obtain X.509 SVIDs.
//
// Note: The InitializeRoutes function should be called only once, usually
// during server initialization.
func InitializeRoutes(source *workloadapi.X509Source) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cid, _ := crypto.RandomString(8)
		if cid == "" {
			cid = "VSECMFHN"
		}

		svid, err := source.GetX509SVID()
		if err != nil {
			log.FatalLn(
				&cid,
				"Unable to get X.509 SVID from source bundle", err.Error(),
			)
		}

		svidId := svid.ID
		if !validation.IsSafe(svidId.String()) {
			log.FatalLn(
				&cid,
				"SpiffeId check: I don't know you, and it's crazy:",
				svidId.String(),
			)
		}

		id, err := spiffeIdFromRequest(r)
		if err != nil {
			log.WarnLn(&cid, "Handler: blocking insecure svid", id, err)

			// Block insecure connection attempt.
			_, err = io.WriteString(w, "")
			if err != nil {
				log.InfoLn(&cid, "Problem writing response:", err.Error())
				return
			}
		}

		sid := id.String()
		p := r.URL.Path

		log.DebugLn(&cid, "Handler: got svid:", sid, "path", p, "method", r.Method)

		// Return the current state of the Keystone secret.
		// Either "initialized", or "pending"
		if r.Method == http.MethodGet && p == "/sentinel/v1/keystone" {
			log.DebugLn(&cid, "Handler: will keystone")
			keystoneRoute.Status(cid, w, r, sid)
		}

		// Route to list secrets.
		// Only VSecM Sentinel is allowed to call this API endpoint.
		// Calling it from anywhere else will error out.
		if r.Method == http.MethodGet && p == "/sentinel/v1/secrets" {
			log.DebugLn(&cid, "Handler: will list")
			listRoute.Masked(cid, w, r, sid)
			return
		}

		if r.Method == http.MethodGet && p == "/sentinel/v1/secrets?reveal=true" {
			log.DebugLn(&cid, "Handler: will list encrypted secrets")
			listRoute.Encrypted(cid, w, r, sid)
			return
		}

		// Route to add secrets to VSecM Safe.
		// Only VSecM Sentinel is allowed to call this API endpoint.
		// Calling it from anywhere else will error out.
		if r.Method == http.MethodPost && p == "/sentinel/v1/secrets" {
			log.DebugLn(&cid, "Handler:/sentinel/v1/secrets will secret")
			secretRoute.Secret(cid, w, r, sid)
			return
		}

		// Route to delete secrets from VSecM Safe.
		// Only VSecM Sentinel is allowed to call this API endpoint.
		// Calling it from anywhere else will error out.
		if r.Method == http.MethodDelete && p == "/sentinel/v1/secrets" {
			log.DebugLn(&cid, "Handler:/sentinel/v1/secrets will delete")
			deleteRoute.Delete(cid, w, r, sid)
			return
		}

		// Route to fetch secrets.
		// Only a VSecM-nominated workload is allowed to
		// call this API endpoint. Calling it from anywhere else will
		// error out.
		if r.Method == http.MethodGet && p == "/workload/v1/secrets" {
			log.DebugLn(&cid, "Handler:/workload/v1/secrets: will fetch")
			fetchRoute.Fetch(cid, w, r, sid)
			return
		}

		// Route to define the root key when VSECM_ROOT_KEY_INPUT_MODE_MANUAL is set.
		// Only VSecM Sentinel is allowed to call this API endpoint.
		// This method works only once. Once a key is set, there is no way to
		// update it. You will have to kill the VSecM Sentinel pod and restart it
		// to be able to set a new key.
		if r.Method == http.MethodPost && p == "/sentinel/v1/keys" {
			log.DebugLn(&cid, "Handler: will receive keys")
			receiveRoute.Keys(cid, w, r, sid)
			return
		}

		log.DebugLn(&cid, "Handler: route mismatch:", r.RequestURI)

		w.WriteHeader(http.StatusBadRequest)
		_, err = io.WriteString(w, "")
		if err != nil {
			log.WarnLn(&cid, "Problem writing response:", err.Error())
			return
		}
	})
}
