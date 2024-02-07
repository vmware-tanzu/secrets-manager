/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package state

import "encoding/json"

type VSecMInternalCommand struct {
	LogLevel int `json:"logLevel"`
}

func evaluate(data string) *VSecMInternalCommand {
	var command VSecMInternalCommand
	err := json.Unmarshal([]byte(data), &command)
	if err != nil {
		return nil
	}
	return &command
}
