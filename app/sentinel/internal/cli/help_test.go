/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package cli

import (
	"log"
	"os"
	"testing"

	"github.com/akamensky/argparse"
	"github.com/stretchr/testify/assert"

	"github.com/vmware-tanzu/secrets-manager/core/constants/sentinel"
)

// captureOutput captures log output. It sets the os.Stdout to a pipe and returns the output as a string. It also restores the original os.Stdout.
func captureFmtOutput(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	f()

	os.Stdout = stdout
	w.Close()

	out := make([]byte, 1000)
	n, err := r.Read(out)
	if err != nil {
		log.Fatal(err)
	}

	return string(out[:n])
}
func TestPrintUsage(t *testing.T) {
	parser := argparse.NewParser(sentinel.CmdName, "Test CLI")
	expectedOutput := parser.Usage(sentinel.CmdName)

	output := captureFmtOutput(func() {
		PrintUsage(parser)
	})

	assert.Equal(t, expectedOutput, output)
}

func TestInputValidationFailure(t *testing.T) {
	tests := []struct {
		name           string
		workload       *[]string
		encrypt        *bool
		inputKeys      *string
		secret         *string
		deleteSecret   *bool
		expectedResult bool
	}{
		{
			name:           "No Workload Name, Not Encrypting, No Input Keys",
			workload:       &[]string{},
			encrypt:        boolPointer(false),
			inputKeys:      stringPointer(""),
			secret:         stringPointer("secret1"),
			deleteSecret:   boolPointer(false),
			expectedResult: true,
		},
		{
			name:           "Secret is Empty, Not Deleting Secret, No Input Keys",
			workload:       &[]string{"workload1"},
			encrypt:        boolPointer(false),
			inputKeys:      stringPointer(""),
			secret:         stringPointer(""),
			deleteSecret:   boolPointer(false),
			expectedResult: true,
		},
		{
			name:           "Valid Input - Workload Provided",
			workload:       &[]string{"workload1"},
			encrypt:        boolPointer(false),
			inputKeys:      stringPointer(""),
			secret:         stringPointer("secret1"),
			deleteSecret:   boolPointer(false),
			expectedResult: false,
		},
		{
			name:           "Valid Input - Input Keys Provided",
			workload:       &[]string{},
			encrypt:        boolPointer(false),
			inputKeys:      stringPointer("key1"),
			secret:         stringPointer("secret1"),
			deleteSecret:   boolPointer(false),
			expectedResult: false,
		},
		{
			name:           "InValid Input - Deleting Secret True",
			workload:       &[]string{},
			encrypt:        boolPointer(false),
			inputKeys:      stringPointer(""),
			secret:         stringPointer("secret1"),
			deleteSecret:   boolPointer(true),
			expectedResult: true,
		},
		{
			name:           "InValid Input - Deleting Secret True, Secret Provided",
			workload:       &[]string{},
			encrypt:        boolPointer(false),
			inputKeys:      stringPointer(""),
			secret:         stringPointer("secret1"),
			deleteSecret:   boolPointer(true),
			expectedResult: true,
		},
		{
			name:           "Valid Input - Deleting Secret True, No Secret Provided",
			workload:       &[]string{"workload1"},
			encrypt:        boolPointer(false),
			inputKeys:      stringPointer(""),
			secret:         stringPointer("secret1"),
			deleteSecret:   boolPointer(true),
			expectedResult: false,
		},
		{
			name:           "Valid Input - Deleting Secret True, Input Keys Provided",
			workload:       &[]string{},
			encrypt:        boolPointer(false),
			inputKeys:      stringPointer("key1"),
			secret:         stringPointer("secret1"),
			deleteSecret:   boolPointer(false),
			expectedResult: false,
		},
		{
			name:           "Valid Input - Encrypt True, No Workload Name, No Input Keys",
			workload:       &[]string{},
			encrypt:        boolPointer(true),
			inputKeys:      stringPointer(""),
			secret:         stringPointer("secret1"),
			deleteSecret:   boolPointer(false),
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InputValidationFailure(tt.workload, tt.encrypt, tt.inputKeys, tt.secret, tt.deleteSecret)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

// Helper functions to create pointers to primitive types
func boolPointer(b bool) *bool {
	return &b
}

func stringPointer(s string) *string {
	return &s
}
