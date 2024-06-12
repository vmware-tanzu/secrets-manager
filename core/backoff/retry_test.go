/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package backoff

import (
	"errors"
	"testing"
)

// mockOperation simulates an operation that fails a specified number of times before succeeding.
func mockOperation(failuresBeforeSuccess int) func() error {
	calls := 0
	return func() error {
		if calls < failuresBeforeSuccess {
			calls++
			return errors.New("operation failed")
		}
		return nil
	}
}

func TestRetry(t *testing.T) {
	tests := []struct {
		name                  string
		failuresBeforeSuccess int
		strategy              Strategy
		expectError           bool
	}{
		{
			name:                  "SuccessWithoutRetry",
			failuresBeforeSuccess: 0,
			strategy:              Strategy{MaxRetries: 5, Delay: 10},
			expectError:           false,
		},
		{
			name:                  "SuccessAfterRetries",
			failuresBeforeSuccess: 3,
			strategy:              Strategy{MaxRetries: 5, Delay: 10},
			expectError:           false,
		},
		{
			name:                  "FailAllRetries",
			failuresBeforeSuccess: 6,
			strategy:              Strategy{MaxRetries: 5, Delay: 10},
			expectError:           true,
		},
		{
			name: "WithDefaultStrategy",
			// Default strategy is MaxRetries: 5, Delay: 1000, Exponential: false, Multiplier: 2
			failuresBeforeSuccess: 3,
			strategy:              Strategy{},
			expectError:           false,
		},
		{
			name:                  "TestSetMaxDurationIfExponentialEnabledAndMaxDurationZero",
			failuresBeforeSuccess: 3,
			strategy:              Strategy{Exponential: true, MaxWait: 0},
			expectError:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := mockOperation(tt.failuresBeforeSuccess)
			err := Retry("test", f, tt.strategy)
			if (err != nil) != tt.expectError {
				t.Errorf("Expected error: %v, got: %v", tt.expectError, err != nil)
			}
		})
	}
}

func TestRetryExponential(t *testing.T) {
	tests := []struct {
		name                  string
		failuresBeforeSuccess int
		expectError           bool
	}{
		{
			name:                  "SuccessAfterRetries",
			failuresBeforeSuccess: 2,
			expectError:           false,
		},
		{
			name:                  "FailAllRetries",
			failuresBeforeSuccess: 12, // Exceeds the default 10 retries
			expectError:           true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//t.Parallel()

			f := mockOperation(tt.failuresBeforeSuccess)
			err := RetryExponential("testExponential: "+tt.name, f)
			if (err != nil) != tt.expectError {
				t.Errorf("%s: expected error: %v, got: %v", tt.name, tt.expectError, err != nil)
			}
		})
	}
}

func TestRetryFixed(t *testing.T) {
	tests := []struct {
		name                  string
		failuresBeforeSuccess int
		expectError           bool
	}{
		{
			name:                  "SuccessAfterRetries",
			failuresBeforeSuccess: 1,
			expectError:           false,
		},
		{
			name:                  "FailAllRetries",
			failuresBeforeSuccess: 12, // Exceeds the default 10 retries
			expectError:           true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			f := mockOperation(tt.failuresBeforeSuccess)
			err := RetryFixed("testFixed: "+tt.name, f)
			if (err != nil) != tt.expectError {
				t.Errorf("%s: expected error: %v, got: %v", tt.name, tt.expectError, err != nil)
			}
		})
	}
}
