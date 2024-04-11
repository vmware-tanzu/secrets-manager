/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package v1

import (
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	data "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
)

type SecretUpsertRequest struct {
	WorkloadIds  []string          `json:"workloads"`
	BackingStore data.BackingStore `json:"backingStore"`
	Namespaces   []string          `json:"namespaces"`
	Value        string            `json:"value"`
	Template     string            `json:"template"`
	Format       data.SecretFormat `json:"format"`
	Encrypt      bool              `json:"encrypt"`
	AppendValue  bool              `json:"appendValue"`
	NotBefore    string            `json:"notBefore"`
	Expires      string            `json:"expires"`

	Err string `json:"err,omitempty"`
}

type KeyInputRequest struct {
	AgeSecretKey string `json:"ageSecretKey"`
	AgePublicKey string `json:"agePublicKey"`
	AesCipherKey string `json:"aesCipherKey"`
	Err          string `json:"err,omitempty"`
}

type SecretUpsertResponse struct {
	Err string `json:"err,omitempty"`
}

type SentinelInitCompleteRequest struct {
	Err string `json:"err,omitempty"`
}

type SentinelInitCompleteResponse struct {
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
	WorkloadIds []string `json:"workloads"`
	Err         string   `json:"err,omitempty"`
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
	Algorithm crypto.Algorithm       `json:"algorithm"`
	Err       string                 `json:"err,omitempty"`
}

type SecretStringTimeListResponse struct {
	Secrets   []data.SecretStringTime `json:"secrets"`
	Algorithm crypto.Algorithm        `json:"algorithm"`
	Err       string                  `json:"err,omitempty"`
}

type KeystoneStatusRequest struct {
	Err string `json:"err,omitempty"`
}

type KeystoneStatusResponse struct {
	Status data.KeystoneInitStatus `json:"status"`
	Err    string                  `json:"err,omitempty"`
}

type GenericRequest struct {
	Err string `json:"err,omitempty"`
}

type GenericResponse struct {
	Err string `json:"err,omitempty"`
}
