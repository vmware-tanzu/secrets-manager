/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package key

import (
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"strings"
)

// RootKeyTriplet splits the RootKey into three components, if it is properly
// formatted.
//
// The function returns a triplet of strings representing the parts of the RootKey,
// separated by newlines. If the RootKey is empty or does not contain exactly
// three parts, the function returns three empty strings.
func RootKeyTriplet() (string, string, string) {
	state.RootKeyLock.RLock()
	defer state.RootKeyLock.RUnlock()

	if state.RootKey == "" {
		return "", "", ""
	}

	parts := strings.Split(state.RootKey, "\n")

	if len(parts) != 3 {
		return "", "", ""
	}

	return parts[0], parts[1], parts[2]
}
