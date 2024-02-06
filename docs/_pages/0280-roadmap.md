---
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title: Roadmap
layout: post
prev_url: /docs/releases/
permalink: /docs/roadmap/
next_url: /
---

<p class="github-button"
><a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0280-roadmap.md"
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

### VSecM v0.22.0 (_codename: Boötes_)

**Sep 12, 2023 – Jan 31, 2024**

This is a relatively longer release because due to the “time-stop” effect of the
holiday season, the majority of the core contributors will be spending quality
time with their loved ones and recharging their batteries for the upcoming year.

This release will be more about enhancing deployment workflows, testing automation
and CI/CD pipelines. We will also focus on improving the overall user experience.

-   Ability for an operator to export secrets (_by providing a public key_),
    to use in other workflows.
-   More documentation updates.
-   More flexibility in SPIFFEID validation.
-   Increased stability.
-   Lots of documentation updates, especially around security and production
    setup.
-   Static code analysis.
-   Website enhancement: Versioned snapshots of the documentation.
-   Option for **VSecM** to run in-memory; without having to rely on any backing
    store.
-   Security: Ability to lock VSecM Safe.

### VSecM v0.23.0 (_codename: Cassiopeia_)

**Feb 01, 2024 – Feb 28, 2024**

This iteration will be focused on improving how **VMware Secrets Manager**
logs and reports errors. We will also focus on improving the performance of the
**VMware Secrets Manager** website.

-   `Secret`less VSecM: Ability to use VMware Secrets Manager **without** relying
    on Kubernetes `Secret`s. This will allow users to use **VMware Secrets Manager**
    without having to create Kubernetes `Secret`s at all—even for the root keys.
-   Ability to use VSecM across clusters (_multi-cluster federation support_).

### VSecM v0.24.0 (_codename: Draco_)

**Feb 29, 2023 – Mar 27, 2023**

This iteration will be focused on making **VMware Secrets Manager** able to
ingest large amounts of secrets, without crashing or slowing down.

-   Ability to generated random password and assign them to workloads based
    on a predefined pattern.
-   Various bugfixes.
-   Ability to deploy **VSecM** to be able to use and existing SPIRE setup.
-   Documentation: A user guide for those who want to develop their own
    language-specific VSecM SDKs (_we’ll continue to support Go only for a while,
    however we still want to make it easier for others to develop their own SDKs_).
-   Security: Ability to customize kubelet verification through helm charts.

### VSecM v0.25.0 (_codename: Eridanus_)

**Mar 28, 2023 – Apr 24, 2024**

In this iteration, our focus will be in-memory usage of **VSecM** and also making
the **VSecM Sidecar** more robust.

-   Stream manipulation: Ability to ingest large amounts of secrets; also
    ability to ingest longer secrets.
-   More automation.
-   Automated generation of coverage reports.
-   Our goal is 90% code coverage by this checkpoint; or as far as we can get.
-   Website optimization and asset cleanup.
-   Enhancements to template transformations.
-   Better, machine-readable logs.
-   Also, better audit logs.

### VSecM v0.26.0 (_codename: Fornax_)

**Apr 25, 2024 – May 22, 2024**

This is an iteration focused on code stability, and community development.

-   Validation and guardrails around VSecM-managed SVIDs.
-   Community development efforts.
-   Focus on observability.
-   Enabling automated tests and static code analysis.
-   Performance improvements on the website.

### VSecM v0.27.0 (_codename: Gemini_)

**May 23, 2024 – Jun 19, 2024**

We’ll create abstractions around certain **VMware Secrets Manager** components
to make further cloud integrations easier.

-   Improving usability and developer experience.
-   Preventing log tampering.
-   More automation and tests.
-   More use-case video lectures.

### VSecM v0.28.0 (_codename: Hydra_)

**Jun 20, 2024 – Jul 17, 2024**

This iteration will be about providing access to **VSecM Sentinel** through
OIDC authentication. We will also focus on various compatibility issues before
we dive into cloud integration in the upcoming iterations.

The goals in this iteration could be a stretch and based on the workload of
the core maintainers, we might have to push some of these goals to the next
iteration, thus impacting the overall roadmap.

-   Focus on auto-scaling.
-   OIDC authentication.
-   Using Redis as a shared backing store.
-   Ability to deploy VSecM to any SPIFFE-compatible cluster that has agents
    that provide SPIFFE Workload API.
-   Using a separate VSecM Safe instance to store the root cryptographic keys
    as VSecM Safe secrets for better security.
-   Introducing refresh hint for sidecars and other consumers.

### VSecM v0.29.0 (_codename: Indus_)

**Jul 18, 2024 – Aug 14, 2024**

This iteration will be about integrating **VMware Secrets Manager** with
**AWS KMS**.

-   Option for the **VSecM Sidecar** to kill the container when the bound secret changes.
-   AWS KMS Integration
-   Using Redis (or PostgreSQL) as a shared backing store.

### VSecM v0.30.0 (_codename: Lupus_)

**Aug 15, 2024 – Sep 11, 2024**

This iteration will be about integrating **VMware Secrets Manager** with
**Azure Key Vault**.

-   Azure Key Vault Integration

### VSecM v0.31.0 (_codename: Mensa_)

**Sep 12, 2024 – Oct 09, 2024**

This release is about security auditing and hardening.

-   Create a self-security assessment.
-   Perform security audit from a third-party security firm, or an internal
    VMware security team.
-   Scalability improvements.

### VSecM v0.32.0 (_codename: Norma_)

**Oct 10, 2024 – Nov 06, 2024**

This iteration will be about integrating **VMware Secrets Manager** with
**Google Cloud KMS**.

-   Google Cloud KMS Integration

### VSecM v0.33.0 (_codename: Orion_)

**Nov 07, 2024 – Dec 04, 2024**

This iteration will be about OIDC support, and improving our OpenSSF conformance.

By this point, we expect to have at lest a silver, if not a gold badge from
OpenSSF.

We will also stretch our research into non-Kubernetes deployment options.

-   Get an OpenSSF Silver Badge.
-   Non-Kubernetes deployment options.
-   Other integration options.
-   Creating other official SDKs, along with VSecM Go SDK.
-   Stability improvements.

### VSecM v0.34.0 (_codename: Perseus_)

**Dec 05, 2024 – Jan 01, 2025**

-   Introducing Operators to auto-inject VSecM sidecars and init containers to
    workloads.
-   OIDC support.
-   CertManager integration.

## Closed Iterations

### VSecM v0.21.0 (_codename: Andromeda_)

**Aug 15, 2023 – Sep, 11, 2023**

This was a stability-focused release. We focused on fixing bugs, improving
stability, and improving workflows and CI/CD pipelines. We also created
missing documentation and generated new video tutorials that feature the current
version of **VMware Secrets Manager**.

[Check out the release notes](/docs/changelog/) to learn more about what has
been added, changed, and fixed in this release.
