/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package data

type SecretFormat string

var (
	Json SecretFormat = "json"
	Yaml SecretFormat = "yaml"
	Raw  SecretFormat = "raw"
)

type Secret struct {
	Name         string   `json:"name"`
	Created      JsonTime `json:"created"`
	Updated      JsonTime `json:"updated"`
	NotBefore    JsonTime `json:"notBefore"`
	ExpiresAfter JsonTime `json:"expiresAfter"`
}

type SecretEncrypted struct {
	Name           string   `json:"name"`
	EncryptedValue []string `json:"value"`
	Created        JsonTime `json:"created"`
	Updated        JsonTime `json:"updated"`
	NotBefore      JsonTime `json:"notBefore"`
	ExpiresAfter   JsonTime `json:"expiresAfter"`
}

type SecretStringTime struct {
	Name           string   `json:"name"`
	EncryptedValue []string `json:"value"`
	Created        string   `json:"created"`
	Updated        string   `json:"updated"`
	NotBefore      JsonTime `json:"notBefore"`
	ExpiresAfter   JsonTime `json:"expiresAfter"`
}

type SecretMeta struct {
	// Defaults to "default"
	Namespaces []string `json:"namespaces"`
	// Go template used to transform the secret.
	// Sample secret:
	// '{"username":"admin","password":"VSecMRocks"}'
	// Sample template:
	// '{"USER":"{{.username}}", "PASS":"{{.password}}"}"
	Template string `json:"template"`
	// Defaults to None
	Format SecretFormat
	// For tracking purposes
	CorrelationId string `json:"correlationId"`
}
