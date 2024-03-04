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
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
)

// PrintEnvironmentInfo prints information about specific environment variables,
// enabled features, and binary version.
func PrintEnvironmentInfo(id *string, envVarsToPrint []string) {
	info := make(map[string]string)

	// Print specific environment variables and their values if envVarsToPrint is provided.
	if len(envVarsToPrint) > 0 {
		printSpecificEnvironmentVariables(id, envVarsToPrint, info)
	}

	// Print additional information.
	printAdditionalInformation(info)

	// Print the information in the desired format.
	printFormattedInfo(id, info)
}

// printSpecificEnvironmentVariables retrieves and prints values of specific
// environment variables.
func printSpecificEnvironmentVariables(id *string, envVarsToPrint []string, info map[string]string) {
	for _, envVar := range envVarsToPrint {
		if value, exists := os.LookupEnv(envVar); exists {
			info[envVar] = value
		} else {
			WarnLn(id, "Warning: Environment variable "+envVar+" not found")
		}
	}
}

// printAdditionalInformation collects and prints additional information,
// such as all environment variables and Go version.
func printAdditionalInformation(info map[string]string) {
	info["ENVIRONMENT_VARIABLES"] = strings.Join(getAllEnvironmentVariables(), ", ")
	info["GO_VERSION"] = runtime.Version()
}

// printFormattedInfo prints the collected information in a formatted way,
// ensuring proper alignment.
func printFormattedInfo(id *string, info map[string]string) {
	infoKeys := sortKeys(info)
	maxLength := getMaxEnvVarLength(infoKeys)
	idp := ""
	if id == nil {
		idp = "<nil>"
	} else {
		idp = *id
	}

	for _, key := range infoKeys {
		padding := strings.Repeat(" ", maxLength-len(key))
		fmt.Printf("%s%s%s: %s\n", idp, padding, toCustomCase(key), info[key])
	}
}

// getMaxEnvVarLength finds the maximum length of environment
// variable names dynamically.
func getMaxEnvVarLength(envVars []string) int {
	maxLength := 0
	for _, envVar := range envVars {
		if len(envVar) > maxLength {
			maxLength = len(envVar)
		}
	}
	return maxLength
}

// getAllEnvironmentVariables retrieves all environment variable keys and
// sorts them alphabetically.
func getAllEnvironmentVariables() []string {
	var envVarKeys []string
	for _, v := range os.Environ() {
		splitEnvVars := strings.Split(v, "=")
		envVarKeys = append(envVarKeys, splitEnvVars[0])
	}

	sort.Strings(envVarKeys)
	return envVarKeys
}

// toCustomCase formats a string to a custom case, replacing underscores
// with spaces and capitalizing words.
func toCustomCase(input string) string {
	return strings.ReplaceAll(strings.Title(strings.ToLower(input)), "_", " ")
}

// sortKeys returns a sorted list of keys from a map.
func sortKeys(m map[string]string) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
