/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package template

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

// ValidJSON checks if the provided string is a valid JSON object.
//
// The function takes a string as input and attempts to unmarshal it
// into a map[string]any using the JSON package. If the unmarshalling
// is successful, it returns true, indicating that the string is a valid JSON
// object. Otherwise, it returns false.
func ValidJSON(s string) bool {
	var js map[string]any
	return json.Unmarshal([]byte(s), &js) == nil
}

// JsonToYaml converts a JSON string into a YAML string.
//
// The function takes a JSON string as input and attempts to unmarshal it
// into an empty interface. If the unmarshalling is successful, it marshals
// the data back into a YAML string using the YAML package.
//
// On success, the function returns the YAML string and a nil error.
// If there is any error during the conversion, it returns an empty string
// and the corresponding error.
func JsonToYaml(js string) (string, error) {
	var jsonObj any
	err := json.Unmarshal([]byte(js), &jsonObj)
	if err != nil {
		return "", err
	}
	yamlBytes, err := yaml.Marshal(jsonObj)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}

// TryParse attempts to parse and execute a template with the given JSON string.
//
// The function takes two string inputs - a template string (tmpStr) and a JSON string.
// It attempts to parse the template string using the "text/template" package.
// If there is any error during parsing, the function returns the original JSON string.
//
// If the template is parsed successfully, the function attempts to execute the template
// using the provided JSON string as input data. If there is any error during execution,
// the function returns the original JSON string.
//
// On successful execution, the function returns the resulting string from the
// executed template.
func TryParse(tmpStr, jason string) string {
	tmpl, err := template.New("secret").Parse(tmpStr)
	if err != nil {
		return jason
	}

	var result map[string]any
	err = json.Unmarshal([]byte(jason), &result)
	if err != nil {
		return jason
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, result)
	if err != nil {
		return jason
	}

	return removeKeyValueWithNoValue(tpl.String())
}

// removeKeyValueWithNoValue takes an input string containing key-value pairs and filters out
// pairs where the value is "<no value>". It splits the input string into key-value pairs,
// iterates through them, and retains only the pairs with values that are not equal to "<no value>".
// The function then joins the filtered pairs back into a string and returns the resulting string.
// This function effectively removes key-value pairs with "<no value>" from the input string.
// Helpful when key-val pairs in template differs from the contents of the secret.
func removeKeyValueWithNoValue(input string) string {
	// Split the input string into key-value pairs
	pairs := strings.Split(input, ",")

	// Initialize a slice to store the filtered pairs
	var filteredPairs []string

	for _, pair := range pairs {
		keyValue := strings.SplitN(pair, ":", 2)
		if len(keyValue) == 2 && keyValue[1] != "<no value>" {
			// Add the pair to the filtered pairs if the value is not "<no value>"
			filteredPairs = append(filteredPairs, pair)
		}
	}

	// Join the filtered pairs back into a string
	result := strings.Join(filteredPairs, ",")
	return result
}
