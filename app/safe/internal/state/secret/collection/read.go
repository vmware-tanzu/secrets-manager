/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package collection

import (
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/io"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/stats"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	data "github.com/vmware-tanzu/secrets-manager/lib/entity"
)

// SecretByName retrieves a secret by its name.
// This function first checks if the secrets have been populated in the cache.
// If not, it populates the secrets using the root key triplet. It then attempts
// to load the secret by name from the populated cache.
//
// Parameters:
//   - cid  string: A correlation ID used to track the request and associated
//     logging. This ID helps in tracing and debugging operations across
//     different components or services that handle the secret data.
//   - name string: The name of the secret to be retrieved.
//
// Returns:
//   - *entity.Secret: A pointer to the Secret entity if found. The Secret
//     structure includes fields such as Name, Created, Updated, NotBefore, and
//     ExpiresAfter. Each of these timestamp fields is converted from the stored
//     format to a JSON compatible format. Returns nil if no secret with the
//     provided name is found in the cache.
//
// Error Handling:
//   - If there is an error in populating the secrets from the disk (e.g., due
//     to read errors or  data corruption), the function logs a warning message
//     with the correlation ID and the error message but continues execution.
//     This does not halt the function, and it subsequently tries to fetch the
//     secret if already available in the cache.
func SecretByName(cid string, name string) *entity.Secret {
	// Check existing stored secrets files.
	// If VSecM pod is evicted and revived, it will not have knowledge about
	// the secret it has. This loop helps it re-populate its cache.
	if !SecretsPopulated() {
		err := PopulateSecrets(cid)
		if err != nil {
			log.WarnLn(&cid,
				"Failed to populate secrets from disk", err.Error())
		}
	}

	s, ok := Secrets.Load(name)
	if !ok {
		return nil
	}

	v := s.(entity.SecretStored)

	return &entity.Secret{
		Name:         v.Name,
		Created:      data.JsonTime(v.Created),
		Updated:      data.JsonTime(v.Updated),
		NotBefore:    data.JsonTime(v.NotBefore),
		ExpiresAfter: data.JsonTime(v.ExpiresAfter),
	}
}

const keystoneWorkloadId = "vsecm-keystone"

// KeystoneInitialized checks whether the keystone secret is registered.
//
// This is a utility function that depends on the SecretByName function to check
// for the presence of the specific secret. A return value of true indicates
// that the keystone is initialized and ready for use, while false indicates it
// is not.
//
// Parameters:
//   - cid string: A correlation ID used for logging and tracing.
//
// Returns:
//   - bool: True if the keystone secret is present, false otherwise.
func KeystoneInitialized(cid string) bool {
	ks := SecretByName(cid, keystoneWorkloadId)
	return ks != nil
}

// AllSecrets returns a slice of entity.Secret containing all secrets
// currently stored. If no secrets are found, an empty slice is
// returned.
func AllSecrets(cid string) []entity.Secret {
	var result []entity.Secret

	// Check existing stored secrets files.
	// If VSecM pod is evicted and revived, it will not have knowledge about
	// the secret it has. This loop helps it re-populate its cache.
	if !SecretsPopulated() {
		err := PopulateSecrets(cid)
		if err != nil {
			log.WarnLn(&cid,
				"Failed to populate secrets from disk", err.Error())
		}
	}

	// Range over all existing secrets.
	Secrets.Range(func(key any, value any) bool {
		v := value.(entity.SecretStored)

		result = append(result, entity.Secret{
			Name:         v.Name,
			Created:      data.JsonTime(v.Created),
			Updated:      data.JsonTime(v.Updated),
			NotBefore:    data.JsonTime(v.NotBefore),
			ExpiresAfter: data.JsonTime(v.ExpiresAfter),
		})

		return true
	})

	if result == nil {
		return []entity.Secret{}
	}

	return result
}

// AllSecretsEncrypted returns a slice of entity.SecretEncrypted containing all
// secrets  currently stored. If no secrets are found, an empty slice is
// returned.
func AllSecretsEncrypted(cid string) []entity.SecretEncrypted {
	var result []entity.SecretEncrypted

	// Check existing stored secrets files.
	// If VSecM pod is evicted and revived, it will not have knowledge about
	// the secret it has. This loop helps it re-populate its cache.
	if !SecretsPopulated() {
		err := PopulateSecrets(cid)
		if err != nil {
			log.WarnLn(&cid,
				"Failed to populate secrets from disk", err.Error())
		}
	}

	// Range over all existing secrets.
	Secrets.Range(func(key any, value any) bool {
		v := value.(entity.SecretStored)

		var vals []string
		for _, val := range v.Values {
			ve, _ := crypto.EncryptValue(val)
			vals = append(vals, ve)
		}

		result = append(result, entity.SecretEncrypted{
			Name:           v.Name,
			EncryptedValue: vals,
			Created:        data.JsonTime(v.Created),
			Updated:        data.JsonTime(v.Updated),
			NotBefore:      data.JsonTime(v.NotBefore),
			ExpiresAfter:   data.JsonTime(v.ExpiresAfter),
		})

		return true
	})

	if result == nil {
		return []entity.SecretEncrypted{}
	}

	return result
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ReadSecret takes a key string and returns a pointer to an entity.SecretStored
// object if the secret exists in the in-memory store. If the secret is not
// found in memory, it attempts to read it from disk, store it in memory, and
// return it. If the secret is not found on disk, it returns nil.
func ReadSecret(cid string, key string) (*entity.SecretStored, error) {
	log.TraceLn(&cid, "ReadSecret:begin")

	result, secretFoundInMemory := Secrets.Load(key)
	if secretFoundInMemory {
		s := result.(entity.SecretStored)
		log.TraceLn(&cid,
			"ReadSecret: returning from memory.", "len", len(s.Values))
		return &s, nil
	}

	stored, err := io.ReadFromDisk(key)

	if err != nil {
		return nil, err
	}
	if stored == nil {
		return nil, nil
	}

	stats.CurrentState.Increment(stored.Name, Secrets.Load)
	Secrets.Store(stored.Name, *stored)

	log.TraceLn(&cid,
		"ReadSecret: returning from disk.", "len", len(stored.Values))
	return stored, nil
}
