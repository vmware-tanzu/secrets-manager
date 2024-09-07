package engine

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
)

// MockHTTPClient is a mock implementation of the HTTPClient interface
type MockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func TestNewAuth(t *testing.T) {
	auth := newAuth()
	if auth == nil {
		t.Error("NewAuth() returned nil")
	}
}

func TestWithHTTPClient(t *testing.T) {
	mockClient := &MockHTTPClient{}
	auth := newAuth(withHTTPClient(mockClient)).(*auth)
	if auth.httpClient != mockClient {
		t.Error("WithHTTPClient() did not set the HTTP client correctly")
	}
}

func TestIsAuthorized(t *testing.T) {
	tests := []struct {
		name           string
		headers        map[string]string
		mockResponse   string
		mockError      error
		expectedResult bool
	}{
		{
			name: "Valid token",
			headers: map[string]string{
				"ClientId":      "client1",
				"ClientSecret":  "secret1",
				"Authorization": "Bearer token1",
				"UserName":      "user1",
			},
			mockResponse:   `{"active": true}`,
			expectedResult: true,
		},
		{
			name: "Invalid token",
			headers: map[string]string{
				"ClientId":      "client2",
				"ClientSecret":  "secret2",
				"Authorization": "Bearer token2",
				"UserName":      "user2",
			},
			mockResponse:   `{"active": false}`,
			expectedResult: false,
		},
		{
			name: "Missing header",
			headers: map[string]string{
				"ClientId":      "client3",
				"ClientSecret":  "secret3",
				"Authorization": "Bearer token3",
			},
			mockResponse:   `{"active": true}`,
			expectedResult: false,
		},
		{
			name: "HTTP client error",
			headers: map[string]string{
				"ClientId":      "client4",
				"ClientSecret":  "secret4",
				"Authorization": "Bearer token4",
				"UserName":      "user4",
			},
			mockError:      io.ErrUnexpectedEOF,
			expectedResult: false,
		},
		{
			name: "Invalid JSON response",
			headers: map[string]string{
				"ClientId":      "client5",
				"ClientSecret":  "secret5",
				"Authorization": "Bearer token5",
				"UserName":      "user5",
			},
			mockResponse:   `{"active": invalid}`,
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &MockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					if tt.mockError != nil {
						return nil, tt.mockError
					}
					return &http.Response{
						StatusCode: 200,
						Body:       io.NopCloser(bytes.NewBufferString(tt.mockResponse)),
					}, nil
				},
			}

			mockLogger := &MockLogger{}
			mockLogger.On("InfoLn", mock.Anything, mock.Anything).Return()
			mockLogger.On("ErrorLn", mock.Anything, mock.Anything).Return()

			auth := newAuth(
				withHTTPClient(mockClient),
				withLogger(mockLogger),
			)

			req, _ := http.NewRequest("GET", "http://example.com", nil)
			for k, v := range tt.headers {
				req.Header.Set(k, v)
			}

			result := auth.IsAuthorized("test-cid", req)
			if result != tt.expectedResult {
				t.Errorf("IsAuthorized() = %v, want %v", result, tt.expectedResult)
			}
		})
	}
}
