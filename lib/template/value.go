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
	"errors"
)

// Value creates a string based on a template with embedded generator
// expressions. The generator expressions specify character ranges and lengths
// for random string parts.
//
// Parameters:
// template - The string template containing generator expressions.
//
// Returns:
// The generated string adhering to the template specifications.
// An error if any generator expression is invalid or if the string
// generation fails.
//
// Example Usage:
//
//	result, _ := Generate(`football[\w]{8}bartender`)
//	log.Printf("result0=%v", result)
//	result, _ = Generate(`admin[a-z0-9]{3}`)
//	log.Printf("result1=%v", result)
//	result, _ = Generate(`admin[a-z0-9]{3}something[\w]{3}`)
//	log.Printf("result1=%v", result)
//	result, _ = Generate(`pass[a-zA-Z0-9]{12}`)
//	log.Printf("result2=%v", result)
//	result, _ = Generate(`pass[a-Z]{8}`)
//	log.Printf("result3=%v", result)
//	result, err := Generate(`pass[z-a]{8}`)
//	log.Printf("result4=%v; err=%v", result, err)
//	result, _ = Generate(`football[\d]{8}bartender`)
//	log.Printf("result5=%v", result)
//	result, _ = Generate(`football[\symbol]{4}`)
//	log.Printf("result5=%v", result)
//
// Example Output:
//
//	2024/01/04 06:37:30 result0=football{A?1o!u9bartender
//	2024/01/04 06:37:30 result1=admin7sg
//	2024/01/04 06:37:30 result1=adminsw8something^5^
//	2024/01/04 06:37:30 result2=passqWv04txU5sKs
//	2024/01/04 06:37:30 result3=passlRxDTdMz
//	2024/01/04 06:37:30 result4=; err=invalid range specified: z-a
//	2024/01/04 06:37:30 result5=football73579557bartender
func Value(template string) (string, error) {
	result := template

	matches := generatorsExp.FindAllStringIndex(template, -1)
	if matches == nil {
		return "", errors.New("no generator expressions found")
	}

	for _, r := range matches {
		ranges, length, err := rangesAndLength(template[r[0]:r[1]])

		if err != nil {
			return "", err
		}

		positions := findExpressionPos(ranges)

		if err := replaceWithGenerated(&result, template[r[0]:r[1]],
			positions, length); err != nil {
			return "", err
		}
	}

	return result, nil
}
