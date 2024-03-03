/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"fmt"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
)

func printGeneratedKeys() {
	privateKey, publicKey, aesSeed, err := crypto.GenerateKeys()

	if err != nil {
		println("Failed to generate keys:")
		println(err.Error())
		return
	}

	println()
	println(crypto.CombineKeys(privateKey, publicKey, aesSeed))
	println()
}
