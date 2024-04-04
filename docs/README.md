---
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title: VMware Secrets Manager for Cloud-Native Apps
layout: post
prev_url: /docs/showcase/
permalink: /
next_url: /docs/community/
---

<p class="badges"><a href="https://www.bestpractices.dev/projects/7793"><img src="https://www.bestpractices.dev/projects/7793/badge" alt="OpenSSF Best Practices"></a>
<a href="https://goreportcard.com/report/github.com/vmware-tanzu/secrets-manager"><img src="https://goreportcard.com/badge/github.com/vmware-tanzu/secrets-manager" alt="Go Report Card"></a>
<a href="https://artifacthub.io/packages/helm/vsecm/vsecm"><img src="https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/vsecm" alt="Artifact Hub"></a>
<a href="https://twitch.tv/vadidekivolkan"><img src="https://img.shields.io/twitch/status/vadidekivolkan" alt="Twitch"></a>
<a href="https://join.slack.com/t/a-101-103-105-s/shared_invite/zt-287dbddk7-GCX495NK~FwO3bh_DAMAtQ"><img src="https://img.shields.io/badge/slack-vsecm-brightgreen.svg?logo=slack" alt="Slack"></a>
<a href="https://github.com/vmware-tanzu/secrets-manager/releases"><img src="https://img.shields.io/github/v/release/vmware-tanzu/secrets-manager?color=blueviolet" alt="Version"></a>
<a href="https://raw.githack.com/wiki/vmware-tanzu/secrets-manager/coverage.html)"><img src="https://github.com/vmware-tanzu/secrets-manager/wiki/coverage.svg" alt="Coverage"></a>
<a href="https://github.com/vmware-tanzu/secrets-manager/graphs/contributors"><img src="https://img.shields.io/github/contributors/vmware-tanzu/secrets-manager.svg?color=orange" alt="Contributors"></a>
<a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/LICENSE"><img src="https://img.shields.io/github/license/vmware-tanzu/secrets-manager" alt="License"></a>
<a href="https://github.com/Everduin94/better-commits" ><img src="https://img.shields.io/badge/better--commits-enabled?style=for-the-badge&logo=git&color=a6e3a1&logoColor=D9E0EE&labelColor=302D41" alt="Using Better Commits"></a></p>

<hr style="margin-bottom:-2em;">

## Keep Your Secrets... Secret

**VSecM** redefines secrets management
with intuitive security for cloud native apps.

If you want to get started quickly, check out the [**quickstart tutorial**](/docs/quickstart).

If you're looking for a specific use case, check out the examples below:

* [Using **VSecM Sidecar**](/docs/use-case-sidecar)
* [Using **VSecM SDK**](/docs/use-case-sdk)
* [Using **VSecM Init Container**](/docs/use-case-init-container)
* [Encrypting Secrets](/docs/use-case-encryption)
* [Transforming Secrets](/docs/use-case-transformation)
* [Mounting Secrets as Volumes](/docs/secrets-as-volumes)
* [Mutating a Template File](/docs/use-case-in-memory-template/)
* [Retrieving Secrets from **VSecM SDK**](/docs/use-case-sdk-retrieves-secrets/)
* [Mounting Kubernetes Secrets](/docs/use-case-mounting-secrets-as-env-vars/)
* [Encrypting Large Files](/docs/use-case-encrypting-large-files/)
* [**VMware Secrets Manager** Showcase](/docs/showcase)

<hr style="margin-bottom:-2em;">

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

[endorsements]: /docs/endorsements/ "Endorsements"

## **VMware Secrets Manager** in Action

[Here is a playlist of videos showcasing **VMware Secrets Manager**][videos].

[videos]: /docs/showcase/ "Showcase"

## Where Can I Use **VMware Secrets Manager**?

**VMware Secrets Manager** is perfect for securely storing arbitrary 
configuration information at a central location and securely dispatching it to 
workloads, offering *centralized* and *secure* secrets store for your clusters.

> **VSecM is Perfect for the Edge**
> 
> The *ease of configuration* and *small footprint* make **VMware Secrets Manager** 
> perfect not only for standard deployments but also for **edge deployments** where 
> resources are limited and efficiency is crucial.
{: .block-tip }

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

<p class="github-button">
    <a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/README.md">
        Suggest edits ✏️ 
    </a>
</p>