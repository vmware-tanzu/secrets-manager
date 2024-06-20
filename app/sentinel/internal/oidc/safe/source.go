/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package safe

import (
	"context"
	"log"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/core/env"
)

func acquireSource(ctx context.Context) (*workloadapi.X509Source, bool) {
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

		if err != nil {
			log.Println(cid, "acquireSource: "+
				"I am having trouble fetching my identity from SPIRE.",
				err.Error())
			log.Println(cid,
				"acquireSource: "+
					"I won't proceed until you put me in a secured container.",
				err.Error())
			errorChan <- err
			return
		}
		resultChan <- source
	}()

	select {
	case source := <-resultChan:
		return source, true
	case err := <-errorChan:
		log.Println(cid, "acquireSource: "+
			"I cannot execute command because I cannot talk to SPIRE.",
			err.Error())
		return nil, false
	case <-ctx.Done():
		log.Println(cid, "acquireSource: Operation was cancelled.")
		return nil, false
	}
}
