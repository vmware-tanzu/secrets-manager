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
	"strconv"
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
	envInterval := constants.GetEnv(constants.VSecMSafeIvInitializationInterval)
	d, _ := strconv.Atoi(string(constants.VSecMSafeIvInitializationIntervalDefault))

	if envInterval == "" {
		return d
	}

	parsedInterval, err := strconv.Atoi(envInterval)
	if err != nil {
		return d
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
	p := constants.GetEnv(constants.VSecMSafeSecretBufferSize)
	d, _ := strconv.Atoi(string(constants.VSecMSafeSecretBufferSizeDefault))

	if p == "" {
		return d
	}

	l, err := strconv.Atoi(p)
	if err != nil {
		return d
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
	p := constants.GetEnv(constants.VSecMSafeK8sSecretBufferSize)
	d, _ := strconv.Atoi(string(constants.VSecMSafeK8sSecretBufferSizeDefault))

	if p == "" {
		return d
	}

	l, err := strconv.Atoi(p)
	if err != nil {
		return d
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
	p := constants.GetEnv(constants.VSecMSafeSecretDeleteBufferSize)
	d, _ := strconv.Atoi(string(constants.VSecMSafeSecretDeleteBufferSizeDefault))

	if p == "" {
		return d
	}

	l, err := strconv.Atoi(p)
	if err != nil {
		return d
	}

	return l
}

// FipsCompliantModeForSafe returns a boolean indicating whether VSecM Safe
// should run in FIPS compliant mode. Note that this is not a guarantee that
// VSecM Safe will run in FIPS compliant mode, as it depends on the underlying
// base image. If you are using one of the official FIPS-complaint
// VMware Secrets Manager Docker images, then it will be FIPS-compliant.
// Check https://vsecm.com/configuration/ for more details.
func FipsCompliantModeForSafe() bool {
	p := constants.GetEnv(constants.VSecMSafeFipsCompliant)

	return constants.True(p)
}

// SecretBackupCountForSafe retrieves the number of backups to keep for VSecM
// Safe secrets. If the environment variable VSECM_SAFE_SECRET_BACKUP_COUNT
// is not set or is not a valid integer, the default value of 3 will be returned.
//
// Note: there are plans to deprecate this feature in the future in favor of
// a more robust database-driven changelog solution for secrets.
func SecretBackupCountForSafe() int {
	p := constants.GetEnv(constants.VSecMSafeSecretBackupCount)
	d, _ := strconv.Atoi(string(constants.VSecMSafeSecretBackupCountDefault))

	if p == "" {
		return d
	}

	l, err := strconv.Atoi(p)
	if err != nil {
		return d
	}

	return l
}

// RootKeyInputModeManual returns a boolean indicating whether to use manual
// cryptographic key input for VSecM Safe, instead of letting it bootstrap
// automatically. If the environment variable is not set or its value is
// not "true", the function returns false. Otherwise, the function returns true.
func RootKeyInputModeManual() bool {
	p := constants.GetEnv(constants.VSecMRootKeyInputModeManual)

	return constants.True(p)
}

// DataPathForSafe returns the path to the safe data directory.
// The path is determined by the VSECM_SAFE_DATA_PATH environment variable.
// If the environment variable is not set, the default path "/var/local/vsecm/data"
// is returned.
func DataPathForSafe() string {
	p := constants.GetEnv(constants.VSecMSafeDataPath)
	if p == "" {
		p = string(constants.VSecMSafeDataPathDefault)
	}
	return p
}

// RootKeyPathForSafe returns the path to the safe age key directory.
// The path is determined by the VSECM_ROOT_KEY_PATH environment variable.
// If the environment variable is not set, the default path "/key/key.txt"
// is returned.
func RootKeyPathForSafe() string {
	p := constants.GetEnv(constants.VSecMRootKeyPath)
	if p == "" {
		p = string(constants.VSecMRootKeyPathDefault)
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
	p := constants.GetEnv(constants.VSecMSafeSourceAcquisitionTimeout)
	if p == "" {
		p = string(constants.VSecMSafeSourceAcquisitionTimeoutDefault)
	}

	i, _ := strconv.ParseInt(p, 10, 32)

	return time.Duration(i) * time.Millisecond
}

// BootstrapTimeoutForSafe returns the allowed time for VSecM Safe to wait
// before killing the pod to retrieve an SVID, in time.Duration.
// The interval is determined by the VSECM_SAFE_BOOTSTRAP_TIMEOUT environment
// variable, with a default value of 300000 milliseconds if the variable is not
// set or if there is an error in parsing the value.
func BootstrapTimeoutForSafe() time.Duration {
	p := constants.GetEnv(constants.VSecMSafeBootstrapTimeout)
	if p == "" {
		p = string(constants.VSecMSafeBootstrapTimeoutDefault)
	}

	i, _ := strconv.ParseInt(p, 10, 32)

	return time.Duration(i) * time.Millisecond
}

// RootKeySecretNameForSafe returns the name of the environment variable that
// holds the VSecM Safe age key. The value is retrieved using the
// "VSECM_ROOT_KEY_NAME" environment variable. If this variable is
// not set or is empty, the default value "vsecm-root-key" is returned.
func RootKeySecretNameForSafe() string {
	p := constants.GetEnv(constants.VSecMRootKeyName)
	if p == "" {
		p = string(constants.VSecMRootKeyNameDefault)
	}
	return p
}
