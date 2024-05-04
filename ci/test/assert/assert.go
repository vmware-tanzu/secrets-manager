/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package assert

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/secrets-manager/ci/test/io"
	"github.com/vmware-tanzu/secrets-manager/ci/test/vsecm"
	"github.com/vmware-tanzu/secrets-manager/ci/test/workload"
)

func SentinelCanEncryptSecret(value string) error {
	sentinel, err := vsecm.Sentinel()
	if err != nil || sentinel == "" || value == "" {
		return errors.Wrap(err,
			"EncryptedSecret: Failed to define sentinel or value",
		)
	}

	// Execute the encryption command within the sentinel pod.
	res, err := io.Exec("kubectl", "exec", sentinel,
		"-n", "vsecm-system", "--", "safe", "-s", value, "-e")
	if err != nil {
		return errors.Wrap(err, "EncryptedSecret: Failed to exec kubectl")
	}

	// Assert that the result exists.
	if strings.TrimSpace(res) == "" {
		return errors.New("EncryptedSecret: Encrypted secret is empty")
	}

	return nil
}

func InitContainerIsRunning() error {
	w, err := workload.Example()
	if err != nil || w == "" {
		return errors.Wrap(err, "InitContainerIsRunning: Failed to define workload")
	}

	podStatus, err := io.Exec("kubectl", "get", "pod", "-n", "default", w,
		"-o", "jsonpath={.status.initContainerStatuses[0].state.running}")
	if err != nil {
		return errors.Wrap(err, "InitContainerIsRunning: Failed to exec kubectl")
	}

	if podStatus == "" {
		return errors.New("InitContainerIsRunning: Init container is not running")
	}

	return nil
}

func WorkloadIsRunning() error {
	fmt.Println("Asserting workload is running...")

	// Define the workload pod.
	w, err := workload.Example()
	if err != nil || w == "" {
		return errors.Wrap(err, "WorkloadIsRunning: Failed to define workload")
	}

	// Fetch all pods in the 'default' namespace and count how many are in
	// 'Running' status for the defined workload.
	cmdOutput, err := io.Exec("kubectl", "get", "po", "-n", "default", "-o",
		"jsonpath={.items[*].status.phase}")
	if err != nil {
		return errors.Wrap(err, "WorkloadIsRunning: Failed to exec kubectl")
	}

	// Count how many times 'Running' appears in the command output.
	podCount := strings.Count(cmdOutput, "Running")
	if podCount == 0 {
		return errors.New("WorkloadIsRunning: No running pods found")
	}

	if podCount != 1 {
		return errors.New(
			fmt.Sprintf("Expected 1 running pod for workload, found %d", podCount),
		)
	}

	return nil
}

func WorkloadSecretHasNoValue() error {
	w, err := workload.Example()
	if err != nil || w == "" {
		return errors.Wrap(err, "WorkloadSecretHasNoValue: Failed to define workload")
	}

	// Execute the command within the workload pod to check the environment or secret.
	res, err := io.Exec("kubectl", "exec", w, "-n", "default",
		"-c", "main", "--", "./env")
	if err != nil {
		return errors.Wrap(err, "WorkloadSecretHasNoValue: Failed to exec kubectl")
	}

	res = strings.TrimSpace(res)

	if len(res) == 0 {
		return nil
	}

	// Check if the response indicates that no secret is set.
	if strings.Contains(res, "NO_SECRET") {
		return nil
	}

	return errors.New("WorkloadSecretHasNoValue: Secret should not have a value")
}

func WorkloadSecretHasValue(expectedValue string) error {
	if expectedValue == "" {
		return errors.New("WorkloadSecretHasValue: Expected value is empty")
	}

	w, err := workload.Example()
	if err != nil {
		return errors.Wrap(err, "WorkloadSecretHasValue: Failed to define workload")
	}

	res, err := io.Exec("kubectl", "exec", w, "-n", "default", "-c", "main", "--", "./env")
	if err != nil {
		return errors.Wrap(err, "WorkloadSecretHasValue: Failed to exec kubectl")
	}

	if strings.Contains(res, expectedValue) {
		return nil
	}

	return errors.New("WorkloadSecretHasValue: Secret value does not match expected")
}
