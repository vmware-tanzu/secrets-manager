/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package state

import (
	"errors"
	"fmt"

	"github.com/vmware-tanzu/secrets-manager/ci/test/io"
	"github.com/vmware-tanzu/secrets-manager/ci/test/vsecm"
	"github.com/vmware-tanzu/secrets-manager/ci/test/wait"
)

func Cleanup() error {
	fmt.Println("----")
	fmt.Println("🧹 Cleanup...")

	// Determine the sentinel pod.
	sentinel, err := vsecm.Sentinel()
	if err != nil {
		return errors.Join(
			err,
			errors.New("cleanup: Failed to determine sentinel"),
		)
	}

	// Execute command within the sentinel pod to delete the secret.
	_, err = io.Exec("kubectl", "exec",
		sentinel, "-n", "vsecm-system", "--", "safe", "-w", "example", "-d")
	if err != nil {
		return errors.Join(
			err,
			errors.New("cleanup: Failed to delete secret"),
		)
	}

	// Check if the deployment exists before attempting to delete.
	_, err = io.Exec("kubectl", "get", "deployment", "example", "-n", "default")
	if err == nil {
		_, err = io.Exec(
			"kubectl", "delete", "deployment", "example", "-n", "default",
		)
		if err != nil {
			return errors.Join(
				err,
				errors.New("cleanup: Failed to delete deployment"),
			)
		}
	}

	// Wait for the workload to be gone.
	if err := wait.ForExampleWorkloadDeletion(); err != nil {
		return errors.Join(
			err,
			errors.New("cleanup: Failed to wait for workload deletion"),
		)
	}

	fmt.Println("✨ All clean and shiny!")
	fmt.Println("----")
	return nil
}
