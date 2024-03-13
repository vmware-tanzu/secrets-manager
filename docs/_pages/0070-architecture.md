---
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title: VSecM Architecture
layout: post
prev_url: /docs/philosophy/
permalink: /docs/architecture/
next_url: /docs/installation
---


## Introduction

This section discusses **VMware Secrets Manager** architecture and building blocks
in greater detail: We will cover **VMware Secrets Manager**'s system design and
project structure.

You don't have to know about these architectural details to use **VMware Secrets Manager**;
however, understanding how **VMware Secrets Manager** works as a system can
prove helpful when you want to extend, augment, or optimize **VMware Secrets Manager**.

Also, if you want to [contribute to the **VMware Secrets Manager** source code][contributor],
knowing what happens under the hood will serve you well.

[contributor]: /docs/contributing/ "Contribute to VMware Secrets Manager"

## Components of VMware Secrets Manager

**VMware Secrets Manager** (*VSecM*), as a system, has the following components.

### SPIRE

[**SPIRE**][spiffe] is not strictly a part of **VMware Secrets Manager**. However,
**VMware Secrets Manager** uses **SPIRE** to establish an *Identity Control Plane*.

**SPIRE** is what makes communication within **VMware Secrets Manager**
components and workloads possible. It dispatches **x.509 SVID Certificates**
to the required parties to make secure **mTLS** communication possible.

[Check out the official SPIFFE/SPIRE documentation][spiffe] for more information
about how **SPIRE** works internally.

> **It Is More SPIFFE than SPIRE**
>
> Technically, any SPIFFE-compatible Identity Control Plane can be used with
> **VMware Secrets Manager**. However, only **SPIRE** comes as part of the
> default installation.

[spiffe]: https://spiffe.io/

### VSecM Safe

[`vsecm-safe`][safe] stores secrets and dispatches them to workloads.

### VSecM Sidecar

[`vsecm-sidecar`][sidecar] is a sidecar that facilitates delivering secrets to 
workloads.

### VSecM Sentinel

[`vsecm-sentinel`][sentinel] is a pod you can shell in and do administrative 
tasks such as registering secrets for workloads.

[safe]: https://github.com/vmware-tanzu/secrets-manager/tree/main/app/safe
[sidecar]: https://github.com/vmware-tanzu/secrets-manager/tree/main/app/sidecar
[sentinel]: https://github.com/vmware-tanzu/secrets-manager/tree/main/app/sentinel

Here is a simplified overview of how various actors on a 
**VMware Secrets Manager** system interact with each other:

![VMware Secrets Manager Components](/assets/actors.jpg "VMware Secrets Manager Component Interaction")

### VSecM Keygen

[`vsecm-keygen`][keygen] is a utility that generates the root key that 
**VSecM Safe** uses, if manual input mode is set. By default, **VSecM Safe**
generates the root key automatically; however, you can opt-out from this
if want to control the root key yourself and provide your own key.

You can check out the [CLI documentation][cli] for more information
about how to use **VSecM Keygen**.

You can also use **VSecM Keygen** to decrypt secrets that **VSecM Safe** stores.
This will require you to provide the root key that **VSecM Safe** uses.
Again, can check out the [CLI documentation][cli] for more information.

[keygen]: https://github.com/vmware-tanzu/secrets-manager/tree/main/app/keygen
[cli]: https://vsecm.com/docs/cli

## High-Level Architecture

### Dispatching Identities

**SPIRE** delivers short-lived X.509 SVIDs to **VMware Secrets Manager**
components and consumer workloads.

**VSecM Sidecar** periodically talks to **VSecM Safe** to check if there is
a new secret to be updated.

![VMware Secrets Manager High Level Architecture](/assets/vsecm-hla.png "VMware Secrets Managers High Level Architecture")

Open the image above in a new tab or window to see it in full size.

### Creating Secrets

**VSecM Sentinel** is the only place where secrets can be created and registered
to **VSecM Safe**.

![Creating Secrets](/assets/vsecm-create-secrets.png "Creating Secrets")

### Component and Workload SPIFFE ID Schemas

SPIFFE ID format wor workloads is as follows:

```text
{% raw %}spiffe://vsecm.com/workload/$workloadName
/ns/{{ .PodMeta.Namespace }}
/sa/{{ .PodSpec.ServiceAccountName }}
/n/{{ .PodMeta.Name }}{% endraw %}
```

For the non-`vsecm-system` workloads that **Safe** injects secrets,
`$workloadName` is determined by the workload's `ClusterSPIFFEID` CRD.

For `vsecm-system` components, we use `vsecm-safe` and `vsecm-sentinel`
for the `$workloadName` (*along with other attestors such as attesting
the service account and namespace*):

```text
{% raw %}spiffe://vsecm.com/workload/vsecm-safe
/ns/{{ .PodMeta.Namespace }}
/sa/{{ .PodSpec.ServiceAccountName }}
/n/{{ .PodMeta.Name }}{% endraw %}
```

```text
{% raw %}spiffe://vsecm.com/workload/vsecm-sentinel
/ns/{{ .PodMeta.Namespace }}
/sa/{{ .PodSpec.ServiceAccountName }}
/n/{{ .PodMeta.Name }}{% endraw %}
```

## Persisting Secrets

**VSecM Safe** uses [`age`][age] encryption by default to securely persist the 
secrets to disk so that when its Pod is replaced by another Pod for any reason
(*eviction, crash, system restart, etc.*) it can retrieve secrets from a 
persistent storage.

> **You Can Swap the Encryption Mechanism**
>
> In **FIPS-compliant** environments, you can use **VMware Secrets Manager**
> FIPS-compliant container images that use **AES-256-GCM** encryption instead
> of **Age**.
>
> Check out the [**Configuration**](/docs/configuration) section for more
> details.
>
> Also, you can opt-out from auto-generating the private and public keys
> and provide your own keys. However, when you do this, you will have to
> manually unlock **VSecM Safe** by providing your keys every time it
> crashes. If you let **VSecM Safe** auto-generate the keys,
> you won't have to do this; so you can `#sleepmore`.
>
> Again, check out the [**Configuration**](/docs/configuration)
> section for more details.
{: .block-tip}

Since decryption is relatively expensive, once a secret is retrieved,
it is kept in memory and served from memory for better performance.
Unfortunately, this also means the amount of secrets you have for all
your workloads **has to** fit in the memory you allocate to **VSecM Safe**.

## **VSecM Safe** Bootstrapping Flow

To persist secrets, **VSecM Safe** needs a way to generate and securely store
the initial cryptographic keys that are utilized for decrypting and encrypting 
the secrets, respectively. After generation, the keys are stored in a Kubernetes 
`Secret` that only **VSecM Safe** can access.

Here is a sequence diagram of the **VSecM Safe** bootstrapping flow:

![VSecM Safe Bootstrapping](/assets/vsecm-bootstrap.png "VSecM Safe Bootstrapping Flow")

Note that, until bootstrapping is complete, **VSecM Safe** will not respond to
any API requests that you make from **VSecM Sentinel**.

[age]: https://github.com/FiloSottile/age

## **VSecM Safe** Pod Layout

Here is what an **VSecM Safe** Pod looks like at a high level:

![VSedM Safe Pod](/assets/vsecm-crypto.jpg "VSecM Safe Pod")

* `spire-agent-socket`: Is a [SPIFFE CSI Driver][csi-driver]-managed volume that
  enables **SPIRE** to distribute **X.509 SVID**s to the Pod.
* `/data` is the volume where secrets are stored in an encrypted format. You are
  **strongly encouraged** to use a **persistent volume** for production setups
  to retrieve the secrets if the Pod crashes and restarts.
* `/key` is where the secret `vsecm-root-key` mounts. For security reasons,
  ensure that **only** the pod **VSecM Safe** can read and write to `vsecm-root-key`
  and no one else has access. In this diagram, this is achieved by assigning
  a `vsecm-secret-readwriter` role to **VSecM Safe** and using that role to update
  the secret. Any pod that does not have the role will be unable to read or
  write to this secret.

If the `main` container does not have a public/private key pair in memory, it
will attempt to retrieve it from the `/key` volume. If that fails, it will
generate a brand new key pair and then store it in the `vsecm-root-key` secret.

[csi-driver]: https://github.com/spiffe/spiffe-csi

## Template Transformation and K8S Secret Generation

Here is a sequence diagram of how the and **VSecM Safe**-managed *secret*
is transformed into a **Kubernetes** `Secret` (*open the image in a
new tab for a larger version*):

![Transforming Secrets](/assets/vsecm-secret-transformation.png "Transforming Secrets")

There are two parts to this:

* Transforming secrets using a Go template transformation
* Updating the relevant **Kubernetes** `Secret`

You can check [**VSecM Sentinel** CLI Documentation](/docs/cli) for
various ways this transformation can be done. In addition, you can check
[**VSecM** Secret Registration Tutorial][register-a-secret] for more information
about how the **Kubernetes** `Secret` object is generated and used in workloads.

[register-a-secret]: /docs/quickstart/#register-a-secret "Register a Secret"

## Liveness and Readiness Probes

**VSecM Safe** and **VSecM Sentinel** use **liveness** and **readiness** probes.
These probes are tiny web servers that serve at ports `8081` and `8082` by
default, respectively.

You can set `VSECM_PROBE_LIVENESS_PORT` (*default `:8081`*) and
`VSECM_PROBE_READINESS_PORT` (*default `:8082`*) environment variables to change
the ports used for these probes.

When the service is healthy, the liveness probe will return an `HTTP 200` success
response. When the service is ready to receive traffic, the readiness
probe will return an `HTTP 200` success response.

## Conclusion

This was a deeper overview of **VMware Secrets Manager** architecture. If you
have further questions, feel free to [join the **VMware Secrets Manager**
community on **Slack**][slack-invite] and ask them out.

[slack-invite]: https://join.slack.com/t/a-101-103-105-s/shared_invite/zt-287dbddk7-GCX495NK~FwO3bh_DAMAtQ "Join VSecM Slack"

<p class="github-button">
  <a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0070-architecture.md">
    Suggest edits ✏️ 
  </a>
</p>