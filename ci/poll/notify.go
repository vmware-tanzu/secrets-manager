package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func notifyBuildFailure() {
	fmt.Println("Build failed, notifying")

	// Read the API key from a file
	apiKeyBytes, err := os.ReadFile("/opt/vsecm/ifttt.key") // apikey.txt contains your API key
	if err != nil {
		fmt.Println("Error reading API key file:", err)
		return
	}
	apiKey := strings.TrimSpace(string(apiKeyBytes))

	urlTemplate := "https://maker.ifttt.com/trigger/{event}/json/with/key/%s"
	event := "vsecm-build-failed"

	url := fmt.Sprintf(urlTemplate, apiKey)
	url = strings.Replace(url, "{event}", event, 1)

	payload := map[string]interface{}{}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing response body:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
