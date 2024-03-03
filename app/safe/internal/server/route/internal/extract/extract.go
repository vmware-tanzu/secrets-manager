/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package extract

import (
	"encoding/json"
	"strings"

	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// WorkloadIDAndParts extracts the workload identifier and its constituent parts
// from a SPIFFE ID string, based on a predefined prefix that is removed from
// the SPIFFE ID.
//
// Parameters:
//   - spiffeid (string): The SPIFFE ID string from which the workload identifier
//     and parts are to be extracted.
//
// Returns:
//   - (string, []string): The first return value is the workload identifier, which
//     is essentially the first part of the SPIFFE ID after removing the prefix. The
//     second return value is a slice of strings representing all parts of the
//     SPIFFE ID after the prefix removal.
func WorkloadIDAndParts(spiffeid string) (string, []string) {
	tmp := strings.Replace(spiffeid, env.SpiffeIdPrefixForWorkload(), "", 1)
	parts := strings.Split(tmp, "/")
	if len(parts) > 0 {
		return parts[0], parts
	}
	return "", nil
}

// SecretValue determines the appropriate representation of a secret's value to use,
// giving precedence to a transformed value over the raw values. This function
// supports both current implementations and backward compatibility by handling
// secrets with multiple values or transformed values.
//
// Parameters:
//   - cid (string): Correlation ID for operation tracing and logging.
//   - secret (*entity.SecretStored): A pointer to the secret entity from which
//     the value is to be retrieved.
//
// Returns:
//   - string: The selected representation of the secret's value. This could be
//     the transformed value, a single raw value, a JSON-encoded string of multiple
//     values, or an empty string in case of an error.
func SecretValue(cid string, secret *entity.SecretStored) string {
	if secret.ValueTransformed != "" {
		log.TraceLn(&cid, "Fetch: using transformed value")
		return secret.ValueTransformed
	}

	// This part is for backwards compatibility.
	// It probably won't execute because `secret.ValueTransformed` will
	// always be set.

	log.TraceLn(&cid, "Fetch: using raw value")

	if len(secret.Values) == 1 {
		return secret.Values[0]
	}

	jsonData, err := json.Marshal(secret.Values)
	if err != nil {
		log.WarnLn(&cid, "Fetch: Problem marshaling values", err.Error())
	} else {
		return string(jsonData)
	}

	return ""
}
