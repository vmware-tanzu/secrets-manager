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
	"net/http"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

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
		cid := crypto.Id()

		validation.EnsureSafe(source)

		id, err := spiffeIdFromRequest(r)

		if err != nil {
			log.WarnLn(&cid, "Handler: blocking insecure svid", id, err)
			return
		}

		sid := id.String()
		p := r.URL.Path
		m := r.Method
		log.DebugLn(&cid, "Handler: got svid:", sid, "path", p, "method", m)

		switch {
		case routeSentinelGetKeystone(cid, r, w, sid):
			return
		case routeSentinelGetSecrets(cid, r, w, sid):
			return
		case routeSentinelGetSecretsReveal(cid, r, w, sid):
			return
		case routeSentinelPostSecrets(cid, r, w, sid):
			return
		case routeSentinelDeleteSecrets(cid, r, w, sid):
			return
		case routeSentinelPostKeys(cid, r, w, sid):
			return
		case routeWorkloadGetSecrets(cid, r, w, sid):
			return
		case routeWorkloadPostSecrets(cid, r, w, sid):
			return
		}

		routeFallback(cid, r, w, sid)
	})
}
