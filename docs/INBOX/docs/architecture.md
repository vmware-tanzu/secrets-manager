---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, architecture, system design, deep dive
title: Aegis Deep Dive
description: <strong>Aegis</strong> under the hood
micro_nav: true
page_nav:
  prev:
    content: <strong>Aegis</strong> Showcase
    url: '/docs/use-case-showcase'
  next:
    content: configuring <strong>Aegis</strong>
    url: '/docs/configuration'
---

<p style="text-align:right;position:relative;top:-40px;"
><a href="https://github.com/ShieldWorks/aegis-web/blob/main/docs/architecture.md"
style="border-bottom: none;background:#e0e0e0;padding:0.5em;display:inline-block;
border-radius:8px;">
edit this page on <strong>GitHub</strong> ✏️</a></p>

## Introduction

This section discusses **Aegis** architecture and building blocks in greater
detail: We will cover **Aegis**’ system design and project structure.

You don’t have to know about these architectural details to use **Aegis**;
however, understanding how **Aegis** works as a system can
prove helpful when you want to extend, augment, or optimize **Aegis**.

Also, if you want to [contribute to the **Aegis** source code][contributor],
knowing what happens under the hood will serve you well.

[contributor]: https://aegis.ist/contact/#i-want-to-be-a-contributor

## Components of Aegis

**Aegis**, as a system, has the following components.

### Aegis SPIRE

[`aegis-spire`][aegis-spire] is what makes communication within **Aegis**
components and workloads possible. It dispatches **x.509 SVID Certificates**
to the required parties to make secure **mTLS** communication possible.

[Check out the official SPIFFE/SPIRE documentation][spiffe] for more information
about how **SPIRE** works internally.

[spiffe]: https://spiffe.io/

### Aegis Safe

[`aegis-safe`][safe] stores secrets and dispatches them to workloads.

### Aegis Sidecar

[`aegis-sidecar`][sidecar] is a sidecar that facilitates delivering secrets to workloads.

### Aegis Sentinel

[`aegis-sentinel`][sentinel] is a pod you can shell in and do administrative tasks such as
registering secrets for workloads.

[aegis-spire]: https://github.com/ShieldWorks/aegis/tree/main/k8s/spire
[safe]: https://github.com/shieldworks/tree/main/app/safe
[sidecar]: https://github.com/shieldworks/tree/main/app/sidecar
[sentinel]: https://github.com/shieldworks/tree/main/app/sentinel

Here is a simplified overview of how various actors on an **Aegis** system
interact with each other:

![Aegis Components](/assets/actors.jpg "Aegis Component Interaction")

## Technologies Used

Without the following technologies, implementing **Aegis** would have been a very
hard, time-consuming, and error-prone endeavor.

* [SPIFFE and **SPIRE**][spire] for establishing an Identity Control Plane.
* [**Age** Encryption][age] to enable out-of-memory encrypted
  backup of the secrets stored for disaster recovery.

[spire]: https://spiffe.io/ "SPIFFE: Secure Production Identity Framework for Everyone"
[sops]: https://github.com/mozilla/sops "Sops: Simple and flexible tool for managing secrets"
[age]: https://github.com/FiloSottile/age "Age: A secure and modern encryption tool"

## High-Level Architecture

### Dispatching Identities

**SPIRE** delivers short-lived X.509 SVIDs to **Aegis**
components and consumer workloads.

**Aegis Sidecar** periodically talks to **Aegis Safe** to check if there is
a new secret to be updated.

![Aegis High Level Architecture](/assets/aegis-hla.png "Aegis High Level Architecture")

### Creating Secrets

**Aegis Sentinel** is the only place where secrets can be created and registered
to **Aegis Safe**.

![Creating Secrets](/assets/aegis-create-secrets.png "Creating Secrets")

### Component and Workload SVID Schemas

SPIFFE ID format wor workloads is as follows:

```text
{% raw %}spiffe://aegis.ist/workload/$workloadName
/ns/{{ .PodMeta.Namespace }}
/sa/{{ .PodSpec.ServiceAccountName }}
/n/{{ .PodMeta.Name }}{% endraw %}
```

For the non-`aegis-system` workloads that **Safe** injects secrets,
`$workloadName` is determined by the workload’s `ClusterSPIFFEID` CRD.

For `aegis-system` components, we use `aegis-safe` and `aegis-sentinel`
for the `$workloadName` (*along with other attestors such as attesting
the service account and namespace*):

```text
{% raw %}spiffe://aegis.ist/workload/aegis-safe
/ns/{{ .PodMeta.Namespace }}
/sa/{{ .PodSpec.ServiceAccountName }}
/n/{{ .PodMeta.Name }}{% endraw %}
```

```text
{% raw %}spiffe://aegis.ist/workload/aegis-sentinel
/ns/{{ .PodMeta.Namespace }}
/sa/{{ .PodSpec.ServiceAccountName }}
/n/{{ .PodMeta.Name }}{% endraw %}
```

## Persisting Secrets

**Aegis Safe** uses [age][age] to securely persist the secrets to disk so that
when its Pod is replaced by another pod for any reason
(*eviction, crash, system restart, etc.*). When that happens, **Aegis Safe**
can retrieve secrets from a persistent storage.

Since decryption is relatively expensive, once a secret is retrieved,
it is kept in memory and served from memory for better performance.
Unfortunately, this also means the amount of secrets you have for all
your workloads **has to** fit in the memory you allocate to **Aegis Safe**.

## **Aegis Safe** Bootstrapping Flow

To persist secrets, **Aegis Safe** needs a way to generate and securely store
the private and public `age` keys that are utilized for decrypting and
encrypting the secrets, respectively.

* Key generation is conveniently provided by `age` Go SDK.
* After generation, the keys are stored in a Kubernetes `Secret` that only
  **Aegis Safe** can access.

Here is a sequence diagram of the **Aegis Safe** bootstrapping flow:

![Aegis Safe Bootstrapping](/assets/bootstrap.jpg "Aegis Safe Bootstrapping Flow")

Note that, until bootstrapping is complete, **Aegis Safe** will not respond to
any API requests that you make from **Aegis Sentinel**.

[age]: https://github.com/FiloSottile/age

## **Aegis Safe** Pod Layout

Here is what an **Aegis Safe** Pod looks like at a high level:

![Aegis Safe Pod](/assets/crypto.jpg "Aegis Safe Pod")

* `spire-agent-socket`: Is a [SPIFFE CSI Driver][csi-driver]-managed volume that
  enables **SPIRE** to distribute **X.509 SVID**s to the Pod.
* `/data` is the volume where secrets are stored in an encrypted format. You are
  **strongly encouraged** to use a **persistent volume** for production setups
  to retrieve the secrets if the Pod crashes and restarts.
* `/key` is where the secret `aegis-safe-age-key` mounts. For security reasons, 
  ensure that **only** the pod **Aegis Safe** can read and write to `aegis-safe-age-key`
  and no one else has access. In this diagram, this is achieved by assigning
  a `secret-readwriter` role to **Aegis Safe** and using that role to update
  the secret. Any pod that does not have the role will be unable to read or
  write to this secret.

If the `main` container does not have a public/private key pair in memory, it
will attempt to retrieve it from the `/key` volume. If that fails, it will
generate a brand new key pair and then store it in the `aegis-safe-age-key` secret.

[csi-driver]: https://github.com/spiffe/spiffe-csi

## Template Transformation and Kubernetes Secret Generation

Here is a sequence diagram of how the and **Aegis Safe**-managed *secret* 
is transformed into a **Kubernetes** `Secret` (*open the image in a 
new tab for a larger version*):

![Transforming Secrets](/assets/secret-transformation.png "Transforming Secrets")

There are two parts to this:

* Transforming secrets using a Go template transformation
* Updating the relevant **Kubernetes** `Secret`

You can check [**Aegis Sentinel** CLI Documentation](/docs/sentinel) for
various ways this transformation can be done. In addition, you can check 
[**Aegis** Secret Registration Tutorial](/docs/register) for more information
about how the **Kubernetes** `Secret` object is generated and used in workloads.

## Liveness and Readiness Probes

**Aegis Safe** and **Aegis Sentinel** use **liveness** and **readiness** probes.
These probes are tiny web servers that serve at ports `8081` and `8082` by
default, respectively.

You can set `AEGIS_PROBE_LIVENESS_PORT` (*default `:8081`*) and 
`AEGIS_PROBE_READINESS_PORT` (*default `:8082`*) environment variables to change
the ports used for these probes.

When the service is healthy, the liveness probe will return an `HTTP 200` success
response. When the service is ready to receive traffic, the readiness
probe will return an `HTTP 200` success response.

## Conclusion

This was a deeper overview of **Aegis** architecture. If you have further
questions, feel free to [join the **Aegis** community on Slack][slack-invite]
and ask them out.

[slack-invite]: https://join.slack.com/t/aegis-6n41813/shared_invite/zt-1myzqdi6t-jTvuRd1zDLbHX0gN8VkCqg "Join aegis.slack.com"
