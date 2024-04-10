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

import reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"

func newInputKeysRequest(ageSecretKey, agePublicKey, aesCipherKey string,
) reqres.KeyInputRequest {
	return reqres.KeyInputRequest{
		AgeSecretKey: ageSecretKey,
		AgePublicKey: agePublicKey,
		AesCipherKey: aesCipherKey,
	}
}

func newInitCompletedRequest() reqres.SentinelInitCompleteRequest {
	return reqres.SentinelInitCompleteRequest{}
}

func newSecretUpsertRequest(workloadIds []string, secret string,
	namespaces []string, backingStore string, template string, format string,
	encrypt, appendSecret bool, notBefore string, expires string,
) reqres.SecretUpsertRequest {
	bs := decideBackingStore(backingStore)
	f := decideSecretFormat(format)

	if notBefore == "" {
		notBefore = "now"
	}

	if expires == "" {
		expires = "never"
	}

	return reqres.SecretUpsertRequest{
		WorkloadIds:  workloadIds,
		BackingStore: bs,
		Namespaces:   namespaces,
		Template:     template,
		Format:       f,
		Encrypt:      encrypt,
		AppendValue:  appendSecret,
		Value:        secret,
		NotBefore:    notBefore,
		Expires:      expires,
	}
}
