/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package wait

import (
	"errors"
	"fmt"
	"time"

	"github.com/vmware-tanzu/secrets-manager/ci/test/io"
)

func ForExampleWorkloadDeletion() error {
	fmt.Println("Waiting for example workload deletion...")

	_, err := io.Exec("kubectl", "wait",
		"--for=delete", "deployment", "-n",
		"default", "--selector=app.kubernetes.io/name=example")

	if err != nil {
		return errors.Join(
			err,
			errors.New("waitForExampleWorkloadDeletion: Failed to wait for deletion"),
		)
	}

	return nil
}

func ForExampleWorkload() error {
	fmt.Println("Waiting for example workload...")

	const maxRetries = 5
	for retries := 0; retries < maxRetries; retries++ {
		_, err := io.Exec("kubectl", "wait", "--for=condition=Ready",
			"pod", "-n", "default", "--selector=app.kubernetes.io/name=example")
		if err == nil {
			return nil // Success
		}

		if retries < maxRetries-1 {
			fmt.Printf("Retry %d/%d: Failed to wait for condition. Retrying in 5 seconds...\n",
				retries+1, maxRetries)
			time.Sleep(5 * time.Second)
			continue
		}
	}

	return errors.New("waitForExampleWorkload: Failed to wait for condition after 5 retries")
}
