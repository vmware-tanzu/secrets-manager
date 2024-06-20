/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package bootstrap

import (
	"context"
	"encoding/json"
	"errors"

	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	e "github.com/vmware-tanzu/secrets-manager/core/constants/env"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/lib/backoff"
)

// PersistRootKeysToRootKeyBackingStore persists the root keys to the
// configured backing store. This is useful to restore VSecM Safe back to
// operation if it crashes or gets temporarily evicted by the scheduler.
//
// If the persist operation succeed, it updates the root keys stored in the
// memory too.
//
// This function is typically called during the first bootstrapping of
// VSecM Safe when there are no keys that have been registered yet.
//
// Note that changing the root key without backing up the existing one means
// the secrets backed up with the old key will be impossible to decrypt.
func PersistRootKeysToRootKeyBackingStore(rkt data.RootKeyCollection) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return errors.Join(
			err,
			errors.New("error creating client config"),
		)
	}

	k8sApi, err := kubernetes.NewForConfig(config)
	if err != nil {
		return errors.Join(
			err,
			errors.New("error creating k8sApi"),
		)
	}

	dd := make(map[string][]byte)
	keysCombined := rkt.Combine()
	dd[string(e.RootKeyText)] = ([]byte)(keysCombined)

	// Serialize the Secret's configuration to JSON
	secretConfigJSON, err := json.Marshal(v1.Secret{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      env.RootKeySecretNameForSafe(),
			Namespace: env.NamespaceForVSecMSystem(),
		},
		Data: dd,
	})
	if err != nil {
		return errors.Join(
			err,
			errors.New("error marshalling the secret"),
		)
	}

	// Update the Secret in the cluster
	err = backoff.RetryFixed(
		env.NamespaceForVSecMSystem(),
		func() error {
			_, err = k8sApi.CoreV1().Secrets(
				env.NamespaceForVSecMSystem()).Update(
				context.Background(),
				&v1.Secret{
					TypeMeta: metaV1.TypeMeta{
						Kind:       "Secret",
						APIVersion: "v1",
					},
					ObjectMeta: metaV1.ObjectMeta{
						Name:      env.RootKeySecretNameForSafe(),
						Namespace: env.NamespaceForVSecMSystem(),
						Annotations: map[string]string{
							"kubectl.kubernetes.io/last-applied-configuration": string(secretConfigJSON),
						},
					},
					Data: dd,
				},
				metaV1.UpdateOptions{
					TypeMeta: metaV1.TypeMeta{
						Kind:       "Secret",
						APIVersion: "v1",
					},
				},
			)
			return err
		},
	)

	if err != nil {
		return errors.Join(
			err,
			errors.New("error creating the secret"),
		)
	}

	crypto.SetRootKeyInMemory(keysCombined)

	return nil
}
