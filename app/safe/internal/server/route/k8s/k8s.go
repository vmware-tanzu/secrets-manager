/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package k8s

import (
	"context"
	"github.com/pkg/errors"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/backoff"
	apiV1 "k8s.io/api/core/v1"
	kErrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func MarkInitializationSecretAsCompleted() error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return errors.Wrap(err, "could not create client config")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "could not create client")
	}

	secretName := "vsecm-sentinel-init-tombstone"
	namespace := "vsecm-system"

	// First, try to get the existing secret
	_, err = clientset.CoreV1().Secrets(namespace).Get(
		context.Background(), secretName, metaV1.GetOptions{})

	if kErrors.IsNotFound(err) {
		return errors.New("initialization secret is expected to have existed.")
	}

	// Update the Secret in the cluster
	err = backoff.RetryFixed(
		namespace,
		func() error {
			_, err = clientset.CoreV1().Secrets(namespace).Update(
				context.Background(),
				&apiV1.Secret{
					TypeMeta: metaV1.TypeMeta{
						Kind:       "Secret",
						APIVersion: "v1",
					},
					ObjectMeta: metaV1.ObjectMeta{
						Name:      secretName,
						Namespace: namespace,
					},
					Data: map[string][]byte{
						"init": []byte("complete"),
					},
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
		return errors.Wrap(err, "error updating the secret")
	}

	return nil
}
