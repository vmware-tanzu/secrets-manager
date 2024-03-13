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

title: Roadmap
layout: post
prev_url: /docs/releases/
permalink: /docs/roadmap/
next_url: /
---

## Introduction

This is a page where we publish our approximate roadmap for **VMware Secrets
Manager** for Cloud-Native Apps. Note that this is not a commitment to deliver
any of the features listed here, and that the roadmap is subject to change at
any time without notice.

Whenever we release a new version of **VMware Secrets Manager**, we will update
this page, and also [the changelog](/docs/changelog/) to reflect the changes.

> **One-Year Window**
> 
> This page will only contain information about the next 12 months of the
> project. We will update the roadmap every release, and remove the completed
> items from the list, and add a new iteration at the end of the list.

## Active Iterations

### VSecM v0.24.0 (*codename: Draco*) 

**Feb 29, 2023 - Mar 27, 2023**

To automate things and be able to dynamically follow issues better, from 
this point on we'll label them and share the GitHub filter here, instead
of outlining them one by one.

This iteration will be focused on demos and documentation.

[Here is a list of issues that are candidate for VSecM vDraco](https://github.com/vmware-tanzu/secrets-manager/issues?q=is%3Aissue+label%3Av0.24.0-candidate+).

### VSecM v0.25.0 (*codename: Eridanus*) 

**Mar 28, 2023 - Apr 24, 2024**

This iteration, our focus will be on testing, test automation, and test coverage.

[Here is a list of issues that are candidate for VSecM 
vEridanus](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.25.0-candidate+).

### VSecM v0.26.0 (*codename: Fornax*) 

**Apr 25, 2024 - May 22, 2024**

This iteration is all about SDKs (C++, Python, Rust, and Java).

[Here is a list of issues that are candidate for VSecM vFornax
](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.26.0-candidate+).

### VSecM v0.27.0 (*codename: Gemini*) 

**May 23, 2024 - Jun 19, 2024**

The sole focus of this iteration is increasing test coverage product-wide, 
as best as we can.

[Here is a list of issues that are candidate for VSecM vGemini
](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.27.0-candidate+).

### VSecM v0.28.0 (*codename: Hydra*) 

**Jun 20, 2024 - Jul 17, 2024**

This iteration is about improving **VSecM Sentinel** to ingest different kinds
of secrets, including very long secrets from a stdin pipe. This will be a 
feature-focused iteration.

[Here is a list of issues that are candidate for VSecM vHydra](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.28.0-candidate+).

### VSecM v0.29.0 (*codename: Indus*) 

**Jul 18, 2024 - Aug 14, 2024**

This is a bug-bash iteration where we'll leave no bug left behind.

[Here is a list of issues that are candidate for VSecM vIndus](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.29.0-candidate+).

### VSecM v0.30.0 (*codename: Lupus*) 

**Aug 15, 2024 - Sep 11, 2024**

This iteration is about enabling external access to **VSecM Safe** in a secure
manner. We will also focus on secure logging and preventing log tampering.

[Here is a list of issues that are candidate for VSecM vLupus](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.30.0-candidate+).

### VSecM v0.31.0 (*codename: Mensa*) 

**Sep 12, 2024 - Oct 09, 2024**

This iteration is about visibility and metrics. We'll create a `/stats` and a
`/health` endpoint for **VSecM Safe** among other observability improvements.

[Here is a list of issues that are candidate for VSecM vMensa](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.31.0-candidate+).

### VSecM v0.32.0 (*codename: Norma*) 

**Oct 10, 2024 - Nov 06, 2024**

In this iteration we'll do a full-blown self-security assessment and publish
our findings. These findings will be used later for a third-party security
audit.

[Here is a list of issues that are candidate for VSecM vNorma](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.32.0-candidate+).

* Google Cloud KMS Integration

### VSecM v0.33.0 (*codename: Orion*)

**Nov 07, 2024 - Dec 04, 2024**

In this iteration, we will focus on non-Kubernetes **VSecM** deployments.

[Here is a list of issues that are candidate for VSecM vOrion](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.33.0-candidate+).

### VSecM v0.34.0 (*codename: Perseus*)

**Dec 05, 2024 - Jan 01, 2025**

This iteration is about UI/UX fixed on the website, and CSS cleanup.

[Here is a list of issues that are candidate for VSecM vPerseus](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.34.0-candidate+).

### VSecM v0.35.0 (*codename: Reticulum*)

**Jan 02, 2025 - Jan 29, 2025**

We'll focus on creating Kubernetes Operators around **VSecM** to better 
coordinate certain workflows.

[Here is a list of issues that are candidate for VSecM vReticulum](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.35.0-candidate+).

### VSecM v0.36.0 (*codename: Sagittarius*)

**Jan 30, 2025 - Feb 26, 2025**

This iteration is focused on better state management.

[Here is a list of issues that are candidate for VSecM vSagittarius](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.36.0-candidate+).

### VSecM v0.37.0 (*codename: Telescopium*)

**Feb 27, 2025 -- Mar 26 2025**

This is a "*catch all*" that contains all remaining documented future plans.
We will create new iterations from it as the time gets closer.

[Here is a list of issues that are candidate for VSecM vTelescopium](https://github.com/vmware-tanzu/secrets-manager/issues?q=+label%3Av0.37.0-candidate+).

## Closed Iterations

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

<p class="github-button">
  <a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0280-roadmap.md">
    Suggest edits ✏️
  </a>
</p>