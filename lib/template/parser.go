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
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Constants for different character sets used in string generation.
const (
	chars      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers    = "0123456789"
	symbols    = "!@#$%^&*()-_+={}[]|<>,.?/"
	everything = chars + numbers + symbols
)

// generatorExprRanges represents a range of characters used in string
// generation.
type generatorExprRanges [][]byte

// seedAndReturnRandom returns a random integer based on a newly seeded source.
// The randomness is seeded each time to enhance unpredictability.
//
// Parameters:
// n - The exclusive upper limit for the generated random number.
//
// Returns:
// A random integer within the range [0, n).
func seedAndReturnRandom(n int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
}

// everythingSlice returns a substring from the `everything` constant.
// It takes in two bytes representing the start and end of the desired range.
//
// Parameters:
// from - The starting byte of the range.
// to   - The ending byte of the range.
//
// Returns:
// A string containing characters from the specified range.
// An error if the specified range is invalid.
func everythingSlice(from, to byte) (string, error) {
	l, r := strings.Index(everything, string(from)),
		strings.LastIndex(everything, string(to))

	if l > r {
		return "", fmt.Errorf("invalid range specified: %s-%s", string(from),
			string(to))
	}

	return everything[l:r], nil
}

// replaceWithGenerated replaces a substring in a given string with randomly
// generated characters.
// The generated string adheres to specified character ranges and lengths.
//
// Parameters:
// s          - A pointer to the string to be modified.
// expression - The substring to be replaced.
// ranges     - A slice of byte slices, each representing a character range.
// length     - The length of the generated string.
//
// Returns:
// An error if the character range is empty or invalid.
func replaceWithGenerated(s *string, expression string, ranges [][]byte,
	length int) error {
	var alphabet string

	for _, r := range ranges {
		switch string(r[0]) + string(r[1]) {
		case `\w`:
			alphabet += everything
		case `\d`:
			alphabet += numbers
		case `\x`:
			alphabet += symbols
		default:
			if slice, err := everythingSlice(r[0], r[1]); err != nil {
				return err
			} else {
				alphabet += slice
			}
		}
	}

	if len(alphabet) == 0 {
		return fmt.Errorf("empty range in expression: %s", expression)
	}

	result := make([]byte, length)

	for i := 0; i <= length-1; i++ {
		result[i] = alphabet[seedAndReturnRandom(len(alphabet))]
	}

	*s = strings.Replace(*s, expression, string(result), 1)

	return nil
}

// Matches things like a-z 1-p, \d, \w, etc.
var rangeExp = regexp.MustCompile(`(\\?[a-zA-Z0-9]-?[a-zA-Z0-9]?)`)

// Matches things like [a-z]{8}, [A-Z]{8} \d \w, etc.
var generatorsExp = regexp.MustCompile(`\[([a-zA-Z0-9\-\\]+)](\{([0-9]+)})`)

// findExpressionPos finds the positions of character ranges in a string.
//
// Parameters:
// s - The string containing the character ranges.
//
// Returns:
// A generatorExprRanges representing the found character ranges.
func findExpressionPos(s string) generatorExprRanges {
	matches := rangeExp.FindAllStringIndex(s, -1)
	result := make(generatorExprRanges, len(matches), len(matches))

	for i, r := range matches {
		result[i] = []byte{s[r[0]], s[r[1]-1]}
	}

	return result
}

// rangesAndLength extracts a character range expression and its specified
// length from a string.
//
// Parameters:
// s - The string containing the expression and length specification.
//
// Returns:
// The extracted character range expression.
// The specified length for string generation.
// An error if the length parsing fails.
func rangesAndLength(s string) (string, int, error) {
	expr := s[0:strings.LastIndex(s, "{")]

	length, err := parseLength(s)

	return expr, length, err
}

// parseLength extracts and parses the length part of a generator expression.
//
// Parameters:
// s - The string containing the length specification.
//
// Returns:
// The parsed length as an integer.
// An error if the length parsing fails.
func parseLength(s string) (int, error) {
	lengthStr := s[strings.LastIndex(s, "{")+1 : len(s)-1]
	if l, err := strconv.Atoi(lengthStr); err != nil {
		return 0, fmt.Errorf("unable to parse length from %v", s)
	} else {
		return l, nil
	}
}
