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

type (
	JsonTime time.Time
)

func (t *JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(*t).Format(time.RFC3339))
	return []byte(stamp), nil
}

func (t *JsonTime) String() string {
	return time.Time(*t).Format(time.RFC3339)
}

func (t *JsonTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = strings.Trim(str, "\"")

	parsedTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	// Set the time value.
	*t = JsonTime(parsedTime)

	return nil
}
