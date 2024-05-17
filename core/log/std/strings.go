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

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

// toCustomCase formats a string to a custom case, replacing underscores
// with spaces and capitalizing words.
func toCustomCase(input string) string {
	caser := cases.Title(language.English)
	return strings.ReplaceAll(caser.String(strings.ToLower(input)), "_", " ")
}
