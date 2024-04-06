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
	"github.com/pkg/errors"
	"time"

	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
	"github.com/vmware-tanzu/secrets-manager/core/backoff"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/spiffe"
)

func ensureApiConnectivity(ctx context.Context, cid *string) {
	terminateAsap := env.TerminateSentinelOnInitCommandConnectivityFailure()

	log.TraceLn(cid, "Before checking api connectivity")

	for {
		s := backoffStrategy()

		err := backoff.Retry("RunInitCommands:CheckConnectivity", func() error {
			log.TraceLn(cid, "RunInitCommands:CheckConnectivity: checking connectivity to safe")

			src, acquired := spiffe.AcquireSourceForSentinel(ctx)
			if !acquired {
				log.TraceLn(cid, "RunInitCommands:CheckConnectivity: failed to acquire source.")
				if terminateAsap {
					panic("RunInitCommands:CheckConnectivity: failed to acquire source")
				}

				return errors.New("RunInitCommands:CheckConnectivity: failed to acquire source")
			}

			log.TraceLn(cid, "RunInitCommands:CheckConnectivity: acquired source successfully")

			if err := safe.Check(ctx, src); err != nil {
				log.TraceLn(cid, "RunInitCommands:CheckConnectivity: failed to verify connection to safe:", err.Error())
				if terminateAsap {
					panic("RunInitCommands:CheckConnectivity: failed to verify connection to safe")
				}

				return errors.Wrap(err, "RunInitCommands:CheckConnectivity: cannot establish connection to safe 001")
			}

			log.TraceLn(cid, "RunInitCommands:CheckConnectivity: success")
			return nil
		}, s)

		if err == nil {
			log.TraceLn(cid, "exiting backoffs")
			break
		}
	}
}

func ensureSourceAcquisition(ctx context.Context, cid *string) {
	// If `true`, instead of retrying with a backoff, kill the pod, and let the
	// deployment controller restart it to initiate a new retry.
	terminateAsap := env.TerminateSentinelOnInitCommandConnectivityFailure()

	waitInterval := env.InitCommandRunnerWaitIntervalForSentinel()
	time.Sleep(waitInterval)

	for {
		log.TraceLn(cid, "RunInitCommands: acquiring source 001")

		s := backoff.Strategy{
			MaxRetries:  20,
			Delay:       1000,
			Exponential: true,
			MaxDuration: 30 * time.Second,
		}

		err := backoff.Retry("RunInitCommands:AcquireSource", func() error {
			log.TraceLn(cid, "RunInitCommands:AcquireSource: acquireSourceForSentinel: 000")
			_, acquired := spiffe.AcquireSourceForSentinel(ctx)
			if !acquired {
				log.TraceLn(cid, "RunInitCommands:AcquireSource: failed to acquire source.")
				if terminateAsap {
					panic("RunInitCommands:AcquireSource: failed to acquire source")
				}

				return errors.New("RunInitCommands:AcquireSource: failed to acquire source 000")
			}

			return nil
		}, s)

		if err == nil {
			log.TraceLn(cid, "RunInitCommands:AcquireSource: got source. breaking.")
			break
		}
	}
}
