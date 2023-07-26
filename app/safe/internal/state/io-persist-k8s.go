/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package state

import (
	"context"
	"github.com/pkg/errors"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	apiV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"time"
)

func saveSecretToKubernetes(secret entity.SecretStored) error {
	// updates the Kubernetes Secret assuming it already exists.

	config, err := rest.InClusterConfig()
	if err != nil {
		return errors.Wrap(err, "could not create client config")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return errors.Wrap(err, "could not create client")
	}

	// Transform the data if there is a transformation defined.
	data := secret.ToMapForK8s()

	// Update the Secret in the cluster
	_, err = clientset.CoreV1().Secrets(secret.Meta.Namespace).Update(
		context.Background(),
		&apiV1.Secret{
			TypeMeta: metaV1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metaV1.ObjectMeta{
				Name:      env.SafeSecretNamePrefix() + secret.Name,
				Namespace: secret.Meta.Namespace,
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

	if err != nil {
		return errors.Wrap(err, "error updating the secret")
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
