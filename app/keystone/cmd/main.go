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
	"fmt"
	"log"
	"os"

	"github.com/vmware-tanzu/secrets-manager/core/system"
)

func main() {
	log.Println(
		"VSecM Keystone",
		fmt.Sprintf("v%s", os.Getenv("APP_VERSION")),
	)

	// Run on the main thread to wait forever.
	system.KeepAlive()
}
