/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package json

import (
	"encoding/json"

	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/v1/reqres/safe"
)

// UnmarshalSecretUpsertRequest takes a JSON-encoded request body and attempts
// to unmarshal it into a SecretUpsertRequest struct. It handles JSON
// unmarshalling errors by logging, responding with an HTTP error, and returning
// nil. This function is typically used in HTTP server handlers to process
// incoming requests for secret upsert operations.
//
// Parameters:
//   - body ([]byte): The JSON-encoded request body to be unmarshalled.
//
// Returns:
//   - *reqres.SecretUpsertRequest: A pointer to the unmarshalled
//     SecretUpsertRequest struct, or nil if unmarshalling fails.
//   - error: An error if unmarshalling fails.
func UnmarshalSecretUpsertRequest(
	body []byte,
) (*reqres.SecretUpsertRequest, error) {
	var sr reqres.SecretUpsertRequest

	if err := json.Unmarshal(body, &sr); err != nil {
		return nil, err
	}

	return &sr, nil
}

// UnmarshalKeyInputRequest takes a JSON-encoded request body and attempts to
// unmarshal it into a KeyInputRequest struct. Similar to
// UnmarshalSecretUpsertRequest, it deals with JSON unmarshalling errors by
// logging, issuing an HTTP error response, and returning nil. This function is
// utilized within HTTP server handlers to parse incoming requests for key input
// operations.
//
// Parameters:
//   - body ([]byte): The JSON-encoded request body to be unmarshalled.
//
// Returns:
//   - *reqres.KeyInputRequest: A pointer to the unmarshalled KeyInputRequest
//     struct, or nil if unmarshalling fails.
//   - error: An error if unmarshalling fails.
func UnmarshalKeyInputRequest(body []byte) (*reqres.KeyInputRequest, error) {
	var sr reqres.KeyInputRequest

	err := json.Unmarshal(body, &sr)
	if err != nil {
		return nil, err
	}

	return &sr, nil
}
