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
	"os"
)

const lockFilePath = "/opt/vsecm/git_poller.lock"

// createLockFile tries to create a lock file and returns an error if it already exists
func createLockFile() error {
	lockFile, err := os.OpenFile(lockFilePath, os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return err
	}
	err = lockFile.Close()
	if err != nil {
		return err
	}
	return nil
}

// removeLockFile deletes the lock file
func removeLockFile() {
	err := os.Remove(lockFilePath)
	if err != nil {
		fmt.Printf("Error removing lock file: %s", err)
	}
}
