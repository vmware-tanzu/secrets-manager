+++
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title = "VSecM SDK"
weight = 3
+++

## SDK

This is the documentation for [VMware Secrets Manager Go SDK][go-sdk].

You can also [check out the Go Docs on `pkg.go.dev` here][go-docs].

[go-docs]: https://pkg.go.dev/github.com/vmware-tanzu/secrets-manager/sdk
[go-sdk]: https://github.com/vmware-tanzu/secrets-manager/tree/main/sdk

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
[**VSecM Safe**][vsecm-safe] regularly. It periodically calls `Fetch()`
behind the scenes to get its work done. Once it fetches the secrets,
it saves them to the location defined in the `VSECM_SIDECAR_SECRETS_PATH`
environment variable (*`/opt/vsecm/secrets.json` by default*).

[vsecm-safe]: https://github.com/vmware-tanzu/secrets-manager/tree/main/app/safe

## Usage Example

Here is a demo workload that uses the `Fetch()` API to retrieve secrets from
**VSecM Safe**.

```go
package main

import (
  "fmt"
  "github.com/vmware-tanzu/secrets-manager/sdk/sentry"
  "time"
)

func main() {
  for {
    // Fetch the secret bound to this workload
    // using VMware Secrets Manager Go SDK:
    data, err := sentry.Fetch()

    if err != nil {
      fmt.Println("Failed. Will retry...")
    } else {
      fmt.Println("secret: '", data, "'")
    }

    time.Sleep(5 * time.Second)
  }
}
```

Here follows a possible Deployment descriptor for such a workload.

Check out [VMware Secrets Manager demo workload manifests][demos] for additional examples.

[demos]: https://github.com/vmware-tanzu/secrets-manager/tree/main/examples/using_sdk_go/k8s "Demo Workloads"

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
  className: vsecm
  spiffeIDTemplate: "spiffe://vsecm.com/workload/example"
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
          image: vsecm/example-using-sdk-go:latest
          volumeMounts:
          - name: spire-agent-socket
            mountPath: /spire-agent-socket
            readOnly: true
          env:
          - name: SPIFFE_ENDPOINT_SOCKET
            value: unix:///spire-agent-socket/spire-agent.sock
      volumes:
      - name: spire-agent-socket
        hostPath:
          path: /run/spire/sockets
          type: Directory
```

## Package `startup`

The current SDK has two public methods under the package `sentry`:

* `func Watch`

### `func Watch(waitTimeBeforeExit time.Duration)`

`Watch` continuously polls the associated secret of the workload to exist.
If the secret exists, and it is not empty, the function exits the process
container with a success status code (`0`).

This is especially useful when used inside an init container.

#### Parameters

* `waitTimeBeforeExit`: The duration to wait before a successful exit from
  the function.

----

You can also [check out the relevant sections of the
**Registering Secrets** article][registering-secrets] for an example of
**VMware Secrets Manager Go SDK** usage.

[registering-secrets]: @/documentation/use-cases/registering-secrets.md "Register a Secret"

{{ edit() }}
