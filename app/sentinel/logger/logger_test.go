package logger

import (
	"io"
	"os"
	"strings"
	"testing"
)

func setup() {
	CreateLogger()
}

func teardown() {
	os.Remove(LogPipePath)
}

func TestCreateLogger_Success(t *testing.T) {
	// given
	setup()
	defer teardown()

	// then
	if _, err := os.Stat(LogPipePath); os.IsNotExist(err) {
		t.Errorf("CreateLogger failed, %s does not exist", LogPipePath)
	}
}

func TestSendLog_WriteSuccess(t *testing.T) {
	// given
	setup()
	defer teardown()
	expectedMessage := "Test message"

	// when
	go SendLog(expectedMessage)
	pipeFile, err := os.OpenFile(LogPipePath, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		t.Fatalf("Failed to open log pipe for reading: %v", err)
	}
	defer pipeFile.Close()

	content, err := io.ReadAll(pipeFile)
	if err != nil {
		t.Fatalf("Failed to read from log pipe: %v", err)
	}

	// then
	if !strings.Contains(string(content), expectedMessage) {
		t.Errorf("Log message not found. Expected '%s', got '%s'", expectedMessage, string(content))
	}
}
