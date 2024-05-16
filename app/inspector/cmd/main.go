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

	"github.com/vmware-tanzu/secrets-manager/core/system"
	"github.com/vmware-tanzu/secrets-manager/sdk/sentry"
)

func main() {
	go system.KeepAlive()

	// Fetch the secret from the VSecM Safe.
	d, err := sentry.Fetch()
	if err != nil {
		fmt.Println("Err:", err.Error())
		return
	}

	if d.Data == "" {
		fmt.Println("<nil>")
		return
	}

	// d.Data is a collection of VSecM secrets.
	fmt.Println(d.Data)
}
