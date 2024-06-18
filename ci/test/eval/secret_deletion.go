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

func SecretDeletion() error {
	fmt.Println("----")
	fmt.Println("🧪     Testing: Secret deletion...")

	if err := sentinel.DeleteSecret(); err != nil {
		return errors.Join(
			err,
			errors.New("deleting secret"),
		)
	}

	if err := assert.WorkloadSecretHasNoValue(); err != nil {
		return errors.Join(
			err,
			errors.New("asserting workload secret no value"),
		)
	}

	fmt.Println("🟢   PASS: Secret deletion successful")
	fmt.Println("----")
	return nil
}
