/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"github.com/vmware-tanzu/secrets-manager/lib/system"
)

func main() {
	// Run on the main thread to wait forever.
	system.KeepAlive()
}
