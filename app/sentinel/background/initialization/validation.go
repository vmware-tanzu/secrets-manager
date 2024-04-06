/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package initialization

import (
	"fmt"
	"os"
	"strings"

	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func initCommandsExecutedAlready(cid *string) bool {
	log.TraceLn(cid, "checking tombstone file")

	// Parse tombstone file first:
	tombstonePath := env.InitCommandTombstonePathForSentinel()
	file, err := os.Open(tombstonePath)
	if err != nil {
		log.InfoLn(
			cid,
			"RunInitCommands: no tombstone file found... skipping custom initialization.",
		)
		return false
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.ErrorLn(cid, "Error closing tombstone file: ", err.Error())
		}
	}(file)

	data, err := os.ReadFile(tombstonePath)

	log.InfoLn(cid, fmt.Sprintf("tombstone:'%s'", string(data)))

	if strings.TrimSpace(string(data)) == "complete" {
		log.InfoLn(
			cid,
			"RunInitCommands: Already initialized. Skipping custom initialization.",
		)
		return false
	}

	return true
}
