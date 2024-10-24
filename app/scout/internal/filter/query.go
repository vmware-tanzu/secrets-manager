/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package filter

import (
	"fmt"
	"strings"
)

func ValueFromPath(data any, path string) (any, error) {
	parts := strings.Split(path, ".")

	// If no path specified, return the data as is.
	if path == "" || !strings.Contains(path, ".") {
		return data, nil
	}

	var current = data
	for _, part := range parts {
		switch v := current.(type) {
		case map[string]interface{}:
			if val, ok := v[part]; ok {
				current = val
			} else {
				return nil, fmt.Errorf("key not found: %s", part)
			}
		case []interface{}:
			return nil, fmt.Errorf("arrays are not supported in path queries")
		default:
			return nil, fmt.Errorf("cannot navigate further from %v", current)
		}
	}

	return current, nil
}
