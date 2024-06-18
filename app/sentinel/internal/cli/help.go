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

import (
	"fmt"

	"github.com/akamensky/argparse"
)

func PrintUsage(parser *argparse.Parser) {
	fmt.Print(parser.Usage("safe"))
}

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

func InputValidationFailure(
	workload *[]string, encrypt *bool, inputKeys *string,
	secret *string, deleteSecret *bool) bool {

	// You need to provide a workload name if you are not encrypting a secret,
	// or if you are not providing input keys.
	if len(*workload) == 0 &&
		!*encrypt &&
		*inputKeys == "" {
		printWorkloadNameNeeded()
		return true
	}

	// You need to provide a secret value if you are not deleting a secret,
	// or if you are not providing input keys.
	if *secret == "" &&
		!*deleteSecret &&
		*inputKeys == "" {
		printSecretNeeded()
		return true
	}

	return false
}
