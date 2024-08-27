/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import "testing"

func TestValue(t *testing.T) {
	t.Run("success_case", func(t *testing.T) {
		t.Setenv("TEST_ENV", "test")
		got := Value("TEST_ENV")
		if got != "test" {
			t.Errorf("Value() = %v, want %v", got, "test")
		}
	})

	t.Run("empty_case", func(t *testing.T) {
		t.Setenv("TEST_ENV", "")
		got := Value("TEST_ENV")
		if got != "" {
			t.Errorf("Value() = %v, want %v", got, "")
		}
	})

	t.Run("not_set_case", func(t *testing.T) {
		got := Value("TEST_ENV")
		if got != "" {
			t.Errorf("Value() = %v, want %v", got, "")
		}
	})
}
