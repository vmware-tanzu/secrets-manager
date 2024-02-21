/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package rpc

import "os"

func SentinelLoggerUrl() string {
	u := os.Getenv("SENTINEL_LOGGER_URL")
	if u == "" {
		return "localhost:50051"
	}
	return u
}
