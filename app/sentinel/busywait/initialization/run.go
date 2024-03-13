/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package initialization

import (
	"bufio"
	"context"
	"github.com/vmware-tanzu/secrets-manager/core/spiffe"
	"strconv"
	"strings"
	"time"

	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"os"
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
// - namespace: (n) Sets the Namespaces field.
// - secret: (s) Sets the Secret field.
// - transformation: (t) Sets the Template field.
// - sleep: (sleep) Enables sleep mode and sets the SleepIntervalMs field.
//
// If the file cannot be opened, the function logs an informational message and
// returns early. Errors encountered while reading the file or closing it are
// logged as errors.
func RunInitCommands(ctx context.Context) {
	cid := ctx.Value("correlationId").(*string)

	// ovolkan: Temporarily commenting out to debug something:
	//
	//src, acquired := spiffe.AcquireSourceForSentinel(ctx)
	//
	//if !acquired {
	//	timeout := env.InitCommandRunnerWaitTimeoutForSentinel()
	//
	//	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	//	defer cancel()
	//
	//	ticker := time.NewTicker(10 * time.Second)
	//	defer ticker.Stop()
	//
	//	for {
	//		select {
	//		case <-timeoutCtx.Done():
	//			log.ErrorLn(cid, "Failed to acquire source at RunInitCommands (1)")
	//			return
	//		case <-ticker.C:
	//			src, acquired = spiffe.AcquireSourceForSentinel(timeoutCtx)
	//			if acquired {
	//				break
	//			}
	//		}
	//	}
	//}
	//
	//if src == nil {
	//	log.ErrorLn(cid, "Failed to acquire source at RunInitCommands (2)")
	//	return
	//}

	// If we are here, then SPIFFE Workload API is functioning as expected.

	// log.ErrorLn(cid, "RunInitCommands: tick: error: ", err.Error())

	foreverCtx := context.WithoutCancel(ctx)
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	canEstablishedConnectivityToSafe := false
	src, acquired := spiffe.AcquireSourceForSentinel(foreverCtx)
	if acquired {
		err := safe.Check(foreverCtx, src)
		if err == nil {
			canEstablishedConnectivityToSafe = true
		}
	}

	for {
		if canEstablishedConnectivityToSafe {
			break
		}

		// Acquiring the source again for defensive programming.
		// At this point the cluster is (likely) still initializing.
		// Although the cached `src` should be valid and ready to be
		// used, there is no harm in acquiring a brand-new source
		// just for the sake of running initialization commands.
		src, acquired := spiffe.AcquireSourceForSentinel(foreverCtx)
		if !acquired {
			log.ErrorLn(
				cid,
				"RunInitCommands: Failed to acquire source... will retry",
			)

			continue
		}

		select {
		case <-ticker.C:
			err := safe.Check(foreverCtx, src)
			if err == nil {
				break
			}

			log.ErrorLn(cid, "RunInitCommands: tick: error: ", err.Error())
		}

		// We got what we need, move out from the infinite loop.
		break
	}

	// Parse tombstone file first:
	tombstonePath := env.InitCommandTombstonePathForSentinel()
	file, err := os.Open(tombstonePath)
	if err != nil {
		log.InfoLn(
			cid,
			"no initialization file found... skipping custom initialization.",
		)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.ErrorLn(cid, "Error closing tombstone file: ", err.Error())
		}
	}(file)

	data, err := os.ReadFile(tombstonePath)

	if strings.TrimSpace(string(data)) == "exit" {
		log.InfoLn(
			cid,
			"Initialization already exit... skipping custom initialization.",
		)
		return
	}

	filePath := env.InitCommandPathForSentinel()
	file, err = os.Open(filePath)

	if err != nil {
		log.InfoLn(
			cid,
			"no initialization file found... skipping custom initialization.",
		)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.ErrorLn(cid, "Error closing initialization file: ", err.Error())
		}
	}(file)

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
		case exit:
			// exit.
			log.InfoLn(
				cid,
				"exit found during initialization.",
				"skipping the rest of the commands.",
				"skipping post initialization.",
			)
			return
		case workload:
			sc.WorkloadId = value
		case namespace:
			sc.Namespaces = strings.SplitN(value, itemSeparator, -1)
		case secret:
			sc.Secret = value
		case transformation:
			sc.Template = value
		case sleep:
			sc.ShouldSleep = true
			intms, err := strconv.Atoi(value)
			if err != nil {
				log.ErrorLn(cid, "Error parsing sleep interval: ", err.Error())
			}
			sc.SleepIntervalMs = intms
		default:
			log.InfoLn(cid, "unknown command: ", key)
		}
	}

	if err := scanner.Err(); err != nil {
		log.ErrorLn(
			cid,
			"Error reading initialization file: ",
			err.Error(),
		)
	}

	safe.PostInitializationComplete(ctx)
}
