/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package safe

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/core/backoff"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/rpc"
	"github.com/vmware-tanzu/secrets-manager/core/spiffe"
)

func backoffStrategy() backoff.Strategy {
	return backoff.Strategy{
		MaxRetries:  20,
		Delay:       1000,
		Exponential: true,
		MaxDuration: 30 * time.Second,
	}
}

// MarkInitializationCompletion is a function that signals the completion of a
// post-initialization process.
// It takes a parent context as an argument and performs several steps involving
// timeout management, source acquisition, error handling, and sending a
// notification about the initialization completion.
//
// In a separate goroutine, it tries to acquire a source and sends the source and
// a proceed signal back to the main function through channels. The main function
// then waits for either a timeout or a source to be returned.
//
// If a timeout occurs, it logs an error depending on whether it's due to deadline
// exceeded or an unknown reason. If a source is received, it checks whether to
// proceed. If not, it returns early.
//
// If proceeding, the function then creates an authorizer and builds a client with
// mutual TLS configuration. It creates a new request payload, marshals it to
// JSON, and sends a POST request to a specified endpoint.
//
// Parameters:
//   - parentContext (context.Context): The parent context from which the function
//     will derive its context.
func MarkInitializationCompletion(parentContext context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(
		parentContext,
		env.SourceAcquisitionTimeoutForSafe(),
	)
	defer cancel()

	cid := ctxWithTimeout.Value("correlationId").(*string)

	log.AuditLn(cid, "Sentinel:MarkInitializationCompletion")

	sourceChan := make(chan *workloadapi.X509Source)
	proceedChan := make(chan bool)

	go func() {
		source, proceed := spiffe.AcquireSourceForSentinel(ctxWithTimeout)
		sourceChan <- source
		proceedChan <- proceed
	}()

	select {
	case <-ctxWithTimeout.Done():
		if errors.Is(ctxWithTimeout.Err(), context.DeadlineExceeded) {
			log.ErrorLn(cid, "PostInit: I cannot execute command because I cannot talk to SPIRE.")
			return
		}

		log.ErrorLn(cid, "PostInit: Operation was cancelled due to an unknown reason.")
	case source := <-sourceChan:
		defer func() {
			if source == nil {
				return
			}
			err := source.Close()
			if err != nil {
				log.ErrorLn(cid, "Post: Problem closing the workload source.")
			}
		}()

		proceed := <-proceedChan

		if !proceed {
			return
		}

		authorizer := createAuthorizer()

		p, err := url.JoinPath(env.EndpointUrlForSafe(), "/sentinel/v1/init-completed")
		if err != nil {
			printEndpointError(cid, err)
			return
		}

		tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}

		sr := newInitCompletedRequest()

		md, err := json.Marshal(sr)
		if err != nil {
			printPayloadError(cid, err)
			return
		}

		// Try forever until success.
		for {
			s := backoffStrategy()

			err := backoff.Retry("sentinel:post", func() error {
				log.TraceLn(cid, "sentinel:post")

				err := doPost(cid, client, p, md)
				if err != nil {
					log.ErrorLn(
						cid,
						"sentinel:post: error:", err.Error(), "will retry.",
					)
				}

				return err
			}, s)

			if err == nil {
				continue
			}
		}
	}
}

// var seed = time.Now().UnixNano()

func Post(parentContext context.Context,
	sc entity.SentinelCommand,
) error {
	ctxWithTimeout, cancel := context.WithTimeout(
		parentContext,
		env.SourceAcquisitionTimeoutForSafe(),
	)
	defer cancel()

	cid := ctxWithTimeout.Value("correlationId").(*string)

	ids := ""
	for _, id := range sc.WorkloadIds {
		ids += id + ", "
	}

	//// TODO: make this optional and disabled by default
	//secret := sc.Secret
	//uniqueData := fmt.Sprintf("%s-%d", secret, seed)
	//dataBytes := []byte(uniqueData)
	//hasher := sha256.New()
	//hasher.Write(dataBytes)
	//hashBytes := hasher.Sum(nil)
	//hashString := hex.EncodeToString(hashBytes)
	hashString := "TBD"
	log.AuditLn(cid, "Sentinel:Post: workloadIds:", ids, "hash", hashString)

	sourceChan := make(chan *workloadapi.X509Source)
	proceedChan := make(chan bool)

	go func() {
		source, proceed := spiffe.AcquireSourceForSentinel(ctxWithTimeout)
		sourceChan <- source
		proceedChan <- proceed
	}()

	select {
	case <-ctxWithTimeout.Done():
		if errors.Is(ctxWithTimeout.Err(), context.DeadlineExceeded) {
			return errors.Wrap(
				ctxWithTimeout.Err(),
				"Post: I cannot execute command because I cannot talk to SPIRE.",
			)
		}

		return errors.New("Post: Operation was cancelled due to an unknown reason.")
	case source := <-sourceChan:
		defer func() {
			if source == nil {
				return
			}
			err := source.Close()
			if err != nil {
				log.ErrorLn(cid, "Post: Problem closing the workload source.")
			}
		}()

		proceed := <-proceedChan
		if !proceed {
			return errors.New("Post: Could not acquire source for Sentinel.")
		}

		authorizer := createAuthorizer()

		if sc.InputKeys != "" {
			p, err := url.JoinPath(env.EndpointUrlForSafe(), "/sentinel/v1/keys")
			if err != nil {
				return errors.New("Post: I am having problem generating VSecM Safe secrets api endpoint URL.")
			}

			tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
			client := &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: tlsConfig,
				},
			}

			parts := strings.Split(sc.InputKeys, "\n")
			if len(parts) != 3 {
				return errors.New("post: Bad data! Very bad data")
			}

			sr := newInputKeysRequest(parts[0], parts[1], parts[2])
			md, err := json.Marshal(sr)
			if err != nil {
				return errors.Wrap(err, "Post: I am having problem generating the payload.")
			}

			return doPost(cid, client, p, md)
		}

		// Generate pattern-based random secrets if the secret has the prefix.
		if strings.HasPrefix(sc.Secret, env.SecretGenerationPrefix()) {
			sc.Secret = strings.Replace(
				sc.Secret, env.SecretGenerationPrefix(), "", 1,
			)
			newSecret, err := crypto.GenerateValue(sc.Secret)
			if err != nil {
				sc.Secret = "ParseError:" + sc.Secret
			} else {
				sc.Secret = newSecret
			}
		}

		p, err := url.JoinPath(env.EndpointUrlForSafe(), "/sentinel/v1/secrets")
		if err != nil {
			return errors.Wrap(
				err,
				"Post: I am having problem generating VSecM Safe secrets api endpoint URL.",
			)
		}

		tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}

		sr := newSecretUpsertRequest(sc.WorkloadIds, sc.Secret, sc.Namespaces,
			sc.BackingStore, sc.Template, sc.Format,
			sc.Encrypt, sc.AppendSecret, sc.NotBefore, sc.Expires)

		md, err := json.Marshal(sr)
		if err != nil {
			return errors.Wrap(err, "Post: I am having problem generating the payload.")
		}

		if sc.DeleteSecret {
			return doDelete(cid, client, p, md)
		}

		return doPost(cid, client, p, md)
	}
}
