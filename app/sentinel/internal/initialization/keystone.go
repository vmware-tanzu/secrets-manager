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

	"github.com/vmware-tanzu/secrets-manager/core/constants/keystone"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/lib/backoff"
)

func (i *Initializer) markKeystone(ctx context.Context, cid *string) bool {
	err := backoff.RetryExponential(
		"RunInitCommands:MarkKeystone",
		func() error {
			i.Logger.TraceLn(cid, "RunInitCommands:MarkKeystone"+
				": retrying with exponential backoff")

			// Assign a secret for VSecM Keystone
			err := i.Safe.Post(ctx, entity.SentinelCommand{
				WorkloadIds: []string{keystone.WorkloadId},
				Namespaces:  []string{i.EnvReader.NamespaceForVSecMSystem()},
				Secret:      keystone.InitSecretValue,
			})

			return err
		})

	if err == nil {
		return true
	}

	i.Logger.ErrorLn(
		cid, "RunInitCommands: error setting keystone secret: ",
		err.Error())
	panic("RunInitCommands: error setting keystone secret")
}
