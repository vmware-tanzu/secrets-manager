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
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func routeSentinelGetKeystone(cid string, r *http.Request, w http.ResponseWriter, sid string) bool {
	p := r.URL.Path
	m := r.Method

	// Return the current state of the Keystone secret.
	// Either "initialized", or "pending"
	if m == http.MethodGet && p == "/sentinel/v1/keystone" {
		log.DebugLn(&cid, "Handler: will keystone")
		routeKeystone.Status(cid, w, r, sid)
		return true
	}
	return false
}

func routeSentinelGetSecrets(cid string, r *http.Request, w http.ResponseWriter, sid string) bool {
	p := r.URL.Path
	m := r.Method

	// Route to list secrets.
	// Only VSecM Sentinel is allowed to call this API endpoint.
	// Calling it from anywhere else will error out.
	if m == http.MethodGet && p == "/sentinel/v1/secrets" {
		log.DebugLn(&cid, "Handler: will list")
		routeList.Masked(cid, w, r, sid)
		return true
	}
	return false
}

func routeSentinelGetSecretsReveal(cid string, r *http.Request, w http.ResponseWriter, sid string) bool {
	p := r.URL.Path
	m := r.Method

	if m == http.MethodGet && p == "/sentinel/v1/secrets?reveal=true" {
		log.DebugLn(&cid, "Handler: will list encrypted secrets")
		routeList.Encrypted(cid, w, r, sid)
		return true
	}
	return false
}

func routeSentinelPostSecrets(cid string, r *http.Request, w http.ResponseWriter, sid string) bool {
	p := r.URL.Path
	m := r.Method

	// Route to add secrets to VSecM Safe.
	// Only VSecM Sentinel is allowed to call this API endpoint.
	// Calling it from anywhere else will error out.
	if m == http.MethodPost && p == "/sentinel/v1/secrets" {
		log.DebugLn(&cid, "Handler:/sentinel/v1/secrets will secret")
		routeSecret.Secret(cid, w, r, sid)
		return false
	}
	return true
}

func routeSentinelDeleteSecrets(cid string, r *http.Request, w http.ResponseWriter, sid string) bool {
	p := r.URL.Path
	m := r.Method

	// Route to delete secrets from VSecM Safe.
	// Only VSecM Sentinel is allowed to call this API endpoint.
	// Calling it from anywhere else will error out.
	if m == http.MethodDelete && p == "/sentinel/v1/secrets" {
		log.DebugLn(&cid, "Handler:/sentinel/v1/secrets will delete")
		routeDelete.Delete(cid, w, r, sid)
		return false
	}
	return true
}

func routeSentinelPostKeys(cid string, r *http.Request, w http.ResponseWriter, sid string) bool {
	p := r.URL.Path
	m := r.Method

	// Route to define the root key.
	// Only VSecM Sentinel is allowed to call this API endpoint.
	if m == http.MethodPost && p == "/sentinel/v1/keys" {
		log.DebugLn(&cid, "Handler: will receive keys")
		routeReceive.Keys(cid, w, r, sid)
		return false
	}
	return true
}

func routeWorkloadGetSecrets(cid string, r *http.Request, w http.ResponseWriter, sid string) bool {
	p := r.URL.Path
	m := r.Method

	// Route to fetch secrets.
	// Only a VSecM-nominated workload is allowed to
	// call this API endpoint. Calling it from anywhere else will
	// error out.
	if m == http.MethodGet && p == "/workload/v1/secrets" {
		log.DebugLn(&cid, "Handler:/workload/v1/secrets: will fetch")
		routeFetch.Fetch(cid, w, r, sid)
		return false
	}
	return true
}

func routeWorkloadPostSecrets(cid string, r *http.Request, w http.ResponseWriter, sid string) bool {
	log.DebugLn(&cid, "Handler:/workload/v1/secrets: will post", r.Method, r.URL.Path, sid, w)

	panic("routeWorkloadPostSecrets not implemented")
}

func routeFallback(cid string, r *http.Request, w http.ResponseWriter, sid string) {
	log.DebugLn(&cid, "Handler: route mismatch:", r.RequestURI, sid)

	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "")
	if err != nil {
		log.WarnLn(&cid, "Problem writing response:", err.Error())
	}
}
