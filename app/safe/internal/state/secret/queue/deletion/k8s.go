/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package deletion

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/vmware-tanzu/secrets-manager/core/backoff"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	kErrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// K8sSecretDeleteQueue contains k8s `Secret`s to be removed.
var K8sSecretDeleteQueue = make(chan entity.SecretStored, env.K8sSecretDeleteBufferSizeForSafe())

// deleteSecretFromKubernetes deletes a given SecretStored entity from a Kubernetes cluster.
// It handles the process of configuring a Kubernetes client, determining the appropriate
// secret name, and deleting the secret in the specified namespace.
func deleteSecretFromKubernetes(secret entity.SecretStored, errChan chan error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		errChan <- errors.Wrap(err, "could not get in-cluster config for k8s client"+err.Error())
		return
	}

	// If the secret does not have the k8s: prefix, then it is not a k8s secret.
	if !strings.HasPrefix(secret.Name, env.StoreWorkloadAsK8sSecretPrefix()) {
		return
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		errChan <- errors.Wrap(err, "could not create k8s client")
		return
	}

	k8sSecretName := secret.Name
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

// ProcessK8sSecretQueue manages the deletion of Kubernetes secrets that
// have been marked for deletion and placed in a queue. This operation is controlled
// by an environment configuration that determines whether the deletion of linked
// K8s secrets is permitted. The function employs goroutines for concurrent deletion
// tasks and asynchronous error reporting, aiming to optimize the deletion
// process and handle potential errors effectively.
func ProcessK8sSecretQueue() {
	errChan := make(chan error)
	id := "AEGIHK8D"

	if env.RemoveLinkedK8sSecretsModeForSafe() {
		go func() {
			for e := range errChan {
				// If the `deleteSecretFromKubernetes` operation spews out an error, log it.
				log.ErrorLn(&id, "processK8sSecretDeleteQueue: error deleting secret:", e.Error())
			}
		}()

		for secret := range K8sSecretDeleteQueue {
			go deleteSecretFromKubernetes(secret, errChan)
		}
	} else {
		log.InfoLn(&id, "processK8sSecretDeleteQueue: Removing safe managed linked k8s secrets is disabled")
	}
}
