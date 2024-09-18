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

title = "Mounting Secrets as Volumes"
weight = 241
+++

## Situation Analysis

Certain apps may require their secrets to be mounted in "*in-memory*" volumes. 
Depending on the app, this could be the only way the app consumes secrets and 
configuration information.

When this is the only way to configure the app, especially when we don't have 
access to the application's configuration and source code, there won't be any 
other feasible way to pass secrets to the app.

## Screencast

Here is a screencast that demonstrates this use case:

```txt
WORK IN PROGRESS
```

## Strategy

Use **VSecM Sidecar** and **VSecM Init Container** to provide the secrets the 
workload needs when needed.

## High-Level Diagram

Open the image in a new tab to see the full-size version:

![High-Level Diagram](/assets/mount-to-volume.png "High-Level Diagram")

## Implementation

We will assume our workload has the name `example`, deployed to the `example-apps` 
namespace, and is associated with the `example-sa` service account, 
having a label `example-app`.

### Initial Application

Here is a sample deployment manifest of such a workload:

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
```

### Create a `ClusterSPIFFEID`

For **VSecM** to communicate with this workload, a `ClusterSPIFFEID` is needed.

Here is how such a ClusterSPIFFEID may look like:

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

When this `ClusterSPIFFEID` is defined, our example pod will get a secure 
x.509 certificate from the [SPIFFE Workload API][workload-api] to talk to VSecM; 
however, to fetch the certificate, we will need to modify the pod's deployment 
manifest slightly. 

[workload-api]: https://github.com/spiffe/spiffe/blob/main/standards/SPIFFE_Workload_API.md "SPIFFE Workload API"

Let's see that in the next section.

### Let the Workload Consume SPIFFE Workload API

Here is the modified deployment manifest to consume SPIFFE Workload API:

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

## <--- BEGIN CHANGE
      volumes:
        - name: spire-agent-socket
          csi:
            driver: "csi.spiffe.io"
            readOnly: true
## <-- END CHANGE
```

We added a special volume, and the [SPIFFE CSI Driver][spiffe-csi-drive] 
will handle the rest of the communication.

[spiffe-csi-driver]: https://github.com/spiffe/spiffe-csi "SPIFFE CSI Driver"

Here is the manifest without the change markers:

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
      volumes:
        - name: spire-agent-socket
          csi:
            driver: "csi.spiffe.io"
            readOnly: true
```

Let's say this workload needs an `/opt/app/credentials` file as an initial 
configuration file to execute its business logic. Based on this assumption, 
let's update the manifest accordingly to provide this file:

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

## <--- BEGIN CHANGE
        volumeMounts:
        - name: credentials-volume
          mountPath: /opt/app/credentials
          subPath: credentials
          readOnly: true
## <--- END CHANGE

      volumes:
      - name: spire-agent-socket
        csi:
          driver: "csi.spiffe.io"
          readOnly: true

## <--- BEGIN CHANGE
      - name: credentials-volume
        emptyDir:
          medium: Memory
## <--- END CHANGE
```

However, the volume is *empty*; our app will likely require it to be populated 
before it can be used. We will address this real soon.

Here's the manifest without the change markers:

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
        - name: credentials-volume
          mountPath: /opt/app/credentials
          subPath: credentials
          readOnly: true
      volumes:
      - name: spire-agent-socket
        csi:
          driver: "csi.spiffe.io"
          readOnly: true
      - name: credentials-volume
        emptyDir:
          medium: Memory
```

We will populate the volume using **VSecM Sidecar**.

### Initializing the Volume Using VSecM Sidecar

Let's add **VSecM Sidecar** to our manifest:

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
        - name: credentials-volume
          mountPath: /opt/app/credentials
          readOnly: true

## <--- BEGIN CHANGE
      - name: sidecar
        image: vsecm/vsecm-ist-sidecar:0.24.1
        volumeMounts:
        - mountPath: /opt/app/credentials
          name: credentials-volume
        - name: spire-agent-socket
          mountPath: /spire-agent-socket
          readOnly: true
        env:
        - name: VSECM_SIDECAR_SECRET_PATH
          value: "/opt/app/credentials/secrets.json"  
        - name: SPIFFE_ENDPOINT_SOCKET
          value: "unix:///spire-agent-socket/spire-agent.sock"
## <--- END CHANGE

      volumes:
      - name: spire-agent-socket
        csi:
          driver: "csi.spiffe.io"
          readOnly: true
      - name: credentials-volume
        emptyDir:
          medium: Memory
```

In this setup, **VSecM Sidecar** will periodically poll **VSecM Safe** to fetch 
the secret associated with the workload and update `/opt/app/credentials`.

> **Help Needed**
>
> **VSecM Sidecar** currently creates a single file. If the application needs 
> more than one file in the volume, you'll need to create a specialized sidecar 
> based on VSecM Sidecar.
> 
> There are upstream issues that will enable **VSecM Sidecar** to parse the 
> incoming secret and create separate files in its associated volume. 
> 
> If you need this functionality, you are welcome to contribute upstream.
> 
> * [VSecM Init Container should be (optionally cofigurable) able to watch the 
>   changes in a volume or a file instead of VSecM Safe `#759`][ticket-759]
> * [VSecM Sidecar should be able to (optionally) create one file per 
>   secret key `#761`][ticket-761]


[ticket-759]: https://github.com/vmware-tanzu/secrets-manager/issues/759
[ticket-761]: https://github.com/vmware-tanzu/secrets-manager/issues/761

Here's the YAML manifest without change markers:

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
        - name: credentials-volume
          mountPath: /opt/app/credentials
          readOnly: true
      - name: sidecar
        image: vsecm/vsecm-ist-sidecar:0.24.1
        volumeMounts:
        - mountPath: /opt/app/credentials
          name: credentials-volume
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

But now, we have another problem: The main application likely assumes the 
credentials are already there when it starts its lifecycle. The app will likely 
crash if the credentials are not there during its bootstrapping.

To fix this, we'll need an init container that watches this volume and initializes 
the main app only after it is populated. We will implement this in the following 
section.

### Adding an Init Container to Wait for The Volume to Populate

Let's add an init container to complete our plan:

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

## <--- BEGIN CHANGE
      initContainers:
      - name: vsecm-init-container
        image: vsecm/vsecm-ist-init-container:latest
        volumeMounts:
        - name: credentials-volume
          mountPath: /opt/vsecm/secrets
          readOnly: true
## <--- END CHANGE

      containers:
      - name: example-container
        image: example-app:0.1.0
        volumeMounts:
        - name: credentials-volume
          mountPath: /opt/app/credentials
          readOnly: true
      - name: sidecar
        image: vsecm/vsecm-ist-sidecar:0.24.1
        volumeMounts:
        - mountPath: /opt/app/credentials
          name: credentials-volume
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

**VSecM Init Container** will now watch the volume and only initialize the main 
app when everything is ready.

> **Help Needed**
>
> Enhancing the VSecM Init Container to monitor volume changes is an outstanding 
> upstream issue. Your contributions are welcome.
> 
> * [VSecM Init Container should be (optionally configurable) able to watch the 
    changes in a volume or a file instead of VSecM Safe `#759`][ticket-759]
>
> Until this issue is resolved, you might need to add a second init container to 
> introduce some delay (*~10-15 seconds, on average*) for **VSecM Sidecar** to 
> consume the secret and create the file needed.


[ticket-759]: https://github.com/vmware-tanzu/secrets-manager/issues/759

### Providing Randomized Secrets to the Workload

VSecM Helm Charts can be configured to provide secure pattern-based random secrets 
for workloads. For example, we can modify [the initCommand stanza][init-command] 
of the sentinel Helm chart's `values.yaml` to register a random username and 
password as secrets for our example workload.

```yaml
# ./charts/sentinel/values.yaml

initCommand:
  enabled: true
  command: |
    --
    w:example
    n:example-apps
    s:gen:{"username":"admin-[a-z0-9]{6}","password":"[a-zA-Z0-9]{12}"}
    t:{"ADMIN_USER":"{{.username}}","ADMIN_PASS":"{{.password}}"}
    --
```

[init-command]: https://github.com/vmware-tanzu/secrets-manager/commit/cbd3b33d9f50569d7b14034a7551ccc6d6575ed7#diff-f4d0a065a7104f6d7271039d0c61dd6ff5155e7a86ca98570c49ecbd48947122R91-R120

## Conclusion

Integrating **VSecM Sidecar** and **VSecM Init Container** for secret management 
represents a forward-thinking solution to the challenge of securely providing 
secrets to Kubernetes workloads. 

This approach embodies the principles of zero trust security by ensuring that 
secrets are dynamically managed and securely injected, thereby minimizing the 
risk of exposure. As Kubernetes continues to evolve, such patterns will be 
pivotal in addressing the complex security needs of cloud-native applications, 
making contributions and enhancements in this space invaluable.

Here are some key points covered in this use case:

* **Strategy and Implementation**: The approach involves using a sidecar and 
  init container pattern to inject secrets securely into a Kubernetes pod.d. The 
  example provided demonstrates how to deploy a workload with this pattern, 
  highlighting the importance of security in modern cloud-native applications.
* **Secure Secret Management**: By leveraging [`ClusterSPIFFEID`s][clusterspiffeid] 
  and the [SPIFFE Workload API][spiffe-workload-api], the system ensures secure 
  communication between the workload and **VSecM**. This method enhances security 
  by providing each pod with a unique identity and secure certificates for 
  authentication.
* **Dynamic Secret Injection**: Introducing a sidecar container allows for dynamic 
  secret fetching and updating, ensuring that workloads can access the most current 
  secrets without restarting pods. This approach mainly benefits applications 
  requiring high security and frequent secret rotations.
* **Initialization and Secret Readiness**: An init container ensures that the 
  main application starts only after the secrets are properly fetched and mounted, 
  addressing the challenge of secret readiness at application startup. This is 
  critical for applications that expect specific configurations to be present 
  before initialization.
* **Customization and Contribution**: The article also suggests customizing the 
  sidecar for applications needing multiple secrets in separate files and invites 
  contributions to enhance the **VSecM Sidecar**'s functionality.

[clusterspiffeid]: https://github.com/spiffe/spire-controller-manager/blob/main/docs/clusterspiffeid-crd.md
[spiffe-workload-api]: https://github.com/spiffe/spiffe/blob/main/standards/SPIFFE_Workload_API.md

## List of Use Cases in This Section

{{ use_cases() }}

{{ edit() }}
