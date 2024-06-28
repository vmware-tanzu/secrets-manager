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
	"io"
	"log"
	"os"
	"time"
)

var lockFilePath = "/opt/vsecm/git_poller.lock"

// createLockFile tries to create a lock file and returns an error if it
// already exists
func createLockFile() error {
	lockFile, err := os.OpenFile(lockFilePath, os.O_CREATE|os.O_EXCL, 0666)
	if err == nil {
		return nil
	}
	defer func(l io.ReadCloser) {
		err := l.Close()
		if err != nil {
			log.Printf("Error closing lock file: %s", err)
		}
	}(lockFile)

	if !os.IsExist(err) {
		return err
	}

	// Check if the lock file is more than one day old
	fileInfo, statErr := os.Stat(lockFilePath)
	if statErr != nil {
		return statErr
	}

	if time.Since(fileInfo.ModTime()) > 24*time.Hour {

		// File is older than one day, attempt to remove it and create a new one
		removeErr := os.Remove(lockFilePath)
		if removeErr != nil {
			return removeErr
		}

		return createLockFile()
	}

	// File is new; return the error instead.
	return err
}

// removeLockFile deletes the lock file
func removeLockFile() {
	err := os.Remove(lockFilePath)
	if err != nil {
		log.Printf("Error removing lock file: %s", err)
	}
}
