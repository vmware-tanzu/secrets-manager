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

title: Changelog
layout: post
prev_url: /
permalink: /docs/changelog/
next_url: /docs/releases/
---

<p class="github-button"
><a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0260-changelog.md"
>edit this page on <strong>GitHub</strong> ✏️</a></p>

## Recent Updates

* We now have a Go-based integration test suite instead of the former bash-based
  one. This change makes the tests more reliable and easier to maintain, while
  we can leverage the Go language’s powerful primitives to make the tests
  readable, maintainable, and scalable.
* Documented all public methods in the codebase. This will help
  contributors to understand the codebase better and make it easier to
  contribute.

## [v0.23.0] - 2024-03-01

### Added

* VSecM Sentinel now waits for VSecM Safe to be ready before running init 
  commands.
* Documentation updates and code refactoring.

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
  i.e. `VSECM_SENTINEL_SVID_PREFIX` to `VSECM_SENTINEL_SPIFFEID_PREFIX`.

### Added

* Documentation updates on security, production installation recommendations,
  and `kind` cluster usage for development.
* Minor code enhancements.

### Security

* Fixed CVE-2023-48795 [Russh vulnerable to Prefix Truncation Attack against ChaCha20-Poly1305 and Encrypt-then-MAC](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/9)

## [v0.21.4] - 2023-11-30

This patch release includes one security update, a minor refactoring, and
documentation updates.

### Security

* This is a patch release to address GHSA-2c7c-3mj9-8fqh [Decryption of malicious
  PBES2 JWE objects can consume unbounded system resources](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/8)

## [v0.21.3] - 2023-11-03

### Added

* Started experimental work on multi-cluster secret federation.
* Various Documentation updates.
* Automated Kubernetes manifest creation from Helm charts.

### Security

* Fixed GHSA-m425-mq94-257g [gRPC-Go HTTP/2 Rapid Reset vulnerability](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/7)

## [v0.21.2] - 2023-10-18

This is a purely security-focused release that fixes several vulnerabilities and
also hardens the AES encryption flow against time-based attacks.

### Security

* Fixed CVE-2023-3978 [Improper rendering of text nodes in golang.org/x/net/html](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/4)
* Fixed CVE-2023-39325 [HTTP/2 rapid reset can cause excessive work in net/http](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/5)
* Fixed CVE-2023-44487 [swift-nio-http2 vulnerable to HTTP/2 Stream Cancellation Attack](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/6)
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
* Updated [contributing guidelines](https://vsecm.com/docs/contributing) to make it easier for first-time
  contributors.
* Published a formal [project governance model](https://vsecm.com/docs/governance/).
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


<!--
Added
Fixed
Changed
Deprecated
Removed
Security
-->
