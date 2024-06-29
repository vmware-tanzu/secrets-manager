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

title = "ADR-0003: VSecM Will Be Scoped to Work on Kubernetes Only"
weight = 3
+++

{{
/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/
}}

- Status: accepted
- Date: 2024-05-09
- Tags: integration

## Context and Problem Statement

**VMware Secrets Manager** leverages [SPIFFE][spiffe] as its identity control plane. 
SPIFFE is platform and infrastructure agnostic; so if we want we can add support for
non-Kubernetes environments too.

[spiffe]: https://spiffe.io/ "SPIFFE"

However, this would mean the project will need to use alternatives to its Kubernetes
tooling (*such as `ClusterSPIFFEID`s, `ServiceAccount`s, Kubernetes RBAC, and similar*)

This will increase the scope of the project a lot.

At least for version 1.0, we shall not be considering a non-Kubernetes solution.

This decision may be revisited when we reach 1.0 and project gains adequate maturity,
and there are not many major features to implement.

## Decision Drivers 

- Increase in scope and complexity
- Increase in the attack surface
- Increase in unit and integration testing needs

## Considered Options

1. No Kubernetes work until version 1.0
2. Provide limited support, offering a subset of features.
3. Plan for non-Kubernetes support anyway.
4. Create an experimental branch and work on it without any commitments.

## Decision Outcome

Chosen option: Option 1, because of increased scope not matching our limited time 
and resources; and also because we’d rather keep the project secure and 
well-tested.

### Positive Consequences

- Less scope means more focus.
- Less complexity in the project.
- The Kubernetes machinery can do a lot of the heavy-lifting.

### Negative Consequences

- The project will not solve secrets management need of those who will not use 
  Kubernetes.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}
