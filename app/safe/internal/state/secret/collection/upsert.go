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
	"strings"
	"time"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/secret/queue/insertion"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/stats"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// UpsertSecret takes an entity.SecretStored object and inserts it into
// the in-memory store if it doesn't exist, or updates it if it does. It also
// handles updating the backing store and Kubernetes secrets if necessary.
// If appendValue is true, the new value will be appended to the existing
// values, otherwise it will replace the existing values.
func UpsertSecret(secretStored entity.SecretStored) {
	cid := secretStored.Meta.CorrelationId

	val := secretStored.Value

	if len(val) == 0 {
		log.InfoLn(&cid,
			"UpsertSecret: nothing to upsert. exiting.", "len(vs)", len(val))
		return
	}

	s, exists := Secrets.Load(secretStored.Name)
	now := time.Now()
	if exists {
		log.TraceLn(&cid, "UpsertSecret: Secret exists. Will update.")

		ss := s.(entity.SecretStored)
		secretStored.Created = ss.Created
	} else {
		secretStored.Created = now
	}
	secretStored.Updated = now

	log.InfoLn(&cid, "UpsertSecret:",
		"created", secretStored.Created, "updated", secretStored.Updated,
		"name", secretStored.Name, "len(vs)", len(val),
	)

	log.TraceLn(&cid, "UpsertSecret: Will parse secret.")

	parsedStr, err := secretStored.Parse()
	if err != nil {
		log.InfoLn(&cid,
			"UpsertSecret: Error parsing secret. Will use fallback value.",
			err.Error(),
		)
	}

	log.TraceLn(&cid,
		"UpsertSecret: Parsed secret. Will set transformed value.")

	secretStored.ValueTransformed = parsedStr
	stats.CurrentState.Increment(secretStored.Name, Secrets.Load)
	Secrets.Store(secretStored.Name, secretStored)

	log.TraceLn(
		&cid, "UpsertSecret: Will push secret. len",
		len(insertion.SecretUpsertQueue),
		"cap", cap(insertion.SecretUpsertQueue))

	// The insertion queue will add the secret to the backing store.
	// The backing store is determined by the env.BackingStoreForSafe()
	// function.
	insertion.SecretUpsertQueue <- secretStored

	log.TraceLn(
		&cid, "UpsertSecret: Pushed secret. len",
		len(insertion.SecretUpsertQueue), "cap",
		cap(insertion.SecretUpsertQueue))

	// A "raw" secret cannot be queried by regular workloads, you will need a
	// special Kubernetes Operator to access it.
	if strings.HasPrefix(secretStored.Name, env.RawSecretPrefix()) {
		log.TraceLn(&cid,
			"UpsertSecret: the secret will not be associated with a workload.",
		)

		return
	}

	// If the "name" of the secret has the prefix "k8s:", then store it as a
	// Kubernetes secret too.
	if strings.HasPrefix(secretStored.Name,
		env.StoreWorkloadAsK8sSecretPrefix()) {
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
