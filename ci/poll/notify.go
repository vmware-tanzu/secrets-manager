package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const apiKeyPath = "/opt/vsecm/ifttt.key"

func notifyBuildFailure() {
	log.Println("Build failed, notifying interested parties...")

	// Read the API key from a file
	apiKeyBytes, err := os.ReadFile(apiKeyPath)
	if err != nil {
		fmt.Println("Error reading API key file:", err)
		return
	}
	apiKey := strings.TrimSpace(string(apiKeyBytes))

	urlTemplate := "https://maker.ifttt.com/trigger/{event}/json/with/key/%s"
	event := "vsecm-build-failed"

	url := fmt.Sprintf(urlTemplate, apiKey)
	url = strings.Replace(url, "{state}", event, 1)

	payload := map[string]interface{}{}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	log.Println("Response:", string(body))
}
