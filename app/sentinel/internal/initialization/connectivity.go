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
	"time"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/lib/backoff"
)

func (i *Initializer) ensureApiConnectivity(ctx context.Context, cid *string) {
	i.Logger.TraceLn(cid, "Before checking api connectivity")

	for {
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

				code, body, err := i.Safe.Check(ctx, src)

				i.Logger.TraceLn(cid, "RunInitCommands:CheckConnectivity",
					"code:", code, "body:", body, "err?", err != nil)

				if err != nil {
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

		// Instead of panicking and exiting, we will wait for 5 minutes and then
		// retry. This approach is useful because when VSecM Safe is using an
		// external database, it might not be ready yet. To configure the
		// database, we need VSecM Sentinel to be up and running. So, if we panic
		// here, there is a slight chance that a human operator might be trying
		// to configure VSecM Safe via VSecM Sentinel, and VSecM Sentinel crashes
		// instead of passing the configuration over.

		i.Logger.ErrorLn(cid, "All retries exhausted. Last error:", err.Error())
		i.Logger.InfoLn(cid, "Entering extended retry mode. "+
			"Will attempt again in 5 minutes.")

		select {
		case <-ctx.Done():
			i.Logger.WarnLn(cid, "Context canceled, stopping retry attempts")
			return
		case <-time.After(5 * time.Minute):
			i.Logger.InfoLn(cid, "Resuming connectivity check after extended wait")
		}
	}
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
