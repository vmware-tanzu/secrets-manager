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
	"strconv"
	"time"

	"github.com/vmware-tanzu/secrets-manager/core/constants/env"
	"github.com/vmware-tanzu/secrets-manager/core/constants/val"
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
	envInterval := env.Value(env.VSecMSafeIvInitializationInterval)
	d, _ := strconv.Atoi(string(env.VSecMSafeIvInitializationIntervalDefault))

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
	p := env.Value(env.VSecMSafeSecretBufferSize)
	d, _ := strconv.Atoi(string(env.VSecMSafeSecretBufferSizeDefault))

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
	p := env.Value(env.VSecMSafeK8sSecretBufferSize)
	d, _ := strconv.Atoi(string(env.VSecMSafeK8sSecretBufferSizeDefault))

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
	p := env.Value(env.VSecMSafeSecretDeleteBufferSize)
	d, _ := strconv.Atoi(string(env.VSecMSafeSecretDeleteBufferSizeDefault))

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
	p := env.Value(env.VSecMSafeFipsCompliant)

	return val.True(p)
}

// SecretBackupCountForSafe retrieves the number of backups to keep for VSecM
// Safe secrets. If the environment variable VSECM_SAFE_SECRET_BACKUP_COUNT
// is not set or is not a valid integer, the default value of 3 will be returned.
//
// Note: there are plans to deprecate this feature in the future in favor of
// a more robust database-driven changelog solution for secrets.
func SecretBackupCountForSafe() int {
	p := env.Value(env.VSecMSafeSecretBackupCount)
	d, _ := strconv.Atoi(string(env.VSecMSafeSecretBackupCountDefault))

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
	p := env.Value(env.VSecMRootKeyInputModeManual)

	return val.True(p)
}

// DataPathForSafe returns the path to the safe data directory.
// The path is determined by the VSECM_SAFE_DATA_PATH environment variable.
// If the environment variable is not set, the default path "/var/local/vsecm/data"
// is returned.
func DataPathForSafe() string {
	p := env.Value(env.VSecMSafeDataPath)
	if p == "" {
		p = string(env.VSecMSafeDataPathDefault)
	}
	return p
}

// RootKeyPathForSafe returns the path to the safe age key directory.
// The path is determined by the VSECM_ROOT_KEY_PATH environment variable.
// If the environment variable is not set, the default path "/key/key.txt"
// is returned.
func RootKeyPathForSafe() string {
	p := env.Value(env.VSecMRootKeyPath)
	if p == "" {
		p = string(env.VSecMRootKeyPathDefault)
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
	p := env.Value(env.VSecMSafeSourceAcquisitionTimeout)
	d, _ := strconv.Atoi(string(env.VSecMSafeSourceAcquisitionTimeoutDefault))
	if p == "" {
		p = string(env.VSecMSafeSourceAcquisitionTimeoutDefault)
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		i = int64(d)
		return time.Duration(i) * time.Millisecond
	}

	return time.Duration(i) * time.Millisecond
}

// BootstrapTimeoutForSafe returns the allowed time for VSecM Safe to wait
// before killing the pod to retrieve an SVID, in time.Duration.
// The interval is determined by the VSECM_SAFE_BOOTSTRAP_TIMEOUT environment
// variable, with a default value of 300000 milliseconds if the variable is not
// set or if there is an error in parsing the value.
func BootstrapTimeoutForSafe() time.Duration {
	p := env.Value(env.VSecMSafeBootstrapTimeout)
	d, _ := strconv.Atoi(string(env.VSecMSafeBootstrapTimeoutDefault))
	if p == "" {
		p = string(env.VSecMSafeBootstrapTimeoutDefault)
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		i = int64(d)
		return time.Duration(i) * time.Millisecond
	}

	return time.Duration(i) * time.Millisecond
}

// RootKeySecretNameForSafe returns the name of the environment variable that
// holds the VSecM Safe age key. The value is retrieved using the
// "VSECM_ROOT_KEY_NAME" environment variable. If this variable is
// not set or is empty, the default value "vsecm-root-key" is returned.
func RootKeySecretNameForSafe() string {
	p := env.Value(env.VSecMRootKeyName)
	if p == "" {
		p = string(env.VSecMRootKeyNameDefault)
	}
	return p
}
