package oidc

import (
	"encoding/json"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"io"
	"net/http"
	"net/url"
	"strings"
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

	req, err := http.NewRequest("POST", env.SentinelOIDCProviderBaseUrl(), strings.NewReader(data.Encode()))
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
	defer resp.Body.Close()

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
