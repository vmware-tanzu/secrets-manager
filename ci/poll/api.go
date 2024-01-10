/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	githubAPIURL = "https://api.github.com/repos/vmware-tanzu/secrets-manager/git/refs/heads/main"
	pollInterval = 15 * time.Minute
)

type GitHubResponse struct {
	Object struct {
		SHA string `json:"sha"`
	} `json:"object"`
}

func getLatestCommitHash() (string, error) {
	resp, err := http.Get(githubAPIURL)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %s", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var gitHubResponse GitHubResponse
	err = json.Unmarshal(body, &gitHubResponse)
	if err != nil {
		return "", err
	}

	return gitHubResponse.Object.SHA, nil
}
