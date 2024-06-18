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

	"github.com/vmware-tanzu/secrets-manager/ci/test/assert"
	"github.com/vmware-tanzu/secrets-manager/ci/test/deploy"
	"github.com/vmware-tanzu/secrets-manager/ci/test/sentinel"
)

func SecretEncryption() error {
	fmt.Println("----")
	fmt.Println("🧪     Testing: Encrypting secrets...")

	value := "!VSecMRocks!"

	if err := assert.SentinelCanEncryptSecret(value); err != nil {
		return errors.Join(
			err,
			errors.New("asserting encrypted secret"),
		)
	}

	if err := deploy.WorkloadUsingSDK(); err != nil {
		return errors.Join(
			err,
			errors.New("deploying workload using SDK"),
		)
	}

	if err := sentinel.SetEncryptedSecret(value); err != nil {
		return errors.Join(
			err,
			errors.New("setting encrypted secret"),
		)
	}

	if err := assert.WorkloadSecretHasValue(value); err != nil {
		return errors.Join(
			err,
			errors.New("asserting workload secret value"),
		)
	}

	fmt.Println("🟢   PASS: Secret encryption successful")
	fmt.Println("----")
	return nil
}
