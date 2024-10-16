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

title = "VSecM Architecture"
weight = 10
+++

## Introduction

This section discusses **VMware Secrets Manager** architecture and building blocks
in greater detail: We will cover **VMware Secrets Manager**'s system design and
project structure.

You don't have to know about these architectural details to use **VMware Secrets Manager**;
however, understanding how **VMware Secrets Manager** works as a system can
prove helpful when you want to extend, augment, or optimize **VMware Secrets Manager**.

Also, if you want to [contribute to the **VMware Secrets Manager** source code][contributor],
knowing what happens under the hood will serve you well.

[contributor]: @/documentation/development/contributing.md "Contribute to VMware Secrets Manager"

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
generates the root key automatically; however, you can opt out from this
if you want to control the root key yourself and provide your own key.

You can check out the [CLI documentation][cli] for more information
about how to use **VSecM Keygen**.

You can also use **VSecM Keygen** to decrypt secrets that **VSecM Safe** stores.
This will require you to provide the root key that **VSecM Safe** uses.
Again, can check out the [CLI documentation][cli] for more information.

[keygen]: https://github.com/vmware-tanzu/secrets-manager/tree/main/app/keygen
[cli]: @/documentation/usage/cli.md

### VSecM Keystone

`vsecm-keystone` is a component that leverages **VSecM Init Container**
behind the scenes.

When the entire **VSecM** system finishes bootstrapping, and it is ready
for operation, **VSecM Keystone** will be in a `Ready` state.

This is useful when you want to configure your orchestrator to wait until
**VSecM** is ready before starting your workloads.

**VSecM Keystone** is an optional component and it may be disabled if you
do not need it.

### VSecM Init Container

`vsecm-init-container` is a sidecar that you can add to a workload to let
it wait until a secrets is registered for it in **VSecM Safe**.

This is useful for use cases where you want to make sure that a workload
does not initialize without the required secrets.

[You can check out the use cases section for more information][use-cases].

> **VSecM Init Container and VSecM Sidecar Coordination**
> 
> **VSecM Init Container** and **VSecM Sidecar** can be used in coordination
> with each other, each serving its own purpose.
> 
> In that case, **VSecM Sidecar** will pre-populate the secret data
> inside an in-memory volume, while **VSecM Init Container** will wait
> until the secret is available before starting the application.

[use-cases]: @/documentation/use-cases/overview.md

### VSecM Inspector

**VSecM Inspector** is a sample application that has an `./env` executable
at the root of it where you can shell into it and execute `./env` to
see the secrets that **VSecM Safe** injects to it.

This is useful for debugging and testing purposes.

## High-Level Architecture

### Dispatching Identities

**SPIRE** delivers short-lived [X.509 SVIDs][svid] to **VMware Secrets Manager**
components and consumer workloads.

[svid]: https://spiffe.io/docs/latest/deploying/svids/

**VSecM Sidecar** periodically talks to **VSecM Safe** to check if there is
a new secret to be updated.

Open the image above in a new tab or window to see it in full size:

![VMware Secrets Manager Sidecar](/assets/vsecm-sidecar.png "VMware Secrets Manager Sidecar")

Alternatively, you can use **VSecM SDK** to retrieve secrets from **VSecM Safe**
programmatically:

![Using VSecM SDK](/assets/vsecm-sdk.png "Using VSecM SDK")

### Creating Secrets

**VSecM Sentinel** is the only place where secrets can be created and registered
to **VSecM Safe**.

![Creating Secrets](/assets/vsecm-create-secrets.png "Creating Secrets")

### Component and Workload SPIFFE ID Schemas

SPIFFE ID format for workloads is as follows:

```txt
spiffe://vsecm.com/workload/$workloadName
  /ns/{{ .PodMeta.Namespace }}
  /sa/{{ .PodSpec.ServiceAccountName }}
  /n/{{ .PodMeta.Name }}
```

For the non-`vsecm-system` workloads that **Safe** injects secrets,
`$workloadName` is determined by the workload's `ClusterSPIFFEID` CRD.

For `vsecm-system` components, we use `vsecm-safe` and `vsecm-sentinel`
for the `$workloadName` (*along with other attestors such as attesting
the service account and namespace*):

```txt
spiffe://vsecm.com/workload/vsecm-safe
  /ns/{{ .PodMeta.Namespace }}
  /sa/{{ .PodSpec.ServiceAccountName }}
  /n/{{ .PodMeta.Name }}
```

```txt
spiffe://vsecm.com/workload/vsecm-sentinel
  /ns/{{ .PodMeta.Namespace }}
  /sa/{{ .PodSpec.ServiceAccountName }}
  /n/{{ .PodMeta.Name }}
```

> **You Can Have Custom SPIFFE ID Formats**
> 
> You can configure your **VSecM** installation to use custom SPIFFE ID
> formats and RegExp-based validation rules. Check out the 
> [**Configuration**](@/documentation/configuration/overview.md) section
> for more details.

## A Note on SPIRE Controller Manager and ClusterSPIFFEIDs

**VSecM** uses **SPIRE** to establish an Identity Control Plane.
And **SPIRE** uses [**SPIRE Controller Manager**][scm] to manage the 
**SPIRE Server** and **SPIRE Agent**. [**SPIRE Controller Manager**][scm] 
uses [**ClusterSPIFFEIDs**][clusterspiffeid] to assign and determine SVIDs 
for workloads.

[scm]: https://github.com/spiffe/spire-controller-manager
[clusterspiffeid]: https://github.com/spiffe/spire-controller-manager/blob/main/docs/clusterspiffeid-crd.md

What this means for **VSecM** is that to assign identities to your workloads,
you need to create **ClusterSPIFFEIDs** for them.

Here is an example:

```yaml
apiVersion: spire.spiffe.io/v1alpha1
kind: ClusterSPIFFEID
metadata:
  name: example-workload
spec:
  className: "vsecm"
  spiffeIDTemplate: "spiffe://vsecm.com\
    /workload/example\
    /ns/default\
    /sa/example\
    /n/{{ .PodMeta.Name }}"
  podSelector:
    matchLabels:
      app.kubernetes.io/name: example-workload
  workloadSelectorTemplates:
  - "k8s:ns:default"
  - "k8s:sa:example-workload-sa"
```

This **ClusterSPIFFEID** will assign the following SPIFFE ID to the workloads
that are in the `default` namespace, have the 
`app.kubernetes.io/name: example-workload`label, and use 
the `example-workload-sa` service account.

Unless you are using a custom configuration for **VSecM** ClusterSPIFFEIDs,
the `className` field should be set to `vsecm` and the `spiifeIDTemplate`
should start wih `spiffe://vsecm.com/workload/` followed by the name you give
to the workload.

You can check out the [**VSecM** quick start guide][quickstart] for more 
information on how to create **ClusterSPIFFEIDs** for your workloads.

[quickstart]: @/documentation/getting-started/overview.md

One additional thing to note about **SPIRE Controller Manager** is that it
takes ownership of all **SPIRE** entries by default. This means that if you
manually create an entry using **SPIRE Server**’s CLI, 
**SPIRE Controller Manager** will override it and reflect its own state that
are defined in the **ClusterSPIFFEIDs**. 

This essentially means that the **ClusterSPIFFEIDs** are the source of truth 
for identities in **SPIRE** and therefore in **VSecM**.

> **Can I Have My Own SPIRE Server Entries**?
> 
> If you want to register your own **SPIRE Server** entries, you will need 
> a custom SPIRE deployment that does not use SPIRE Controller Manager.
> 
> **VMware Secrets Manager** can work with any SPIFFE-compatible Identity
> Control Plane, so you can use your own SPIRE deployment if you want to.
> 
> However, using `ClusterSPIFFEID`s is the recommended way to assign identities 
> to your workloads in **VSecM**. They are easier to manage and maintain because
> they are Kubernetes CRDs and use the declarative nature of Kubernetes.

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
> Check out the [**Configuration**](@/documentation/configuration/overview.md) 
> section for more details.
>
> Also, you can opt out from auto-generating the private and public keys
> and provide your own keys. However, when you do this, you will have to
> manually unlock **VSecM Safe** by providing your keys every time it
> crashes. If you let **VSecM Safe** auto-generate the keys,
> you won't have to do this; so you can `#sleepmore`.
>
> Again, check out the [**Configuration**](@/documentation/configuration/overview.md)
> section for more details.

Since decryption is relatively expensive, once a secret is retrieved,
it is kept in memory and served from memory for better performance.
Unfortunately, this also means the amount of secrets you have for all
your workloads **has to** fit in the memory you allocate to **VSecM Safe**.

## Bootstrapping Flow of **VSecM Safe**

To persist secrets, **VSecM Safe** needs a way to generate and securely store
the initial cryptographic keys that are utilized for decrypting and encrypting
the secrets, respectively. After generation, the keys are stored in a Kubernetes
`Secret` that only **VSecM Safe** can access.

Here is a sequence diagram of the **VSecM Safe** bootstrapping flow:

![VSecM Safe Bootstrapping](/assets/vsecm-bootstrap.png "VSecM Safe Bootstrapping Flow")

Note that, until bootstrapping is complete, **VSecM Safe** will not respond to
any API requests that you make from **VSecM Sentinel**.

Here is a simplified version of the bootstrapping flow without taking the
"manual key input" option into account:

![VSecM Safe Bootstrapping](/assets/vsecm-bootstrap-simplified.png "VSecM Safe Bootstrapping Flow (Simplified)")

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

## Template Transformation and Kubernetes Secret Generation

**VSecM** enables you to transform the secrets you register with **VSecM Safe**
using Go template transformations before storing them. This is useful when
you want to store a secret in a format that is different from the format
you want to use in your workloads.

You can also prefix `k8s:` to the name of the workload to create a Kubernetes
`Secret` object with the transformed secret. This is useful when you want to
inject the secret to a workload using a Kubernetes `Secret` object instead
of **VSecM Sidecar** or **VSecM SDK**.

Check out the [examples folder][github-examples] for examples of how to use
**VSecM** to transform secrets and generate **Kubernetes** `Secret` objects.

[github-examples]: https://github.com/vmware-tanzu/secrets-manager/tree/main/examples "VSecM Examples"

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

## **VSecM** and SPIRE Deployment Diagram

Here is a different look at how **VSecM** and **SPIRE** components are deployed
in a Kubernetes cluster, focusing on the Services exposed by the components and
volumes mounted to the Pods.

Open the image in a new tab to see it in full size:

![VSecM and SPIRE Deployment](/assets/vsecm-infra.png "VSecM and SPIRE Deployment Diagram")

## Conclusion

This was a deeper overview of **VMware Secrets Manager** architecture. If you
have further questions, feel free to [join the **VMware Secrets Manager**
community on **Slack**][slack-invite] and ask them out.

[slack-invite]: https://join.slack.com/t/a-101-103-105-s/shared_invite/zt-287dbddk7-GCX495NK~FwO3bh_DAMAtQ "Join VSecM Slack"

{{ edit() }}
