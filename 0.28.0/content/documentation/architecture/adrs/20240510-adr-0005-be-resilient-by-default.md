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

title = "ADR-0005: Be Resilient by Default"
weight = 5
+++

- Status: accepted 
- Date: 2024-05-11
- Tags: quality, stability

## Context and Problem Statement

When an VMware Secrets Manager component crashes or when an VMware Secrets Manager 
component is evicted, the workloads can still function with the existing secrets 
they have without having to rely on the existence of an active secrets store.

When an VMware Secrets Manager component restarts, it seamlessly recovers its 
state from an encrypted backup without requiring manual intervention.

## Decision Drivers

- Resilience is also related to [being practically secure][practically-secure]
- A resilient system is easy to operate, maintain, and troubleshoot.
- To have a highly-available system, we need to be resilient first.

[practically-secure]: @/documentation/architecture/adrs/20240510-adr-0004-be-practically-secure.md

## Considered Options

1. Be resilient by default.
2. Think about resilience only when we have time.

## Decision Outcome

Chosen option: "option 1", because we cannot afford to have a system that
is not resilient.

### Positive Consequences

- DevOps will sleep more.

### Negative Consequences

- Additional work upfront.
- Additional complexity in the system.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}

