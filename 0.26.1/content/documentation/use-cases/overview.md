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

title = "Overview"
weight = 2
+++

## Introduction

This section provides the necessary prerequisites for the upcoming tutorials,
as well as a high-level overview of the architecture. This information should
suffice to get you started and familiar with the basics.

## Prerequisites

To complete the tutorials listed here, you will need the following:

* A **Kubernetes** cluster that you have sufficient admin rights.
* **VMware Secrets Manager** up and running on that cluster.
* [The `vmware-tanzu/secrets-manager` repository][repo] cloned inside a workspace
  folder (such as `/home/WORKSPACE/secrets-manager`)

> **How Do I Set Up VMware Secrets Manager**?
>
> To set up **VMware Secrets Manager**, [follow the instructions in this quickstart guide][quickstart].


[quickstart]: @/documentation/getting-started/overview.md
[repo]: https://github.com/vmware-tanzu/secrets-manager

## Minikube Instructions

For your Kubernetes cluster, you can use [**minikube**][minikube] for development
purposes.

To use **minikube**, as your cluster, make sure you have
[**Docker**][docker] up and running first--while there are other ways, using
Minikube's Docker driver is the fastest and painless way to get started.

Once you have **Docker** up and running, execute the following script to
install **minikube**. Note that you will also need [`git`][git] and [`make`][make]
installed on your system.

```bash
# Switch to your workspace folder (e.g., `~/Desktop/WORKSPACE`).
cd $WORKSPACE
# Clone VMware Secrets Manager repository if you haven't already done so:
git clone https://github.com/vmware-tanzu/secrets-manager.git
# cd into the cloned project folder
cd secrets-manager
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
> we'll only provide instructions for **Minikube** in this document.


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
* **VSecM Safe** is where secrets are stored.
* **VSecM Sentinel** can be considered a bastion host.
* **Demo Workload** is a typical Kubernetes Pod that needs secrets.

> **Want a Deeper Dive**?
>
> In this tutorial, we cover only the amount of information necessary
> to follow through the steps and make sense of how things tie together
> from a platform operator's perspective.
>
> [You can check out this "**VMware Secrets Manager** Deep Dive" article][architecture]
> to learn more about these components.


[architecture]: @/documentation/architecture/overview.md

The **Demo Workload** fetches secrets from **VSecM Safe**. This is either
indirectly done through a **sidecar** or directly by using
[**VMware Secrets Manager Go SDK**][go-sdk].

Using **VSecM Sentinel**, an admin operator or ar CI/CD pipeline can register
secrets to **VSecM Safe** for the **Demo Workload** to consume.

All the above workload-to-safe and sentinel-to-safe communication are
encrypted through **mTLS** using the **X.509 SVID**s that **SPIRE**
dispatches to all the actors.

[go-sdk]: https://github.com/vmware-tanzu/secrets-manager/tree/main/sdk

## List of Use Cases in This Section

{{ use_cases() }}

{{ edit() }}
