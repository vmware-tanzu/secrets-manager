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
	"github.com/vmware-tanzu/secrets-manager/core/constants"
)

// SpiffeIdPrefixForSentinel returns the prefix for the Safe SPIFFE ID.
// The prefix is obtained from the environment variable
// VSECM_SPIFFEID_PREFIX_SENTINEL. If the variable is not set, the default
// prefix is used.
func SpiffeIdPrefixForSentinel() string {
	p := constants.GetEnv(constants.VSecMSpiffeIdPrefixSentinel)
	if p == "" {
		p = string(constants.VSecMSpiffeIdPrefixSentinelDefault)
	}
	return p
}

// SpiffeIdPrefixForSafe returns the prefix for the Safe SPIFFE ID.
// The prefix is obtained from the environment variable
// VSECM_SPIFFEID_PREFIX_SAFE. If the variable is not set, the default prefix is
// used.
func SpiffeIdPrefixForSafe() string {
	p := constants.GetEnv(constants.VSecMSpiffeIdPrefixSafe)
	if p == "" {
		p = string(constants.VSecMSpiffeIdPrefixSafeDefault)
	}
	return p
}

// SpiffeIdPrefixForWorkload returns the prefix for the Workload's SPIFFE ID.
// The prefix is obtained from the environment variable
// VSECM_SPIFFEID_PREFIX_WORKLOAD.
// If the variable is not set, the default prefix is used.
func SpiffeIdPrefixForWorkload() string {
	p := constants.GetEnv(constants.VSecMSpiffeIdPrefixWorkload)
	if p == "" {
		p = string(constants.VSecMSpiffeIdPrefixWorkloadDefault)
	}
	return p
}
