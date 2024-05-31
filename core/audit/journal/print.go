/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

import (
	"github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func printAudit(correlationId string, e state.Event,
	method, url, spiffeid, payload string) {
	std.AuditLn(
		&correlationId,
		string(e),
		"{{"+
			"method:[["+method+"]],"+
			"url:[["+url+"]],"+
			"spiffeid:[["+spiffeid+"]],"+
			"payload:[["+payload+"]]}}",
	)
}