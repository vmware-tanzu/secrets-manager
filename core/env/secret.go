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

func SecretGenerationPrefix() string {
	p := os.Getenv("VSECM_SAFE_SECRET_GENERATION_PREFIX")
	if p == "" {
		return "gen:"
	}
	return p
}
