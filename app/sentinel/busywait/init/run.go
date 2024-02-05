/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package init

import (
	"bufio"
	"context"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"os"
	"strconv"
	"strings"
)

// RunInitCommands reads and processes initialization commands from a file.
//
// This function is designed to execute an initial set of commands that are
// declaratively defined in a file. The end result will be as if an operator
// manually entered those commands using the Sentinel CLI.
//
// The function opens the file and reads it line by line using a `bufio.Scanner`.
// Each line is expected to be a command in a specific format, typically
// key-value pairs separated by a defined separator. Lines that do not conform
// to the expected format are ignored.
//
// Special handling is applied for commands that require sleeping (pause
// execution for a specified duration) or processing a block of commands.
//
// The function supports dynamic commands, which are defined in the
// 'entity.SentinelCommand' struct.
//
// Key commands include:
// - workload: (w) Sets the WorkloadId field in the SentinelCommand.
// - namespace: (n) Sets the Namespace field.
// - secret: (s) Sets the Secret field.
// - transformation: (t) Sets the Template field.
// - sleep: (sleep) Enables sleep mode and sets the SleepIntervalMs field.
//
// If the file cannot be opened, the function logs an informational message and
// returns early. Errors encountered while reading the file or closing it are
// logged as errors.
func RunInitCommands() {
	cid := "VSECMSENTINEL"
	filePath := env.SentinelInitCommandPath()
	file, err := os.Open(filePath)

	if err != nil {
		log.InfoLn(
			&cid,
			"no initialization file found… skipping custom initialization.",
		)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.ErrorLn(&cid, "Error closing initialization file: ", err.Error())
		}
	}(file)

	ctx := context.Background()

	scanner := bufio.NewScanner(file)
	var sc entity.SentinelCommand

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		parts := strings.SplitN(line, separator, 2)

		if len(parts) != 2 && line != delimiter {
			continue
		}

		if line == delimiter {
			if sc.ShouldSleep {
				doSleep(sc.SleepIntervalMs)
				sc = entity.SentinelCommand{}
				continue
			}

			processCommandBlock(ctx, sc)
			sc = entity.SentinelCommand{}
			continue
		}

		key := parts[0]
		value := parts[1]

		switch command(key) {
		case workload:
			sc.WorkloadId = value
		case namespace:
			sc.Namespace = value
		case secret:
			sc.Secret = value
		case transformation:
			sc.Template = value
		case sleep:
			sc.ShouldSleep = true
			intms, _ := strconv.Atoi(value)
			sc.SleepIntervalMs = intms
		default:
			log.InfoLn(&cid, "unknown command: ", key)
		}
	}

	if err := scanner.Err(); err != nil {
		log.ErrorLn(
			&cid,
			"Error reading initialization file: ",
			err.Error(),
		)
	}
}
