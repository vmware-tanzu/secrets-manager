package main

import (
	"fmt"
	"os"
	"testing"
)

func TestReadWriteCommitHashToFile(t *testing.T) {
	// Create a temporary file for testing
	tempDir := os.TempDir()

	tempFile, err := os.CreateTemp(tempDir, "commit-hash")
	if err != nil {
		t.Fatalf("Error creating temporary file: %s", err)
	}
	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(tempFile)
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			fmt.Println(err.Error())
		}
	}(tempFile.Name())

	commitHashFile = tempFile.Name()

	tests := []struct {
		name        string
		commitHash  string
		expectError bool
	}{
		{
			name:       "Valid commit hash",
			commitHash: "d3b07384d113edec49eaa6238ad5ff00",
		},
		{
			name:        "Empty commit hash",
			commitHash:  "",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Write the commit hash to the file
			err := writeCommitHashToFile(tt.commitHash)
			if (err != nil) != tt.expectError {
				t.Fatalf("writeCommitHashToFile() error = %v, expectError %v", err, tt.expectError)
			}

			// Read the commit hash from the file
			readCommitHash, err := readCommitHashFromFile()
			if (err != nil) != tt.expectError {
				t.Fatalf("readCommitHashFromFile() error = %v, expectError %v", err, tt.expectError)
			}

			// Compare the written and read commit hash
			if readCommitHash != tt.commitHash {
				t.Errorf("Expected commit hash %s, but got %s", tt.commitHash, readCommitHash)
			}
		})
	}
}
