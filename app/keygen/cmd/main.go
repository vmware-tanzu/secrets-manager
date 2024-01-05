/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"github.com/vmware-tanzu/secrets-manager/core/env"
)

func main() {
	d := env.KeyGenDecrypt()

	if d {
		printDecryptedKeys()
		return
	}

	printGeneratedKeys()
}
