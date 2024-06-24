/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package bootstrap

import (
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/queue"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/probe"
)

func completeInitialization(correlationId *string) {
	queue.Initialize()
	log.DebugLn(correlationId, "Creating readiness probe.")

	<-probe.CreateReadiness()

	log.AuditLn(correlationId, "VSecM Safe is ready to serve.")
}
