package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitRootKeys(t *testing.T) {
	tests := []struct {
		name         string
		command      SentinelCommand
		expectedKeys []string
	}{
		{
			name: "multiple_keys",
			command: SentinelCommand{
				SerializedRootKeys: "key1\nkey2\nkey3",
			},
			expectedKeys: []string{"key1", "key2", "key3"},
		},
		{
			name: "single_key",
			command: SentinelCommand{
				SerializedRootKeys: "singlekey",
			},
			expectedKeys: []string{"singlekey"},
		},
		{
			name: "empty_string",
			command: SentinelCommand{
				SerializedRootKeys: "",
			},
			expectedKeys: []string{""},
		},
		{
			name: "keys_with_empty_lines",
			command: SentinelCommand{
				SerializedRootKeys: "key1\n\nkey2\n\nkey3",
			},
			expectedKeys: []string{"key1", "", "key2", "", "key3"},
		},
		{
			name: "keys_with_spaces",
			command: SentinelCommand{
				SerializedRootKeys: "key 1\nkey 2\nkey 3",
			},
			expectedKeys: []string{"key 1", "key 2", "key 3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.command.SplitRootKeys()
			assert.Equal(t, tt.expectedKeys, result,
				"SplitRootKeys() returned unexpected keys")
		})
	}
}

func TestSentinelCommandStructure(t *testing.T) {
	tests := []struct {
		name    string
		command SentinelCommand
	}{
		{
			name: "complete_command",
			command: SentinelCommand{
				WorkloadIds:        []string{"workload1", "workload2"},
				Namespaces:         []string{"namespace1", "namespace2"},
				Secret:             "secret123",
				Template:           "template123",
				DeleteSecret:       true,
				AppendSecret:       false,
				Format:             "json",
				Encrypt:            true,
				NotBefore:          "2024-01-01",
				Expires:            "2025-01-01",
				SerializedRootKeys: "key1\nkey2",
				ShouldSleep:        true,
				SleepIntervalMs:    1000,
			},
		},
		{
			name: "minimal_command",
			command: SentinelCommand{
				WorkloadIds: []string{"workload1"},
				Secret:      "secret123",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test that WorkloadIds is a slice
			assert.IsType(t, []string{}, tt.command.WorkloadIds,
				"WorkloadIds should be a string slice")

			// Test that Namespaces is a slice
			assert.IsType(t, []string{}, tt.command.Namespaces,
				"Namespaces should be a string slice")

			// Test boolean fields
			assert.IsType(t, true, tt.command.DeleteSecret,
				"DeleteSecret should be a boolean")
			assert.IsType(t, true, tt.command.AppendSecret,
				"AppendSecret should be a boolean")
			assert.IsType(t, true, tt.command.Encrypt,
				"Encrypt should be a boolean")
			assert.IsType(t, true, tt.command.ShouldSleep,
				"ShouldSleep should be a boolean")

			// Test integer field
			assert.IsType(t, 0, tt.command.SleepIntervalMs,
				"SleepIntervalMs should be an integer")
		})
	}
}

func TestVSecMInternalCommandStructure(t *testing.T) {
	tests := []struct {
		name     string
		command  VSecMInternalCommand
		logLevel int
	}{
		{
			name: "debug_level",
			command: VSecMInternalCommand{
				LogLevel: 0,
			},
			logLevel: 0,
		},
		{
			name: "info_level",
			command: VSecMInternalCommand{
				LogLevel: 1,
			},
			logLevel: 1,
		},
		{
			name: "warning_level",
			command: VSecMInternalCommand{
				LogLevel: 2,
			},
			logLevel: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test that LogLevel is an integer
			assert.IsType(t, 0, tt.command.LogLevel,
				"LogLevel should be an integer")

			// Test that LogLevel matches expected value
			assert.Equal(t, tt.logLevel, tt.command.LogLevel,
				"LogLevel should match expected value")
		})
	}
}
