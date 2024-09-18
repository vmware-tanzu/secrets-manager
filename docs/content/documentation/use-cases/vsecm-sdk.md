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

title = "Retrieving Secrets Via VSecM SDK"
weight = 243
+++

## Situation Analysis

If you are the creator of an app, and you have access to its source code, it would 
be beneficial to retrieve the secrets the app needs whenever it needs them through 
the [**VSecM SDK**][sdk].

[sdk]: @/documentation/usage/sdk.md "VSecM SDK"

Fetching secrets using VSecM SDK will also enable you to…

* Rotate the app's secrets without needing to restart or evict the app.
* Get meta-information about the secrets that are otherwise inaccessible.

## Screencast

Here is a screencast that demonstrates this use case:

```txt
WORK IN PROGRESS
```

## High-Level Diagram

Open the images in a new tab to see the full-size versions:

![High-Level Diagram](/assets/using-sdk.png "High-Level Diagram")

![High-Level Diagram](/assets/using-vsecm-sdk.png "High-Level Diagram")

## Implementation

We'll define the `ClusterSPIFFEID` and "*SPIRE Agent Socket*" for our workload, 
similar to the [Mounting Secrets as Volumes][secrets-as-volumes] use case.

[secrets-as-volumes]: @/documentation/use-cases/mounting-secrets.md "Mounting Secrets as Volumes"

### Prepare Kubernetes Manifests

Here's the `ClusterSPIFFEID`:

```yaml 
apiVersion: spire.spiffe.io/v1alpha1
kind: ClusterSPIFFEID
metadata:
  name: example
spec:
  className: vsecm
  spiffeIDTemplate: "spiffe://vsecm.com\
    /workload/example\
    /ns/{{ .PodMeta.Namespace }}\
    /sa/{{ .PodSpec.ServiceAccountName }}\
    /n/{{ .PodMeta.Name }}"
  podSelector:
    matchLabels:
      app: example-app
  workloadSelectorTemplates:
    - "k8s:ns:example-apps"
    - "k8s:sa:example-sa"
```

Here's our deployment manifest:

```yaml 
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example
  namespace: example-apps
  labels:
    app: example
spec:
  replicas: 1
  selector:
    matchLabels:
      app: example-app
  template:
    metadata:
      labels:
        app: example-app
    spec:
      serviceAccountName: example-sa
      containers:
      - name: example-container
        image: example-app:0.1.0
        volumeMounts:
        - name: spire-agent-socket
          mountPath: /spire-agent-socket
          readOnly: true
        env:
        - name: VSECM_SIDECAR_SECRET_PATH
          value: "/opt/app/credentials/secrets.json"
        - name: SPIFFE_ENDPOINT_SOCKET
          value: "unix:///spire-agent-socket/spire-agent.sock"
      volumes:
      - name: spire-agent-socket
        csi:
          driver: "csi.spiffe.io"
          readOnly: true
      - name: credentials-volume
        emptyDir:
          medium: Memory
```

### Application Code

Since we have access to source code, our application can directly fetch its 
secrets as follows:

```go
package main

import (
	"fmt"
	"github.com/vmware-tanzu/secrets-manager/sdk/sentry" // <- SDK
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	d, err := sentry.Fetch()

	if err != nil {
		fmt.Println("Failed to read the secrets file. Try again.")
		fmt.Println(err.Error())
		return
	}

	if d.Data == "" {
		fmt.Println("no secret yet... Check again later.")
		return
	}

	fmt.Printf(
		"secret: updated: %s, created: %s, value: %s\n",
		d.Updated, d.Created, d.Data,
	)
}
```

> **Want an SDK in Your Favorite Language?**
>
> VMware Secrets Manager only has an official Go SDK at this time.
> If you have application code in another language, you can contribute to the 
> following SDK initiatives:
> 
> * [VSecM Java SDK][sdk-java]
> * [VSecM C++ SDK][sdk-cpp]
> * [VSecM Rust SDK][sdk-rust]
> * [VSecM Python SDK][sdk-python]


### The Benefit of Using **VSecM SDK**

**VSecM SDK** gives direct control of **VSecM Safe** to your workload.

The advantage of this approach is: you are in charge.
The downside of it is: Well, you are in charge 🙂.

But, jokes aside, your application will have to be
more tightly bound to **VMware Secrets Manager** without a sidecar.

However, when you use a sidecar, your application does not have any idea of
**VMware Secrets Manager**'s existence. From its perspective, it is merely
reading from a file that something magically updates every once in a while. This
"*separation of concerns*" can make your application architecture more
adaptable to changes.

As in anything, there is no one true way to do it. Your approach will depend
on your project's requirements.

## Conclusion

Integrating **VSecM Go SDK** into application development workflows offers a 
robust and dynamic approach to handling secrets management. By leveraging the 
capabilities of VSecM SDK, developers can ensure that their applications have 
secure and efficient access to necessary secrets, facilitating smoother operations 
and bolstering security. 

The ability to rotate secrets without impacting the application's availability 
and access meta-information about the secrets are critical benefits that can 
significantly enhance the security posture of any application.

The **VSecM SDK**, while currently officially supporting only Go, highlights a 
significant moment of growth and opportunity within the developer community. 
This moment isn't a limitation but a clarion call to action.

Such collaboration not only diversifies the toolkit available to developers but 
also strengthens the bonds within the community, fostering an environment where 
innovation thrives on the principles of security and efficiency.

By contributing to the development of the **VSecM SDK** in various programming 
languages, you are not merely coding; you are pioneering a movement towards a 
more secure, efficient, and community-driven approach in application development. 

[sdk-java]: https://github.com/vmware-tanzu/secrets-manager/issues/448
[sdk-cpp]: https://github.com/vmware-tanzu/secrets-manager/issues/450
[sdk-rust]: https://github.com/vmware-tanzu/secrets-manager/issues/556
[sdk-python]: https://github.com/vmware-tanzu/secrets-manager/issues/505

## List of Use Cases in This Section

{{ use_cases() }}

{{ edit() }}
