/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package std

import (
	"fmt"
	"log"
	"strings"

	"github.com/vmware-tanzu/secrets-manager/core/log/level"
)

// getMaxEnvVarLength finds the maximum length of environment
// variable names dynamically.
func maxLen(envVars []string) int {
	maxLength := 0
	for _, envVar := range envVars {
		if len(envVar) > maxLength {
			maxLength = len(envVar)
		}
	}
	return maxLength
}

// printFormattedInfo prints the collected information in a formatted way,
// ensuring proper alignment.
func printFormattedInfo(id *string, info map[string]string) {
	infoKeys := sortKeys(info)
	maxLength := maxLen(infoKeys)
	idp := ""
	if id == nil {
		idp = "<nil>"
	} else {
		idp = *id
	}

	for _, key := range infoKeys {
		padding := strings.Repeat(" ", maxLength-len(key))
		fmt.Printf("%s %s%s: %s\n", idp, padding, toCustomCase(key), info[key])
	}
}

// logMessage logs a message with the specified level, correlation ID, and
// message arguments. It checks the current log level to decide if the message
// should be logged.
func logMessage(l level.Level, prefix string, correlationID *string, v ...any) {
	if l != level.Audit && level.Get() < l {
		return
	}

	args := make([]any, 0, len(v)+2)
	args = append(args, prefix)
	if correlationID != nil {
		args = append(args, *correlationID)
	}
	args = append(args, v...)

	if l == level.Fatal {
		log.Fatalln(args...)
		return
	}

	log.Println(args...)
}
