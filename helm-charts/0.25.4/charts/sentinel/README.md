# sentinel

![Version: 0.25.4](https://img.shields.io/badge/Version-0.25.4-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.25.4](https://img.shields.io/badge/AppVersion-0.25.4-informational?style=flat-square)

Helm chart for sentinel

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| autoscaling.enabled | bool | `false` |  |
| autoscaling.maxReplicas | int | `100` |  |
| autoscaling.minReplicas | int | `1` |  |
| autoscaling.targetCPUUtilizationPercentage | int | `80` |  |
| environments[0].name | string | `"SPIFFE_ENDPOINT_SOCKET"` |  |
| environments[0].value | string | `"unix:///spire-agent-socket/agent.sock"` |  |
| environments[1].name | string | `"VSECM_LOG_LEVEL"` |  |
| environments[1].value | string | `"7"` |  |
| environments[2].name | string | `"VSECM_PROBE_LIVENESS_PORT"` |  |
| environments[2].value | string | `":8081"` |  |
| environments[3].name | string | `"VSECM_SAFE_TLS_PORT"` |  |
| environments[3].value | string | `":8443"` |  |
| environments[4].name | string | `"VSECM_SENTINEL_INIT_COMMAND_PATH"` |  |
| environments[4].value | string | `"/opt/vsecm-sentinel/init/data"` |  |
| environments[5].name | string | `"VSECM_SENTINEL_LOGGER_URL"` |  |
| environments[5].value | string | `"localhost:50051"` |  |
| environments[6].name | string | `"VSECM_SENTINEL_SECRET_GENERATION_PREFIX"` |  |
| environments[6].value | string | `"gen:"` |  |
| environments[7].name | string | `"VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL"` |  |
| environments[7].value | string | `"http://0.0.0.0:8080/auth/realms/XXXXX/protocol/openid-connect/token/introspect"` |  |
| environments[8].name | string | `"VSECM_SENTINEL_ENABLE_OIDC_RESOURCE_SERVER"` |  |
| environments[8].value | string | `"false"` |  |
| fullnameOverride | string | `""` |  |
| imagePullSecrets | list | `[]` |  |
| initCommand.command | string | `"exit:true\n--\n"` |  |
| initCommand.enabled | bool | `true` |  |
| livenessPort | int | `8081` |  |
| nameOverride | string | `""` |  |
| podAnnotations | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| replicaCount | int | `1` |  |
| resources.requests.cpu | string | `"5m"` |  |
| resources.requests.memory | string | `"20Mi"` |  |
| securityContext | object | `{}` |  |
| serviceAccount.annotations | object | `{}` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `"vsecm-sentinel"` |  |

