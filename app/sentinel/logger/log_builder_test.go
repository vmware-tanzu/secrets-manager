package logger

import (
	"testing"
	"time"
)

func mockTime() string {
	t, _ := time.Parse(time.DateTime, "2022-01-02 15:04:05")
	return t.Format(time.DateTime)
}

func TestLogTextBuilder(t *testing.T) {
	oldTimeNowFunc := currentTime
	defer func() { currentTime = oldTimeNowFunc }()
	currentTime = mockTime

	tests := []struct {
		name     string
		args     []any
		expected string
	}{
		{
			name:     "Single argument",
			args:     []any{"Hello, World!"},
			expected: "[LOG][2022-01-02 15:04:05] Hello, World!\n",
		},
		{
			name:     "No arguments",
			args:     []any{},
			expected: "[LOG][2022-01-02 15:04:05] \n",
		},
		{
			name:     "Multiple arguments",
			args:     []any{"Multiple", 123, true},
			expected: "[LOG][2022-01-02 15:04:05] Multiple 123 true\n",
		},
		{
			name:     "Special characters",
			args:     []any{"Special @#$%^&*()"},
			expected: "[LOG][2022-01-02 15:04:05] Special @#$%^&*()\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LogTextBuilder(tt.args...)
			if result != tt.expected {
				t.Errorf("Expected %q but got %q", tt.expected, result)
			}
		})
	}
}
