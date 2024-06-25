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
	"regexp"
	"strings"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/core/env"
)

// Any SPIFFE ID regular expression matcher shall start with the
// `^spiffe://$trustDomain` prefix for extra security.
//
// This variable shall be treated as constant and should not be modified.
var spiffeRegexPrefixStart = "^spiffe://" + env.SpiffeTrustDomain() + "/"

// IsSentinel checks if a given SPIFFE ID belongs to VSecM Sentinel.
//
// A SPIFFE ID (SPIFFE IDentifier) is a URI that uniquely identifies a workload
// in a secure, interoperable way. This function verifies if the provided
// SPIFFE ID meets the criteria to be classified as a workload ID based on
// certain environmental settings.
//
// The function performs the following checks:
//  1. If the `spiffeid` starts with a "^", it assumed that it is a regular
//     expression pattern, it compiles the expression and checks if the SPIFFE
//     ID matches it.
//  2. Otherwise, it checks if the SPIFFE ID starts with the proper  prefix.
//
// Parameters:
//
//	spiffeid (string): The SPIFFE ID to be checked.
//
// Returns:
//
//	bool: `true` if the SPIFFE ID belongs to VSecM Sentinel, `false` otherwise.
func IsSentinel(spiffeid string) bool {
	if !IsWorkload(spiffeid) {
		return false
	}

	prefix := env.SpiffeIdPrefixForSentinel()

	if strings.HasPrefix(prefix, spiffeRegexPrefixStart) {
		re, err := regexp.Compile(prefix)
		if err != nil {
			return false
		}

		return re.MatchString(spiffeid)
	}

	if !strings.HasPrefix(
		spiffeid,
		"spiffe://"+env.SpiffeTrustDomain()+"/") {
		return false
	}

	return strings.HasPrefix(spiffeid, prefix)
}

// IsSafe checks if a given SPIFFE ID belongs to VSecM Safe.
//
// A SPIFFE ID (SPIFFE IDentifier) is a URI that uniquely identifies a workload
// in a secure, interoperable way. This function verifies if the provided
// SPIFFE ID meets the criteria to be classified as a workload ID based on
// certain environmental settings.
//
// The function performs the following checks:
//  1. If the `spiffeid` starts with a "^", it assumed that it is a regular
//     expression pattern, it compiles the expression and checks if the SPIFFE
//     ID matches it.
//  2. Otherwise, it checks if the SPIFFE ID starts with the proper prefix.
//
// Parameters:
//
//	spiffeid (string): The SPIFFE ID to be checked.
//
// Returns:
//
//	bool: `true` if the SPIFFE ID belongs to VSecM Safe, `false` otherwise.
func IsSafe(spiffeid string) bool {
	if !IsWorkload(spiffeid) {
		return false
	}

	prefix := env.SpiffeIdPrefixForSafe()

	if strings.HasPrefix(prefix, spiffeRegexPrefixStart) {
		re, err := regexp.Compile(prefix)
		if err != nil {
			return false
		}
		return re.MatchString(spiffeid)
	}

	if !strings.HasPrefix(
		spiffeid,
		"spiffe://"+env.SpiffeTrustDomain()+"/") {
		return false
	}

	return strings.HasPrefix(spiffeid, prefix)
}

// IsWorkload checks if a given SPIFFE ID belongs to a workload.
//
// A SPIFFE ID (SPIFFE IDentifier) is a URI that uniquely identifies a workload
// in a secure, interoperable way. This function verifies if the provided
// SPIFFE ID meets the criteria to be classified as a workload ID based on
// certain environmental settings.
//
// The function performs the following checks:
//  1. If the `spiffeid` starts with a "^", it assumed that it is a regular
//     expression pattern, it compiles the expression and checks if the SPIFFE
//     ID matches it.
//  2. Otherwise, it checks if the SPIFFE ID starts with the proper prefix.
//
// Parameters:
//
//	spiffeid (string): The SPIFFE ID to be checked.
//
// Returns:
//
//	bool: `true` if the SPIFFE ID belongs to a workload, `false` otherwise.
func IsWorkload(spiffeid string) bool {
	prefix := env.SpiffeIdPrefixForWorkload()

	if strings.HasPrefix(prefix, spiffeRegexPrefixStart) {
		re, err := regexp.Compile(prefix)
		if err != nil {
			return false
		}

		return re.MatchString(spiffeid)
	}

	if !strings.HasPrefix(
		spiffeid,
		"spiffe://"+env.SpiffeTrustDomain()+"/") {
		return false
	}

	wre := regexp.MustCompile(env.NameRegExpForWorkload())
	match := wre.FindStringSubmatch(spiffeid)
	if len(match) == 0 {
		return false
	}

	return strings.HasPrefix(spiffeid, prefix)
}

// EnsureSafe checks the safety of the SPIFFE ID from the provided X509Source.
// It retrieves an X.509 SVID (SPIFFE Verifiable Identity Document) from the
// source, and validates the SPIFFE ID against a predefined safety check.
//
// If the X509Source fails to provide an SVID, the function will panic with an
// error message specifying the inability to retrieve the SVID.
//
// Similarly, if the SPIFFE ID from the retrieved SVID does not pass the safety
// check, the function will panic with an error message indicating that the
// SPIFFE ID is not recognized.
//
// Panicking in this function indicates severe issues with identity verification
// that require immediate attention and resolution.
//
// Usage:
//
//	var source *workloadapi.X509Source // Assume source is properly initialized
//	EnsureSafe(source)
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
