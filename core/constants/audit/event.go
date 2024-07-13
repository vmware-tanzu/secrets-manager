/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package audit

type Event string

const BadPayload Event = "vsecm-bad-payload"
const BadPeerSvid Event = "vsecm-bad-peer-spiffeid"
const BadSpiffeId Event = "vsecm-bad-spiffeid"
const BrokenBody Event = "vsecm-broken-body"
const DecryptionFailed Event = "vsecm-decryption-failed"
const EncryptionFailed Event = "vsecm-encryption-failed"
const Enter Event = "vsecm-enter"
const NoSecret Event = "vsecm-no-secret"
const NoValue Event = "vsecm-no-value"
const NotWorkload = "vsecm-not-workload"
const NoWorkloadId Event = "vsecm-no-wl-id"
const NotSentinel = "vsecm-not-sentinel"
const Ok Event = "vsecm-ok"
const RequestTypeMismatch Event = "vsecm-request-type-mismatch"
const RootKeyNotSet = "vsecm-root-key-not-set"
