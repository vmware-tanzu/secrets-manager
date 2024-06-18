package main

import (
	"os"
	"testing"
)

func Test_createLockFile(t *testing.T) {
	tests := []struct {
		name    string
		setup   func()
		wantErr bool
	}{

		{
			name: "Create lock file Exist Error",
			setup: func() {
				tempDir := os.TempDir()

				tempFile, err := os.CreateTemp(tempDir, "git_poller.lock")
				if err != nil {
					t.Fatalf("Error creating temporary file: %s", err)
				}

				lockFilePath = tempFile.Name()

				err = tempFile.Close()
				if err != nil {
					return
				}
			},
			wantErr: true,
		},
		{
			name: "Create lock file Success",
			setup: func() {
				tempDir := os.TempDir()

				lockFilePath = tempDir + "/git_poller.lock"

			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			if err := createLockFile(); (err != nil) != tt.wantErr {
				t.Errorf("createLockFile() error = %v, wantErr %v", err, tt.wantErr)
			}

			removeLockFile()
		})
	}
}
