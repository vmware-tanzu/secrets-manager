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
	"os"
)

// SecretGenerationPrefix returns a prefix that's used by VSecM Sentinel to
// generate random pattern-based secrets. If a secret is prefixed with this value,
// then VSecM Sentinel will consider it as a "template" rather than a literal value.
//
// It retrieves this prefix from the environment variable
// "VSECM_SENTINEL_SECRET_GENERATION_PREFIX".
// If the environment variable is not set or is empty, it defaults to "gen:".
func SecretGenerationPrefix() string {
	p := os.Getenv("VSECM_SENTINEL_SECRET_GENERATION_PREFIX")
	if p == "" {
		return "gen:"
	}
	return p
}

// StoreWorkloadAsK8sSecretPrefix retrieves the prefix for storing workload data
// as a Kubernetes secret.
//
// It fetches the value of the environment variable
// VSECM_SAFE_STORE_WORKLOAD_AS_K8S_SECRET_PREFIX.
// If this environment variable is not set or is empty, it defaults to "k8s:".
//
// This way, you can use VSecM to generate Kubernetes Secrets instead of
// associating secrets to workloads. This approach is especially useful in
// legacy use case where you cannot use VSecM SDK, or VSecM Sidecar
// to associate secrets to workloads, or doing so is not feasible because it
// would introduce deviation from the upstream dependencies.
//
// Returns:
//   - A string representing the prefix for Kubernetes secrets.
//     The default value is "k8s:" if the environment variable is not set or empty.
func StoreWorkloadAsK8sSecretPrefix() string {
	p := os.Getenv("VSECM_SAFE_STORE_WORKLOAD_AS_K8S_SECRET_PREFIX")
	if p == "" {
		return "k8s:"
	}
	return p
}
