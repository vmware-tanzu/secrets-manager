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

const projectDirectory = "/home/aegis/WORKSPACE/VSecM"
const osPath = "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/snap/bin:/usr/local/go/bin:/usr/local/go/bin:/usr/local/go/bin"

func main() {
	err := createLockFile()
	if err != nil {
		fmt.Println("Another instance is already running. Exiting.")
		return
	}
	defer removeLockFile()

	err = os.Setenv("PATH", osPath)
	if err != nil {
		fmt.Println("Error setting PATH:", err)
		return
	}

	err = os.Chdir(projectDirectory)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	lastKnownCommitHash, err := readCommitHashFromFile()
	if err != nil && !os.IsNotExist(err) {
		fmt.Println("Error reading from file:", err)
		return
	}

	currentCommitHash, err := getLatestCommitHash()
	if err != nil {
		fmt.Println("Error fetching latest commit hash:", err)
		return
	}

	if currentCommitHash == lastKnownCommitHash {
		// Nothing to do here.
		return
	}

	fmt.Println("New commit detected:", currentCommitHash)

	if err := runCommand("make", "k8s-delete"); err != nil {
		fmt.Printf("make build-local failed: %s", err)
		return
	}
	if err := runCommand("make", "k8s-start"); err != nil {
		fmt.Printf("make build-local failed: %s", err)
		return
	}
	if err := setMinikubeDockerEnv(); err != nil {
		fmt.Printf("Failed to set Minikube Docker environment: %s", err)
	}
	if err := runCommand("make", "build-local"); err != nil {
		fmt.Printf("make build-local failed: %s", err)
		return
	}
	if err := runCommand("make", "deploy-local"); err != nil {
		fmt.Printf("make deploy-local failed: %s", err)
		return
	}
	if err := runCommand("make", "test-local"); err != nil {
		fmt.Printf("make test-local failed: %s", err)
		return
	}

	err = writeCommitHashToFile(currentCommitHash)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
