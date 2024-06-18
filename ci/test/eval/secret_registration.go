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
	"github.com/vmware-tanzu/secrets-manager/ci/test/sentinel"
)

func SecretRegistration() error {
	fmt.Println("----")
	fmt.Println("🧪     Testing: Secret registration...")

	value := "!VSecMRocks!"

	if err := sentinel.SetSecret(value); err != nil {
		return errors.Join(
			err,
			errors.New("setting secret"),
		)
	}

	if err := assert.WorkloadSecretHasValue(value); err != nil {
		return errors.Join(
			err,
			errors.New("asserting workload secret value"),
		)
	}

	fmt.Println("🟢   PASS: Secret registration successful")
	fmt.Println("----")
	return nil
}
