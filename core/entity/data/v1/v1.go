/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	tpl "github.com/vmware-tanzu/secrets-manager/core/template"
	"strings"
	"text/template"
	"time"
)

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RubyDate))
	return []byte(stamp), nil
}

type Secret struct {
	Name    string   `json:"name"`
	Created JsonTime `json:"created"`
	Updated JsonTime `json:"updated"`
}

type BackingStore string

var File BackingStore = "file"
var Memory BackingStore = "memory"

type SecretFormat string

var Json SecretFormat = "json"
var Yaml SecretFormat = "yaml"

type SecretMeta struct {
	// Overrides Env.SafeUseKubernetesSecrets()
	UseKubernetesSecret bool `json:"k8s"`
	// Overrides Env.SafeBackingStoreType()
	BackingStore BackingStore `json:"storage"`
	// Defaults to "default"
	Namespace string `json:"namespace"`
	// Go template used to transform the secret.
	// Sample secret:
	// '{"username":"admin","password":"VSecMRocks"}'
	// Sample template:
	// '{"USER":"{{.username}}", "PASS":"{{.password}}"}"
	Template string `json:"template"`
	// Defaults to None
	Format SecretFormat
	// For tracking purposes
	CorrelationId string `json:"correlationId"`
}

type SecretStored struct {
	// Name of the secret.
	Name string
	// Raw values. A secret can have multiple values. Sentinel returns
	// a single value if there is a single value in this array. Sentinel
	// will return an array of values if there are multiple values in the array.
	Values []string `json:"values"`
	// Transformed values. This value is the value that workloads see.
	//
	// Apply transformation (if needed) and then store the value in
	// one of the supported formats. If the format is json, ensure that
	// a valid JSON is stored here. If the format is yaml, ensure that
	// a valid YAML is stored here. If the format is none, then just
	// apply transformation (if needed) and do not do any validity check.
	ValueTransformed string `json:"valuesTransformed"`
	// Additional information that helps formatting and storing the secret.
	Meta SecretMeta
	// Timestamps
	Created time.Time
	Updated time.Time
}

// parseForK8sSecret parses the provided `SecretStored` and applies a template
// if one is defined.
//
// Args:
//
//	secret: A SecretStored struct containing the secret data and metadata.
//
// Returns:
//
//	A map of string keys to string values, containing the parsed secret data.
//
//	If there is an error during parsing or applying the template, an error
//	will be returned.
//
// Note that this function will consider only the first value in the `Values`
// collection. If there are multiple values, only the first value will be
// parsed and transformed.
func parseForK8sSecret(secret SecretStored) (map[string]string, error) {
	// cannot move this to /core/template because of circular dependency.

	secretData := make(map[string]string)

	if len(secret.Values) == 0 {
		return secretData, fmt.Errorf("no values found for secret %s", secret.Name)
	}

	jsonData := strings.TrimSpace(secret.Values[0])
	tmpStr := strings.TrimSpace(secret.Meta.Template)

	err := json.Unmarshal([]byte(jsonData), &secretData)
	if err != nil {
		return secretData, err
	}

	if tmpStr == "" {
		return secretData, err
	}

	tmpl, err := template.New("secret").Parse(tmpStr)
	if err != nil {
		return secretData, err
	}

	var t bytes.Buffer
	err = tmpl.Execute(&t, secretData)
	if err != nil {
		return secretData, err
	}

	output := make(map[string]string)
	err = json.Unmarshal(t.Bytes(), &output)
	if err != nil {
		return output, err
	}

	return output, nil
}

// ToMapForK8s returns a map that can be used to create a Kubernetes secret.
//
//  1. If there is no template, attempt to unmarshal the secret’ss value
//     into a map. If that fails, store the secret’s value under the "VALUE" key.
//  2. If there is a template, attempt to parse it. If parsing is successful,
//     create a new map with the parsed data. If parsing fails, follow the same
//     logic as in case 1, attempting to unmarshal the secret’s value into a map,
//     and if that fails, storing the secret’s value under the "VALUE" key.
func (secret SecretStored) ToMapForK8s() map[string][]byte {
	data := make(map[string][]byte)

	// If there are no values, return an empty map.
	if len(secret.Values) == 0 {
		return data
	}

	// If there is no template, use the secret’s value as is.
	if secret.Meta.Template == "" {
		err := json.Unmarshal(([]byte)(secret.Values[0]), &data)
		if err != nil {
			value := secret.Values[0]
			data["VALUE"] = ([]byte)(value)
		}

		return data
	}

	// Otherwise, apply the template.
	newData, err := parseForK8sSecret(secret)
	if err == nil {
		data = make(map[string][]byte)
		for k, v := range newData {
			data[k] = ([]byte)(v)
		}

		return data
	}

	// If the template fails, use the secret’s value as is.
	err = json.Unmarshal(([]byte)(secret.Values[0]), &data)
	if err != nil {
		value := secret.Values[0]
		data["VALUE"] = ([]byte)(value)
	}

	return data
}

// ToMap converts the SecretStored struct to a map[string]any.
// The resulting map contains the following key-value pairs:
//
//	"Name": the Name field of the SecretStored struct
//	"Values": the Values field of the SecretStored struct
//	"Created": the Created field of the SecretStored struct
//	"Updated": the Updated field of the SecretStored struct
func (secret SecretStored) ToMap() map[string]any {
	return map[string]any{
		"Name":    secret.Name,
		"Values":  secret.Values,
		"Created": secret.Created,
		"Updated": secret.Updated,
	}
}

func transform(secret SecretStored, value string) (string, error) {
	jsonData := strings.TrimSpace(value)
	tmpStr := strings.TrimSpace(secret.Meta.Template)

	parsedString := ""
	if tmpStr == "" {
		parsedString = jsonData
	} else {
		parsedString = tpl.TryParse(tmpStr, jsonData)
	}

	switch secret.Meta.Format {
	case Json:
		// If the parsed string is a valid JSON, return it as is.
		// Otherwise, assume the parsing failed and return the original JSON string.
		if tpl.ValidJSON(parsedString) {
			return parsedString, nil
		} else {
			return jsonData, nil
		}
	case Yaml:
		if tpl.ValidJSON(parsedString) {
			yml, err := tpl.JsonToYaml(parsedString)
			if err != nil {
				return parsedString, err
			}
			return yml, nil
		} else {
			// Parsed string is not a valid JSON, so return it as is.
			// It can be either a valid YAML already, or some random string.
			// There is not much can be done at this point other than returning it.
			return parsedString, nil
		}
	default:
		// The program flow shall never enter here.
		return parsedString, fmt.Errorf("unknown format: %s", secret.Meta.Format)
	}
}

// Parse takes a data.SecretStored type as input and returns the parsed
// string or an error.
//
// It parses all the `.Values` of the secret, and for each value tries to apply
// a template transformation.
//
// Here is how the template transformation is applied:
//
//  1. Compute parsedString:
//     If the Meta.Template field is empty, then parsedString is the original value.
//     Otherwise, parsedString is the result of applying the template transformation
//     to the original value.
//
// 2.	Compute the output string:
//   - If the Meta.Format field is Json, then the output string is parsedString
//     if parsedString is a valid JSON, otherwise it’s the original value.
//   - If the Meta.Format field is Yaml, then the output string is the result of
//     transforming parsedString into Yaml if parsedString is a valid JSON,
//     otherwise it’s parsedString.
func (secret SecretStored) Parse() (string, error) {
	if len(secret.Values) == 0 {
		return "", fmt.Errorf("no values found for secret %s", secret.Name)
	}

	parseFailed := false
	var results []string
	for _, v := range secret.Values {
		transformed, err := transform(secret, v)
		if err != nil {
			parseFailed = true
			continue
		}
		if transformed == "" {
			continue
		}
		results = append(results, transformed)
	}

	if results == nil {
		return "", fmt.Errorf("failed to parse secret %s", secret.Name)
	}

	if len(results) == 1 {
		if parseFailed {
			return results[0], fmt.Errorf("failed to parse secret %s", secret.Name)
		}

		return results[0], nil
	}

	marshaled, err := json.Marshal(results)
	if err != nil {
		return "", err
	}

	if parseFailed {
		return string(marshaled), fmt.Errorf("failed to parse secret %s", secret.Name)
	}

	return string(marshaled), nil
}
