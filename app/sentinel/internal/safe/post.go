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
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	u "github.com/vmware-tanzu/secrets-manager/core/constants/url"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/rpc"
	"github.com/vmware-tanzu/secrets-manager/core/spiffe"
	"github.com/vmware-tanzu/secrets-manager/lib/template"
)

var seed = time.Now().UnixNano()

// Post handles the posting of secrets to the VSecM Safe API using the
// provided SentinelCommand.
//
// This function performs the following steps:
//  1. Creates a context with a timeout based on the parent context and
//     environment settings.
//  2. Computes a hash of the secret for logging purposes if configured to do so.
//  3. Acquires a workload source and proceeds only if the source acquisition
//     is successful.
//  4. Depending on the SentinelCommand, it either posts new secrets or deletes
//     existing ones.
//
// Parameters:
//   - parentContext: The parent context for the request, used for tracing and
//     cancellation.
//   - sc: The SentinelCommand containing details for the secret management
//     operation.
//
// Returns:
//   - An error if the operation fails, or nil if successful.
//
// Example usage:
//
//	parentContext := context.Background()
//	sc := entity.SentinelCommand{
//	    WorkloadIds:        []string{"workload1"},
//	    Secret:             "my-secret",
//	    Namespaces:         []string{"namespace1"},
//	    SerializedRootKeys: "key1\nkey2\nkey3",
//	}
//	err := Post(parentContext, sc)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Error Handling:
//   - If the context times out or is canceled, it logs the error and returns
//     an appropriate message.
//   - If there is an error during source acquisition, secret generation, or
//     payload processing, it returns an error with details.
func Post(parentContext context.Context,
	sc entity.SentinelCommand,
) error {
	ctxWithTimeout, cancel := context.WithTimeout(
		parentContext,
		env.SourceAcquisitionTimeoutForSafe(),
	)
	defer cancel()

	cid := ctxWithTimeout.Value(key.CorrelationId).(*string)

	ids := ""
	for _, id := range sc.WorkloadIds {
		ids += id + ", "
	}

	hashString := "<>"
	if env.LogSecretFingerprints() {
		secret := sc.Secret
		uniqueData := fmt.Sprintf("%s-%d", secret, seed)
		dataBytes := []byte(uniqueData)
		hasher := sha256.New()
		hasher.Write(dataBytes)
		hashBytes := hasher.Sum(nil)
		hashString = hex.EncodeToString(hashBytes)
	}

	log.AuditLn(cid, "Sentinel:Post: ws:", ids, "h:", hashString)

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
			return errors.Join(
				ctxWithTimeout.Err(),
				errors.New("post:"+
					" I cannot execute command because I cannot talk to SPIRE"),
			)
		}

		return errors.New("post: Operation was cancelled due to an unknown reason")
	case source := <-sourceChan:
		defer func(s *workloadapi.X509Source) {
			if s == nil {
				return
			}
			err := s.Close()
			if err != nil {
				log.ErrorLn(cid, "post: Problem closing the workload source")
			}
		}(source)

		proceed := <-proceedChan
		if !proceed {
			return errors.New("post: Could not acquire source for Sentinel")
		}

		authorizer := createAuthorizer()

		if sc.SerializedRootKeys != "" {
			log.InfoLn(cid, "Post: I am going to post the root keys.")

			p, err := url.JoinPath(env.EndpointUrlForSafe(), "/sentinel/v1/keys")
			if err != nil {
				return errors.New("post: I am having problem" +
					" generating VSecM Safe secrets api endpoint URL")
			}

			tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
			client := &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: tlsConfig,
				},
			}

			parts := sc.SplitRootKeys()
			if len(parts) != 3 {
				return errors.New("post: Bad data! Very bad data")
			}

			sr := newRootKeyUpdateRequest(parts[0], parts[1], parts[2])
			md, err := json.Marshal(sr)
			if err != nil {
				return errors.Join(
					err,
					errors.New("post: I am having problem generating the payload"),
				)
			}

			return doPost(cid, client, p, md)
		}

		log.InfoLn(cid, "Post: I am going to post the secrets.")

		// Generate pattern-based random secrets if the secret has the prefix.
		if strings.HasPrefix(sc.Secret, env.SecretGenerationPrefix()) {
			sc.Secret = strings.Replace(
				sc.Secret, env.SecretGenerationPrefix(), "", 1,
			)
			newSecret, err := template.Value(sc.Secret)
			if err != nil {
				sc.Secret = "ParseError:" + sc.Secret
			} else {
				sc.Secret = newSecret
			}
		}

		p, err := url.JoinPath(env.EndpointUrlForSafe(), u.SentinelSecrets)
		if err != nil {
			return errors.Join(
				err,
				errors.New("post: I am having problem "+
					"generating VSecM Safe secrets api endpoint URL"),
			)
		}

		tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}

		sr := newSecretUpsertRequest(sc.WorkloadIds, sc.Secret, sc.Namespaces,
			sc.Template, sc.Format,
			sc.Encrypt, sc.AppendSecret, sc.NotBefore, sc.Expires)

		md, err := json.Marshal(sr)
		if err != nil {
			return errors.Join(
				err,
				errors.New("post: I am having problem generating the payload"),
			)
		}

		if sc.DeleteSecret {
			return doDelete(cid, client, p, md)
		}

		return doPost(cid, client, p, md)
	}
}
