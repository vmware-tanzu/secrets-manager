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
	"github.com/vmware-tanzu/secrets-manager/sdk/core/log/level"
	"log"
)

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
