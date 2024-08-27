/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/vmware-tanzu/secrets-manager/sdk/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/sdk/core/constants/symbol"
	tpl "github.com/vmware-tanzu/secrets-manager/sdk/core/template"
)

// convertMapToStringBytes converts a map[string]string into a map[string][]byte,
// by converting each string value into a []byte, and returns the resulting map.
func convertMapToStringBytes(inputMap map[string]string) map[string][]byte {
	data := make(map[string][]byte)
	for k, v := range inputMap {
		data[k] = []byte(v)
	}
	return data
}

// handleTemplateFailure is used when applying a template to the secret's value
// fails. It attempts to unmarshal the 'value' string as JSON into the 'data'
// map. If the unmarshalling fails, it creates a new empty 'data' map and
// populates it with a single entry, "VALUE", containing the original 'value' as
// []byte.
func convertValueToMap(values []string) map[string][]byte {
	var data map[string][]byte

	val := ""
	if len(values) == 1 {
		val = values[0]
	} else {
		val = strings.Join(values, symbol.CollectionDelimiter)
	}

	err := json.Unmarshal([]byte(val), &data)
	if err != nil {
		data = map[string][]byte{}
		data[key.SecretDataValue] = []byte(val)
	}

	return data
}

// handleNoTemplate is used when there is no template defined.
// It attempts to unmarshal the 'value' string as JSON. If successful, it
// returns a map with the JSON key-value pairs converted to []byte values;
// otherwise, it returns a map with a single entry, "VALUE", containing the
// original 'value' as []byte.
func convertValueNoTemplate(values []string) map[string][]byte {
	var data map[string][]byte
	var jsonData map[string]string

	val := ""
	if len(values) == 1 {
		val = values[0]
	} else {
		val = strings.Join(values, symbol.CollectionDelimiter)
	}

	err := json.Unmarshal(([]byte)(val), &jsonData)
	if err != nil {
		//If error in unmarshalling, add the whole as a part of VALUE
		data[key.SecretDataValue] = ([]byte)(val)
	} else {
		//Use the secret's value as a key-val pair
		return convertMapToStringBytes(jsonData)
	}

	return data
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
		return secretData, fmt.Errorf("no values found for secret %s",
			secret.Name)
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

func transform(
	value string, tmpStr string, f SecretFormat,
) (string, error) {
	jsonData := strings.TrimSpace(value)

	parsedString := ""
	if tmpStr == "" {
		parsedString = jsonData
	} else {
		parsedString = tpl.TryParse(tmpStr, jsonData)
	}

	switch f {
	case Json:
		// If the parsed string is a valid JSON, return it as is.
		// Otherwise, assume the parsing failed and return the original
		// JSON string.
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
			// There is not much can be done at this point other than
			// returning it.
			return parsedString, nil
		}
	case Raw:
		// If the format is Raw, return the parsed string as is.
		return parsedString, nil
	default:
		// The program flow shall never enter here.
		return parsedString, fmt.Errorf("unknown format: %s", f)
	}
}
