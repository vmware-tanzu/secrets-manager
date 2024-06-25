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
	"github.com/vmware-tanzu/secrets-manager/core/constants/env"
)

// SpiffeIdPrefixForSentinel returns the prefix for the Safe SPIFFE ID.
// The prefix is obtained from the environment variable
// VSECM_SPIFFEID_PREFIX_SENTINEL. If the variable is not set, the default
// prefix is used.
func SpiffeIdPrefixForSentinel() string {
	p := env.Value(env.VSecMSpiffeIdPrefixSentinel)
	if p == "" {
		p = string(env.VSecMSpiffeIdPrefixSentinelDefault)
	}
	return p
}

// SpiffeIdPrefixForSafe returns the prefix for the Safe SPIFFE ID.
// The prefix is obtained from the environment variable
// VSECM_SPIFFEID_PREFIX_SAFE. If the variable is not set, the default prefix is
// used.
func SpiffeIdPrefixForSafe() string {
	p := env.Value(env.VSecMSpiffeIdPrefixSafe)
	if p == "" {
		p = string(env.VSecMSpiffeIdPrefixSafeDefault)
	}
	return p
}

// SpiffeIdPrefixForWorkload returns the prefix for the Workload's SPIFFE ID.
// The prefix is obtained from the environment variable
// VSECM_SPIFFEID_PREFIX_WORKLOAD.
// If the variable is not set, the default prefix is used.
func SpiffeIdPrefixForWorkload() string {
	p := env.Value(env.VSecMSpiffeIdPrefixWorkload)
	if p == "" {
		p = string(env.VSecMSpiffeIdPrefixWorkloadDefault)
	}
	return p
}

// NameRegExpForWorkload returns the regular expression pattern for extracting
// the workload name from the SPIFFE ID.
// The prefix is obtained from the environment variable
// VSECM_NAME_REGEXP_FOR_WORKLOAD.
// If the variable is not set, the default pattern is used.
func NameRegExpForWorkload() string {
	p := env.Value(env.VSecMWorkloadNameRegExp)
	if p == "" {
		p = string(env.VSecMNameRegExpForWorkloadDefault)
	}
	return p
}
