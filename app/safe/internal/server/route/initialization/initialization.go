/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package initialization

import (
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/internal/handle"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/k8s"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

// InitComplete is called when the Sentinel has completed its initialization
// process. It is responsible for marking the initialization process as
// complete in the Kubernetes cluster, by updating the value of a
// "vsecm-sentinel-init-tombstone" Secret in the "vsecm-system" namespace.
//
// See ./app/sentinel/internal/safe/post.go:PostInitializationComplete for the
// corresponding sentinel-side implementation.
func InitComplete(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	if !state.RootKeySet() {
		log.InfoLn(&cid, "InitComplete: Root key not set")
		return
	}

	j := audit.JournalEntry{
		CorrelationId: cid,
		Entity:        reqres.SentinelInitCompleteRequest{},
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         event.Enter,
	}

	audit.Log(j)

	if !validation.IsSentinel(spiffeid) {
		handle.BadSvidResponse(cid, w, spiffeid, j)
		return
	}

	log.DebugLn(&cid, "InitComplete: sending response")

	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.InfoLn(&cid, "InitComplete: Problem closing body")
		}
	}()

	log.DebugLn(&cid, "InitComplete: preparing request")

	err := k8s.MarkInitializationSecretAsCompleted()
	if err != nil {
		log.WarnLn(
			&cid,
			"InitComplete: Problem creating initialization secret",
			err.Error(),
		)
	}

	icr := reqres.SentinelInitCompleteResponse{}

	handle.InitCompleteSuccessResponse(cid, w, j, icr)
}
