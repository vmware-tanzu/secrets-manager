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

title = "VSecM Changelog"
weight = 11
+++

## Recent Updates

* Added the ability to have regex-based SPIFFE ID matchers.
* Enabled stricter validation on SPIFFE IDs to reduce configuration errors.
* Added ability to optionally use multiple worker nodes for the development
  clusters.
* Introduced `helm-docs` to automatically augment the documentation with the
  Helm chart's values.yaml.
* Added the ability to deploy VSecM without SPIRE Controller Manager. In this
  mode, the operator will need to manually create SPIRE Server registration
  entries.
* Added the ability to not create `ClusterSPIFFEID`s for the VSecM components
  automatically. In this mode, the operator will need to manually create those
  required `ClusterSPIFFEID`s.
* Documentation updates.
* Introduced `helm-docs` for helm chart documentation.
* Replaced `github.com/pkg/errors` with the native `errors` package 
  to reduce the number of dependencies and the codebase more
  secure and maintainable.

## [0.25.3] - 2024-05-17

* Removed some configuration options including
  `VSECM_MANUAL_ROOT_KEY_UPDATES_K8S_SECRET` because how the root key will
  be updated will be depending on backing store implementation. And it does
  not make sense for an operator updating the root key in memory but not
  updating the backing Kubernetes secret. That could bring inconsistencies
  to the system.
* Removed `VSECM_SAFE_REMOVE_LINKED_K8S_SECRETS` since we have long
  deprecated and removed the `-k` flag that was dealing with the linked
  Kubernetes secrets. Again, future behavior will be contingent upon the
  backing store plugins that will be implemented.
* Removed Kubernetes secrets deletion queue because we do not link Kubernetes
  secrets to workloads anymore. Deletion of ad-hoc VSecM-generated Kubernetes
  `Secret`s will be handled by upcoming configuration options. Right now,
  VSecM Safe can only create and update, but not delete Kubernetes `Secret`s.
* Stability improvements, including adding "exponential backoff"s to places
  where requests can be retried before giving up; also letting the apps
  crash (*and be re-crated by the scheduler*) if certain critical requests fail
  even after a fair amount exponentially-backed-off of retries (*10 by default*).
* An entire overhaul of the documentation website: It is now faster, more 
  accessible, more usable, easier to navigate and follow.
* Added an experimental Java SDK. The keyword here is: **experimental**; we
  do know that it does not work out-of-the box, so we are not providing any
  documentation yet: Feel free to join our Slack channel to learn more about
  how best you can use it.
* Refactorings and improvements across the entire codebase.
* Introduced [Architectural Decision Records](https://vsecm.com/documentation/architecture/adr-intro/)

## [0.25.2] - 2024-05-06

This release introduced many structural changes. The functionality remains
the same, but the codebase is more organized and easier to maintain.
We had to temporarily disable some of the unit tests to make the release
happen on time. We will re-enable them before the next release.

### Changed

* Simplified audit journaling
* Refactoring and code organization
* Now helm-chart can deploy VSecM and SPIRE to any namespace, before it had
  to be `vsecm-system` and `spire-system` respectively.
* removed "backing store" from secret meta info; backing store will be set
  at a global level.
* removed `-b` (backing store) flag from VSecM Sentinel's CLI too.
* Added certain useful methods from `internal` packages to the `core` package
  to make it more reusable. These functionalities may be part of the SDK too,
  later.
* Organized imports and functions according to the project standards.
* Renamed certain modules and functions for clarity.
* Introduced certain environment variables whose functionalities will be
  implemented later.
* updated helm charts, removed hard coded namespace references from service URLs.

## [0.25.1] - 2024-04-26

This was a stability and reliability release. We have made several improvements
to VSecM Sentinel, helm charts, and Kubernetes manifests to make the system
more reliable and resilient.

### Changed

* Converted VSecM Safe and SPIRE Server to StatefulSets (because they are stateful).
* VSecM Sentinel "init command" loop now exits the container if it cannot execute
  commands after exponential backoff. The former behavior was to retry forever,
  and that was not a cloud-native way of handling the situation. Panicking
  early and thus killing the pod fixed issues with things like persistent volumes
  and CSI drivers.

### Fixed

* Minor bug fixes in the VSecM Sentinel init command workflow.

## [0.25.0] - 2024-04-24

### Added

* Documentation updates.
* Added liveness and readiness probes to SPIRE Server and SPIRE Agent.
* Added pod priority classes to SPIRE Server, SPIRE Agent, and VSecM pods
  to ensure that VSecM components are prioritized and maintained in the
  event of resource constraints.
* VSecM Sentinel Init Commands can now wait a configurable amount of time
  before running. This feature is useful when you want to delay the execution
  of the init commands to ensure that other components are ready.
* VSecM Sentinel can now wait before marking Init Commands as successful.
  This feature is useful when you want to delay the readiness of VSecM Sentinel
  until other components are ready.
* VSecM Sentinel Init Command can now parse and understand all VSecM Sentinel
  commands.
* Added Generated protobuffer files into the source code for ease of maintenance.

### Changed

* Removed the tombstone feature, we use VSecM Keystone instead of tombstone,
  which is more reliable, secure, and under our control.
* Reliability improvements in VSecM Sentinel. For example, VSecM Sentinel does
  not wait forever in a loop for VSecM Safe to be ready. Instead, it crashes
  after a grace period, and the orchestrator can restart it in a more cloud-native way.
* SPIRE Server is now a `StatefulSet` by default instead of a `Deployment`.
  This change ensures that SPIRE Server has a stable identity across restarts.
* VSecM Keystone, and VSecM Keystone secrets are being used instead of tombstone.
* Various other stabilization improvements.

### Fixed

* Minor bug fixes and feature enhancements.

### Security

* Fixed CVE-2023-45288 [net/http, x/net/http2: close connections when receiving too many headers](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/12)
* Fixed GHSA-pxvg-2qj5-37jq [Nokogiri updates packaged libxml2 to v2.10.4 to resolve multiple CVEs](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/13)
* Fixed GHSA-xc9x-jj77-9p9j [Nokogiri update packaged libxml2 to v2.12.5 to resolve CVE-2024-25062](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/14)
* Fixed GHSA-vcc3-rw6f-jv97 [Use-after-free in libxml2 via Nokogiri::XML::Reader](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/15)
* Addressed CVE-2020-8559 [Privilege Escalation in Kubernetes](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/16)

## [v0.24.1] - 2024-03-31

### Added

* Added 4 new use cases to the documentation.

### Fixed

* VSecM Sentinel was not honoring the tombstone secret, now it is fixed.

## [v0.24.0] - 2024-03-28

### Added

* Kickstarted SDK work for Java.
* `VSECM_MANUAL_ROOT_KEY_UPDATES_K8S_SECRET` environment variable added for
  giving an option to updating internal k8s secrets when manual root key provided.
* Added additional logs to VSecM Sentinel to help with debugging.

### Fixed

* Quickstart guide on the website was not working as expected, now it is fixed.

## [v0.23.3] - 2024-03-24

### Added

* Added the Helm charts the ability for SPIRE Server to use Persistent Volumes
  for its data.
* Introduced **VSecM Keystone** a workload that waits until all the "init commands"
  that VSecM Sentinel runs are completed. This feature is useful, especially when
  an orchestrator watches for the readiness of VSecM Keystone to bring up other
  workloads that depend on the secrets that VSecM Sentinel initializes.
* Now, one secret can be associated with multiple workloads in multiple
  namespaces. This feature is useful when you want to share a secret across
  multiple workloads in different namespaces.
* Added image pull secrets to SPIRE Server and SPIRE Agent Helm charts.
* Added [Kampus Discord Server][kampus] as a welcoming and supporting community
  for VSecM users and contributors. This is an additional community that
  augments the official **Slack** workspace. The **Slack** workspace is still
  the primary community for VSecM. Kampus is a global community; however, its
  core audience is Turkish-speaking people. The community is open to everyone.
* By adding Kampus as a supported community, we aim to:
    * Acknowledge and express gratitude for the Kampus community's ongoing
      support and contributions.
    * Facilitate a more integrated and cohesive ecosystem for current and future
      contributors.
    * Enhance accessibility for new contributors seeking guidance or looking to
      engage with the project community.
    * Foster a diverse and inclusive environment where all members can share,
      learn, and contribute to the project's success.

[kampus]: https://discord.gg/kampus

### Changed

* **BREAKING**: Removed the `-k` flag from Sentinel, as the `k8s:` prefix was
  a better way that does an identical job. This change also simplified the
  internal workings of VSecM Safe, making it more efficient, reliable, and
  easier to maintain.
* VSecM documentation now has a dark mode for better readability. In addition,
  the typography and layout of the documentation have been improved for a more
  consistent and user-friendly experience.

### Fixed

* Integration tests were failing, now they are fixed.
* Various minor bugfixes.
* Performance improvements and asset cleanup in the documentation website.

### Security

* SPIRE Server Helm charts was using `NodePort`; we defaulted it to the more
  secure `ClusterIP` in the Helm charts.

## [v0.23.2] - 2024-03-13

### Added

* VSecM Sentinel can now act as an OIDC Resource Server (experimental). This
  feature is disabled by default, and can be enabled by an environment variable.
  When you enable it, you should also ensure the security of the OIDC Server
  as breaching it will give direct access to VSecM. This feature changes the
  attack surface of the system and should be implemented only if you are
  extremely sure of what you are doing.
* Documented all public methods in the codebase. This will help
  contributors to understand the codebase better and make it easier to
  contribute.
* We now have an official "VSecM Inspector" container image that can be used
  to inspect the secrets bound to workloads without having to shell into
  the workloads. This is especially helpful when you want to debug a workload’s
  secrets without needing to uninstall or change the source code of the workload.
* Unit tests to increase coverage.

### Changed

* We now have a Go-based integration test suite instead of the former bash-based
  one. This change makes the tests more reliable and easier to maintain, while
  we can leverage the Go language’s powerful primitives to make the tests
  readable, maintainable, and scalable.
* VSecM components have sensible "memory" lower limits in helm charts (before
  it was left for the end-user to decide, now we provide a starting point
  while encouraging the user to do their own benchmarks to update the
  resource limits to their production needs.)
* Updated the log level of all VSecM components to the highest (7, TRACE).
  This setting is to help VSecM users to diagnose and debug potential
  installation issues during initial deployment. Once you are sure that things
  work as expected, you are encouraged to change the log level to a more
  sensible value (like, 3, DEBUG).
* Refactorings to make the code easier to follow.

### Fixed

* VSecM Sentinel's "Init Command" loop had a logic error that was preventing the
  initialization command to function under certain edge conditions. It’s now
  fixed.

### Security

* Updated SPIRE Server, SPIRE Client, and SPIRE Controller Manager images to
  their latest version.
* Increased the Go version to the recent stable.
* Fixed CVE-2024-28180 [Go JOSE vulnerable to Improper Handling of Highly
  Compressed Data](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/11)

## [v0.23.0] - 2024-03-01

### Added

* VSecM Sentinel now waits for VSecM Safe to be ready before running init
  commands.
* Documentation updates and code refactoring.
* Updated dependencies for better security.
* Updated Go version to 1.22.0--the latest stable version.

## [v0.22.5] - 2024-02-26

### Added

* Provisioned an public ECR registry to deploy and test VSecM on EKS.
* Added a GitHub Actions workflow to generate a test coverage badge, and
  coverage reports.
* Added the ability to use a persistent volume for VSecM Safe.

### Changed

* Bumped SPIRE Server and SPIRE Agent to the latest versions (1.9.0).
* VSecM Sentinel logs now have a correlation ID to make it easier to trace
  logs initiated by different requests.
* Improvements to the logging-and-auditing-related code.
* Deleting a VSecM Safe "secret" now also deletes the associated Kubernetes
  secret, if it exists.
* VSecM Safe now has a more robust retry strategy for creating and updating
  Kubernetes secrets.

## [v0.22.4] - 2024-02-17

### Added

* Added the ability to associate multiple namespaces with a single VSecM secret.
* Added a tombstone feature to VSecM Sentinel, so that when the init commands run to
  completion, they will not run again if VSecM Sentinel is evicted and restarted.
* Created an ECR repository to test edge versions of VSecM container images that
  have not been released yet.
* Added audit logging capabilities to VSecM Sentinel.

### Fixed

* Secrets creation now has a backoff policy and will retry if the first attempt fails.
* `VSECM_LOG_LEVEL` was left at `7` (verbose) in the charts, defaulting to `3` (warn).

### Changed

* Moved "*VMware, Inc.*" from the copyright headers, replacing it with "*VMware
  Secrets Manager contributors*".
* Default resource limits for Minikube initialization scripts to a more reasonable
  values for development. These are still configurable via environment variables.

### Security

* Fixed CVE-2024-25062 [When using the XML Reader interface with DTD validation
  and XInclude expansion enabled, processing crafted XML documents can lead to
  an xmlValidatePopElement use-after-free](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/10)

## [v0.22.3] - 2024-02-04

### Added

* Added the ability to run init commands during bootstrap to VSecM Sentinel.
* Added more test cases to the project.
* Added coverage targets to tests.
* Added scripts to test the project on a cloud AWS EKS cluster.

### Fixed

* Bug fixes and performance improvements.
* `make h` and `make help` had a cosmetic regression, which is now fixed.

### Changed

* Upgraded SPIRE Controller Manager to v0.4.1.
* Documentation updates, especially around establishing a secure production deployment.

## [v0.22.2] - 2024-01-14

### Added

* Documentation updates.
* Ability to create and update Kubernetes secrets without attaching the secret
  to a workload. This is useful for legacy use cases, or when you don't have
  direct access to the app's source code or deployment manifests.

## [v0.22.1] - 2024-01-11

### Added

* Added expiration and "invalid before" dates to secrets.
* Implemented a basic CI automation that runs test whenever there is a change
  in the `main` branch. The automation runs unit and integration tests and
  send status updates upon failure.
* Upgraded SPIRE and SPIFFE CSI Driver to the latest versions.
* Minor fixes and documentation updates.

## [v0.22.0] - 2024-01-08

### Added

* Documentation updated, especially around production usage and security.
* Added a `make commit` helper for a `better-commits` workflow.
* Added a PR template.
* Achieved great progress towards Open SSF Best Practices compliance; reaching
  93% of the requirements.
* Added ability to generate random secrets based on a pattern.
* Added ability to export encrypted secrets.

### Changed

* **BREAKING**: Certain environment variables are renamed to be more consistent
  with the rest of the project. The old variables are not supported anymore.
  check out the **configuration** section of the documentation for more details.
* Updated SPIRE, SPIRE Controller Manager, and SPIFFE CSI Driver to the latest
  versions.
* Moved older versions of the manifests to a `k8s` branch, and older snapshots
  of documentation to a `docs` branch to keep the `main` branch clean.

### Fixed

* Fixes on workflow scripts to have a more streamlined build process and
  development experience.
* Minor bugfixes and code enhancements.

## [v0.21.5] - 2023-12-18

### Changed

* **BREAKING**: Environment variables related to SPIFFEID are renamed from
  i.e. `VSECM_SENTINEL_SVID_PREFIX` to `VSECM_SPIFFEID_PREFIX_SENTINEL`.

### Added

* Documentation updates on security, production installation recommendations,
  and `kind` cluster usage for development.
* Minor code enhancements.

### Security

* Fixed CVE-2023-48795 [Russh vulnerable to Prefix Truncation Attack against
  ChaCha20-Poly1305 and Encrypt-then-MAC](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/9)

## [v0.21.4] - 2023-11-30

This patch release includes one security update, a minor refactoring, and
documentation updates.

### Security

* This is a patch release to address GHSA-2c7c-3mj9-8fqh [Decryption of malicious
  PBES2 JWE objects can consume unbounded system
  resources](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/8)

## [v0.21.3] - 2023-11-03

### Added

* Started experimental work on multi-cluster secret federation.
* Various Documentation updates.
* Automated Kubernetes manifest creation from Helm charts.

### Security

* Fixed GHSA-m425-mq94-257g [gRPC-Go HTTP/2 Rapid Reset
  vulnerability](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/7)

## [v0.21.2] - 2023-10-18

This is a purely security-focused release that fixes several vulnerabilities and
also hardens the AES encryption flow against time-based attacks.

### Security

* Fixed CVE-2023-3978 [Improper rendering of text nodes in
  golang.org/x/net/html](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/4)
* Fixed CVE-2023-39325 [HTTP/2 rapid reset can cause excessive work
  in net/http](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/5)
* Fixed CVE-2023-44487 [swift-nio-http2 vulnerable to HTTP/2 Stream Cancellation
  Attack](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/6)
* Fixed an issue with possible memory overflow when doing a cryptographic size
  computation.
* Added a configurable throttle to AES IV computation to make it harder to
  perform time-based attacks.
* The computed AES IV is zeroed out after use for additional security.

## [v0.21.1] - 2023-10-11

### Added

* Fixed `spire-controller-manager`'s version. The older setup was fixed on
  `nightly` which was causing ad-hoc issues.

### Changed

* Performance update: VSecM Sentinel now honors `SIGTERM` and `SIGINT` signals
  and gracefully shuts down when the pod is killed.
* Performance update: VSecM Safe is now leveraging several goroutines to speed
  up some of the blocking code paths during bootstrapping and initialization.
* Minor updates to the documentation.

### Security

* VSecM Safe has stricter validation routines for its identity.
* Added VSecM Keygen: a utility application that generates VSecM Safe's
  bootstrapping keys if you want an extra level of security and control the
  creation of the root key.

## [v0.21.0] - 2023-09-08

### Added

* Documentation updates to make the project align with the current status of
  **VSecM**.
* Migrate existing Aegis documentation to the [new VMware Secrets Manager
  documentation site](https://vsecm.com).
* Updated [contributing guidelines](@/documentation/development/contributing.md) to make it easier for first-time
  contributors.
* Published a formal [project governance model](@/documentation/operations/governance.md).
* Added a [blog section](https://vsecm.com/docs/blog/) to the website.
* Decided to add a new helm chart per each release.
* Added instructional video content to the [showcase section](https://vsecm.com/docs/showcase/).

### Fixed

* Minor bugfixes after migration; ensuring feature and behavior parity with
  Aegis.
* Implemented stricter matchers for VSecM Sentinel and VSecM Safe's
  `Identity.yaml`s.

### Security

* Updated the [security policy](https://vsecm.com/docs/security/), clarifying
  our ideal response time for security vulnerabilities.
* Fixed a minor vulnerability in `activesupport` dependency:
  ([CVE-2023-38037](https://access.redhat.com/security/cve/cve-2023-38037)).
  [fix](https://github.com/vmware-tanzu/secrets-manager/pull/215);
  [dependabot](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/2).
  The vulnerability affects only the website build process, not the **VSecM**
  codebase itself. It is not exploitable in our case, but we still wanted to
  fix it.

## [v0.20.0] - 2023-07-27

### Added

* Migrated the source code from <https://github.com/shieldworks/aegis> to
  <https://github.com/vmware-tanzu/secrets-manager>
* Did necessary changes for the project to run build and pass tests.
* Created new container image repositories at <https://hub.docker.com/u/vsecm>.

### Changed

* Minor changes to build and deployment scripts.
* **BREAKING**: The binary that `vsecm-sentinel` uses is called `safe` right
  now (*formerly it was `aegis`*).

{{ edit() }}
