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
	"time"

	"github.com/vmware-tanzu/secrets-manager/ci/test/assert"
	"github.com/vmware-tanzu/secrets-manager/ci/test/sentinel"
)

func SecretRegistrationYAMLFormatSidecar() error {
	fmt.Println("----")
	fmt.Println("🧪     Testing Secret registration (YAML transformation)...")

	value := `{"username": "*root*", "password": "*CasHC0w*"}`
	transform := `{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}`

	// Simulate setting a YAML secret with transformation.
	if err := setYAMLSecret(value, transform); err != nil {
		return errors.Join(
			err,
			errors.New("setYAMLSecret failed"),
		)
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
		return errors.Join(
			err,
			errors.New("assertWorkloadSecretValue failed"),
		)
	}

	// Delete the secret as part of cleanup.
	if err := sentinel.DeleteSecret(); err != nil {
		return errors.Join(
			err,
			errors.New("deleteSecret failed"),
		)
	}

	fmt.Println("🟢   PASS: Secret registration (YAML transformation sidecar) successful")
	fmt.Println("----")
	return nil
}
