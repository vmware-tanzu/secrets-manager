/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package rpc

import (
	"os"
	"testing"
)

func TestSentinelLoggerUrl(t *testing.T) {
	tests := []struct {
		envValue string
		expected string
	}{
		{"", "localhost:50051"},
		{"testLoggerUrl", "testLoggerUrl"},
	}

	for _, test := range tests {
		_ = os.Setenv("VSECM_SENTINEL_LOGGER_URL", test.envValue)
		actual := SentinelLoggerUrl()
		if actual != test.expected {
			t.Errorf("SentinelLoggerUrl() with env value %q = %v; expected %v", test.envValue, actual, test.expected)
		}
	}
}
