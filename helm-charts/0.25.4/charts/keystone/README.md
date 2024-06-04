# keystone

![Version: 0.25.4](https://img.shields.io/badge/Version-0.25.4-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.25.4](https://img.shields.io/badge/AppVersion-0.25.4-informational?style=flat-square)

Helm chart for keystone

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
| environments[2].name | string | `"VSECM_WORKLOAD_SPIFFEID_PREFIX"` |  |
| environments[2].value | string | `"spiffe://vsecm.com/workload/"` |  |
| environments[3].name | string | `"VSECM_SAFE_TLS_PORT"` |  |
| environments[3].value | string | `":8443"` |  |
| fullnameOverride | string | `""` |  |
| imagePullSecrets | list | `[]` |  |
| initEnvironments[0].name | string | `"SPIFFE_ENDPOINT_SOCKET"` |  |
| initEnvironments[0].value | string | `"unix:///spire-agent-socket/agent.sock"` |  |
| initEnvironments[1].name | string | `"VSECM_LOG_LEVEL"` |  |
| initEnvironments[1].value | string | `"7"` |  |
| initEnvironments[2].name | string | `"VSECM_WORKLOAD_SPIFFEID_PREFIX"` |  |
| initEnvironments[2].value | string | `"spiffe://vsecm.com/workload/"` |  |
| initEnvironments[3].name | string | `"VSECM_INIT_CONTAINER_POLL_INTERVAL"` |  |
| initEnvironments[3].value | string | `"5000"` |  |
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
| serviceAccount.name | string | `"vsecm-keystone"` |  |

