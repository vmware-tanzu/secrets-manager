/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package http_test

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	http_vmware "github.com/vmware-tanzu/secrets-manager/app/safe/internal/server/route/base/http" // Replace with the actual path to the `http` package
)

func TestReadBody_Success(t *testing.T) {
	// Prepare the test data
	cid := "test-cid"
	expectedBody := []byte("test body content")
	r := &http.Request{
		Body: ioutil.NopCloser(bytes.NewBuffer(expectedBody)),
	}

	// Call the function
	body, err := http_vmware.ReadBody(cid, r)

	// Assertions
	require.NoError(t, err, "Expected no error reading the body")
	assert.Equal(t, expectedBody, body, "Expected the body to match the input")
}

func TestReadBody_ErrorReadingBody(t *testing.T) {
	// Prepare the test data
	cid := "test-cid"
	r := &http.Request{
		Body: ioutil.NopCloser(&errorReader{}),
	}

	// Call the function
	body, err := http_vmware.ReadBody(cid, r)

	// Assertions
	assert.Nil(t, body, "Expected body to be nil on error")
	assert.Error(t, err, "Expected an error reading the body")
}

func TestReadBody_ErrorClosingBody(t *testing.T) {
	// Prepare the test data
	cid := "test-cid"
	expectedBody := []byte("test body content")
	r := &http.Request{
		Body: &closerWithError{
			ReadCloser: ioutil.NopCloser(bytes.NewBuffer(expectedBody)),
		},
	}

	// Call the function
	body, err := http_vmware.ReadBody(cid, r)

	// Assertions
	require.NoError(t, err, "Expected no error reading the body")
	assert.Equal(t, expectedBody, body, "Expected the body to match the input")

	// If you want to test log output, you need to set up log capturing and assertions
	// Example (assuming a log package that allows this):
	// var logOutput bytes.Buffer
	// log.SetOutput(&logOutput)
	// assert.Contains(t, logOutput.String(), "ReadBody: Problem closing body", "Expected log to contain close error message")
}

// Helper types for testing error cases
type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("read error")
}

type closerWithError struct {
	io.ReadCloser
}

func (c *closerWithError) Close() error {
	return errors.New("close error")
}
