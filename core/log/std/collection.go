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
	"sort"
	"strings"
)

// sortKeys returns a sorted list of keys from a map.
func sortKeys(m map[string]string) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func envVars() []string {
	var envVarKeys []string

	for _, v := range os.Environ() {
		splitEnvVars := strings.Split(v, "=")
		envVarKeys = append(envVarKeys, splitEnvVars[0])
	}

	sort.Strings(envVarKeys)

	return envVarKeys
}
