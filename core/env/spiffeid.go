/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import "os"

// SentinelSpiffeIdPrefix returns the prefix for the Safe SPIFFE ID.
// The prefix is obtained from the environment variable
// VSECM_SENTINEL_SPIFFEID_PREFIX. If the variable is not set, the default prefix is
// used.
func SentinelSpiffeIdPrefix() string {
	p := os.Getenv("VSECM_SENTINEL_SPIFFEID_PREFIX")
	if p == "" {
		p = "spiffe://vsecm.com/workload/vsecm-sentinel/ns/vsecm-system/sa/vsecm-sentinel/n/"
	}
	return p
}

// SafeSpiffeIdPrefix returns the prefix for the Safe SPIFFE ID.
// The prefix is obtained from the environment variable
// VSECM_SAFE_SPIFFEID_PREFIX. If the variable is not set, the default prefix is
// used.
func SafeSpiffeIdPrefix() string {
	p := os.Getenv("VSECM_SAFE_SPIFFEID_PREFIX")
	if p == "" {
		p = "spiffe://vsecm.com/workload/vsecm-safe/ns/vsecm-system/sa/vsecm-safe/n/"
	}
	return p
}

// WorkloadSpiffeIdPrefix returns the prefix for the WorkloadId’s SPIFFE ID.
// The prefix is obtained from the environment variable VSECM_WORKLOAD_SPIFFEID_PREFIX.
// If the variable is not set, the default prefix is used.
func WorkloadSpiffeIdPrefix() string {
	p := os.Getenv("VSECM_WORKLOAD_SPIFFEID_PREFIX")
	if p == "" {
		p = "spiffe://vsecm.com/workload/"
	}
	return p
}
