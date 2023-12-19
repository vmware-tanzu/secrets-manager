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
	"github.com/vmware-tanzu/secrets-manager/sdk/sentry"
)

func main() {
	id := "AEGSSDCR"
	log.InfoLn(&id, "Starting VSecM Sidecar")
	go sentry.Watch()
	// Keep the main routine alive:
	system.KeepAlive()
}
