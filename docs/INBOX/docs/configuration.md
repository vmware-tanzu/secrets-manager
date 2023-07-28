---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, architecture, configuration, environment
title: Configuring Aegis
description: buttons, levers, knobs, nuts, and bolts…
micro_nav: false
page_nav:
  prev:
    content: <strong>Aegis</strong> Deep Dive
    url: '/docs/architecture'
  next:
    content: design decisions
    url: '/docs/philosophy'
---

<p style="text-align:right;position:relative;top:-40px;"
><a href="https://github.com/ShieldWorks/aegis-web/blob/main/docs/configuration.md"
style="border-bottom: none;background:#e0e0e0;padding:0.5em;display:inline-block;
border-radius:8px;">
edit this page on <strong>GitHub</strong> ✏️</a></p>

## Introduction

**Aegis** system components can be configured using environment variables.

The following section contain a breakdown of all of these environment variables.

> **Looking for Aegis Production Tips**?
> 
> For **production setup**, check out [**Aegis Production Deployment**](/production).

## Environment Variables

### SPIFFE_ENDPOINT_SOCKET

`SPIFFE_ENDPOINT_SOCKET` is required for **Aegis Sentinel** to talk to
**Aegis SPIRE**.

If not provided, a default value of `"unix:///spire-agent-socket/agent.sock"`
will be used.

### AEGIS_LOG_LEVEL

`AEGIS_LOG_LEVEL` determines the verbosity of the logs in **Aegis Safe**.

**Aegis Sidecar** also uses this configuration; however, unlike **Aegis Safe**, 
it is not dynamic. While you can dynamically configure this at runtime for 
**Aegis Safe** without having to restart **Aegis Safe**, for **Aegis Sidecar** 
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

### AEGIS_WORKLOAD_SVID_PREFIX

Both **Aegis Safe** and **workloads** use this environment variable.

`AEGIS_WORKLOAD_SVID_PREFIX` is required for validation. If not provided,
it will default to: `"spiffe://aegis.ist/workload/"`

### AEGIS_SENTINEL_SVID_PREFIX

Both **Aegis Safe** and **Aegis Sentinel** use this environment variable.

`AEGIS_SENTINEL_SVID_PREFIX` is required for validation.

If not provided, it will default to:
`"spiffe://aegis.ist/workload/aegis-sentinel/ns/aegis-system/sa/aegis-sentinel/n/"`

### AEGIS_SAFE_FIPS_COMPLIANT

`AEGIS_SAFE_FIPS_COMPLIANT` is required for **Aegis Safe** to run in FIPS-compliant
mode. Defaults to `"false"`, which means **Aegis Safe** will run in non-FIPS-compliant
mode. Setting it to `"true"` will make **Aegis Safe** run in FIPS-compliant mode.

Note that this is not a guarantee that Aegis Safe will actually
run in FIPS compliant mode, as it depends on the underlying base image.

If you are using one of the official FIPS-complaint Aegis Docker images,
then it will be FIPS-compliant.

As a FIPS-compliant base image you can choose from the following:

* [aegishub/aegis-ist-fips-safe][aegis-safe-istanbul-fips] (*using a Distroless base*)
* [aegishub/aegis-photon-fips-safe][aegis-safe-photon-fips] (*using VMware Photon OS as a base*)

[aegis-safe-istanbul-fips]: https://hub.docker.com/repository/docker/aegishub/aegis-ist-fips-safe/general
[aegis-safe-photon-fips]: https://hub.docker.com/repository/docker/aegishub/aegis-photon-fips-safe/general

### AEGIS_SAFE_SVID_PREFIX

Both **Aegis Sentinel**, **Aegis Safe**, and **workloads** use this environment
variable.

`AEGIS_SAFE_SVID_PREFIX` is required for validation.

If not provided, it will default to:
`"spiffe://aegis.ist/workload/aegis-safe/ns/aegis-system/sa/aegis-safe/n/"`

### AEGIS_SAFE_DATA_PATH

`AEGIS_SAFE_DATA_PATH` is where **Aegis Safe** stores its encrypted secrets.

If not given, defaults to `"/data"`.

### AEGIS_CRYPTO_KEY_PATH

`AEGIS_CRYPTO_KEY_PATH` is where **Aegis Safe** will fetch the `"key.txt"`
that contains the encryption keys.

If not given, it will default to `"/key/key.txt"`.

### AEGIS_CRYPTO_KEY_NAME

`AEGIS_CRYPTO_KEY_NAME` is how the age secret key is referenced by
name inside **Aegis Safe**’s code. If not set, defaults to `"aegis-safe-age-key"`.

If you change the value of this environment variable, make sure to change the 
relevant `Secret` and `Deployment` YAML manifests too. The easiest way to do 
this is to do a project wide search and find and replace places where reference
`"aegis-safe-age-key"` to your new name of choice.

### AEGIS_MANUAL_KEY_INPUT

`AEGIS_MANUAL_KEY_INPUT` is used to tell **Aegis Safe** to bypass generating
the master crypto key and instead use the key provided by the operator using
**Aegis Sentinel**. Defaults to `"false"`.

When set to `"true"`, **Aegis Safe** will **not** store the master key in a 
Kubernetes `Secret`; the master key will reside soley in the memory of
**Aegis Safe**. 

The control offered by this approach regarding the threat boundary of
**Aegis Safe** provides enhanced security compared to the default behavior where 
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

Also note that when this variable is set to `"true"`, **Aegis Safe** will **not** 
respond to API requests until a master key is provided, using **Aegis Sentinel**.

### AEGIS_SAFE_SECRET_NAME_PREFIX

`AEGIS_SAFE_SECRET_NAME_PREFIX` is the prefix that is used to prepend to the
secret names that **Aegis Safe** stores in the cluster as `Secret` objects when
the `-k` option in **Aegis Sentinel** is selected.

If this variable is not set or is empty, the default value `"aegis-secret-"` 
is used.

### AEGIS_SAFE_ENDPOINT_URL

`AEGIS_SAFE_ENDPOINT_URL` is the **REST API** endpoint that **Aegis Safe**
exposes from its `Service`.

**Aegis Sentinel**, **Aegis Sidecar** and
**workloads** need this URL configured.

If not provided, it will default to:
`"https://aegis-safe.aegis-system.svc.cluster.local:8443/"`.

### AEGIS_PROBE_LIVENESS_PORT

**Aegis Safe** and **Aegis Sentinel** use this configuration.

`AEGIS_PROBE_LIVENESS_PORT` is the port where the liveness probe
will serve.

Defaults to `:8081`.

### AEGIS_PROBE_READINESS_PORT

**Aegis Safe** uses this configuration.

`AEGIS_PROBE_READINESS_PORT` is the port where the readiness probe
will serve.

Defaults to `:8082`.

### AEGIS_SAFE_SVID_RETRIEVAL_TIMEOUT

**Aegis Safe** uses this configuration.

`AEGIS_SAFE_SVID_RETRIEVAL_TIMEOUT` is how long (*in milliseconds*) **Aegis Safe**
will wait for an *SPIRE X.509 SVID* bundle before giving up and crashing.

The default value is `30000` milliseconds.

### AEGIS_SAFE_TLS_PORT

`AEGIS_SAFE_TLS_PORT` is the port that **Aegis Safe** serves its API endpoints.

When you change this port, you will likely need to make changes in more 
than one manifest, and restart or redeploy **Aegis** and **SPIRE**.

Defaults to `":8443"`.

### AEGIS_SAFE_SECRET_BUFFER_SIZE

`AEGIS_SAFE_SECRET_BUFFER_SIZE` is the amount of secret insertion operations
to be buffered until **Aegis Safe API** blocks and waits for the buffer to have 
an empty slot.

If the environment variable is not set, this buffer size defaults to `10`.

Two separate buffers of the same size are used for IO operations, and 
Kubernetes `Secret` creation (*depending on the type of the API request*). The
Kubernetes Secrets buffer, and File IO buffer work asynchronously and
independent of each other int two separate goroutines.

### AEGIS_SAFE_K8S_SECRET_BUFFER_SIZE

`AEGIS_SAFE_K8S_SECRET_BUFFER_SIZE` is the buffer size for the **Aegis Safe**
Kubernetes secret queue.

If the environment variable is not set, the default buffer size is `10`.

### AEGIS_SAFE_SECRET_DELETE_BUFFER_SIZE

`AEGIS_SAFE_SECRET_DELETE_BUFFER_SIZE` isd the buffer size for the **Aegis Safe** 
secret deletion queue.

If the environment variable is not set, the default buffer size is `10`.

### AEGIS_SAFE_K8S_SECRET_DELETE_BUFFER_SIZE

`AEGIS_SAFE_K8S_SECRET_DELETE_BUFFER_SIZE` the buffer size for the 
**Aegis Safe** Kubernetes secret deletion queue.

If the environment variable is not set, the default buffer size is `10`.

### AEGIS_SAFE_BACKING_STORE

This environment variable is used by **Aegis Sentinel** to let **Aegis Safe**
know where to persist the secret. To reiterate, this environment variable shall
be defined for **Aegis Sentinel** deployment; defining it for **Aegis Safe**
has no effect.

`AEGIS_SAFE_BACKING_STORE` is the type of the storage where the secrets
will be encrypted and persisted. 

* If not given, defaults to `"file"`. 
* The other option is  `"in-memory"`.

A `"file"` backing store means **Aegis Safe** persists an encrypted version
of its state in a volume (*ideally a `PersistedVolume`*).

An `"in-memory"` backing store means **Aegis Safe** does not persist backups
of the secrets it created to disk. When that option is selected, you will
lose all of your secrets if **Aegis Safe** is evicted by the scheduler or
manually restarted by an operator.

### AEGIS_SAFE_SECRET_BACKUP_COUNT

`AEGIS_SAFE_SECRET_BACKUP_COUNT` indicates the number of backups to keep for
**Aegis Safe** secrets. 

If the environment variable AEGIS_SAFE_SECRET_BACKUP_COUNT is not set or is not 
a valid integer, the default value of `"3"` will be used.

This configuration is **not** effective when `AEGIS_SAFE_BACKING_STORE` is
set to `"in-memory"`.

### AEGIS_SAFE_USE_KUBERNETES_SECRETS

`AEGIS_SAFE_USE_KUBERNETES_SECRETS` is a flag indicating whether to create a
plain text Kubernetes secret for the workloads registered. 

If the environment variable is not set or its value is not `"true"`, it will
be assumed `"false"`.

There are two things to note about this approach:

First, by design, and for security reasons, the original Kubernetes `Secret` 
should exist, and it should be initiated to a default data as follows before
it can be synced by **Aegis Safe**:

```text
{% raw %}apiVersion: v1
kind: Secret
metadata:
  # The string after `aegis-secret-` must match the workload’s name.
  # For example, this is an Aegis-managed secret for `example`
  # with the SPIFFE ID 
  # `"spiffe://aegis.ist/workload/example\
  #  /ns/{{ .PodMeta.Namespace }}\
  #  /sa/{{ .PodSpec.ServiceAccountName }}\
  #  /n/{{ .PodMeta.Name }}"`
  name: aegis-secret-example
  namespace: default
type: Opaque{% endraw %}
```

Secondly this approach is **less** secure, and it is meant to be used for 
**legacy** systems where directly using the **Safe Sidecar** or 
**Safe SDK** are not feasible. For example, you might not have direct control
over the source code to enable a tighter **Safe** integration. Or, you might
temporarily want to establish behavior parity of your legacy system before 
starting a more canonical **Aegis** implementation.

### AEGIS_SIDECAR_POLL_INTERVAL

`AEGIS_SIDECAR_POLL_INTERVAL` is the interval (*in milliseconds*) 
that the sidecar polls **Aegis Safe** for new secrets. 

Defaults to `20000` milliseconds, if not provided.

### AEGIS_SIDECAR_MAX_POLL_INTERVAL

**Aegis Sidecar** has an **exponential backoff** algorithm to execute fetch
in longer intervals when an error occurs. `AEGIS_SIDECAR_MAX_POLL_INTERVAL`
is the maximum wait time (*in milliseconds*) before executing the next.

Defaults to `300000` milliseconds, if not provided.

### AEGIS_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER

**Aegis Sidecar** uses this environment variable.

`AEGIS_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER` configures how fast the algorithm
backs off when there is a failure. Defaults to `2`, which means when there are
enough failures to trigger a backoff, the next wait interval will be twice the
current one.

### AEGIS_SIDECAR_SUCCESS_THRESHOLD

**Aegis Sidecar** uses this environment variable.

`AEGIS_SIDECAR_SUCCESS_THRESHOLD` configures the number of successful poll 
results before reducing the poll interval. Defaults to `3`.

The next interval is calculated by dividing the current interval with
`AEGIS_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER`.

### AEGIS_SIDECAR_ERROR_THRESHOLD

**Aegis Sidecar** uses this environment variable.

`AEGIS_SIDECAR_ERROR_THRESHOLD` configures the number of fetch failures before
increasing the poll interval. Defaults to `2`.

The next interval is calculated by multiplying the current interval with
`AEGIS_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER`.
