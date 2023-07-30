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

title: What is VSecM?
layout: post
next_url: /docs/endorsements/
prev_url: /
permalink: /docs/about/
---

> **Welcome** 🛡
> 
> Welcome to *VMware Secrets Manager* (**VSecM**) for Cloud-Native Apps.️
> 
> **VSecM** keeps your secrets… secret, so you can `#sleepmore`.
{: .block-tip }

## Hey, This Looks Like a GitBook 📖

Yes, this website is intentionally created like a book. We wanted to make
sure that you have a great experience reading our documentation, and we
believe that following the documentation cover-to-cover as if it were a book
is the best way to learn about **VMware Secrets Manager** (*VSecM*) for
Cloud-Native Apps.

We use a heavily customized version of [the Jekyll GitBook Theme][gitbook-theme]
to achieve this. [You can check the source code for this website on
GitHub][github] to see how that is done.

## Terminology: A Tale of Two Secrets

There are two kinds of “*secret*”s mentioned throughout this documentation:

* Secrets that are stored in **Aegis Safe**: When discussing these, they will
  be used like a regular word “secret” or, emphasized “**secret**”; however,
  you will never see them in `monotype text`.
* The other kind of secret is Kubernetes `Secret` objects. Those types
  will be explicitly mentioned as “Kubernetes `Secret`s” in the documentation.

We hope this will clarify any confusion going forward.

## Wait, What’s Wrong With Kubernetes `Secret`s?

Kubernetes `Secret`s have legitimate use cases; however,
the out-of-the-box security provided by Kubernetes `Secret`s might not always
meet the stringent security and flexibility demands of modern applications.

In the Kubernetes ecosystem, the handling of secrets is facilitated through a
specialized resource known as a `Secret`. The `Secret` resource allows Kubernetes
to manage and store key-value pairs of sensitive data within a designated
namespace in the cluster.

Kubernetes `Secrets` can be widespread across the cluster into various namespaces
which makes the management and access control to them tricky. In addition,
when you update a Kubernetes `Secret` it is hard to make the workloads be aware
of the change. Moreover, due to namespace isolation, you cannot define a cluster-wide
or cross-cluster-federated secrets: You have to tie your secrets to a single
namespace, which, at times, can be limiting.

## The **Aegis** Difference

Cloud-native secret management can be more secure, centralized, and easy-to-use.
This is where **Aegis**, comes into play:

> **Aegis** offers a *secure*, *user-friendly*, *cloud-native* secrets store that’s
> robust yet *lightweight*, requiring almost zero DevOps skills for installation
> and maintenance.

In addition, **Aegis**…

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

## Where **NOT** to Use Aegis

**Aegis** is **not** a Database, nor is it a distributed caching layer. Of course,
you may tweak it to act like one if you try hard enough, yet, that is
generally not a good use of the tool.

**Aegis** is suitable for storing secrets and dispatching them; however, it
is a *terrible* idea to use it as a centralized database to store everything
but the kitchen sink.

Use **Aegis** to store service keys, database credentials, access tokens,
etc.

## How Do I Get the Root Token? Where Do I Store It?

Unlike other “*vault*”-style secrets stores, **Aegis** requires no admin token
for operation—a clear advantage that lets your Ops team `#sleepmore` due to
automation and eliminates manual unlocking after system crashes.

However, there’s no free lunch, and as the operator of a production system,
your homework is to secure access to your cluster. [Check out the **Production
Deployment Guidelines**][production] for further instructions about hardening
your cluster to securely use **Aegis**.

[production]: /production

## Installation

First, ensure that you have sufficient administrative rights on your
**Kubernetes** cluster. Then create a workspace folder
(*such as `$HOME/Desktop/WORKSPACE`*) and clone the project.
And finally execute `./hack/deploy.sh` as follows.

```bash 
mkdir $HOME/Desktop/WORKSPACE
export $WORKSPACE=$HOME/Desktop/WORKSPACE

./hack/deploy.sh
```

## Verifying the Installation

To verify installation, check out the `aegis-system` namespace:

```bash
kubectl get deployment -n aegis-system

# Output:
#
# NAME             READY   UP-TO-DATE   AVAILABLE
# aegis-safe       1/1     1            1
# aegis-sentinel   1/1     1            1
```

That’s it. You are all set 🤘.

## Uninstalling Aegis

Uninstallation can be done by running a script:

```bash 
cd $WORKSPACE/aegis 
./hack/uninstall.sh
```

## Next Steps

Since you have **Aegis** up and running, here is a list of topics that you can
explore next:

* [A Deeper Dive into **Aegis** Architecture](/docs/architecture)
* [**Aegis** Design Decisions](/docs/philosophy)
* [How to Register Secrets to A Workload Using **Aegis**](/docs/register)
* [**Aegis Sentinel** CLI Documentation](/docs/sentinel)
* [**Aegis** Go SDK](/docs/sdk)
* [Configuring **Aegis**](/docs/configuration)
* [Local Development](/docs/contributing)
* [**Aegis** Production Deployment Guidelines](/production)

In addition, these topics might pique your interest too:

* [Umm… How Do I Pronounce “**Aegis**”](/pronounciation)?
* [Who’s Behind **Aegis**](/maintainers)?
* [What’s Coming Up Next](/timeline)?
* [Can I See a Change Log](/changelog)?

## We’d Love to Hear Back From You

If you have comments, suggestions, and ideas to share; if you have found
a bug; or if you want to contribute to **Aegis**, these links might be what
you are looking for:

* [I Want to Contribute to **Aegis**](/contact#i-want-to-be-a-contributor)
* [I Have Something to Say](/contact#comments-and-suggestions)
* [Can I Buy You A Coffee](/contact#coffee)?


## VMware Secrets Manager for Cloud-Native Apps

*VMware Secrets Manager for Cloud-Native Apps is a cloud-native secure store
for secrets management. It provides a minimal and intuitive API, ensuring
practical security without compromising user experience.*

[Endorsed by **industry experts**](/docs/endorsements/), **VMware Secrets Manager** 
is a ground-up re-imagination of secrets management, leveraging SPIRE for 
authentication and providing a cloud-native way to manage secrets end-to-end.

**VMware Secrets Manager** is resilient and secure by default, storing sensitive
data in memory and encrypting any data saved to disk.

With **VMware Secrets Manager**, you can rest assured that your sensitive data is
always secure and protected.

## Maintainers

**VMware Secrets Manager** is maintained by [a group of dedicated individuals 
listed in the `CODEOWNERS` file on the **VMware Secrets Manager** GitHub 
repository][codeowners].

[codeowners]: https://github.com/vmware-tanzu/secrets-manager/blob/main/CODEOWNERS "VMware Secrets Manager CODEOWNERS"
[gitbook-theme]: https://github.com/sighingnow/jekyll-gitbook "Jekyll GitBook Theme"
[github]: https://github.com/vmware-tanzu/secrets-manager/tree/main/docs "VMware Secrets Manager Documentation on GitHub"