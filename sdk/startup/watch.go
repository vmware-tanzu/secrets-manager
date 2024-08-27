/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package startup

import (
	"os"
	"time"

	"github.com/vmware-tanzu/secrets-manager/sdk/core/env"
	log "github.com/vmware-tanzu/secrets-manager/sdk/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/sdk/lib/crypto"
)

// Watch continuously polls the associated secret of the workload to exist.
// If the secret exists, and it is not empty, the function exits the init
// container with a success status code (0).
//
//   - waitTimeBeforeExit: The duration to wait before a successful exit from
//     the function.
func Watch(waitTimeBeforeExit time.Duration) {
	interval := env.PollIntervalForInitContainer()
	ticker := time.NewTicker(interval)

	cid, _ := crypto.RandomString(8)
	if cid == "" {
		panic("Unable to create a secure correlation id.")
	}

	for {
		select {
		case <-ticker.C:
			log.InfoLn(&cid, "init:: tick")
			if initialized() {
				log.InfoLn(&cid, "initialized... exiting the init process")

				time.Sleep(waitTimeBeforeExit)

				os.Exit(0)
			}
		}
	}
}
