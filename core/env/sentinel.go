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

func SentinelInitCommandPath() string {
	p := os.Getenv("VSECM_SENTINEL_INIT_COMMAND_PATH")
	if p == "" {
		p = "/opt/vsecm-sentinel/init/data"
	}
	return p
}

// TODO: add docs and also add godocs.

func SentinelInitCommandTombstonePath() string {
	p := os.Getenv("VSECM_SENTINEL_INIT_COMMAND_TOMBSTONE_PATH")
	if p == "" {
		p = "/opt/vsecm-sentinel/tombstone/init"
	}
	return p
}
