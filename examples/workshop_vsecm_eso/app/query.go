package main

import (
	"fmt"
	"strings"
)

func getValueFromPath(data interface{}, path string) (interface{}, error) {
	parts := strings.Split(path, ".")

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
