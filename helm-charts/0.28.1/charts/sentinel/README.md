# sentinel

![Version: 0.28.1](https://img.shields.io/badge/Version-0.28.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.28.1](https://img.shields.io/badge/AppVersion-0.28.1-informational?style=flat-square)

Helm chart for sentinel

## Values

| Key                        | Type      | Default | Description |
|----------------------------|-----------|---------|-------------|
| environments               | list      | `[{"name":"SPIFFE_ENDPOINT_SOCKET","value":"unix:///spire-agent-socket/spire-agent.sock"},{"name":"VSECM_BACKOFF_DELAY","value":"1000"},{"name":"VSECM_BACKOFF_MAX_RETRIES","value":"10"},{"name":"VSECM_BACKOFF_MAX_WAIT","value":"10000"},{"name":"VSECM_BACKOFF_MODE","value":"exponential"},{"name":"VSECM_LOG_LEVEL","value":"7"},{"name":"VSECM_LOG_SECRET_FINGERPRINTS","value":"false"},{"name":"VSECM_PROBE_LIVENESS_PORT","value":":8081"},{"name":"VSECM_SENTINEL_OIDC_ENABLE_RESOURCE_SERVER","value":"false"},{"name":"VSECM_SENTINEL_INIT_COMMAND_PATH","value":"/opt/vsecm-sentinel/init/data"},{"name":"VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE","value":"0"},{"name":"VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC","value":"0"},{"name":"VSECM_SENTINEL_LOGGER_URL","value":"localhost:50051"},{"name":"VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL","value":"http://0.0.0.0:8080/auth/realms/XXXXX/protocol/openid-connect/token/introspect"},{"name":"VSECM_SENTINEL_SECRET_GENERATION_PREFIX","value":"gen:"}]` | See https://vsecm.com/configuration for more information about these environment variables. |
| environments[0]            | object    | `{"name":"SPIFFE_ENDPOINT_SOCKET","value":"unix:///spire-agent-socket/spire-agent.sock"}` | The SPIFFE endpoint socket. This is used to communicate with the SPIRE. The name of the socket should match spireAgent.socketName in values.yaml of the SPIRE chart. |
| environments[10]           | object    | `{"name":"VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE","value":"0"}` | The amount of time to wait (in milliseconds) after all initialization commands are executed. |
| environments[11]           | object    | `{"name":"VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC","value":"0"}` | The amount of time to wait (in milliseconds) before executing the initialization commands. |
| environments[12]           | object    | `{"name":"VSECM_SENTINEL_LOGGER_URL","value":"localhost:50051"}` | VSecM Sentinel uses a gRPC logger to log audit events. This is the URL of the gRPC logger. |
| environments[13]           | object    | `{"name":"VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL","value":"http://0.0.0.0:8080/auth/realms/XXXXX/protocol/openid-connect/token/introspect"}` | The OIDC provider's base URL. This is the URL that VSecM Sentinel will use to introspect the token. |
| environments[14]           | object    | `{"name":"VSECM_SENTINEL_SECRET_GENERATION_PREFIX","value":"gen:"}` | The prefix to hint to generate secrets randomly based on regex-like patterns. |
| environments[1]            | object    | `{"name":"VSECM_BACKOFF_DELAY","value":"1000"}` | The interval between retries (in milliseconds) for the default backoff strategy. |
| environments[2]            | object    | `{"name":"VSECM_BACKOFF_MAX_RETRIES","value":"10"}` | The maximum number of retries for the default backoff strategy before it gives up. |
| environments[3]            | object    | `{"name":"VSECM_BACKOFF_MAX_WAIT","value":"10000"}` | The maximum wait time (in milliseconds) for the default backoff strategy. |
| environments[4]            | object    | `{"name":"VSECM_BACKOFF_MODE","value":"exponential"}` | The backoff mode. The default is "exponential". Allowed values: "exponential", "linear" |
| environments[5]            | object    | `{"name":"VSECM_LOG_LEVEL","value":"7"}` | The log level. 0: Logs are off (only audit events will be logged), 7: TRACE level logging (maximum verbosity). |
| environments[6]            | object    | `{"name":"VSECM_LOG_SECRET_FINGERPRINTS","value":"false"}` | Useful for debugging. This will log cryptographic fingerprints of secrets without revealing the secret itself. It is recommended to keep this "false" in production. |
| environments[7]            | object    | `{"name":"VSECM_PROBE_LIVENESS_PORT","value":":8081"}` | The port that the liveness probe listens on. |
| environments[8]            | object    | `{"name":"VSECM_SENTINEL_OIDC_ENABLE_RESOURCE_SERVER","value":"false"}` | Enable or disable OIDC resource server. When enabled, VSecM Sentinel will act as an OIDC resource server. Note that exposing VSecM Sentinel's functionality through a server significantly alters the attack surface, and the decision should be considered carefully. This option will create a RESTful API around VSecM Sentinel. Since VSecM Sentinel is the main entry point to the system, the server's security is important. Ideally, do not expose this server to the public Internet and protect it with tight security controls. |
| environments[9]            | object    | `{"name":"VSECM_SENTINEL_INIT_COMMAND_PATH","value":"/opt/vsecm-sentinel/init/data"}` | The path where the initialization commands are mounted. |
| fullnameOverride           | string    | `""` | The fullname override of the chart. |
| imagePullSecrets           | list      | `[]` |  |
| initCommand                | object    | `{"command":"exit:true\n--\n","enabled":true}` | The custom initialization commands that will be executed by the VSecM Sentinel during its initial bootstrapping. The commands are executed in the order they are provided. See the official documentation for more information: https://vsecm.com/configuration |
| initCommand.enabled        | bool      | `true` | Specifies whether the custom initialization commands are enabled. If set to 'false', the custom initialization commands will not be executed. |
| livenessPort               | int       | `8081` | The port that the liveness probe listens on. |
| nameOverride               | string    | `""` | The name override of the chart. |
| podAnnotations             | object    | `{}` | Additional pod annotations. |
| podSecurityContext         | object    | `{}` | Pod security context overrides. |
| resources.requests.cpu     | string    | `"5m"` |  |
| resources.requests.memory  | string    | `"20Mi"` |  |
| serviceAccount             | object    | `{"annotations":{},"create":true,"name":"vsecm-sentinel"}` | The service account to use. |
| serviceAccount.annotations | object    | `{}` | Annotations to add to the service account |
| serviceAccount.create      | bool      | `true` | Specifies whether a service account should be created |
| serviceAccount.name        | string    | `"vsecm-sentinel"` | The name of the service account to use. If not set and create is true, a name is generated using the fullname template |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.13.1](https://github.com/norwoodj/helm-docs/releases/v1.13.1)
