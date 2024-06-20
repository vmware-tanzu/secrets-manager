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

	"github.com/vmware-tanzu/secrets-manager/core/constants/sentinel"
)

// PrintUsage prints the usage of the CLI.
func PrintUsage(parser *argparse.Parser) {
	fmt.Print(parser.Usage(sentinel.CmdName))
}

// InputValidationFailure checks if the input provided by the user is valid.
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
