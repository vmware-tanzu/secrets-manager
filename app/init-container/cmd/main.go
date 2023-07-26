/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"github.com/vmware-tanzu/secrets-manager/core/system"
	"github.com/vmware-tanzu/secrets-manager/sdk/startup"
)

func main() {
	id := "AEGIICNT"

	log.InfoLn(&id, "Starting VMware Secrets Manager Init Container")
	go startup.Watch()

	// Block the process from exiting, but also be graceful and honor the
	// termination signals that may come from the orchestrator.
	system.KeepAlive()
}
