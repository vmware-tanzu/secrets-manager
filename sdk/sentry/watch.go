/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package sentry

import (
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"github.com/vmware-tanzu/secrets-manager/sdk/internal/timer"
	"time"
)

// Watch synchronizes the internal state of the sidecar by talking to
// VMware Secrets Manager Safe regularly. It periodically calls Fetch behind-the-scenes to
// get its work done. Once it fetches the secrets, it saves it to
// the location defined in the `VSECM_SIDECAR_SECRETS_PATH` environment
// variable (`/opt/vsecm/secrets.json` by default).
func Watch() {
	interval := timer.InitialInterval
	successCount := int64(0)
	errorCount := int64(0)

	cid, _ := crypto.RandomString(8)
	if cid == "" {
		cid = "VSECMSDK"
	}

	for {
		ticker := time.NewTicker(interval)
		select {
		case <-ticker.C:
			err := fetchSecrets()

			// Update parameters based on success/failure.
			interval, successCount, errorCount = timer.ExponentialBackoff(
				err == nil, interval, successCount, errorCount,
			)

			if err != nil {
				log.InfoLn(&cid, "Could not fetch secrets", err.Error(),
					". Will retry in", interval, ".")
			}

			ticker.Stop()
		}
	}
}
