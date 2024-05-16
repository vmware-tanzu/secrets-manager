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
	"io"
	"math"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/secrets-manager/core/backoff"
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
		return errors.Wrap(err, "saveSecretToDisk: failed to marshal secret")
	}

	file, err := os.Create(dataPath)
	if err != nil {
		return errors.Wrap(err, "saveSecretToDisk: failed to create file")
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

// PersistToDisk saves a given secret to disk and also creates a backup copy
// of the secret. The function is designed to enhance data durability through
// retries and backup management based on environmental configurations.
//
// Parameters:
//   - secret (entity.SecretStored): The secret to be saved, which is a
//     structured entity containing the secret's name and possibly other
//     metadata or the secret data itself.
//   - errChan (chan<- error): A channel through which errors are reported. This
//     channel allows the function to operate asynchronously, notifying the
//     caller of any issues in the process of persisting the secret.
func PersistToDisk(secret entity.SecretStored, errChan chan<- error) {
	backupCount := env.SecretBackupCountForSafe()

	// Save the secret
	dataPath := path.Join(env.DataPathForSafe(), secret.Name+".age")

	err := saveSecretToDisk(secret, dataPath)
	if err != nil {
		// Retry once more.
		time.Sleep(500 * time.Millisecond)
		err := saveSecretToDisk(secret, dataPath)
		if err != nil {
			errChan <- err
		}
	}

	lastBackupIndexLock.Lock()
	index, found := lastBackedUpIndex[secret.Name]
	if !found {
		lastBackedUpIndex[secret.Name] = 0
		index = 0
	}
	lastBackupIndexLock.Unlock()

	newIndex := math.Mod(float64(index+1), float64(backupCount))

	// Save a copy
	dataPath = path.Join(
		env.DataPathForSafe(),
		secret.Name+"-"+strconv.Itoa(int(newIndex))+"-"+".age.backup",
	)

	err = backoff.RetryExponential("PersistToDisk", func() error {
		return saveSecretToDisk(secret, dataPath)
	})

	if err != nil {
		errChan <- err
		// Do not change lastBackedUpIndex
		// since the backup was not successful.
		return
	}

	lastBackupIndexLock.Lock()
	lastBackedUpIndex[secret.Name] = int(newIndex)
	lastBackupIndexLock.Unlock()
}
