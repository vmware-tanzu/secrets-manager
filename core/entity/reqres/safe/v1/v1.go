/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package v1

import (
	data "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
)

type SecretUpsertRequest struct {
	WorkloadId    string            `json:"key"`
	BackingStore  data.BackingStore `json:"backingStore"`
	UseKubernetes bool              `json:"useKubernetes"`
	Namespace     string            `json:"namespace"`
	Value         string            `json:"value"`
	Template      string            `json:"template"`
	Format        data.SecretFormat `json:"format"`
	Encrypt       bool              `json:"encrypt"`
	AppendValue   bool              `json:"appendValue"`
	Err           string            `json:"err,omitempty"`
}

type KeyInputRequest struct {
	AgeSecretKey string `json:"ageSecretKey"`
	AgePublicKey string `json:"agePublicKey"`
	AesCipherKey string `json:"aesCipherKey"`
}

type SecretUpsertResponse struct {
	Err string `json:"err,omitempty"`
}

type SecretFetchRequest struct {
	Err string `json:"err,omitempty"`
}

type SecretFetchResponse struct {
	Data    string `json:"data"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Err     string `json:"err,omitempty"`
}

type SecretDeleteRequest struct {
	WorkloadId string `json:"key"`
	Err        string `json:"err,omitempty"`
}

type SecretDeleteResponse struct {
	Err string `json:"err,omitempty"`
}

type SecretListRequest struct {
	Err string `json:"err,omitempty"`
}

type SecretListResponse struct {
	Secrets []data.Secret `json:"secrets"`
	Err     string        `json:"err,omitempty"`
}

type SecretEncryptedListResponse struct {
	Secrets   []data.SecretEncrypted `json:"secrets"`
	Algorithm string                 `json:"algorithm"`
}

type GenericRequest struct {
	Err string `json:"err,omitempty"`
}

type GenericResponse struct {
	Err string `json:"err,omitempty"`
}
