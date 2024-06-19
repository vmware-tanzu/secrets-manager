/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package data

// VSecMInternalCommand is the command that VSecM uses to perform
// internal operations.
type VSecMInternalCommand struct {
	LogLevel int `json:"logLevel"`
}

// SentinelCommand is the command that VSecM Sentinel uses to perform
// REST API operations on VSecM Safe.
type SentinelCommand struct {
	WorkloadIds        []string
	Namespaces         []string
	Secret             string
	Template           string
	DeleteSecret       bool
	AppendSecret       bool
	Format             string
	Encrypt            bool
	NotBefore          string
	Expires            string
	SerializedRootKeys string
	ShouldSleep        bool
	SleepIntervalMs    int
}
