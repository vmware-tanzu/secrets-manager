package data

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitStatus(t *testing.T) {
	tests := []struct {
		name   string
		status InitStatus
		want   string
	}{
		{
			name:   "pending_status",
			status: Pending,
			want:   "pending",
		},
		{
			name:   "ready_status",
			status: Ready,
			want:   "ready",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, InitStatus(tt.want), tt.status)
		})
	}
}

func TestStatus_Increment(t *testing.T) {
	tests := []struct {
		name          string
		initialCount  int
		secretName    string
		mockLoader    func(name any) (any, bool)
		expectedCount int
	}{
		{
			name:          "increment_new_secret",
			initialCount:  0,
			secretName:    "new-secret",
			mockLoader:    func(name any) (any, bool) { return nil, false },
			expectedCount: 1,
		},
		{
			name:          "no_increment_existing_secret",
			initialCount:  1,
			secretName:    "existing-secret",
			mockLoader:    func(name any) (any, bool) { return "secret-value", true },
			expectedCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := &Status{
				NumSecrets: tt.initialCount,
				Lock:       sync.RWMutex{},
			}

			status.Increment(tt.secretName, tt.mockLoader)

			assert.Equal(t, tt.expectedCount, status.NumSecrets)
		})
	}
}

func TestStatus_Decrement(t *testing.T) {
	tests := []struct {
		name          string
		initialCount  int
		secretName    string
		mockLoader    func(name any) (any, bool)
		expectedCount int
	}{
		{
			name:          "decrement_existing_secret",
			initialCount:  1,
			secretName:    "existing-secret",
			mockLoader:    func(name any) (any, bool) { return "secret-value", true },
			expectedCount: 0,
		},
		{
			name:          "no_decrement_nonexistent_secret",
			initialCount:  1,
			secretName:    "nonexistent-secret",
			mockLoader:    func(name any) (any, bool) { return nil, false },
			expectedCount: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := &Status{
				NumSecrets: tt.initialCount,
				Lock:       sync.RWMutex{},
			}

			status.Decrement(tt.secretName, tt.mockLoader)

			assert.Equal(t, tt.expectedCount, status.NumSecrets)
		})
	}
}

func TestStatus_ConcurrentOperations(t *testing.T) {
	status := &Status{
		NumSecrets: 0,
		Lock:       sync.RWMutex{},
	}

	// Mock loaders
	mockLoaderNotFound := func(name any) (any, bool) { return nil, false }
	mockLoaderFound := func(name any) (any, bool) { return "secret-value", true }

	// Number of concurrent operations
	numOperations := 1000

	// Use WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup
	wg.Add(numOperations * 2) // For both increment and decrement operations

	// Concurrent increments
	for i := 0; i < numOperations; i++ {
		go func() {
			defer wg.Done()
			status.Increment("test-secret", mockLoaderNotFound)
		}()
	}

	// Concurrent decrements
	for i := 0; i < numOperations; i++ {
		go func() {
			defer wg.Done()
			status.Decrement("test-secret", mockLoaderFound)
		}()
	}

	// Wait for all operations to complete
	wg.Wait()

	// The final count should be 0 since we performed equal numbers of
	// increments and decrements
	assert.Equal(t, 0, status.NumSecrets)
}

func TestStatus_QueueProperties(t *testing.T) {
	tests := []struct {
		name          string
		status        Status
		expectedProps struct {
			secretQueueLen int
			secretQueueCap int
			k8sQueueLen    int
			k8sQueueCap    int
		}
	}{
		{
			name: "empty_queues",
			status: Status{
				SecretQueueLen: 0,
				SecretQueueCap: 10,
				K8sQueueLen:    0,
				K8sQueueCap:    10,
			},
			expectedProps: struct {
				secretQueueLen int
				secretQueueCap int
				k8sQueueLen    int
				k8sQueueCap    int
			}{
				secretQueueLen: 0,
				secretQueueCap: 10,
				k8sQueueLen:    0,
				k8sQueueCap:    10,
			},
		},
		{
			name: "partially_filled_queues",
			status: Status{
				SecretQueueLen: 5,
				SecretQueueCap: 10,
				K8sQueueLen:    3,
				K8sQueueCap:    10,
			},
			expectedProps: struct {
				secretQueueLen int
				secretQueueCap int
				k8sQueueLen    int
				k8sQueueCap    int
			}{
				secretQueueLen: 5,
				secretQueueCap: 10,
				k8sQueueLen:    3,
				k8sQueueCap:    10,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedProps.secretQueueLen, tt.status.SecretQueueLen)
			assert.Equal(t, tt.expectedProps.secretQueueCap, tt.status.SecretQueueCap)
			assert.Equal(t, tt.expectedProps.k8sQueueLen, tt.status.K8sQueueLen)
			assert.Equal(t, tt.expectedProps.k8sQueueCap, tt.status.K8sQueueCap)
		})
	}
}
