package engine

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/rpc"
)

// TokenIntrospectionResponse represents the JSON structure received from an
// OAuth 2.0 Token Introspection endpoint.
// It contains the active state of the token which determines if it is valid
// or expired.
type TokenIntrospectionResponse struct {
	Active bool `json:"active"`
}

// AuthorizedJWT determines if the provided JWT (access token) is authorized
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
//	authorized := AuthorizedJWT("cid123", req)
//	if authorized {
//	    // proceed with authorized logic
//	}
func AuthorizedJWT(cid string, r *http.Request) bool {
	client := &http.Client{}
	data := url.Values{}
	accessToken := r.Header.Get("Authorization")
	clientId := r.Header.Get("ClientId")
	clientSecret := r.Header.Get("ClientSecret")
	username := r.Header.Get("UserName")

	data.Set("client_id", strings.TrimSpace(clientId))
	data.Set("client_secret", strings.TrimSpace(clientSecret))
	data.Set("token", strings.TrimSpace(accessToken))
	data.Set("username", strings.TrimSpace(username))

	if accessToken == "" && clientId == "" && clientSecret == "" &&
		username == "" {
		log.ErrorLn(&cid,
			"AuthorizedJWT please check your sending request headers!")
		return false
	}

	req, err := http.NewRequest("POST",
		env.OIDCProviderBaseUrlForSentinel(), strings.NewReader(data.Encode()))
	if err != nil {
		log.ErrorLn(&cid,
			"AuthorizedJWT an error occurred when creating request:", err)
		return false
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.ErrorLn(&cid,
			"AuthorizedJWT an error occurred when sending request:", err)
		return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.ErrorLn(&cid, "Error closing response body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.ErrorLn(&cid,
			"AuthorizedJWT an error occurred when reading response body:",
			err)
		return false
	}

	var tokenResponse TokenIntrospectionResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		log.ErrorLn(&cid,
			"AuthorizedJWT an error occurred when unmarshalling response:",
			err)
		return false
	}

	log.InfoLn(&cid, "AuthorizedJWT token is active:", tokenResponse.Active)
	return tokenResponse.Active
}
