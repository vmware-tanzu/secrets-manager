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
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/lib/backoff"
	"github.com/vmware-tanzu/secrets-manager/lib/crypto"
)

// Watch synchronizes the internal state of the sidecar by talking to
// VSecM Safe regularly. It periodically calls Fetch behind-the-scenes to
// get its work done. Once it fetches the secrets, it saves it to
// the location defined in the `VSECM_SIDECAR_SECRETS_PATH` environment
// variable (`/opt/vsecm/secrets.json` by default).
func Watch() {
	interval := env.PollIntervalForSidecar()

	cid, _ := crypto.RandomString(8)
	if cid == "" {
		panic("Unable to create a secure correlation id.")
	}

	for {
		_ = backoff.Retry("sentry.Watch", func() error {
			err := fetchSecrets()
			if err != nil {
				log.InfoLn(&cid, "Could not fetch secrets", err.Error(),
					". Will retry in", interval, ".")
			}
			return err
		}, backoff.Strategy{
			MaxRetries:  10,
			Delay:       interval,
			Exponential: false,
		})
	}
}
