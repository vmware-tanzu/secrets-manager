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
	"context"
	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"github.com/vmware-tanzu/secrets-manager/core/probe"
	"github.com/vmware-tanzu/secrets-manager/core/system"
	"os"
	"strconv"
	"strings"
	"time"
)

const delimiter = "--"
const separator = ":"

type command string

const workload command = "w"
const namespace command = "n"
const secret command = "s"
const transformation command = "t"
const sleep = "sleep"

func processCommandBlock(ctx context.Context, sc entity.SentinelCommand) {
	// TODO: change the signature of the function to accept the context.
	safe.Post(ctx, sc)
}

func doSleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Millisecond)
}

func executeInitCommand() {
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

func main() {
	id := "VSECMSENTINEL"

	go probe.CreateLiveness()

	//Print the diagnostic information about the environment.
	envVarsToPrint := []string{"APP_VERSION", "VSECM_LOG_LEVEL",
		"VSECM_SENTINEL_SECRET_GENERATION_PREFIX"}
	log.PrintEnvironmentInfo(&id, envVarsToPrint)

	log.InfoLn(&id, "Executing the initialization commands (if any)")
	executeInitCommand()
	log.InfoLn(&id, "Initialization commands executed successfully")

	// Run on the main thread to wait forever.
	system.KeepAlive()
}
