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

// isValidSecretModification validates the secret modification request.
// It returns true if the request is valid, false otherwise.
func isValidSecretModification(
	workloads []string,
	encrypt bool,
	serializedRootKeys string,
	secret string,
	delete bool,
) bool {
	// You need to provide a workload collection if you are not encrypting
	// a secret, or if you are not providing input keys.
	if len(workloads) == 0 && !encrypt && serializedRootKeys == "" {
		return false
	}

	// You need to provide a secret value if you are not deleting a secret,
	// or if you are not providing input keys.
	if secret == "" && !delete && serializedRootKeys == "" {
		return false
	}

	return true
}
