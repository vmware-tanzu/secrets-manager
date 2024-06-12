/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	githubAPIURL = "https://api.github.com/repos/vmware-tanzu/secrets-manager/git/refs/heads/main"
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
			log.Printf("Error closing response body: %s", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error fetching latest commit hash: %s", body)
	}

	var gitHubResponse GitHubResponse
	err = json.Unmarshal(body, &gitHubResponse)
	if err != nil {
		return "", err
	}

	return gitHubResponse.Object.SHA, nil
}
