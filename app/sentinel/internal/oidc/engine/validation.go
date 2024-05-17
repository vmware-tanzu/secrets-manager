/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package engine

func invalidInput(workloads []string, encrypt bool,
	inputKeys string, secret string, deleteSecret bool) bool {

	// You need to provide a workload collection if you are not encrypting
	// a secret, or if you are not providing input keys.
	if workloads != nil && len(workloads) == 0 && !encrypt && inputKeys == "" {
		return true
	}

	// You need to provide a secret value if you are not deleting a secret,
	// or if you are not providing input keys.
	if secret == "" && !deleteSecret && inputKeys == "" {
		return true
	}

	return false
}
