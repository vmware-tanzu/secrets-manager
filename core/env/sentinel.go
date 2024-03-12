/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import (
	"os"
	"strconv"
	"time"
)

// InitCommandPathForSentinel returns the path to the initialization commands file
// for VSecM Sentinel.
//
// It checks for an environment variable "VSECM_SENTINEL_INIT_COMMAND_PATH" and
// uses its value as the path. If the environment variable is not set, it
// defaults to "/opt/vsecm-sentinel/init/data".
//
// Returns:
//
//	string: The path to the Sentinel initialization commands file.
func InitCommandPathForSentinel() string {
	p := os.Getenv("VSECM_SENTINEL_INIT_COMMAND_PATH")
	if p == "" {
		p = "/opt/vsecm-sentinel/init/data"
	}
	return p
}

// InitCommandTombstonePathForSentinel returns the path for the VSecM Sentinel
// initialization command tombstone file.
//
// It looks for the environment variable "VSECM_SENTINEL_INIT_COMMAND_TOMBSTONE_PATH"
// and uses its value as the path. If the environment variable is not set, it
// defaults to "/opt/vsecm-sentinel/tombstone/init".
//
// This path is usually used to store a "tombstone" file or data indicating that
// the initialization command has been executed or is no longer valid.
//
// Returns:
//
//	string: The path to the Sentinel initialization command tombstone.
func InitCommandTombstonePathForSentinel() string {
	p := os.Getenv("VSECM_SENTINEL_INIT_COMMAND_TOMBSTONE_PATH")
	if p == "" {
		p = "/opt/vsecm-sentinel/tombstone/init"
	}
	return p
}

// InitCommandRunnerWaitTimeoutForSentinel initializes and returns the timeout
// duration for waiting for Sentinel to acquire an SVID.
//
// If the environment variable "VSECM_SENTINEL_INIT_COMMAND_RUNNER_WAIT_TIMEOUT"
// is set and valid, it uses the value provided by the environment variable as
// the timeout duration.
// If the environment variable is not set or invalid, a default timeout of
// 300,000 milliseconds (5 minutes)
//
// Returns:
//
//	time.Duration: The max time duration that the Sentinel will wait for an SVID.
func InitCommandRunnerWaitTimeoutForSentinel() time.Duration {
	p := os.Getenv("VSECM_SENTINEL_INIT_COMMAND_RUNNER_WAIT_TIMEOUT")
	if p == "" {
		p = "300000"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 300000 * time.Millisecond
	}
	return time.Duration(i) * time.Millisecond
}

// SentinelOIDCProviderBaseUrl returns the prefix to be used for the names of secrets that
// VSecM Safe stores, when it is configured to persist the secret in the Kubernetes
// cluster as Kubernetes `Secret` objects.
//
// The prefix is retrieved using the "VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL"
// environment variable. If this variable is not set or is empty, the default
// value "" is returned.
func SentinelOIDCProviderBaseUrl() string {
	p := os.Getenv("VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL")
	return p
}

// SentinelEnableOIDCResourceServer returns the prefix to be used for the names of secrets that
// VSecM Safe stores, when it is configured to persist the secret in the Kubernetes
// cluster as Kubernetes `Secret` objects.
//
// The prefix is retrieved using the "VSECM_SENTINEL_ENABLE_OIDC_RESOURCE_SERVER"
// environment variable. If this variable is not set or is empty, the default
// value "FALSE" is returned.
func SentinelEnableOIDCResourceServer() bool {
	p := os.Getenv("VSECM_SENTINEL_ENABLE_OIDC_RESOURCE_SERVER")
	if p == "" {
		p = "false"
	}
	return p == "true"
}
