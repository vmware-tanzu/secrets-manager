---
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title: Using VSecM Sidecar
layout: post
next_url: /docs/use-case-sdk/
prev_url: /docs/use-cases-overview/
permalink: /docs/use-case-sidecar/
---

## Using With **Aegis Sidecar**

Let’s deploy our demo workload that will use **Aegis Sidecar**.

You can find the deployment manifests inside the
[`./examples/workload-using-sidecar/k8s`][workload-yaml] folder of your
cloned **Aegis** folder.

[workload-yaml]: https://github.com/shieldworks/aegis/tree/main/examples/workload-using-sidecar/k8s

To deploy our workload using that manifest, execute the following:

```bash
# Switch to the Aegis repo:
cd $WORKSPACE/aegis
# Install the workload:
make example-sidecar-deploy
# If you are building from the source, 
# use `make example-sidecar-deploy-local` instead.
```

And that’s it. You have your demo workload up and running.

## Read the Source

Make sure [you examine the manifests][workload-yaml] to gain an understanding
of what kinds of entities you’ve deployed to your cluster.

You’ll see that there are two images in the `Deployment` object declared inside
that folder:

* `aegishub/example`: This is the container that has the business logic.
* `aegishub/aegis-sidecar`: This **Aegis**-managed container injects
  secrets to a place that our demo container can consume.

## The Demo App

[Here is the source code of the demo container’s app][workload-src] for the
sake of completeness.

[workload-src]: https://github.com/shieldworks/aegis/blob/main/examples/workload-using-sidecar/main.go

When you check the source code, you’ll see that our demo app tries to read a
secret file every 5 seconds forever:

```go 
for {
    dat, err := os.ReadFile(sidecarSecretsPath())
    if err != nil {
        fmt.Println("Failed to read the secrets file. Will retry in 5 seconds…")
        fmt.Println(err.Error())
    } else {
        fmt.Println("secret: '", string(dat), "'")
    }

    time.Sleep(5 * time.Second)
}
```

## ClusterSPIFFEID

Yet, how do we tell **Aegis** about our app so that it can identify it to
deliver secrets?

For this, there is an identity file that defines a `ClusterSPIFFEID` for
the workload:

```yaml
# ./examples/workload-using-sidecar/k8s/Identity.yaml

{% raw %}kind: ClusterSPIFFEID
metadata:
  name: example
spec:
  # SPIFFE ID `MUST` start with "spiffe://aegis.ist/workload/$workloadName/ns/"
  # for `safe` to recognize the workload and dispatch secrets to it.
  spiffeIDTemplate: "spiffe://aegis.ist\
    /workload/example\
    /ns/{{ .PodMeta.Namespace }}\
    /sa/{{ .PodSpec.ServiceAccountName }}\
    /n/{{ .PodMeta.Name }}"
  podSelector:
    matchLabels:
      app.kubernetes.io/name: example
  workloadSelectorTemplates:
    - "k8s:ns:default"
    - "k8s:sa:example"{% endraw %}
```

This identity descriptor, tells **Aegis** that the workload:

* Lives under a certain namespace,
* Is bound to a certain service account,
* And as a certain name.

When the time comes, **Aegis** will read this identity and learn about which
workload is requesting secrets. Then it can decide to deliver
the secrets (*because the workload is registered*) or deny dispatching them
(*because the workload is unknown/unregistered*).

> **ClusterSPIFFEID is an Abstraction**
>
> Please note that `Identity.yaml` is not a random YAML file:
> It is a binding contract and abstracts a host of operations
> behind the scenes.
>
> For every `ClusterSPIFFEID` created this way,
> `SPIRE` (*Aegis’ identity control plane*) will deliver an **X.509 SVID**
> bundle to the workload.
>
> Therefore, creating a `ClusterSPIFFEID` is a way to **irrefutably**,
> **securely**, and **cryptographically** identify a workload.

## Verifying the Deployment

If you have been following along so far, when you execute `kubectl get po` will
give you something like this:

```bash 
{% raw %}kubectl get po

NAME                                  READY   STATUS    RESTARTS   AGE
example-5d564458b6-vsmtm  2/2     Running   0          9s{% endraw %}
```

Let’s check the logs of our pod:

```bash 
{% raw %}kubectl logs example-5d564458b6-vsmtm -f

Failed to read the secrets file. Will retry in 5 seconds…
open /opt/aegis/secrets.json: no such file or directory
Failed to read the secrets file. Will retry in 5 seconds…
open /opt/aegis/secrets.json: no such file or directory
Failed to read the secrets file. Will retry in 5 seconds…

…{% endraw %}
```

What we see here that our workload checks for the secrets file and cannot
find it for a while, and displays a failure message.

## Registering a Secret

Let’s register a secret and see how the logs change:

```bash 
{% raw %}# Find the name of the Aegis Sentinel pod.
kubectl get po -n aegis-system

# register a secret to our workload using Aegis Sentinel
kubectl exec aegis-sentinel-778b7fdc78-86v6d -n aegis-system -- aegis \
  -w "example" \
  -s "AegisRocks!"
  
# Response: 
# OK{% endraw %}
```

Now let’s check the logs again:

```bash 
{% raw %}kubectl logs example-5d564458b6-vsmtm -f

secret: ' AegisRocks! '
secret: ' AegisRocks! '
secret: ' AegisRocks! '
secret: ' AegisRocks! '

…{% endraw %}
```

So we registered our first secret to a workload using **Aegis Sentinel**.
The secret is stored in **Aegis Safe** and dispatched to the workload
through **Aegis Sidecar** behind the scenes.

> **What Is Aegis Sentinel**?
>
> For all practical purposes, you can think of **Aegis Sentinel** as the
> “*bastion host*” you log in and execute sensitive operations.
>
> In our case, we will register secrets to workloads using it.

## Registering Multiple Secrets

If needed, you can associate more than one secret to a worklad, for this, you’ll
need to use the `-a` (for “*append*”) flag.

```bash 
{% raw %}kubectl exec aegis-sentinel-778b7fdc78-86v6d -n aegis-system -- aegis \
  -w "example" \
  -s "AegisRocks!" \
  -a
  
# Response:
# OK
  
kubectl exec aegis-sentinel-778b7fdc78-86v6d -n aegis-system -- aegis \
  -w "example" \
  -s "YouRockToo!" \
  -a
  
# Response: 
# OK{% endraw %}
```

Now, let’s check our logs:

```bash
k logs example-5d564458b6-sx9sj -f

secret: ' ["YouRockToo!","AegisRocks!"] '
secret: ' ["YouRockToo!","AegisRocks!"] '
secret: ' ["YouRockToo!","AegisRocks!"] '
secret: ' ["YouRockToo!","AegisRocks!"] '
````

Yes, we have two secrets in an array.

**Aegis Safe** returns a single string if there is a single secret associated
with the workload, and a JSON Array of strings if the workload has more than
one secret registered.

## More About ClusterSPIFFEID

Let’s dig a bit deeper.

[`ClusterSPIFFEID`][clusterspiffeid] is a Kubernetes Custom Resource that enables distributing
[**SPIRE**](https://spiffe.io/) identities to workloads in a cloud-native
and declarative way.

Assuming you’ve had a chance to review the deployment manifests as recommended
at the start of this tutorial, you might have noticed something similar to what’s
presented below in the [`Identity.yaml`][identity-yaml]."

[identity-yaml]: https://github.com/shieldworks/aegis/blob/main/examples/workload-using-sidecar/k8s/Identity.yaml
[clusterspiffeid]: https://github.com/spiffe/spire-controller-manager/blob/main/docs/clusterspiffeid-crd.md

```text
{% raw %}spiffeIDTemplate: "spiffe://aegis.ist\
  /workload/example\
  /ns/{{ .PodMeta.Namespace }}\
  /sa/{{ .PodSpec.ServiceAccountName }}\
  /n/{{ .PodMeta.Name }}"{% endraw %}
```

The `example` part from that template is the **name** that **Aegis**
will identify this workload as. That is the name we used when we registered
the secret to our workload.

## **Aegis Sentinel** Commands

You can execute
`kubectl exec -it $sentinelPod -n aegis-sytem -- aegis --help`
for a list of all available commands and command-line flags
that **Aegis Sentinel** has.

Also, [Check out **Aegis Sentinel CLI Reference**][sentinel-ref] for more
information and usage examples on **Aegis Sentinel**.

[sentinel-ref]: /docs/sentinel

## Conclusion

Yay 🎉. That was our first secret.

In the next tutorial, we will do something similar; however, this time we
will leverage [**Aegis SDK**](/docs/sdk) instead of **Aegis Sidecar**.