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

title = "ADR-0006: Be Secure by Default"
weight = 6
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
- Date: 2024-05-11 
- Tags: security, guidelines

## Context and Problem Statement

**VMware Secrets Manager** stores your sensitive data in memory. None of your 
secrets are stored as plain text on disk. Any secret that **VMware Secrets 
Manager** saves to any other medium is encrypted.

Yes, that brings up resource limitations, since the amount of secrets you can
store per **VSecM** instance, is limited by the amount of memory you can allocate
to the **VSecM Safe** Kubernetes Pod. You cannot store a gorilla holding a 
banana and the entire jungle in your store; however, a couple of gigabytes of RAM 
can store a lot of plain text secrets, so it's good enough for most practical 
purposes.

More importantly, almost all modern instruction set architectures and operating 
systems implement [memory protection][memory-protection]. The primary purpose of 
memory protection is to prevent a process from accessing memory that has not 
been allocated to it. This prevents a bug or malware within a process from 
affecting other processes or the operating system itself.

Therefore, reading a variable's value from a process's memory is practically 
impossible unless you attach a debugger to it. And that makes keeping plain text 
secrets in memory (*and nowhere else than memory*) crucial.

For disaster recovery, VMware Secrets Manager (*by default*) backs up encrypted 
version of its state on the file system; however, the plain text secrets that 
**VSecM Safe** dispatches to workloads will always be stored in memory.

The above line of thought should be applied to any and every architectural 
decision we make in **VMware Secrets Manager**. If a decision is not secure by
default, it **MUST** be revisited.

[memory-protection]: https://en.wikipedia.org/wiki/Memory_protection "Linux Memory Protection"

## Decision Drivers <!-- optional -->

- Security is a primary concern.
- Users will be inclined to trust **VMware Secrets Manager** more if it is 
  secure by default.
- People will typically use the defaults, so the defaults should be secure.
- Not everyone reads the documentation, so the defaults should be secure.

## Considered Options

1. Be secure by default.
2. Be secure by configuration.

## Decision Outcome

Chosen option: "option 1", because we want to make it easy for users to trust
**VMware Secrets Manager**. Hope is not a strategy, and we cannot hope that
users will configure **VSecM** securely.

### Positive Consequences <!-- optional -->

- Users will trust **VMware Secrets Manager** more.
- Users will not have to worry about configuring **VSecM** securely.
- Users will not have to worry about the defaults being insecure.

### Negative Consequences <!-- optional -->

- If we make a mistake in the defaults, it will affect everyone.
- Users still shall trust, but verify; and file issues if they find something
  that is not secure by default.
- We will have to be extra careful when making changes to the defaults.
- This approach requires additional effort upfront and ongoing.

<p>&nbsp;</p>
<p>&nbsp;</p>

## ADRs

You can view the ADRs by browsing this following list:

{{ adrs() }}

{{ edit() }}

