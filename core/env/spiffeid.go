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

import "os"

// SpiffeIdPrefixForSentinel returns the prefix for the Safe SPIFFE ID.
// The prefix is obtained from the environment variable
// VSECM_SENTINEL_SPIFFEID_PREFIX. If the variable is not set, the default prefix is
// used.
func SpiffeIdPrefixForSentinel() string {
	p := os.Getenv("VSECM_SENTINEL_SPIFFEID_PREFIX")
	if p == "" {
		p = "spiffe://vsecm.com/workload/vsecm-sentinel/ns/vsecm-system/sa/vsecm-sentinel/n/"
	}
	return p
}

// SpiffeIdPrefixForSafe returns the prefix for the Safe SPIFFE ID.
// The prefix is obtained from the environment variable
// VSECM_SAFE_SPIFFEID_PREFIX. If the variable is not set, the default prefix is
// used.
func SpiffeIdPrefixForSafe() string {
	p := os.Getenv("VSECM_SAFE_SPIFFEID_PREFIX")
	if p == "" {
		p = "spiffe://vsecm.com/workload/vsecm-safe/ns/vsecm-system/sa/vsecm-safe/n/"
	}
	return p
}

// SpiffeIdPrefixForWorkload returns the prefix for the WorkloadId's SPIFFE ID.
// The prefix is obtained from the environment variable VSECM_WORKLOAD_SPIFFEID_PREFIX.
// If the variable is not set, the default prefix is used.
func SpiffeIdPrefixForWorkload() string {
	p := os.Getenv("VSECM_WORKLOAD_SPIFFEID_PREFIX")
	if p == "" {
		p = "spiffe://vsecm.com/workload/"
	}
	return p
}
