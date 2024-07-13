/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package journal

import (
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/core/constants/audit"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

// Log prints an audit log entry to the standard output. The log entry includes
// the correlation ID, event type, request method, request URI, SPIFFE ID, and
// payload. The function is intended for use in logging and auditing
// request-related events, providing a standardized format for capturing and
// recording essential request handling information.
func Log(e data.JournalEntry) {
	printAudit(
		e.CorrelationId, e.Event,
		e.Method, e.Url, e.SpiffeId, e.Payload,
	)
}

// CreateDefaultEntry constructs a default audit journal entry for HTTP requests.
// This entry includes basic request information and identifiers, serving as a
// foundational record for auditing and logging purposes. The function
// encapsulates details like the request method, URI, SPIFFE ID, and correlation
// ID into an audit journal entry.
//
// Parameters:
//   - cid (string): The correlation ID associated with the request, used for
//     tracking and correlating logs and audit entries.
//   - spiffeid (string): The SPIFFE ID associated with the requestor or the
//     service, providing a security context.
//   - r (*http.Request): The HTTP request from which method and request URI
//     information are extracted.
//
// Returns:
//   - audit.JournalEntry: An audit journal entry populated with the request's
//     method, URI, correlation ID, SPIFFE ID, and a default event type of
//     'Enter'.
//
// The returned audit.JournalEntry object is intended for use in logging and
// audit trails, providing essential context about the request handling process.
// It serves as a standardized format for capturing request-related information,
// facilitating easier analysis and review of logged events.
func CreateDefaultEntry(
	cid, spiffeid string, r *http.Request,
) data.JournalEntry {
	return data.JournalEntry{
		CorrelationId: cid,
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         audit.Enter,
	}
}
