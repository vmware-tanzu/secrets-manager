/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package sentry

import (
	"github.com/vmware-tanzu/secrets-manager/core/backoff"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"time"
)

// Watch synchronizes the internal state of the sidecar by talking to
// VSecM Safe regularly. It periodically calls Fetch behind-the-scenes to
// get its work done. Once it fetches the secrets, it saves it to
// the location defined in the `VSECM_SIDECAR_SECRETS_PATH` environment
// variable (`/opt/vsecm/secrets.json` by default).
func Watch() {
	interval := backoff.InitialInterval
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
			interval, successCount, errorCount = backoff.ExponentialBackoff(
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
