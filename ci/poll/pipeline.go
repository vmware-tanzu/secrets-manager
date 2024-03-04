/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"log"
	"os"
)

func proceed() (bool, string) {
	err := os.Setenv("PATH", osPath)
	if err != nil {
		log.Println("Error setting PATH:", err)
		return false, ""
	}

	err = os.Chdir(projectDirectory)
	if err != nil {
		log.Println("Error changing directory:", err)
		return false, ""
	}

	lastKnownCommitHash, err := readCommitHashFromFile()
	if err != nil && !os.IsNotExist(err) {
		log.Println("Error reading from file:", err)
		return false, ""
	}

	currentCommitHash, err := getLatestCommitHash()
	if err != nil {
		log.Println("Error fetching latest commit hash:", err)
		return false, ""
	}

	if currentCommitHash == lastKnownCommitHash {
		// Nothing to do here.
		return false, ""
	}

	log.Println("New commit detected:", currentCommitHash)
	return true, currentCommitHash
}

func runPipeline() bool {
	if err := runCommand("make", "k8s-delete"); err != nil {
		log.Printf("make build-local failed: %s", err)
		return false
	}
	if err := runCommand("make", "k8s-start"); err != nil {
		log.Printf("make build-local failed: %s", err)
		return false
	}
	if err := setMinikubeDockerEnv(); err != nil {
		log.Printf("Failed to set Minikube Docker environment: %s", err)
	}
	if err := runCommand("make", "build-local"); err != nil {
		log.Printf("make build-local failed: %s", err)
		return false
	}
	if err := runCommand("make", "deploy-local"); err != nil {
		log.Printf("make deploy-local failed: %s", err)
		return false
	}
	if err := runCommand("make", "test-local-ci"); err != nil {
		log.Printf("make test-local failed: %s", err)
		return false
	}

	return true
}
