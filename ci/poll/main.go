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
	"log"
)

func main() {
	if err := createLockFile(); err != nil {
		log.Println("Another instance is already running. Exiting.")
		return
	}

	defer removeLockFile()

	move, currentCommitHash := proceed()
	if !move {
		log.Println("No commit hash change... exiting.")
		return
	}

	done := runPipeline()
	if !done {
		log.Println("Pipeline failed: exiting.")
		notifyBuildFailure()
		return
	}

	if err := writeCommitHashToFile(currentCommitHash); err != nil {
		log.Println("Error writing to file:", err)
		return
	}
}
