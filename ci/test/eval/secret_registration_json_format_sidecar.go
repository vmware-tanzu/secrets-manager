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
	"fmt"
	"github.com/pkg/errors"
	"github.com/vmware-tanzu/secrets-manager/ci/test/assert"
	"github.com/vmware-tanzu/secrets-manager/ci/test/sentinel"
	"time"
)

func SecretRegistrationJSONFormatSidecar() error {
	println("----")
	println("🧪     Testing: Secret registration (JSON transformation)...")

	value := `{"username": "*root*", "password": "*CasHC0w*"}`
	transform := `{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"}`

	// Simulate setting a JSON secret with a transformation.
	if err := sentinel.SetJSONSecret(value, transform); err != nil {
		return errors.Wrap(err, "setJSONSecret failed")
	}

	// Pause to allow time for the secret to be processed by the system.
	time.Sleep(5 * time.Second)

	transformed := `{"USERNAME":"*root*", "PASSWORD":"*CasHC0w*"}`

	// Assert the transformed secret's value.
	if err := assert.WorkloadSecretHasValue(transformed); err != nil {
		return errors.Wrap(err, "assertWorkloadSecretValue failed")
	}

	// Delete the secret as part of cleanup.
	if err := sentinel.DeleteSecret(); err != nil {
		return fmt.Errorf("deleteSecret failed: %w", err)
	}

	println("🟢   PASS: Secret registration (JSON transformation sidecar) successful")
	println("----")
	return nil
}
