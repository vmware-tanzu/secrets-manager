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

title: Roadmap
layout: post
prev_url: /docs/releases/
permalink: /docs/roadmap/
next_url: /
---

<p class="github-button"
><a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0007-roadmap.md"
>edit this page on <strong>GitHub</strong> ✏️</a></p>

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

### VSecM v0.22.0 (*codename: Boötes*) 

**Sep 12, 2023 – Oct 9, 2023**

This release will be more about enhancing deployment workflows, testing automation
and CI/CD pipelines. We will also focus on improving the overall user experience.

* Ability for an operator to export secrets (*by providing a public key*), 
  to use in other workflows.
* Enhancements to template transformations.
* Better, machine-readable logs.
* Also, better audit logs.
* More documentation updates.
* Enabling automated tests and static code analysis.
* Preventing log tampering.
* Performance improvements on the website.

### VSecM v0.23.0 (*codename: Cassiopeia*) 

**Oct 10, 2023 – Nov 6, 2023**

This iteration will be focused on improving how **VMware Secrets Manager** 
logs and reports errors. We will also focus on improving the performance of the
**VMware Secrets Manager** website.

* `Secret`less VSecM: Ability to use VMware Secrets Manager **without** relying
  on Kubernetes `Secret`s. This will allow users to use **VMware Secrets Manager**
  without having to create Kubernetes `Secret`s at all—even for the master keys.
* Ability to use VSecM across clusters (*multi-cluster federation support*).
* Static code analysis.
* More automation and tests.
* More use-case video lectures.
* Website enhancement: Versioned snapshots of the documentation.

### VSecM v0.24.0 (*codename: Draco*) 

**Nov 7, 2023 – Dec 4, 2023**

This iteration will be focused on making **VMware Secrets Manager** able to 
ingest large amounts of secrets, without crashing or slowing down.

* Stream manipulation: Ability to ingest large amounts of secrets; also 
  ability to ingest longer secrets.
* More automation.
* Our goal is 90% code coverage by this checkpoint; or as far as we can get.
* Security: Ability to lock VSecM Safe.
* Documentation: A user guide for those who want to develop their own
  language-specific VSecM SDKs (*we’ll continue to support Go only for a while,
  however we still want to make it easier for others to develop their own SDKs*).
* Website optimization and asset cleanup.

### VSecM v0.25.0 (*codename: Eridanus*) 

**Dec 5, 2023 – Jan 1, 2024**

In this iteration, our focus will be in-memory usage of **VSecM** and also making
the **VSecM Sidecar** more robust.

* Option for **VSecM** to run in-memory; without having to rely on any backing store.
* Option for the **VSecM Sidecar** to kill the container when the bound secret changes.
* Introducing Operators to auto-inject VSecM sidecars and init containers to workloads.

### VSecM v0.26.0 (*codename: Fornax*) 

**Jan 2, 2024 – Jan 29, 2024**

This is an iteration focused on code stability, and community development.

* Validation and guardrails around VSecM-managed SVIDs.
* Community development efforts.
* Focus on observability.

### VSecM v0.27.0 (*codename: Gemini*) 

**Jan 30, 2024 – Feb 26, 2024**

We’ll create abstractions around certain **VMware Secrets Manager** components
to make further cloud integrations easier.

* Creating custom resources (`ClusterVSecMId`) for better abstraction.
* Improving usability and developer experience.

### VSecM v0.28.0 (*codename: Hydra*) 

**Feb 27, 2024 – Mar 25, 2024**

This iteration will be about providing access to **VSecM Sentinel** through
OIDC authentication. We will also focus on various compatibility issues before
we dive into cloud integration in the upcoming iterations.

The goals in this iteration could be a stretch and based on the workload of
the core maintainers, we might have to push some of these goals to the next
iteration, thus impacting the overall roadmap.

* Focus on auto-scaling.
* OIDC authentication.
* Using Redis as a shared backing store.
* Ability to deploy VSecM to any SPIFFE-compatible cluster that has agents
  that provide SPIFFE Workload API.

### VSecM v0.29.0 (*codename: Indus*) 

**Mar 26, 2024 – Apr 22, 2024**

This iteration will be about integrating **VMware Secrets Manager** with
**AWS KMS**.

* AWS KMS Integration

### VSecM v0.30.0 (*codename: Lupus*) 

**Apr 23, 2024 – May 20, 2024**

This iteration will be about integrating **VMware Secrets Manager** with
**Azure Key Vault**.

* Azure Key Vault Integration

### VSecM v0.31.0 (*codename: Mensa*) 

**May 21, 2024 – Jun 17, 2024**

This release is about security auditing and hardening.

* Create a self-security assessment.
* Perform security audit from a third-party security firm, or an internal
  VMware security team.

### VSecM v0.32.0 (*codename: Norma*) 

**Jun 18, 2024 – Jul 15, 2024**

This iteration will be about integrating **VMware Secrets Manager** with
**Google Cloud KMS**.

* Google Cloud KMS Integration

### VSecM v0.32.0 (*codename: Orion*)

**Jul 16, 2024 – Aug 12, 2024**

This iteration will be about OIDC support, and improving our OpenSSF conformance.

By this point, we expect to have at lest a silver, if not a gold badge from
OpenSSF.

We will also stretch our research into non-Kubernetes deployment options.

* Get an OpenSSF Silver Badge.
* Non-Kubernetes deployment options.
* Other integration options.
* Creating other official SDKs, along with VSecM Go SDK.
* Stability improvements.

## Closed Iterations

### VSecM v0.21.0 (*codename: Andromeda*)

**Aug 15, 2023 – Sep, 11, 2023**

This is a stability-focused release. We will focus on fixing bugs, improving
stability, and improving workflows and CI/CD pipelines. We will also create
missing documentation and generate new video tutorials that feature the current
version of **VMware Secrets Manager**.

[Check out the release notes](/docs/changelog/) to learn more about what has
been added, changed, and fixed in this release.