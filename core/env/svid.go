/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import "os"

// SentinelSvidPrefix returns the prefix for the Safe
// SVID (Short-lived Verifiable Identity Document) used in the VSecM system.
// The prefix is obtained from the environment variable
// VSECM_SENTINEL_SVID_PREFIX. If the variable is not set, the default prefix is
// used.
func SentinelSvidPrefix() string {
	p := os.Getenv("VSECM_SENTINEL_SVID_PREFIX")
	if p == "" {
		p = "spiffe://vsecm.com/workload/vsecm-sentinel/ns/vsecm-system/sa/vsecm-sentinel/n/"
	}
	return p
}

// SafeSvidPrefix returns the prefix for the Safe
// SVID (Short-lived Verifiable Identity Document) used in the VSecM system.
// The prefix is obtained from the environment variable
// VSECM_SAFE_SVID_PREFIX. If the variable is not set, the default prefix is
// used.
func SafeSvidPrefix() string {
	p := os.Getenv("VSECM_SAFE_SVID_PREFIX")
	if p == "" {
		p = "spiffe://vsecm.com/workload/vsecm-safe/ns/vsecm-system/sa/vsecm-safe/n/"
	}
	return p
}

// WorkloadSvidPrefix returns the prefix for the Workload SVID
// (SPIFFE Verifiable Identity Document) used in the VSecM system.
// The prefix is obtained from the environment variable VSECM_WORKLOAD_SVID_PREFIX.
// If the variable is not set, the default prefix is used.
func WorkloadSvidPrefix() string {
	p := os.Getenv("VSECM_WORKLOAD_SVID_PREFIX")
	if p == "" {
		p = "spiffe://vsecm.com/workload/"
	}
	return p
}
