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

title: What is VSecM?
layout: post
prev_url: /docs/navigation/
permalink: /docs/about/
next_url: /docs/endorsements/
---

## Why Do We Need a Cloud-Native Secrets Manager?

Before we begin, let's think about why we need a cloud-native secrets manager in 
the first place. Can't we just use Kubernetes `Secret`s? The following section 
will answer this question.

## Wait, What's Wrong With Kubernetes `Secret`s?

Kubernetes `Secret`s have legitimate use cases; however,
the out-of-the-box security provided by Kubernetes `Secret`s might not always
meet the stringent security and flexibility demands of modern applications.

In the Kubernetes ecosystem, the handling of secrets is facilitated through a
specialized resource known as a `Secret`. The `Secret` resource allows Kubernetes
to manage and store key-value pairs of sensitive data within a designated
namespace in the cluster.

Kubernetes `Secrets` can be widespread across the cluster into various namespaces
which makes the management and access control to them tricky. 

In addition, when you update a Kubernetes `Secret` it is hard to make the 
workloads be aware of the change. 

Moreover, due to namespace isolation, you cannot define a cluster-wide or 
cross-cluster-federated secrets: You have to tie your secrets to a single 
namespace, which, at times, can be limiting.

All of these (*and more*) is possible with **VMware Secrets Manager**.

## The **VMware Secrets Manager** Difference

Cloud-native secret management can be more secure, centralized, and easy-to-use.
This is where **VMware Secrets Manager**, comes into play:

> **VMware Secrets Manager** offers a *secure*, *user-friendly*, *cloud-native* 
> secrets store that's robust yet *lightweight*, requiring almost zero DevOps 
> skills for installation and maintenance.

In addition, **VMware Secrets Manager**...

* Has the ability to change secrets dynamically at runtime without having to
  reboot your workloads,
* Keeps encrypted backups of your secrets,
* Records last creation and last update timestamps for your secrets,
* Has a version history for your secrets,
* Stores backups of your secrets encrypted at rest,
* Enables GoLang transformations on your secrets,
* Can interpolate your stored secrets onto Kubernetes `Secret`s,
* Enables federation of your secrets across namespaces and clusters,
* and more.

These are not achievable by using Kubernetes `Secret`s only.

## Where **NOT** to Use VMware Secrets Manager

VMware Secrets Manager is **not** a Database, nor is it a distributed caching 
layer. Of course, you may tweak it to act like one if you try hard enough, yet, 
that is generally not a good use of the tool.

VMware Secrets Manager is suitable for storing secrets and dispatching them; 
however, it is a *terrible* idea to use it as a centralized database to store 
everything but the kitchen sink.

Use **VMware Secrets Manager** to store service keys, database credentials, 
access tokens, etc.

## How Do I Get the Root Token? Where Do I Store It?

Unlike other "*vault*"-style secrets stores, **VMware Secrets Manager** requires 
no admin token for operation--a clear advantage that lets your Ops team 
`#sleepmore` due to automation and eliminates manual unlocking after system 
crashes.

However, there's no free lunch, and as the operator of a production system,
your homework is to secure access to your cluster. [Check out the **Production
Deployment Guidelines**][production] for further instructions about hardening
your cluster to securely use **VMware Secrets Manager**.

> **Still Want Your Root Tokens?**
> 
> Although **VMware Secrets Manager** does not require a root token, you can
> still provide one if you want to. Though, when you do that, you will have
> to manually unlock the system after a crash.
> 
> If you let **VMware Secrets Manager** generate the root token for you, you
> will not have to worry about this, and when the system crashes, it will
> automatically unlock itself, so you can `#sleepmore`.
{: .block-tip }

[production]: /docs/production
[gitbook-theme]: https://github.com/sighingnow/jekyll-gitbook "Jekyll GitBook Theme"
[github]: https://github.com/vmware-tanzu/secrets-manager/tree/main/docs "VMware Secrets Manager Documentation on GitHub"

<p class="github-button">
  <a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0021-about.md">
    Suggest edits ✏️ 
  </a>
</p>