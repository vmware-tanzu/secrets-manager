/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package probe

import (
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestCreateLiveness(t *testing.T) {
	os.Setenv("VSECM_PROBE_LIVENESS_PORT", ":8095")
	defer os.Unsetenv("VSECM_PROBE_LIVENESS_PORT")

	t.Logf("Creating liveness probe...")
	go CreateLiveness()

	t.Logf("Sending request to mock server...")
	resp, err := http.Get("http://localhost" + env.ProbeLivenessPort())
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	t.Logf("Checking response status code...")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}

	t.Logf("Checking response body...")
	expected := "OK"
	actual, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	actualStr := string(actual)
	if actualStr != expected {
		t.Errorf("Unexpected response body: got %v, want %v", actual, expected)
	}
}

func TestCreateReadiness(t *testing.T) {
	os.Setenv("VSECM_PROBE_READINESS_PORT", ":8096")
	defer os.Unsetenv("VSECM_PROBE_READINESS_PORT")

	t.Logf("Creating readiness probe...")
	go CreateReadiness()

	t.Logf("Sending request to mock server...")
	resp, err := http.Get("http://localhost" + env.ProbeReadinessPort())
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	t.Logf("Checking response status code...")
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}

	t.Logf("Checking response body...")
	expected := "OK"
	actual, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}

	actualStr := string(actual)
	if actualStr != expected {
		t.Errorf("Unexpected response body: got %v, want %v", actual, expected)
	}
}
