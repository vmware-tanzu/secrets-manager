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
	"github.com/vmware-tanzu/secrets-manager/core/constants/val"
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/v1/reqres/safe"
)

func newRootKeyUpdateRequest(
	ageSecretKey, agePublicKey, aesCipherKey string,
) reqres.KeyInputRequest {
	return reqres.KeyInputRequest{
		AgeSecretKey: ageSecretKey,
		AgePublicKey: agePublicKey,
		AesCipherKey: aesCipherKey,
	}
}

func newSecretUpsertRequest(workloadIds []string, secret string,
	namespaces []string, template string, format string,
	encrypt, appendSecret bool, notBefore string, expires string,
) reqres.SecretUpsertRequest {
	f := decideSecretFormat(format)

	if notBefore == "" {
		notBefore = val.TimeNow
	}

	if expires == "" {
		expires = val.TimeNever
	}

	return reqres.SecretUpsertRequest{
		WorkloadIds: workloadIds,
		Namespaces:  namespaces,
		Template:    template,
		Format:      f,
		Encrypt:     encrypt,
		AppendValue: appendSecret,
		Value:       secret,
		NotBefore:   notBefore,
		Expires:     expires,
	}
}
