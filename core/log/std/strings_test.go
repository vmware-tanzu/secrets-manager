/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package std

import "testing"

func TestToCustomCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello_world", "Hello world"},
		{"test_case", "Test case"},
		{"ANOTHER_test_case", "Another test case"},
		{"already Properly Formatted", "Already Properly Formatted"},
	}

	for _, test := range tests {
		actual := toCustomCase(test.input)
		if actual != test.expected {
			t.Errorf("toCustomCase(%q) = %v; expected %v", test.input, actual, test.expected)
		}
	}
}
