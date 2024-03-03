/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package state

import (
	"bytes"
	"encoding/base64"

	"strings"
	"sync"
	"time"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/io/crypto"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/io/persistence"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/secret"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/secret/queue/deletion"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/secret/queue/insertion"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/stats"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

const InitialSecretValue = `{"empty":true}`
const BlankRootKeyValue = "{}"

// RootKey is the key used for encryption, decryption, backup, and restore.
var RootKey = ""
var RootKeyLock sync.RWMutex

// Initialize starts two goroutines: one to process the secret queue and
// another to process the Kubernetes secret queue. These goroutines are
// responsible for handling queued secrets and persisting them to disk.
func Initialize() {
	go insertion.ProcessSecretQueue()
	go insertion.ProcessK8sSecretQueue()
	go deletion.ProcessSecretDeleteQueue()
	go deletion.ProcessK8sSecretDeleteQueue()
}

// SetRootKey sets the age key to be used for encryption and decryption.
func SetRootKey(k string) {
	id := "AEGSAK"

	RootKeyLock.Lock()
	defer RootKeyLock.Unlock()

	if RootKey != "" {
		log.WarnLn(&id, "Root key already set")
		return
	}
	RootKey = k
}

// RootKeySet returns true if the root key has been set.
func RootKeySet() bool {
	RootKeyLock.RLock()
	defer RootKeyLock.RUnlock()

	return RootKey != ""
}

// EncryptValue takes a string value and returns an encrypted and base64-encoded
// representation of the input value. If the encryption process encounters any
// error, it will return an empty string and the corresponding error.
func EncryptValue(value string) (string, error) {
	var out bytes.Buffer

	fipsMode := env.FipsCompliantModeForSafe()

	if fipsMode {
		err := crypto.EncryptToWriterAes(&out, value)
		if err != nil {
			return "", err
		}
	} else {
		err := crypto.EncryptToWriterAge(&out, value)
		if err != nil {
			return "", err
		}
	}

	base64Str := base64.StdEncoding.EncodeToString(out.Bytes())

	return base64Str, nil
}

// DecryptValue takes a base64-encoded and encrypted string value and returns
// the original, decrypted string. If the decryption process encounters any
// error, it will return an empty string and the corresponding error.
func DecryptValue(value string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	if env.FipsCompliantModeForSafe() {
		decrypted, err := crypto.DecryptBytesAes(decoded)
		if err != nil {
			return "", err
		}
		return string(decrypted), nil
	}

	decrypted, err := crypto.DecryptBytes(decoded)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

// AllSecrets returns a slice of entity.Secret containing all secrets
// currently stored. If no secrets are found, an empty slice is
// returned.
func AllSecrets(cid string) []entity.Secret {
	var result []entity.Secret

	// Check existing stored secrets files.
	// If VSecM pod is evicted and revived, it will not have knowledge about
	// the secret it has. This loop helps it re-populate its cache.
	if !secret.SecretsPopulated() {
		err := secret.PopulateSecrets(cid)
		if err != nil {
			log.WarnLn(&cid, "Failed to populate secrets from disk", err.Error())
		}
	}

	// Range over all existing secrets.
	secret.Secrets.Range(func(key any, value any) bool {
		v := value.(entity.SecretStored)

		result = append(result, entity.Secret{
			Name:         v.Name,
			Created:      entity.JsonTime(v.Created),
			Updated:      entity.JsonTime(v.Updated),
			NotBefore:    entity.JsonTime(v.NotBefore),
			ExpiresAfter: entity.JsonTime(v.ExpiresAfter),
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
	if !secret.SecretsPopulated() {
		err := secret.PopulateSecrets(cid)
		if err != nil {
			log.WarnLn(&cid, "Failed to populate secrets from disk", err.Error())
		}
	}

	// Range over all existing secrets.
	secret.Secrets.Range(func(key any, value any) bool {
		v := value.(entity.SecretStored)

		var vals []string
		for _, val := range v.Values {
			ve, _ := EncryptValue(val)
			vals = append(vals, ve)
		}

		result = append(result, entity.SecretEncrypted{
			Name:           v.Name,
			EncryptedValue: vals,
			Created:        entity.JsonTime(v.Created),
			Updated:        entity.JsonTime(v.Updated),
			NotBefore:      entity.JsonTime(v.NotBefore),
			ExpiresAfter:   entity.JsonTime(v.ExpiresAfter),
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

// UpsertSecret takes an entity.SecretStored object and inserts it into
// the in-memory store if it doesn't exist, or updates it if it does. It also
// handles updating the backing store and Kubernetes secrets if necessary.
// If appendValue is true, the new value will be appended to the existing values,
// otherwise it will replace the existing values.
func UpsertSecret(secretStored entity.SecretStored, appendValue bool) {
	cid := secretStored.Meta.CorrelationId

	vs := secretStored.Values

	if len(vs) == 0 {
		log.InfoLn(&cid, "UpsertSecret: nothing to upsert. exiting.", "len(vs)", len(vs))
		return
	}

	var nonEmptyValues []string
	for _, value := range secretStored.Values {
		if value != "" {
			nonEmptyValues = append(nonEmptyValues, value)
		}
	}

	if nonEmptyValues == nil {
		log.InfoLn(&cid, "UpsertSecret: nothing to upsert. exiting.", "len(vs)", len(vs))
		return
	}

	secretStored.Values = nonEmptyValues

	s, exists := secret.Secrets.Load(secretStored.Name)
	now := time.Now()
	if exists {
		log.TraceLn(&cid, "UpsertSecret: Secret exists. Will update.")

		ss := s.(entity.SecretStored)
		secretStored.Created = ss.Created

		if appendValue {
			log.TraceLn(&cid, "UpsertSecret: Will append value.")

			for _, v := range ss.Values {
				if contains(secretStored.Values, v) {
					continue
				}
				if len(v) == 0 {
					continue
				}
				secretStored.Values = append(secretStored.Values, v)
			}
		}
	} else {
		secretStored.Created = now
	}
	secretStored.Updated = now

	log.InfoLn(&cid, "UpsertSecret:",
		"created", secretStored.Created, "updated", secretStored.Updated,
		"name", secretStored.Name, "len(vs)", len(vs),
	)

	log.TraceLn(&cid, "UpsertSecret: Will parse secret.")

	parsedStr, err := secretStored.Parse()
	if err != nil {
		log.InfoLn(&cid,
			"UpsertSecret: Error parsing secret. Will use fallback value.", err.Error())
	}

	log.TraceLn(&cid, "UpsertSecret: Parsed secret. Will set transformed value.")

	secretStored.ValueTransformed = parsedStr
	stats.CurrentState.Increment(secretStored.Name)
	secret.Secrets.Store(secretStored.Name, secretStored)

	store := secretStored.Meta.BackingStore

	switch store {
	case entity.File:
		log.TraceLn(
			&cid, "UpsertSecret: Will push secret. len",
			len(insertion.SecretUpsertQueue),
			"cap", cap(insertion.SecretUpsertQueue))
		insertion.SecretUpsertQueue <- secretStored
		log.TraceLn(
			&cid, "UpsertSecret: Pushed secret. len",
			len(insertion.SecretUpsertQueue), "cap",
			cap(insertion.SecretUpsertQueue))
	}

	useK8sSecrets := secretStored.Meta.UseKubernetesSecret

	// If useK8sSecrets is not set, use the value from the environment.
	// The environment value defaults to false, too, if not set.
	// If the "name" of the secret has the prefix "k8s:", then store it as a
	// Kubernetes secret too.
	if useK8sSecrets ||
		env.UseKubernetesSecretsModeForSafe() ||
		strings.HasPrefix(secretStored.Name, env.StoreWorkloadAsK8sSecretPrefix()) {
		log.TraceLn(
			&cid,
			"UpsertSecret: will push Kubernetes secret. len",
			len(insertion.K8sSecretUpsertQueue),
			"cap", cap(insertion.K8sSecretUpsertQueue),
		)
		insertion.K8sSecretUpsertQueue <- secretStored
		log.TraceLn(
			&cid,
			"UpsertSecret: pushed Kubernetes secret. len",
			len(insertion.K8sSecretUpsertQueue),
			"cap", cap(insertion.K8sSecretUpsertQueue),
		)
	}
}

// DeleteSecret orchestrates the deletion of a specified secret from both the
// application's internal cache and its persisted storage locations, which may
// include local filesystem and Kubernetes secrets. The deletion process is
// contingent upon the secret's metadata, specifically its backing store and
// whether it is used as a Kubernetes secret.
//
// Parameters:
//   - secretToDelete (entity.SecretStored): The secret entity marked for deletion,
//     containing necessary metadata such as the name of the secret, its correlation
//     ID for logging, and metadata specifying where and how the secret is stored.
func DeleteSecret(secretToDelete entity.SecretStored) {
	cid := secretToDelete.Meta.CorrelationId

	s, exists := secret.Secrets.Load(secretToDelete.Name)
	if !exists {
		log.WarnLn(&cid,
			"DeleteSecret: Secret does not exist. Cannot delete.",
			secretToDelete.Name)
		return
	}

	ss := s.(entity.SecretStored)

	store := ss.Meta.BackingStore

	switch store {
	case entity.File:
		log.TraceLn(
			&cid, "DeleteSecret: Will delete secret. len",
			len(deletion.SecretDeleteQueue),
			"cap", cap(deletion.SecretDeleteQueue))
		deletion.SecretDeleteQueue <- secretToDelete
		log.TraceLn(
			&cid, "DeleteSecret: Pushed secret to delete. len",
			len(deletion.SecretDeleteQueue), "cap",
			cap(deletion.SecretDeleteQueue))
	}

	useK8sSecrets := secretToDelete.Meta.UseKubernetesSecret

	// If useK8sSecrets is not set, use the value from the environment.
	// The environment value defaults to false, too, if not set.
	if useK8sSecrets || env.UseKubernetesSecretsModeForSafe() {
		log.TraceLn(
			&cid,
			"DeleteSecret: will push Kubernetes secret to delete. len",
			len(deletion.K8sSecretDeleteQueue),
			"cap", cap(deletion.K8sSecretDeleteQueue),
		)
		deletion.K8sSecretDeleteQueue <- secretToDelete
		log.TraceLn(
			&cid,
			"DeleteSecret: pushed Kubernetes secret to delete. len",
			len(deletion.K8sSecretDeleteQueue),
			"cap", cap(deletion.K8sSecretDeleteQueue),
		)
	}

	// Remove the secret from the memory.
	stats.CurrentState.Decrement(secretToDelete.Name)
	secret.Secrets.Delete(secretToDelete.Name)
}

// ReadSecret takes a key string and returns a pointer to an entity.SecretStored
// object if the secret exists in the in-memory store. If the secret is not
// found in memory, it attempts to read it from disk, store it in memory, and
// return it. If the secret is not found on disk, it returns nil.
func ReadSecret(cid string, key string) (*entity.SecretStored, error) {
	log.TraceLn(&cid, "ReadSecret:begin")

	result, secretFoundInMemory := secret.Secrets.Load(key)
	if secretFoundInMemory {
		s := result.(entity.SecretStored)
		log.TraceLn(&cid, "ReadSecret: returning from memory.", "len", len(s.Values))
		return &s, nil
	}

	stored, err := persistence.ReadFromDisk(key)
	if err != nil {
		return nil, err
	}
	if stored == nil {
		return nil, nil
	}

	stats.CurrentState.Increment(stored.Name)
	secret.Secrets.Store(stored.Name, *stored)

	log.TraceLn(&cid, "ReadSecret: returning from disk.", "len", len(stored.Values))
	return stored, nil
}
