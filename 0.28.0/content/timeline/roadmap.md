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

title = "VSecM Roadmap"
weight = 10
+++

## Introduction

This is a page where we publish our approximate roadmap for **VMware Secrets
Manager** for Cloud-Native Apps. Note that this is not a commitment to deliver
any of the features listed here, and that the roadmap is subject to change at
any time without notice.

Whenever we release a new version of **VMware Secrets Manager**, we will update
this page, and also [the changelog](@/timeline/changelog.md) to reflect the changes.

> **One-Year Window**
>
> This page will only contain information about the next 12 months of the
> project. We will update the roadmap every release, and remove the completed
> items from the list, and add a new iteration at the end of the list.

## Active Iterations

### VSecM v0.29.0 (*codename: Indus*)

**Oct 16, 2024 - Jan, 31, 2025**

This iteration focuses on several enhancements and fixes aimed at improving 
flexibility and security. Key areas include:

* Introducing secret kind ‘raw’ for external operators.
* Enhancements to VSecM Sentinel with nested encryption.
* Removing autoscaling from Helm charts.
* Configuring sleep intervals for better resource control.
* Improving PostgreSQL table configurations via environment variables.
* Utilizing Web Crypto API for secure secret relay.
* Additionally, we are addressing good first issues like increasing test 
  coverage and optimizing the landing page CSS.

[Here is a list of issues that are candidate for VSecM vIndus](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.29.0-candidate+).

### VSecM v0.30.0 (*codename: Lupus*)

**Feb 01, 2025 - Feb 28, 2025**

This iteration focuses on several key features and enhancements:

* Fetching data in the GoLang SDK using a `jq` filter.
* Better support for secret versioning.
* A new hub-spoke secrets relay architecture.
* Adding multi-tenancy capabilities.
* Integration tests and in-memory persistence improvements.
* New video tutorials and updated documentation for clarity.

The goal is to solidify multi-tenancy, streamline secret handling, and 
improve system stability through better testing and documentation.

[Here is a list of issues that are candidate for VSecM 
vLupus](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.30.0-candidate+).

### VSecM v0.31.0 (*codename: Mensa*)

**Mar 01, 2025 - Apr 11, 2025**

This iteration focuses on enhancing security and user-requested features:

* Audit logging and log streaming.
* Rate limiting to control access.
* Hierarchical secrets management and better root key support from external users.
* A `/purge` API to remove orphaned secret backups.
* Features like secret versioning, ACL for secret access, and storing large 
  Kubeconfig files.
* Several CLI enhancements for improved usability.
* Enhanced support for Java SDK, alongside multi-tenancy and operator-specific 
  encryption capabilities.

[Here is a list of issues that are candidate for VSecM 
vMensa](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.31.0-candidate+).

### VSecM v0.32.0 (*codename: Norma*)

**Apr 12, 2025 - May 09, 2025**

This iteration brings major enhancements for system security and SDK expansion:

* Development of a UI for VSecM Sentinel leveraging OIDC functionality.
* Introduction of user auditing capabilities in Sentinel.
* New SDKs in Rust and Python.
* Enhancements for SPIRE Helm charts, supporting x509 node attestation.
* Optional support for AWS KMS for master key storage.
* Consideration of Redis as a memory backing store for VSecM Safe.

These improvements focus on expanding SDKs, improving audit capabilities, and 
integrating new storage options.

[Here is a list of issues that are candidate for VSecM vNorma](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.32.0-candidate+).

### VSecM v0.33.0 (*codename: Orion*)

**May 10, 2025 - Jun 06, 2025**

This iteration focuses on improving security, scalability, and coverage:

* ClusterSPIFFEID management added to Sentinel.
* RBAC/ABAC policy support for enhanced access control.
* Scalability improvements with multiple VSecM Safe instances.
* 60% test coverage target across the project.
* New hashing of log lines to prevent tampering and increase security.
* Documentation on federating identity control planes for VSecM.

These enhancements aim to solidify security, scalability, and manageability 
while enhancing project test coverage.

[Here is a list of issues that are candidate for VSecM vOrion](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.33.0-candidate+).

### VSecM v0.34.0 (*codename: Perseus*)

**Jun 07, 2025 - Jul 04, 2025**

This iteration emphasizes enhancements around Helm charts and key management for SPIRE:

* Support for key rotation.
* Expanded Helm chart capabilities, including customizing node attestors, 
  key managers, and data stores for both SPIRE server and agent.
* Configurable options for telemetry and federation in SPIRE.
* Key storage on persistent volumes (PVs) and custom upstream authorities.
* Focus on integrating SSH node attestation and improving system resilience via 
  retry logic.
* These changes aim to enhance flexibility, security, and scalability in managing 
  SPIRE deployments.

[Here is a list of issues that are candidate for VSecM 
vPerseus](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.34.0-candidate+).

### VSecM v0.35.0 (*codename: Reticulum*)

**Jul 05, 2025 - Aug 01, 2025**

This release emphasizes improved flexibility and synchronization for key and 
secret management:

* Support for VSecM Safe as an alternative to the in-memory store, with two-way 
  sync.
* Integration with cloud KMS and databases for secret backups, root key storage, 
  and versioning.
* New `/stats` and `/health` endpoints for VSecM Safe.
* Persistent root key storage across cloud KMS and persistent volumes, with 
  automatic updates to VSecM memory.

These features enhance resilience and scalability across various storage backends.

[Here is a list of issues that are candidate for VSecM 
vReticulum](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.35.0-candidate+).

### VSecM v0.36.0 (*codename: Sagittarius*)

**Aug 02, 2025 - Aug 29, 2025**

This release continues to enhance security and storage capabilities:

* Use a separate VSecM Safe to store root keys, improving security by avoiding 
  reliance on Kubernetes secrets.
* Focus on workflows and improving overall security measures within VSecM 
  infrastructure.

These updates aim to bolster the system’s integrity by leveraging dedicated, 
secure storage solutions for critical keys.

[Here is a list of issues that are candidate for VSecM 
vSagittarius](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.36.0-candidate+).

### VSecM v0.37.0 (*codename: Telescopium*)

**Aug 30, 2025 - Nov 03 2025**

This iteration focuses on demonstrating key features through extensive video 
tutorials:

* Demos on key management, secret decryption, root key changes, and large file 
  encryption.
* Showcasing integrations with tools like Keycloak and Cassandra.
* Use cases for federated SPIRE, three-way federation, and GitOps.
* Secrets handling across multiple VSecM instances, namespaces, and workloads.
* Deploying VSecM on Kubernetes clusters and EKS.

* These video demonstrations enhance understanding of VSecM’s advanced features 
  and integrations.

[Here is a list of issues that are candidate for VSecM 
vTelescopium](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.37.0-candidate+).

### VSecM v0.38.0 (*codename: Ursa*)

**Nov 04, 2025 - Dec 01 2025**

This release focuses on improving automated testing and security demos:

* Achieve 90% test coverage using FLOSS automated test suites.
* A demo showcasing the integration of OPA (Open Policy Agent) with VSecM.
* Further enhancements to the workflow and project infrastructure.

* These updates are crucial for strengthening the project's testing capabilities 
* and demonstrating VSecM's integration with modern policy management tools like OPA.

[Here is a list of issues that are candidate for VSecM 
vUrsa](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.38.0-candidate+).

### VSecM v0.39.0 (*codename: Virgo*)

**Dec 02, 2025 - Dec 29 2025**

This release focuses on enhancing security, replication, and integration:

* `ValidatingAdmissionWebhook` to ensure proper ClusterSPIFFEID templates.
* Secrets rotation demo with a sidecar.
* Replication support for multiple VSecM Safe instances.
* Improved audit logging with separation options.
* Kubernetes Operator for automating VSecM sidecar and init container injection.

These enhancements aim to improve multi-cloud integration, security, and 
cluster management capabilities.

[Here is a list of issues that are candidate for VSecM 
vVirgo](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.39.0-candidate+).

### VSecM v0.40.0 (*codename: Antlia*)

**Dec 30, 2025 - Jan 26 2026**

This release enhances key integration and testing functionalities:

* Expanded secrets rotation demo with sidecar integration.
* Sentinel OIDC Resource Server functionality included in integration tests.
* Documentation for VSecM Sentinel OIDC authentication.
* Replication support for multiple VSecM Safe instances.
* Exploration of minimal disruption secret refresh strategies.
* Customizable kubelet verification in Helm charts.
* Self-security assessment and independent security audit.

This release focuses on improving authentication, replication, and security 
testing.

[Here is a list of issues that are candidate for VSecM 
vAntlia](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.40.0-candidate+).

### VSecM v0.41.0 (*codename: Bellatrix*)

**Jan 27, 2026 - Feb 23 2026**

This is a "*catch all*" that contains all remaining documented future plans.
We will create new iterations from it as the time gets closer.

[Here is a list of issues that are candidate for VSecM 
vBellatrix](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.41.0-candidate+).

## Closed Iterations

### VSecM v0.28.0 (*codename: Hydra*)

**Dec 30, 2025 - Jan 26, 2026**

This iteration was about increasing coverage. We will focus on unit tests.

In addition, we are targeted to fix certain low-hanging bugs and improve
stability.

[Here is a list of issues that were candidate for VSecM vHydra
](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.28.0-candidate+).

### VSecM v0.27.0 (*codename: Gemini*)

**May 23, 2024 - Jun 19, 2024**

The sole focus of this iteration was increasing unit test coverage and adding
more integration tests. 

We also introduced improvements too; however, stability will be our main
focus.

[Here is a list of issues that are candidate for VSecM vGemini
](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.27.0-candidate+).

### VSecM v0.26.1 (*codename: Fornax*)

**Apr 25, 2024 - May 22, 2024**

This iteration will was about stability and documentation updates.

We also introduced a lot of flexibility such as ability to use custom 
namespaces, trust domains, and regex-based SPIFFEID validation.

[Here is a list of issues that are candidate for VSecM vFornax
](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.26.1-candidate+).

### VSecM v0.25.0 (*codename: Eridanus*)

**Mar 28, 2023 - Apr 24, 2024**

This iteration was mostly about security and stability.

[Here is a list of issues that were closed in vEridanus](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.25.0-candidate+).

### VSecM v0.24.0 (*codename: Draco*)

**Feb 29, 2023 - Mar 27, 2023**

To automate things and be able to dynamically follow issues better, from
this point on we started labeling them and share the GitHub filter here.

This iteration was mainly focused on demos and documentation.

[Here is a list of issues that were closed in  vDraco](https://github.com/vmware-tanzu/secrets-manager/issues?q=is%3Aissue+label%3Av0.24.0-candidate+).

### VSecM v0.23.0 (*codename: Cassiopeia*)

**Feb 01, 2024 - Feb 28, 2024**

This iteration was focused on improving how **VMware Secrets Manager**
logs and reports errors. We will also focus on improving the performance of the
**VMware Secrets Manager** website.

* `Secret`less VSecM: Ability to use VMware Secrets Manager **without** relying
  on Kubernetes `Secret`s. This will allow users to use **VMware Secrets Manager**
  without having to create Kubernetes `Secret`s at all--even for the root keys.
* Ability to use VSecM across clusters (*multi-cluster federation support*).
* More automation, and stability improvements.
* Ability to use an "*init command*" for **VSecM Sentinel** to run before the
  world starts.
* Ability to generate pattern-based random secrets.
* The operator shall be able to export secrets in an encrypted format, and
  can decrypt them, if they have the right permissions.
* A public ECR registry to share the untested "*edge*" versions of **VMware
  Secrets Manager**, for those who like living dangerously.
* Focus on increasing test coverage.
* Ability to create Kubernetes `Secret`s without necessarily associating them
  with a workload.
* Adding "invalid before" and "expires after" timestamps to secrets, to
  help with secret rotation.
* Progress towards Open SSF Best Practices compliance; reaching 97% of the  
  requirements.

### VSecM v0.22.0 (*codename: Boötes*)

**Sep 12, 2023 - Jan 31, 2024**

This was a relatively longer release because due to the "*time-stop*" effect of
the holiday season, the majority of the core contributors will be spending quality
time with their loved ones and recharging their batteries for the upcoming year.

This release will be more about enhancing deployment workflows, testing automation
and CI/CD pipelines. We will also focus on improving the overall user experience.

* Ability for an operator to export secrets (*by providing a public key*),
  to use in other workflows.
* More documentation updates.
* More flexibility in SPIFFEID validation.
* Increased stability.
* Lots of documentation updates, especially around security and production
  setup.
* Static code analysis.
* Website enhancement: Versioned snapshots of the documentation.
* Option for **VSecM** to run in-memory; without having to rely on any backing
  store.
* Security: Ability to lock VSecM Safe.

### VSecM v0.21.0 (*codename: Andromeda*)

**Aug 15, 2023 - Sep, 11, 2023**

This was a stability-focused release. We focused on fixing bugs, improving
stability, and improving workflows and CI/CD pipelines. We also created
missing documentation and generated new video tutorials that feature the current
version of **VMware Secrets Manager**.

[Check out the release notes](/docs/changelog/) to learn more about what has
been added, changed, and fixed in this release.

{{ edit() }}
