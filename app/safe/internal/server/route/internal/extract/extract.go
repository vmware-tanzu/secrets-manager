/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package extract

import (
	"encoding/json"
	"strings"

	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func WorkloadIDAndParts(spiffeid string) (string, []string) {
	tmp := strings.Replace(spiffeid, env.WorkloadSpiffeIdPrefix(), "", 1)
	parts := strings.Split(tmp, "/")
	if len(parts) > 0 {
		return parts[0], parts
	}
	return "", nil
}

func SecretValue(cid string, secret *entity.SecretStored) string {
	if secret.ValueTransformed != "" {
		log.TraceLn(&cid, "Fetch: using transformed value")
		return secret.ValueTransformed
	}

	// This part is for backwards compatibility.
	// It probably won’t execute because `secret.ValueTransformed` will
	// always be set.

	log.TraceLn(&cid, "Fetch: using raw value")

	if len(secret.Values) == 1 {
		return secret.Values[0]
	}

	jsonData, err := json.Marshal(secret.Values)
	if err != nil {
		log.WarnLn(&cid, "Fetch: Problem marshaling values", err.Error())
	} else {
		return string(jsonData)
	}

	return ""
}
