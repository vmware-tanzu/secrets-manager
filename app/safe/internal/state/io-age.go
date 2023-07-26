/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package state

import "strings"

func ageKeyTriplet() (string, string, string) {
	if masterKey == "" {
		return "", "", ""
	}

	parts := strings.Split(masterKey, "\n")

	if len(parts) != 3 {
		return "", "", ""
	}

	return parts[0], parts[1], parts[2]
}
