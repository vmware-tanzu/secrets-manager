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

title = "Production Setup"
weight = 1
+++

## Introduction

You need to pay attention to certain aspects and parts of the system that
you'd need to harden for a **production** *VMware Secrets Manager* setup.
This article will overview them.

## OpenShift and VMware Secrets Manager

Before we dive into the rest of the production setup, let's talk about
**OpenShift**. 

**OpenShift** has its own security model and its own way of doing things.
To enable **VMware Secrets Manager** to work with **OpenShift**, you need to
make sure that you have the necessary permissions and configurations in place.

The best way to do this is to use [**VSecM Helm Charts**][helm-charts] to 
deploy **VMware Secrets Manager** on **OpenShift**. The **Helm Charts** will 
take care of most of the configurations and permissions for you.

Just make sure you set up `global.enableOpenShift` to `true` in your 
`values.yaml` file. Keep in mind that this value is **false** by default, and
you need to explicitly set it to `true`.

[helm-charts]: https://artifacthub.io/packages/helm/vsecm/vsecm

## Version Compatibility

We test **VMware Secrets Manager** with the recent stable version of Kubernetes 
and Minikube.

As long as there isn't a change in the **major** version number your
Kubernetes client and server you use, things will likely work just fine.

## Resource Requirements

**VMware Secrets Manager** is designed from the ground up to work in 
environments with limited resources, such as edge computing and IoT.

That being said, **VMware Secrets Manager**, by design, is a memory-intensive 
application. However, even when you throw all your secrets at it, 
**VSecM Safe**'s peak memory consumption will be in the order or 10-20 megabytes 
of RAM. The CPU consumption will be within reasonable limits too.

However, it's crucial to understand that every system and user profile is unique.
Factors such as the number and size of secrets, concurrent processes, and system
specifications can influence these averages. Therefore, it is always advisable to
benchmark **VMware Secrets Manager** and **SPIRE** on your own system under your 
specific usage conditions to accurately gauge the resource requirements to 
ensure optimal performance.

Benchmark your system usage and set **CPU** and **Memory** limits to the
**VSecM Safe** pod.

We recommend you to:

* Set a memory **request** and **limit** for **VSecM Safe**,
* Set a CPU **request**; but **not** set a CPU limit for **VSecM Safe**
  (*i.e., the **VSecM Safe** pod will ask for a baseline CPU;
  yet burst for more upon need*).

As in any secrets management solution, your compute and memory requirements
will depend on several factors, such as:

* The number of workloads in the cluster
* The number of secrets **Safe** (*VMware Secrets Manager's Secrets Store*) has to manage
  (*see [architecture details][architecture] for more context*)
* The number of workloads interacting with **Safe**
  (*see [architecture details][architecture] for more context*)
* **Sidecar** poll frequency (*see [architecture details][architecture] for
  more context*)
* etc.

[architecture]: @/documentation/architecture/overview.md

We recommend you benchmark with a realistic production-like
cluster and allocate your resources accordingly.

That being said, here are the resource allocation reported by `kubectl top`
for a demo setup on a single-node minikube cluster to give an idea:

```txt 
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

```yaml
  # Resource allocation will highly depend on the system.
  # Benchmark your deployment, monitor your resource utilization,
  # and adjust these values accordingly.
  resources:
    requests:
      memory: "128Mi"
      cpu: "250m"
    limits:
      memory: "128Mi"
      # We recommend "NOT" setting a CPU limit.
      # As long as you have configured your CPU "requests"
      # correctly, everything would work fine.
```

## Ensure Your Clusters Up-to-Date

Although not directly related to VSecM, keeping your clusters updated is a
fundamental aspect of maintaining a secure and robust production environment.
By implementing a proactive update strategy, you not only protect your 
infrastructure from known threats but also maintain compliance with industry 
standards and regulations.

Timely updates ensure that the cluster is safeguarded against known 
vulnerabilities, which can prevent potential security leaks.

Regularly updating your cluster components ensures that you benefit from the 
latest security patches and performance improvements. These updates often 
include fixes for security flaws that, if exploited, could lead to unauthorized 
access, data breaches, or loss of service. Failing to apply updates can leave 
your cluster vulnerable to attacks that exploit outdated software 
vulnerabilities.

## Back Up Your Cluster Regularly

**VMware Secrets Manager** is designed to be resilient; however, losing access
to your sensitive data is possible by inadvertently deleting a Kubernetes
`Secret` that you are not supposed to delete. Or, your backing store that
contains the secrets can get corrupted for any reason.

Cloud Native or not, you rely on hardware which--intrinsically--is unreliable.

Things happen. Make sure you back up your cluster [using a tool like
**Velero**][velero], so that when things do happen, you can revert your
cluster's last known good state.

> **Make Sure You Back Up `vsecm-root-key`**
>
> The Kubernetes `Secret` names `vsecm-root-key` that resides in the
> `vsecm-system` namespace is especially important, and needs to be
> securely backed up.
>
> The reason is; if you lose this secret, you will lose access to all the
> encrypted secret backups, and you will not be able to restore your secrets.
{: .block-warning }

**Set up your backups from day zero**.

[velero]: https://velero.io/ "Velero"

## Double-Check Your VSecM Configuration

**VMware Secrets Manager** uses *environment variables* to configure various
aspects of its components. Although, using environment variables to configure
**VSecM** provides flexibility, it requires additional care and
attention---especially when your configuration deviates from the defaults
provided.

For example, if you have a permissive environment variable for 
`VSECM_SPIFFEID_PREFIX_WORKLOAD`, then certain apps can incorrectly identify
as workloads. Similarly, a misconfiguration of an environmnet variable or a
`ClusterSPIFFEID` may result in an unauthorized workload gaining access to 
another workload's secrets.

Treat SPIFFE IDs, and `ClusterSPIFFEID`s similar to how you treat any identitiy:
Make sure they are assigned to the apps and workloads that they ought to have
been assigned. Make sure that you are matching the correct selectors and labels
when defining `ClusterSPIFFEID`s for your workloads. And make sure that
your `*_SPIFFEID_PREFIX` environment variables are restrictive enough: This is
especially important if you use regular expression matchers in those 
environment variables.

For mor information [check out the **VSecM** configuration reference][config].

[config]: @/documentation/configuration/overview.md

## Restrict Access To `vsecm-root-key`

The `vsecm-root-key` secret that **VSecM Safe** stores in the `vsecm-system`
namespace contains the keys to encrypt and decrypt secret data on the data
volume of **VSecM Safe**.

While reading the secret alone is not enough to plant an attack on the secrets
(*because the attacker also needs to access the VSecM Safe Pod or the `/data`
volume in that Pod*), it is still **crucial** to follow the **principle of least
privilege** guideline and do not allow anyone on the cluster read or write
to the `vsecm-root-key` secret.

The only entity allowed to have read/write (*but not delete*) access to
`vsecm-root-key` should be the **VSecM Safe** Pod inside the `vsecm-system`
namespace with an `vsecm-safe` service account.

> **With Great Power Comes Great Responsibility**
>
> It is worth noting that a **Cluster Administrator** due to their elevated
> privileges can read/write to any Kubernetes `Secret` in the cluster.
>
> This includes access to the `vsecm-root-key` secret. Therefore, it is
> highly recommended that you grant the `cluster-admin` role to a **very**
> small group of trusted individuals only.
>
> Although, access to `vsecm-root-key` does not give the attacker direct
> access to the secrets, due to their sheer power, a determined Cluster
> Administrator can still read the secrets by accessing the `/data` volume.
>
> Their actions will be recorded in the audit logs, so they can, and will be
> held responsible; however, it is still a bad idea to have more than an
> absolute minimum number of Cluster Administrators in your cluster.

Kubernetes Secrets are, by default, stored **unencrypted** in the API server's
underlying data store (`etcd`). Anyone with API access and sufficient RBAC
credentials can retrieve or modify a `Secret`, as can anyone with access
to `etcd`.

> **`Secret`less VSecM**
>
> For an additional layer of security, you can opt out of using Kubernetes
> `Secret`s altogether and use **VMware Secrets Manager** without any
> Kubernetes secrets to protect the *root keys. In this mode, you'll have to
> manually provide the root keys to **VSecM Safe**; and you'll need to
> re-provide the root keys every time you restart the **VSecM Safe** Pod or
> the pod is evicted, crashed, or rescheduled.
>
> This added layer of security comes with a cost of added complexity and
> operational overhead. You will need to manually intervene when **VSemM Safe**
> crashes or restarts.
>
> That said, **VSecM Safe** is designed to be resilient, and it rarely crashes.
>
> If you let **VMware Secrets Manager** generate the root token for you, you
> will not have to worry about this, and when the system crashes, it will
> automatically unlock itself, so you can `#sleepmore`.
>
> Our honest recommendation is to let **VMware Secrets Manager** manage your
> keys unless you have special conformance or compliance requirements that
> necessitate you to do otherwise.
>
> Check ou the [Configuration Reference][config-ref] for more information.

[config-ref]: @/documentation/configuration/overview.md "Configuration Reference"

If you are **only** using **VMware Secrets Manager** for your configuration and
secret storage needs, and your workloads do **not** bind any Kubernetes `Secret`
(*i.e., instead of using Kubernetes `Secret` objects, you use tools like 
**VSecM SDK**or **VSecM Sidecar** to securely dispatch secrets to your 
workloads*) then as long as you secure access to the secret `vsecm-root-key` 
inside the `vsecm-system` namespace, you should be good to go.

With the help of **VSecM SDK**, **VSecM Sidecar**, and **VSecM Init Container**,
and with some custom coding/shaping
of your data, you should be able to use it.

However, **VMware Secrets Manager** also has the option to[persist the secrets stored in
**VSecM Safe** as Kubernetes `Secret` objects][VSecM-k]. This approach can
help support **legacy** systems where you want to start using
**VMware Secrets Manager** without introducing much code and infrastructure change to the
existing cluster--at least initially.

[VSecM-k]: /documentation/cli/#creating-kubernetes-secrets "VSecM Sentinel: Creating Kubernetes Secrets"

If you are using **VMware Secrets Manager** to generate Kubernetes `Secrets` for 
the workloads to consume, then take regular precautions around those secrets,
such as [*implementing restrictive RBACs*][rbac], and even [considering using
a KMS to encrypt `etcd` at rest][kms] if your security posture requires it.

### **Do I Really Need to Encrypt `etcd`?**

#### tl;dr:

Using plain Kubernetes `Secret`s is good enough, and it is not the
end of the world if you keep your `etcd` unencrypted.

> **VMware Secrets Manager Keeps Your Secrets Safe**
>
> If you use **VMware Secrets Manager** to store your sensitive data, your 
> secrets will be securely stored in **VSecM Safe** (*instead of `etcd`*),
> so you will have even fewer reasons to encrypt `etcd` 😉.

#### Details

This is an excellent question. And as in any profound answer to
good questions, the answer is: "*it depends*" 🙂.

`Secret`s are, by default, stored unencrypted in `etcd`. So if an adversary
can read `etcd` in any way, it's game over.

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
many, *many*, **many** concerns that you'll have to worry about.

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

secured your virtual infrastructure **and** physical data center. And
if you haven't secured your virtual and physical assets, then you are in big
trouble at day zero, even before you set up your cluster, so encrypting
`etcd` will not save you the slightest from losing other valuable data
elsewhere anyway.

##### Security Is a Layered Cake

That being said, we are humans, and $#!% does happen: If a node is compromised
due to a misconfiguration, it would be nice to make the job harder for the
attacker.

[rbac]: https://kubernetes.io/docs/reference/access-authn-authz/rbac/
[kms]: https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/

## Restrict Access to `vsecm-system` and `spire-system` Namespaces

Rigorously define and enforce access policies for the `vsecm-system` and
`spire-system` namespaces. These namespaces contain the **VSecM Safe** and
**SPIRE** components, respectively, and are critical to the security of
**VMware Secrets Manager**. Only a **Cluster Administrator** should have
access to these namespaces.

In addition, make sure you implement continuous monitoring and auditing
mechanisms to ensure that the access policies are not violated.

## Restrict Access to VSecM Sentinel

All **VMware Secrets Manager** images are based on [distroless][distroless] 
containers for an additional layer of security. Thus, an operator cannot execute 
a shell on the Pod to try a privilege escalation or container escape attack. 
However, this does not mean you can leave the `vsecm-system` namespace like an 
open buffet.

Always take a **principle of least privilege** stance. For example, do not let
anyone who does not need to fiddle with the `vsecm-system` namespace see and use
the resources there.

This stance is especially important for the **VSecM Sentinel** Pod since an
attacker with access to that pod can override (*but not read*) secrets on
workloads.

**VMware Secrets Manager** leverages Kubernetes security primitives and modern
cryptography to secure access to secrets. And **VSecM Sentinel** is the **only**
system part that has direct write access to the **VSecM Safe** secrets store.
Therefore, once you secure access to **VSecM Sentinel** with [proper RBAC and
policies][rbac], you secure access to your secrets.

> **You Can Delete `vsecm-sentinel` When You No Longer Need It**
>
> For an added layer of security and to reduce the attack surface, you can
> delete the `vsecm-sentinel` Pod after registering your secrets to 
> **VSecM Safe**.

## The Importance of Securing Access to VSecM Sentinel

Securing access to the **VSecM Sentinel** is crucial for maintaining the security
of your **VMware Secrets Manager** deployment. **VSecM Sentinel** is the primary
interface for managing secrets and interacting with the **VSecM Safe**.

The architecture of **VMware Secrets Manager** includes several key components: 

* **SPIRE** for identity control, 
* **VSecM Safe** for storing and dispatching secrets, 
* **VSecM Sidecar** for delivering secrets to workloads, 
* **VSecM Sentinel** for administrative tasks and registering secrets. 
* The **SPIFFE ID** format is used for workload identification. 

The secrets stored in **VSecM Safe** are encrypted using `age` or `AES-256-GCM`
(*in FIPS-compliant environments*) and stored in memory for performance
and security, necessitating sufficient memory allocation for all secrets.

In this architecture any pod in the `vsecm-system` namespace that has the
`vsecm-sentinel` service account can access **VSecM Safe**.

Securing access to the `vsecm-system` namespace is important because if an 
attacker gains access to this namespace, they could potentially introduce a fake 
*VSecM Safe*. 

Although the default configuration of **VMware Secrets Manager** is secure,
and mitigates many common attack vectors, it is essential to follow best 
practices to further enhance security.

To secure your system even further, follow these best practices:

* **Restrict Namespace Access**: Tighten the security around the 
  **vsecm-system** namespace to prevent unauthorized access. This could involve 
  network policies, stricter RBAC rules, and audit logging to monitor for 
  unusual activities.
* **Least Privilege Principle**: Ensuring that only the minimal necessary 
  permissions are granted to users and service accounts, reducing the attack 
  surface.
* **Regular Audits and Anomaly Detection**: Conduct regular audits of your 
  Kubernetes environment and implement anomaly detection mechanisms to identify 
  and respond to unauthorized access or changes in the **vsecm-system** 
  namespace.
* **Security Scanning and Compliance Tools**: Integrate security scanning and 
  compliance tools to continuously monitor for vulnerabilities and enforce 
  security best practices.
* **Immutable Infrastructure**: Use immutable infrastructure principles for 
  critical components, where any change to the running configuration would be 
  a red flag indicating potential tampering.
* **Open Policy Agent (OPA)/Gatekeeper**: Implement policies using 
  OPA/Gatekeeper to enforce custom rules, like restricting which registries 
  pods can pull images from or enforcing labeling standards.

By combining these measures, enhance the security of you **VSecM** installation
even further.

Also, note that, these practices are Kubernetes security best practices.
They are not only applicable to **VSecM Sentinel**, but also to any other 
critical components in your system.

## Scaling SPIRE

**SPIRE** is designed to scale horizontally. This means that you can add more
**SPIRE Server** and **SPIRE Agent** instances to your cluster to increase
the capacity of your **SPIRE** deployment.

Although **VMware Secrets Manager** comes with a default SPIRE configuration,
depending on your deployment needs, you may need to scale SPIRE to meet your
specific requirements.

SPIRE supports:

* Horizontal scaling with multiple SPIRE server,
* Nested topologies to have separate failure domains,
* Federated deployments with multiple trust roots,
* And more.

Check out [Scaling SPIRE][scaling-spire] and [Extending SPIRE][extending-spire]
sections in the official SPIRE documentation for more information.

[scaling-spire]: https://spiffe.io/docs/latest/planning/scaling_spire/ "Scaling SPIRE"
[extending-spire]: https://spiffe.io/docs/latest/planning/extending/ "Extending SPIRE"

## Use of Attestation

In **VSecM**, the security of your secrets depends on the `ClusterSPIFFEID`s
that you assign to your workloads. Therefore, it is crucial to ensure that
you specify proper attestors in your `ClusterSPIFFEID`s.

You can check the [SPIRE Documentation][spire-docs] and
also [VSecM Usage Examples][usage-examples] for examples on how to create
`ClusterSPIFFEID`s with proper attestors.

## SPIRE Configuration

**VMware Secrets Manager** uses [SPIRE][spire] as its underlying identity control
plane. The default SPIRE configuration bundled with **VMware Secrets Manager**
is secure enough for most use cases.

While **VSecM** uses sane defaults for SPIRE installation, it can be further
hardened according to specific deployment needs, providing a more robust and
secure environment.

Here are some suggestions to consider; as always, you should consult
the [SPIRE documentation][spire-docs] for more details.

[usage-examples]: https://github.com/vmware-tanzu/secrets-manager/tree/main/examples "VSecM Usage Examples"

### Enabling Kubelet Verification

For ease of installation the **SPIRE Agent** is configured to trust all kubelets
by setting `skip_kubelet_verification` to `true` in the `agent.conf` file.

The `skip_kubelet_verification` flag is used when **SPIRE** is validating the
identity of workloads running in Kubernetes.

Normally, **SPIRE** interacts with the kubelet API to verify the identity of a 
workload. This includes validating the serving certificate of the kubelet.
When `skip_kubelet_verification` is set to `true`, **SPIRE** does not validate
the kubelet's serving certificate. This can be useful in environments where the
kubelet's serving certificate is not properly configured or cannot be trusted for
some reason.

That being said, skipping kubelet verification reduces security. It should be used
cautiously and only in environments where the risks are understood and deemed
acceptable.

To ensure kubelet verification is enabled:

* **Flag Setup**: Ensure the `skip_kubelet_verification` flag is either set to
  `false` or omitted.
  By default, if the flag is not specified, kubelet verification is enabled.
* **Kubelet Certificate**: Make sure the kubelet's serving certificate is properly
  configured and trusted within your Kubernetes cluster. This may involve 
  configuring the Kubernetes cluster to issue valid serving certificates for 
  kubelets.
* **Restart SPIRE Agent**: After making changes to the configuration, restart
  the **SPIRE Agent** to apply the new settings.

> **Plan Carefully**
>
> Remember, enabling kubelet verification might require updates to your cluster's
> configuration and careful planning to avoid disruption to existing workloads.

### Configuration Files

**SPIRE Server** and **SPIRE Agent** are configured using `server.conf` and
`agent.conf` files, respectively. For Kubernetes deployments, these can be stored
in `ConfigMaps` and mounted into containers. This ensures configuration consistency
and ease of updating.

To secure these configuration files, you can:

* Use a `Secret` instead of a `ConfigMap` to store the configuration files.
* Encrypt `etcd` at rest using a **KMS** (*which is the most robust method
  proposed [in the Kubernetes guides][kms]*).
* **Audit Logs**: Enable and monitor audit logs to track access and changes to `ConfigMaps`.
  This helps in identifying unauthorized access or modifications.
  * Audit logs are also useful to track who executed commands like `kubectl exec $SENTINEL ...`
    to ensure that only authorized personnel are accessing the **VSecM Sentinel** to 
    manage secrets.
* **Regular Reviews and Updates**: Periodically review and update the access policies
  and configurations to ensure they remain secure and relevant.
* **Minimize Configuration**: Only include necessary configuration settings in
  `server.conf` and `agent.conf`. Avoid any sensitive data unless absolutely
  necessary.

[spire]: https://spiffe.io/spire/
[spire-docs]: https://spiffe.io/docs/latest/

### Trust Domain Configuration

Set the `trust_domain` parameter in both server and agent `ConfigMaps`. This 
parameter is crucial for ensuring that all workloads in the trust domain are 
issued identity documents that can be verified against the trust domain's root 
keys.

### Port Configuration

The `bind_port` parameter in the server `ConfigMap` sets the port on which the
**SPIRE Server** listens for **SPIRE Agent** connections. Ensure this port is
securely configured and matches the setting on the agents.

### Node Attestation

Choose and configure appropriate Node Attestor plugins for both **SPIRE Server** 
and**SPIRE Agent**. This is critical for securely identifying and attesting 
agents.

### Data Storage

For SPIRE runtime data, set the `data_dir` in both server and agent `ConfigMap`s.
Use **absolute paths** in production for stability and security.

Consider the choice of database for storing SPIRE Server data, especially in
high-availability configurations.

By default, SPIRE uses SQLite, but for production, an alternative SQL-compatible
storage like MySQL can be a better fit.

### Key Management

**SPIRE** supports *in-memory* and *on-disk* storage strategies for keys and 
certificates.

For production, the **on-disk** strategy may offer advantages in terms of 
persistence across restarts but requires additional security measures to 
protect the stored keys.

### Trust Root/Upstream Authority Configuration

Configure the `UpstreamAuthority` section in the server `ConfigMap`.

This is pivotal for maintaining the integrity of the SPIRE Server's root signing
key, which is central to establishing trust and generating identities.

### SPIRE Needs hostPath Access for SPIRE Agent DaemonSets

**SPIRE Agent** primarily uses `hostPath` for managing [Unix domain
socket][unix-domain-socket]s on Linux systems. This specific usage is focused
on facilitating communication between the **SPIRE Agent** and **workloads**
running on the same host.

The **Unix domain socket** used by the **SPIRE Agent** is typically configured
to be read-only for workloads. This read-only configuration is an important
security feature for several reasons:

* **Principle of Least Privilege**: Setting the Unix domain socket to read-only
  for workloads adheres to the principle of least privilege. Workloads generally
  only need to read data from the socket (*such as fetching SVIDs*) and do not
  require write permissions. Limiting these permissions reduces the risk of
  unauthorized actions.
* **Mitigating Risks of Tampering**: By making the socket **read-only** for
  workloads, the risk of these workloads tampering with the socket's data or
  behavior is minimized. This is crucial as SPIRE Agents deal with sensitive
  identity credentials.
* **Reducing Attack Surface**: A **read-only** configuration limits the potential
  actions an attacker can perform if they compromise a workload.
* **Ensuring Data Integrity**: **Read-only** access helps ensure the integrity of
  the data being transmitted through the socket. Workloads receive the data as
  intended by the SPIRE Agent without the risk of accidental or malicious alteration.
* **Compliance with Security Best Practices**: This configuration aligns with
  broader security best practices in systems design, where components are given only
  the permissions necessary for their function, reducing potential vulnerabilities.

> **OpenShift Support**
>
> For Kubernetes deployments such as [OpenShift][openshift] where enabling `hostPath`
> requires additional permissions 
> [you can follow SPIRE's official documentation][spire-openshift].


[spire-openshift]: https://spiffe.io/docs/latest/deploying/spire_agent/#openshift-support
[openshift]: https://www.openshift.com/ "OpenShift"

To make the `hostPath` binding extra secure, you can:

* Use **Pod Security Admission** and custom **Admission Controllers** to restrict
  the use of `hostPath` to certain paths and ensure that only authorized pods have
  access to those paths.
* **Node-Level Security**: Ensure that the nodes themselves are secure. This includes
  regular updates, patch management, and following best practices for host security.
  Secure nodes reduce the risk of compromising the directories accessed through `hostPath`.
* **Network Policies**: Configure network policies to control the traffic to and from
  the SPIRE Agent pods. This can limit the potential for network-based attacks
  against the agents.
* **Regular Security Reviews**: Regularly review and update your security configurations.
  This includes checking for updates in Kubernetes security recommendations and ensuring
  your configurations align with the latest best practices.

[unix-domain-socket]: https://en.wikipedia.org/wiki/Unix_domain_socket

[distroless]: https://github.com/GoogleContainerTools/distroless

### Further Security and Configuration Considerations

As in any distributed system, regularly monitor and audit **SPIRE** and **VSecM**
operations to detect any unusual or suspicious activity. This includes monitoring
the issuance and use of `SVID`s, as well as the performance and status of the
**SPIRE Server** and **SPIRE Agent*(s.

Regularly conduct security audits of your **SPIRE** deployment to identify and
address any vulnerabilities.

To reduce the blast radius in unlikely breaches, if needed, use a **nested
topology** and **federated deployments** to segment failure domains and provide
multiple roots of trust.

## Keep the SPIRE Server Alive

This is more of a **stability** than a **security** concern; however, it is
important.

If **SPIRE Server** if offline for a long time then its root certificate will
expire. The expiry time of the root certificate is configurable, but by default
it's CA TTL is *24 hours*.

> **SPIRE Is Designed to Be Resilient**
>
> Occasional disruptions, evictions, and restarts of SPIRE Server are not
> a problem. **SPIRE Server** is designed to be resilient and it will
> automatically recover from such disruptions.
>
> However, if the **SPIRE Server** is offline for more than its TTL, then
> it will not be able to renew its root certificate, and this will disrupt
> the trust mechanism within the **SPIRE** environment.
{: .block-info }

Regarding the implications of the **SPIRE Server** being offline for more than
its TTL, it's important to understand the role of the server's CA certificate in
the SPIRE architecture:

The CA certificate is central to the trust establishment in the **SPIRE**
infrastructure. If the server is offline and unable to renew its CA certificate
before expiration, this will disrupt the trust mechanism within the SPIRE
environment. Agents and workloads will not be able to validate the authenticity
of new SVIDs issued after the CA certificate has expired, leading to trust and
authentication issues across the system.

> **Know Your TTLs**
>
> Although **SPIRE** has sane defaults, it is still important to know your
> tolerance and set TTLs (*both CA TTLs, and agent SVID TTLs*) accordingly.
{: .block-warning }

From the **VMware Secrets Manager** perspective, this will result in workloads
not being able to receive secrets from the **VSecM Safe**, and the **VSecM Safe**
failing to respond to the requests made by the **VSecM Sentinel**.

Therefore, it's important to ensure that the **SPIRE Server** is online and able
to renew its CA certificate before it expires. Otherwise, manual intervention
will be required to fix the trust issue.

## Volume Selection for VSecM Safe Backing Store

[**VSecM Safe** default deployment descriptor][vsecm-safe-deployment-yaml]
uses [`HostPath`][k8s-pv] to store encrypted backups for secrets.

It is highly recommended to ensure that the backing store **VSecM Safe**
uses is **durable**, **performant**, and **reliable**.

It is a best practice to avoid `HostPath` volumes for production deployments.
You are strongly encouraged to [choose a `PersistentVolume` that suits your
needs][k8s-pv] for production setups.

If you are using Helm Charts, you can modify the `values.yaml` file to use
a `PersistentVolume` instead of `HostPath` for **VSecm Safe**:

```yaml
# ./charts/safe/values.yaml

# -- How persistence is handled.
data:
  # -- If `persistent` is true, a PersistentVolumeClaim is used.
  # Otherwise, a hostPath is used.
  persistent: false
  # -- PVC settings (if `persistent` is true).
  persistentVolumeClaim:
    storageClass: ""
    accessMode: ReadWriteOnce
    size: 1Gi
```

[vsecm-safe-deployment-yaml]: https://github.com/vmware-tanzu/secrets-manager/blob/main/helm-charts/charts/safe/templates/Deployment.yaml
[k8s-pv]: https://kubernetes.io/docs/concepts/storage/volumes/

## Volume Selection for SPIRE Server Data

Similar to **VSecM Safe**, it is also recommended for **SPIRE Server** to use
a `PersistentVolume` for storing its data for production deployments.

With Helm charts, you can modify the `values.yaml` file to use a `PersistentVolume`
instead of in-memory storage for **SPIRE Server**:

```yaml
# ./charts/spire/values.yaml

# -- Persistence settings for the SPIRE Server.
data:
  # -- Persistence is disabled by default. You are recommended to provide a
  # persistent volume.
  persistent: false
  # -- Define the PVC if `persistent` is true.
  persistentVolumeClaim:
    storageClass: ""
    accessMode: ReadWriteOnce
    size: 1Gi
```

## High Availability of VSecM Safe

> **tl;dr:**
>
> **VSecM Safe** may not emphasize high-availability, but its robustness is
> so outstanding that the need for high-availability becomes almost negligible.


Since **VSecM Safe** keeps all of it state in memory, using a pod with enough
memory and compute resources is the most effective way to leverage it. Although,
with some effort, it might be possible to make it highly available, the effort
will likely bring unnecessary complexity without much added benefit.

**VSecM Safe** is, by design, a single pod; so technically-speaking, it is
not highly-available. So in the rare case when **VSecM Safe** crashes, or
gets evicted due to a resource contention, there will be minimal disruption
until it restarts. However, **VSecM Safe** restarts fairly quickly, so the
time window where it is unreachable will hardly be an issue.

Moreover **VSecM Safe** employs "*lazy learning*" and does not load everything
into memory all at once, allowing **very** fast restarts. In addition, its
lightweight and focused code ensures that crashes are infrequent, making
**VSecM Safe** *practically* highly available.

While it is possible to modify the current architecture to include more than one
**VSecM Safe** pod and place it behind a service to ensure high-availability,
this would be a significant undertaking, with not much benefit to merit it:

First of all, for that case to happen, the state would need to be moved away
from the memory, and centralized into a common in-memory store (*such as Redis,
or etcd*). This will introduce another moving part to manage. Or alternatively
all **VSecM Safe** pods could be set up to broadcast their operations and reach
a quorum. A quorum-based solution would be more complex than using a share store,
besides reaching a quorum means a performance it (*both in terms of decision time
and also compute required*).

On top of all these bootstrapping coordination would be necessary to prevent
two pods from creating different bootstrap secrets simultaneously.

Also, for a backing store like Redis, the data would need to be encrypted
(*and Redis, for example, does not support encryption at rest by default*).

When considering all these, **VSecM Safe** has not been created highly-available
**by design**; however, it is so robust, and it restarts from crashes so fast that
it's "*as good as*" highly-available.

## DO NOT LIMIT CPU on VSecM Pods

**VSecM Safe** uses CPU resources only when it needs it. It is designed to be
lightweight and it does not consume CPU resources unless it needs to. So
unless you have a very specific reason to limit CPU on **VSecM Safe** pods,
it is recommended to let it burst when it needs.

Moreover, **VSecM Safe** is a go-based application. Limiting CPU on Go-based
workloads can be problematic due to the nature of [Go's garbage collector
(*GC*) and concurrency management][go-gc].


In Go, a significant portion of CPU usage can be attributed to the garbage collector
(*GC*). It's designed to be fast and optimized, so altering its behavior is generally
**not** recommended.

Limiting CPU directly for Go-based workloads might not be the best approach due
to the intricacies of Go's garbage collection and concurrency model. And
**VSecM Safe** is no exception to this.

Instead, profiling it to understand its specific needs in your cluster and
adjusting the relevant environment variables (like `GOGC` and `GOMAXPROCS`) can
lead to better overall performance.

Having said that, please note that each cluster has its own characteristics
and this is not a one-size-fits-all recommendation. Kubernetes is a complex
machine and there are many factors that can influence the performance of
**VSecM Safe** including, but not limited to Node Capacity, Node Utilization,
CPU Throttling and Overcommitment, QoS Classes, and so on.

[go-gc]: https://tip.golang.org/doc/gc-guide.html

## Update VMware Secrets Manager's Log Levels

**VSecM Safe** and **VSecM Sidecar** are configured to log at `TRACE` level by
default. This is to help you debug issues with **VMware Secrets Manager**. 
However, this can cause a lot of noise in your logs. Once you are confident 
that **VMware Secrets Manager** works as expected, you can reduce the log level 
to `INFO` or `WARN`.

For this, you will need to modify the `VSECM_LOG_LEVEL` environment variable
in the **VSecM Safe** and **VSecM Sidecar** Deployment manifests.

See [**Configuring VMware Secrets Manager**][config] for details.

[config]: @/documentation/configuration/overview.md

## Update SPIRE's Log Levels

The default VSeCM **SPIRE** installation is configured to log at `DEBUG` level.
This might be too verbose for a production environment. You can reduce the log
level to `INFO` or `WARN` to reduce the noise in your logs.

For this, you will need to modify the update the `log_level` parameter in the
`server.conf` and `agent.conf` ConfigMaps.

Here is a relevant agent.conf snippet from **VSecM** Helm charts:

```yaml
# ./charts/spire/templates/spire-agent-configmap.yaml

data:
  agent.conf: |
    agent {
      data_dir = "/run/spire"
      log_level = {{ .Values.global.spire.logLevel | quote }}
      server_address = {{ .Values.global.spire.serverAddress | quote }}
      server_port = {{ .Values.global.spire.serverPort | quote }}
      socket_path = "/run/spire/sockets/spire-agent.sock"
      trust_bundle_path = "/run/spire/bundle/bundle.crt"
      trust_domain = {{ .Values.global.spire.trustDomain | quote }}
    }
```

If you are using Helm charts to install **VSecM**, you can also provide 
`global.spire.logLevel` in your `values.yaml` file to override the default 
SPIRE log level.

## A Note on FIPS Compliance with VSecM

VMware Secrets Manager (VSecM) is designed with security and compliance in mind,
providing a robust platform for managing secrets and sensitive information in
a FIPS-compliant manner. Here are key features and practices already implemented
in VSecM that contribute to its FIPS compliance:

### FIPS-Compliant Cryptographic Modules 

VSecM utilizes cryptographic modules that are either FIPS-certified or comply 
with FIPS standards for encryption algorithms and key management practices. 

This ensures that cryptographic operations within **VSecM** adhere to the 
rigorous requirements set forth by FIPS standards.

### Encryption at Rest and In Transit

**VSecM** ensures that all secrets and sensitive data are encrypted both at 
rest and in transit using FIPS-approved algorithms. This safeguards the data 
against unauthorized access and exposure during storage and transmission.

### Secure Key Management

VSecM implements secure key management practices, including the secure 
generation, storage, and handling of cryptographic keys. This minimizes the risk 
of key compromise and ensures the integrity of cryptographic operations.

### Role-Based Access Control (RBAC)

**VSecM** employs strict Kubernetes RBAC policies to control access to secrets 
and cryptographic keys. This allows for fine-grained access control, ensuring 
that only authorized personnel can access sensitive information based on their 
roles.

### Ongoing Enhancements for Improved Compliance

To further enhance its compliance and security capabilities, 
we are actively working on the following initiatives:

#### Hardware Security Module (HSM) Integration

Integration with HSMs is underway to provide an additional layer of security 
for cryptographic key management. HSMs offer hardware-based key storage and 
cryptographic operations, providing superior protection against key compromise 
and enhancing the overall security of cryptographic practices within VSecM.

#### Cloud Key Management Service (KMS) Integration

We are also working on integrating with cloud-based KMS solutions. This allows 
for the centralized management of cryptographic keys in the cloud, offering 
scalability, high availability, and the convenience of cloud-based key 
management, while still adhering to FIPS standards.

### Recommendations for End-Users to Enhance Practical Compliance

End-users of **VSecM** can take additional steps to improve their practical 
compliance with FIPS standards, including:

* *Encrypt Kubernetes Secrets*: For applications deployed in Kubernetes 
  environments, ensure that Kubernetes Secrets used for storing root keys and 
  other sensitive information are encrypted at rest using a KMS provider that is 
  FIPS-compliant.
* *Implement Audit Logging and Monitoring*: Establish comprehensive audit 
  logging and monitoring for access to secrets and cryptographic operations. 
  This helps in identifying unauthorized access attempts and ensuring compliance 
  with security policies.
* *Regular Security Assessments*: Conduct regular security assessments of your 
  applications and infrastructure to identify and address potential 
  vulnerabilities. This includes penetration testing and vulnerability scanning.
* *Disaster Recovery and Key Rotation*: Develop and maintain disaster recovery 
  plans and implement key rotation policies to minimize risks associated with 
  key compromise. Regular key rotation and having a robust disaster recovery 
  plan in place are critical for maintaining a secure and resilient cryptographic 
  infrastructure.
* *Document Compliance Efforts*: Maintain detailed documentation of your security 
  controls, policies, and procedures, including how you use VSecM and other tools 
  in a manner that aligns with FIPS standards. This documentation is invaluable 
  for internal review and compliance audits.

By leveraging the FIPS-compliant features of VSecM and adopting these recommended 
practices, you can significantly enhance the security and compliance of their 
secret and key management practices, ensuring the protection of sensitive 
information and adherence to FIPS standards.

Please note that FIPS compliance is a shared responsibility between VSecM and
the end-user. While VSecM provides a secure platform for managing secrets and
sensitive data, you are responsible for configuring and using VSecM in a
manner that aligns with FIPS standards and best practices.

In addition, **VSecM** helm charts do **not** enable FIPS mode by default. If
you need to enable FIPS mode, you can do so by setting the 
`VSECM_SAFE_FIPS_COMPLIANT` environment variable to `true` in the `environments`
section of the **VSecM Safe** helm chart.

## Conclusion

Since **VMware Secrets Manager** is a *Kubernetes-native* framework, its 
security is strongly related to how you secure your cluster. You should be safe 
if you keep your cluster and the `vsecm-system` namespace secure and follow
"*the principle of least privilege*" as a guideline.

**VMware Secrets Manager** is a lightweight secrets manager; however, that does 
not mean it runs on water: It needs CPU and Memory resources. The amount of 
resources you need will depend on the criteria outlined in the former sections. 
You can either benchmark your system and set your resources accordingly. Or set 
generous-enough limits and adjust your settings as time goes by.

Also, you are strongly encouraged **not** to set a limit on **VMware Secrets 
Manager** Pods' CPU usage. Instead, it is recommended to let **VSecM Safe** 
burst the CPU when it needs.

On the same topic, you are encouraged to set a **request** for **VSecM Safe**
to guarantee a baseline compute allocation.

{{ edit() }}
