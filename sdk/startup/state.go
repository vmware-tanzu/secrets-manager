/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package startup

import "github.com/vmware-tanzu/secrets-manager/sdk/sentry"

func initialized() bool {
	r, _ := sentry.Fetch()
	v := r.Data
	return v != ""
}
