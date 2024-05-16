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
)

func SecretRegistrationJSONFormat() error {
	fmt.Println("----")
	fmt.Println("🧪     Testing: Secret registration (JSON transformation)...")

	value := `{"username": "*root*", "password": "*CasHC0w*"}`
	transform := `{"USERNAME":"*root*", "PASSWORD":"*CasHC0w*"}`

	if err := sentinel.SetJSONSecret(value, transform); err != nil {
		return errors.Wrap(err, "setJSONSecret failed")
	}

	if err := assert.WorkloadSecretHasValue(transform); err != nil {
		return errors.Wrap(err, "assertWorkloadSecretValue failed")
	}

	if err := sentinel.DeleteSecret(); err != nil {
		return errors.Wrap(err, "deleteSecret failed")
	}

	fmt.Println("🟢   PASS: Secret registration (JSON transformation) successful")
	fmt.Println("----")
	return nil
}
