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

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortKeys(t *testing.T) {
	m := map[string]string{
		"b": "2",
		"a": "1",
		"c": "3",
	}

	expected := []string{"a", "b", "c"}
	actual := sortKeys(m)

	if !assert.ElementsMatch(t, actual, expected) {
		t.Errorf("sortKeys() = %v; expected %v", actual, expected)
	}
}

func TestEnvVars(t *testing.T) {
	// Temporarily set some environment variables for testing
	_ = os.Setenv("VAR1", "value1")
	_ = os.Setenv("VAR2", "value2")
	defer func() {
		err := os.Unsetenv("VAR1")
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
	defer func() {
		err := os.Unsetenv("VAR2")
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	expected := []string{"VAR1", "VAR2"}
	actual := envVars()

	// Note: actual might contain other environment variables set by the system
	// So we check for expected variables being present in actual
	for _, envVar := range expected {
		if !contains(actual, envVar) {
			t.Errorf("envVars() missing %v; got %v", envVar, actual)
		}
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
