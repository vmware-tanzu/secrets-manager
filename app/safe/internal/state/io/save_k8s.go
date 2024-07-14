/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package io

import (
	"context"
	"errors"
	"strings"

	apiV1 "k8s.io/api/core/v1"
	kErrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	ec "github.com/vmware-tanzu/secrets-manager/core/constants/env"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/lib/backoff"
)

// saveSecretToKubernetes saves a given SecretStored entity to a Kubernetes
// cluster. It handles the process of configuring a Kubernetes client,
// determining the appropriate  secret name, and either creating or updating the
// secret in the specified namespace.
//
// The secret data is prepared by converting the input secret entity into a
// map suitable for Kubernetes. The namespace for the secret is extracted from
// the secret's metadata.
//
// Parameters:
//   - secret: An entity.SecretStored object containing the secret data to be
//     stored.
//
// Returns:
//   - error: An error object that will be non-nil if an error occurs at any
//     step of the process.
//
// Note: This function assumes it is running within a Kubernetes cluster as it
// uses InClusterConfig to generate the Kubernetes client configuration.
func saveSecretToKubernetes(secret entity.SecretStored) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return errors.Join(
			err,
			errors.New("could not create client config"),
		)
	}

	// If the secret does not have the k8s: prefix, then it is not a k8s secret;
	// do not save it in the cluster.
	if !strings.HasPrefix(secret.Name, env.StoreWorkloadAsK8sSecretPrefix()) {
		return errors.New("secret does not have k8s: prefix")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return errors.Join(
			err,
			errors.New("could not create client"),
		)
	}

	k8sSecretName := secret.Name

	// If the secret has k8s: prefix, then do not append a prefix; use the name
	// as is.
	if strings.HasPrefix(secret.Name, env.StoreWorkloadAsK8sSecretPrefix()) {
		k8sSecretName = strings.TrimPrefix(
			secret.Name, env.StoreWorkloadAsK8sSecretPrefix(),
		)
	}

	// Transform the data if there is a transformation defined.
	data := secret.ToMapForK8s()
	namespaces := secret.Meta.Namespaces

	for i, ns := range namespaces {
		if ns == "" {
			namespaces[i] = string(ec.Default)
		}

		// First, try to get the existing secret
		_, err = clientset.CoreV1().Secrets(ns).Get(
			context.Background(), k8sSecretName, metaV1.GetOptions{})

		if kErrors.IsNotFound(err) {
			// Create the Secret in the cluster with a backoff.
			err = backoff.RetryFixed(
				ns,
				func() error {
					// Create the Secret in the cluster
					_, err = clientset.CoreV1().Secrets(ns).Create(
						context.Background(),
						&apiV1.Secret{
							TypeMeta: metaV1.TypeMeta{
								Kind:       "Secret",
								APIVersion: "v1",
							},
							ObjectMeta: metaV1.ObjectMeta{
								Name:      k8sSecretName,
								Namespace: ns,
								Labels: map[string]string{
									"app.kubernetes.io/operated-by": "vsecm",
								},
							},
							Data: data,
						},
						metaV1.CreateOptions{
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

			continue
		}

		// Secret is found in the cluster.

		// Update the Secret in the cluster:
		err = backoff.RetryFixed(
			ns,
			func() error {
				_, err = clientset.CoreV1().Secrets(ns).Update(
					context.Background(),
					&apiV1.Secret{
						TypeMeta: metaV1.TypeMeta{
							Kind:       "Secret",
							APIVersion: "v1",
						},
						ObjectMeta: metaV1.ObjectMeta{
							Name:      k8sSecretName,
							Namespace: ns,
							Labels: map[string]string{
								"app.kubernetes.io/operated-by": "vsecm",
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
			return errors.Join(
				err,
				errors.New("error updating the secret"),
			)
		}
	}

	return nil
}
