/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

// Package std provides a simple and flexible logging library with various
// log levels.
package std

import "github.com/vmware-tanzu/secrets-manager/sdk/core/log/level"

// InfoLn logs an info level message.
func InfoLn(correlationID *string, v ...any) {
	logMessage(level.Info, "[INFO]", correlationID, v...)
}

// TraceLn logs a trace level message.
func TraceLn(correlationID *string, v ...any) {
	logMessage(level.Trace, "[TRACE]", correlationID, v...)
}
