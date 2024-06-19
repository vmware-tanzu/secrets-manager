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
	"github.com/vmware-tanzu/secrets-manager/sdk/sentry"
	"strings"
)

func main() {
	d, err := sentry.Fetch()
	if err != nil {
		msg := err.Error()

		if strings.Contains(strings.ToLower(msg),
			"secret does not exist",
		) {
			fmt.Print("NO_SECRET")
			return
		}

		fmt.Print("ERR_SENTRY_FETCH_FAILED")
		fmt.Print(" ", err.Error())
		return
	}

	if strings.TrimSpace(d.Data) == "" {
		fmt.Print("NO_SECRET")
	}

	fmt.Print(d.Data)
}
