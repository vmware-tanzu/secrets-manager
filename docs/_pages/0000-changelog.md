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

title: VSecM Changelog
layout: post
prev_url: /
next_url: /docs/about/
permalink: /docs/changelog/
---

<p class="github-button"
><a
href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0000-changelog.md"
>edit this page on <strong>GitHub</strong> ✏️</a></p>

## Recent Updates

* Documentation updates to make the project align with the current status of
  **VSecM**.

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