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

- Status: draft
- Date: 2024-07-14
- Tags: security, secrets, expiration

## Context and Problem Statement

it is up to the receiver of the secret to validate if its expired, and use a 
mechanism to rotate the secret (similar to how JWT works)

VSecM is a tool to manage secrets. If the end user wants to have the secret
they should get it even if it is expired. The end user shall check for
the expiration timestamp and decide if they want to use it or not, or
rotate it.

Otherwise the least you need would be to synchronize clocks between the
VSecM and the end user, and things will get unnecessarily complicated.

Granted, we can have a secret type "token", which will be guaranteed to
expire and get deleted after a certain time, but this is a different
kind of secret, and a different use case.