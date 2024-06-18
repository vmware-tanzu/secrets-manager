/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package initialization

type command string

// Special Commands
const sleep = "sleep"
const exit = "exit"

// Sentinel Commands
// These should match `./app/sentinel/cmd/parse.go` values.
const encrypt command = "e"
const expires command = "E"
const format command = "f"
const join command = "a"
const keys command = "i"
const namespace command = "n"
const notBefore command = "N"
const remove command = "d"
const secret command = "s"
const transformation command = "t"
const workload command = "w"
