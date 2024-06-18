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
	"time"

	"github.com/vmware-tanzu/secrets-manager/ci/test/assert"
	"github.com/vmware-tanzu/secrets-manager/ci/test/sentinel"
)

func SecretRegistrationAppendSidecar() error {
	fmt.Println("----")
	fmt.Println("🧪     Testing Secret registration (append mode)...")

	secret1 := "!VSecM"
	secret2 := "Rocks!"
	value := fmt.Sprintf(`["%s","%s"]`, secret2, secret1)

	if err := sentinel.AppendSecret(secret1); err != nil {
		return errors.Join(
			err,
			errors.New("appendSecret failed for secret1"),
		)
	}

	if err := sentinel.AppendSecret(secret2); err != nil {
		return errors.Join(
			err,
			errors.New("appendSecret failed for secret2"),
		)
	}

	// Pause to allow time for the system to process the appended secrets.
	time.Sleep(5 * time.Second)

	// Assert the combined value of the appended secrets.
	if err := assert.WorkloadSecretHasValue(value); err != nil {
		return errors.Join(
			err,
			errors.New("assertWorkloadSecretValue failed"),
		)
	}

	// Delete the secrets as part of cleanup.
	if err := sentinel.DeleteSecret(); err != nil {
		return errors.Join(
			err,
			errors.New("deleteSecret failed"),
		)
	}

	fmt.Println("🟢   PASS: Secret registration (append sidecar mode) successful")
	fmt.Println("----")
	return nil
}
