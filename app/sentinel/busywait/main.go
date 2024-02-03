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
	"fmt"
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
	fmt.Println("####### 000 in executeInitCommand !!!!!")

	cid := "VSECMSENTINEL"
	filePath := env.SentinelInitCommandPath()
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("####### 001 init command file not found")

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

	fmt.Println("####### 002 init command file FOUND")

	ctx := context.Background()

	scanner := bufio.NewScanner(file)
	var sc entity.SentinelCommand

	fmt.Println("####### 003 beginning scan")

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		fmt.Println("##### LINE")
		fmt.Println(line)
		fmt.Println("###### LINE")

		if line == "" {
			fmt.Println("####### 004 skipping empty line")
			continue
		}

		parts := strings.SplitN(line, separator, 2)

		if len(parts) != 2 && line != delimiter {
			fmt.Println("####### 006 part count mismatch")
			continue
		}

		//fmt.Println("####### 007 delimiter found")
		//fmt.Println("####### PARTS[0]: ", parts[0])
		//
		//if len(parts) >= 2 {
		//	fmt.Println("##### PARTS[1]: ", parts[1])
		//} else {
		//	fmt.Println("##### PARTS[1]: <empty>")
		//}
		//
		//if parts[0] == sleep {
		//	fmt.Println("####### 007 sleeping")
		//	milliSeconds, _ := strconv.Atoi(parts[1])
		//
		//	doSleep(milliSeconds)
		//	continue
		//}

		if line == delimiter {
			if sc.ShouldSleep {
				fmt.Println("####### 007 sleeping")
				doSleep(sc.SleepIntervalMs)
				fmt.Println("Should have slept")
				continue
			}

			fmt.Println("####### 005 processing command block")
			fmt.Println("inputKeys: ", sc.InputKeys)
			fmt.Println("workloadId: ", sc.WorkloadId)
			fmt.Println("namespace: ", sc.Namespace)
			fmt.Println("secret: ", sc.Secret)
			fmt.Println("template: ", sc.Template)
			fmt.Println("format: ", sc.Format)
			fmt.Println("encrypt: ", sc.Encrypt)
			fmt.Println("appendSecret: ", sc.AppendSecret)
			fmt.Println("notBefore: ", sc.NotBefore)
			fmt.Println("expires: ", sc.Expires)
			fmt.Println("deleteSecret: ", sc.DeleteSecret)
			fmt.Println("shouldSleep: ", sc.ShouldSleep)
			fmt.Println("###### 005 processing command block")
			processCommandBlock(ctx, sc)
			fmt.Println("####### 006 processed command block")
			sc = entity.SentinelCommand{}
			continue
		}

		key := parts[0]
		value := parts[1]

		fmt.Println("####### 008 key: ", key, " value: ", value)

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
			fmt.Println("####### 009 unknown command: ", key)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("####### 011 error reading init file")

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
