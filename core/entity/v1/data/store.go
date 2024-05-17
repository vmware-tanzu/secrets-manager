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

// BackingStore is the backing store for the data for VSecM Safe.
type BackingStore string

var (
	Memory           BackingStore = "memory"
	File             BackingStore = "file"
	AwsSecretStore   BackingStore = "aws-secret"
	AzureSecretStore BackingStore = "azure-secret"
	GcpSecretStore   BackingStore = "gcp-secret"
	Kubernetes       BackingStore = "k8s"
)
