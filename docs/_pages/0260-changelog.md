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

* Documentation updates to make the project align with the current status of
  **VSecM**.
* Migrate existing Aegis documentation to the [new VMware Secrets Manager
  documentation site](https://vsecm.com).
* Minor bugfixes after migration; ensuring feature and behavior parity with
  Aegis.
* Updated the [security policy](https://vsecm.com/docs/security/), clarifying 
  our ideal response time for security vulnerabilities.
* Updated [contributing guidelines](https://vsecm.com/docs/contributing) to 
  make it easier for first-time contributors.
* Published a formal [project governance model](https://vsecm.com/docs/governance/)
* Improvements in helm charts.
* [Fixed a minor vulnerability in `activesupport` dependency
  (CVE-2023-38037)](https://github.com/vmware-tanzu/secrets-manager/pull/215).
  The vulnerability affects only the website build process, and not the
  **VSecM** codebase itself. It is not exploitable in our case, but we still
  wanted to fix it.

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
Changed
Deprecated
Removed
Security
-->