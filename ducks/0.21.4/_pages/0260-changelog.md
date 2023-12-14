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

* TBD

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

* Fixed `spire-controller-manager`’s version. The older setup was fixed on
  `nightly` which was causing ad-hoc issues.

### Changed

* Performance update: VSecM Sentinel now honors `SIGTERM` and `SIGINT` signals 
  and gracefully shuts down when the pod is killed.
* Performance update: VSecM Safe is now leveraging several goroutines to speed 
  up some of the blocking code paths during bootstrapping and initialization.
* Minor updates to the documentation.

### Security

* VSecM Safe has stricter validation routines for its identity.
* Added VSecM Keygen: a utility application that generates VSecM Safe’s
  bootstrapping keys if you want an extra level of security and control the
  creation of the master key.

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
* Implemented stricter matchers for VSecM Sentinel and VSecM Safe’s 
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