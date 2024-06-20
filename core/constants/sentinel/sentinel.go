/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package sentinel

const CmdName = "safe"

type Command string

// Special Commands

const Sleep Command = "sleep"
const Exit Command = "exit"

// Sentinel Commands
// These should match `./app/sentinel/cmd/parse.go` values.

const List Command = "l"
const ListExp Command = "list"
const Encrypt Command = "e"
const EncryptExp Command = "encrypt"
const Expires Command = "E"
const ExpiresExp Command = "exp"
const Format Command = "f"
const FormatExp Command = "format"
const Join Command = "a"
const JoinExp Command = "append"
const Keys Command = "i"
const KeysExp Command = "input-keys"
const Namespace Command = "n"
const NamespaceExp Command = "namespace"
const NotBefore Command = "N"
const NotBeforeExp Command = "nbf"
const Remove Command = "d"
const RemoveExp Command = "delete"
const Secret Command = "s"
const SecretExp Command = "secret"
const Transformation Command = "t"
const TransformationExp Command = "template"
const Workload Command = "w"
const WorkloadExp Command = "workload"
