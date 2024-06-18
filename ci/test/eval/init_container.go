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
	"github.com/vmware-tanzu/secrets-manager/ci/test/deploy"
	"github.com/vmware-tanzu/secrets-manager/ci/test/sentinel"
	"github.com/vmware-tanzu/secrets-manager/ci/test/workload"
)

func InitContainer() error {
	fmt.Println("----")
	fmt.Println("🧪     Testing: Init Container...")

	if err := deploy.WorkloadUsingInitContainer(); err != nil {
		return errors.Join(
			err,
			errors.New("deployWorkloadUsingInitContainer failed"),
		)
	}

	// Pause for deployment to reconcile.
	time.Sleep(10 * time.Second)

	_, err := workload.Example()
	if err != nil {
		return errors.Join(
			err,
			errors.New("workload.Example failed"),
		)
	}

	if err := assert.InitContainerIsRunning(); err != nil {
		return errors.Join(
			err,
			errors.New("assertInitContainerRunning failed"),
		)
	}

	// Additional pause just in case.
	time.Sleep(30 * time.Second)

	// Set a Kubernetes secret via Sentinel.
	if err := sentinel.SetKubernetesSecretToTriggerInitContainer(); err != nil {
		return errors.Join(
			err,
			errors.New("setKubernetesSecret failed"),
		)
	}

	// Assert the workload is running.
	if err := assert.WorkloadIsRunning(); err != nil {
		return errors.Join(
			err,
			errors.New("assertWorkloadIsRunning failed"),
		)
	}

	fmt.Println("🟢   PASS: Init Container test passed.")
	fmt.Println("----")
	return nil
}
