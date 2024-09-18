package engine

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/vmware-tanzu/secrets-manager/core/env"
)

// httpClient is a minimal interface for making HTTP requests.
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Auth is the default implementation of the Authorizer interface.
type auth struct {
	httpClient httpClient
	log        logger
}

// Ensure Auth implements Authorizer
var _ authorizer = (*auth)(nil)

// newAuth creates a new Auth instance with the provided options.
func newAuth(opts ...authOption) authorizer {
	a := &auth{httpClient: &http.Client{}}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// authOption is a functional option type for configuring Auth.
type authOption func(*auth)

// withHTTPClient sets a custom HTTP client for Auth.
func withHTTPClient(client httpClient) authOption {
	return func(a *auth) {
		if client != nil {
			a.httpClient = client
		}
	}
}

// withLogger sets a custom logger for Auth.
func withLogger(logger logger) authOption {
	return func(a *auth) {
		if logger != nil {
			a.log = logger
		}
	}
}

// IsAuthorized checks if the JWT (access token) is authorized based on the
// OAuth 2.0 Token Introspection standard.
func (a *auth) IsAuthorized(id string, r *http.Request) bool {
	return a.isAuthorizedJWT(id, r)
}

// isAuthorizedJWT determines if the provided JWT (access token) is authorized
// based on the OAuth 2.0 Token Introspection standard. It makes an HTTP POST
// request to an introspection endpoint with the required credentials and token,
// expecting a TokenIntrospectionResponse.
//
// Parameters:
//   - cid: a client identifier used for logging purposes.
//   - r: the http.Request containing necessary headers such as Authorization
//     (Bearer token), ClientId, ClientSecret, and UserName.
//
// Returns true if the token is active and authorized, false otherwise.
// This function logs detailed error messages using a custom logger and handles
// the request lifecycle, including closing response bodies and error handling.
//
// Usage:
//
// a := NewAuth()
//
//	if a.IsAuthorized("client_id", r) {
//	    // Proceed with the request
//	} else {
//
//	    // Handle unauthorized access
//	}
func (a *auth) isAuthorizedJWT(cid string, r *http.Request) bool {
	// headers is a list of key-value pairs that are required to be sent to the
	// introspection endpoint.
	headers := []struct {
		key, value string
	}{
		{"client_id", "ClientId"},
		{"client_secret", "ClientSecret"},
		{"token", "Authorization"},
		{"username", "UserName"},
	}

	// data is a list of key-value pairs that are required to be sent to the
	// introspection endpoint.
	data := url.Values{}
	for _, h := range headers {
		if value := strings.TrimSpace(r.Header.Get(h.value)); value != "" {
			data.Set(h.key, value)
		}
	}

	// If the number of headers is not equal to the number of data, then some
	// required headers are missing.
	if len(data) != len(headers) {
		a.log.ErrorLn(&cid, "isAuthorizedJWT: missing required headers")
		return false
	}

	// Create a new HTTP request to the introspection endpoint with the required
	// data.
	req, err := http.NewRequest("POST",
		env.OIDCProviderBaseUrlForSentinel(), strings.NewReader(data.Encode()))
	if err != nil {
		a.log.ErrorLn(&cid, "isAuthorizedJWT: error creating request:", err)
		return false
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		a.log.ErrorLn(&cid, "isAuthorizedJWT: error sending request:", err)
		return false
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			a.log.ErrorLn(&cid,
				"isAuthorizedJWT: error closing response body:", err)
		}
	}()

	// TokenIntrospectionResponse represents the JSON structure received from an
	// OAuth 2.0 Token Introspection endpoint.
	// It contains the active state of the token which determines if it is valid
	// or expired.
	var tokenResponse struct {
		Active bool `json:"active"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		a.log.ErrorLn(&cid, "isAuthorizedJWT: error decoding response:", err)
		return false
	}

	a.log.InfoLn(&cid,
		"isAuthorizedJWT: token is active:", tokenResponse.Active)
	return tokenResponse.Active
}
