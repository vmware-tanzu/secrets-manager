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

	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/rpc"
	"github.com/vmware-tanzu/secrets-manager/core/spiffe"
)

var seed = time.Now().UnixNano()

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

			parts := strings.Split(sc.SerializedRootKeys, "\n")
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
			newSecret, err := crypto.GenerateValue(sc.Secret)
			if err != nil {
				sc.Secret = "ParseError:" + sc.Secret
			} else {
				sc.Secret = newSecret
			}
		}

		p, err := url.JoinPath(env.EndpointUrlForSafe(), "/sentinel/v1/secrets")
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
