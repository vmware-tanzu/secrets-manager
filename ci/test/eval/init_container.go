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
	"time"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/secrets-manager/ci/test/assert"
	"github.com/vmware-tanzu/secrets-manager/ci/test/deploy"
	"github.com/vmware-tanzu/secrets-manager/ci/test/sentinel"
	"github.com/vmware-tanzu/secrets-manager/ci/test/workload"
)

func InitContainer() error {
	println("----")
	println("🧪     Testing: Init Container...")

	if err := deploy.WorkloadUsingInitContainer(); err != nil {
		return errors.Wrap(err, "deployWorkloadUsingInitContainer failed")
	}

	// Pause for deployment to reconcile.
	time.Sleep(10 * time.Second)

	_, err := workload.Example()
	if err != nil {
		return errors.Wrap(err, "workload.Example failed")
	}

	if err := assert.InitContainerIsRunning(); err != nil {
		return errors.Wrap(err, "assertInitContainerRunning failed")
	}

	// Additional pause just in case.
	time.Sleep(30 * time.Second)

	// Set a Kubernetes secret via Sentinel.
	if err := sentinel.SetKubernetesSecretToTriggerInitContainer(); err != nil {
		return errors.Wrap(err, "setKubernetesSecret failed")
	}

	// Assert the workload is running.
	if err := assert.WorkloadIsRunning(); err != nil {
		return errors.Wrap(err, "assertWorkloadIsRunning failed")
	}

	println("🟢   PASS: Init Container test passed.")
	println("----")
	return nil
}
