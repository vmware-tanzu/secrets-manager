/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package rpc

import (
	"fmt"
	"strings"
	"time"
)

func currentTime() string {
	return time.Now().Local().Format(time.DateTime)
}

func build(logHeader string, correlationId *string, a ...any) string {
	logPrefix := fmt.Sprintf("%s[%s]", logHeader, currentTime())
	var messageParts []string

	if correlationId != nil {
		messageParts = append(messageParts, *correlationId)
	}

	for _, element := range a {
		messageParts = append(messageParts, fmt.Sprintf("%v", element))
	}

	message := strings.Join(messageParts, " ")
	finalLog := fmt.Sprintf("%s %s\n", logPrefix, message)

	return finalLog
}
