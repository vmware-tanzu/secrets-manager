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
	routeFallback "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/fallback"
	"net/http"

	routeDelete "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/delete"
	routeFetch "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/fetch"
	routeKeystone "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/keystone"
	routeList "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/list"
	routeReceive "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/receive"
	routeSecret "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/secret"
	"github.com/vmware-tanzu/secrets-manager/core/constants/url"
)

type handler func(string, *http.Request, http.ResponseWriter)

func factory(p, m string) handler {
	switch {
	case m == http.MethodGet && p == url.SentinelKeystone:
		return routeKeystone.Status
	case m == http.MethodGet && p == url.SentinelSecretsWithReveal:
		return routeList.Encrypted
	case m == http.MethodPost && p == url.SentinelSecrets:
		return routeSecret.Secret
	// Route to delete secrets from VSecM Safe.
	// Only VSecM Sentinel is allowed to call this API endpoint.
	// Calling it from anywhere else will error out.
	case m == http.MethodDelete && p == url.SentinelSecrets:
		return routeDelete.Delete
	// Route to define the root key.
	// Only VSecM Sentinel is allowed to call this API endpoint.
	case m == http.MethodPost && p == url.SentinelKeys:
		return routeReceive.Keys
	// Route to fetch secrets.
	// Only a VSecM-nominated workload is allowed to
	// call this API endpoint. Calling it from anywhere else will
	// error out.
	case m == http.MethodGet && p == url.WorkloadSecrets:
		return routeFetch.Fetch
	case m == http.MethodPost && p == url.WorkloadSecrets:
		panic("routeWorkloadPostSecrets not implemented")
	default:
		return routeFallback.Fallback
	}
}

func route(
	cid string, r *http.Request, w http.ResponseWriter,
) {
	factory(r.URL.Path, r.Method)(cid, r, w)
}
