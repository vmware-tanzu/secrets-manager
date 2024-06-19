/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package cli

import "fmt"

func printWorkloadNameNeeded() {
	fmt.Println("Please provide a workload name.")
	fmt.Println("")
	fmt.Println("type `safe -h` (without backticks) and press return for help.")
	fmt.Println("")
}

func printSecretNeeded() {
	fmt.Println("Please provide a secret.")
	fmt.Println("")
	fmt.Println("type `safe -h` (without backticks) and press return for help.")
	fmt.Println("")
}
