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

title = "Using VSecM Init Container"
weight = 220
+++

{{
/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/
}}

## Situation Analysis

In certain cases you might not have full control over the source code
of your workloads. For example, your workload can be a containerized third
party binary executable that you don't have the source code of. It might
be consuming Kubernetes `Secret`s through injected environment variables,
and the like.

Luckily, with **VSecM Init Container** you can let the workload
wait until you register secrets to it. In addition, you can create
Kubernetes `Secret`s for the workload and let it consume them as environment
variables or file mounts similar to how it did before.

## Strategy

Use **VSecM Sentinel** to register a secret; use the `k8s:` prefix in the
name of the secret to make it a Kubernetes `Secret` instead of a VSecM secret.

The workload will consume the Kubernetes `Secret` in the usual way that it
consumes Kubernetes `Secret`s, like environment variables or file mounts.

## High-Level Diagram

Open the image in a new tab to see its full-size version:

![High-Level Diagram](/assets/init-container.png "High-Level Diagram")

## Implementation

Our approach is twofold:

1. Create a Kubernetes `Secret` for the workload using **VSecM Sentinel**.
2. Let the workload wait for the secret using **VSecM Init Container**.

Let's begin.

### Cleanup

Let's remove our workload and its associated secret to start with a
clean slate:

```bash
# Remove the workload deployment:
kubectl delete deployment example
# Find the sentinel pod's name:
kubectl get po -n vsecm-system 
# Delete the secret:
kubectl exec vsecm-sentinel-778b7fdc78-86v6d -n \
  vsecm-system -- safe -w example -d
# Make sure that the secret is gone:
kubectl exec vsecm-sentinel-778b7fdc78-86v6d -n \
  vsecm-system -- safe -l
# Output:
# {"secrets":[]}
```

### Read the Source

Make sure [you examine the manifests][workload-yaml] to gain an understanding
of what kinds of entities you've deployed to your cluster.

[workload-yaml]: https://github.com/vmware-tanzu/secrets-manager/tree/main/examples/using_init_container/k8s

### Demo Workload

Here are certain important code pieces from the demo workload that we are
going to deploy soon.

The following is the main application that the workload runs:

```go
// ./examples/using_init_container/main.go

func main() {
    // ... Truncated ...
    
	for {
		fmt.Printf("My secret: '%s'.\n", os.Getenv("SECRET"))
		fmt.Printf("My creds: username:'%s' password:'%s'.\n",
			os.Getenv("USERNAME"), os.Getenv("PASSWORD"),
		)
		fmt.Println("")

		time.Sleep(5 * time.Second)
	}
}
```

As you see, the code tries to parse several environment variables. But,
where does it get them?

For that let's look into the [`Deployment.yaml`][deployment-yaml] manifest:

```yaml
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
        image: vsecm/example-using-init-container:latest

# ... Truncated  ... 
```

[deployment-yaml]: https://github.com/vmware-tanzu/secrets-manager/blob/main/examples/using_init_container/k8s/Deployment.yaml
[secret-yaml]: https://github.com/vmware-tanzu/secrets-manager/blob/main/examples/using_init_container/k8s/Secret.yaml

### Deploy the Demo Workload

To begin, let's deploy our demo workload:

```bash 
# Switch to the project folder:
cd $WORKSPACE/secrets-manager
# Deploy the demo workload:
# Install the workload:
make example-init-container-deploy
# If you are building from the source, 
# use `make example-init-container-deploy-local` instead.
```

When we list the pods, you'll see that it's not ready yet because
**VSecM Init Container** is waiting for a secret to be registered to this pod.

```bash
kubectl get po 

NAME                                  READY   STATUS
example-5d8c6c4865-dlt8r   0/1     Init:0/1   0
```

Here are the containers in that [`Deployment.yaml`][deployment-yaml]

```yaml
      containers:
      - name: main
        image: vsecm/example-using-init-container:latest
      
      # ... Truncated  ... 
      
      initContainers:
      - name: init-container
        image: vsecm/vsecm-ist-init-container:latest
```

It's the `init-container` that waits until the workload acquires a secret.

### Registering Secrets to the Workload

To make the init container exit successfully and initialize the main
container of the Pod, execute the following script:

```bash
# ./examples/using_init_container/register.sh

# Find a Sentinel node.
SENTINEL=$(kubectl get po -n vsecm-system \
  | grep "vsecm-sentinel-" | awk '{print $1}')

# Execute the command needed to interpolate the secret.
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
-w "k8s:example-secret" \
-n "default" \
-s '{"username": "root", \
  "password": "SuperSecret", "value": "VSecMRocks"}' \
-t '{"USERNAME":"{{.username}}", \
  "PASSWORD":"{{.password}}", "VALUE": "{{.value}}"}'
```

Here are the meanings of the parameters in the above command:

* `-w` is the name of the workload.
* `-n` identifies the namespace of the Kubernetes `Secret`.
* `-s` is the actual value of the secret.
* `-t` is the template that will transform this secret to the `data` portion 
  of the Kubernetes `Secret`.

### Initializing the Workload

We have created a Kubernetes `Secret`.

Next, we will register a dummy secret to the workload using **VSecM Sentinel**.
This will make the init container exit successfully and let the main container
of the workload initialize.

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "example" \
  -s "trigger" \
  -n default
```

> **Why Do We Need a Dummy Secret?**
> 
> This use case assumes that we do not have access to the source code of the
> workload. The workload is designed to consume secrets as environment variables.
> 
> If we had access to the source code, we could have modified it to consume
> secrets from [**VSecM SDK**][vsecm-sdk] or [**VSecM Sidecar**][vsecm-sidecar].
> In that case, we would not need to register a dummy secret to the workload.

[vsecm-sdk]: @/documentation/use-cases/vsecm-sdk.md
[vsecm-sidecar]: @/documentation/use-cases/sidecar.md

Now, let's check if our pod has initialized:

```bash 
kubectl get po

NAME                                 READY   STATUS
example-5d8c6c4865-dlt8r   1/1     Running   0
```

It looks like it did. So, let's check its logs:

```bash
kubectl logs example-5d8c6c4865-dlt8r

My secret: 'VSecMRocks'.
My creds: username:'root' password:'SuperSecret'.

My secret: 'VSecMRocks'.
My creds: username:'root' password:'SuperSecret'.

My secret: 'VSecMRocks'.
My creds: username:'root' password:'SuperSecret'.
```

## What Happened?

In summary, the `Pod` that your `Deployment` manages will not initialize until
you register secrets to your workload.

Once you register secrets using the above command, **VSecM Init Container** will
exit with a success status code and let the main container initialize.

Here is a sequence diagram of how the secret is transformed (*open the image
in a new tab for a larger version*):

![Transforming Secrets](/assets/transform-secrets.png "Transforming Secrets")

## Conclusion

That's how you can register secrets as environment variables to workloads and
halt bootstrapping of the main container until the secrets are registered to
the workload.

This approach is marginally **less** secure, because it creates interim secrets
which are not strictly necessary if we were to use **VSecM Sidecar** or
**VSecM Safe**. It is meant to be used for **legacy** systems where directly
using the **Safe Sidecar** or **Safe SDK** are not feasible.

For example, you might not have direct control over the source code to enable a
tighter **Safe** integration. Or, you might temporarily want to establish
behavior parity of your legacy system before starting a more canonical **VMware Secrets Manager**
implementation.

For modern workloads that you have more control, we highly encourage you to
use [**VSecM SDK**][tutorial-sdk] or [**VSecM Sidecar**][tutorial-sidecar] instead.

That being said, it's good to have this option, and we are sure you can find
other creative ways to leverage it too.

[tutorial-sdk]: @/documentation/use-cases/vsecm-sdk.md
[tutorial-sidecar]: @/documentation/use-cases/sidecar.md

## List of Use Cases in This Section

{{ use_cases() }}

{{ edit() }}
