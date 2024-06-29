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

title = "Configuring VSecM"
weight = 10
+++

> **Alphabetic Order**
>
> The environment variables mentioned in this document are listed in alphabetic
> order.

> **Looking for VMware Secrets Manager Production Tips**?
>
> For **production setup**, check out [**VMware Secrets Manager Production
> Deployment**][prod].

## Introduction

**VMware Secrets Manager** system components can be configured using environment
variables.

The following section contain a breakdown of all of these environment variables.

> **With Great Flexibility Comes Responsibility**
>
> Using environment variables to configure **VSecM** offers operators 
> significant flexibility. However, it's important to be mindful that 
> incorrect configurations can potentially reduce your apps' security.
>
> For detailed guidance on secure configurations, please refer to 
> the [**Production Deployment Instructions**][prod].

[prod]: @/documentation/production/overview.md

## Environment Variables

> **Using VSecM Helm Charts**?
>
> If you are using [**VMware Secrets Manager Helm Charts**][helm-charts],
> you can configure these environment variables using the `values.yaml` file.
>
> You can check out the details [from the official VSecM Helm Charts
> documentation][helm-charts].

[helm-charts]: https://vmware-tanzu.github.io/secrets-manager/

### SPIFFE_ENDPOINT_SOCKET

**Used By**: *VSecM Sentinel*, *VSecM Sidecar*,
*VSecM Init Container*, *VSecM Safe*, *Workloads*.

`SPIFFE_ENDPOINT_SOCKET` is required for **VSecM Sentinel** to talk to
**SPIRE**.

If not provided, a default value of `"unix:///spire-agent-socket/agent.sock"`
will be used.

### SPIFFE_TRUST_DOMAIN

**Used By**: *VSecM Safe*, *VSecM Sentinel*, *VSecM Sidecar*, *Workloads*.

`SPIFFE_TRUST_DOMAIN` specifies which trust domain that the Kubernetes cluster
**VSecM** deployed on is configured with.

By default, it is `"vsecm.com"`.

The workload SPIFFE IDs will have this trust domain as per the SPIFFE
standard.

For example, for the SPIFFE ID 
`spiffe://vsecm.com/workload/app/s/app/n/app-1234`, `vsecm.com` is the 
*trust domain*.

For more information about SPIFFE IDs and trust domains, 
[check out the official SPIFFE documentation][spiffe-docs].

[spiffe-docs]: https://spiffe.io/

### VSECM_BACKOFF_DELAY

**Used By**: *VSecM Sentinel*, *VSecM Safe*, *VSecM Sidecar*, 
*VSecM Init Container*.

`VSECM_BACKOFF_DELAY` configures the initial delay for backoff algorithms. This
introduces a waiting period before retrying an operation after a failure.

If not given, defaults to `"1000"` milliseconds.``

### VSECM_BACKOFF_MAX_RETRIES

**Used By**: *VSecM Sentinel*, *VSecM Safe*, *VSecM Sidecar*, 
*VSecM Init Container*.

`VSECM_BACKOFF_MAX_RETRIES` configures the maximum number of retries in internal
backoff algorithms. This is used in scenarios where operations might fail
transiently and require repeated attempts to succeed.

If not given, defaults to `"10"`.

### VSECM_BACKOFF_MAX_WAIT

**Used By**: *VSecM Sentinel*, *VSecM Safe*, *VSecM Sidecar*, 
*VSecM Init Container*.

`VSECM_BACKOFF_MAX_WAIT` configures the maximum duration the backoff
algorithm can wait before retrying an operation. This is used to prevent
excessive waiting times in case of repeated failures.

If the environment variable is not set or cannot be parsed, a default maximum
duration of `"30000"` milliseconds is used.

This variable is crucial for defining the upper limit on the duration to which
backoff delay can grow, ensuring that retry mechanisms do not result in
excessively long wait times.

### VSECM_BACKOFF_MODE

**Used By**: *VSecM Sentinel*, *VSecM Safe*, *VSecM Sidecar*, 
*VSecM Init Container*.

`VSECM_BACKOFF_MODE` determines the backoff strategy to be used. Available
options are `"exponential"` and `"linear"`.

If the environment variable is not set, or if its value is `"exponential"`,
`"exponential"` is used. For any other non-empty value, `"linear"` is used.

### VSECM_INIT_CONTAINER_POLL_INTERVAL

**Used By**: *VSecM Init Container*.

`VSECM_INIT_CONTAINER_POLL_INTERVAL` determines the time interval between each
poll in the `Watch()` function. The interval is specified in milliseconds.
If the environment variable is not set or is not a valid integer value,
a default interval of `5000` milliseconds is used.

### VSECM_INIT_CONTAINER_WAIT_BEFORE_EXIT

**Used By**: *VSecM Init Container*.

`VSECM_INIT_CONTAINER_WAIT_BEFORE_EXIT` determines the wait time before the
init container exits. The duration is determined by the environment variable
`"VSECM_INIT_CONTAINER_WAIT_BEFORE_EXIT"` and defaults to zero if the variable
is not set or cannot be parsed.

The environment variable is expected to be an integer value representing the
wait time in milliseconds. If parsing fails, the function a default value
of `"0"` milliseconds will be assumed.

### VSECM_KEYGEN_DECRYPT

**Used By**: *VSecM Keygen*.

`VSECM_KEYGEN_DECRYPT` determines if **VSecM Keygen** should decrypt the secrets
JSON file instead of generation a root key (*which is its default behavior*).

If this value is  anything but `"true"`, **VSecM Keygen** will generate a new
root key. Otherwise, it will attempt to decrypt the secrets provided to it.

Defaults to `"false"`.

### VSECM_KEYGEN_EXPORTED_SECRET_PATH

**Used By**: *VSecM Keygen*.

`VSECM_KEYGEN_EXPORTED_SECRET_PATH` is the path where the exported secrets are
stored. This needs to be mounted to the container that you run **VSecM Keygen**.

If the environment variable is not set, it defaults to
`"/opt/vsecm/secrets.json"`.

### VSECM_KEYGEN_ROOT_KEY_PATH

**Used By**: *VSecM Keygen*.

`VSECM_KEYGEN_ROOT_KEY_PATH` is the path where the root key is stored. This
needs to be mounted to the container that you run **VSecM Keygen**.

If not given, it defaults to `"/opt/vsecm/keys.txt"`.

### VSECM_LOG_LEVEL

**Used By**: *VSecM Sentinel*, *VSecM Safe*.

`VSECM_LOG_LEVEL` determines the verbosity of the logs in **VSecM Safe**.

**VSecM Sidecar** also uses this configuration; however, unlike **VSecM Safe**,
it is not dynamic. While you can dynamically configure this at runtime for
**VSecM Safe** without having to restart **VSecM Safe**, for **VSecM Sidecar**
you'll have to restart the **workload**'s pod for any changes to take effect.

`0`: logs are off, `7`: highest verbosity. default: `3`

Here are what various log levels correspond to:

```txt
Off   = 0
Fatal = 1
Error = 2
Warn  = 3
Info  = 4
Audit = 5
Debug = 6
Trace = 7
```

### VSECM_LOG_SECRET_FINGERPRINTS

**Used By**: *VSecM Safe*.

`VSECM_LOG_SECRET_FINGERPRINTS` is used to determine whether secret fingerprints
should be logged in the **VSecM Safe** logs.

If the environment variable is not set or is not `"true"`, defaults to `"false"`.

When `"true"`, **VSecM Safe** logs will include partial hashes for the secrets.
This approach will be useful to verify changes to a secret without revealing it
in the logs.

The partial hash is a cryptographically secure string, and there is no way to
retrieve the original secret from it.

### VSECM_NAMESPACE_SYSTEM

**Used By**: *VSecM Safe*.

`VSECM_NAMESPACE_SYSTEM` environment variable specifies the namespace in
which a **VSecM** instance is deployed.

Ensure this is set as an environment variable for your containers; it's a
critical piece. **VSecM Safe** and Sentinel rely on it to precisely locate the
deployment's namespace. For instance, Safe leverages this information to
securely store age keys within a designated secret, as specified by the
`VSECM_ROOT_KEY_NAME` configuration.

### VSECM_PROBE_LIVENESS_PORT

**Used By**: *VSecM Sentinel*, *VSecM Safe*.

**VSecM Safe** and **VSecM Sentinel** use this configuration.

`VSECM_PROBE_LIVENESS_PORT` is the port where the liveness probe
will serve.

Defaults to `:8081`.

### VSECM_PROBE_READINESS_PORT

**Used By**: *VSecM Safe*.

**VSecM Safe** uses this configuration.

`VSECM_PROBE_READINESS_PORT` is the port where the readiness probe
will serve.

Defaults to `:8082`.

### VSECM_ROOT_KEY_INPUT_MODE_MANUAL

**Used By**: *VSecM Safe*.

`VSECM_ROOT_KEY_INPUT_MODE_MANUAL` is a boolean indicating whether to use manual
cryptographic key input for **VSecM Safe**, instead of letting the bootstrap
flow automatically compute cryptographic keys.

If the environment variable is not set or its value is not `"true"`, the
bootstrap flow will **not** compute the cryptographic keys automatically.

If this variable is set to `"true"` then a human operator has to provide the
necessary cryptographic keys using **VSecM Sentinel** for **VSecM** Safe** to
unlock itself and start serving API requests.

> **Setting the Root Key Manually**
>
> You can [set the root key programmatically using
> **VSecM Sentinel**][vsecm-sentinel-root].

[vsecm-sentinel-root]: @/documentation/usage/cli.md

The control offered by this approach changes the threat boundary of
**VSecM Safe**. With this approach, the responsibility of securing the
*root key* is on you as the operator. This is different than the default
behavior of **VSecM Safe** where the *root key* is randomly-generated in a
cryptographically secure way and stored in a Kubernetes `Secret`.

Using a Kubernetes `Secret` to store the *root key* is still secure,
especially if you encrypt your `etcd` and establish a **tight RBAC** over the
Kubernetes `Secret`that stores the *root key*.

Also note that when this variable is set to `"true"`, **VSecM Safe** will **not**
respond to API requests until a *root key* is provided, using **VSecM Sentinel**.

[vsecm-sentinel-cli]: @/documentation/usage/cli.md

### VSECM_ROOT_KEY_NAME

**Used By**: *VSecM Safe*.

`VSECM_ROOT_KEY_NAME` is how the age secret key is referenced by
name inside **VSecM Safe**'s code. If not set, defaults to `"vsecm-root-key"`.

If you change the value of this environment variable, make sure to change the
relevant `Secret` and `Deployment` YAML manifests too. The easiest way to do
this is to do a project wide search and find and replace places where reference
`"vsecm-root-key"` to your new name of choice.

### VSECM_ROOT_KEY_PATH

**Used By**: *VSecM Safe*.

`VSECM_ROOT_KEY_PATH` is where **VSecM Safe** will fetch the `"key.txt"`
that contains the encryption keys.

If not given, it will default to `"/key/key.txt"`.

### VSECM_SAFE_BACKING_STORE

**Used By**: *VSecM Safe*.

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

### VSECM_SAFE_BOOTSTRAP_TIMEOUT

**Used By**: *VSecM Safe*.

**VSecM Safe** uses this configuration.

`VSECM_SAFE_BOOTSTRAP_TIMEOUT` is how long (*in milliseconds*) **VSecM Safe**
will wait for an *SPIRE X.509 SVID* bundle before giving up and crashing.

The default value is `300000` milliseconds.

### VSECM_SAFE_DATA_PATH

**Used By**: *VSecM Safe*.

`VSECM_SAFE_DATA_PATH` is where **VSecM Safe** stores its encrypted secrets.

If not given, defaults to `"/var/local/vsecm/data"`.

### VSECM_SAFE_ENDPOINT_URL

**Used By**: *VSecM Sentinel*, *VSecM Sidecar*, *VSecM Init Container*,
*VSecM Safe*, *Workloads*.

`VSECM_SAFE_ENDPOINT_URL` is the **REST API** endpoint that **VSecM Safe**
exposes from its `Service`.

**VSecM Sentinel**, **VSecM Sidecar** and
**workloads** need this URL configured.

If not provided, it will default to:
`"https://vsecm-safe.vsecm-system.svc.cluster.local:8443/"`.

### VSECM_SAFE_FIPS_COMPLIANT

**Used By**: *VSecM Safe*.

`VSECM_SAFE_FIPS_COMPLIANT` is required for **VSecM Safe** to run in 
FIPS-compliant mode. Defaults to `"false"`, which means **VSecM Safe** will run 
in non-FIPS-compliant mode. Setting it to `"true"` will make **VSecM Safe** run 
in FIPS-compliant mode.

> **You Need Host Support for FIPS-Compliant Mode**
>
> Note that this is not a guarantee that **VSecM Safe** will actually
> run in FIPS compliant mode, as it depends on the underlying base image.
>
> In addition, the host environment will need to be compliant too.
>
> Furthermore, compliance with FIPS (*Federal Information Processing Standards*)
> requires not just a technically compliant binary and system but also
> **formal approval** by the United States National Institute of Standards and
> Technology (*NIST*). This involves a certification process where **NIST**
> validates that the cryptographic modules used in the software meet the
> **FIPS 140-2 standard**.
>
> Therefore, even if the binary and the underlying system are FIPS-compliant,
> **VSecM Safe** will still need to undergo this formal FIPS approval process
> to be officially recognized as FIPS compliant.

If you are using one of the official FIPS-complaint VSecM Docker images,
then it will be FIPS-compliant.

As a FIPS-compliant base image you can use the following:

* [vsecm/vsecm-ist-fips-safe][vsecm-safe-istanbul-fips] (*using a Distroless 
  base*)

[vsecm-safe-istanbul-fips]: https://hub.docker.com/repository/docker/vsecm/vsecm-ist-fips-safe/general

### VSECM_SAFE_IV_INITIALIZATION_INTERVAL

**Used By**: *VSecM Safe*.

`VSECM_SAFE_IV_INITIALIZATION_INTERVAL` is used as a security measure to
time-based attacks where too frequent call of a function can be used to
generate less-randomized AES IV values.

If the environment variable is not set or contains an invalid integer, it
defaults to `50` milliseconds.

The value in the environment variable is in milliseconds.

### VSECM_SAFE_K8S_SECRET_BUFFER_SIZE

**Used By**: *VSecM Safe*.

`VSECM_SAFE_K8S_SECRET_BUFFER_SIZE` is the buffer size for the **VSecM Safe**
Kubernetes secret queue.

If the environment variable is not set, the default buffer size is `10`.

### VSECM_SAFE_ROOT_KEY_STORE

**Used By**: *VSecM Safe*.

`VSECM_SAFE_ROOT_KEY_STORE` determines the root key store type for **VSecM
Safe**.

As of this version, the only valid value is `"k8s"`, meaning that the **VSecM
Safe** will store its root keys in a Kubernetes Secret.

Defaults to `"k8s"` when not provided.

### VSECM_SAFE_SECRET_BACKUP_COUNT

**Used By**: *VSecM Safe*.

`VSECM_SAFE_SECRET_BACKUP_COUNT` indicates the number of backups to keep for
**VSecM Safe** secrets.

If the environment variable VSECM_SAFE_SECRET_BACKUP_COUNT is not set or is not
a valid integer, the default value of `"3"` will be used.

This configuration is **not** effective when `VSECM_SAFE_BACKING_STORE` is
set to `"in-memory"`.

> **Plan to Deprecate**
>
> There are plans to deprecate this feature in the future in favor of
> a more robust database-driven changelog solution for secrets.

### VSECM_SAFE_SECRET_BUFFER_SIZE

**Used By**: *VSecM Safe*.

`VSECM_SAFE_SECRET_BUFFER_SIZE` is the amount of secret insertion operations
to be buffered until **VSecM Safe API** blocks and waits for the buffer to have
an empty slot.

If the environment variable is not set, this buffer size defaults to `10`.

Two separate buffers of the same size are used for IO operations, and
Kubernetes `Secret` creation (*depending on the type of the API request*). The
Kubernetes Secrets buffer, and File IO buffer work asynchronously and
independent of each other int two separate goroutines.

### VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE

**Used By**: *VSecM Safe*.

`VSECM_SAFE_SECRET_DELETE_BUFFER_SIZE` isd the buffer size for the **VSecM Safe**
secret deletion queue.

If the environment variable is not set, the default buffer size is `10`.

### VSECM_SAFE_SOURCE_ACQUISITION_TIMEOUT

**Used By**: *VSecM Safe*.

`VSECM_SAFE_SOURCE_ACQUISITION_TIMEOUT` is the timeout duration for acquiring
a SPIFFE source bundle.

If the environment variable is not set, or cannot be parsed, defaults to
`10000` milliseconds.

> **RegEx Prefixes**
>
> You can use regular expressions for:
> * `VSECM_SPIFFEID_PREFIX_SENTINEL`,
> * `VSECM_SPIFFEID_PREFIX_SAFE`,
> * and `VSECM_SPIFFEID_PREFIX_WORKLOAD`.
> 
> When the prefix string starts with a "^", VSecM will validate the
> SPIFFE ID based on the given regular expression.
>
> For example, a pattern `"^spiffe://vsecm\.com/workload/(app1|app2)$`
> will treat `"spiffe://vsecm.com/workload/app1"` and
> `"spiffe://vsecm.com/workload/app2"` as SPIFFE IDs that can be trusted.

### VSECM_SAFE_STORE_WORKLOAD_SECRET_AS_K8S_SECRET_PREFIX

**Used By**: *VSecM Safe*.

`VSECM_SAFE_STORE_WORKLOAD_SECRET_AS_K8S_SECRET_PREFIX` retrieves the prefix that
indicates the secret shall be stored as a Kubernetes `Secret`. The prefix
is as part of the `-w` flag (*as in `-w k8s:my-secret`*).

If this environment variable is not set or is empty, it defaults to `"k8s:"`.

### VSECM_SAFE_SYNC_DELETED_SECRETS

**Used By**: *VSecM Safe*.

This configuration has no effect as of this version.
It is a work in progress.

### VSECM_SAFE_SYNC_EXPIRED_SECRETS

**Used By**: *VSecM Safe*.

This configuration has no effect as of this version.
It is a work in progress.

### VSECM_SAFE_SYNC_INTERPOLATED_K8S_SECRETS

**Used By**: *VSecM Safe*.

This configuration has no effect as of this version.
It is a work in progress.

### VSECM_SAFE_SYNC_ROOT_KEY_INTERVAL

**Used By**: "VSecM Safe".

As of this version, this configuration has no effect.
It is a work in progress.

### VSECM_SAFE_SYNC_SECRETS_INTERVAL

**Used By**: "VSecM Safe".

As of this version, this configuration has no effect.
It is a work in progress.

### VSECM_SAFE_TLS_PORT

**Used By**: *VSecM Safe*, *VSecM Sentinel*, *Workloads*.

Both **VSecM Sentinel**, **VSecM Safe**, and **workload**s use this environment
variable.

`VSECM_SAFE_TLS_PORT` is the port that **VSecM Safe** serves its API endpoints.

When you change this port, you will likely need to make changes in more
than one manifest, and restart or redeploy **VMware Secrets Manager** and
**SPIRE**.

Defaults to `":8443"`.

### VSECM_SENTINEL_INIT_COMMAND_PATH

**Used By**: *VSecM Sentinel*.

`VSECM_SENTINEL_INIT_COMMAND_PATH` returns the path to the initialization commands
file for **VSecM Sentinel**.

This file is used to automate the initialization of **VSecM Sentinel** entries
during the first deployment of **VSecM**.

If not given, defaults to `"/opt/vsecm-sentinel/init/data"`.

### VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE

**Used By**: *VSecM Sentinel*.

`VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE` is the interval to wait 
after executing an **init command** stanza of Sentinel.

The interval is an integer value representing the wait time in milliseconds.

If the environment variable is not set or cannot be parsed, it defaults to
zero milliseconds.

### VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC

**Used By**: *VSecM Sentinel*.

`VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC` is the interval to wait before 
executing an **init command** stanza of Sentinel. 

The interval is an integer value representing the wait time in milliseconds.

If the environment variable is not set or cannot be parsed, it defaults to 
zero milliseconds.

### VSECM_SENTINEL_LOGGER_URL

**Used By**: *VSecM Sentinel*.

`VSECM_SENTINEL_LOGGER_URL` ise the URL for the **VSecM Sentinel** Logger.

If this environment variable is not set, it defaults to `"localhost:50051"`.

This url used to configure gRPC logging service, which enables
**VSecM Sentinel**'s `safe` CLI command to send audit logs to the container's
standard output.

### VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL

> **Experimental**
> 
> **VSecM Sentinel** can provide an OIDC resource server so that you can
> link **VSecM** to your OIDC provider. This feature is experimental,
> have not been thoroughly tested, and is subject to change.

**Used By**: *VSecM Sentinel*.

`VSECM_SENTINEL_OIDC_PROVIDER_BASE_URL` is the base URL for the OIDC provider
server that **VSecM Sentinel** uses. This url is used when
`VSECM_SENTINEL_OIDC_ENABLE_RESOURCE_SERVER` is set to `"true"`.

Defaults to `""`.

### VSECM_SENTINEL_OIDC_ENABLE_RESOURCE_SERVER

> **Experimental**
>
> **VSecM Sentinel** can provide an OIDC resource server so that you can
> link **VSecM** to your OIDC provider. This feature is experimental,
> have not been thoroughly tested, and is subject to change.

**Used By**: *VSecM Sentinel*.

`VSECM_SENTINEL_OIDC_ENABLE_RESOURCE_SERVER` is a flag that enables the OIDC
resource server functionality in **VSecM Sentinel**.

Defaults to "false".

### VSECM_SENTINEL_SECRET_GENERATION_PREFIX

**Used By**: *VSecM Sentinel*.

`VSECM_SENTINEL_SECRET_GENERATION_PREFIX` is a prefix that's used by
that's used by **VSecM Sentinel** to generate random pattern-based secrets.

If a secret is prefixed with this value, then **VSecM Sentinel** will consider
it as a "*template*" rather than a literal value.

If the environment variable is not set or is empty, it defaults to `"gen:"`.

### VSECM_SIDECAR_ERROR_THRESHOLD

**Used By**: *VSecM Sidecar*.

**VSecM Sidecar** uses this environment variable.

`VSECM_SIDECAR_ERROR_THRESHOLD` configures the number of fetch failures before
increasing the poll interval. Defaults to `2`.

The next interval is calculated by multiplying the current interval with
`VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER`.

### VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER

**Used By**: *VSecM Sidecar*.

**VSecM Sidecar** uses this environment variable.

`VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER` configures how fast the algorithm
backs off when there is a failure. Defaults to `2`, which means when there are
enough failures to trigger a backoff, the next wait interval will be twice the
current one.

### VSECM_SIDECAR_MAX_POLL_INTERVAL

**Used By**: *VSecM Sidecar*.

**VSecM Sidecar** has an **exponential backoff** algorithm to execute fetch
in longer intervals when an error occurs. `VSECM_SIDECAR_MAX_POLL_INTERVAL`
is the maximum wait time (*in milliseconds*) before executing the next.

Defaults to `300000` milliseconds, if not provided.

### VSECM_SIDECAR_POLL_INTERVAL

**Used By**: *VSecM Sidecar*.

`VSECM_SIDECAR_POLL_INTERVAL` is the interval (*in milliseconds*)
that the sidecar polls **VSecM Safe** for new secrets.

Defaults to `20000` milliseconds, if not provided.

### VSECM_SIDECAR_SECRETS_PATH

**Used By**: *VSecM Sidecar*

`VSECM_SIDECAR_SECRETS_PATH` is path to the secrets file used by the
**VSecM Sidecar**.

If not specified, it has a default value of `"/opt/vsecm/secrets.json"`.

### VSECM_SIDECAR_SUCCESS_THRESHOLD

**Used By**:  *VSecM Sidecar*.

**VSecM Sidecar** uses this environment variable.

`VSECM_SIDECAR_SUCCESS_THRESHOLD` configures the number of successful poll
results before reducing the poll interval. Defaults to `3`.

The next interval is calculated by dividing the current interval with
`VSECM_SIDECAR_EXPONENTIAL_BACKOFF_MULTIPLIER`.

### VSECM_SPIFFEID_PREFIX_SAFE

**Used By**: *VSecM Safe*, *VSecM Sentinel*, *Workloads*.

`VSECM_SPIFFEID_PREFIX_SAFE` is the prefix for the Safe SPIFFE ID.

`VSECM_SPIFFEID_PREFIX_SAFE` is required for validation.

If the prefix starts with `"^"`, it will be treated as a regular expression.

If the prefix is a regular expression, it MUST start with 
`"^spiffe://<trust-domain>/"` where `<trust-domain>` is the trust domain
(`vsecm.com` by default).

If not provided, it will default to:
`"^spiffe://vsecm.com/workload/vsecm-safe/ns/vsecm-system/sa/vsecm-safe/n/[^/]+$"`

### VSECM_SPIFFEID_PREFIX_SENTINEL

**Used By**: *VSecM Safe*, *VSecM Sentinel*.

Both **VSecM Safe** and **VSecM Sentinel** use this environment variable.

`VSECM_SPIFFEID_PREFIX_SENTINEL` is required for validation.

If the prefix starts with `"^"`, it will be treated as a regular expression.

If the prefix is a regular expression, it MUST start with
`"^spiffe://<trust-domain>/"` where `<trust-domain>` is the trust domain
(`vsecm.com` by default).

If not provided, it will default to:
`"^spiffe://vsecm.com/workload/vsecm-sentinel/ns/vsecm-system/sa/vsecm-sentinel/n/[^/]+$"`

### VSECM_SPIFFEID_PREFIX_WORKLOAD

**Used By**: *VSecM Safe*, *VSecM Sentinel*, *Workloads*.

Both **VSecM Safe** and **workloads** use this environment variable.

`VSECM_SPIFFEID_PREFIX_WORKLOAD` is required for validation. If not provided,
it will default to: `"spiffe://vsecm.com/workload/"`

If the prefix starts with `"^"`, it will be treated as a regular expression.

If the prefix is a regular expression, it MUST start with
`"^spiffe://<trust-domain>/"` where `<trust-domain>` is the trust domain
(`vsecm.com` by default).

If not provided, it will default to:
`"^spiffe://vsecm.com/workload/[^/]+/ns/[^/]+/sa/[^/]+/n/[^/]+$"`

### VSECM_WORKLOAD_NAME_REGEXP

**Used By**: *VSecM Safe*, *VSecM Sentinel*, *Workloads*.

`VSECM_WORKLOAD_NAME_REGEXP` is the regular expression pattern for extracting
the workload name from the SPIFFE ID.

The first capture group of the regular expression is used as the workload name.

If not provided, it will default to:
`"^spiffe://vsecm.com/workload/([^/]+)/ns/[^/]+/sa/[^/]+/n/[^/]+$"`

The regular expression MUST start with `"^spiffe://<trust-domain>/"` where 
`<trust-domain>` is the trust domain (`vsecm.com` by default).

{{ edit() }}
