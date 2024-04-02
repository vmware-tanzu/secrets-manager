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
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"

	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
	"github.com/vmware-tanzu/secrets-manager/core/backoff"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/spiffe"
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
// - workload: (w) Sets the WorkloadIds field in the SentinelCommand.
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

	// TODO: remove InitCommandRunnerWaitTimeoutForSentinel()
	// we don't need it. If the init commands cannot run either the
	// commands are corrupt (then it needs fixing) or there is a
	// connectivity issue (then it needs retry).
	// Init commands should be reliable, NOT half-baked.

	for {
		s := backoff.Strategy{
			MaxRetries:  20,
			Delay:       1000,
			Exponential: true,
			MaxDuration: 30 * time.Second,
		}

		err := backoff.Retry("vsecm-system", func() error {
			_, acquired := spiffe.AcquireSourceForSentinel(ctx)
			if !acquired {
				return errors.New("failed to acquire source")
			}

			return nil
		}, s)

		if err == nil {
			break
		}
	}

	// Now, we are sure that we can acquire a source.
	// Try to do a fetch with the source.

	for {
		s := backoff.Strategy{
			MaxRetries:  20,
			Delay:       1000,
			Exponential: true,
			MaxDuration: 30 * time.Second,
		}

		err := backoff.Retry("vsecm-system", func() error {
			src, acquired := spiffe.AcquireSourceForSentinel(ctx)
			if !acquired {
				return errors.New("failed to acquire source")
			}

			if err := safe.Check(ctx, src); err == nil {
				return errors.New("cannot establish connection to safe")
			}

			return nil
		}, s)

		if err == nil {
			break
		}
	}

	// Now we know that we can establish a connection to VSecM Safe.
	// We can safely run init commands.

	// Parse tombstone file first:
	tombstonePath := env.InitCommandTombstonePathForSentinel()
	file, err := os.Open(tombstonePath)
	if err != nil {
		log.InfoLn(
			cid,
			"RunInitCommands: no tombstone file found... skipping custom initialization.",
		)

		// TODO: optionally crash.
		// the crash should be based on a config var, and it should be off by default.

		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.ErrorLn(cid, "Error closing tombstone file: ", err.Error())
		}
	}(file)

	data, err := os.ReadFile(tombstonePath)

	log.InfoLn(cid, fmt.Sprintf("tombstone:'%s'", string(data)))

	if strings.TrimSpace(string(data)) == "complete" {
		log.InfoLn(
			cid,
			"RunInitCommands: Already initialized. Skipping custom initialization.",
		)
		return
	}

	filePath := env.InitCommandPathForSentinel()
	file, err = os.Open(filePath)

	if err != nil {
		log.InfoLn(
			cid,
			"RunInitCommands: no initialization file found... skipping custom initialization.",
		)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.ErrorLn(cid, "RunInitCommands: Error closing initialization file: ", err.Error())
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var sc entity.SentinelCommand

dance:
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

			// TODO: get some of these from env vars.
			s := backoff.Strategy{
				MaxRetries:  20,
				Delay:       1000,
				Exponential: true,
				MaxDuration: 30 * time.Second,
			}

			err = backoff.Retry("vsecm-system", func() error {
				log.TraceLn(cid, "RunInitCommands: processCommandBlock: retrying with exponential backoff")

				return processCommandBlock(ctx, sc)
			}, s)

			if err != nil {
				log.ErrorLn(cid, "RunInitCommands: error processing command block: ", err.Error())

				// TODO: optionally crash.
				// the crash should be based on a config var, and it should be off by default.
			}

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
			// Move out of the loop to allow the keystone secret to be registered.
			break dance
		case workload:
			sc.WorkloadIds = strings.SplitN(value, itemSeparator, -1)
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
				log.ErrorLn(cid, "RunInitCommands: Error parsing sleep interval: ", err.Error())
			}
			sc.SleepIntervalMs = intms
		default:
			log.InfoLn(cid, "RunInitCommands: unknown command: ", key)
		}
	}

	if err := scanner.Err(); err != nil {
		log.ErrorLn(
			cid,
			"RunInitCommands: Error reading initialization file: ",
			err.Error(),
		)
	}

	// TODO: get some of these from env vars.
	s := backoff.Strategy{
		MaxRetries:  20,
		Delay:       1000,
		Exponential: true,
		MaxDuration: 30 * time.Second,
	}

	err = backoff.Retry("vsecm-system", func() error {
		log.TraceLn(cid, "RunInitCommands: processCommandBlock: retrying with exponential backoff")

		// Assign a secret for VSecM Keystone
		return processCommandBlock(ctx, entity.SentinelCommand{
			WorkloadIds: []string{"vsecm-keystone"},
			Namespaces:  []string{"vsecm-system"},
			Secret:      "keystone-init",
		})
	}, s)

	if err != nil {
		log.ErrorLn(cid, "RunInitCommands: error setting keystone secret: ", err.Error())

		// TODO: optionally crash.
		// the crash should be based on a config var, and it should be off by default.
	} else {
		log.InfoLn(cid, "RunInitCommands: keystone secret set successfully.")
		safe.PostInitializationComplete(ctx)
	}
}
