# spire

![Version: 0.27.2](https://img.shields.io/badge/Version-0.27.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.27.2](https://img.shields.io/badge/AppVersion-0.27.2-informational?style=flat-square)

Helm chart for spire

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| data | object | `{"persistent":true,"persistentVolumeClaim":{"accessMode":"ReadWriteOnce","size":"1Gi","storageClass":""}}` | Persistence settings for the SPIRE Server. |
| data.persistent | bool | `true` | Persistence is enabled by default. However, you are recommended to provide your own storage class if you are using a cloud provider or a storage solution that supports dynamic provisioning. |
| data.persistentVolumeClaim | object | `{"accessMode":"ReadWriteOnce","size":"1Gi","storageClass":""}` | Define the PVC if `persistent` is true. |
| enableSpireMintedDefaultClusterSpiffeIds | bool | `false` | SPIRE assigns a default Cluster SPIFFE ID to all workloads in the cluster. The SPIFFEID SPIRE assigns by default is not aligned with the SPIFFE ID format that VSecM Safe expects. Also, you might not want SPIRE to assign SPIFFE IDs to every single workload you have in your cluster if you are not using SPIRE to attest those workloads. Therefore, this option is set to false by default.  If you set this to true, make sure you update `safeSpiffeIdTemplate` `sentinelSpiffeIdTemplate`, `keystoneSpiffeIdTemplate`, `workloadNameRegExp`, `workloadSpiffeIdPrefix`, `safeSpiffeIdPrefix`, `sentinelSpiffeIdPrefix` and other relevant configurations to match with what SPIRE assigns. |
| experimental | object | `{"eventsBasedCache":false}` | Experimental settings. |
| experimental.eventsBasedCache | bool | `false` | eventsBasedCache is known to significantly improve SPIRE Server performance. It is set to `false` by default, just in case. |
| fullnameOverride | string | `""` | The fullname override of the chart. |
| imagePullSecrets | list | `[]` | Override it with an image pull secret that you need as follows: imagePullSecrets:  - name: my-registry-secret |
| nameOverride | string | `""` | The name override of the chart. |
| resources | object | `{"agent":{"requests":{"cpu":"50m","memory":"512Mi"}},"server":{"requests":{"cpu":"100m","memory":"1Gi"}},"spiffeCsiDriver":{"requests":{"cpu":"50m","memory":"128Mi"}}}` | These are the default resources suitable for a moderate SPIRE usage. Of course, it's best to do your own benchmarks and update these requests and limits to your production needs accordingly. That being said, as a rule of thumb, do not limit the CPU request on SPIRE Agent and SPIRE server. It's best to let them leverage the available excess CPU, if available. |
| resources.agent | object | `{"requests":{"cpu":"50m","memory":"512Mi"}}` | SPIRE Agent resource requests and limits. |
| resources.server | object | `{"requests":{"cpu":"100m","memory":"1Gi"}}` | SPIRE Server resource requests and limits. |
| resources.spiffeCsiDriver | object | `{"requests":{"cpu":"50m","memory":"128Mi"}}` | SPIFFE CSI Driver resource requests and limits. |
| spireAgent | object | `{"hostSocketDir":"/run/spire/agent-sockets","internalAdminSocketDir":"/tmp/spire-agent/private","internalPublicSocketDir":"/tmp/spire-agent/public","socketName":"spire-agent.sock"}` | SPIRE Agent settings. |
| spireAgent.hostSocketDir | string | `"/run/spire/agent-sockets"` | The corresponding SPIRE Agent socket directory on the host. SPIRE Agents and SPIFFE CSI Driver shares this directory. |
| spireAgent.internalAdminSocketDir | string | `"/tmp/spire-agent/private"` | The corresponding SPIRE Agent internal admin directory in the container. The configuration should match the SPIRE Agent configuration and SPIRE Agent DaemonSet. You are advised not to change this value. |
| spireAgent.internalPublicSocketDir | string | `"/tmp/spire-agent/public"` | The corresponding SPIRE Agent internal socket directory in the container. The configuration should match the SPIRE Agent configuration and SPIRE Agent DaemonSet. |
| spireAgent.socketName | string | `"spire-agent.sock"` | The SPIRE Agent socket name. |
| spireServer | object | `{"configDir":"/run/spire/config","dataDir":"/run/spire/data","privateSocketDir":"/tmp/spire-server/private","service":{"type":"ClusterIP"}}` | SPIRE Server settings. |
| spireServer.configDir | string | `"/run/spire/config"` | The configuration directory for the SPIRE Server. |
| spireServer.dataDir | string | `"/run/spire/data"` | The data directory for the SPIRE Server. SPIRE Server’s ConfigMap and StatefulSet should agree on this directory. |
| spireServer.privateSocketDir | string | `"/tmp/spire-server/private"` | The private socket directory for the SPIRE Server. SPIRE Server’s ConfigMap and StatefulSet should agree on this directory. |
| spireServer.service | object | `{"type":"ClusterIP"}` | Service details for the SPIRE Server. |
| spireServer.service.type | string | `"ClusterIP"` | Service type. Possible values are: ClusterIP, NodePort, LoadBalancer. Defaults to `ClusterIP`. |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.13.1](https://github.com/norwoodj/helm-docs/releases/v1.13.1)
