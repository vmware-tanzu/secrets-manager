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

	routeDelete "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/delete"
	routeFetch "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/fetch"
	routeKeystone "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/keystone"
	routeList "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/list"
	routeReceive "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/receive"
	routeSecret "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/secret"
	"github.com/vmware-tanzu/secrets-manager/core/constants/url"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func routeSentinelGetKeystone(
	cid string, r *http.Request, w http.ResponseWriter,
) bool {
	p := r.URL.Path
	m := r.Method

	// Return the current state of the Keystone secret.
	// Either "initialized", or "pending"
	if m == http.MethodGet && p == url.SentinelKeystone {
		log.DebugLn(&cid, "Handler:routeSentinelGetKeystone")
		routeKeystone.Status(cid, w, r)

		return true
	}

	return false
}

func routeSentinelGetSecrets(
	cid string, r *http.Request, w http.ResponseWriter,
) bool {
	p := r.URL.Path
	m := r.Method

	// Route to list secrets.
	// Only VSecM Sentinel is allowed to call this API endpoint.
	// Calling it from anywhere else will error out.
	if m == http.MethodGet && p == url.SentinelSecrets {
		log.DebugLn(&cid, "Handler:routeSentinelGetSecrets")
		routeList.Masked(cid, w, r)

		return true
	}

	return false
}

func routeSentinelGetSecretsReveal(
	cid string, r *http.Request, w http.ResponseWriter,
) bool {
	p := r.URL.Path
	m := r.Method

	if m == http.MethodGet && p == url.SentinelSecretsWithReveal {
		log.DebugLn(&cid, "Handler:routeSentinelGetSecretsReveal")
		routeList.Encrypted(cid, w, r)

		return true
	}

	return false
}

func routeSentinelPostSecrets(
	cid string, r *http.Request, w http.ResponseWriter,
) bool {
	p := r.URL.Path
	m := r.Method

	// Route to add secrets to VSecM Safe.
	// Only VSecM Sentinel is allowed to call this API endpoint.
	// Calling it from anywhere else will error out.
	if m == http.MethodPost && p == url.SentinelSecrets {
		log.DebugLn(&cid, "Handler:routeSentinelPostSecrets")
		routeSecret.Secret(cid, w, r)

		return true
	}

	return false
}

func routeSentinelDeleteSecrets(
	cid string, r *http.Request, w http.ResponseWriter,
) bool {
	p := r.URL.Path
	m := r.Method

	// Route to delete secrets from VSecM Safe.
	// Only VSecM Sentinel is allowed to call this API endpoint.
	// Calling it from anywhere else will error out.
	if m == http.MethodDelete && p == url.SentinelSecrets {
		log.DebugLn(&cid, "Handler:routeSentinelDeleteSecrets")
		routeDelete.Delete(cid, w, r)

		return true
	}

	return false
}

func routeSentinelPostKeys(
	cid string, r *http.Request, w http.ResponseWriter,
) bool {
	p := r.URL.Path
	m := r.Method

	// Route to define the root key.
	// Only VSecM Sentinel is allowed to call this API endpoint.
	if m == http.MethodPost && p == url.SentinelKeys {
		log.DebugLn(&cid, "Handler:routeSentinelPostKeys")
		routeReceive.Keys(cid, w, r)

		return true
	}

	return false
}

func routeWorkloadGetSecrets(
	cid string, r *http.Request, w http.ResponseWriter,
) bool {
	p := r.URL.Path
	m := r.Method

	// Route to fetch secrets.
	// Only a VSecM-nominated workload is allowed to
	// call this API endpoint. Calling it from anywhere else will
	// error out.
	if m == http.MethodGet && p == url.WorkloadSecrets {
		log.DebugLn(&cid, "Handler:routeWorkloadGetSecrets")
		routeFetch.Fetch(cid, w, r)

		return true
	}

	return false
}

func routeWorkloadPostSecrets(
	cid string, r *http.Request, w http.ResponseWriter,
) bool {
	log.DebugLn(&cid,
		"Handler:routeWorkloadPostSecrets: will post", r.Method, r.URL.Path)

	panic("routeWorkloadPostSecrets not implemented")
}

func routeFallback(
	cid string, r *http.Request, w http.ResponseWriter,
) {
	log.DebugLn(&cid, "Handler: route mismatch:", r.RequestURI)

	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "")
	if err != nil {
		log.WarnLn(&cid, "Problem writing response:", err.Error())
	}
}
