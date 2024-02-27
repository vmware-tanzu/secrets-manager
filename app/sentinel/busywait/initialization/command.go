/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package initialization

type command string

const workload command = "w"
const namespace command = "n"
const secret command = "s"
const transformation command = "t"
const sleep = "sleep"
const done = "done"
