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
	"context"
	"github.com/pkg/errors"
	"strings"
	"time"

	apiV1 "k8s.io/api/core/v1"
	kErrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/backoff"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// saveSecretToKubernetes saves a given SecretStored entity to a Kubernetes cluster.
// It handles the process of configuring a Kubernetes client, determining the
// appropriate  secret name, and either creating or updating the secret in the
// specified namespace.
//
// The secret name is derived from the input secret entity. If the secret’s name
// has a specific prefix  (determined by env.StoreWorkloadAsK8sSecretPrefix), that
// prefix is removed. Otherwise, a default prefix (from env.SecretNamePrefixForSafe)
// is appended to the secret name.
//
// The secret data is prepared by converting the input secret entity into a
// map suitable for Kubernetes. The namespace for the secret is extracted from
// the secret’s metadata.
//
// Parameters:
// - secret: An entity.SecretStored object containing the secret data to be stored.
//
// Returns:
//   - error: An error object that will be non-nil if an error occurs at any step of
//     the process.
//
// Example:
// err := saveSecretToKubernetes(mySecret)
//
//	if err != nil {
//	    log.Fatalf("Failed to save secret: %v", err)
//	}
//
// Note: This function assumes it is running within a Kubernetes cluster as it
// uses InClusterConfig to generate the Kubernetes client configuration.
func saveSecretToKubernetes(secret entity.SecretStored) error {
	config, err := rest.InClusterConfig()
	if err != nil {
		return errors.Wrap(err, "could not create client config")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "could not create client")
	}

	k8sSecretName := env.SecretNamePrefixForSafe() + secret.Name

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
			namespaces[i] = "default"
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
				return errors.Wrap(err, "error creating the secret")
			}

			continue
		}

		// Secret is found in the cluster.

		// Update the Secret in the cluster
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
			return errors.Wrap(err, "error updating the secret")
		}
	}

	return nil
}

func persistK8s(secret entity.SecretStored, errChan chan<- error) {
	cid := secret.Meta.CorrelationId

	log.TraceLn(&cid, "persistK8s: Will persist k8s secret.")

	if len(secret.Values) == 0 {
		secret.Values = append(secret.Values, InitialSecretValue)
	}

	// Defensive coding:
	// secret’s value is never empty because when the value is set to an
	// empty secret, it is scheduled for deletion and not persisted to the
	// file system or the cluster. However, it that happens, we would at least
	// want an indicator that it happened.
	if secret.Values[0] == "" {
		secret.Values[0] = InitialSecretValue
	}

	log.TraceLn(&cid, "persistK8s: Will try saving secret to k8s.")
	err := saveSecretToKubernetes(secret)
	log.TraceLn(&cid, "persistK8s: should have saved secret to k8s.")
	if err != nil {
		log.TraceLn(&cid, "persistK8s: Got error while trying to save, will retry.")
		// Retry once more.
		time.Sleep(500 * time.Millisecond)
		log.TraceLn(&cid, "persistK8s: Retrying saving secret to k8s.")
		err := saveSecretToKubernetes(secret)
		log.TraceLn(&cid, "persistK8s: Should have saved secret.")
		if err != nil {
			log.TraceLn(&cid, "persistK8s: still error, pushing the error to errchan")
			errChan <- err
		}
	}
}
