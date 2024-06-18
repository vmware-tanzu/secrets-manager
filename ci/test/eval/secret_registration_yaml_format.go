/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package eval

import (
	"errors"
	"fmt"
	"strings"

	"github.com/vmware-tanzu/secrets-manager/ci/test/assert"
	"github.com/vmware-tanzu/secrets-manager/ci/test/io"
	"github.com/vmware-tanzu/secrets-manager/ci/test/sentinel"
	"github.com/vmware-tanzu/secrets-manager/ci/test/vsecm"
)

func setYAMLSecret(value, transform string) error {
	if value == "" || transform == "" {
		return errors.New("setYAMLSecret: Value or transform is empty")
	}

	s, err := vsecm.Sentinel()
	if err != nil || s == "" {
		return errors.Join(
			err,
			errors.New("setYAMLSecret: Failed to get sentinel"),
		)
	}

	// Executing command within the sentinel pod to set the YAML secret with
	// transformation.
	_, err = io.Exec("kubectl", "exec", s, "-n", "vsecm-system",
		"--", "safe", "-w", "example", "-n", "default", "-s", value,
		"-t", transform, "-f", "yaml")
	if err != nil {
		return errors.Join(
			err,
			errors.New("setYAMLSecret: Failed to exec kubectl"),
		)
	}

	return nil
}

func SecretRegistrationYAMLFormat() error {
	fmt.Println("----")
	fmt.Println("🧪     Testing: Secret registration (YAML transformation)...")

	value := `{"username": "*root*", "password": "*CasHC0w*"}`
	transform := `{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}`

	err := setYAMLSecret(value, transform)
	if err != nil {
		return errors.Join(
			err,
			errors.New("setYAMLSecret failed"),
		)
	}

	expectedTransformed := `
PASSWORD: '*CasHC0w*'
USERNAME: '*root*'
`
	if err := assert.WorkloadSecretHasValue(
		strings.TrimSpace(expectedTransformed),
	); err != nil {
		return errors.Join(
			err,
			errors.New("assertWorkloadSecretValue failed"),
		)
	}

	if err := sentinel.DeleteSecret(); err != nil {
		return errors.Join(
			err,
			errors.New("deleteSecret failed"),
		)
	}

	fmt.Println("🟢   PASS: Secret registration (YAML transformation) successful")
	fmt.Println("----")
	return nil
}
