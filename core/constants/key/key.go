/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package key

type Identifier string

// CorrelationId is the identifier for the correlation ID.
// It is used to correlate log messages and other data across
// services while logging.
const CorrelationId Identifier = "correlationId"

const SecretDataValue = "VALUE"

const EnvVars = "ENV_VARS"
const GoVersion = "GO_VERSION"
