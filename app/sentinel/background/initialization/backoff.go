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
	"github.com/vmware-tanzu/secrets-manager/core/env"
)

func backoffStrategy() backoff.Strategy {
	return backoff.Strategy{
		MaxRetries:  env.BackoffMaxRetries(),
		Delay:       env.BackoffDelay(),
		Exponential: env.BackoffMode() != "linear",
		MaxDuration: env.BackoffMaxDuration(),
	}
}
