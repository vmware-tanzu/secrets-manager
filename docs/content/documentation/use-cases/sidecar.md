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

title = "Using VSecM Sidecar"
weight = 200
+++

## Situation Analysis

There might be times you don't have direct control over the application code
to integrate **VMware Secrets Manager** SDK. In such cases, you can use
**VSecM Sidecar** to inject secrets into your application.

**VSceM Sidecar** is a container that runs alongside your application container
and injects secrets into a shared in-memory volume. Your application can then
read the secrets from that volume.

## Strategy

Use **VSecM Sidecar** to inject secrets into your application container.

## High-Level Diagram

Open the image in a new tab to see its full-size version:

![High-Level Diagram](/assets/vsecm-sidecar.png "High-Level Diagram")

## Implementation

Let's deploy our demo workload that will use **VSecM Sidecar**.

You can find the deployment manifests inside the
[`./examples/workload-using-sidecar/k8s`][workload-yaml] folder of your
cloned **VMware Secrets Manager** project.

[workload-yaml]: https://github.com/vmware-tanzu/secrets-manager/tree/main/examples/using_sidecar/k8s

### Deploying the Example Workload

To deploy our workload using that manifest, execute the following:

```bash
# Switch to the VSecM repo:
cd $WORKSPACE/secrets-manager
# Install the workload:
make example-sidecar-deploy
# If you are building from the source, 
# use `make example-sidecar-deploy-local` instead.
```

And that's it. You have your demo workload up and running.

### Read the Source

Make sure [you examine the manifests][workload-yaml] to gain an understanding
of what kinds of entities you've deployed to your cluster.

You'll see that there are two images in the `Deployment` object declared inside
that folder:

* `vsecm/example`: This is the container that has the business logic.
* `vsecm/vsecm-ist-sidecar`: This **VMware Secrets Manager**-managed container 
  injects secrets to a place that our demo container can consume.

### The Demo App

[Here is the source code of the demo container's app][workload-src] for the
sake of completeness.

[workload-src]: https://github.com/vmware-tanzu/secrets-manager/blob/main/examples/using_sidecar/main.go

When you check the source code, you'll see that our demo app tries to read a
secret file every 5 seconds forever:

```go 
for {
    dat, err := os.ReadFile(sidecarSecretsPath())
    if err != nil {
        fmt.Println("Failed to read. Will retry in 5 seconds...")
        fmt.Println(err.Error())
    } else {
        fmt.Println("secret: '", string(dat), "'")
    }

    time.Sleep(5 * time.Second)
}
```

### ClusterSPIFFEID

Yet, how do we tell **VMware Secrets Manager** about our app so that it can 
identify it to deliver secrets?

For this, there is an identity file that defines a `ClusterSPIFFEID` for
the workload:

```yaml
# ./examples/using_sidecar/k8s/Identity.yaml

kind: ClusterSPIFFEID
metadata:
  name: example
spec:
  # SPIFFE ID `MUST` start with 
  # "spiffe://vsecm.com/workload/$workloadName/ns/"
  # for `safe` to recognize the workload and 
  # dispatch secrets to it.
  spiffeIDTemplate: "spiffe://vsecm.com\
    /workload/example\
    /ns/{{ .PodMeta.Namespace }}\
    /sa/{{ .PodSpec.ServiceAccountName }}\
    /n/{{ .PodMeta.Name }}"
  podSelector:
    matchLabels:
      app.kubernetes.io/name: example
  workloadSelectorTemplates:
    - "k8s:ns:default"
    - "k8s:sa:example"
```

This identity descriptor, tells **VMware Secrets Manager** that the workload:

* Lives under a certain namespace,
* Is bound to a certain service account,
* And as a certain name.

When the time comes, **VMware Secrets Manager** will read this identity and learn
about which workload is requesting secrets. Then it can decide to deliver
the secrets (*because the workload is registered*) or deny dispatching them
(*because the workload is unknown/unregistered*).

> **ClusterSPIFFEID is an Abstraction**
>
> Please note that `Identity.yaml` is not a random YAML file:
> It is a binding contract and abstracts a host of operations
> behind the scenes.
>
> For every `ClusterSPIFFEID` created this way,
> `SPIRE` (*VSecM identity control plane*) will deliver an **X.509 SVID**
> bundle to the workload.
>
> Therefore, creating a `ClusterSPIFFEID` is a way to **irrefutably**,
> **securely**, and **cryptographically** identify a workload.

### Verifying the Deployment

If you have been following along so far, when you execute `kubectl get po` will
give you something like this:

```bash 
kubectl get po

NAME                              STATUS    AGE
example-5d564458b6-vsmtm  2/2     Running   9s
```

Let's check the logs of our pod:

```bash
kubectl logs example-5d564458b6-vsmtm -f
```

The output will be something like this:

```txt
Failed to read the secrets file. Will retry in 5 seconds...
open /opt/vsecm/secrets.json: no such file or directory
Failed to read the secrets file. Will retry in 5 seconds...
open /opt/vsecm/secrets.json: no such file or directory
Failed to read the secrets file. Will retry in 5 seconds...
...
```

What we see here that our workload checks for the secrets file and cannot
find it for a while, and displays a failure message.

### Registering a Secret

Let's register a secret and see how the logs change:

```bash 
# Find the name of the VSecM Sentinel pod.
kubectl get po -n vsecm-system

# register a secret to our workload using VSecM Sentinel
kubectl exec vsecm-sentinel-778b7fdc78-86v6d -n vsecm-system \
  -- safe \
  -w "example" \
  -s "VSecMRocks!"
  
# Response: 
# OK
```

Now let's check the logs again:

```bash 
kubectl logs example-5d564458b6-vsmtm -f

secret: ' VSecMRocks! '
secret: ' VSecMRocks! '
secret: ' VSecMRocks! '
secret: ' VSecMRocks! '

...
```

So we registered our first secret to a workload using **VSecM Sentinel**.
The secret is stored in **VSecM Safe** and dispatched to the workload
through **VSecM Sidecar** behind the scenes.

> **What Is VSecM Sentinel**?
>
> For all practical purposes, you can think of **VSecM Sentinel** as the
> "*bastion host*" you log in and execute sensitive operations.
>
> In our case, we will register secrets to workloads using it.

### Registering Multiple Secrets

If needed, you can associate more than one secret to a workload, for this, you'll
need to use the `-a` (for "*append*") flag.

```bash 
kubectl exec vsecm-sentinel-778b7fdc78-86v6d -n vsecm-system \
  -- safe \
  -w "example" \
  -s "VSecMRocks!" \
  -a
  
# Response:
# OK
  
kubectl exec vsecm-sentinel-778b7fdc78-86v6d -n vsecm-system \
  -- safe \
  -w "example" \
  -s "YouRockToo!" \
  -a
  
# Response: 
# OK
```

Now, let's check our logs:

```bash
k logs example-5d564458b6-sx9sj -f

secret: ' ["YouRockToo!","VSecMRocks!"] '
secret: ' ["YouRockToo!","VSecMRocks!"] '
secret: ' ["YouRockToo!","VSecMRocks!"] '
secret: ' ["YouRockToo!","VSecMRocks!"] '
````

Yes, we have two secrets in an array.

**VSecM Safe** returns a single string if there is a single secret associated
with the workload, and a JSON Array of strings if the workload has more than
one secret registered.

### More About `ClusterSPIFFEID`s

Let's dig a bit deeper.

[`ClusterSPIFFEID`][clusterspiffeid] is a Kubernetes Custom Resource that enables 
distributing [**SPIRE**](https://spiffe.io/) identities to workloads in a cloud-native
and declarative way.

Assuming you've had a chance to review the deployment manifests as recommended
at the start of this tutorial, you might have noticed something similar to what's
presented below in the [`Identity.yaml`][identity-yaml].

[identity-yaml]: https://github.com/vmware-tanzu/secrets-manager/blob/main/examples/using_sidecar/k8s/Identity.yaml
[clusterspiffeid]: https://github.com/spiffe/spire-controller-manager/blob/main/docs/clusterspiffeid-crd.md

```txt
spiffeIDTemplate: "spiffe://vsecm.com\
  /workload/example\
  /ns/{{ .PodMeta.Namespace }}\
  /sa/{{ .PodSpec.ServiceAccountName }}\
  /n/{{ .PodMeta.Name }}"
```

The `example` part from that template is the **name** that **VMware Secrets Manager**
will identify this workload as. That is the name we used when we registered
the secret to our workload.

### **VSecM Sentinel** Commands

You can execute
`kubectl exec -it $sentinelPod -n vsecm-system -- safe --help`
for a list of all available commands and command-line flags
that **VSecM Sentinel** has.

Also, [Check out **VSecM Sentinel CLI Reference**][sentinel-ref] for more
information and usage examples on **VSecM Sentinel**.

[sentinel-ref]: @/documentation/usage/cli.md

## Conclusion

**VSecM Sidecar** presents a robust solution for securely managing and injecting 
secrets into application containers without directly modifying the application code. 
This approach leverages the sidecar pattern, effectively decoupling the secret 
management from the application's business logic, thereby enhancing security 
and maintainability.

This approach not only simplifies the integration of secret management into 
existing Kubernetes deployments but also strengthens security protocols by 
minimizing direct access to sensitive information. 

As observed, the seamless functionality of registering and appending multiple 
secrets highlights the flexibility and scalability of the 
**VMware Secrets Manager** ecosystem.

By utilizing **VSecM Sidecar**, organizations can achieve a higher level of 
security assurance and operational efficiency in managing secrets, which is 
crucial for maintaining the integrity and confidentiality of application data 
in dynamic and complex cloud environments. 

## List of Use Cases in This Section

{{ use_cases() }}

{{ edit() }}
