# safe

![Version: 0.25.4](https://img.shields.io/badge/Version-0.25.4-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.25.4](https://img.shields.io/badge/AppVersion-0.25.4-informational?style=flat-square)

Helm chart for VMware Secrets Manager (VSecM) Safe

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| autoscaling.enabled | bool | `false` |  |
| autoscaling.maxReplicas | int | `100` |  |
| autoscaling.minReplicas | int | `1` |  |
| autoscaling.targetCPUUtilizationPercentage | int | `80` |  |
| data.hostPath.path | string | `"/var/local/vsecm/data"` |  |
| data.persistent | bool | `false` |  |
| data.persistentVolumeClaim.accessMode | string | `"ReadWriteOnce"` |  |
| data.persistentVolumeClaim.size | string | `"1Gi"` |  |
| data.persistentVolumeClaim.storageClass | string | `""` |  |
| environments[0].name | string | `"SPIFFE_ENDPOINT_SOCKET"` |  |
| environments[0].value | string | `"unix:///spire-agent-socket/agent.sock"` |  |
| environments[10].name | string | `"VSECM_SAFE_IV_INITIALIZATION_INTERVAL"` |  |
| environments[10].value | string | `"50"` |  |
| environments[11].name | string | `"VSECM_SAFE_K8S_SECRET_BUFFER_SIZE"` |  |
| environments[11].value | string | `"10"` |  |
| environments[12].name | string | `"VSECM_SAFE_K8S_SECRET_DELETE_BUFFER_SIZE"` |  |
| environments[12].value | string | `"10"` |  |
| environments[13].name | string | `"VSECM_ROOT_KEY_INPUT_MODE_MANUAL"` |  |
| environments[13].value | string | `"false"` |  |
| environments[14].name | string | `"VSECM_SAFE_SECRET_BACKUP_COUNT"` |  |
| environments[14].value | string | `"3"` |  |
| environments[15].name | string | `"VSECM_SAFE_SECRET_BUFFER_SIZE"` |  |
| environments[15].value | string | `"10"` |  |
| environments[16].name | string | `"VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE"` |  |
| environments[16].value | string | `"10"` |  |
| environments[17].name | string | `"VSECM_SAFE_SOURCE_ACQUISITION_TIMEOUT"` |  |
| environments[17].value | string | `"10000"` |  |
| environments[18].name | string | `"VSECM_SAFE_STORE_WORKLOAD_SECRET_AS_K8S_SECRET_PREFIX"` |  |
| environments[18].value | string | `"k8s:"` |  |
| environments[19].name | string | `"VSECM_SAFE_TLS_PORT"` |  |
| environments[19].value | string | `":8443"` |  |
| environments[1].name | string | `"VSECM_LOG_LEVEL"` |  |
| environments[1].value | string | `"7"` |  |
| environments[20].name | string | `"VSECM_WORKLOAD_SPIFFEID_PREFIX"` |  |
| environments[20].value | string | `"spiffe://vsecm.com/workload/"` |  |
| environments[21].name | string | `"VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL"` |  |
| environments[21].value | string | `"http://0.0.0.0:8080/auth/realms/XXXXX/protocol/openid-connect/token/introspect"` |  |
| environments[22].name | string | `"VSECM_SENTINEL_ENABLE_OIDC_RESOURCE_SERVER"` |  |
| environments[22].value | string | `"false"` |  |
| environments[2].name | string | `"VSECM_PROBE_LIVENESS_PORT"` |  |
| environments[2].value | string | `":8081"` |  |
| environments[3].name | string | `"VSECM_PROBE_READINESS_PORT"` |  |
| environments[3].value | string | `":8082"` |  |
| environments[4].name | string | `"VSECM_SAFE_BACKING_STORE"` |  |
| environments[4].value | string | `"file"` |  |
| environments[5].name | string | `"VSECM_SAFE_BOOTSTRAP_TIMEOUT"` |  |
| environments[5].value | string | `"300000"` |  |
| environments[6].name | string | `"VSECM_ROOT_KEY_NAME"` |  |
| environments[6].value | string | `"vsecm-root-key"` |  |
| environments[7].name | string | `"VSECM_ROOT_KEY_PATH"` |  |
| environments[7].value | string | `"/key/key.txt"` |  |
| environments[8].name | string | `"VSECM_SAFE_DATA_PATH"` |  |
| environments[8].value | string | `"/var/local/vsecm/data"` |  |
| environments[9].name | string | `"VSECM_SAFE_FIPS_COMPLIANT"` |  |
| environments[9].value | string | `"false"` |  |
| fullnameOverride | string | `""` |  |
| imagePullSecrets | list | `[]` |  |
| livenessPort | int | `8081` |  |
| nameOverride | string | `""` |  |
| podAnnotations | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| readinessPort | int | `8082` |  |
| replicaCount | int | `1` |  |
| resources.requests.cpu | string | `"5m"` |  |
| resources.requests.memory | string | `"20Mi"` |  |
| rootKeySecretName | string | `"vsecm-root-key"` |  |
| securityContext | object | `{}` |  |
| service.port | int | `8443` |  |
| service.targetPort | int | `8443` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.annotations | object | `{}` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `"vsecm-safe"` |  |

