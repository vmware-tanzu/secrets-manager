---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, installation, deployment, faq, quickstart
title: Using Aegis Go SDK
description: directly consume the <strong>Aegis Safe</strong> API
micro_nav: true
page_nav:
  prev:
    content: local development
    url: '/docs/contributing'
  next:
    content: <strong>Aegis</strong> Sentinel CLI
    url: '/docs/sentinel'
---

<p style="text-align:right;position:relative;top:-40px;"
><a href="https://github.com/ShieldWorks/aegis-web/blob/main/docs/sdk.md"
style="border-bottom: none;background:#e0e0e0;padding:0.5em;display:inline-block;
border-radius:8px;">
edit this page on <strong>GitHub</strong> ✏️</a></p>

This is the documentation for [Aegis Go SDK][go-sdk].

[go-sdk]: https://github.com/shieldworks/aegis/tree/main/sdk

## Package `sentry`

The current SDK has two public methods under the package `sentry`:

* `func Fetch`
* `func Watch`

### `func Fetch() (string, error)`

`Fetch` fetches the up-to-date secret that has been registered to the workload.

```go
secret, err := sentry.Fetch()
```

In case of a problem, `Fetch` will return an empty string and an error 
explaining what went wrong.


### `func Watch()`

`Watch` synchronizes the internal state of the workload by talking to 
[**Aegis Safe**][aegis-safe] regularly. It periodically calls `Fetch()` 
behind the scenes to get its work done. Once it fetches the secrets, 
it saves them to the location defined in the `AEGIS_SIDECAR_SECRETS_PATH` 
environment variable (*`/opt/aegis/secrets.json` by default*).

[aegis-safe]: https://github.com/shieldworks/aegis-safe

## Usage Example

Here is a demo workload that uses the `Fetch()` API to retrieve secrets from 
**Aegis Safe**.

```go
package main

import (
	"fmt"
	"github.com/shieldworks/aegis-sdk-go/sentry"
	"time"
)

func main() {
	for {
		// Fetch the secret bound to this workload
		// using Aegis Go SDK:
		data, err := sentry.Fetch()

		if err != nil {
			fmt.Println("Failed. Will retry…")
		} else {
			fmt.Println("secret: '", data, "'")
		}

		time.Sleep(5 * time.Second)
	}
}
```

Here follows a possible Deployment descriptor for such a workload. 

Check out [Aegis demo workload manifests][demos] for additional examples.

[demos]: https://github.com/shieldworks/aegis/tree/main/install/k8s/demo-workload "Demo Workloads"

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: example
  namespace: default
automountServiceAccountToken: false
---
apiVersion: spire.spiffe.io/v1alpha1
kind: ClusterSPIFFEID
metadata:
  name: example
spec:
  spiffeIDTemplate: "spiffe://aegis.ist/workload/example"
  podSelector:
    matchLabels:
      app.kubernetes.io/name: example
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example
  namespace: default
  labels:
    app.kubernetes.io/name: example
spec:
  replicas: 1
  selector:
    matchLabels:
      app.kubernetes.io/name: example
  template:
    metadata:
      labels:
        app.kubernetes.io/name: example
    spec:
      serviceAccountName: example
      containers:
        - name: main
          image: aegishub/example-using-sdk:0.7.0
          volumeMounts:
          - name: spire-agent-socket
            mountPath: /spire-agent-socket
            readOnly: true
          env:
          - name: SPIFFE_ENDPOINT_SOCKET
            value: unix:///spire-agent-socket/agent.sock
      volumes:
      - name: spire-agent-socket
        hostPath:
          path: /run/spire/sockets
          type: Directory
```

You can also [check out the relevant sections of the 
**Registering Secrets** article][registering-secrets] for an example of 
**Aegis Go SDK** usage.

[registering-secrets]: /docs/register

