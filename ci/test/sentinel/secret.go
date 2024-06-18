/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package sentinel

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/vmware-tanzu/secrets-manager/ci/test/io"
	"github.com/vmware-tanzu/secrets-manager/ci/test/vsecm"
	"github.com/vmware-tanzu/secrets-manager/ci/test/wait"
)

func DeleteSecret() error {
	sentinel, err := vsecm.Sentinel()
	if err != nil || sentinel == "" {
		return errors.Join(
			err,
			errors.New("deleteSecret: Failed to define sentinel"),
		)
	}

	_, err = io.Exec("kubectl", "exec", sentinel, "-n", "vsecm-system",
		"--", "safe", "-w", "example", "-n", "default", "-d")
	if err != nil {
		return errors.Join(
			err,
			errors.New("deleteSecret: Failed to exec kubectl"),
		)
	}

	return nil
}

func SetKubernetesSecretToTriggerInitContainer() error {

	// Define the sentinel pod.
	sentinel, err := vsecm.Sentinel()
	if err != nil {
		return fmt.Errorf(
			"SetKubernetesSecretToTriggerInitContainer: "+
				"Failed to define sentinel: %w", err)
	}
	if sentinel == "" {
		return errors.New("SetKubernetesSecretToTriggerInitContainer: " +
			"Failed to define sentinel")
	}

	// Construct the command to execute within the sentinel pod.
	secretData := `{"username": "root", "password": "SuperSecret", "value": "VSecMRocks"}`
	transformTemplate := `{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}", "VALUE": "{{.value}}"}`

	_, err = io.Exec("kubectl", "exec", sentinel, "-n", "vsecm-system",
		"--", "safe", "-w", "example", "-n", "default",
		"-s", secretData, "-t", transformTemplate,
	)
	if err != nil {
		return errors.Join(
			err,
			errors.New("setKubernetesSecretToTriggerInitContainer:"+
				" Failed to exec kubectl"),
		)
	}

	// Wait for the workload to be ready.
	if err := wait.ForExampleWorkload(); err != nil {
		return fmt.Errorf("set_kubernetes_secret:"+
			" Failed to wait for workload readiness: %w", err)
	}

	fmt.Println("done: set_kubernetes_secret()")
	return nil
}

func SetSecret(value string) error {
	if value == "" {
		return errors.New("SetSecret: Value is empty")
	}

	sentinel, err := vsecm.Sentinel()
	if err != nil || sentinel == "" {
		return errors.Join(
			err,
			errors.New("setSecret: Failed to define sentinel"),
		)
	}

	// Executing command within the sentinel pod to set the secret.
	_, err = io.Exec("kubectl", "exec", sentinel, "-n", "vsecm-system",
		"--", "safe", "-w", "example", "-n", "default", "-s", value)
	if err != nil {
		return errors.Join(
			err,
			errors.New("setSecret: Failed to exec kubectl"),
		)
	}

	return nil
}

func SetEncryptedSecret(value string) error {
	if value == "" {
		return errors.New("SetEncryptedSecret: Value is empty")
	}

	sentinel, err := vsecm.Sentinel()
	if err != nil || sentinel == "" {
		return errors.Join(
			err,
			errors.New("setEncryptedSecret: Failed to define sentinel"),
		)
	}

	// Execute the command to encrypt and set the secret.
	res, err := io.Exec("kubectl", "exec", sentinel, "-n", "vsecm-system",
		"--", "safe", "-s", value, "-e")
	if err != nil {
		return errors.Join(
			err,
			errors.New("setEncryptedSecret: Failed to exec kubectl"),
		)
	}

	if res == "" {
		return errors.New("SetEncryptedSecret: Encrypted secret is empty")
	}

	lines := strings.Split(res, "\n")
	out := ""
	// Remove the lines that do not contain the secret to encrypt.
	for _, line := range lines {
		logLinePattern := regexp.MustCompile(`\[(INFO|DEBUG|WARN|ERROR)]`)
		if !logLinePattern.MatchString(line) && strings.TrimSpace(line) != "" {
			out += line
		}
	}

	// Assuming res is the encrypted secret, now setting it.
	_, err = io.Exec("kubectl", "exec", sentinel, "-n", "vsecm-system",
		"--", "safe", "-w", "example", "-n", "default", "-s", out, "-e")
	if err != nil {
		return errors.Join(
			err,
			errors.New("setEncryptedSecret: Failed to set encrypted secret"),
		)
	}

	return nil
}

func SetJSONSecret(value, transform string) error {
	if value == "" || transform == "" {
		return errors.New("setJSONSecret: Value or transform is empty")
	}

	sentinel, err := vsecm.Sentinel()
	if err != nil || sentinel == "" {
		return errors.Join(
			err,
			errors.New("setJSONSecret: Failed to define sentinel"),
		)
	}

	// Executing command within the sentinel pod to set the JSON secret with transformation.
	_, err = io.Exec("kubectl", "exec", sentinel, "-n", "vsecm-system",
		"--", "safe", "-w", "example", "-n", "default",
		"-s", value, "-t", transform, "-f", "json")
	if err != nil {
		return errors.Join(
			err,
			errors.New("setJSONSecret: Failed to exec kubectl"),
		)
	}

	return nil
}

func AppendSecret(value string) error {
	if value == "" {
		return errors.New("AppendSecret: Value is empty")
	}

	sentinel, err := vsecm.Sentinel()
	if err != nil || sentinel == "" {
		return errors.Join(
			err,
			errors.New("appendSecret: Failed to define sentinel"),
		)
	}

	// Executing command within the sentinel pod to append the secret.
	_, err = io.Exec("kubectl", "exec", sentinel, "-n", "vsecm-system",
		"--", "safe", "-w", "example", "-n", "default", "-a", "-s", value)
	if err != nil {
		return errors.Join(
			err,
			errors.New("appendSecret: Failed to exec kubectl"),
		)
	}

	return nil
}
