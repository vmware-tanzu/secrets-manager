/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package workload

import (
	"errors"
	"strings"

	"github.com/vmware-tanzu/secrets-manager/ci/test/io"
)

func Example() (string, error) {
	cmd := "kubectl"
	args := []string{"get", "pods", "-n", "default", "-o", "name"}

	output, err := io.Exec(cmd, args...)
	if err != nil {
		return "", errors.Join(
			err,
			errors.New("example: Failed to execute command"),
		)
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "example-") {
			podName := strings.TrimPrefix(line, "pod/")
			return podName, nil
		}
	}

	return "", errors.New("example: No workload matching 'example-' prefix found")
}
