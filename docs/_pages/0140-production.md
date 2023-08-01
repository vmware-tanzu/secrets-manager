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

title: Production Setup
layout: post
next_url: /docs/use-cases/
prev_url: /docs/contributing/
permalink: /docs/production/
---

<p class="github-button"
><a
href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0140-production.md"
>edit this page on <strong>GitHub</strong> ✏️</a></p>

## VMware Secrets Manager for Cloud Native Apps

## Introduction

You need to pay attention to certain aspects and parts of the system that
you’d need to harden for a **production** *VMware Secrets Manager* setup.
This article will overview them.

## Version Compatibility

We test **VMware Secrets Manager** with the recent stable version of Kubernetes and Minikube.

As long as there isn’t a change in the **major** version number your
Kubernetes client and server you use, things will likely work just fine.

## Resource Requirements

**VMware Secrets Manager** is designed from the ground up to work in environments with limited
resources, such as edge computing and IoT.

That being said, **VMware Secrets Manager**, by design, is a memory-intensive application.
However, even when you throw all your secrets at it, **VMware Secrets Manager Safe**’s peak
memory consumption will be in the order or 10-20 megabytes of RAM. The CPU
consumption will be within reasonable limits too.

However, it’s crucial to understand that every system and user profile is unique.
Factors such as the number and size of secrets, concurrent processes, and system
specifications can influence these averages. Therefore, it is always advisable to
benchmark **VMware Secrets Manager** and **SPIRE** on your own system under your specific usage
conditions to accurately gauge the resource requirements to ensure optimal
performance.

Benchmark your system usage and set **CPU** and **Memory** limits to the
**VMware Secrets Manager Safe** pod.

We recommend you to:

* Set a memory **request** and **limit** for **VMware Secrets Manager Safe**,
* Set a CPU **request**; but **not** set a CPU limit for **VMware Secrets Manager Safe**
  (*i.e., the **VMware Secrets Manager Safe** pod will ask for a baseline CPU;
  yet burst for more upon need*).

As in any secrets management solution, your compute and memory requirements
will depend on several factors, such as:

* The number of workloads in the cluster
* The number of secrets **Safe** (*VMware Secrets Manager’ Secrets Store*) has to manage
  (*see [architecture details][architecture] for more context*)
* The number of workloads interacting with **Safe**
  (*see [architecture details][architecture] for more context*)
* **Sidecar** poll frequency (*see [architecture details][architecture] for
  more context*)
* etc.

[architecture]: /docs/architecture

We recommend you benchmark with a realistic production-like
cluster and allocate your resources accordingly.

That being said, here are the resource allocation reported by `kubectl top`
for a demo setup on a single-node minikube cluster to give an idea:

```text 
NAMESPACE     WORKLOAD            CPU(cores) MEMORY(bytes)
vsecm-system  vsecm-safe          1m         9Mi
vsecm-system  vsecm-sentinel      1m         3Mi
default       example 2m         7Mi
spire-system  spire-agent         4m         35Mi
spire-system  spire-server        6m         41Mi
```

Note that 1000m is 1 full CPU core.

Based on these findings, the following resource and limit allocations can be
a starting point for **VMware Secrets Manager**-managed containers:

```text
  # Resource allocation will highly depend on the system.
  # Benchmark your deployment, monitor your resource utilization,
  # and adjust these values accordingly.
  resources:
    requests:
      memory: "128Mi"
      cpu: "250m"
    limits:
      memory: "128Mi"
      # We recommend “NOT” setting a CPU limit.
      # As long as you have configured your CPU “requests”
      # correctly, everything would work fine.
```

## Back Up Your Cluster Regularly

**VMware Secrets Manager** is designed to be resilient; however, losing access to your sensitive
data is possible by inadvertently deleting a Kubernetes `Secret` that you are
not supposed to delete. Or, your backing store that contains the secrets can get
corrupted for any reason.

Cloud Native or not, you rely on hardware which—intrinsically—is unreliable.

Things happen. Make sure you back up your cluster [using a tool like
**Velero**][velero], so that when things do happen, you can revert your
cluster’s last known good state.

**Set up your backups from day zero**.

[velero]: https://velero.io/ "Velero"

## Restrict Access To `vsecm-safe-age-key`

The `vsecm-safe-age-key` secret that **VSecM Safe** stores in the `vsecm-system`
namespace contains the keys to encrypt and decrypt secret data on the data
volume of **VSecM Safe**.

While reading the secret alone is not enough to plant an attack on the secrets
(*because the attacker also needs to access the VMware Secrets Manager Safe Pod or the `/data`
volume in that Pod*), it is still **crucial** to follow the **principle of least
privilege** guideline and do not allow anyone on the cluster read or write
to the `vsecm-safe-age-key` secret.

The only entity allowed to have read/write (*but not delete*) access to
`vsecm-safe-age-key` should be the **VSecM Safe** Pod inside the `vsecm-system`
namespace with an `vsecm-safe` service account.

> **With Great Power Comes Great Responsibility**
>
> It is worth noting that a **Cluster Administrator** due to their elevated
> privileges can read/write to any Kubernetes `Secret` in the cluster.
>
> This includes access to the `vsecm-safe-age-key` secret. Therefore, it is
> highly recommended that you grant the `cluster-admin` role to a **very**
> small group of trusted individuals only.
>
> Although, access to `vsecm-safe-age-key` does not give the attacker direct
> access to the secrets, due to their sheer power, a determined Cluster
> Administrator can still read the secrets by accessing the `/data` volume.
>
> Their actions will be recorded in the audit logs, so they can, and will be
> held responsible; however, it is still a bad idea to have more than an
> absolute minimum number of Cluster Administrators in your cluster.

Kubernetes Secrets are, by default, stored **unencrypted** in the API server’s
underlying data store (`etcd`). Anyone with API access and sufficient RBAC
credentials can retrieve or modify a `Secret`, as can anyone with access
to `etcd`.

If you are **only** using **VMware Secrets Manager** for your configuration and secret storage
needs, and your workloads do **not** bind any Kubernetes `Secret` (*i.e.,
instead of using Kubernetes `Secret` objects, you use tools like **VMware Secrets Manager SDK**
or **VMware Secrets Manager Sidecar** to securely dispatch secrets to your workloads*) then
as long as you secure access to the secret `vsecm-safe-age-key` inside the
`vsecm-system` namespace, you should be good to go.

With the help of **VMware Secrets Manager SDK**, **VMware Secrets Manager Sidecar**, and **VMware Secrets Manager Init Container**,
and with some custom coding/shaping of your data, you should be able to use it.
[Even when you don’t have direct access to the application’s source code, there
are still ways to inject environment variables without Kubernetes `Secret`s
or `ConfigMap`s][david-copperfield].

<!-- TODO: Update below link with new link -->
<!-- [david-copperfield]: https://github.com/shieldworks/aegis/issues/231 -->

<!-- TODO: Fix below link for [VSecM-k] -->
However, **VMware Secrets Manager** also has the option to[persist the secrets stored in
**VMware Secrets Manager Safe** as Kubernetes `Secret` objects][VSecM-k]. This approach can
help support **legacy** systems where you want to start using
**VMware Secrets Manager** without introducing much code and infrastructure change to the
existing cluster—at least initially.

[VSecM-k]: https://vsecm.com/docs/sentinel/#creating-kubernetes-secrets "VMware Secrets Manager Sentinel: Creating Kubernetes Secrets"

If you are using **VMware Secrets Manager** to generate Kubernetes `Secrets` for the workloads
to consume, then take regular precautions around those secrets,
such as [*implementing restrictive RBACs*][rbac], and even [considering using
a KMS to encrypt `etcd` at rest][kms] if your security posture requires it.

### **Do I Really Need to Encrypt `etcd`?**

#### tl;dr:

Using plain Kubernetes `Secret`s is good enough, and it is not the
end of the world if you keep your `etcd` unencrypted.

> **VMware Secrets Manager Keeps Your Secrets Safe**
>
> If you use **VMware Secrets Manager** to store your sensitive data, your secrets
> will be securely stored in **VMware Secrets Manager Safe** (*instead of `etcd`*),
> so you will have even fewer reasons to encrypt `etcd` 😉.

#### Details

This is an excellent question. And as in any profound answer to
good questions, the answer is: “*it depends*” 🙂.

`Secret`s are, by default, stored unencrypted in `etcd`. So if an adversary
can read `etcd` in any way, it’s game over.

##### Threat Analysis

Here are some ways this could happen:

1. Root access to a control plane node.
2. Root access to a worker node.
3. Access to the actual physical server (*i.e., physically removing the disk*).
4. Possible zero day attacks.

For `1`, and `2`, server hardening, running secure Linux instances, patching,
and preventing privileged pods from running in the cluster are the usual ways
to mitigate the threat. Unfortunately, it is a relatively complex attack vector
to guard against. Yet, once your node is compromised, you have **a lot**
of things to worry about. In that case, `etcd` exposure will be just one of
many, *many*, **many** concerns that you’ll have to worry about.

For `3`, assuming your servers are in a data center, there should already be
physical security to secure your servers. So the attack is **unlikely**
to happen. In addition, your disks are likely encrypted, so unless the attacker
can shell into the operating system, your data is already safe: Encrypting
`etcd` once more will not provide any additional advantage in this particular
case, given the disk is encrypted, and root login is improbable.

For `4.`, the simpler your setup is, the lesser moving parts you have, and the
lesser the likelihood of bumping into a zero-day. And Kubernetes `Secret`s
are as simple as it gets.

Even when you encrypt `etcd` at rest using a **KMS** (*which is the most robust
method proposed [in the Kubernetes guides][kms]*), an attacker can still
impersonate `etcd` and decrypt the secrets: As long as you provide the correct
encrypted DEK to KMS, the KMS will be more than happy to decrypt that DEK with
its KEK and provide a plain text secret to the attacker.

##### Secure Your House Before Securing Your Jewelry

So, **yes**, securing `etcd` will marginally increase your security posture.
Yet, it does not make too much of a difference unless you have **already**
secured your virtual infrastructure **and** physical data center. And
if you haven’t secured your virtual and physical assets, then you are in big
trouble at day zero, even before you set up your cluster, so encrypting
`etcd` will not save you the slightest from losing other valuable data
elsewhere anyway.

##### Security Is a Layered Cake

That being said, we are humans, and $#!% does happen: If a node is compromised
due to a misconfiguration, it would be nice to make the job harder for the
attacker.

[rbac]: https://kubernetes.io/docs/reference/access-authn-authz/rbac/
[kms]: https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/

## Restrict Access to VMware Secrets Manager Sentinel

All **VMware Secrets Manager** images are based on [distroless][distroless] containers for an
additional layer of security. Thus, an operator cannot execute a shell on the
Pod to try a privilege escalation or container escape attack. However, this does
not mean you can leave the `vsecm-system` namespace like an open buffet.

Always take a **principle of least privilege** stance. For example, do not let
anyone who does not need to fiddle with the `vsecm-system` namespace see and use
the resources there.

This stance is especially important for the **VMware Secrets Manager Sentinel** Pod since an
attacker with access to that pod can override (*but not read*) secrets on
workloads.

**VMware Secrets Manager** leverages Kubernetes security primitives and modern cryptography
to secure access to secrets. And **VMware Secrets Manager Sentinel** is the **only** system
part that has direct write access to the **VMware Secrets Manager Safe** secrets store. Therefore,
once you secure access to **VMware Secrets Manager Sentinel** with [proper RBAC and
policies][rbac], you secure access to your secrets.

[distroless]: https://github.com/GoogleContainerTools/distroless

## Volume Selection for VMware Secrets Manager Safe Backing Store

[**VMware Secrets Manager Safe** default deployment descriptor][vsecm-safe-deployment-yaml]
uses [`HostPath`][k8s-pv] to store encrypted backups for secrets.

It is highly recommended to ensure that the backing store **VMware Secrets Manager Safe**
uses is **durable**, **performant**, and **reliable**.

It is a best practice to avoid `HostPath` volumes for production deployments.
You are strongly encouraged to [choose a `PersistentVolume` that suits your
needs][k8s-pv] for production setups.

[vsecm-safe-deployment-yaml]: https://github.com/vmware-tanzu/secrets-manager/blob/main/helm-charts/charts/safe/templates/Deployment.yaml
[k8s-pv]: https://kubernetes.io/docs/concepts/storage/volumes/

## High Availability of VMware Secrets Manager Safe

> **tl;dr:**
>
> **VMware Secrets Manager Safe** may not emphasize high-availability, but its robustness is
> so outstanding that the need for high-availability becomes almost negligible.

Since **VMware Secrets Manager Safe** keeps all of it state in memory, using a pod with enough
memory and compute resources is the most effective way to leverage it. Although,
with some effort, it might be possible to make it highly available, the effort
will likely bring unnecessary complexity without much added benefit.

**VMware Secrets Manager Safe** is, by design, a single pod; so technically-speaking, it is
not highly-available. So in the rare case when **VMware Secrets Manager Safe** crashes, or
gets evicted due to a resource contention, there will be minimal disruption
until it restarts. However, **VMware Secrets Manager Safe** restarts fairly quickly, so the
time window where it is unreachable will hardly be an issue.

Moreover **VMware Secrets Manager Safe** employs “*lazy learning*” and does not load everything
into memory all at once, allowing **very** fast restarts. In addition, its
lightweight and focused code ensures that crashes are infrequent, making
**VMware Secrets Manager Safe** *practically* highly available.

While it is possible to modify the current architecture to include more than one
**VMware Secrets Manager Safe** pod and place it behind a service to ensure high-availability,
this would be a significant undertaking, with not much benefit to merit it:

First of all, for that case to happen, the state would need to be moved away
from the memory, and centralized into a common in-memory store (*such as Redis,
or etcd*). This will introduce another moving part to manage. Or alternatively
all **VMware Secrets Manager Safe** pods could be set up to broadcast their operations and reach
a quorum. A quorum-based solution would be more complex than using a share store,
besides reaching a quorum means a performance it (*both in terms of decision time
and also compute required*).

On top of all these bootstrapping coordination would be necessary to prevent
two pods from creating different bootstrap secrets simultaneously.

Also, for a backing store like Redis, the data would need to be encrypted
(*and Redis, for example, does not support encryption at rest by default*).

When considering all these, **VMware Secrets Manager Safe** has not been created highly-available
**by design**; however, it is so robust, and it restarts from crashes so fast that
it’s “*as good as*” highly-available.



## Update VMware Secrets Manager’ Log Levels

**VMware Secrets Manager Safe** and **VMware Secrets Manager Sidecar** are configured to log at `TRACE` level by
default. This is to help you debug issues with **VMware Secrets Manager**. However, this can
cause a lot of noise in your logs. Once you are confident that **VMware Secrets Manager**
works as expected, you can reduce the log level to `INFO` or `WARN`.

For this, you will need to modify the `VSECM_LOG_LEVEL` environment variable
in the **VMware Secrets Manager Safe** and **VMware Secrets Manager Sidecar** Deployment manifests.

See [**Configuring VMware Secrets Manager**](/docs/configuration) for details.

## Conclusion

Since **VMware Secrets Manager** is a *Kubernetes-native* framework, its security is strongly
related to how you secure your cluster. You should be safe if you keep your
cluster and the `vsecm-system` namespace secure and follow
“*the principle of least privilege*” as a guideline.

**VMware Secrets Manager** is a lightweight secrets manager; however, that does not mean it
runs on water: It needs CPU and Memory resources. The amount of resources you
need will depend on the criteria outlined in the former sections. You can either
benchmark your system and set your resources accordingly. Or set generous-enough
limits and adjust your settings as time goes by.

Also, you are strongly encouraged **not** to set a limit on **VMware Secrets Manager** Pods’ CPU
usage. Instead, it is recommended to let **VMware Secrets Manager Safe** burst the CPU when
it needs.

On the same topic, you are encouraged to set a **request** for **VMware Secrets Manager Safe**
to guarantee a baseline compute allocation.
