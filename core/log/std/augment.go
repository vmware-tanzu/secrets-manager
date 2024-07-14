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
	"os"
)

func updateInfoWithExpectedEnvVars(
	envVarsToPrint []string, info map[string]string) []string {
	var nf []string

	for _, envVar := range envVarsToPrint {
		if value, exists := os.LookupEnv(envVar); exists {
			info[envVar] = value
			continue
		}

		nf = append(nf, envVar)
	}

	return nf
}
