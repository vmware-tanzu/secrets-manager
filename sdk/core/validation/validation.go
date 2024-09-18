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
	"regexp"
	"strings"

	e "github.com/vmware-tanzu/secrets-manager/sdk/core/constants/env"
	"github.com/vmware-tanzu/secrets-manager/sdk/core/env"
)

// Any SPIFFE ID regular expression matcher shall start with the
// `^spiffe://$trustDomain` prefix for extra security.
//
// This variable shall be treated as constant and should not be modified.
var spiffeRegexPrefixStart = "^spiffe://" + env.SpiffeTrustDomain() + "/"
var spiffeIdPrefixStart = "spiffe://" + env.SpiffeTrustDomain() + "/"

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
			panic(
				"Failed to compile the regular expression pattern " +
					"for SPIFFE ID." +
					" Check the " + string(e.VSecMSpiffeIdPrefixWorkload) +
					" environment variable. " +
					" val: " + env.SpiffeIdPrefixForWorkload() +
					" trust: " + env.SpiffeTrustDomain(),
			)
			return false
		}

		nrw := env.NameRegExpForWorkload()
		wre, err := regexp.Compile(nrw)
		if err != nil {
			panic(
				"Failed to compile the regular expression pattern " +
					"for SPIFFE ID." +
					" Check the " + string(e.VSecMWorkloadNameRegExp) +
					" environment variable." +
					" val: " + env.NameRegExpForWorkload() +
					" trust: " + env.SpiffeTrustDomain(),
			)
			return false
		}

		match := wre.FindStringSubmatch(spiffeid)
		if len(match) == 0 {
			return false
		}

		return re.MatchString(spiffeid)
	}

	if !strings.HasPrefix(spiffeid, spiffeIdPrefixStart) {
		return false
	}

	nrw := env.NameRegExpForWorkload()
	if !strings.HasPrefix(nrw, spiffeRegexPrefixStart) {

		// Insecure configuration detected.
		// Panic to prevent further issues:
		panic(
			"Invalid regular expression pattern for SPIFFE ID." +
				" Expected: ^spiffe://<trust_domain>/..." +
				" Check the " + string(e.VSecMWorkloadNameRegExp) +
				" environment variable." +
				" val: " + env.NameRegExpForWorkload() +
				" trust: " + env.SpiffeTrustDomain(),
		)
		return false
	}

	wre, err := regexp.Compile(nrw)
	if err != nil {
		panic(
			"Failed to compile the regular expression pattern " +
				"for SPIFFE ID." +
				" Check the " + string(e.VSecMWorkloadNameRegExp) +
				" environment variable." +
				" val: " + env.NameRegExpForWorkload() +
				" trust: " + env.SpiffeTrustDomain(),
		)
		return false
	}

	match := wre.FindStringSubmatch(spiffeid)
	if len(match) == 0 {
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
			panic(
				"Failed to compile the regular expression pattern " +
					"for Sentinel SPIFFE ID." +
					" Check the " + string(e.VSecMSpiffeIdPrefixSafe) +
					" environment variable." +
					" val: " + env.SpiffeIdPrefixForSafe() +
					" trust: " + env.SpiffeTrustDomain(),
			)
		}

		return re.MatchString(spiffeid)
	}

	return strings.HasPrefix(spiffeid, prefix)
}
