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
	"errors"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/lib/backoff"
)

func (i *Initializer) ensureApiConnectivity(ctx context.Context, cid *string) {
	i.Logger.TraceLn(cid, "Before checking api connectivity")

	err := backoff.RetryExponential(
		"RunInitCommands:CheckConnectivity",
		func() error {
			i.Logger.TraceLn(cid,
				"RunInitCommands:CheckConnectivity: checking connectivity to safe")

			src, acquired := i.Spiffe.AcquireSourceForSentinel(ctx)
			if !acquired {
				i.Logger.TraceLn(cid,
					"RunInitCommands:CheckConnectivity: failed to acquire source.")

				return errors.New(
					"RunInitCommands:CheckConnectivity: failed to acquire source")
			}

			i.Logger.TraceLn(cid,
				"RunInitCommands:CheckConnectivity"+
					": acquired source successfully")

			if err := i.Safe.Check(ctx, src); err != nil {
				i.Logger.TraceLn(cid,
					"RunInitCommands:CheckConnectivity: "+
						"failed to verify connection to safe:", err.Error())

				return errors.New("runInitCommands:CheckConnectivity:" +
					" cannot establish connection to safe 001")
			}

			i.Logger.TraceLn(cid, "RunInitCommands:CheckConnectivity: success")
			return nil
		})

	if err == nil {
		i.Logger.TraceLn(cid, "exiting backoffs")
		return
	}

	panic("RunInitCommands:CheckConnectivity:" +
		" failed to verify connection to safe")
}

func (i *Initializer) ensureSourceAcquisition(ctx context.Context) *workloadapi.X509Source {
	cid := ctx.Value(key.CorrelationId).(*string)

	i.Logger.TraceLn(cid, "RunInitCommands: acquiring source 001")

	var src *workloadapi.X509Source

	err := backoff.RetryExponential("RunInitCommands:AcquireSource",
		func() error {
			i.Logger.TraceLn(cid, "RunInitCommands:AcquireSource"+
				": acquireSourceForSentinel: 000")

			acq, acquired := i.Spiffe.AcquireSourceForSentinel(ctx)
			src = acq

			if !acquired {
				i.Logger.TraceLn(cid, "RunInitCommands:AcquireSource"+
					": failed to acquire source.")

				return errors.New("RunInitCommands:AcquireSource" +
					": failed to acquire source 000")
			}

			return nil
		})

	if err == nil {
		i.Logger.TraceLn(cid, "RunInitCommands:AcquireSource"+
			": got source. breaking.")
		return src
	}

	panic("RunInitCommands:AcquireSource: failed to acquire source")
}
