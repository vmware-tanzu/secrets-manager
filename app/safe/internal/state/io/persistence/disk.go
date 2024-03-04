/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package persistence

import (
	"encoding/json"
	"math"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/io/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

var lastBackedUpIndex = make(map[string]int)

func saveSecretToDisk(secret entity.SecretStored, dataPath string, rootKeyTriplet []string) error {
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

	if env.FipsCompliantModeForSafe() {
		return crypto.EncryptToWriterAes(file, string(data), rootKeyTriplet)
	}

	return crypto.EncryptToWriterAge(file, string(data), rootKeyTriplet)
}

// PersistToDisk saves a given secret to disk and also creates a backup copy of the
// secret. The function is designed to enhance data durability through retries and
// backup management based on environmental configurations.
//
// Parameters:
//   - secret (entity.SecretStored): The secret to be saved, which is a structured
//     entity containing the secret's name and possibly other metadata or the secret
//     data itself.
//   - errChan (chan<- error): A channel through which errors are reported. This
//     channel allows the function to operate asynchronously, notifying the caller
//     of any issues in the process of persisting the secret.
func PersistToDisk(secret entity.SecretStored, rootKeyTriplet []string, errChan chan<- error) {
	backupCount := env.SecretBackupCountForSafe()

	k1, k2, k3 := rootKeyTriplet[0], rootKeyTriplet[1], rootKeyTriplet[2]
	rkt := []string{k1, k2, k3}

	// Save the secret
	dataPath := path.Join(env.DataPathForSafe(), secret.Name+".age")

	err := saveSecretToDisk(secret, dataPath, rkt)
	if err != nil {
		// Retry once more.
		time.Sleep(500 * time.Millisecond)
		err := saveSecretToDisk(secret, dataPath, rkt)
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
		env.DataPathForSafe(),
		secret.Name+"-"+strconv.Itoa(int(newIndex))+"-"+".age.backup",
	)

	err = saveSecretToDisk(secret, dataPath, rkt)
	if err != nil {
		// Retry once more.
		time.Sleep(500 * time.Millisecond)
		err := saveSecretToDisk(secret, dataPath, rkt)
		if err != nil {
			errChan <- err
			// Do not change lastBackedUpIndex
			// since the backup was not successful.
			return
		}
	}

	lastBackedUpIndex[secret.Name] = int(newIndex)
}
