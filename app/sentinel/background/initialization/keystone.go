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

	"github.com/vmware-tanzu/secrets-manager/core/backoff"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func markKeystone(ctx context.Context, cid *string) bool {
	terminateAsap := env.TerminateSentinelOnInitCommandConnectivityFailure()

	s := backoffStrategy()
	err := backoff.Retry("RunInitCommands:MarkKeystone", func() error {
		log.TraceLn(cid, "RunInitCommands:MarkKeystone: retrying with exponential backoff")

		// Assign a secret for VSecM Keystone
		err := processCommandBlock(ctx, entity.SentinelCommand{
			WorkloadIds: []string{"vsecm-keystone"},
			Namespaces:  []string{"vsecm-system"},
			Secret:      "keystone-init",
		})

		if err != nil {
			if terminateAsap {
				panic("RunInitCommands: error setting keystone secret")
			}
		}

		return err
	}, s)

	if err != nil {
		log.ErrorLn(cid, "RunInitCommands: error setting keystone secret: ", err.Error())
		if terminateAsap {
			panic("RunInitCommands: error setting keystone secret")
		}
		return false
	}

	return true
}
