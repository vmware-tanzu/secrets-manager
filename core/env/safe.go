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
	data "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"os"
	"strconv"
	"strings"
	"time"
)

// IvInitializationIntervalForSafe fetches the Initialization Vector (IV) interval
// from an environment variable. IV is used in AES encryption.
//
// The environment variable used is VSECM_SAFE_IV_INITIALIZATION_INTERVAL.
// If the environment variable is not set or contains an invalid integer, the
// function returns a default value of 50.
// The returned value is intended to be used for rate-limiting or throttling the
// initialization of IVs.
//
// Returns:
// int: The IV initialization interval in milliseconds.
func IvInitializationIntervalForSafe() int {
	envInterval := os.Getenv("VSECM_SAFE_IV_INITIALIZATION_INTERVAL")
	if envInterval == "" {
		return 50
	}
	parsedInterval, err := strconv.Atoi(envInterval)
	if err != nil {
		return 50
	}
	return parsedInterval
}

// SecretBufferSizeForSafe returns the buffer size for the VSecM Safe secret queue.
//
// The buffer size is determined by the environment variable
// VSECM_SAFE_SECRET_BUFFER_SIZE.
//
// If the environment variable is not set, the default buffer size is 10.
// If the environment variable is set and can be parsed as an integer,
// it will be used as the buffer size.
// If the environment variable is set but cannot be parsed as an integer,
// the default buffer size is used.
func SecretBufferSizeForSafe() int {
	p := os.Getenv("VSECM_SAFE_SECRET_BUFFER_SIZE")
	if p == "" {
		return 10
	}
	l, err := strconv.Atoi(p)
	if err != nil {
		return 10
	}
	return l
}

// K8sSecretBufferSizeForSafe returns the buffer size for the VSecM Safe Kubernetes
// secret queue.
//
// The buffer size is determined by the environment variable
// VSECM_SAFE_K8S_SECRET_BUFFER_SIZE.
//
// If the environment variable is not set, the default buffer size is 10.
// If the environment variable is set and can be parsed as an integer,
// it will be used as the buffer size.
// If the environment variable is set but cannot be parsed as an integer,
// the default buffer size is used.
func K8sSecretBufferSizeForSafe() int {
	p := os.Getenv("VSECM_SAFE_K8S_SECRET_BUFFER_SIZE")
	if p == "" {
		return 10
	}
	l, err := strconv.Atoi(p)
	if err != nil {
		return 10
	}
	return l
}

// SecretDeleteBufferSizeForSafe returns the buffer size for the VSecM Safe secret
// deletion queue.
//
// The buffer size is determined by the environment variable
// VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE.
//
// If the environment variable is not set, the default buffer size is 10.
// If the environment variable is set and can be parsed as an integer,
// it will be used as the buffer size.
// If the environment variable is set but cannot be parsed as an integer,
// the default buffer size is used.
func SecretDeleteBufferSizeForSafe() int {
	p := os.Getenv("VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE")
	if p == "" {
		return 10
	}
	l, err := strconv.Atoi(p)
	if err != nil {
		return 10
	}
	return l
}

// K8sSecretDeleteBufferSizeForSafe returns the buffer size for the VSecM Safe
// Kubernetes secret deletion queue.
//
// The buffer size is determined by the environment variable
// VSECM_SAFE_K8S_SECRET_DELETE_BUFFER_SIZE.
//
// If the environment variable is not set, the default buffer size is 10.
// If the environment variable is set and can be parsed as an integer,
// it will be used as the buffer size.
// If the environment variable is set but cannot be parsed as an integer,
// the default buffer size is used.
func K8sSecretDeleteBufferSizeForSafe() int {
	p := os.Getenv("VSECM_SAFE_K8S_SECRET_DELETE_BUFFER_SIZE")
	if p == "" {
		return 10
	}
	l, err := strconv.Atoi(p)
	if err != nil {
		return 10
	}
	return l
}

// RemoveLinkedK8sSecretsModeForSafe returns a boolean indicating whether VSecM Safe
// should delete linked Kubernetes secrets when as safe managed secret is deleted.
//
// The removal of linked Kubernetes secrets is determined by the environment variable
// VSECM_SAFE_REMOVE_LINKED_K8S_SECRETS.
//
// If the environment variable is not set or its value is not "true",
// the function returns false. Otherwise, the function returns true.
func RemoveLinkedK8sSecretsModeForSafe() bool {
	p := strings.ToLower(os.Getenv("VSECM_SAFE_REMOVE_LINKED_K8S_SECRETS"))
	if p == "" {
		return false
	}

	return p == "true"
}

// FipsCompliantModeForSafe returns a boolean indicating whether VSecM Safe should run in
// FIPS compliant mode. Note that this is not a guarantee that VSecM Safe will
// run in FIPS compliant mode, as it depends on the underlying base image.
// If you are using one of the official FIPS-complaint VMware Secrets Manager Docker images,
// then it will be FIPS-compliant. Check https://vsecm.com/configuration/
// for more details.
func FipsCompliantModeForSafe() bool {
	p := strings.ToLower(os.Getenv("VSECM_SAFE_FIPS_COMPLIANT"))
	if p == "" {
		return false
	}

	return p == "true"
}

// BackingStoreForSafe returns the storage type for the data,
// as specified in the VSECM_SAFE_BACKING_STORE environment variable.
// If the environment variable is not set, it defaults to "file".
// Any value that is not "file" will mean VSecM Safe will store
// its state in-memory
func BackingStoreForSafe() data.BackingStore {
	s := os.Getenv("VSECM_SAFE_BACKING_STORE")
	if s == "" {
		return data.File
	}

	if data.BackingStore(s) == data.File {
		return data.File
	}

	return data.Memory
}

// UseKubernetesSecretsModeForSafe returns a boolean indicating whether to create a
// plain text Kubernetes secret for the workloads registered. There are two
// things to note about this approach:
//
// 1. By design, and for security the original kubernetes `Secret` should exist,
// and it should be initiated to a default data as follows:
//
//	data:
//	  # '{}' (e30=) is a special placeholder to tell Safe that the Secret
//	  # is not initialized. DO NOT deletion or change it.
//	  KEY_TXT: "e30="
//
// 2. This approach is LESS secure, and it is meant to be used for LEGACY
// systems where directly using the Safe Sidecar or Safe SDK are not feasible.
// It should be left as a last resort.
//
// If the environment variable is not set or its value is not "true",
// the function returns false. Otherwise, the function returns true.
func UseKubernetesSecretsModeForSafe() bool {
	p := os.Getenv("VSECM_SAFE_USE_KUBERNETES_SECRETS")
	if p == "" {
		return false
	}
	if strings.ToLower(p) == "true" {
		return true
	}
	return false
}

// SecretBackupCountForSafe retrieves the number of backups to keep for VSecM Safe
// secrets. If the environment variable VSECM_SAFE_SECRET_BACKUP_COUNT is not
// set or is not a valid integer, the default value of 3 will be returned.
func SecretBackupCountForSafe() int {
	p := os.Getenv("VSECM_SAFE_SECRET_BACKUP_COUNT")
	if p == "" {
		return 3
	}
	l, err := strconv.Atoi(p)
	if err != nil {
		return 3
	}
	return l
}

// RootKeyInputModeManual returns a boolean indicating whether to use manual
// cryptographic key input for VSecM Safe, instead of letting it bootstrap
// automatically. If the environment variable is not set or its value is
// not "true", the function returns false. Otherwise, the function returns true.
func RootKeyInputModeManual() bool {
	p := os.Getenv("VSECM_ROOT_KEY_INPUT_MODE_MANUAL")
	if p == "" {
		return false
	}
	if strings.ToLower(p) == "true" {
		return true
	}
	return false
}

// ManualRootKeyUpdatesK8sSecret returns a boolean indicating whether to
// update the Kubernetes secret when the root key is provided manually to VSecM Safe.
// If the environment variable is not set or its value is not "true", the function
// returns false. Otherwise, the function returns true.
func ManualRootKeyUpdatesK8sSecret() bool {
	p := os.Getenv("VSECM_MANUAL_ROOT_KEY_UPDATES_K8S_SECRET")
	if p == "" {
		return false
	}
	if strings.ToLower(p) == "true" {
		return true
	}
	return false

}

// DataPathForSafe returns the path to the safe data directory.
// The path is determined by the VSECM_SAFE_DATA_PATH environment variable.
// If the environment variable is not set, the default path "/data" is returned.
func DataPathForSafe() string {
	p := os.Getenv("VSECM_SAFE_DATA_PATH")
	if p == "" {
		p = "/data"
	}
	return p
}

// RootKeyPathForSafe returns the path to the safe age key directory.
// The path is determined by the VSECM_ROOT_KEY_PATH environment variable.
// If the environment variable is not set, the default path "/key/key.txt"
// is returned.
func RootKeyPathForSafe() string {
	p := os.Getenv("VSECM_ROOT_KEY_PATH")
	if p == "" {
		p = "/key/key.txt"
	}
	return p
}

// SourceAcquisitionTimeoutForSafe returns the timeout duration for acquiring
// a SPIFFE source bundle.
// It reads an environment variable `VSECM_SAFE_SOURCE_ACQUISITION_TIMEOUT`
// to determine the timeout.
// If the environment variable is not set, or cannot be parsed, it defaults to
// 10000 milliseconds.
//
// The returned duration is in milliseconds.
//
// Returns:
//
//	time.Duration: The time duration in milliseconds for acquiring the source.
func SourceAcquisitionTimeoutForSafe() time.Duration {
	p := os.Getenv("VSECM_SAFE_SOURCE_ACQUISITION_TIMEOUT")
	if p == "" {
		p = "10000"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 10000 * time.Millisecond
	}
	return time.Duration(i) * time.Millisecond
}

// BootstrapTimeoutForSafe returns the allowed time for VSecM Safe to wait
// before killing the pod to retrieve an SVID, in time.Duration.
// The interval is determined by the VSECM_SAFE_BOOTSTRAP_TIMEOUT environment
// variable, with a default value of 300000 milliseconds if the variable is not
// set or if there is an error in parsing the value.
func BootstrapTimeoutForSafe() time.Duration {
	p := os.Getenv("VSECM_SAFE_BOOTSTRAP_TIMEOUT")
	if p == "" {
		p = "300000"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 300000 * time.Millisecond
	}
	return time.Duration(i) * time.Millisecond
}

// RootKeySecretNameForSafe returns the name of the environment variable that holds
// the VSecM Safe age key. The value is retrieved using the
// "VSECM_ROOT_KEY_NAME" environment variable. If this variable is
// not set or is empty, the default value "vsecm-root-key" is returned.
func RootKeySecretNameForSafe() string {
	p := os.Getenv("VSECM_ROOT_KEY_NAME")
	if p == "" {
		p = "vsecm-root-key"
	}
	return p
}

// SecretNamePrefixForSafe returns the prefix to be used for the names of secrets that
// VSecM Safe stores, when it is configured to persist the secret in the Kubernetes
// cluster as Kubernetes `Secret` objects.
//
// The prefix is retrieved using the "VSECM_SAFE_SECRET_NAME_PREFIX"
// environment variable. If this variable is not set or is empty, the default
// value "vsecm-secret-" is returned.
func SecretNamePrefixForSafe() string {
	p := os.Getenv("VSECM_SAFE_SECRET_NAME_PREFIX")
	if p == "" {
		p = "vsecm-secret-"
	}
	return p
}
