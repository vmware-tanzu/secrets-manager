/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package http

import (
	"io"
	"net/http"

	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// ReadBody reads the body from an HTTP request and returns it as a byte slice.
// If there is an error reading the body, it returns the error.
//
// The function takes the following parameters:
//   - cid: A string representing the correlation ID for logging purposes.
//   - r: A pointer to an http.Request from which the body is to be read.
//
// It returns:
//   - A byte slice containing the body of the request.
//   - An error if there was an issue reading the body or closing the request 
//   body.
//
// Example usage:
//
//	body, err := ReadBody(cid, request)
//	if err != nil {
//	  log.Error("Failed to read body:", err)
//	  // Handle error
//	}
//	// Use body
func ReadBody(cid string, r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.InfoLn(&cid, "ReadBody: Problem closing body", err.Error())
		}
	}(r.Body)

	return body, nil
}
