/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package vsecm

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/vmware-tanzu/secrets-manager/ci/test/io"
)

func Sentinel() (string, error) {
	const maxRetries = 5
	var sentinel string

	for retryCount := 0; retryCount < maxRetries; retryCount++ {
		output, err := io.Exec("kubectl", "get", "pods", "-n", "vsecm-system",
			"--selector=app.kubernetes.io/name=vsecm-sentinel",
			"--output=jsonpath={.items[*].metadata.name}")

		if err != nil {
			fmt.Printf("Attempt %d failed: %v\n", retryCount+1, err)
			time.Sleep(10 * time.Second) // Wait before retrying
			continue
		}

		// Split output to handle potential multiple pod names
		sentinelNames := strings.Fields(output)
		if len(sentinelNames) == 1 {
			sentinel = sentinelNames[0]
			break // Successfully defined sentinel
		} else if len(sentinelNames) > 1 {
			// Handle case where multiple sentinel pods are returned
			fmt.Println("Multiple sentinel pods found, selecting the first one.")
			sentinel = sentinelNames[0]
			break
		}

		// If no sentinel pod was found, wait before retrying
		fmt.Println("No sentinel pod found, retrying...")
		time.Sleep(10 * time.Second)
	}

	if sentinel == "" {
		return "", errors.New(
			"defineSentinel: Maximum retries reached without defining a sentinel pod",
		)
	}

	return sentinel, nil
}
