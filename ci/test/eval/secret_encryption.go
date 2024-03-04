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
	"github.com/vmware-tanzu/secrets-manager/ci/test/deploy"
	"github.com/vmware-tanzu/secrets-manager/ci/test/sentinel"
)

func SecretEncryption() error {
	println("----")
	println("🧪 Testing: Encrypting secrets...")

	value := "!VSecMRocks!"

	if err := assert.SentinelCanEncryptSecret(value); err != nil {
		return errors.Wrap(err, "asserting encrypted secret")
	}

	if err := deploy.WorkloadUsingSDK(); err != nil {
		return errors.Wrap(err, "deploying workload using SDK")
	}

	if err := sentinel.SetEncryptedSecret(value); err != nil {
		return errors.Wrap(err, "setting encrypted secret")
	}

	if err := assert.WorkloadSecretHasValue(value); err != nil {
		return errors.Wrap(err, "asserting workload secret value")
	}

	println("🟢 PASS: Secret encryption successful")
	println("----")
	return nil
}
