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

import "os"

var commitHashFile = "/opt/vsecm/commit-hash"

func readCommitHashFromFile() (string, error) {
	data, err := os.ReadFile(commitHashFile)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func writeCommitHashToFile(commitHash string) error {
	return os.WriteFile(commitHashFile, []byte(commitHash), 0644)
}
