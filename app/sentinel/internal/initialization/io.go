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
	"os"
	"strconv"
	"strings"

	"github.com/vmware-tanzu/secrets-manager/core/constants/sentinel"
	"github.com/vmware-tanzu/secrets-manager/core/constants/symbol"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/lib/backoff"
)

func (i *Initializer) commandFileScanner(
	cid *string) (*os.File, *bufio.Scanner) {
	filePath := i.EnvReader.InitCommandPathForSentinel()
	file, err := i.FileOpener.Open(filePath)

	if err != nil {
		i.Logger.InfoLn(
			cid,
			"RunInitCommands: no initialization file found... "+
				"skipping custom initialization.",
		)
		return nil, nil
	}

	i.Logger.TraceLn(cid, "Before parsing commands 001")

	return file, bufio.NewScanner(file)
}

func (i *Initializer) parseCommandsFile(
	ctx context.Context, cid *string, scanner *bufio.Scanner,
) {
	i.Logger.TraceLn(cid, "Before parsing commands 002")

	sc := entity.SentinelCommand{}

	if scanner == nil {
		panic("RunInitCommands: error scanning commands file")
	}

	i.Logger.TraceLn(cid, "beginning scan")
dance:
	for scanner.Scan() {
		i.Logger.TraceLn(cid, "scan:for")

		line := strings.TrimSpace(scanner.Text())
		i.Logger.TraceLn(cid, "line:", line)

		if line == "" {
			continue
		}

		parts := strings.SplitN(line, symbol.Separator, 2)

		if len(parts) == 1 && sentinel.Command(parts[0]) == sentinel.Exit {
			i.Logger.InfoLn(
				cid,
				"exit found during initialization.",
				"skipping the rest of the commands.",
				"skipping post initialization.",
			)

			// Move out of the loop to allow the keystone secret to be
			// registered.
			break dance
		}
		if len(parts) != 2 && line != symbol.LineDelimiter {
			continue
		}

		if line == symbol.LineDelimiter {
			i.Logger.TraceLn(cid, "scanner: delimiter found")
			if sc.ShouldSleep {
				i.doSleep(sc.SleepIntervalMs)
				sc = entity.SentinelCommand{}
				continue
			}

			err := backoff.RetryExponential(
				"RunInitCommands:ProcessCommandBlock",
				func() error {
					i.Logger.TraceLn(
						cid,
						"RunInitCommands:ProcessCommandBlock"+
							": retrying with exponential backoff",
					)

					err := i.Safe.Post(ctx, sc)
					if err != nil {
						i.Logger.ErrorLn(
							cid,
							"RunInitCommands:ProcessCommandBlock:error:",
							err.Error(),
						)
					}

					return err
				})

			if err != nil {
				i.Logger.ErrorLn(
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

			i.Logger.TraceLn(cid, "scanner: after delimiter")

			sc = entity.SentinelCommand{}
			continue
		}

		key := parts[0]
		value := parts[1]

		i.Logger.TraceLn(cid, "command found.", "key", key, "value", value)

		switch sentinel.Command(key) {
		case sentinel.Workload:
			sc.WorkloadIds = strings.SplitN(value, symbol.ItemSeparator, -1)
		case sentinel.Namespace:
			sc.Namespaces = strings.SplitN(value, symbol.ItemSeparator, -1)
		case sentinel.Secret:
			sc.Secret = value
		case sentinel.Transformation:
			sc.Template = value
		case sentinel.Encrypt:
			sc.Encrypt = true
		case sentinel.Remove:
			sc.DeleteSecret = true
		case sentinel.Join:
			sc.AppendSecret = true
		case sentinel.Format:
			sc.Format = value
		case sentinel.Keys:
			sc.SerializedRootKeys = value
		case sentinel.NotBefore:
			sc.NotBefore = value
		case sentinel.Expires:
			sc.Expires = value
		case sentinel.Sleep:
			sc.ShouldSleep = true
			intervalMs, err := strconv.Atoi(value)
			if err != nil {
				i.Logger.ErrorLn(cid, "RunInitCommands"+
					": Error parsing sleep interval: ", err.Error())
			}
			sc.SleepIntervalMs = intervalMs
		default:
			i.Logger.InfoLn(cid, "RunInitCommands: unknown command: ", key)
		}
	}

	i.Logger.TraceLn(cid, "scan finished")

	if err := scanner.Err(); err != nil {
		i.Logger.ErrorLn(
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
