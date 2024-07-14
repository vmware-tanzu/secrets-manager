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

// PrintEnvironmentInfo prints information about specific environment variables,
// enabled features, and app's version.
func PrintEnvironmentInfo(id *string, envVarsToExpect []string) {
	info := make(map[string]string)

	notFound := updateInfoWithExpectedEnvVars(envVarsToExpect, info)
	for _, v := range notFound {
		WarnLn(id, "Environment variable '"+v+"' not found")
	}

	printFormattedInfo(id, info)
}
