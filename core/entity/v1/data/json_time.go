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
	"fmt"
	"strings"
	"time"
)

// JsonTime wraps the standard time.Time type to provide JSON serialization and
// deserialization in RFC3339 format. This type ensures that the JSON
// representation of dates and times in Go applications follows a standard and
// easily interchangeable format.
type JsonTime time.Time

// MarshalJSON converts the JsonTime value to a JSON-formatted string in
// RFC3339 format. This method ensures JsonTime can be directly marshaled into
// a JSON string.
//
// Returns:
//   - A byte slice containing the JSON-formatted date and time string.
//   - An error if the formatting fails, though in practice this method should
//     not error out since the time formatting used (RFC3339) is a valid and
//     supported format.
func (t *JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(*t).Format(time.RFC3339))
	return []byte(stamp), nil
}

// String returns the JsonTime as a string formatted according to RFC3339.
// This method provides a standard way to convert a JsonTime object to a
// human-readable string.
func (t *JsonTime) String() string {
	return time.Time(*t).Format(time.RFC3339)
}

// UnmarshalJSON parses a JSON-formatted string in RFC3339 format and sets
// the JsonTime accordingly. This method enables JsonTime to directly receive
// and parse time information from JSON data.
//
// Parameters:
//   - data: a byte slice containing the JSON string to be parsed.
//
// Returns:
//   - An error if the string is not in valid RFC3339 format or if the parsing
//     fails.
func (t *JsonTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = strings.Trim(str, "\"")

	parsedTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	*t = JsonTime(parsedTime)

	return nil
}
