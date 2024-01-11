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
	"fmt"
)

func main() {
	if err := createLockFile(); err != nil {
		fmt.Println("Another instance is already running. Exiting.")
		return
	}

	defer removeLockFile()

	move, currentCommitHash := proceed()
	if !move {
		fmt.Println("No commit hash change… exiting.")
		return
	}

	done := runPipeline()
	if !done {
		fmt.Println("Pipeline failed: exiting.")
		notifyBuildFailure()
		return
	}

	if err := writeCommitHashToFile(currentCommitHash); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
