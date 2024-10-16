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

title = "ADR-0004: Be Practically Secure"
weight = 4
+++

- Status: accepted 
- Date: 2024-05-11
- Tags: guidelines, principles

## Context and Problem Statement

**Corollary**: Do not be annoyingly secure.

Provide a delightful user experience while taking security seriously.

VMware Secrets Manager is a secure solution, yet still delightful to operate.

You won’t have to jump over the hoops or wake up in the middle of the night to 
keep it up and running. 

Instead, VMware Secrets Manager will work seamlessly, as if it doesn't exist 
at all.

## Decision Drivers

- Certain solutions out in the wild are hard to configure operate and maintain.
- Security solutions are often seen as a hindrance to productivity.
- A security solution can help DevOps, SREs, and security engineers
  instead of being a burden.

## Considered Options

1. Be practically secure, and enable delightful user experience. Provide
   additional levels of security as optional features to those who need them.
2. Be annoyingly secure, maintenance-heavy, and hard to operate.

## Decision Outcome

Chosen option: "option 1", because we want DevOps teams to have a life. 

### Positive Consequences

- Happy users.
- Secure solution.
- Less maintenance.
- More time to focus on other things.
- Ability to harden the solution when needed.

### Negative Consequences

- Those who need a more secure solution may need to enable additional features.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}
