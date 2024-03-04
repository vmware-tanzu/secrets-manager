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
)

func SecretDeletion() error {
	println("----")
	println("🧪 Testing: Secret deletion...")

	if err := sentinel.DeleteSecret(); err != nil {
		return errors.Wrap(err, "deleting secret")
	}

	if err := assert.WorkloadSecretHasNoValue(); err != nil {
		return errors.Wrap(err, "asserting workload secret no value")
	}

	println("🟢   PASS: Secret deletion successful")
	println("----")
	return nil
}
