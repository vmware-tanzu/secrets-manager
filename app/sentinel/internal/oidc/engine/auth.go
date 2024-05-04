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

type TokenIntrospectionResponse struct {
	Active bool `json:"active"`
}

func IsAuthorizedJWT(cid string, r *http.Request) bool {
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

	if accessToken == "" && clientId == "" && clientSecret == "" && username == "" {
		log.ErrorLn(&cid, "IsAuthorizedJWT please check your sending request headers!")
		return false
	}

	req, err := http.NewRequest("POST", env.OIDCProviderBaseUrlForSentinel(), strings.NewReader(data.Encode()))
	if err != nil {
		log.ErrorLn(&cid, "IsAuthorizedJWT an error occurred when creating request:", err)
		return false
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.ErrorLn(&cid, "IsAuthorizedJWT an error occurred when sending request:", err)
		return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.ErrorLn(&cid, "IsAuthorizedJWT an error occurred when reading response body:", err)
		return false
	}

	var tokenResponse TokenIntrospectionResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		log.ErrorLn(&cid, "IsAuthorizedJWT an error occurred when unmarshalling response:", err)
		return false
	}

	log.InfoLn(&cid, "IsAuthorizedJWT token is active:", tokenResponse.Active)
	return tokenResponse.Active
}
