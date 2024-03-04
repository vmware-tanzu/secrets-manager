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
	"time"
)

func SecretRegistrationYAMLFormatSidecar() error {
	println("----")
	println("🧪     Testing Secret registration (YAML transformation)...")

	value := `{"username": "*root*", "password": "*CasHC0w*"}`
	transform := `{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}`

	// Simulate setting a YAML secret with transformation.
	if err := setYAMLSecret(value, transform); err != nil {
		return errors.Wrap(err, "setYAMLSecret failed")
	}

	// Simulated transformed YAML as a string
	transformed := `
PASSWORD: '*CasHC0w*'
USERNAME: '*root*'
`

	// Pause to allow time for the system to process the secret.
	time.Sleep(5 * time.Second)

	// Assert the transformed secret's value.
	if err := assert.WorkloadSecretHasValue(
		strings.TrimSpace(transformed),
	); err != nil {
		return errors.Wrap(err, "assertWorkloadSecretValue failed")
	}

	// Delete the secret as part of cleanup.
	if err := sentinel.DeleteSecret(); err != nil {
		return errors.Wrap(err, "deleteSecret failed")
	}

	println("🟢   PASS: Secret registration (YAML transformation sidecar) successful")
	println("----")
	return nil
}
