/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package cli

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintWorkloadNameNeeded(t *testing.T) {
	output := captureFmtOutput(printWorkloadNameNeeded)

	expectedOutput := `Please provide a workload name.

type ` + "`safe -h`" + ` (without backticks) and press return for help.

`
	assert.Equal(t, expectedOutput, output)

	// Additional checks
	assert.True(t, strings.Contains(output, "Please provide a workload name."))
	assert.True(t, strings.Contains(output, "type `safe -h`"))
	assert.Equal(t, 4, strings.Count(output, "\n"), "Expected 4 newline characters")
}

func TestPrintSecretNeeded(t *testing.T) {
	output := captureFmtOutput(printSecretNeeded)

	expectedOutput := `Please provide a secret.

type ` + "`safe -h`" + ` (without backticks) and press return for help.

`
	assert.Equal(t, expectedOutput, output)

	// Additional checks
	assert.True(t, strings.Contains(output, "Please provide a secret."))
	assert.True(t, strings.Contains(output, "type `safe -h`"))
	assert.Equal(t, 4, strings.Count(output, "\n"), "Expected 4 newline characters")
}
