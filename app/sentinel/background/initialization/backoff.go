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
	"github.com/vmware-tanzu/secrets-manager/core/backoff"
	"time"
)

// TODO: get some of these from env vars.
func backoffStrategy() backoff.Strategy {
	return backoff.Strategy{
		MaxRetries:  20,
		Delay:       1000,
		Exponential: true,
		MaxDuration: 30 * time.Second,
	}
}
