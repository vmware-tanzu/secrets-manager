/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package spiffe

import (
	"context"
	"errors"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware-tanzu/secrets-manager/core/constants/key"

	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
)

// AcquireSourceForSentinel initiates an asynchronous operation to obtain an
// X509Source from the SPIFFE workload API, using the context for cancellation
// and a correlation ID for logging purposes.
//
// It attempts to create a new X509Source configured with the SPIRE server
// address from the environment, fetches the X509SVID from the source, and
// validates the SVID against a known VSecM Sentinel value to ensure the caller
// is operating within a trusted environment.
//
// Parameters:
//   - ctx: A context.Context object used for cancellation and to carry metadata
//     across API boundaries, including a correlation ID for tracking the
//     operation in logs.
//
// Returns:
//   - A pointer to a workloadapi.X509Source object if the source is
//     successfully acquired and validated. This object can be used to obtain
//     X.509 SVIDs for secure communication.
//   - A boolean flag indicating whether the source was successfully acquired
//     (true) or not (false). If false, the source pointer will be nil.
func AcquireSourceForSentinel(
	ctx context.Context,
) (*workloadapi.X509Source, bool) {
	resultChan := make(chan *workloadapi.X509Source)
	errorChan := make(chan error)

	cid := ctx.Value(key.CorrelationId).(*string)

	go func() {
		source, err := workloadapi.NewX509Source(
			ctx, workloadapi.WithClientOptions(
				workloadapi.WithAddr(env.SpiffeSocketUrl()),
			),
		)

		if err != nil {
			errorChan <- err
			return
		}

		svid, err := source.GetX509SVID()
		if err != nil {
			log.ErrorLn(cid,
				"acquireSource: trouble fetching my identity from SPIRE.")
			log.ErrorLn(cid,
				"acquireSource: not in a secured container.")
			errorChan <- err
			return
		}

		// Make sure that the binary is enclosed in a Pod that we trust.
		if !validation.IsSentinel(svid.ID.String()) {
			log.ErrorLn(cid,
				"acquireSource: I don't know you, and it's crazy: '"+
					svid.ID.String()+"'")
			log.ErrorLn(cid,
				"acquireSource: "+
					"`safe` can only run from within the Sentinel container.")
			errorChan <- errors.New(
				"acquireSource: I don't know you, and it's crazy: '" +
					svid.ID.String() + "'")
			return
		}

		resultChan <- source
	}()

	select {
	case source := <-resultChan:
		log.InfoLn(cid, "acquireSource: Source acquired.")
		return source, true
	case err := <-errorChan:
		log.ErrorLn(cid, "acquireSource: "+
			"I cannot execute command because I cannot talk to SPIRE.",
			err.Error())
		return nil, false
	case <-ctx.Done():
		log.ErrorLn(cid, "acquireSource: Operation was cancelled.")
		return nil, false
	}
}
