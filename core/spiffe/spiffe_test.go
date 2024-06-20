/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package spiffe

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/stretchr/testify/assert"
	"github.com/vmware-tanzu/secrets-manager/lib/spiffe"
	"net/http/httptest"
	"net/url"
	"testing"
)

// Mock certificate generation
func generateMockCertificate(t *testing.T) *x509.Certificate {
	uri, err := spiffeid.FromString("spiffe://example.org/service")
	if err != nil {
		t.Fatalf("failed to create SPIFFE ID: %v", err)
	}

	// Create a mock certificate with the URI
	return &x509.Certificate{
		URIs: []*url.URL{uri.URL()},
	}
}

func TestIdFromRequest(t *testing.T) {
	tests := []struct {
		name        string
		tlsState    *tls.ConnectionState
		expectedID  *spiffeid.ID
		expectedErr string
	}{
		{
			name: "No peer certificates",
			tlsState: &tls.ConnectionState{
				PeerCertificates: nil,
			},
			expectedID:  nil,
			expectedErr: "no peer certs",
		},
		{
			name: "Valid peer certificate",
			tlsState: &tls.ConnectionState{
				PeerCertificates: []*x509.Certificate{
					generateMockCertificate(t),
				},
			},
			expectedID: func() *spiffeid.ID {
				id, _ := spiffeid.FromString("spiffe://example.org/service")
				return &id
			}(),
			expectedErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "https://example.com", nil)
			req.TLS = tt.tlsState

			id, err := spiffe.IdFromRequest(req)
			if tt.expectedErr != "" {
				assert.Nil(t, id)
				assert.EqualError(t, err, tt.expectedErr)
			} else {
				assert.NotNil(t, id)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedID, id)
			}
		})
	}
}

func TestIdAsString(t *testing.T) {
	req := httptest.NewRequest("GET", "https://example.com", nil)
	req.TLS = &tls.ConnectionState{
		PeerCertificates: []*x509.Certificate{
			generateMockCertificate(t),
		},
	}

	spiffeIDString := spiffe.IdAsString(req)

	expectedID := "spiffe://example.org/service"
	assert.Equal(t, expectedID, spiffeIDString)
}
