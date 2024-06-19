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
	"github.com/vmware-tanzu/secrets-manager/lib/backoff"
	"os"
	"strconv"
	"strings"

	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func commandFileScanner(cid *string) (*os.File, *bufio.Scanner) {
	filePath := env.InitCommandPathForSentinel()
	file, err := os.Open(filePath)

	if err != nil {
		log.InfoLn(
			cid,
			"RunInitCommands: no initialization file found... "+
				"skipping custom initialization.",
		)
		return nil, nil
	}

	log.TraceLn(cid, "Before parsing commands 001")

	// Parse the commands file and execute the commands in it.
	return file, bufio.NewScanner(file)
}

func parseCommandsFile(
	ctx context.Context, cid *string, scanner *bufio.Scanner,
) {
	log.TraceLn(cid, "Before parsing commands 002")

	sc := entity.SentinelCommand{}

	if scanner == nil {
		panic("RunInitCommands: error scanning commands file")
	}

	log.TraceLn(cid, "beginning scan")
dance:
	for scanner.Scan() {
		log.TraceLn(cid, "scan:for")

		line := strings.TrimSpace(scanner.Text())
		log.TraceLn(cid, "line:", line)

		if line == "" {
			continue
		}

		parts := strings.SplitN(line, separator, 2)

		if len(parts) != 2 && line != delimiter {
			continue
		}

		if line == delimiter {
			log.TraceLn(cid, "scanner: delimiter found")
			if sc.ShouldSleep {
				doSleep(sc.SleepIntervalMs)
				sc = entity.SentinelCommand{}
				continue
			}

			err := backoff.RetryExponential(
				"RunInitCommands:ProcessCommandBlock",
				func() error {
					log.TraceLn(
						cid,
						"RunInitCommands:ProcessCommandBlock"+
							": retrying with exponential backoff",
					)

					err := processCommandBlock(ctx, sc)
					if err != nil {
						log.ErrorLn(
							cid,
							"RunInitCommands:ProcessCommandBlock:error:",
							err.Error(),
						)
					}

					return err
				})

			if err != nil {
				log.ErrorLn(
					cid,
					"RunInitCommands: error processing command block: ",
					err.Error(),
				)

				// If command failed, then the initialization is not totally
				// successful.
				// Thus, it is best to crash the container to restart the
				// initialization.
				panic("RunInitCommands:ProcessCommandBlock failed")
			}

			log.TraceLn(cid, "scanner: after delimiter")

			sc = entity.SentinelCommand{}
			continue
		}

		key := parts[0]
		value := parts[1]

		log.TraceLn(cid, "command found.", "key", key, "value", value)

		switch command(key) {
		case exit:
			// exit.
			log.InfoLn(
				cid,
				"exit found during initialization.",
				"skipping the rest of the commands.",
				"skipping post initialization.",
			)

			// Move out of the loop to allow the keystone secret to be
			// registered.
			break dance
		case workload:
			sc.WorkloadIds = strings.SplitN(value, itemSeparator, -1)
		case namespace:
			sc.Namespaces = strings.SplitN(value, itemSeparator, -1)
		case secret:
			sc.Secret = value
		case transformation:
			sc.Template = value
		case encrypt:
			sc.Encrypt = true
		case remove:
			sc.DeleteSecret = true
		case join:
			sc.AppendSecret = true
		case format:
			sc.Format = value
		case keys:
			sc.SerializedRootKeys = value
		case notBefore:
			sc.NotBefore = value
		case expires:
			sc.Expires = value
		case sleep:
			sc.ShouldSleep = true
			intervalMs, err := strconv.Atoi(value)
			if err != nil {
				log.ErrorLn(cid, "RunInitCommands"+
					": Error parsing sleep interval: ", err.Error())
			}
			sc.SleepIntervalMs = intervalMs
		default:
			log.InfoLn(cid, "RunInitCommands: unknown command: ", key)
		}
	}

	log.TraceLn(cid, "scan finished")

	if err := scanner.Err(); err != nil {
		log.ErrorLn(
			cid,
			"RunInitCommands: Error reading initialization file: ",
			err.Error(),
		)

		// If command failed, then the initialization is not totally successful.
		// Thus, it is best to crash the container to restart the
		// initialization.
		panic("RunInitCommands: Error in scanning the file")
	}
}
