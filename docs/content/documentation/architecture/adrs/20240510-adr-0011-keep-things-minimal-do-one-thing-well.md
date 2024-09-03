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

title = "ADR-0011: Keep Things Minimal: Do One Thing Well"
weight = 11
+++

# Keep Things Minimal: Do One Thing Well

- Status: accepted
- Date: 2024-05-10
- Tags: design, philosophy

## Context and Problem Statement

At a 5000-feet level, VMware Secrets Manager is a secure Key-Value store.

It can securely store arbitrary values that you, as an administrator, associate 
with keys. It does that, and it does that well.

Any other feature that we add to the project should be in service of this core.
And any additional feature introduced should be evaluated with utmost scrutiny.

We do not add new features just because they are cool. Nor do we add new
features to reach feature parity with other projects.

## Decision Drivers

- We want to keep the project simple and easy to understand.
- We want to keep the project attack surface minimal.
- We want to limit the scope of the project based on the core team and the 
  community's bandwidth.
- We believe that a project that does one thing well is better than a project 
  that does many things poorly.
- We also believe that other tools can complement VMware Secrets Manager, 
  and we do not need to add all features in this project. For example,
  instead of creating an internal policy engine, we can rely on OPA. Or
  instead of creating a new identity control plane, we can rely on SPIFFE.

## Considered Options

1. Say yes to every feature request.
2. Prioritize features based on the core team's bandwidth, the community's 
   bandwidth, and the project's vision.

## Decision Outcome

Chosen option: "option 2", because it is common sense.

### Positive Consequences

- The project will remain simple and easy to understand.
- The project will have a minimal attack surface.
- The project will be able to focus on doing one thing well.
- The project will be able to leverage other tools in the ecosystem.
- The project will be able to keep the scope in check.
- The project will be able to deliver value to the users.

### Negative Consequences

- The project may not be able to cater to every feature request.
- The project may not be able to reach feature parity with other projects.
- The project may not be able to keep everyone happy.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}
