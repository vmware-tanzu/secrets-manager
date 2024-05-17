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
	"flag"
	"fmt"
	"os"
)

func main() {
	origin := flag.String("origin", "local",
		"The origin of the operation, can be 'remote' or 'eks'.")
	ci := flag.String("ci", "", "CI mode")
	flag.Parse()

	// Validate origin
	if *origin != "remote" && *origin != "eks" && *origin != "local" {
		fmt.Println("Invalid origin. Must be 'remote', 'eks', or 'local'.")
		os.Exit(1)
	}

	fmt.Println("---- VSecM Integration Tests ----")
	fmt.Printf("Running tests for %s origin\n", *origin)

	if *ci == "" {
		fmt.Println(`
This script assumes that you have a local minikube cluster running,
and you have already installed SPIRE and VMware Secrets Manager.
Also, make sure you have executed 'eval $(minikube docker-env)'
before running this script.

Press Enter to proceed...`)
		// Wait for user to press enter
		_, _ = fmt.Scanln()
	}

	// Run tests
	run()
}
