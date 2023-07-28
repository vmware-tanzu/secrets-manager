---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, tutorial, secrets, overview
title: Overview and Prerequisites
description: set up your environment
micro_nav: true
page_nav:
  prev:
    content: registering secrets
    url: '/docs/register'
  next:
    content: using <strong>Aegis</strong> sidecar
    url: '/docs/use-case-sidecar'
---

<p style="text-align:right;position:relative;top:-40px;"
><a href="https://github.com/ShieldWorks/aegis-web/blob/main/docs/use-case-overview.md"
style="border-bottom: none;background:#e0e0e0;padding:0.5em;display:inline-block;
border-radius:8px;">
edit this page on <strong>GitHub</strong> ✏️</a></p>

## Introduction

This section provides the necessary prerequisites for the upcoming tutorials, 
as well as a high-level overview of the architecture. This information should 
suffice to get you started and familiar with the basics.

## Prerequisites

To complete the tutorials listed here, you will need the following:

* A **Kubernetes** cluster that you have sufficient admin rights.
* **Aegis** up and running on that cluster.
* [The `shieldworks/aegis` repository][repo] cloned inside a workspace
  folder (such as `/home/WORKSPACE/aegis`)

> **How Do I Set Up Aegis**?
>
> To set up **Aegis**, [follow the instructions in this quickstart guide][quickstart].

[quickstart]: /docs/
[repo]: https://github.com/shieldworks/aegis

## Minikube Instructions

For your Kubernetes cluster, you can use [**minikube**][minikube] for development 
purposes.

To use **minikube**, as your cluster, make sure you have 
[**Docker**][docker] up and running first—while there are other ways, using
Minikube’s Docker driver is the fastest and painless way to get started.

Once you have **Docker** up and running, execute the following script to 
install **minikube**. Note that you will also need [`git`][git] and [`make`][make]
installed on your system.

```bash 
# Switch to your workspace folder (e.g., `~/Desktop/WORKSPACE`).
cd $WORKSPACE
# Clone Aegis repository if you haven’t already done so:
git clone https://github.com/shieldworks/aegis.git
# cd into the cloned project folder
cd aegis
# Test if `make` is working, if it fails, install `make` first
make help 
# Install minikube
make k8s-start
```

> **Can I Use This Other Thing Instead**?
>
> You can of course use other tools such as [microk8s], or [kind], [k38][k3s]
> or even a full-blown managed Kubernetes cluster; however it will be virtually
> impossible to cover all possible tooling and OS combinations. Therefore,
> we’ll only provide instructions for **minikube** in this document.

[minikube]: https://minikube.sigs.k8s.io/docs/ "minikube"
[docker]: https://www.docker.com/ "Docker"
[git]: https://git-scm.com/
[make]: https://www.gnu.org/software/make/
[microk8s]: https://microk8s.io/
[kind]: https://kind.sigs.k8s.io/
[k3s]: https://k3s.io/

## High-Level Overview

Here is a high-level overview of various components that will interact with
each other in the upcoming tutorials:

![High-Level Overview](/assets/actors.jpg "High-Level Overview")

On the above diagram:

* **SPIRE** is the identity provider for all intents and purposes.
* **Aegis Safe** is where secrets are stored.
* **Aegis Sentinel** can be considered a bastion host.
* **Demo Workload** is a typical Kubernetes Pod that needs secrets.

> **Want a Deeper Dive**?
>
> In this tutorial, we cover only the amount of information necessary
> to follow through the steps and make sense of how things tie together
> from a platform operator’s perspective.
>
> [You can check out this “**Aegis** Deep Dive” article][architecture]
> to learn more about these components.

[architecture]: /docs/architecture

The **Demo Workload** fetches secrets from **Aegis Safe**. This is either
indirectly done through a **sidecar** or directly by using
[**Aegis Go SDK**][go-sdk].

Using **Aegis Sentinel**, an admin operator or ar CI/CD pipeline can register
secrets to **Aegis Safe** for the **Demo Workload** to consume.

All the above workload-to-safe and sentinel-to-safe communication are
encrypted through **mTLS** using the **X.509 SVID**s that **SPIRE**
dispatches to all the actors.

[go-sdk]: https://github.com/shieldworks/aegis-sdk-go

After this high-level overview of your system, let’s create a workload next.