/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import (
	"os"
)

type VarName string

const AppVersion VarName = "APP_VERSION"
const SpiffeEndpointSocket VarName = "SPIFFE_ENDPOINT_SOCKET"
const SpiffeTrustDomain VarName = "SPIFFE_TRUST_DOMAIN"
const VSecMBackoffDelay VarName = "VSECM_BACKOFF_DELAY"
const VSecMBackoffMaxRetries VarName = "VSECM_BACKOFF_MAX_RETRIES"
const VSecMBackoffMaxWait VarName = "VSECM_BACKOFF_MAX_WAIT"
const VSecMBackoffMode VarName = "VSECM_BACKOFF_MODE"
const VSecMInitContainerPollInterval VarName = "VSECM_INIT_CONTAINER_POLL_INTERVAL"
const VSecMInitContainerWaitBeforeExit VarName = "VSECM_INIT_CONTAINER_WAIT_BEFORE_EXIT"
const VSecMKeygenDecrypt VarName = "VSECM_KEYGEN_DECRYPT"
const VSecMKeygenExportedSecretPath VarName = "VSECM_KEYGEN_EXPORTED_SECRET_PATH"
const VSecMKeygenRootKeyPath VarName = "VSECM_KEYGEN_ROOT_KEY_PATH"
const VSecMLogLevel VarName = "VSECM_LOG_LEVEL"
const VSecMLogSecretFingerprints VarName = "VSECM_LOG_SECRET_FINGERPRINTS"
const VSecMNamespaceSystem VarName = "VSECM_NAMESPACE_SYSTEM"
const VSecMProbeLivenessPort VarName = "VSECM_PROBE_LIVENESS_PORT"
const VSecMProbeReadinessPort VarName = "VSECM_PROBE_READINESS_PORT"
const VSecMRootKeyInputModeManual VarName = "VSECM_ROOT_KEY_INPUT_MODE_MANUAL"
const VSecMRootKeyName VarName = "VSECM_ROOT_KEY_NAME"
const VSecMRootKeyPath VarName = "VSECM_ROOT_KEY_PATH"
const VSecMSafeBackingStore VarName = "VSECM_SAFE_BACKING_STORE"
const VSecMSafeBootstrapTimeout VarName = "VSECM_SAFE_BOOTSTRAP_TIMEOUT"
const VSecMSafeDataPath VarName = "VSECM_SAFE_DATA_PATH"
const VSecMSafeEndpointUrl VarName = "VSECM_SAFE_ENDPOINT_URL"
const VSecMSafeFipsCompliant VarName = "VSECM_SAFE_FIPS_COMPLIANT"
const VSecMSafeIvInitializationInterval VarName = "VSECM_SAFE_IV_INITIALIZATION_INTERVAL"
const VSecMSafeK8sSecretBufferSize VarName = "VSECM_SAFE_K8S_SECRET_BUFFER_SIZE"
const VSecMSafeRootKeyStore VarName = "VSECM_SAFE_ROOT_KEY_STORE"
const VSecMSafeSecretBackupCount VarName = "VSECM_SAFE_SECRET_BACKUP_COUNT"
const VSecMSafeSecretBufferSize VarName = "VSECM_SAFE_SECRET_BUFFER_SIZE"
const VSecMSafeSecretDeleteBufferSize VarName = "VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE"
const VSecMSafeSourceAcquisitionTimeout VarName = "VSECM_SAFE_SOURCE_ACQUISITION_TIMEOUT"
const VSecMSafeStoreWorkloadSecretAsK8sSecretPrefix = "VSECM_SAFE_STORE_WORKLOAD_SECRET_AS_K8S_SECRET_PREFIX"
const VSecMSafeSyncDeletedSecrets VarName = "VSECM_SAFE_SYNC_DELETED_SECRETS"
const VSecMSafeSyncExpiredSecrets VarName = "VSECM_SAFE_SYNC_EXPIRED_SECRETS"
const VSecMSafeSyncInterpolatedK8sSecrets VarName = "VSECM_SAFE_SYNC_INTERPOLATED_K8S_SECRETS"
const VSecMSafeSyncRootKeyInterval VarName = "VSECM_SAFE_SYNC_ROOT_KEY_INTERVAL"
const VSecMSafeSyncSecretsInterval VarName = "VSECM_SAFE_SYNC_SECRETS_INTERVAL"
const VSecMSafeTlsPort VarName = "VSECM_SAFE_TLS_PORT"
const VSecMSentinelInitCommandPath VarName = "VSECM_SENTINEL_INIT_COMMAND_PATH"
const VSecMSentinelInitCommandWaitAfterInitComplete VarName = "VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE"
const VSecMSentinelInitCommandWaitBeforeExec VarName = "VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC"

// See ADR-0017 for a discussion about important security considerations when
// using OIDC.

const VSecMSentinelOidcEnableResourceServer VarName = "VSECM_SENTINEL_OIDC_ENABLE_RESOURCE_SERVER"
const VSecMSentinelOidcProviderBaseUrl VarName = "VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL"
const VSecMSentinelOidcResourceServerPort VarName = "VSECM_SENTINEL_OIDC_RESOURCE_SERVER_PORT"

const VSecMSentinelSecretGenerationPrefix VarName = "VSECM_SENTINEL_SECRET_GENERATION_PREFIX"
const VSecMSidecarErrorThreshold VarName = "VSECM_SIDECAR_ERROR_THRESHOLD"
const VSecMSidecarExponentialBackoffMultiplier VarName = "VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER"
const VSecMSidecarMaxPollInterval VarName = "VSECM_SIDECAR_MAX_POLL_INTERVAL"
const VSecMSidecarPollInterval VarName = "VSECM_SIDECAR_POLL_INTERVAL"
const VSecMSidecarSecretsPath VarName = "VSECM_SIDECAR_SECRETS_PATH"
const VSecMSidecarSuccessThreshold VarName = "VSECM_SIDECAR_SUCCESS_THRESHOLD"
const VSecMSpiffeIdPrefixSafe VarName = "VSECM_SPIFFEID_PREFIX_SAFE"
const VSecMSpiffeIdPrefixSentinel VarName = "VSECM_SPIFFEID_PREFIX_SENTINEL"
const VSecMSpiffeIdPrefixWorkload VarName = "VSECM_SPIFFEID_PREFIX_WORKLOAD"
const VSecMWorkloadNameRegExp VarName = "VSECM_WORKLOAD_NAME_REGEXP"

type VarValue string

const SpiffeEndpointSocketDefault VarValue = "unix:///spire-agent-socket/spire-agent.sock"
const SpiffeTrustDomainDefault VarValue = "vsecm.com"
const VSecMBackoffDelayDefault VarValue = "1000"
const VSecMBackoffMaxRetriesDefault VarValue = "10"
const VSecMBackoffMaxWaitDefault VarValue = "30000"
const VSecMInitContainerPollIntervalDefault VarValue = "5000"
const VSecMInitContainerWaitBeforeExitDefault VarValue = "0"
const VSecMKeygenExportedSecretPathDefault VarValue = "/opt/vsecm/secrets.json"
const VSecMKeygenRootKeyPathDefault VarValue = "/opt/vsecm/keys.txt"
const VSecMProbeLivenessPortDefault VarValue = ":8081"
const VSecMProbeReadinessPortDefault VarValue = ":8082"
const VSecMRootKeyNameDefault VarValue = "vsecm-root-key"
const VSecMRootKeyPathDefault VarValue = "/key/key.txt"
const VSecMSafeBootstrapTimeoutDefault VarValue = "300000"
const VSecMSafeDataPathDefault VarValue = "/var/local/vsecm/data"
const VSecMSafeEndpointUrlDefault VarValue = "https://vsecm-safe.vsecm-system.svc.cluster.local:8443/"
const VSecMSafeIvInitializationIntervalDefault VarValue = "50"
const VSecMSafeK8sSecretBufferSizeDefault VarValue = "10"
const VSecMSafeSecretBackupCountDefault VarValue = "3"
const VSecMSafeSecretBufferSizeDefault VarValue = "10"
const VSecMSafeSecretDeleteBufferSizeDefault VarValue = "10"
const VSecMSafeSourceAcquisitionTimeoutDefault VarValue = "10000"
const VSecMSafeStoreWorkloadSecretAsK8sSecretPrefixDefault VarValue = "k8s:"
const VSecMSafeTlsPortDefault VarValue = ":8443"
const VSecMSentinelInitCommandPathDefault VarValue = "/opt/vsecm-sentinel/init/data"
const VSecMSentinelInitCommandWaitAfterInitCompleteDefault VarValue = "0"
const VSecMSentinelInitCommandWaitBeforeExecDefault VarValue = "0"
const VSecMSentinelSecretGenerationPrefixDefault VarValue = "gen:"
const VSecMSentinelOidcResourceServerPortDefault VarValue = ":8085"
const VSecMSidecarErrorThresholdDefault VarValue = "3"
const VSecMSidecarExponentialBackoffMultiplierDefault VarValue = "2"
const VSecMSidecarMaxPollIntervalDefault VarValue = "300000"
const VSecMSidecarPollIntervalDefault VarValue = "20000"
const VSecMSidecarSecretsPathDefault VarValue = "/opt/vsecm/secrets.json"
const VSecMSidecarSuccessThresholdDefault VarValue = "3"
const VSecMSpiffeIdPrefixSafeDefault VarValue = "^spiffe://vsecm.com/workload/vsecm-safe/ns/vsecm-system/sa/vsecm-safe/n/[^/]+$"
const VSecMSpiffeIdPrefixSentinelDefault VarValue = "^spiffe://vsecm.com/workload/vsecm-sentinel/ns/vsecm-system/sa/vsecm-sentinel/n/[^/]+$"
const VSecMSpiffeIdPrefixWorkloadDefault VarValue = "^spiffe://vsecm.com/workload/[^/]+/ns/[^/]+/sa/[^/]+/n/[^/]+$"
const VSecMNameRegExpForWorkloadDefault VarValue = "^spiffe://vsecm.com/workload/([^/]+)/ns/[^/]+/sa/[^/]+/n/[^/]+$"

type Namespace string

const Default Namespace = "default"
const VSecMSystem Namespace = "vsecm-system"
const SpireSystem Namespace = "spire-system"
const SpireServer Namespace = "spire-server"

func Value(name VarName) string {
	return os.Getenv(string(name))
}

type FieldName string

const RootKeyText FieldName = "KEY_TXT"
