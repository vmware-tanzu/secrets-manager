/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package constants

type Identifier string

// CorrelationId is the identifier for the correlation ID.
// It is used to correlate log messages and other data across
// services while logging.
const CorrelationId Identifier = "correlationId"

type EnvVarName string

const AppVersion EnvVarName = "APP_VERSION"
const VSecMLogLevel EnvVarName = "VSECM_LOG_LEVEL"
const VSecMSafeSpiffeIdPrefix EnvVarName = "VSECM_SAFE_SPIFFEID_PREFIX"
const VSecMSafeEndpointUrl EnvVarName = "VSECM_SAFE_ENDPOINT_URL"
const VSecMKeyGenDecryptMode EnvVarName = "VSECM_KEYGEN_DECRYPT"

type FieldName string

const RootKeyText FieldName = "KEY_TXT"
