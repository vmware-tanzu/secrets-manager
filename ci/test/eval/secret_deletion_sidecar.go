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
	"time"
)

func SecretDeletionSidecar() error {
	println("----")
	println("🧪 Testing: Secret deletion (sidecar)...")

	if err := sentinel.DeleteSecret(); err != nil {
		return errors.Wrap(err, "deleteSecret failed")
	}

	// Pause to simulate waiting for the system to process the secret deletion.
	time.Sleep(5 * time.Second) // Adjust the duration as needed for your environment.

	if err := assert.WorkloadSecretHasNoValue(); err != nil {
		return errors.Wrap(err, "assertWorkloadSecretNoValue failed")
	}

	println("🟢   PASS: Secret deletion (sidecar) successful")
	println("----")
	return nil
}
