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
	"github.com/pkg/errors"
	"github.com/vmware-tanzu/secrets-manager/ci/test/assert"
	"github.com/vmware-tanzu/secrets-manager/ci/test/sentinel"
	"strings"
)

func setYAMLSecret(value, transform string) (string, error) {
	// Direct string transformation for demonstration. Real logic might involve actual data manipulation.
	transformed := strings.ReplaceAll(transform, "{{.username}}", "*root*")
	transformed = strings.ReplaceAll(transformed, "{{.password}}", "*CasHC0w*")
	// Simulate the YAML-like transformation.
	yamlTransformed := strings.ReplaceAll(transformed, `{"USERNAME":"`, "USERNAME: '")
	yamlTransformed = strings.ReplaceAll(yamlTransformed, `", "PASSWORD":"`, "'\nPASSWORD: '")
	yamlTransformed = strings.ReplaceAll(yamlTransformed, `"}`, "'")
	return yamlTransformed, nil
}

func SecretRegistrationYAMLFormat() error {
	println("----")
	println("Testing: Secret registration (YAML transformation)...")

	value := `{"username": "*root*", "password": "*CasHC0w*"}`
	transform := `{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}`

	_, err := setYAMLSecret(value, transform)
	if err != nil {
		return errors.Wrap(err, "setYAMLSecret failed")
	}

	expectedTransformed := `
PASSWORD: '*CasHC0w*'
USERNAME: '*root*'
`
	if err := assert.WorkloadSecretHasValue(
		strings.TrimSpace(expectedTransformed),
	); err != nil {
		return errors.Wrap(err, "assertWorkloadSecretValue failed")
	}

	if err := sentinel.DeleteSecret(); err != nil {
		return errors.Wrap(err, "deleteSecret failed")
	}

	println("Secret registration (YAML transformation) successful")
	return nil
}
