/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package index

type RootKeyIndex int

const PrivateKey RootKeyIndex = 0
const PublicKey RootKeyIndex = 1
const AesSeed RootKeyIndex = 2
