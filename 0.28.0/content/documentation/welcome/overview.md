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

title = "Introducing VSecM"
weight = 10
+++

{{ badges() }}

## Keep Your Secrets... Secret

**VMware Secrets Manager** (*VSecM*) redefines secrets management for 
cloud native apps.

If you want to get started quickly, check out the [**quickstart tutorial**](@/documentation/getting-started/overview.md).

If you're looking for a specific use case, check out the examples below:

{{ use_cases() }}

> **A Tale of Two Secrets**
> 
> There are two kinds of "*secret*"s mentioned throughout this documentation:
> 
> * Secrets that are stored in **VSecM Safe**: When discussing these, they will
>   be used like a regular word "secret" or, emphasized "**secret**"; however,
>   you will never see them in `monotype text`.
> * The other kind of secret is Kubernetes `Secret` objects. Those types
>   will be explicitly mentioned as "*Kubernetes `Secret`s*" in the documentation.
> 
> We hope this distinction helps you navigate the documentation more easily.

## **VMware Secrets Manager** in Action

[Here is a playlist of videos showcasing **VMware Secrets Manager**][videos].

[videos]: @/showcase/vsecm.md "Showcase"

## Delightfully Secure Secrets Management

**VMware Secrets Manager** is a secure store for secrets management. It provides
a minimal and intuitive API, ensuring practical security without compromising user
experience.

[Endorsed by industry experts][endorsements], **VMware Secrets Manager** is a
ground-up re-imagination of secrets management, leveraging [**SPIRE**][spire]
for authentication and providing a cloud-native way to manage secrets end-to-end.

**VMware Secrets Manager** is resilient and secure by default, storing sensitive
data in memory and encrypting any data saved to disk.

With **VMware Secrets Manager**, you can rest assured that your sensitive data is
always secure and protected.

[endorsements]: @/community/endorsements.md "Endorsements"

## Where Can I Use **VMware Secrets Manager**?

**VMware Secrets Manager** is perfect for securely storing arbitrary
configuration information at a central location and securely dispatching it to
workloads, offering *centralized* and *secure* secrets store for your clusters.

> **VSecM is Perfect for the Edge**
>
> The *ease of configuration* and *small footprint* make **VMware Secrets Manager**
> perfect not only for standard deployments but also for **edge deployments** where
> resources are limited and efficiency is crucial.


By leveraging Kubernetes security primitives, [**SPIRE**][spire], and strong,
industry-standard encryption, **VMware Secrets Manager** ensures that your
secrets are **only** accessible to **trusted** and **authorized** workloads.
**VMware Secrets Manager**'s Cloud Native--secure by default--foundation helps
you safeguard your business and protect against data breaches.

[Check out **VMware Secret Manager**'s *GitHub* for details][vsecm-github].

[spire]: https://spiffe.io/spire
[vsecm-github]: https://github.com/vmware-tanzu/secrets-manager

If you haven't watched them yet, now might be a good time 🙂.

## Thanks ❤️

Hope you enjoy using **VMware Secrets Manager** as much as we do and find it
helpful in making your ops teams `#sleepmore`.

Browse the rest of this website to learn more about **VSecM**.

May the source be with you 🦄.

{{ edit() }}
