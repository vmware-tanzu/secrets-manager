/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package deploy

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/vmware-tanzu/secrets-manager/ci/test/io"
)

func WorkloadUsingSDK() error {
	fmt.Println("Deploying workload that uses the SDK...")

	origin := os.Getenv("ORIGIN")
	var deployCommand string

	switch origin {
	case "remote":
		deployCommand = "make example-sdk-deploy"
	case "eks":
		deployCommand = "make example-sdk-deploy-eks"
	default:
		deployCommand = "make example-sdk-deploy-local"
	}

	// Splitting command and arguments
	parts := strings.Fields(deployCommand)
	command := parts[0]
	args := parts[1:]

	if output, err := io.Exec(command, args...); err != nil {
		return fmt.Errorf("deploy command failed: %w\nOutput: %s", err, output)
	}

	// Wait for the workload to be ready.
	if err := io.Wait(10); err != nil {
		return fmt.Errorf("waiting for workload failed: %w", err)
	}

	fmt.Println("Deployed workload that uses the SDK.")
	return nil
}

func WorkloadUsingSidecar() error {
	fmt.Println("Deploying workload that uses the sidecar...")

	origin := os.Getenv("ORIGIN")
	var command string

	switch origin {
	case "remote":
		command = "make example-sidecar-deploy"
	case "eks":
		command = "make example-sidecar-deploy-eks"
	default:
		command = "make example-sidecar-deploy-local"
	}

	// Splitting command and arguments for executeCommand
	parts := strings.Fields(command)
	cmd, args := parts[0], parts[1:]

	if output, err := io.Exec(cmd, args...); err != nil {
		return fmt.Errorf("deploy command failed: %v\nOutput: %s", err, output)
	}

	// Wait for the workload to be ready.
	if err := io.Wait(10); err != nil {
		return fmt.Errorf("waiting for workload failed: %w", err)
	}

	fmt.Println("Deployed workload that uses the sidecar.")
	return nil
}

func WorkloadUsingInitContainer() error {
	fmt.Println("Deploying workload that uses the init container...")

	origin := os.Getenv("ORIGIN")
	var command string

	// Determine the deployment command based on the ORIGIN environment variable.
	switch origin {
	case "remote":
		command = "make example-init-container-deploy"
	case "eks":
		command = "make example-init-container-deploy-eks"
	default:
		command = "make example-init-container-deploy-local"
	}

	// Split the command to pass to executeCommand.
	parts := strings.Fields(command)
	cmd, args := parts[0], parts[1:]

	// Execute the deployment command.
	output, err := io.Exec(cmd, args...)
	if err != nil {
		return fmt.Errorf("deployment command failed: %v\nOutput: %s", err, output)
	}

	// Pause for deployment to settle.
	time.Sleep(10 * time.Second)

	return nil
}
