/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import "os"

// SentinelInitCommandPath returns the path to the initialization commands file
// for VSecM Sentinel.
//
// It checks for an environment variable "VSECM_SENTINEL_INIT_COMMAND_PATH" and
// uses its value as the path. If the environment variable is not set, it
// defaults to "/opt/vsecm-sentinel/init/data".
//
// Returns:
//
//	string: The path to the Sentinel initialization commands file.
func SentinelInitCommandPath() string {
	p := os.Getenv("VSECM_SENTINEL_INIT_COMMAND_PATH")
	if p == "" {
		p = "/opt/vsecm-sentinel/init/data"
	}
	return p
}

// SentinelInitCommandTombstonePath returns the path for the VSecM Sentinel
// initialization command tombstone file.
//
// It looks for the environment variable "VSECM_SENTINEL_INIT_COMMAND_TOMBSTONE_PATH"
// and uses its value as the path. If the environment variable is not set, it
// defaults to "/opt/vsecm-sentinel/tombstone/init".
//
// This path is usually used to store a “tombstone” file or data indicating that
// the initialization command has been executed or is no longer valid.
//
// Returns:
//
//	string: The path to the Sentinel initialization command tombstone.
func SentinelInitCommandTombstonePath() string {
	p := os.Getenv("VSECM_SENTINEL_INIT_COMMAND_TOMBSTONE_PATH")
	if p == "" {
		p = "/opt/vsecm-sentinel/tombstone/init"
	}
	return p
}
