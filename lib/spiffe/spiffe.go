package spiffe

import (
	"errors"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/svid/x509svid"
	"net/http"
)

// IdFromRequest extracts the SPIFFE ID from the TLS peer certificate of
// an HTTP request.
// It checks if the incoming request has a valid TLS connection and at least one
// peer certificate.
// The first certificate in the chain is used to extract the SPIFFE ID.
//
// Params:
//
//	r *http.Request - The HTTP request from which the SPIFFE ID is to be
//	extracted.
//
// Returns:
//
//	 *spiffeid.ID - The SPIFFE ID extracted from the first peer certificate,
//	 or nil if extraction fails.
//	 error - An error object indicating the failure reason. Possible errors
//	include the absence of peer certificates or a failure in extracting the
//	SPIFFE ID from the certificate.
//
// Note:
//
//	This function assumes that the request is already over a secured TLS
//	connection and will fail if the TLS connection state is not available or
//	the peer certificates are missing.
func IdFromRequest(r *http.Request) (*spiffeid.ID, error) {
	tlsConnectionState := r.TLS
	if len(tlsConnectionState.PeerCertificates) == 0 {
		return nil, errors.New("no peer certs")
	}

	id, err := x509svid.IDFromCert(tlsConnectionState.PeerCertificates[0])
	if err != nil {
		return nil, errors.Join(
			err,
			errors.New("problem extracting svid"),
		)
	}

	return &id, nil
}

// IdAsString retrieves the SPIFFE ID from an HTTP request and returns it as
// a string.
//
// Parameters:
// - cid: A string representing the context identifier.
// - r: A pointer to an http.Request from which the SPIFFE ID will be extracted.
//
// Returns:
//   - A string representing the SPIFFE ID if it can be successfully retrieved;
//     otherwise, an empty string.
//
// If the SPIFFE ID cannot be retrieved from the request, it logs an
// informational message and returns an empty string.
//
// Example usage:
//
//	func handler(w http.ResponseWriter, r *http.Request) {
//	    cid := "exampleContextID"
//	    spiffeID := IdAsString(cid, r)
//	    fmt.Fprintf(w, "SPIFFE ID: %s", spiffeID)
//	}
func IdAsString(r *http.Request) string {
	sr, err := IdFromRequest(r)
	if err != nil {
		return ""
	}
	return sr.String()
}
