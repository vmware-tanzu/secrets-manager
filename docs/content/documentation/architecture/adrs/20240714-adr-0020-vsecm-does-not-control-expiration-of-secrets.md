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

title = "ADR-0020: VSecM will not enforce expiration and invalid before time for secrets"
weight = 20
+++

- Status: accepted
- Date: 2024-09-02
- Tags: security, secrets, expiration

## Context and Problem Statement

VMware Secrets Manager (VSecM) is designed to manage secrets efficiently. The 
question arises whether VSecM should enforce expiration and invalid before time 
for secrets, or if this responsibility should lie with the secret consumers.

## Decision Drivers

* Simplicity and flexibility of the system
* Avoiding unnecessary complexity in time synchronization
* Allowing users to have full control over their secrets
* Maintaining consistency with common practices in secret management

## Considered Options

* Option 1: VSecM enforces expiration and invalid before time
* Option 2: Secret consumers handle expiration and invalid before time
* Option 3: Implement a hybrid approach with optional enforcement

## Decision Outcome

Chosen option: "Option 2: Secret consumers handle expiration and invalid 
before time", because:

1. It provides more flexibility to the end users.
2. It avoids the need for clock synchronization between VSecM and secret consumers.
3. It aligns with common practices in secret management (e.g., JWT handling).
4. It keeps VSecM's core functionality simple and focused.

## Consequences

Positive:

- Simplifies VSecM's architecture and reduces potential points of failure.
- Gives users full control over how they handle secret expiration.
- Allows for easier integration with various systems and use cases.

Negative:

- Requires secret consumers to implement their own expiration checks.
- May lead to inconsistent handling of expired secrets across different consumers.

## Implementation Details

- VSecM will store expiration and invalid before time as metadata with secrets.
- VSecM will always return secrets when requested, regardless of their expiration 
  status.
- Secret consumers are responsible for checking expiration and invalid before time.
- Consumers should implement their own mechanisms for secret rotation when needed.

## Related Decisions

* For specific use cases requiring strict expiration enforcement, a separate 
  "token" secret type may be considered in the future.

## Notes

It is up to the receiver of the secret to validate if it's expired and use a 
mechanism to rotate the secret (similar to how JWT works). If the end user wants 
to have the secret, they should get it even if it is expired. The end user shall 
check for the expiration timestamp and decide if they want to use it or not, 
or rotate it.

This approach avoids the need to synchronize clocks between VSecM and the end 
user, preventing unnecessary complications.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}
