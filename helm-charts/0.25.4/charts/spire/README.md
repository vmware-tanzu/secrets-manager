# spire

![Version: 0.25.4](https://img.shields.io/badge/Version-0.25.4-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.25.4](https://img.shields.io/badge/AppVersion-0.25.4-informational?style=flat-square)

Helm chart for spire

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| autoscaling.enabled | bool | `false` |  |
| autoscaling.maxReplicas | int | `100` |  |
| autoscaling.minReplicas | int | `1` |  |
| autoscaling.targetCPUUtilizationPercentage | int | `80` |  |
| bundleEndpoint.annotations | object | `{}` |  |
| bundleEndpoint.port | int | `8443` |  |
| bundleEndpoint.type | string | `"ClusterIP"` |  |
| data.persistent | bool | `false` |  |
| data.persistentVolumeClaim.accessMode | string | `"ReadWriteOnce"` |  |
| data.persistentVolumeClaim.size | string | `"1Gi"` |  |
| data.persistentVolumeClaim.storageClass | string | `""` |  |
| experimental.eventsBasedCache | bool | `false` |  |
| fullnameOverride | string | `""` |  |
| imagePullSecrets | list | `[]` |  |
| nameOverride | string | `""` |  |
| podAnnotations | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| replicaCount | int | `1` |  |
| resources.agent.requests.cpu | string | `"50m"` |  |
| resources.agent.requests.memory | string | `"512Mi"` |  |
| resources.server.requests.cpu | string | `"100m"` |  |
| resources.server.requests.memory | string | `"1Gi"` |  |
| resources.spiffeCsiDriver.requests.cpu | string | `"50m"` |  |
| resources.spiffeCsiDriver.requests.memory | string | `"128Mi"` |  |
| securityContext | object | `{}` |  |
| service.annotations | object | `{}` |  |
| service.port | int | `8081` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.annotations | object | `{}` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.name | string | `""` |  |
| spireAgent.annotations."helm.sh/hook" | string | `"post-install"` |  |
| spireAgent.annotations."helm.sh/hook-delete-policy" | string | `"hook-succeeded"` |  |

