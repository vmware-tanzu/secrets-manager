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

title = "ADR-0008: Have a Minimal and Intuitive API"
weight = 8
+++

- Status: accepted
- Date: 2024-05-11
- Tags: api, usability, interface

## Context and Problem Statement

At its essence, **VMware Secrets Manager** is an in-memory key-value store that
is designed to store sensitive data securely. Since what **VSecM** is fairly
intuitive, the API should be as well. We don't want to bloat the API with
unnecessary complexity or features that are not essential to the core 
functionality. 

The same reasoning also applies for **VSecM SDK**s and **VSecM CLI**, as well,
since they are a kind of API to the **VMware Secrets Manager**.

The API of **VMware Secrets Manager** should be minimal and intuitive to use, 
allowing developers to interact with the service without needing to refer to 
extensive documentation.

## Decision Drivers

- Simplicity and ease of use are key to adoption.
- A minimal API reduces the cognitive load on developers.
- An intuitive API improves developer productivity.
- A minimal API is easier to maintain and evolve.
- A minimal API is less prone to bugs and errors.
- A minimal API is easier to test and validate.
- A minimal API is easier to document and support.
- A minimal API is more likely to be consistent and coherent.
- A minimal API is more likely to be performant and efficient.
- A minimal API is more likely to be secure and reliable.
- A minimal API is more likely to be extensible and scalable.
- A minimal API is more likely to be future-proof and adaptable.
- A minimal API is more likely to be user-friendly and accessible.
- A minimal API is more likely to be well-received and appreciated.
- A minimal API is more likely to be successful and sustainable.

## Considered Options

1. Have a minimal and intuitive API.
2. Have a feature-rich and complex API.
3. Have a complex API but provide extensive documentation and examples.

## Decision Outcome

Chosen option: "option 1", because a minimal and intuitive API aligns with the
core principles of **VMware Secrets Manager** and will help drive adoption and
success. In addition, a minimal API will require less maintenance, be easier to
test, and be more likely to meet the needs of developers without overwhelming
them with unnecessary complexity. This would help the core team and contributors
to focus on the core functionality and quality of the service.

### Positive Consequences

- Improved developer experience.
- Faster adoption and integration.
- Reduced cognitive load on developers.
- Easier to maintain and evolve.
- Easier to test and validate.
- Easier to document and support.

### Negative Consequences

- Potential limitations in functionality.
- Potential reduced flexibility for advanced use cases.
- Potential resistance from users who prefer a feature-rich API.
- Potential need for additional features in the future that may conflict with
  the need for a minimal API.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}

