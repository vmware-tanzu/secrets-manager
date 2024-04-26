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
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"time"

	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
	"github.com/vmware-tanzu/secrets-manager/core/backoff"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/spiffe"
)

func ensureApiConnectivity(ctx context.Context, cid *string) {
	log.TraceLn(cid, "Before checking api connectivity")

	s := backoffStrategy()
	err := backoff.Retry("RunInitCommands:CheckConnectivity", func() error {
		log.TraceLn(cid, "RunInitCommands:CheckConnectivity: checking connectivity to safe")

		src, acquired := spiffe.AcquireSourceForSentinel(ctx)
		if !acquired {
			log.TraceLn(cid, "RunInitCommands:CheckConnectivity: failed to acquire source.")

			return errors.New("RunInitCommands:CheckConnectivity: failed to acquire source")
		}

		log.TraceLn(cid, "RunInitCommands:CheckConnectivity: acquired source successfully")

		if err := safe.Check(ctx, src); err != nil {
			log.TraceLn(cid, "RunInitCommands:CheckConnectivity: failed to verify connection to safe:", err.Error())

			return errors.Wrap(err, "RunInitCommands:CheckConnectivity: cannot establish connection to safe 001")
		}

		log.TraceLn(cid, "RunInitCommands:CheckConnectivity: success")
		return nil
	}, s)

	if err == nil {
		log.TraceLn(cid, "exiting backoffs")
		return
	}

	// I shouldn't be here.
	panic("RunInitCommands:CheckConnectivity: failed to verify connection to safe")
}

func ensureSourceAcquisition(ctx context.Context, cid *string) *workloadapi.X509Source {
	// If `true`, instead of retrying with a backoff, kill the pod, and let the
	// deployment controller restart it to initiate a new retry.

	log.TraceLn(cid, "RunInitCommands: acquiring source 001")

	s := backoff.Strategy{
		MaxRetries:  20,
		Delay:       1000,
		Exponential: true,
		MaxDuration: 30 * time.Second,
	}

	var src *workloadapi.X509Source

	err := backoff.Retry("RunInitCommands:AcquireSource", func() error {
		log.TraceLn(cid, "RunInitCommands:AcquireSource: acquireSourceForSentinel: 000")

		acq, acquired := spiffe.AcquireSourceForSentinel(ctx)
		src = acq

		if !acquired {
			log.TraceLn(cid, "RunInitCommands:AcquireSource: failed to acquire source.")

			return errors.New("RunInitCommands:AcquireSource: failed to acquire source 000")
		}

		return nil
	}, s)

	if err == nil {
		log.TraceLn(cid, "RunInitCommands:AcquireSource: got source. breaking.")
		return src
	}

	// I shouldn't be here.
	panic("RunInitCommands:AcquireSource: failed to acquire source")
}
