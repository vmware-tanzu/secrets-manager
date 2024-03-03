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

import "os"

// SentinelLoggerUrl retrieves the URL for the VSecM Sentinel Logger from the
// environment variable VSECM_SENTINEL_LOGGER_URL. If this environment variable
// is not set, it defaults to "localhost:50051".
//
// This url is used to configure gRPC logging service, which enables
// VSecM Sentinel's `safe` CLI command to send audit logs to the container's
// standard output.
func SentinelLoggerUrl() string {
	u := os.Getenv("VSECM_SENTINEL_LOGGER_URL")
	if u == "" {
		return "localhost:50051"
	}
	return u
}
