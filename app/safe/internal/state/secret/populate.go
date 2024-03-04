/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package secret

import (
	"os"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/io/persistence"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/stats"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// Secrets is where all the secrets are stored.
var Secrets sync.Map

var secretsPopulated = false
var secretsPopulatedLock = sync.Mutex{}

// SecretsPopulated returns a boolean indicating whether the secrets have been
// populated.
func SecretsPopulated() bool {
	secretsPopulatedLock.Lock()
	defer secretsPopulatedLock.Unlock()

	return secretsPopulated
}

// PopulateSecrets scans the designated secrets storage directory on disk, reading each secret file
// that is not marked as a backup, and loads the secrets into a global store if they have not already
// been loaded. This ensures that the application's current session has access to all persisted secrets.
// It uses a locking mechanism to prevent concurrent execution and ensure data consistency.
//
// Parameters:
//   - cid (string): A correlation ID that is used for logging purposes, allowing for the tracing of
//     the populate operation through logs.
//
// Returns:
//   - error: If an error occurs during the directory reading or secret reading process, it returns
//     an error wrapped with context about the failure point. If no errors occur, it returns nil to
//     indicate successful completion.
func PopulateSecrets(cid string, rootKeyTriplet []string) error {
	log.TraceLn(&cid, "populateSecrets: populating secrets...")
	secretsPopulatedLock.Lock()
	defer secretsPopulatedLock.Unlock()

	root := env.DataPathForSafe()
	files, err := os.ReadDir(root)
	if err != nil {
		return errors.Wrap(err, "populateSecrets: problem reading secrets directory")
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fn := file.Name()
		if strings.HasSuffix(fn, ".backup") {
			continue
		}

		key := strings.Replace(fn, ".age", "", 1)

		_, exists := Secrets.Load(key)
		if exists {
			continue
		}

		secretOnDisk, err := persistence.ReadFromDisk(key, rootKeyTriplet)
		if err != nil {
			log.ErrorLn(&cid, "populateSecrets: problem reading secret from disk:", err.Error())
			continue
		}
		if secretOnDisk != nil {
			stats.CurrentState.Increment(key, Secrets.Load)
			Secrets.Store(key, *secretOnDisk)
		}
	}

	secretsPopulated = true
	log.TraceLn(&cid, "populateSecrets: secrets populated.")
	return nil
}
