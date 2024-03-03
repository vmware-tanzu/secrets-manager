/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package eval

import (
	"github.com/pkg/errors"
	"github.com/vmware-tanzu/secrets-manager/ci/test/assert"
	"github.com/vmware-tanzu/secrets-manager/ci/test/sentinel"
	"time"
)

func SecretRegistrationSidecar() error {
	println("Testing: Secret registration (sidecar)…")

	value := "!VSecMRocks!"

	if err := sentinel.SetSecret(value); err != nil {
		return errors.Wrap(err, "setSecret failed")
	}

	// Pause to simulate waiting for the secret to propagate or system to update.
	time.Sleep(5 * time.Second)

	if err := assert.WorkloadSecretHasValue(value); err != nil {
		return errors.Wrap(err, "assertWorkloadSecretValue failed")
	}

	println("Secret registration (sidecar) successful")
	return nil
}
