package logger

import (
	"testing"
)

func TestSetAndGetLevel(t *testing.T) {
	tests := []struct {
		name     string
		setLevel Level
		want     Level
	}{
		{"Set to Off", Off, Off},
		{"Set to Error", Error, Error},
		{"Set to Debug", Debug, Debug},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			SetLevel(tc.setLevel)
			if got := GetLevel(); got != tc.want {
				t.Errorf("After SetLevel(%v), GetLevel() = %v, want %v", tc.setLevel, got, tc.want)
			}
		})
	}
}

func TestLogTextBuilder(t *testing.T) {
	// Temporarily override currentTime for consistent testing
	originalCurrentTime := currentTime
	currentTime = func() string { return "2024-02-13T12:00:00" }
	defer func() { currentTime = originalCurrentTime }()

	tests := []struct {
		name      string
		logHeader string
		messages  []any
		want      string
	}{
		{
			name:      "Info level log",
			logHeader: "[SENTINEL_INFO]",
			messages:  []any{"Test", "message"},
			want:      "[SENTINEL_INFO][2024-02-13T12:00:00] Test message\n",
		},
		{
			name:      "Debug level log",
			logHeader: "[SENTINEL_DEBUG]",
			messages:  []any{"Another", "test", 123},
			want:      "[SENTINEL_DEBUG][2024-02-13T12:00:00] Another test 123\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LogTextBuilder(tc.logHeader, tc.messages...)
			if got != tc.want {
				t.Errorf("LogTextBuilder() = %q, want %q", got, tc.want)
			}
		})
	}
}
