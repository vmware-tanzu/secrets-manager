/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package audit

import (
	reqres "github.com/vmware-tanzu/secrets-manager/core/entity/reqres/safe/v1"
	"github.com/vmware-tanzu/secrets-manager/core/log"
)

type Event string

const EventEnter Event = "vsecm-enter"
const EventBadSvid Event = "vsecm-bad-spiffeid"
const EventBrokenBody Event = "vsecm-broken-body"
const EventRequestTypeMismatch Event = "vsecm-request-type-mismatch"
const EventBadPeerSvid Event = "vsecm-bad-peer-spiffeid"
const EventNoSecret Event = "vsecm-no-secret"
const EventOk Event = "vsecm-ok"
const EventNoWorkloadId Event = "vsecm-no-wl-id"
const EventNoValue Event = "vsecm-no-value"
const EventEncryptionFailed Event = "vsecm-encryption-failed"
const EventDecryptionFailed Event = "vsecm-decryption-failed"
const EventBadPayload Event = "vsecm-bad-payload"

type JournalEntry struct {
	CorrelationId string
	Entity        any
	Method        string
	Url           string
	SpiffeId      string
	Event         Event
}

func printAudit(correlationId, entityName, method, url, spiffeid, message string) {
	log.AuditLn(
		&correlationId,
		entityName,
		"{{"+
			"method:[["+method+"]],"+
			"url:[["+url+"]],"+
			"spiffeid:[["+spiffeid+"]],"+
			"msg:[["+message+"]]}}",
	)
}

func Log(e JournalEntry) {
	if e.Entity == nil {
		printAudit(
			e.CorrelationId,
			"nil",
			e.Method, e.Url, e.SpiffeId, string(e.Event),
		)
	}

	switch v := e.Entity.(type) {
	case reqres.SecretDeleteRequest:
		printAudit(
			e.CorrelationId,
			"SecretDeleteRequest",
			e.Method, e.Url, e.SpiffeId,
			"w:'"+v.WorkloadId+"',e:'"+v.Err+"',m:'"+string(e.Event)+"'",
		)
	case reqres.SecretDeleteResponse:
		printAudit(
			e.CorrelationId,
			"SecretDeleteResponse",
			e.Method, e.Url, e.SpiffeId,
			"e:'"+v.Err+"',m:'"+string(e.Event)+"'",
		)
	case reqres.SecretFetchRequest:
		printAudit(
			e.CorrelationId,
			"SecretFetchRequest",
			e.Method, e.Url, e.SpiffeId,
			"e:'"+v.Err+"',m:'"+string(e.Event)+"'",
		)
	case reqres.SecretFetchResponse:
		printAudit(
			e.CorrelationId,
			"SecretFetchResponse",
			e.Method, e.Url, e.SpiffeId,
			"e:'"+v.Err+",'c:'"+v.Created+",'u:'"+v.Updated+",'m:'"+string(e.Event)+"'",
		)
	case reqres.SecretUpsertRequest:
		printAudit(
			e.CorrelationId,
			"SecretUpsertRequest",
			e.Method, e.Url, e.SpiffeId,
			"e:'"+v.Err+"',m:'"+string(e.Event)+"'",
		)
	case reqres.SecretUpsertResponse:
		printAudit(
			e.CorrelationId,
			"SecretUpsertResponse",
			e.Method, e.Url, e.SpiffeId,
			"e:'"+v.Err+"',m:'"+string(e.Event)+"'",
		)
	case reqres.SecretListRequest:
		printAudit(
			e.CorrelationId,
			"SecretListRequest",
			e.Method, e.Url, e.SpiffeId,
			"e:'"+v.Err+"',m:'"+string(e.Event)+"'",
		)
	case reqres.SecretListResponse:
		printAudit(
			e.CorrelationId,
			"SecretListResponse",
			e.Method, e.Url, e.SpiffeId,
			"e:'"+v.Err+"',m:'"+string(e.Event)+"'",
		)
	case reqres.SecretEncryptedListResponse:
		printAudit(
			e.CorrelationId,
			"SecretEncryptedListResponse",
			e.Method, e.Url, e.SpiffeId,
			"e:'"+v.Err+"',m:'"+string(e.Event)+"'",
		)
	case reqres.KeyInputRequest:
		printAudit(
			e.CorrelationId,
			"KeyInputRequest",
			e.Method, e.Url, e.SpiffeId,
			"e:'"+v.Err+"',m:'"+string(e.Event)+"'",
		)
	default:
		printAudit(
			e.CorrelationId,
			"UnknownEntity",
			e.Method, e.Url, e.SpiffeId,
			"e: UNKNOWN ENTITY in AUDIT LOG",
		)
	}
}
