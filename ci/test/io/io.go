/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package io

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// Wait is a placeholder for the workload readiness check.
// Placeholder for the workload readiness check. Replace this with your actual readiness check logic.
func Wait(seconds time.Duration) error {
	// This is a simplification. In a real scenario, you would check the workload's readiness more robustly.
	fmt.Println("Waiting for the workload to be ready...")
	time.Sleep(seconds * time.Second) // Simulate waiting for readiness with a sleep.
	fmt.Println("Workload is now ready.")
	return nil
}

func Exec(command string, args ...string) (string, error) {
	fmt.Println("Executing: `", command, strings.Join(args, " "), "`")

	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}
