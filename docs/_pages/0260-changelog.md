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
><a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0005-changelog.md"
>edit this page on <strong>GitHub</strong> ✏️</a></p>

## Recent Updates

## [v0.21.0] - 2023-09-08

### Added

* Documentation updates to make the project align with the current status of
  **VSecM**.
* Migrate existing Aegis documentation to the [new VMware Secrets Manager
  documentation site](https://vsecm.com).
* Updated [contributing guidelines](https://vsecm.com/docs/contributing) to make it easier for first-time 
  contributors.
* Published a formal [project governance model](https://vsecm.com/docs/governance/)
* Added a [blog section](https://vsecm.com/docs/blog/) to the website.
* Decided to add a new helm chart per each release.
* Added instructional video content to the [showcase section](https://vsecm.com/docs/showcase/).

### Fixed

* Minor bugfixes after migration; ensuring feature and behavior parity with
  Aegis.
* Sentinel and Safe’s Identity.yaml need stricter matchers.

### Security

* Updated the [security policy](https://vsecm.com/docs/security/), clarifying 
  our ideal response time for security vulnerabilities.
* Fixed: [Active Support Possibly Discloses Locally Encrypted Files](https://github.com/vmware-tanzu/secrets-manager/security/dependabot/2)

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