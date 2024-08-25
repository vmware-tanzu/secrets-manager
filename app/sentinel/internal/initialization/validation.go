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
	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/lib/backoff"
)

func (i *Initializer) initCommandsExecutedAlready(
	ctx context.Context, src *workloadapi.X509Source,
) bool {
	cid := ctx.Value(key.CorrelationId).(*string)

	i.Logger.TraceLn(cid, "check:initCommandsExecutedAlready")

	initialized := false

	err := backoff.RetryExponential(
		"RunInitCommands:CheckConnectivity",
		func() error {
			var err error
			initialized, err = i.Safe.CheckInitialization(ctx, src)
			return err
		})

	if err == nil {
		return initialized
	}

	// I shouldn't be here.
	panic("RunInitCommands" +
		":initCommandsExecutedAlready: failed to check command initialization")
}
