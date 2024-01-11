/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"fmt"
	"os"
)

func proceed() (bool, string) {
	err := os.Setenv("PATH", osPath)
	if err != nil {
		fmt.Println("Error setting PATH:", err)
		return false, ""
	}

	err = os.Chdir(projectDirectory)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return false, ""
	}

	lastKnownCommitHash, err := readCommitHashFromFile()
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error reading from file:", err)
		return false, ""
	}

	currentCommitHash, err := getLatestCommitHash()
	if err != nil {
		fmt.Println("Error fetching latest commit hash:", err)
		return false, ""
	}

	if currentCommitHash == lastKnownCommitHash {
		// Nothing to do here.
		return false, ""
	}

	fmt.Println("New commit detected:", currentCommitHash)
	return true, currentCommitHash
}

func runPipeline() bool {
	if err := runCommand("make", "k8s-delete"); err != nil {
		fmt.Printf("make build-local failed: %s", err)
		return false
	}
	if err := runCommand("make", "k8s-start"); err != nil {
		fmt.Printf("make build-local failed: %s", err)
		return false
	}
	if err := setMinikubeDockerEnv(); err != nil {
		fmt.Printf("Failed to set Minikube Docker environment: %s", err)
	}
	if err := runCommand("make", "build-local"); err != nil {
		fmt.Printf("make build-local failed: %s", err)
		return false
	}
	if err := runCommand("make", "deploy-local"); err != nil {
		fmt.Printf("make deploy-local failed: %s", err)
		return false
	}
	if err := runCommand("make", "test-local-ci"); err != nil {
		fmt.Printf("make test-local failed: %s", err)
		return false
	}

	return true
}
