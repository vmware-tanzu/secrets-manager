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

title = "ADR-0018: Use Asciinema for Use Case Examples in the Documentation"
weight = 18
+++

- Status: accepted
- Date: 2024-09-02
- Tags: documentation, content, asciinema, community

## Context and Problem Statement

The VMware Secrets Manager documentation is a living document that evolves with 
the product. Part of that documentation includes video tutorials and walkthroughs.

The issue with the video tutorials is they are not easily editable. If a change
is made to the product, the video tutorial must be re-recorded. This is not
scalable and can lead to outdated content.

An alternative to video tutorials is using [Asciinema]. Asciinema is a free and
open-source solution for recording terminal sessions and sharing them on the web.
It is lightweight, easy to use, and can be embedded in the documentation.

Of course there will be cases that require video tutorials, but for most use cases,
Asciinema is a better solution.

[Asciinema]: https://asciinema.org/

## Decision Drivers

- Ease of updating documentation
- Maintainability of content
- User experience for readers
- Accessibility of content
- Consistency across documentation

## Considered Options

- Continue with video tutorials
- Use Asciinema for terminal-based examples
- Use static code snippets and screenshots

## Decision Outcome

Chosen option: "Use Asciinema for terminal-based examples", because it provides 
a good balance between interactivity and maintainability.

### Positive Consequences

* Easier to update and maintain documentation
* Provides an interactive experience for users
* Reduces the need for large video files
* Improves accessibility as text can be copied from Asciinema recordings
* Consistent look and feel across terminal-based examples

### Negative Consequences

* May not be suitable for all types of demonstrations (e.g., GUI interactions)
* Requires JavaScript to be enabled for full functionality

## Pros and Cons of the Options

### Continue with video tutorials

* Good, because it provides a visual representation of the process
* Good, because it's familiar to many users
* Bad, because it's difficult to update and maintain
* Bad, because it can lead to outdated content
* Bad, because video files are large and may slow down page loading

### Use Asciinema for terminal-based examples

* Good, because it's easy to update and maintain
* Good, because it provides an interactive experience
* Good, because it's lightweight and fast to load
* Good, because it allows text to be copied from the recording
* Bad, because it may not be suitable for all types of demonstrations

### Use static code snippets and screenshots

* Good, because it's simple to implement
* Good, because it's easy to update
* Bad, because it lacks interactivity
* Bad, because it may not provide enough context for complex operations

## Implementation

1. Gradually replace existing video tutorials with Asciinema recordings where 
   appropriate
2. We can still use video tutorials where Asciinema is not suitable.

## References

* [Asciinema website](https://asciinema.org/)
* [Asciinema embedding documentation](https://asciinema.org/docs/embedding)

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}
