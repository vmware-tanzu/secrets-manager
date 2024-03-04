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
	"os"
	"os/exec"
)

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout // Forward stdout
	cmd.Stderr = os.Stderr // Forward stderr

	log.Printf("Executing command: %s %v", command, args)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
