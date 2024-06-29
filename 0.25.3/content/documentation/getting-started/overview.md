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

title = "Prerequisites"
weight = 5
+++

> **Get Your Hands Dirty**
>
> This is a quickstart guide to get you up and running with **VSecM**.

## Prerequisites

* [**Minikube**][minikube]: You can install **VMware Secrets Manager** on any
  Kubernetes cluster, but we'll use *Minikube* in this quickstart example.
  Minikube is a tool that makes it easy to run Kubernetes locally.
* [**make**][make]: You'll need `make` to run certain build tasks. You can
  install `make` using your favorite package manager.
* [**Docker**][docker]: This quickstart guide assumes that **Minikube** uses
  the *Docker* driver. If you use a different driver, things will still likely
  work, but you might need to tweak some of the commands and configuration.
* [**Helm**][helm]: Helm is a package manager for Kubernetes. It allows developers 
  and operators to more easily package, configure, and deploy applications and 
  services onto Kubernetes clusters. In this quickstart guide, it is expected 
  that you have `helm` installed on your system. Helm will be used to manage 
  the deployment of the **VMware Secrets Manager** onto your Kubernetes cluster. 
  If you don't have Helm installed, you can follow the instructions on the 
  [official Helm website](https://helm.sh) to install it.

> **I Have a Kubernetes Cluster Already**
>
> If you are already have a cluster and a `kubectl` that you can use on that
> cluster, you won't need Minikube, so you can skip the steps related to
> initializing Minikube and configuring Minikube-related environment variables.

Also, if you are not using minikube, you will not need a local docker instance either.

[minikube]: https://minikube.sigs.k8s.io/docs/ "Minikube"
[make]: https://www.gnu.org/software/make/ "GNU Make"
[docker]: https://www.docker.com "Docker"
[helm]: https://helm.sh "Helm"

## A Video Is Worth A Lot of Words

Here's a video that walks you through the steps in this quickstart guide:

<div style="padding:56.25% 0 0 0;position:relative;"><iframe
src="https://player.vimeo.com/video/849328819?h=46caa595f7&amp;badge=0&amp;autopause=0&amp;player_id=0&amp;app_id=58479"
frameborder="0" allow="autoplay; fullscreen; picture-in-picture" allowfullscreen
style="position:absolute;top:0;left:0;width:100%;height:100%;"
title="VMware Secrets Manager (for Cloud-Native Apps) Quickstart"></iframe></div>
<script src="https://player.vimeo.com/api/player.js"></script>

{{ edit() }}
