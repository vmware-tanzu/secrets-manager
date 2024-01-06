/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import (
	"os"
)

// SecretGenerationPrefix returns a prefix that’s used by VSecM Sentinel to
// generate random pattern-based secrets. If a secret is prefixed with this value,
// then VSecM sentinel will consider it as a “template” rather than a literal value.
//
// It retrieves this prefix from the environment variable
// "VSECM_SAFE_SECRET_GENERATION_PREFIX".
// If the environment variable is not set or is empty, it defaults to "gen:".
func SecretGenerationPrefix() string {
	p := os.Getenv("VSECM_SAFE_SECRET_GENERATION_PREFIX")
	if p == "" {
		return "gen:"
	}
	return p
}
