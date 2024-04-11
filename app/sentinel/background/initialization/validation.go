/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package initialization

import (
	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
	"github.com/vmware-tanzu/secrets-manager/core/backoff"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func initCommandsExecutedAlready(cid *string) bool {
	log.TraceLn(cid, "check:initCommandsExecutedAlready")

	s := backoffStrategy()
	for {
		err := backoff.Retry("RunInitCommands:CheckConnectivity", func() error {
			initialized, err := safe.CheckInitialization()
			if err != nil {
				return err
			}
			return initialized
		}, s)

		if err != nil {
			log.ErrorLn(cid, "check:backoff:error", err.Error())
		}
	}
}
