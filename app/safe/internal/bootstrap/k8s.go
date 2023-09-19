/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package bootstrap

import (
	"context"
	"github.com/pkg/errors"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	id string = "VSECMSAFE"
)

func ValidateEnvironment() {
	// getting metadata.namespace, passed through environment variable VSECM_SYSTEM_NAMESPACE
	if len(env.SystemNamespace()) == 0 {
		log.FatalLn(&id, "Failed to get pod namespace",
			"Pod namespace should be exported into environment as VSECM_SYSTEM_NAMESPACE")
	}
}

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
	keysCombined := privateKey + "\n" + publicKey + "\n" + aesSeed
	data["KEY_TXT"] = ([]byte)(keysCombined)

	// Update the Secret in the cluster
	_, err = k8sApi.CoreV1().Secrets(env.SystemNamespace()).Update(
		context.Background(),
		&v1.Secret{
			TypeMeta: metaV1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metaV1.ObjectMeta{
				Name:      env.SafeAgeKeySecretName(),
				Namespace: env.SystemNamespace(),
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
		return errors.Wrap(err, "Error creating the secret")
	}

	state.SetMasterKey(keysCombined)

	return nil
}
