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
	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"golang.org/x/net/context"
	"time"
)

func processCommandBlock(ctx context.Context, sc entity.SentinelCommand) error {
	return safe.Post(ctx, sc)
}

func doSleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Millisecond)
}
