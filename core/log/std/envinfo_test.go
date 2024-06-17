/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package std

//func TestSortKeys(t *testing.T) {
//	tests := []struct {
//		name     string
//		input    map[string]string
//		expected []string
//	}{
//		{
//			name:     "Empty map",
//			input:    map[string]string{},
//			expected: []string{},
//		},
//		{
//			name: "Map with one key",
//			input: map[string]string{
//				"key1": "value1",
//			},
//			expected: []string{"key1"},
//		},
//		{
//			name: "Map with multiple keys",
//			input: map[string]string{
//				"key3": "value3",
//				"key1": "value1",
//				"key2": "value2",
//			},
//			expected: []string{"key1", "key2", "key3"},
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			result := sortKeys(tt.input)
//			if !reflect.DeepEqual(result, tt.expected) {
//				if len(result) == 0 && len(tt.expected) == 0 {
//					return
//				}
//				t.Errorf("Unexpected result for %s. Got %v, expected %v", tt.name, result, tt.expected)
//			}
//		})
//	}
//}
//
//func TestToCustomCase(t *testing.T) {
//	tests := []struct {
//		name     string
//		input    string
//		expected string
//	}{
//		{
//			name:     "Empty string",
//			input:    "",
//			expected: "",
//		},
//		{
//			name:     "Single word",
//			input:    "hello",
//			expected: "Hello",
//		},
//		{
//			name:     "Multiple words separated by underscores",
//			input:    "snake_case_example",
//			expected: "Snake case example",
//		},
//		{
//			name:     "Uppercase letters",
//			input:    "UPPER_CASE",
//			expected: "Upper case",
//		},
//		{
//			name:     "Mixed case and underscores",
//			input:    "Camel_Case_Example",
//			expected: "Camel case example",
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			result := toCustomCase(tt.input)
//			if result != tt.expected {
//				t.Errorf("Unexpected result for %s. Got %s, expected %s", tt.name, result, tt.expected)
//			}
//		})
//	}
//}
//
//func TestGetAllEnvironmentVariables(t *testing.T) {
//	// Set temporary environment variables for testing
//	os.Setenv("TEST_VAR1", "value1")
//	os.Setenv("TEST_VAR2", "value2")
//	os.Setenv("TEST_VAR3", "value3")
//
//	// Ensure the environment variables are unset after the test
//	defer func() {
//		os.Unsetenv("TEST_VAR1")
//		os.Unsetenv("TEST_VAR2")
//		os.Unsetenv("TEST_VAR3")
//	}()
//
//	// Expected result based on the temporary environment variables
//	expected := []string{"TEST_VAR1", "TEST_VAR2", "TEST_VAR3"}
//
//	// Run the function
//	result := getAllEnvironmentVariables()
//
//	// Check if all expected variables exist in the result
//	for _, key := range expected {
//		found := false
//		for _, resultKey := range result {
//			if key == resultKey {
//				found = true
//				break
//			}
//		}
//		if !found {
//			t.Errorf("Expected variable %s not found in result %v", key, result)
//		}
//	}
//}
//
//func TestGetMaxEnvVarLength(t *testing.T) {
//	// Test case 1: Empty slice
//	envVars1 := []string{}
//	result1 := getMaxEnvVarLength(envVars1)
//	if result1 != 0 {
//		t.Errorf("Expected 0 for an empty slice, but got %d", result1)
//	}
//
//	// Test case 2: Slice with various lengths
//	envVars2 := []string{"VAR1", "VAR22", "VAR333", "VAR4444"}
//	result2 := getMaxEnvVarLength(envVars2)
//	expected2 := 7 // "VAR4444" has the maximum length
//	if result2 != expected2 {
//		t.Errorf("Expected %d for the second test case, but got %d", expected2, result2)
//	}
//
//	// Test case 3: Slice with equal lengths
//	envVars3 := []string{"ABC", "DEF", "GHI", "JKL"}
//	result3 := getMaxEnvVarLength(envVars3)
//	expected3 := 3 // All have the same length
//	if result3 != expected3 {
//		t.Errorf("Expected %d for the third test case, but got %d", expected3, result3)
//	}
//}
//
//func TestPrintFormattedInfo(t *testing.T) {
//	// Prepare the input data
//	id := "testID"
//	info := map[string]string{
//		"var1":      "value1",
//		"variable2": "value2",
//		"env_var3":  "value3",
//	}
//
//	// Create a pipe to capture stdout
//	r, w, err := os.Pipe()
//	if err != nil {
//		t.Fatal(err)
//	}
//	defer r.Close()
//	defer w.Close()
//
//	// Redirect stdout to the pipe
//	oldStdout := os.Stdout
//	os.Stdout = w
//
//	// Call the function
//	printFormattedInfo(&id, info)
//
//	// Restore stdout
//	w.Close()
//	os.Stdout = oldStdout
//
//	// Read from the pipe to get the captured output
//	var capturedOutput strings.Builder
//	_, err = io.Copy(&capturedOutput, r)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	// Expected output
//	expectedOutput := "testID  Env var3: value3\ntestID      Var1: value1\ntestID Variable2: value2\n"
//
//	// Compare actual vs expected output
//	if capturedOutput.String() != expectedOutput {
//		t.Errorf("Unexpected output:\nExpected: \n%s\nActual: \n%s", expectedOutput, capturedOutput.String())
//	}
//}


//func mockLogger(t *testing.T) (*bufio.Scanner, *os.File, *os.File) {
//	reader, writer, err := os.Pipe()
//	if err != nil {
//		log.Println(err)
//	}
//	log.SetOutput(writer)
//
//	return bufio.NewScanner(reader), reader, writer
//}
//
//func resetLogger(reader *os.File, writer *os.File) {
//	err := reader.Close()
//	if err != nil {
//		fmt.Println("error closing reader was ", err)
//	}
//	if err = writer.Close(); err != nil {
//		fmt.Println("error closing writer was ", err)
//	}
//	log.SetOutput(os.Stderr)
//}
//
//func TestPrintSpecificEnvironmentVariables(t *testing.T) {
//
//	scanner, reader, writer := mockLogger(t)
//	defer resetLogger(reader, writer)
//	// Prepare the input data
//	id := "testID"
//	envVarsToPrint := []string{"EXISTING_VAR", "NON_EXISTING_VAR"}
//	info := make(map[string]string)
//
//	// Set an existing environment variable for testing
//	os.Setenv("EXISTING_VAR", "existing_value")
//
//	// Call the function
//	printSpecificEnvironmentVariables(&id, envVarsToPrint, info)
//
//	scanner.Scan()                   // blocks until a new line is written to the pipe
//	capturedOutput := scanner.Text() // the last line written to the scanner
//
//	// Verify that the existing environment variable is captured
//	existingVarValue, exists := info["EXISTING_VAR"]
//	if !exists || existingVarValue != "existing_value" {
//		t.Errorf("Expected existing environment variable 'EXISTING_VAR' to be captured, but it was not.")
//	}
//
//	// Verify that the non-existing environment variable triggers a warning in the mockWarnLogger
//	nonExistingVarWarning := "Warning: Environment variable NON_EXISTING_VAR not found"
//	if !strings.Contains(capturedOutput, nonExistingVarWarning) {
//		t.Errorf("Expected warning message for non-existing variable:\nExpected: %s\nActual: %s",
//			nonExistingVarWarning, capturedOutput)
//	}
//}
//
//func TestPrintEnvironmentInfo(t *testing.T) {
//	// Prepare the input data
//	id := "testID"
//	envVarsToPrint := []string{"EXISTING_VAR", "NON_EXISTING_VAR"}
//
//	// Set an existing environment variable for testing
//	os.Setenv("EXISTING_VAR", "existing_value")
//
//	scanner, reader, writer := mockLogger(t)
//	defer resetLogger(reader, writer)
//
//	// Redirect stdout to capture the printed output
//	oldStdout := os.Stdout
//	r, w, _ := os.Pipe()
//	os.Stdout = w
//
//	// Call the function
//	PrintEnvironmentInfo(&id, envVarsToPrint)
//
//	scanner.Scan() // blocks until a new line is written to the pipe
//	loggerOutput := scanner.Text()
//
//	// Close the write end of the pipe and restore stdout
//	w.Close()
//	os.Stdout = oldStdout
//
//	// Read from the pipe to get the captured output
//	var capturedOutput strings.Builder
//	_, _ = io.Copy(&capturedOutput, r)
//
//	// Verify that the non-existing environment variable triggers a warning
//	nonExistingVarWarning := "Warning: Environment variable NON_EXISTING_VAR not found"
//	if !strings.Contains(loggerOutput, nonExistingVarWarning) {
//		t.Errorf("Expected warning message for non-existing variable in the captured output:\nExpected: %s\nActual: %s",
//			nonExistingVarWarning, loggerOutput)
//	}
//
//	if !strings.Contains(capturedOutput.String(), "EXISTING_VAR") {
//		t.Errorf("Expected existing environment variable 'EXISTING_VAR' to be captured, but it was not.:\nExpected: %s\nActual: %s",
//			nonExistingVarWarning, capturedOutput.String())
//	}
//
//}
