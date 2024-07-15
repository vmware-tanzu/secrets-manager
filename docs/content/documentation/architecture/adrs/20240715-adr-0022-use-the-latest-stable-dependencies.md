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

title = "ADR-0022: VSecM Shall Use the Latest Stable Dependencies"
weight = 22
+++

- Status: draft
- Date: 2024-07-14
- Tags: security, dependencies

## Context and Problem Statement

We should use the latest stable dependencies to ensure that we are not using
vulnerable dependencies. This includes the GoLang version that we are using
too.