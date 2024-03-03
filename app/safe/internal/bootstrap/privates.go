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
	"github.com/vmware-tanzu/secrets-manager/core/backoff"

	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pkg/errors"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func persistKeys(privateKey, publicKey, aesSeed string) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return errors.Wrap(err, "Error creating client config")
	}

	k8sApi, err := kubernetes.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "Error creating k8sApi")
	}

	data := make(map[string][]byte)
	keysCombined := crypto.CombineKeys(privateKey, publicKey, aesSeed)
	data["KEY_TXT"] = ([]byte)(keysCombined)

	// Serialize the Secret's configuration to JSON
	secretConfigJSON, err := json.Marshal(v1.Secret{
		TypeMeta: metaV1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metaV1.ObjectMeta{
			Name:      env.RootKeySecretNameForSafe(),
			Namespace: env.SystemNamespace(),
		},
		Data: data,
	})
	if err != nil {
		return errors.Wrap(err, "Error marshalling the secret")
	}

	// Update the Secret in the cluster
	err = backoff.RetryFixed(
		env.SystemNamespace(),
		func() error {
			_, err = k8sApi.CoreV1().Secrets(env.SystemNamespace()).Update(
				context.Background(),
				&v1.Secret{
					TypeMeta: metaV1.TypeMeta{
						Kind:       "Secret",
						APIVersion: "v1",
					},
					ObjectMeta: metaV1.ObjectMeta{
						Name:      env.RootKeySecretNameForSafe(),
						Namespace: env.SystemNamespace(),
						Annotations: map[string]string{
							"kubectl.kubernetes.io/last-applied-configuration": string(secretConfigJSON),
						},
					},
					Data: data,
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
		return errors.Wrap(err, "Error creating the secret")
	}

	state.SetRootKey(keysCombined)

	return nil
}
