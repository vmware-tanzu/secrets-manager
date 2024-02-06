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

import (
	"encoding/json"
	"github.com/pkg/errors"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"math"
	"os"
	"path"
	"strconv"
	"time"
)

var lastBackedUpIndex = make(map[string]int)

func saveSecretToDisk(secret entity.SecretStored, dataPath string) error {
	data, err := json.Marshal(secret)
	if err != nil {
		return errors.Wrap(err, "saveSecretToDisk: failed to marshal secret")
	}

	file, err := os.Create(dataPath)
	if err != nil {
		return errors.Wrap(err, "saveSecretToDisk: failed to create file")
	}
	defer func() {
		err := file.Close()
		if err != nil {
			id := "AEGIIOCL"
			log.InfoLn(&id, "saveSecretToDisk: problem closing file", err.Error())
		}
	}()

	if env.SafeFipsCompliant() {
		return encryptToWriterAes(file, string(data))
	}

	return encryptToWriterAge(file, string(data))
}

// Only one goroutine accesses this function at any given time.
func persist(secret entity.SecretStored, errChan chan<- error) {
	backupCount := env.SafeSecretBackupCount()

	// Save the secret
	dataPath := path.Join(env.SafeDataPath(), secret.Name+".age")

	err := saveSecretToDisk(secret, dataPath)
	if err != nil {
		// Retry once more.
		time.Sleep(500 * time.Millisecond)
		err := saveSecretToDisk(secret, dataPath)
		if err != nil {
			errChan <- err
		}
	}

	index, found := lastBackedUpIndex[secret.Name]
	if !found {
		lastBackedUpIndex[secret.Name] = 0
		index = 0
	}

	newIndex := math.Mod(float64(index+1), float64(backupCount))

	// Save a copy
	dataPath = path.Join(
		env.SafeDataPath(),
		secret.Name+"-"+strconv.Itoa(int(newIndex))+"-"+".age.backup",
	)

	err = saveSecretToDisk(secret, dataPath)
	if err != nil {
		// Retry once more.
		time.Sleep(500 * time.Millisecond)
		err := saveSecretToDisk(secret, dataPath)
		if err != nil {
			errChan <- err
			// Do not change lastBackedUpIndex
			// since the backup was not successful.
			return
		}
	}

	lastBackedUpIndex[secret.Name] = int(newIndex)
}
