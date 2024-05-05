/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package validation

import (
	"fmt"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"strings"
)

// IsSentinel returns true if the given SPIFFEID is a Sentinel ID.
// It does this by checking if the SPIFFEID has the SpiffeIdPrefixForSentinel as its prefix.
func IsSentinel(spiffeid string) bool {
	return strings.HasPrefix(spiffeid, env.SpiffeIdPrefixForSentinel())
}

// IsSafe returns true if the given SPIFFEID is a Safe ID.
// It does this by checking if the SPIFFEID has the SpiffeIdPrefixForSafe as its prefix.
func IsSafe(spiffeid string) bool {
	return strings.HasPrefix(spiffeid, env.SpiffeIdPrefixForSafe())
}

// IsWorkload returns true if the given SPIFFEID is a WorkloadIds ID.
// It does this by checking if the SPIFFEID has the SpiffeIdPrefixForWorkload as its prefix.
func IsWorkload(spiffeid string) bool {
	return strings.HasPrefix(spiffeid, env.SpiffeIdPrefixForWorkload())
}

func EnsureSafe(source *workloadapi.X509Source) {
	svid, err := source.GetX509SVID()
	if err != nil {
		panic(
			fmt.Sprintf(
				"Unable to get X.509 SVID from source bundle: %s",
				err.Error(),
			),
		)
	}

	svidId := svid.ID
	if !IsSafe(svidId.String()) {
		panic(
			fmt.Sprintf(
				"SpiffeId check: I don't know you, and it's crazy: %s",
				svidId.String(),
			),
		)
	}
}
