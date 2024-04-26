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

func InitCommandRunnerWaitBeforeExecIntervalForSentinel() time.Duration {
	p := os.Getenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC")
	if p == "" {
		p = "0"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 0 * time.Millisecond
	}
	return time.Duration(i) * time.Millisecond
}

func InitCommandRunnerWaitIntervalBeforeInitComplete() time.Duration {
	p := os.Getenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE")
	if p == "" {
		p = "0"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 0 * time.Millisecond
	}
	return time.Duration(i) * time.Millisecond
}

// OIDCProviderBaseUrlForSentinel returns the prefix to be used for the names of secrets that
// VSecM Safe stores, when it is configured to persist the secret in the Kubernetes
// cluster as Kubernetes `Secret` objects.
//
// The prefix is retrieved using the "VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL"
// environment variable. If this variable is not set or is empty, the default
// value "" is returned.
func OIDCProviderBaseUrlForSentinel() string {
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
		return false
	}
	return p == "true"
}
