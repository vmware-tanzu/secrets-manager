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

title = "Register a Secret"
weight = 40
+++

## Register a Secret

Now, let's register a secret and see what happens.

To register a secret we'll need to find the `vsecm-sentinel` pod in the
`vsecm-system` namespace and execute a command inside the pod.

Let's get the pod first:

```bash
kubectl get po -n vsecm-system
```

Here's a sample output:

```txt
NAME                              READY   STATUS    RESTARTS   
vsecm-keystone-c54d99d7b-c4jk4    1/1     Running   0          
vsecm-safe-6cc477f58f-x6wc9       1/1     Running   0          
vsecm-sentinel-74d648675b-8zdn2   1/1     Running   0          
```

`vsecm-sentinel-74d648675b-8zdn2` is what we need here.

Let's use it and register a secret to our example workload:

```bash
kubectl exec vsecm-sentinel-74d648675b-8zdn2 -n vsecm-system -- \
safe -w "example" -n "default" -s "VSecMRocks"
```

> **Sentinel Command Line Help**
>
> **VSecM Sentinel** comes with a command line tool
> called `safe`. `safe` allows you to register secrets to
> **VSecM Safe**, delete secrets, or list existing secrets.
>
> You can execute `safe -h` or `safe --help` to get a list of available
> commands and options.


You'll get an `OK` as a response:

```txt
OK
```

For the command `safe -w "example" -n "default" -s "VSecMRocks"`

* `-w` is the workload name
* `-n` is the namespace
* `-s` is the secret value

But how do you know what the workload name is?

That's where **ClusterSPIFFEID** comes in:

```bash
kubectl get ClusterSPIFFEID
```

And here's the output:

```txt
NAME             AGE
example          73m
vsecm-keystone   73m
vsecm-safe       73m
vsecm-sentinel   73m
```
> **ClusterSPIFFEID with an Analogy**
>
> Imagine the **ClusterSPIFFEID** as a **badge maker** for an organization.
>
> If anyone could create or modify badges (*SVIDs*), they could make one for
> themselves that mimics the CEO's badge, gaining access to restricted areas.
>
> Hence, only trusted personnel (*with elevated privileges*) are allowed to
> manage the badge maker.
>
> Make sure your guard your **ClusterSPIFFEID** with proper RBAC rules.

Let's see the details of this `example` SPIFFE ID:

```bash
kubectl describe ClusterSPIFFEID example
```

And the output:

```txt
Name:         example
Namespace:    
Labels:       <none>
Annotations:  <none>
API Version:  spire.spiffe.io/v1alpha1
Kind:         ClusterSPIFFEID
Metadata:
  Creation Timestamp:  2024-03-25T17:17:58Z
  Generation:          1
  Resource Version:    1651
  UID:                 e8af0138-7b3a-438e-9d58-21ab35a97b15
Spec:
  Pod Selector:
    Match Labels:
      app.kubernetes.io/name:  example
  Spiffe ID Template:          
  spiffe://vsecm.com/workload/example/
  ns/{{ .PodMeta.Namespace }}/
  sa/{{ .PodSpec.ServiceAccountName }}/
  n/{{ .PodMeta.Name }}
  Workload Selector Templates:
    k8s:ns:default
    k8s:sa:example
Status:
  Stats:
    Entries Masked:             0
    Entries To Set:             1
    Entry Failures:             0
    Namespaces Ignored:         4
    Namespaces Selected:        6
    Pod Entry Render Failures:  0
    Pods Selected:              1
Events:                         <none>
```

For the sake of keeping things simple because this is a quickstart, we can
assume that someone has created this `example` SPIFFE ID for us, and using
this SPIFFE ID, our example workload can securely communicate with the
**VSecM Safe**.

## Verifying Secret Registration

Since we've registered a secret, let's see if our example workload can fetch
the secret now and display it in its logs.

```bash
kubectl logs example-6cbb96b768-dhm7c
```

And the output would be something like this:

```txt
2024/03/25 17:36:13 fetch
2024/03/25 17:36:13 [TRACE] ZjmdoNn9 Sentry:Fetch https://vsecm-safe.
vsecm-system.svc.cluster.local:8443/workload/v1/secrets
2024/03/25 17:36:13 [TRACE] ZjmdoNn9 Sentry:Fetch svid:id:  spiffe://vsecm.com/
workload/example/ns/default/sa/example/n/example-6cbb96b768-dhm7c
secret: updated: "2024-03-25T17:34:25Z", created: 
"2024-03-25T17:34:25Z", value: VSecMRocks
2024/03/25 17:36:18 fetch
2024/03/25 17:36:18 [TRACE] kcOZQXeH Sentry:Fetch https://vsecm-safe.
vsecm-system.svc.cluster.local:8443/workload/v1/secrets
2024/03/25 17:36:18 [TRACE] kcOZQXeH Sentry:Fetch svid:id:  spiffe://vsecm.com/
workload/example/ns/default/sa/example/n/example-6cbb96b768-dhm7c
secret: updated: "2024-03-25T17:34:25Z", created: "2024-03-25T17:34:25Z", 
value: VSecMRocks
```

As you can see, the secret is now fetched and displayed in the logs.

The beauty of this approach is when we change the secret using
**VSecM Sentinel**, the workload will automatically fetch the
new value, without having to restart itself.

## Where to Go From Here

This quickstart is meant to give you a quick overview of how you can use
**VMware Secrets Manager** to securely manage secrets in your Kubernetes
clusters.

After successfully completing this quickstart, you can try the following:

* [Join the **VMware Secrets Manager** Community on **Slack**][slack-invite]
  where helpful community members and **VMware Secrets Manager** engineers
  hang out and answer questions.
* Navigate this website to learn more about **VMware Secrets Manager**, starting
  with [its architecture, and design philosophy](@/documentation/architecture/philosophy.md).
* [Follow a more detailed tutorial that contains multiple use cases](@/documentation/use-cases/overview.md).

[slack-invite]: https://join.slack.com/t/a-101-103-105-s/shared_invite/zt-287dbddk7-GCX495NK~FwO3bh_DAMAtQ "Join VSecM Slack"

{{ edit() }}
