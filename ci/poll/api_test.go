package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getLatestCommitHash(t *testing.T) {
	tests := []struct {
		name        string
		expectedSha string
		response    *httptest.Server
		want        string
		wantErr     bool
	}{
		{
			name:        "Test getLatestCommitHash Success",
			expectedSha: "d3b07384d113edec49eaa6238ad5ff00",
			response: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"object": {"sha": "d3b07384d113edec49eaa6238ad5ff00"}}`))
			})),
			want:    "d3b07384d113edec49eaa6238ad5ff00",
			wantErr: false,
		},
		{
			name:        "Test getLatestCommitHash Failure",
			expectedSha: "",
			response: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(`{"message": "Internal Server Error"}`))
			})),
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.response.Close()

			// Override the githubAPIURL constant to point to the mock server
			oldGithubAPIURL := githubAPIURL
			defer func() { githubAPIURL = oldGithubAPIURL }()
			githubAPIURL = tt.response.URL

			// Call the function to test
			sha, err := getLatestCommitHash()
			if (err != nil) != tt.wantErr {
				t.Errorf("getLatestCommitHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if sha != tt.want {
				t.Errorf("getLatestCommitHash() got = %v, want %v", sha, tt.want)
			}
		})
	}
}

//	// Mocked response data
//	mockSHA := "d3b07384d113edec49eaa6238ad5ff00"
//	mockResponse := fmt.Sprintf(`{"object": {"sha": "%s"}}`, mockSHA)
//
//	// Create a new server to mock the GitHub API response
//	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(mockResponse))
//	}))
//	defer server.Close()
//
//	// Override the githubAPIURL constant to point to the mock server
//	oldGithubAPIURL := githubAPIURL
//	defer func() { githubAPIURL = oldGithubAPIURL }()
//	githubAPIURL = server.URL
//
//	// Call the function to test
//	sha, err := getLatestCommitHash()
//	if err != nil {
//		t.Fatalf("expected no error, got %v", err)
//	}
//
//	if sha != mockSHA {
//		t.Errorf("expected %s, got %s", mockSHA, sha)
//	}
