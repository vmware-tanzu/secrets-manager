# keystone

![Version: 0.25.4](https://img.shields.io/badge/Version-0.25.4-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.25.4](https://img.shields.io/badge/AppVersion-0.25.4-informational?style=flat-square)

Helm chart for keystone

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| autoscaling | object | `{"enabled":false,"maxReplicas":100,"minReplicas":1,"targetCPUUtilizationPercentage":80}` | Autoscaling settings. Note that, by default, autoscaling is disabled. It does not typically make sense to autoscale VSecM Keystone as it is a control plane component with minimal resource requirements. |
| environments | list | `[{"name":"VSECM_LOG_LEVEL","value":"7"}]` | See https://vsecm.com/configuration for more information about these environment variables. |
| environments[0] | object | `{"name":"VSECM_LOG_LEVEL","value":"7"}` | The log level. 0: Logs are off (only audit events will be logged) 7: TRACE level logging (maximum verbosity). |
| fullnameOverride | string | `""` | The fullname override of the chart. |
| imagePullSecrets | list | `[]` | Override it with an image pull secret that you need as follows: imagePullSecrets:  - name: my-registry-secret |
| initEnvironments | list | `[{"name":"SPIFFE_ENDPOINT_SOCKET","value":"unix:///spire-agent-socket/agent.sock"},{"name":"VSECM_LOG_LEVEL","value":"7"},{"name":"VSECM_INIT_CONTAINER_POLL_INTERVAL","value":"5000"}]` | See https://vsecm.com/configuration for more information about these environment variables. |
| initEnvironments[0] | object | `{"name":"SPIFFE_ENDPOINT_SOCKET","value":"unix:///spire-agent-socket/agent.sock"}` | The SPIFFE endpoint socket. This is used to communicate with the SPIRE  agent. If you change this, you will need to change the associated  volumeMount in the Deployment.yaml too. |
| initEnvironments[1] | object | `{"name":"VSECM_LOG_LEVEL","value":"7"}` | The log level. 0: Logs are off (only audit events will be logged) 7: TRACE level logging (maximum verbosity). |
| initEnvironments[2] | object | `{"name":"VSECM_INIT_CONTAINER_POLL_INTERVAL","value":"5000"}` | The interval (in milliseconds) that the VSecM Init Container will poll the VSecM Safe for secrets. |
| livenessPort | int | `8081` |  |
| nameOverride | string | `""` | The name override of the chart. |
| podAnnotations | object | `{}` | Additional pod annotations. |
| podSecurityContext | object | `{}` | Pod security context overrides. |
| replicaCount | int | `1` |  |
| resources | object | `{"requests":{"cpu":"5m","memory":"20Mi"}}` | Resource limits and requests. |
| serviceAccount | object | `{"annotations":{},"create":true,"name":"vsecm-keystone"}` | The service account to use. |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account. |
| serviceAccount.create | bool | `true` | Specifies whether a service account should be created. |
| serviceAccount.name | string | `"vsecm-keystone"` | The name of the service account to use. If not set and 'create' is true, a name is generated using the fullname template. |

