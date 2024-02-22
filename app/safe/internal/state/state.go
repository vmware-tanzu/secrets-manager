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
	"bytes"
	"encoding/base64"
	"strings"
	"sync"
	"time"

	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

const InitialSecretValue = `{"empty":true}`
const BlankAgeKeyValue = "{}"

// masterKey is set only once during initialization; we don’t need to lock
// access to it.
var masterKey = ""

// Initialize starts two goroutines: one to process the secret queue and
// another to process the Kubernetes secret queue. These goroutines are
// responsible for handling queued secrets and persisting them to disk.
func Initialize() {
	go processSecretQueue()
	go processK8sSecretQueue()
	go processSecretDeleteQueue()
	go processK8sSecretDeleteQueue()
}

var masterKeyLock sync.RWMutex

// SetMasterKey sets the age key to be used for encryption and decryption.
func SetMasterKey(k string) {
	id := "AEGSAK"

	masterKeyLock.Lock()
	defer masterKeyLock.Unlock()

	if masterKey != "" {
		log.WarnLn(&id, "master key already set")
		return
	}
	masterKey = k
}

// MasterKeySet returns true if the master key has been set.
func MasterKeySet() bool {
	masterKeyLock.RLock()
	defer masterKeyLock.RUnlock()

	return masterKey != ""
}

// EncryptValue takes a string value and returns an encrypted and base64-encoded
// representation of the input value. If the encryption process encounters any
// error, it will return an empty string and the corresponding error.
func EncryptValue(value string) (string, error) {
	var out bytes.Buffer

	fipsMode := env.SafeFipsCompliant()

	if fipsMode {
		err := encryptToWriterAes(&out, value)
		if err != nil {
			return "", err
		}
	} else {
		err := encryptToWriterAge(&out, value)
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

	if env.SafeFipsCompliant() {
		decrypted, err := decryptBytesAes(decoded)
		if err != nil {
			return "", err
		}
		return string(decrypted), nil
	}

	decrypted, err := decryptBytes(decoded)
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
	if !secretsPopulated {
		err := populateSecrets(cid)
		if err != nil {
			log.WarnLn(&cid, "Failed to populate secrets from disk", err.Error())
		}
	}

	// Range over all existing secrets.
	secrets.Range(func(key any, value any) bool {
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
	if !secretsPopulated {
		err := populateSecrets(cid)
		if err != nil {
			log.WarnLn(&cid, "Failed to populate secrets from disk", err.Error())
		}
	}

	// Range over all existing secrets.
	secrets.Range(func(key any, value any) bool {
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
func UpsertSecret(secret entity.SecretStored, appendValue bool) {
	cid := secret.Meta.CorrelationId

	vs := secret.Values

	if len(vs) == 0 {
		log.InfoLn(&cid, "UpsertSecret: nothing to upsert. exiting.", "len(vs)", len(vs))
		return
	}

	var nonEmptyValues []string
	for _, value := range secret.Values {
		if value != "" {
			nonEmptyValues = append(nonEmptyValues, value)
		}
	}

	if nonEmptyValues == nil {
		log.InfoLn(&cid, "UpsertSecret: nothing to upsert. exiting.", "len(vs)", len(vs))
		return
	}

	secret.Values = nonEmptyValues

	s, exists := secrets.Load(secret.Name)
	now := time.Now()
	if exists {
		log.TraceLn(&cid, "UpsertSecret: Secret exists. Will update.")

		ss := s.(entity.SecretStored)
		secret.Created = ss.Created

		if appendValue {
			log.TraceLn(&cid, "UpsertSecret: Will append value.")

			for _, v := range ss.Values {
				if contains(secret.Values, v) {
					continue
				}
				if len(v) == 0 {
					continue
				}
				secret.Values = append(secret.Values, v)
			}
		}
	} else {
		secret.Created = now
	}
	secret.Updated = now

	log.InfoLn(&cid, "UpsertSecret:",
		"created", secret.Created, "updated", secret.Updated, "name", secret.Name,
		"len(vs)", len(vs),
	)

	log.TraceLn(&cid, "UpsertSecret: Will parse secret.")

	parsedStr, err := secret.Parse()
	if err != nil {
		log.InfoLn(&cid,
			"UpsertSecret: Error parsing secret. Will use fallback value.", err.Error())
	}

	log.TraceLn(&cid, "UpsertSecret: Parsed secret. Will set transformed value.")

	secret.ValueTransformed = parsedStr
	currentState.Increment(secret.Name)
	secrets.Store(secret.Name, secret)

	store := secret.Meta.BackingStore

	switch store {
	case entity.File:
		log.TraceLn(
			&cid, "UpsertSecret: Will push secret. len", len(secretQueue), "cap", cap(secretQueue))
		secretQueue <- secret
		log.TraceLn(
			&cid, "UpsertSecret: Pushed secret. len", len(secretQueue), "cap", cap(secretQueue))
	}

	useK8sSecrets := secret.Meta.UseKubernetesSecret

	// If useK8sSecrets is not set, use the value from the environment.
	// The environment value defaults to false, too, if not set.
	// If the "name" of the secret has the prefix "k8s:", then store it as a
	// Kubernetes secret too.
	if useK8sSecrets ||
		env.SafeUseKubernetesSecrets() ||
		strings.HasPrefix(secret.Name, env.StoreWorkloadAsK8sSecretPrefix()) {
		log.TraceLn(
			&cid,
			"UpsertSecret: will push Kubernetes secret. len", len(k8sSecretQueue),
			"cap", cap(k8sSecretQueue),
		)
		k8sSecretQueue <- secret
		log.TraceLn(
			&cid,
			"UpsertSecret: pushed Kubernetes secret. len", len(k8sSecretQueue),
			"cap", cap(k8sSecretQueue),
		)
	}
}

func DeleteSecret(secret entity.SecretStored) {
	cid := secret.Meta.CorrelationId

	s, exists := secrets.Load(secret.Name)
	if !exists {
		log.WarnLn(&cid, "DeleteSecret: Secret does not exist. Cannot delete.", secret.Name)
		return
	}

	ss := s.(entity.SecretStored)

	store := ss.Meta.BackingStore

	switch store {
	case entity.File:
		log.TraceLn(
			&cid, "DeleteSecret: Will delete secret. len", len(secretDeleteQueue), "cap", cap(secretDeleteQueue))
		secretDeleteQueue <- secret
		log.TraceLn(
			&cid, "DeleteSecret: Pushed secret to delete. len", len(secretDeleteQueue), "cap", cap(secretDeleteQueue))
	}

	useK8sSecrets := secret.Meta.UseKubernetesSecret

	// If useK8sSecrets is not set, use the value from the environment.
	// The environment value defaults to false, too, if not set.
	if useK8sSecrets || env.SafeUseKubernetesSecrets() {
		log.TraceLn(
			&cid,
			"DeleteSecret: will push Kubernetes secret to delete. len", len(k8sSecretDeleteQueue),
			"cap", cap(k8sSecretDeleteQueue),
		)
		k8sSecretDeleteQueue <- secret
		log.TraceLn(
			&cid,
			"DeleteSecret: pushed Kubernetes secret to delete. len", len(k8sSecretDeleteQueue),
			"cap", cap(k8sSecretDeleteQueue),
		)
	}

	// Remove the secret from the memory.
	currentState.Decrement(secret.Name)
	secrets.Delete(secret.Name)
}

// ReadSecret takes a key string and returns a pointer to an entity.SecretStored
// object if the secret exists in the in-memory store. If the secret is not
// found in memory, it attempts to read it from disk, store it in memory, and
// return it. If the secret is not found on disk, it returns nil.
func ReadSecret(cid string, key string) (*entity.SecretStored, error) {
	log.TraceLn(&cid, "ReadSecret:begin")

	result, ok := secrets.Load(key)
	if !ok {
		stored, err := readFromDisk(key)
		if err != nil {
			return nil, err
		}

		if stored == nil {
			return nil, nil
		}
		currentState.Increment(stored.Name)
		secrets.Store(stored.Name, *stored)
		secretQueue <- *stored

		log.TraceLn(&cid, "ReadSecret: returning from disk.", "len", len(stored.Values))
		return stored, nil
	}

	s := result.(entity.SecretStored)
	log.TraceLn(&cid, "ReadSecret: returning from memory.", "len", len(s.Values))
	return &s, nil
}
