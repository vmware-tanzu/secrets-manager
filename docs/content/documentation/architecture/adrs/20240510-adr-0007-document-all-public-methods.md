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

title = "ADR-0007: Document All Public Methods"
weight = 7
+++

- Status: accepted 
- Date: 2024-05-11
- Tags: documentation, code-quality

## Context and Problem Statement

In software development, documentation plays a crucial role in understanding, 
maintaining, and using code effectively. Public methods, being the interfaces 
through which modules interact, require clear and precise documentation to 
ensure developers and users of the software can understand their purpose, 
usage, and behavior without needing to delve into the underlying implementation.

We will document all public methods in our codebase. 

This documentation will include:

- **Purpose**: A brief description of what the method does.
- **Parameters**: List each parameter, its type, and a description of its role.
- **Returns**: Describe the return type and the semantics of the returned value.
- **Exceptions**: List all exceptions that the method can throw, with a 
  description of the conditions under which each exception is thrown.
- **Examples (optional)**: Provide at least one example of how the method can be 
  used, which can also serve as a snippet for tests.

## Decision Drivers

- Improve code readability and maintainability.
- Facilitate onboarding of new developers.
- Ensure consistent and clear documentation across the codebase.
- Reduce the cognitive load on developers by providing clear and concise 
  documentation.
- Enhance the overall quality of the codebase.
- Meet the expectations of developers and users regarding documentation quality.
- Ensure that the codebase is self-explanatory and easy to understand.
- Reduce the time required to understand the functionality of each method.
- Code alone cannot be the source of truth for the behavior of the software.

## Considered Options

1. Document all public methods.
2. Document only complex or non-obvious public methods.
3. Document only public methods that are part of the public API.
4. Make documentation optional for public methods.

## Decision Outcome

Chosen option: "option 1", because documenting all public methods will create a 
more maintainable and understandable codebase, despite the initial investment 
in time and the need for ongoing maintenance of the documentation.

In addition, this is a simpler approach that ensures we don't have to think
about which methods to document and which to skip. It also sets a clear
standard for documentation quality across the codebase.

### Positive Consequences

- **Enhanced Clarity**: New and existing developers will spend less time 
  understanding the functionality of each method.
- **Increased Maintainability**: Well-documented code is easier to maintain and 
  debug.
- **Better Onboarding**: New team members can onboard faster as they have clear 
  documentation to refer to.

### Negative Consequences 

- **Increased Initial Development Time**: Writing detailed documentation can 
  increase the time required for initial development.
- **Potential for Documentation Decay**: If not maintained, documentation can 
  become outdated as the code evolves, leading to potential misunderstandings.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}

