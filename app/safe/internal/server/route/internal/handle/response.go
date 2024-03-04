/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package handle

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/core/audit"
	event "github.com/vmware-tanzu/secrets-manager/core/audit/state"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// BadSvidResponse logs an event for a bad SPIFFE ID and sends an HTTP 400 Bad
// Request response. This function is typically invoked when the SPIFFE ID provided
// in a request is invalid or malformed.
//
// Parameters:
// - cid (string): Correlation ID for operation tracing and logging.
// - w (http.ResponseWriter): The HTTP response writer to send back the response.
// - spiffeid (string): The SPIFFE ID that was determined to be invalid.
// - j (audit.JournalEntry): An audit journal entry for recording the event.
func BadSvidResponse(cid string, w http.ResponseWriter, spiffeid string,
	j audit.JournalEntry,
) {
	j.Event = event.BadSpiffeId
	audit.Log(j)

	log.DebugLn(&cid, "Fetch: bad spiffeid", spiffeid)

	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "")
	if err != nil {
		log.InfoLn(&cid, "Fetch: Problem sending response", err.Error())
	}
}

// BadPeerSvidResponse logs an event for a bad peer SPIFFE ID and sends an
// HTTP 400 Bad Request response. This function is used when the peer SPIFFE ID
// in a mutual TLS session is found to be invalid or unacceptable.
//
// Parameters:
// - cid (string): Correlation ID for operation tracing and logging.
// - w (http.ResponseWriter): The HTTP response writer to send back the response.
// - spiffeid (string): The peer's SPIFFE ID that was found to be invalid.
// - j (audit.JournalEntry): An audit journal entry for recording the event.
func BadPeerSvidResponse(cid string, w http.ResponseWriter,
	spiffeid string, j audit.JournalEntry,
) {
	j.Event = event.BadPeerSvid
	audit.Log(j)

	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "")
	if err != nil {
		log.InfoLn(&cid, "Fetch: Problem with spiffeid", spiffeid)
	}
}

// NoSecretResponse logs an event indicating that no secret was found and sends
// an HTTP 404 Not Found response. This function is invoked when a request for
// a secret results in no matching secret being available.
//
// Parameters:
// - cid (string): Correlation ID for operation tracing and logging.
// - w (http.ResponseWriter): The HTTP response writer to send back the response.
// - j (audit.JournalEntry): An audit journal entry for recording the event.
func NoSecretResponse(cid string, w http.ResponseWriter,
	j audit.JournalEntry,
) {
	j.Event = event.NoSecret
	audit.Log(j)

	w.WriteHeader(http.StatusNotFound)
	_, err2 := io.WriteString(w, "")
	if err2 != nil {
		log.InfoLn(&cid, "Fetch: Problem sending response", err2.Error())
	}
}

// InitCompleteSuccessResponse logs a successful initialization event and sends
// the initialization completion response back to the client. It marshals and
// sends a structured response indicating successful initialization.
//
// Parameters:
//   - cid (string): Correlation ID for operation tracing and logging.
//   - w (http.ResponseWriter): The HTTP response writer to send back the response.
//   - j (audit.JournalEntry): An audit journal entry for recording the event.
//   - sfr (reqres.SentinelInitCompleteResponse): The response payload to be
//     marshaled and sent.
func InitCompleteSuccessResponse(cid string, w http.ResponseWriter,
	j audit.JournalEntry, sfr reqres.SentinelInitCompleteResponse) {
	j.Event = event.Ok
	j.Entity = sfr
	audit.Log(j)

	resp, err := json.Marshal(sfr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err2 := io.WriteString(w, "Problem unmarshalling response")
		if err2 != nil {
			log.InfoLn(&cid, "Problem sending response", err2.Error())
		}
		return
	}

	log.DebugLn(&cid, "before response")

	_, err = io.WriteString(w, string(resp))
	if err != nil {
		log.InfoLn(&cid, "Problem sending response", err.Error())
	}

	log.DebugLn(&cid, "after response")
}

// SuccessResponse logs a successful operation event and sends a structured
// success response back to the client. It marshals and sends a secret fetch
// response, indicating the successful retrieval of a secret.
//
// Parameters:
//   - cid (string): Correlation ID for operation tracing and logging.
//   - w (http.ResponseWriter): The HTTP response writer to send back the response.
//   - j (audit.JournalEntry): An audit journal entry for recording the event.
//   - sfr (reqres.SecretFetchResponse): The secret fetch response payload to be
//     marshaled and sent.
func SuccessResponse(cid string, w http.ResponseWriter,
	j audit.JournalEntry, sfr reqres.SecretFetchResponse) {
	j.Event = event.Ok
	j.Entity = sfr
	audit.Log(j)

	resp, err := json.Marshal(sfr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err2 := io.WriteString(w, "Problem unmarshalling response")
		if err2 != nil {
			log.InfoLn(&cid, "Fetch: Problem sending response", err2.Error())
		}
		return
	}

	log.DebugLn(&cid, "Fetch: before response")

	_, err = io.WriteString(w, string(resp))
	if err != nil {
		log.InfoLn(&cid, "Problem sending response", err.Error())
	}

	log.DebugLn(&cid, "Fetch: after response")
}
