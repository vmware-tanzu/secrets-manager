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

import (
	"os"
	"testing"

	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

func TestRootKeyStoreTypeForSafe(t *testing.T) {
	tests := []struct {
		envValue    string
		expected    data.BackingStore
		shouldPanic bool
	}{
		{"", data.Kubernetes, false},
		{"k8s", data.Kubernetes, false},
		{"invalid", data.Kubernetes, true},
	}

	for _, test := range tests {
		_ = os.Setenv("VSECM_SAFE_ROOT_KEY_STORE", test.envValue)
		if test.shouldPanic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("RootKeyStoreTypeForSafe() with env value %q did not panic as expected", test.envValue)
				}
			}()
		}
		actual := RootKeyStoreTypeForSafe()
		if !test.shouldPanic && actual != test.expected {
			t.Errorf("RootKeyStoreTypeForSafe() with env value %q = %v; expected %v", test.envValue, actual, test.expected)
		}
	}
}

func TestBackingStoreForSafe(t *testing.T) {
	tests := []struct {
		envValue    string
		expected    data.BackingStore
		shouldPanic bool
	}{
		{"", data.File, false},
		{"file", data.File, false},
		{"invalid", data.File, true},
	}

	for _, test := range tests {
		_ = os.Setenv("VSECM_SAFE_BACKING_STORE", test.envValue)
		if test.shouldPanic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("BackingStoreForSafe() with env value %q did not panic as expected", test.envValue)
				}
			}()
		}
		actual := BackingStoreForSafe()
		if !test.shouldPanic && actual != test.expected {
			t.Errorf("BackingStoreForSafe() with env value %q = %v; expected %v", test.envValue, actual, test.expected)
		}
	}
}
