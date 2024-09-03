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

### VSecM v0.28.0 (*codename: Hydra*)

**Jun 29, 2024 - Sep 30, 2024**

This iteration is about increasing coverage. We will focus on
unit tests.

Our goal is to reach at least 50% coverage.

In addition, we are targeting to fix certain low-hanging bugs and improve 
stability. 

[Here is a list of issues that are candidate for VSecM vHydra](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.28.0-candidate+).

### VSecM v0.29.0 (*codename: Indus*)

**Oct 01, 2024 - Oct 28, 2024**

In this iteration, we will focus on adding use cases and tutorials, along with
any stability and security improvement that may come our way.

[Here is a list of issues that are candidate for VSecM vIndus](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.29.0-candidate+).

### VSecM v0.30.0 (*codename: Lupus*)

**Oc 29, 2024 - Nov 25, 2024**

This iteration will be about adding more features that may be immediately
useful around **VSecM Sentinel** CLI.

[Here is a list of issues that are candidate for VSecM vLupus](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.30.0-candidate+).

### VSecM v0.31.0 (*codename: Mensa*)

**Nov 26, 2024 - Jan 6, 2025**

This iteration is about SDKs and KMS integration.

[Here is a list of issues that are candidate for VSecM vMensa](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.31.0-candidate+).

### VSecM v0.32.0 (*codename: Norma*)

**Jan 07, 2024 - Feb 03, 2024**

The overall theme of this iteration is issues labeled as enhancing the system's 
scalability, high availability (HA), and integration capabilities. Key areas of 
focus include improving HA modes for various components, enabling state 
federation and synchronization with external storage systems, enhancing 
documentation, and expanding the flexibility of key management through integration 
with cloud Key Management Services (KMS) and databases. 

These efforts aim to make VMware Secrets Manager more robust, reliable, 
and easier to integrate with other systems and environments.

This iteration is about visibility and metrics. We'll create a `/stats` and a
`/health` endpoint for **VSecM Safe** among other observability improvements.

[Here is a list of issues that are candidate for VSecM vNorma](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.32.0-candidate+).

### VSecM v0.33.0 (*codename: Orion*)

**Feb 04, 2024 - Mar 03, 2024**

This iteration is centered around enhancing the system's capabilities in high 
availability, scalability, and integration. Key areas include:

* **High Availability (HA)**: Demonstrating and improving HA modes for various 
  components like SPIRE and VSecM.
* **Federation and State Management**: Implementing federation of the identity
  control plane and considering state federation for VSecM.
* **Storage and Synchronization**: Enhancing the flexibility of key management 
  by enabling the use of external storage solutions, such as separate VSecM Safe 
  instances, cloud Key Management Services (KMS), databases, and persistent 
  volumes, with a focus on two-way synchronization.
* **Documentation and Usability**: Improving documentation for various features 
  and providing better configuration options in Helm charts.

[Here is a list of issues that are candidate for VSecM vOrion](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.33.0-candidate+).

### VSecM v0.34.0 (*codename: Perseus*)

**Mar 04, 2024 - Mar 31, 2025**

This iteration centers on enhancing the system's scalability, high availability, 
and integration capabilities. Key areas include:

* **High Availability (HA)**: Enhancing HA modes for components like SPIRE and 
  VSecM to ensure system reliability.
* **Federation and State Management**: Implementing federation for identity 
  control and state management to improve scalability and integration with other 
  systems.
* **Storage and Synchronization**: Increasing flexibility in key management by 
  enabling external storage solutions such as cloud KMS, databases, and 
  persistent volumes, along with ensuring two-way synchronization.
* **Helm Charts Customization**: Providing extensive customization options in 
  Helm charts for SPIRE server and agents, including the ability to use 
  different data stores, key managers, and telemetry configurations.
* **Documentation and User Requests**: Improving documentation for better 
  usability and addressing user requests for new features such as stats and 
  health endpoints, and key rotation capabilities.

[Here is a list of issues that are candidate for VSecM vPerseus](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.34.0-candidate+).

### VSecM v0.35.0 (*codename: Reticulum*)

**Apr 01, 2024 - Apr 28, 2025**

This iteration aims at enhancing security, workflow, and documentation. 
The open issues include:

* **Security Enhancements**: Use a separate VSecM Safe to store root keys instead of  
  a Kubernetes secret, improving security by isolating critical keys from the 
  main application infrastructure.

* **Workflow Improvements**: Ensuring the project includes an automated test 
  suite that provides at least 90% statement coverage, which aims to improve the 
  reliability and maintainability of the codebase.

* **Documentation Updates:**: Documenting the VSecM Sentinel OIDC authentication 
  feature, which will help users understand and implement this feature more 
  effectively.

[Here is a list of issues that are candidate for VSecM vReticulum](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.35.0-candidate+).

### VSecM v0.36.0 (*codename: Sagittarius*)

**Apr 31, 2025 - May 26, 2025**

This iteration revolves around enhancing functionality, improving integration 
capabilities, and expanding test coverage. The key areas include:

* **Enhanced Integration and Configuration**: Ability to configure SPIRE's key 
  manager in Helm charts. 
  * Configure **VSecM Sidecar** for dynamic secret updates.
  * Use AWS KMS as an alternate backup store. 
  * Demo and integrate OPA and VSecM.

* **Workflow and Automation Improvements**:
  * Integration of Sentinel OIDC Resource Server functionality into tests.
    
* **Use Cases and Examples**:
  * Demonstration of GitOps use cases.
  * Sample configurations and documentation for various use cases.

[Here is a list of issues that are candidate for VSecM vSagittarius](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.36.0-candidate+).

### VSecM v0.37.0 (*codename: Telescopium*)

**May 27, 2025 -- Jul 31 2025**

This iteration focuses on enhancing security, improving integration and workflow, 
and providing new features for better user experience. Here are the main points:

* **Security Enhancements**:
  * Conducting security reviews and independent security audits.
  * Introducing dynamic code analysis using fuzzing to enhance security measures.
* **Integration and Workflow**:
  * Facilitating integration with AWS and GCP through placeholders.
  * Creating a dedicated user for EKS provisioning with just-enough privileges.
  * Supporting the separation of audit logs from other logs for better security 
    management.
* New Features and Usability Improvements:
  * Allowing the VSecM Init container to decrypt mounted files using AES or age 
    decryption keys.
  * Enabling the creation and registration of AES keys or age key pairs for 
    workloads to decrypt shared encrypted files.
  * Demonstrating autoscaling use cases.
  * Adding new endpoints for statistics and health monitoring.

[Here is a list of issues that are candidate for VSecM vTelescopium](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.37.0-candidate+).

### VSecM v0.38.0 (*codename: Ursa*)

**Aug 01, 2025 -- Aug 28 2025**

This iteration is focused on enhancing the system's security, improving 
integration capabilities, and refining user experience. Key points include:

* **Security Enhancements**:
  * Introduction of high-trust modes, such as using a `PKCS#11` interface to 
    secure root keys.
  * Customizable kubelet verification for enhanced security.
* **Integration and Configuration**:
  * Ability to use Kubernetes as a backing store.
  * Configurable audit targets and customizable Helm charts.
* **User Experience Improvements**:
  * Creating Kubernetes Operators to inject Init Containers and Sidecars based 
    on annotations.
  * Adding optional policies for secrets and considering methods for seamless 
    secret rotation.
* **Documentation and Workflow**:
  * Documenting self-security assessments and OIDC authentication features.
  * Release workflow management for new versions.

[Here is a list of issues that are candidate for VSecM vUrsa](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.38.0-candidate+).

### VSecM v0.39.0 (*codename: Virgo*)

**Aug 29, 2025 -- Sep 25 2025**

This is mainly a security-focused iteration. Here are the main points:

* There is a plan to create a UI for VSecM Sentinel, leveraging the existing 
  OIDC Server functionality.
* Integration tests will include Sentinel OIDC Resource Server functionality.
* Implementation of a `ValidatingAdmissionWebhook` to ensure `clusterspiffeid`s 
  and `clusterspiffeid` templates have the correct format.
* Development of a self-security assessment documentation.
* Exploration of secure methods for sharing root key material across VSecM 
  instances in different clusters.
* Support for replication across multiple VSecM Safe instances.
* Creation of Kubernetes Operator to inject VSecM Init Container and 
  VSecM Sidecar-based on annotations.

[Here is a list of issues that are candidate for VSecM vVirgo](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.39.0-candidate+).

### VSecM v0.40.0 (*codename: Antlia*)

**Sep 26, 2025 -- Oct 32 2025**

This is a "*catch all*" that contains all remaining documented future plans.
We will create new iterations from it as the time gets closer.

[Here is a list of issues that are candidate for VSecM vUrsa](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.40.0-candidate+).

## Closed Iterations

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
