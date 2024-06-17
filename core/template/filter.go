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

import "strings"

const separator = ","
const delimiter = ":"
const empty = "<empty>"

// removeKeyValueWithNoValue takes an input string containing key-value pairs
// and filters out pairs where the value is "<empty>". It splits the input
// string into key-value pairs, iterates through them, and retains only the
// pairs with values that are not equal to "<empty>".
// The function then joins the filtered pairs back into a string and returns the
// resulting string. This function effectively removes key-value pairs with
// "<empty>" from the input string. Helpful for data cleaning and filtering
// when you want to omit certain key/value pairs from a template.
func removeKeyValueWithNoValue(input string) string {
	// Split the input string into key-value pairs
	pairs := strings.Split(input, separator)

	// Initialize a slice to store the filtered pairs
	var filteredPairs []string

	for _, pair := range pairs {
		keyValue := strings.SplitN(pair, delimiter, 2)
		if len(keyValue) == 2 && keyValue[1] != empty {
			// Add the pair to the filtered pairs if the value is not
			// "<empty>"
			filteredPairs = append(filteredPairs, pair)
		}
	}

	// Join the filtered pairs back into a string
	result := strings.Join(filteredPairs, ",")
	return result
}
