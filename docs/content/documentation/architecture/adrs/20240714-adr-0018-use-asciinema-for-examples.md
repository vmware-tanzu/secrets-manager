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

- Status: draft
- Date: 2024-07-14
- Tags: documentation, content, asciinema, community

## Context and Problem Statement

The VMware Secrets Manager documentation is a living document that evolves with 
the product. Part of that documentation includes video tutorials and walkthroughs.

The issue with the video tutorials is they are not easily editable. If a change
is made to the product, the video tutorial must be re-recorded. This is not
scalable and can lead to outdated content.

An alternative to video tutorials is using Asciinema. Asciinema is a free and
open-source solution for recording terminal sessions and sharing them on the web.
It is lightweight, easy to use, and can be embedded in the documentation.

Of course there will be cases that require video tutorials, but for most use cases,
Asciinema is a better solution.

