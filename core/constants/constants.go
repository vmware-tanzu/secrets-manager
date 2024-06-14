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

import (
	"os"
	"strings"
)

type Identifier string

// CorrelationId is the identifier for the correlation ID.
// It is used to correlate log messages and other data across
// services while logging.
const CorrelationId Identifier = "correlationId"

type EnvVarName string

const AppVersion EnvVarName = "APP_VERSION"
const VSecMLogLevel EnvVarName = "VSECM_LOG_LEVEL"
const VSecMSpiffeIdPrefixSafe EnvVarName = "VSECM_SPIFFEID_PREFIX_SAFE"
const VSecMSpiffeIdPrefixWorkload EnvVarName = "VSECM_SPIFFEID_PREFIX_WORKLOAD"
const VSecMSpiffeIdPrefixSentinel EnvVarName = "VSECM_SPIFFEID_PREFIX_SENTINEL"
const VSecMSafeEndpointUrl EnvVarName = "VSECM_SAFE_ENDPOINT_URL"
const VSecMKeygenDecryptMode EnvVarName = "VSECM_KEYGEN_DECRYPT"
const VSecMBackoffMaxRetries EnvVarName = "VSECM_BACKOFF_MAX_RETRIES"
const VSecMBackoffDelay EnvVarName = "VSECM_BACKOFF_DELAY"
const VSecMBackoffMode EnvVarName = "VSECM_BACKOFF_MODE"
const VSecMBackoffMaxWait EnvVarName = "VSECM_BACKOFF_MAX_WAIT"
const VSecMInitContainerPollInterval EnvVarName = "VSECM_INIT_CONTAINER_POLL_INTERVAL"
const VSecMInitContainerWaitBeforeExit EnvVarName = "VSECM_INIT_CONTAINER_WAIT_BEFORE_EXIT"
const VSecMKeygenRootKeyPath EnvVarName = "VSECM_KEYGEN_ROOT_KEY_PATH"
const VSecMKeygenExportedSecretPath EnvVarName = "VSECM_KEYGEN_EXPORTED_SECRET_PATH"
const VSecMKeygenDecrypt EnvVarName = "VSECM_KEYGEN_DECRYPT"
const VSecMLogSecretFingerprints EnvVarName = "VSECM_LOG_SECRET_FINGERPRINTS"
const VSecMNamespaceSystem EnvVarName = "VSECM_NAMESPACE_SYSTEM"
const VSecMSidecarMaxPollInterval EnvVarName = "VSECM_SIDECAR_MAX_POLL_INTERVAL"
const VSecMSidecarExponentialBackoffMultiplier EnvVarName = "VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER"
const VSecMSidecarSuccessThreshold EnvVarName = "VSECM_SIDECAR_SUCCESS_THRESHOLD"
const VSecMSidecarErrorThreshold EnvVarName = "VSECM_SIDECAR_ERROR_THRESHOLD"
const VSecMSidecarPollInterval EnvVarName = "VSECM_SIDECAR_POLL_INTERVAL"
const VSecMProbeLivenessPort EnvVarName = "VSECM_PROBE_LIVENESS_PORT"
const VSecMProbeReadinessPort EnvVarName = "VSECM_PROBE_READINESS_PORT"
const VSecMSafeIvInitializationInterval EnvVarName = "VSECM_SAFE_IV_INITIALIZATION_INTERVAL"
const VSecMSafeSecretBufferSize EnvVarName = "VSECM_SAFE_SECRET_BUFFER_SIZE"
const VSecMSafeK8sSecretBufferSize EnvVarName = "VSECM_SAFE_K8S_SECRET_BUFFER_SIZE"
const VSecMSafeSecretDeleteBufferSize EnvVarName = "VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE"
const VSecMSafeFipsCompliant EnvVarName = "VSECM_SAFE_FIPS_COMPLIANT"
const VSecMSafeSecretBackupCount EnvVarName = "VSECM_SAFE_SECRET_BACKUP_COUNT"
const VSecMRootKeyInputModeManual EnvVarName = "VSECM_ROOT_KEY_INPUT_MODE_MANUAL"
const VSecMSafeDataPath EnvVarName = "VSECM_SAFE_DATA_PATH"
const VSecMRootKeyPath EnvVarName = "VSECM_ROOT_KEY_PATH"
const VSecMSafeSourceAcquisitionTimeout EnvVarName = "VSECM_SAFE_SOURCE_ACQUISITION_TIMEOUT"
const VSecMSafeBootstrapTimeout EnvVarName = "VSECM_SAFE_BOOTSTRAP_TIMEOUT"
const VSecMRootKeyName EnvVarName = "VSECM_ROOT_KEY_NAME"
const VSecMSentinelSecretGenerationPrefix EnvVarName = "VSECM_SENTINEL_SECRET_GENERATION_PREFIX"
const VSecMSafeStoreWorkloadSecretAsK8sSecretPrefix = "VSECM_SAFE_STORE_WORKLOAD_SECRET_AS_K8S_SECRET_PREFIX"
const VSecMSentinelInitCommandPath EnvVarName = "VSECM_SENTINEL_INIT_COMMAND_PATH"
const VSecMSentinelInitCommandWaitBeforeExec EnvVarName = "VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC"
const VSecMSentinelInitCommandWaitAfterInitComplete EnvVarName = "VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE"
const VSecMSentinelOidcProviderBaseUrl EnvVarName = "VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL"
const VSecMSentinelEnableOidcResourceServer EnvVarName = "VSECM_SENTINEL_ENABLE_OIDC_RESOURCE_SERVER"
const VSecMSidecarSecretsPath EnvVarName = "VSECM_SIDECAR_SECRETS_PATH"
const SpiffeEndpointSocket EnvVarName = "SPIFFE_ENDPOINT_SOCKET"
const SpiffeTrustDomain EnvVarName = "SPIFFE_TRUST_DOMAIN"
const VSecMSafeRootKeyStore EnvVarName = "VSECM_SAFE_ROOT_KEY_STORE"
const VSecMSafeBackingStore EnvVarName = "VSECM_SAFE_BACKING_STORE"
const VSecMSafeSyncSecretsInterval EnvVarName = "VSECM_SAFE_SYNC_SECRETS_INTERVAL"
const VSecMSafeSyncDeletedSecrets EnvVarName = "VSECM_SAFE_SYNC_DELETED_SECRETS"
const VSecMSafeSyncInterpolatedK8sSecrets EnvVarName = "VSECM_SAFE_SYNC_INTERPOLATED_K8S_SECRETS"
const VSecMSafeSyncExpiredSecrets EnvVarName = "VSECM_SAFE_SYNC_EXPIRED_SECRETS"
const VSecMSafeTlsPort EnvVarName = "VSECM_SAFE_TLS_PORT"

type EnvVarValue string

const VSecMSpiffeIdPrefixSafeDefault EnvVarValue = "spiffe://vsecm.com/workload/vsecm-safe/ns/vsecm-system/sa/vsecm-safe/n/"
const VSecMSpiffeIdPrefixWorkloadDefault EnvVarValue = "spiffe://vsecm.com/workload/"
const VSecMSpiffeIdPrefixSentinelDefault EnvVarValue = "spiffe://vsecm.com/workload/vsecm-sentinel/ns/vsecm-system/sa/vsecm-sentinel/n/"
const VSecMBackoffMaxRetriesDefault EnvVarValue = "10"
const VSecMBackoffDelayDefault EnvVarValue = "1000"
const VSecMBackoffMaxWaitDefault EnvVarValue = "30000"
const VSecMSafeEndpointUrlDefault EnvVarValue = "https://vsecm-safe.vsecm-system.svc.cluster.local:8443/"
const VSecMInitContainerPollIntervalDefault EnvVarValue = "5000"
const VSecMInitContainerWaitBeforeExitDefault EnvVarValue = "0"
const VSecMKeygenRootKeyPathDefault EnvVarValue = "/opt/vsecm/keys.txt"
const VSecMKeygenExportedSecretPathDefault EnvVarValue = "/opt/vsecm/secrets.json"
const VSecMSidecarMaxPollIntervalDefault EnvVarValue = "300000"
const VSecMSidecarExponentialBackoffMultiplierDefault EnvVarValue = "2"
const VSecMSidecarSuccessThresholdDefault EnvVarValue = "3"
const VSecMSidecarErrorThresholdDefault EnvVarValue = "3"
const VSecMSidecarPollIntervalDefault EnvVarValue = "20000"
const VSecMProbeLivenessPortDefault EnvVarValue = ":8081"
const VSecMProbeReadinessPortDefault EnvVarValue = ":8082"
const VSecMSafeIvInitializationIntervalDefault EnvVarValue = "50"
const VSecMSafeSecretBufferSizeDefault EnvVarValue = "10"
const VSecMSafeK8sSecretBufferSizeDefault EnvVarValue = "10"
const VSecMSafeSecretDeleteBufferSizeDefault EnvVarValue = "10"
const VSecMSafeSecretBackupCountDefault EnvVarValue = "3"
const VSecMSafeDataPathDefault EnvVarValue = "/var/local/vsecm/data"
const VSecMRootKeyPathDefault EnvVarValue = "/key/key.txt"
const VSecMSafeSourceAcquisitionTimeoutDefault EnvVarValue = "10000"
const VSecMSafeBootstrapTimeoutDefault EnvVarValue = "300000"
const VSecMRootKeyNameDefault EnvVarValue = "vsecm-root-key"
const VSecMSentinelSecretGenerationPrefixDefault EnvVarValue = "gen:"
const VSecMSafeStoreWorkloadSecretAsK8sSecretPrefixDefault EnvVarValue = "k8s:"
const VSecMSentinelInitCommandPathDefault EnvVarValue = "/opt/vsecm-sentinel/init/data"
const VSecMSentinelInitCommandWaitBeforeExecDefault EnvVarValue = "0"
const VSecMSentinelInitCommandWaitAfterInitCompleteDefault EnvVarValue = "0"
const VSecMSidecarSecretsPathDefault EnvVarValue = "/opt/vsecm/secrets.json"
const SpiffeEndpointSocketDefault EnvVarValue = "unix:///spire-agent-socket/agent.sock"
const SpiffeTrustDomainDefault EnvVarValue = "vsecm.com"
const VSecMSafeTlsPortDefault EnvVarValue = ":8443"

type Namespace string

const VSecMSystem Namespace = "vsecm-system"
const SpireSystem Namespace = "spire-system"
const SpireServer Namespace = "spire-server"

func Never(s string) bool {
	return "never" == strings.ToLower(strings.TrimSpace(s))
}

func True(s string) bool {
	return "true" == strings.ToLower(strings.TrimSpace(s))
}

func GetEnv(name EnvVarName) string {
	return os.Getenv(string(name))
}

type FieldName string

const RootKeyText FieldName = "KEY_TXT"
