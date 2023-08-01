---
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title: VSecM Design Philosophy
layout: post
permalink: /docs/philosophy/
prev_url: /docs/quickstart/
next_url: /docs/architecture/
---

## VMware Secrets Manager for Cloud Native Apps

## Introduction

We follow the **guidelines** outlined in the next few sections as an
architectural baseline.

## Do One Thing Well

At a 5000-feet level, **VMware Secrets Manager** is a secure Key-Value store.

It can securely store arbitrary values that you, as an administrator, associate
with keys. It does that, and it does that well.

If you are searching for a solution to create and store X.509 certificates,
create dynamic secrets, automate your PKI infrastructure, federate your
identities, use as an OTP generator, policy manager, in short, anything other
than a secure key-value store, then VMware Secrets Manager is likely not the solution you are
looking for.

## Be Kubernetes-Native

**VMware Secrets Manager** is designed to run on Kubernetes and **only** on Kubernetes.

That helps us leverage Kubernetes concepts like *Operators*, *Custom Resources*,
and *Controllers* to our advantage to simplify workflow and state management.

If you are looking for a solution that runs outside Kubernetes or as a
standalone binary, then VMware Secrets Manager is not the Secrets Store you’re looking for.

## Have a Minimal and Intuitive API

As an administrator, there is a limited set of API endpoints that you can
interact with **VMware Secrets Manager**. This makes **VMware Secrets Manager** easy to manage.

In addition, a minimal set of APIs means a smaller attack surface, a smaller
footprint, and a codebase that is easy to understand, test, audit, and
develop; all good things.

## Be Practically Secure

Corollary: Do not be annoyingly secure.

Provide a delightful user experience while taking security seriously.

**VMware Secrets Manager** is a secure solution, yet still delightful to operate.

You won’t have to jump over the hoops or wake up in the middle of the night
to keep it up and running. Instead, **VMware Secrets Manager** will work seamlessly, as if it
doesn’t exist at all.

## Secure By Default

**VMware Secrets Manager** stores your sensitive data **in memory**. None of your secrets
are stored as plain text on disk. Anything that **VMware Secrets Manager** saves to disk
is encrypted.

Yes, that brings up resource limitations since you cannot store a gorilla holding
a banana and the entire jungle in your store; however, a couple of gigabytes of
RAM can store a lot of plain text secrets, so it’s good enough for most
practical purposes.

More importantly, almost all modern instruction set architectures and
operating systems implement [*memory protection*][memory-protection]. The primary
purpose of memory protection is to prevent a process from accessing memory that
has not been allocated to it. This prevents a bug or malware within a process
from affecting other processes or the operating system itself.

[memory-protection]: https://en.wikipedia.org/wiki/Memory_protection "Memory Protection (Wikipedia)"

Therefore, reading a variable’s value from a process’s memory is practically
impossible unless you attach a debugger to it. And that makes keeping
plain text secrets in memory (*and nowhere else than memory*) crucial.

For disaster recovery, VMware Secrets Manager (*by default*) backs up encrypted version of
its state on the file system; however, the
plain text secrets that **VMware Secrets Manager Safe** dispatches to
workloads will always be stored in memory.

## Resilient By Default

When an **VMware Secrets Manager** component crashes or when an **VMware Secrets Manager** component is evicted,
the workloads can still function with the existing secrets they have without
having to rely on the existence of an active secrets store.

When an **VMware Secrets Manager** component restarts, it seamlessly recovers its state from an
encrypted backup without requiring manual intervention.

## Conclusion

While **VMware Secrets Manager** might not have all the bells and whistles of a full-blown
security suite that integrates everything but the kitchen sink, we are sure
that by following these guiding principles, it will remain a
**delightfully secure** secrets manager that a lot of people will love to use.