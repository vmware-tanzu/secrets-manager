/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package startup

import (
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"github.com/vmware-tanzu/secrets-manager/sdk/sentry"
	"os"
	"time"
)

func initialized() bool {
	r, _ := sentry.Fetch()
	v := r.Data
	return v != ""
}

// Watch continuously polls the associated secret of the workload to exist.
// If the secret exists, and it is not empty, the function exits the init
// container with a success status code (0).
func Watch() {
	interval := env.InitContainerPollInterval()
	ticker := time.NewTicker(interval)

	cid, _ := crypto.RandomString(8)
	if cid == "" {
		cid = "VSECMSDK"
	}

	for {
		select {
		case <-ticker.C:
			log.InfoLn(&cid, "init:: tick")
			if initialized() {
				log.InfoLn(&cid, "initialized… exiting the init process")
				os.Exit(0)
			}
		}
	}
}
