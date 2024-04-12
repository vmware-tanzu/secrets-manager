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
	"context"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
	"github.com/vmware-tanzu/secrets-manager/core/backoff"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func initCommandsExecutedAlready(ctx context.Context, src *workloadapi.X509Source) bool {
	cid := ctx.Value("correlationId").(*string)

	log.TraceLn(cid, "check:initCommandsExecutedAlready")

	s := backoffStrategy()
	initialized := false
	for {
		err := backoff.Retry("RunInitCommands:CheckConnectivity", func() error {
			i, err := safe.CheckInitialization(ctx, src)
			if err != nil {
				return err
			}
			initialized = i
			return nil
		}, s)

		if err != nil {
			log.ErrorLn(cid, "check:backoff:error", err.Error())
			continue
		}

		log.TraceLn(cid, "check:return", initialized)
		return initialized
	}
}
