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
	"strings"

	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/backoff"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	kErrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// The secrets put here are synced with their Kubernetes Secret counterparts.
var k8sSecretDeleteQueue = make(chan entity.SecretStored, env.SafeK8sSecretDeleteBufferSize())

func processK8sSecretDeleteQueue() {
	errChan := make(chan error)
	id := "AEGIHK8D"

	if env.SafeRemoveLinkedK8sSecrets() {
		go func() {
			for e := range errChan {
				// If the `deleteSecretFromKubernetes` operation spews out an error, log it.
				log.ErrorLn(&id, "processK8sSecretDeleteQueue: error deleting secret:", e.Error())
			}
		}()

		for secret := range k8sSecretDeleteQueue {
			go deleteSecretFromKubernetes(secret, errChan)
		}
	} else {
		log.InfoLn(&id, "processK8sSecretDeleteQueue: Removing safe managed linked k8s secrets is disabled")
	}
}

// deleteSecretFromKubernetes deletes a given SecretStored entity from a Kubernetes cluster.
// It handles the process of configuring a Kubernetes client, determining the appropriate
// secret name, and deleting the secret in the specified namespace.
func deleteSecretFromKubernetes(secret entity.SecretStored, errChan chan error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		errChan <- errors.Wrap(err, "could not get in-cluster config for k8s client"+err.Error())
		return
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		errChan <- errors.Wrap(err, "could not create k8s client")
		return
	}

	k8sSecretName := env.SafeSecretNamePrefix() + secret.Name
	// If the secret has k8s: prefix, then do not append a prefix; use the name
	// as is.
	if strings.HasPrefix(secret.Name, env.StoreWorkloadAsK8sSecretPrefix()) {
		k8sSecretName = strings.TrimPrefix(
			secret.Name, env.StoreWorkloadAsK8sSecretPrefix(),
		)
	}

	namespaces := secret.Meta.Namespaces
	for i, ns := range namespaces {
		if ns == "" {
			namespaces[i] = "default"
		}

		// Delete the Secret in the cluster with a backoff. If the secret is not found, it is a no-op.
		err = backoff.RetryFixed(
			ns,
			func() error {
				// First, try to get the existing secret
				_, err := clientSet.CoreV1().Secrets(ns).Get(
					context.Background(),
					k8sSecretName,
					metaV1.GetOptions{},
				)
				if kErrors.IsNotFound(err) {
					// Secret is not found in the cluster. No need to call delete.
					return nil
				}

				err = clientSet.CoreV1().Secrets(ns).Delete(
					context.Background(),
					k8sSecretName,
					metaV1.DeleteOptions{},
				)

				return err
			},
		)

		if err != nil {
			errChan <- err
		}
	}
}
