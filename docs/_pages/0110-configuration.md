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

title: Configuring VSecM
layout: post
prev_url: /docs/sdk/
permalink: /docs/configuration/
next_url: /docs/use-the-source/
---

<p class="github-button"
><a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0110-configuration.md"
>edit this page on <strong>GitHub</strong> ✏️</a></p>

## Introduction

**VMware Secrets Manager** system components can be configured using environment variables.

The following section contain a breakdown of all of these environment variables.

> **Looking for VMware Secrets Manager Production Tips**?
>
> For **production setup**, check out [**VMware Secrets Manager Production Deployment**](/docs/production).
{: .block-warning}

## Environment Variables

> **Using VSecM Helm Charts**?
> 
> If you are using [**VMware Secrets Manager Helm Charts**][helm-charts],
> you can configure these environment variables using the `values.yaml` file.
{: .block-tip}

[helm-charts]: https://vmware-tanzu.github.io/secrets-manager/


### SPIFFE_ENDPOINT_SOCKET

`SPIFFE_ENDPOINT_SOCKET` is required for **VSecM Sentinel** to talk to
**SPIRE**.

If not provided, a default value of `"unix:///spire-agent-socket/agent.sock"`
will be used.

### VSECM_LOG_LEVEL

`VSECM_LOG_LEVEL` determines the verbosity of the logs in **VSecM Safe**.

**VSecM Sidecar** also uses this configuration; however, unlike **VSecM Safe**,
it is not dynamic. While you can dynamically configure this at runtime for
**VSecM Safe** without having to restart **VSecM Safe**, for **VSecM Sidecar**
you’ll have to restart the **workload**’s pod for any changes to take effect.

`0`: logs are off, `7`: highest verbosity. default: `3`

Here are what various log levels correspond to:

```text
Off   = 0
Fatal = 1
Error = 2
Warn  = 3
Info  = 4
Audit = 5
Debug = 6
Trace = 7
```

### VSECM_WORKLOAD_SVID_PREFIX

Both **VSecM Safe** and **workloads** use this environment variable.

`VSECM_WORKLOAD_SVID_PREFIX` is required for validation. If not provided,
it will default to: `"spiffe://vsecm.com/workload/"`

### VSECM_SENTINEL_SVID_PREFIX

Both **VSecM Safe** and **VSecM Sentinel** use this environment variable.

`VSECM_SENTINEL_SVID_PREFIX` is required for validation.

If not provided, it will default to:
`"spiffe://vsecm.com/workload/vsecm-sentinel/ns/vsecm-system/sa/vsecm-sentinel/n/"`

### VSECM_SAFE_FIPS_COMPLIANT

`VSECM_SAFE_FIPS_COMPLIANT` is required for **VSecM Safe** to run in FIPS-compliant
mode. Defaults to `"false"`, which means **VSecM Safe** will run in non-FIPS-compliant
mode. Setting it to `"true"` will make **VSecM Safe** run in FIPS-compliant mode.

> **You Need Host Support for FIPS-Compliant Mode**
> 
> Note that this is not a guarantee that VSecM Safe will actually
> run in FIPS compliant mode, as it depends on the underlying base image.
> 
> In addition, the host environment will need to be compliant too.

If you are using one of the official FIPS-complaint VSecM Docker images,
then it will be FIPS-compliant.

As a FIPS-compliant base image you can choose from the following:

* [vsecm/vsecm-ist-fips-safe][vsecm-safe-istanbul-fips] (*using a Distroless base*)
* [vsecm/vsecm-photon-fips-safe][vsecm-safe-photon-fips] (*using VMware Photon OS as a base*)

[vsecm-safe-istanbul-fips]: https://hub.docker.com/repository/docker/vsecm/vsecm-ist-fips-safe/general
[vsecm-safe-photon-fips]: https://hub.docker.com/repository/docker/vsecm/vsecm-photon-fips-safe/general

### VSECM_SAFE_SVID_PREFIX

Both **VSecM Sentinel**, **VSecM Safe**, and **workloads** use this environment
variable.

`VSECM_SAFE_SVID_PREFIX` is required for validation.

If not provided, it will default to:
`"spiffe://vsecm.com/workload/vsecm-safe/ns/vsecm-system/sa/vsecm-safe/n/"`

### VSECM_SAFE_DATA_PATH

`VSECM_SAFE_DATA_PATH` is where **VSecM Safe** stores its encrypted secrets.

If not given, defaults to `"/data"`.

### VSECM_CRYPTO_KEY_PATH

`VSECM_CRYPTO_KEY_PATH` is where **VSecM Safe** will fetch the `"key.txt"`
that contains the encryption keys.

If not given, it will default to `"/key/key.txt"`.

### VSECM_CRYPTO_KEY_NAME

`VSECM_CRYPTO_KEY_NAME` is how the age secret key is referenced by
name inside **VSecM Safe**’s code. If not set, defaults to `"vsecm-safe-age-key"`.

If you change the value of this environment variable, make sure to change the
relevant `Secret` and `Deployment` YAML manifests too. The easiest way to do
this is to do a project wide search and find and replace places where reference
`"vsecm-safe-age-key"` to your new name of choice.

### VSECM_MANUAL_KEY_INPUT

`VSECM_MANUAL_KEY_INPUT` is used to tell **VSecM Safe** to bypass generating
the master crypto key and instead use the key provided by the operator using
**VSecM Sentinel**. Defaults to `"false"`.

When set to `"true"`, **VSecM Safe** will **not** store the master key in a
Kubernetes `Secret`; the master key will reside solely in the memory of
**VSecM Safe**.

The control offered by this approach regarding the threat boundary of
**VSecM Safe** provides enhanced security compared to the default behavior where
the master key is randomly-generated in a cryptographically secure way and
stored in a Kubernetes `Secret` (*which is still **pretty secure**, especially if
you encrypt your `etcd` and establish a tight RBAC over the Kubernetes `Secret`
that stores the master key*).

That being said, in manual input mode, all the cryptographic material will be kept
in memory, and no Kubernetes `Secret`s will be harmed (*which is even more secure*).

However, this approach does place the onus of securing the master key on you.

Yet, don’t worry, storing the master key can be easily handled depending on your
infrastructure: With the right setup, managing the master key is just another cog
in the machinery of your security measures, not an additional burden.

That’s why, although it’s not enabled by default, adopting this measure could be
an additional step in bolstering your system’s security.

Also note that when this variable is set to `"true"`, **VSecM Safe** will **not**
respond to API requests until a master key is provided, using **VSecM Sentinel**.

### VSECM_SAFE_SECRET_NAME_PREFIX

`VSECM_SAFE_SECRET_NAME_PREFIX` is the prefix that is used to prepend to the
secret names that **VSecM Safe** stores in the cluster as `Secret` objects when
the `-k` option in **VSecM Sentinel** is selected.

If this variable is not set or is empty, the default value `"vsecm-secret-"`
is used.

### VSECM_SAFE_ENDPOINT_URL

`VSECM_SAFE_ENDPOINT_URL` is the **REST API** endpoint that **VSecM Safe**
exposes from its `Service`.

**VSecM Sentinel**, **VSecM Sidecar** and
**workloads** need this URL configured.

If not provided, it will default to:
`"https://vsecm-safe.vsecm-system.svc.cluster.local:8443/"`.

### VSECM_PROBE_LIVENESS_PORT

**VSecM Safe** and **VSecM Sentinel** use this configuration.

`VSECM_PROBE_LIVENESS_PORT` is the port where the liveness probe
will serve.

Defaults to `:8081`.

### VSECM_PROBE_READINESS_PORT

**VSecM Safe** uses this configuration.

`VSECM_PROBE_READINESS_PORT` is the port where the readiness probe
will serve.

Defaults to `:8082`.

### VSECM_SAFE_SVID_RETRIEVAL_TIMEOUT

**VSecM Safe** uses this configuration.

`VSECM_SAFE_SVID_RETRIEVAL_TIMEOUT` is how long (*in milliseconds*) **VSecM Safe**
will wait for an *SPIRE X.509 SVID* bundle before giving up and crashing.

The default value is `30000` milliseconds.

### VSECM_SAFE_TLS_PORT

`VSECM_SAFE_TLS_PORT` is the port that **VSecM Safe** serves its API endpoints.

When you change this port, you will likely need to make changes in more
than one manifest, and restart or redeploy **VMware Secrets Manager** and **SPIRE**.

Defaults to `":8443"`.

### VSECM_SAFE_SECRET_BUFFER_SIZE

`VSECM_SAFE_SECRET_BUFFER_SIZE` is the amount of secret insertion operations
to be buffered until **VSecM Safe API** blocks and waits for the buffer to have
an empty slot.

If the environment variable is not set, this buffer size defaults to `10`.

Two separate buffers of the same size are used for IO operations, and
Kubernetes `Secret` creation (*depending on the type of the API request*). The
Kubernetes Secrets buffer, and File IO buffer work asynchronously and
independent of each other int two separate goroutines.

### VSECM_SAFE_K8S_SECRET_BUFFER_SIZE

`VSECM_SAFE_K8S_SECRET_BUFFER_SIZE` is the buffer size for the **VSecM Safe**
Kubernetes secret queue.

If the environment variable is not set, the default buffer size is `10`.

### VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE

`VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE` isd the buffer size for the **VSecM Safe**
secret deletion queue.

If the environment variable is not set, the default buffer size is `10`.

### VSECM_SAFE_K8S_SECRET_DELETE_BUFFER_SIZE

`VSECM_SAFE_K8S_SECRET_DELETE_BUFFER_SIZE` the buffer size for the
**VSecM Safe** Kubernetes secret deletion queue.

If the environment variable is not set, the default buffer size is `10`.

### VSECM_SAFE_BACKING_STORE

This environment variable is used by **VSecM Sentinel** to let **VSecM Safe**
know where to persist the secret. To reiterate, this environment variable shall
be defined for **VSecM Sentinel** deployment; defining it for **VSecM Safe**
has no effect.

`VSECM_SAFE_BACKING_STORE` is the type of the storage where the secrets
will be encrypted and persisted.

* If not given, defaults to `"file"`.
* The other option is  `"in-memory"`.

A `"file"` backing store means **VSecM Safe** persists an encrypted version
of its state in a volume (*ideally a `PersistedVolume`*).

An `"in-memory"` backing store means **VSecM Safe** does not persist backups
of the secrets it created to disk. When that option is selected, you will
lose all of your secrets if **VSecM Safe** is evicted by the scheduler or
manually restarted by an operator.

### VSECM_SAFE_SECRET_BACKUP_COUNT

`VSECM_SAFE_SECRET_BACKUP_COUNT` indicates the number of backups to keep for
**VSecM Safe** secrets.

If the environment variable VSECM_SAFE_SECRET_BACKUP_COUNT is not set or is not
a valid integer, the default value of `"3"` will be used.

This configuration is **not** effective when `VSECM_SAFE_BACKING_STORE` is
set to `"in-memory"`.

### VSECM_SAFE_USE_KUBERNETES_SECRETS

`VSECM_SAFE_USE_KUBERNETES_SECRETS` is a flag indicating whether to create a
plain text Kubernetes secret for the workloads registered.

If the environment variable is not set or its value is not `"true"`, it will
be assumed `"false"`.

There are two things to note about this approach:

First, by design, and for security reasons, the original Kubernetes `Secret`
should exist, and it should be initiated to a default data as follows before
it can be synced by **VSecM Safe**:

```yaml
{% raw %}apiVersion: v1
kind: Secret
metadata:
  # The string after `vsecm-secret-` must match the 
  # workload’s name.
  # For example, this is an VSecM-managed secret for `example`
  # with the SPIFFE ID 
  # `"spiffe://vsecm.com/workload/example\
  #  /ns/{{ .PodMeta.Namespace }}\
  #  /sa/{{ .PodSpec.ServiceAccountName }}\
  #  /n/{{ .PodMeta.Name }}"`
  name: vsecm-secret-example
  namespace: default
type: Opaque{% endraw %}
```

Secondly this approach is **less** secure, and it is meant to be used for
**legacy** systems where directly using the **Safe Sidecar** or
**Safe SDK** are not feasible. For example, you might not have direct control
over the source code to enable a tighter **Safe** integration. Or, you might
temporarily want to establish behavior parity of your legacy system before
starting a more canonical **VMware Secrets Manager** implementation.

### VSECM_SIDECAR_POLL_INTERVAL

`VSECM_SIDECAR_POLL_INTERVAL` is the interval (*in milliseconds*)
that the sidecar polls **VSecM Safe** for new secrets.

Defaults to `20000` milliseconds, if not provided.

### VSECM_SIDECAR_MAX_POLL_INTERVAL

**VSecM Sidecar** has an **exponential backoff** algorithm to execute fetch
in longer intervals when an error occurs. `VSECM_SIDECAR_MAX_POLL_INTERVAL`
is the maximum wait time (*in milliseconds*) before executing the next.

Defaults to `300000` milliseconds, if not provided.

### VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER

**VSecM Sidecar** uses this environment variable.

`VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER` configures how fast the algorithm
backs off when there is a failure. Defaults to `2`, which means when there are
enough failures to trigger a backoff, the next wait interval will be twice the
current one.

### VSECM_SIDECAR_SUCCESS_THRESHOLD

**VSecM Sidecar** uses this environment variable.

`VSECM_SIDECAR_SUCCESS_THRESHOLD` configures the number of successful poll
results before reducing the poll interval. Defaults to `3`.

The next interval is calculated by dividing the current interval with
`VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER`.

### VSECM_SIDECAR_ERROR_THRESHOLD

**VSecM Sidecar** uses this environment variable.

`VSECM_SIDECAR_ERROR_THRESHOLD` configures the number of fetch failures before
increasing the poll interval. Defaults to `2`.

The next interval is calculated by multiplying the current interval with
`VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER`.

### VSECM_SYSTEM_NAMESPACE

`VSECM_SYSTEM_NAMESPACE` environment variable specifies the namespace in
which a VSecM instance is deployed.

Ensure this is set as an environment variable for your containers; it's a
critical piece. VSecM Safe and Sentinel rely on it to precisely locate the
deployment's namespace. For instance, Safe leverages this information to securely
store age keys within a designated secret, as specified by the `VSECM_CRYPTO_KEY_NAME`
configuration.
