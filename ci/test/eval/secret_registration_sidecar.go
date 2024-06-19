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

func SecretRegistrationSidecar() error {
	fmt.Println("----")
	fmt.Println("🧪     Testing: Secret registration (sidecar)...")

	value := "!VSecMRocks!"

	if err := sentinel.SetSecret(value); err != nil {
		return errors.Join(
			err,
			errors.New("setSecret failed"),
		)
	}

	// Pause to simulate waiting for the secret to propagate or system to update.
	time.Sleep(5 * time.Second)

	if err := assert.WorkloadSecretHasValue(value); err != nil {
		return errors.Join(
			err,
			errors.New("assertWorkloadSecretValue failed"),
		)
	}

	fmt.Println("🟢   PASS: Secret registration (sidecar) successful")
	fmt.Println("----")
	return nil
}
