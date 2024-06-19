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
	"fmt"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestCreateLiveness(t *testing.T) {
	_ = os.Setenv("VSECM_PROBE_LIVENESS_PORT", ":8095")
	defer func() {
		err := os.Unsetenv("VSECM_PROBE_LIVENESS_PORT")
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	t.Logf("Creating liveness probe...")
	<-CreateLiveness()

	t.Logf("Sending request to mock server...")
	resp, err := http.Get("http://localhost" + env.ProbeLivenessPort())
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		t.Logf("Closing response body...")
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(resp.Body)

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
	_ = os.Setenv("VSECM_PROBE_READINESS_PORT", ":8096")
	defer func() {
		err := os.Unsetenv("VSECM_PROBE_READINESS_PORT")
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	t.Logf("Creating readiness probe...")
	<-CreateReadiness()

	t.Logf("Sending request to mock server...")
	resp, err := http.Get("http://localhost" + env.ProbeReadinessPort())
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		t.Logf("Closing response body...")
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(resp.Body)

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
