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
	"bufio"
	"os"
	"os/exec"
	"strings"
)

func setMinikubeDockerEnv() error {
	cmd := exec.Command("minikube", "-p", "minikube", "docker-env")
	out, err := cmd.Output()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()

		// Parse lines that look like export commands
		if strings.HasPrefix(line, "export ") {
			parts := strings.SplitN(line, " ", 3)
			if len(parts) == 3 {
				keyValue := strings.SplitN(parts[1], "=", 2)
				if len(keyValue) == 2 {
					key := keyValue[0]
					value := strings.Trim(keyValue[1], "\"") // Remove any quotes
					err := os.Setenv(key, value)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
}
