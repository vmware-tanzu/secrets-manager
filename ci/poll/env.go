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
	"fmt"
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

		fmt.Println("LINE")
		fmt.Println(line)

		if !strings.HasPrefix(line, "export ") {
			continue
		}

		parts := strings.SplitN(line, " ", 3)

		if len(parts) != 2 {
			continue
		}

		keyValue := strings.SplitN(parts[1], "=", 2)
		if len(keyValue) != 2 {
			continue
		}

		key := keyValue[0]
		value := strings.Trim(keyValue[1], "\"")

		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}
