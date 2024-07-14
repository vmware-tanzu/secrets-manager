/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package io

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

var lastBackedUpIndex = make(map[string]int)

// Only one thread reaches lastBackupIndex at a time;
// but still using this lock for defensive programming.
var lastBackupIndexLock = sync.Mutex{}

func saveSecretToDisk(secret entity.SecretStored, dataPath string) error {
	data, err := json.Marshal(secret)
	if err != nil {
		return errors.Join(
			err,
			errors.New("saveSecretToDisk: failed to marshal secret"),
		)
	}

	file, err := os.Create(dataPath)
	if err != nil {
		return errors.Join(
			err,
			errors.New("saveSecretToDisk: failed to create file"),
		)
	}
	defer func(f io.ReadCloser) {
		err := f.Close()
		if err != nil {
			id := crypto.Id()
			log.InfoLn(&id, "saveSecretToDisk: problem closing file", err.Error())
		}
	}(file)

	if env.FipsCompliantModeForSafe() {
		return crypto.EncryptToWriterAes(file, string(data))
	}

	return crypto.EncryptToWriterAge(file, string(data))
}
