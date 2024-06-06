# safe

![Version: 0.25.4](https://img.shields.io/badge/Version-0.25.4-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.25.4](https://img.shields.io/badge/AppVersion-0.25.4-informational?style=flat-square)

Helm chart for VMware Secrets Manager (VSecM) Safe

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| autoscaling | object | `{"enabled":false,"maxReplicas":10,"minReplicas":1,"targetCPUUtilizationPercentage":80}` | Autoscaling settings. Note that autoscaling is not supported for VSecM Safe yet. For proper operation there should always be a single VSecM Safe pod at all times. |
| data | object | `{"hostPath":{"path":"/var/local/vsecm/data"},"persistent":false,"persistentVolumeClaim":{"accessMode":"ReadWriteOnce","size":"1Gi","storageClass":""}}` | How persistence is handled. |
| data.hostPath | object | `{"path":"/var/local/vsecm/data"}` | hostPath if `persistent` is false. |
| data.persistent | bool | `false` | If `persistent` is true, a PersistentVolumeClaim is used. Otherwise, a hostPath is used. |
| data.persistentVolumeClaim | object | `{"accessMode":"ReadWriteOnce","size":"1Gi","storageClass":""}` | PVC settings (if `persistent` is true). |
| environments | list | `[{"name":"SPIFFE_ENDPOINT_SOCKET","value":"unix:///spire-agent-socket/agent.sock"},{"name":"VSECM_LOG_LEVEL","value":"7"},{"name":"VSECM_PROBE_LIVENESS_PORT","value":":8081"},{"name":"VSECM_PROBE_READINESS_PORT","value":":8082"},{"name":"VSECM_SAFE_BACKING_STORE","value":"file"},{"name":"VSECM_SAFE_BOOTSTRAP_TIMEOUT","value":"300000"},{"name":"VSECM_ROOT_KEY_NAME","value":"vsecm-root-key"},{"name":"VSECM_ROOT_KEY_PATH","value":"/key/key.txt"},{"name":"VSECM_SAFE_DATA_PATH","value":"/var/local/vsecm/data"},{"name":"VSECM_SAFE_FIPS_COMPLIANT","value":"false"},{"name":"VSECM_SAFE_IV_INITIALIZATION_INTERVAL","value":"50"},{"name":"VSECM_SAFE_K8S_SECRET_BUFFER_SIZE","value":"10"},{"name":"VSECM_ROOT_KEY_INPUT_MODE_MANUAL","value":"false"},{"name":"VSECM_SAFE_SECRET_BACKUP_COUNT","value":"3"},{"name":"VSECM_SAFE_SECRET_BUFFER_SIZE","value":"10"},{"name":"VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE","value":"10"},{"name":"VSECM_SAFE_SOURCE_ACQUISITION_TIMEOUT","value":"10000"},{"name":"VSECM_SAFE_STORE_WORKLOAD_SECRET_AS_K8S_SECRET_PREFIX","value":"k8s:"},{"name":"VSECM_SAFE_TLS_PORT","value":":8443"}]` | See https://vsecm.com/configuration for more information about these environment variables. |
| environments[0] | object | `{"name":"SPIFFE_ENDPOINT_SOCKET","value":"unix:///spire-agent-socket/agent.sock"}` | The SPIFFE endpoint socket. This is used to communicate with the SPIRE agent. If you change this, you will need to change the associated volumeMount in the Deployment.yaml too. |
| environments[10] | object | `{"name":"VSECM_SAFE_IV_INITIALIZATION_INTERVAL","value":"50"}` | The IV initialization interval (in milliseconds) for the VSecM Safe. |
| environments[11] | object | `{"name":"VSECM_SAFE_K8S_SECRET_BUFFER_SIZE","value":"10"}` | The number of secrets VSecM Safe can buffer before blocking further operations until the buffer has space. |
| environments[12] | object | `{"name":"VSECM_ROOT_KEY_INPUT_MODE_MANUAL","value":"false"}` | Whether to automatically generate root cryptographic material or expect it to be provided through VSecM Sentinel CLI by the operator. If set to "false", VSecM Safe will automatically generate the root keys, which will make the operator's life easier. |
| environments[13] | object | `{"name":"VSECM_SAFE_SECRET_BACKUP_COUNT","value":"3"}` | How many versions of older secrets should be kept. |
| environments[14] | object | `{"name":"VSECM_SAFE_SECRET_BUFFER_SIZE","value":"10"}` | The number of secrets VSecM Safe can buffer before blocking further operations until the buffer has space. |
| environments[15] | object | `{"name":"VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE","value":"10"}` | The number of secrets VSecM Safe can buffer before blocking further operations until the buffer has space. |
| environments[16] | object | `{"name":"VSECM_SAFE_SOURCE_ACQUISITION_TIMEOUT","value":"10000"}` | The timeout (in milliseconds) for the VSecM Safe to acquire a source. After this timeout, the VSecM Safe will bail out. |
| environments[17] | object | `{"name":"VSECM_SAFE_STORE_WORKLOAD_SECRET_AS_K8S_SECRET_PREFIX","value":"k8s:"}` | The prefix to use for the workload names, when storing workload secrets as Kubernetes secrets. |
| environments[18] | object | `{"name":"VSECM_SAFE_TLS_PORT","value":":8443"}` | The port that the VSecM Safe will listen on. |
| environments[1] | object | `{"name":"VSECM_LOG_LEVEL","value":"7"}` | The log level. 0: Logs are off (only audit events will be logged) 7: TRACE level logging (maximum verbosity). |
| environments[2] | object | `{"name":"VSECM_PROBE_LIVENESS_PORT","value":":8081"}` | The port that the liveness probe listens on. |
| environments[3] | object | `{"name":"VSECM_PROBE_READINESS_PORT","value":":8082"}` | The port that the readiness probe listens on. |
| environments[4] | object | `{"name":"VSECM_SAFE_BACKING_STORE","value":"file"}` | The backing store for VSecM Safe. Possible values are: "memory", "file", "aws-secret", "azure-secret", "gcp-secret", "k8s". Currently, only "memory" and "file" are supported. |
| environments[5] | object | `{"name":"VSECM_SAFE_BOOTSTRAP_TIMEOUT","value":"300000"}` | The interval (in milliseconds) that the VSecM Safe will wait during bootstrapping before it bails out. |
| environments[6] | object | `{"name":"VSECM_ROOT_KEY_NAME","value":"vsecm-root-key"}` | The name of the VSecM Root Key Secret. |
| environments[7] | object | `{"name":"VSECM_ROOT_KEY_PATH","value":"/key/key.txt"}` | The path where the VSecM Root Key will be mounted. |
| environments[8] | object | `{"name":"VSECM_SAFE_DATA_PATH","value":"/var/local/vsecm/data"}` | The path where the VSecM Safe will store its data (if the backing store is "file"). |
| environments[9] | object | `{"name":"VSECM_SAFE_FIPS_COMPLIANT","value":"false"}` | Should VSecM Safe use FIPS-compliant encryption? |
| fullnameOverride | string | `""` | The fullname override of the chart. |
| imagePullSecrets | list | `[]` | Override it with an image pull secret that you need as follows: imagePullSecrets:  - name: my-registry-secret |
| livenessPort | int | `8081` | The port that the liveness probe listens on. `environments.VSECM_PROBE_LIVENESS_PORT` should match this value. |
| nameOverride | string | `""` | The name override of the chart. |
| podAnnotations | object | `{}` | Additional pod annotations. |
| podSecurityContext | object | `{}` | Pod security context overrides. |
| readinessPort | int | `8082` | The port that the readiness probe listens on. `environments.VSECM_PROBE_READINESS_PORT` should match this value. |
| replicaCount | int | `1` | Number of replicas to deploy. Note that values greater than 1 are not supported yet. |
| resources | object | `{"requests":{"cpu":"5m","memory":"20Mi"}}` | Resource limits and requests. |
| rootKeySecretName | string | `"vsecm-root-key"` | The name of the root key secret. |
| service | object | `{"port":8443,"targetPort":8443,"type":"ClusterIP"}` | Service settings. |
| serviceAccount | object | `{"annotations":{},"create":true,"name":"vsecm-safe"}` | The service account to use. |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account |
| serviceAccount.create | bool | `true` | Specifies whether a service account should be created |
| serviceAccount.name | string | `"vsecm-safe"` | The name of the service account to use. If not set and create is true, a name is generated using the fullname template |

