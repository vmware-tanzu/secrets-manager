/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package list

import (
	"net/http"
)

// Masked returns all registered workloads to the system with some metadata
// that is secure to share. For example, it returns secret names but not values.
//
// - cid: A string representing the client identifier.
// - w: An http.ResponseWriter used to write the HTTP response.
// - r: A pointer to an http.Request representing the received HTTP request.
// - spiffeid: spiffe id of the caller.
func Masked(
	cid string, r *http.Request, w http.ResponseWriter,
) {
	doList(cid, w, r, false)
}

// Encrypted returns all registered workloads to the system. Similar to `Masked`
// it return meta information; however, it also returns encrypted secret values
// where an operator can decrypt if they know the VSecM root key.
//
// - cid: A string representing the client identifier.
// - w: An http.ResponseWriter used to write the HTTP response.
// - r: A pointer to an http.Request representing the received HTTP request.
// - spiffeid: spiffe id of the caller.
func Encrypted(
	cid string, r *http.Request, w http.ResponseWriter,
) {
	doList(cid, w, r, true)
}
