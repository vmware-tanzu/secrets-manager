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
	"strings"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/secrets-manager/ci/test/io"
)

func Example() (string, error) {
	cmd := "kubectl"
	args := []string{"get", "pods", "-n", "default", "-o", "name"}

	output, err := io.Exec(cmd, args...)
	if err != nil {
		return "", errors.Wrap(err, "Example: Failed to execute command")
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.Contains(line, "example-") {
			podName := strings.TrimPrefix(line, "pod/")
			return podName, nil
		}
	}

	return "", errors.New("Example: No workload matching 'example-' prefix found")
}
