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
	"strconv"
	"time"

	"github.com/vmware-tanzu/secrets-manager/core/constants/env"
	"github.com/vmware-tanzu/secrets-manager/core/constants/val"
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
	p := env.Value(env.VSecMSentinelInitCommandPath)
	if p == "" {
		p = string(env.VSecMSentinelInitCommandPathDefault)
	}
	return p
}

// InitCommandRunnerWaitBeforeExecIntervalForSentinel retrieves the interval
// to wait before executing an init command stanza of Sentinel. The interval is
// determined by the environment variable
// "VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC", which is expected to contain
// an integer value representing the wait time in milliseconds.
// If the environment variable is not set or cannot be parsed, it defaults to
// zero milliseconds.
//
// Returns:
//
//	time.Duration: The wait interval in milliseconds before executing an init command.
func InitCommandRunnerWaitBeforeExecIntervalForSentinel() time.Duration {
	p := env.Value(env.VSecMSentinelInitCommandWaitBeforeExec)

	if p == "" {
		p = string(env.VSecMSentinelInitCommandWaitBeforeExecDefault)
	}

	i, _ := strconv.ParseInt(p, 10, 32)

	return time.Duration(i) * time.Millisecond
}

// InitCommandRunnerWaitIntervalBeforeInitComplete retrieves the interval
// to wait after the init command stanza of Sentinel has been completed. The
// interval is determined by the environment variable
// "VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE",
// which is expected to contain an integer value representing the wait time
// in milliseconds. If the environment variable is not set or cannot be parsed,
// it defaults to zero milliseconds.
//
// Returns:
//
//	time.Duration: The wait interval in milliseconds after initialization is
//	complete.
func InitCommandRunnerWaitIntervalBeforeInitComplete() time.Duration {
	p := env.Value(env.VSecMSentinelInitCommandWaitAfterInitComplete)
	if p == "" {
		p = string(env.VSecMSentinelInitCommandWaitAfterInitCompleteDefault)
	}

	i, _ := strconv.ParseInt(p, 10, 32)

	return time.Duration(i) * time.Millisecond
}

// OIDCProviderBaseUrlForSentinel returns the url to be used for the
// OIDC provider base URL for VSecM	Sentinel. This url is used when
// VSECM_SENTINEL_OIDC_ENABLE_RESOURCE_SERVER is set to "true".
func OIDCProviderBaseUrlForSentinel() string {
	p := env.Value(env.VSecMSentinelOidcProviderBaseUrl)
	return p
}

// SentinelEnableOIDCResourceServer is a flag that enables the OIDC resource
// server functionality in VSecM Sentinel.
func SentinelEnableOIDCResourceServer() bool {
	p := env.Value(env.VSecMSentinelOidcEnableResourceServer)
	return val.True(p)
}
