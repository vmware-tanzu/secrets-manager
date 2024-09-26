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
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type TrustDomainInfo struct {
	Name              string `json:"name"`
	BundleEndpointURL string `json:"bundleEndpointUrl"`
	TrustDomain       string `json:"trustDomain"`
	EndpointSPIFFEID  string `json:"endpointSPIFFEID"`
}

type Config struct {
	Diablo       TrustDomainInfo `json:"diablo"`
	Mephisto     TrustDomainInfo `json:"mephisto"`
	Baal         TrustDomainInfo `json:"baal"`
	Azmodan      TrustDomainInfo `json:"azmodan"`
	FederateWith []string        `json:"federateWith"`
}

const crTemplate = `apiVersion: spire.spiffe.io/v1alpha1
kind: ClusterFederatedTrustDomain
metadata:
  name: {{ .Name }}
spec:
  trustDomain: {{ .TrustDomain }}
  bundleEndpointURL: {{ .BundleEndpointURL }}
  bundleEndpointProfile:
    type: https_spiffe
    endpointSPIFFEID: {{ .EndpointSPIFFEID }}
  trustDomainBundle: |-
{{ .TrustDomainBundle | indent 4 }}
`

func main() {
	// Read and parse the JSON configuration file
	configFile, err := os.Open("endpoints.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer configFile.Close()

	var config Config
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Error parsing config file:", err)
		return
	}

	// Create ClusterFederatedTrustDomain CRs for each federation target
	for _, target := range config.FederateWith {
		var info TrustDomainInfo
		switch target {
		case "diablo":
			info = config.Diablo
		case "mephisto":
			info = config.Mephisto
		case "baal":
			info = config.Baal
		case "azmodan":
			info = config.Azmodan
		default:
			fmt.Printf("Unknown federation target: %s\n", target)
			continue
		}

		// Fetch trust domain bundle
		bundle, err := fetchTrustDomainBundle(info.BundleEndpointURL)
		if err != nil {
			fmt.Printf("Error fetching trust domain bundle for %s: %v\n", target, err)
			continue
		}

		// Create CR
		cr, err := createCR(target, info, bundle)
		if err != nil {
			fmt.Printf("Error creating CR for %s: %v\n", target, err)
			continue
		}

		// Apply CR
		err = applyCR(cr)
		if err != nil {
			fmt.Printf("Error applying CR for %s: %v\n", target, err)
			continue
		}

		fmt.Printf("Successfully created and applied CR for %s\n", target)
	}
}

func fetchTrustDomainBundle(url string) (string, error) {
	// Create a new HTTP client with InsecureSkipVerify set to true
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// Make the GET request
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func createCR(name string, info TrustDomainInfo, bundle string) (string, error) {
	tmpl, err := template.New("cr").Funcs(template.FuncMap{
		"indent": func(spaces int, v string) string {
			pad := strings.Repeat(" ", spaces)
			return pad + strings.Replace(v, "\n", "\n"+pad, -1)
		},
	}).Parse(crTemplate)
	if err != nil {
		return "", err
	}

	// Pretty print the JSON bundle
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(bundle), "", "  ")
	if err != nil {
		return "", fmt.Errorf("error formatting JSON: %v", err)
	}

	data := struct {
		Name              string
		TrustDomain       string
		BundleEndpointURL string
		EndpointSPIFFEID  string
		TrustDomainBundle string
	}{
		Name:              name,
		TrustDomain:       info.TrustDomain,
		BundleEndpointURL: info.BundleEndpointURL,
		EndpointSPIFFEID:  info.EndpointSPIFFEID,
		TrustDomainBundle: prettyJSON.String(),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func applyCR(cr string) error {
	// Create a temporary file to store the CR
	tmpfile, err := ioutil.TempFile("", "cr-*.yaml")
	if err != nil {
		return fmt.Errorf("error creating temporary file: %v", err)
	}
	// defer os.Remove(tmpfile.Name())

	// Write the CR to the temporary file
	if _, err := tmpfile.Write([]byte(cr)); err != nil {
		return fmt.Errorf("error writing to temporary file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		return fmt.Errorf("error closing temporary file: %v", err)
	}

	// Apply the CR using microk8s kubectl
	cmd := exec.Command("microk8s", "kubectl", "apply", "-f", tmpfile.Name())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error applying CR: %v, output: %s", err, string(output))
	}

	fmt.Printf("Successfully applied CR. Output: %s\n", string(output))
	return nil
}
