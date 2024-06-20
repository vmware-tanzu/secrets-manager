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

import (
	"github.com/vmware-tanzu/secrets-manager/core/constants/audit"
)

// JournalEntry represents a single entry in the audit journal.
type JournalEntry struct {
	CorrelationId string
	Payload       string
	Method        string
	Url           string
	SpiffeId      string
	Event         audit.Event
}
