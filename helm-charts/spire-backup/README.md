# spire

![Version: 0.26.1](https://img.shields.io/badge/Version-0.26.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.26.1](https://img.shields.io/badge/AppVersion-0.26.1-informational?style=flat-square)

Helm chart for spire

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| autoscaling | object | `{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80}` | Autoscaling settings. Note that for autoscaling to work, you need to have a proper setup for the SPIRE Server database. Check out the official documentation for more information: https://spiffe.io/docs/latest/setup/ |
| bundleEndpoint | object | `{"annotations":{},"port":8443,"type":"ClusterIP"}` | Service details for the SPIRE Server Bundle Endpoint. The bundle endpoint is typically used for federating  |
| bundleEndpoint.annotations | object | `{}` | Additional Service annotations. |
| bundleEndpoint.port | int | `8443` | The port that the bundle endpoint serves. |
| bundleEndpoint.type | string | `"ClusterIP"` | Valid values are: ClusterIP, NodePort, LoadBalancer |
| data | object | `{"persistent":false,"persistentVolumeClaim":{"accessMode":"ReadWriteOnce","size":"1Gi","storageClass":""}}` | Persistence settings for the SPIRE Server. |
| data.persistent | bool | `false` | Persistence is disabled by default. You are recommended to provide a persistent volume. |
| data.persistentVolumeClaim | object | `{"accessMode":"ReadWriteOnce","size":"1Gi","storageClass":""}` | Define the PVC if `persistent` is true. |
| experimental | object | `{"eventsBasedCache":false}` | Experimental settings. |
| experimental.eventsBasedCache | bool | `false` | eventsBasedCache is known to significantly improve SPIRE Server performance. It is set to `false` by default, just in case. |
| fullnameOverride | string | `""` | The fullname override of the chart. |
| imagePullSecrets | list | `[]` | Override it with an image pull secret that you need as follows: imagePullSecrets:  - name: my-registry-secret |
| nameOverride | string | `""` | The name override of the chart. |
| replicaCount | int | `1` | replicaCount SPIRE server currently runs with a sqlite database. Scaling to multiple instances will not work until we use an external database. |
| resources | object | `{"agent":{"requests":{"cpu":"50m","memory":"512Mi"}},"server":{"requests":{"cpu":"100m","memory":"1Gi"}},"spiffeCsiDriver":{"requests":{"cpu":"50m","memory":"128Mi"}}}` | These are the default resources suitable for a moderate SPIRE usage. Of course, it's best to do your own benchmarks and update these requests and limits to your production needs accordingly. That being said, as a rule of thumb, do not limit the CPU request on SPIRE Agent and SPIRE server. It's best to let them leverage the available excess CPU, if available. |
| resources.agent | object | `{"requests":{"cpu":"50m","memory":"512Mi"}}` | SPIRE Agent resource requests and limits. |
| resources.server | object | `{"requests":{"cpu":"100m","memory":"1Gi"}}` | SPIRE Server resource requests and limits. |
| resources.spiffeCsiDriver | object | `{"requests":{"cpu":"50m","memory":"128Mi"}}` | SPIFFE CSI Driver resource requests and limits. |
| service | object | `{"annotations":{},"type":"ClusterIP"}` | Service details for the SPIRE Server. |
| service.annotations | object | `{}` | Additional Service annotations. |
| service.type | string | `"ClusterIP"` | Service type. Possible values are: ClusterIP, NodePort, LoadBalancer. Defaults to `ClusterIP`. |
| serviceAccount | object | `{"annotations":{},"create":true,"name":""}` | Service Account details for the SPIRE Server. |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account |
| serviceAccount.create | bool | `true` | Specifies whether a service account should be created |
| serviceAccount.name | string | `""` | The name of the service account to use. If not set and create is true, a name is generated using the fullname template. |
| spireAgent | object | `{"agentSocketDir":"/tmp/spire-agent/public","annotations":{"helm.sh/hook":"post-install","helm.sh/hook-delete-policy":"hook-succeeded"},"hostSocketDir":"/var/run/spire/sockets","socketName":"spire-agent.sock"}` | SPIRE Agent settings. |
| spireAgent.agentSocketDir | string | `"/tmp/spire-agent/public"` | The SPIRE Agent socket directory. |
| spireAgent.annotations | object | `{"helm.sh/hook":"post-install","helm.sh/hook-delete-policy":"hook-succeeded"}` | Annotations to add to the SPIRE Agent pod. |
| spireAgent.annotations."helm.sh/hook" | string | `"post-install"` | Define a helm hook to make spire-agent daemonSet deploy after spire-server statefulSet. |
| spireAgent.annotations."helm.sh/hook-delete-policy" | string | `"hook-succeeded"` | Define the policy to delete the hook after it has succeeded. |
| spireAgent.hostSocketDir | string | `"/var/run/spire/sockets"` | The corresponding SPIRE Agent socket directory on the host. |
| spireAgent.socketName | string | `"spire-agent.sock"` | The SPIRE Agent socket name. |
| spireServer | object | `{"configDir":"/run/spire/server/config","containerPort":"8081","dataDir":"/run/spire/data","privateSocketDir":"/tmp/spire-server/private"}` | SPIRE Server settings. |
| spireServer.configDir | string | `"/run/spire/server/config"` | The configuration directory for the SPIRE Server. |
| spireServer.containerPort | string | `"8081"` | The internal port that the server serves. |
| spireServer.dataDir | string | `"/run/spire/data"` | The data directory for the SPIRE Server. |
| spireServer.privateSocketDir | string | `"/tmp/spire-server/private"` | The private socket directory for the SPIRE Server. |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.13.1](https://github.com/norwoodj/helm-docs/releases/v1.13.1)
