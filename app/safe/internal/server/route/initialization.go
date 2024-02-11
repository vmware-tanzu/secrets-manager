/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package route

import (
	"context"
	"github.com/pkg/errors"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/backoff"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state"
	"github.com/vmware-tanzu/secrets-manager/core/audit"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
	apiV1 "k8s.io/api/core/v1"
	kErrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"net/http"
)

func markInitializationSecretAsCompleted() error {
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
	err = backoff.RetryLinear(
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

// InitComplete is called when the Sentinel has completed its initialization
// process. It is responsible for marking the initialization process as
// complete in the Kubernetes cluster, by updating the value of a
// "vsecm-sentinel-init-tombstone" Secret in the "vsecm-system" namespace.
//
// See ./app/sentinel/internal/safe/post.go:PostInitializationComplete for the
// corresponding sentinel-side implementation.
func InitComplete(cid string, w http.ResponseWriter, r *http.Request, spiffeid string) {
	if env.SafeManualKeyInput() && !state.MasterKeySet() {
		log.InfoLn(&cid, "InitComplete: Master key not set")
		return
	}

	j := audit.JournalEntry{
		CorrelationId: cid,
		Entity:        reqres.SentinelInitCompleteRequest{},
		Method:        r.Method,
		Url:           r.RequestURI,
		SpiffeId:      spiffeid,
		Event:         audit.EventEnter,
	}

	audit.Log(j)

	if !validation.IsSentinel(spiffeid) {
		handleBadSvidResponse(cid, w, spiffeid, j)
		return
	}

	log.DebugLn(&cid, "InitComplete: sending response")

	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.InfoLn(&cid, "InitComplete: Problem closing body")
		}
	}()

	log.DebugLn(&cid, "InitComplete: preparing request")

	err := markInitializationSecretAsCompleted()
	if err != nil {
		log.WarnLn(
			&cid,
			"InitComplete: Problem creating initialization secret",
			err.Error(),
		)
	}

	icr := reqres.SentinelInitCompleteResponse{}

	handleInitCompleteSuccessResponse(cid, w, j, icr)
}
