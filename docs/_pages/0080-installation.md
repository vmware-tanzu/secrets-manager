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

title: VSecM Installation
layout: post
next_url: /docs/cli/
prev_url: /docs/architecture/
permalink: /docs/installation/
---

There are several ways to install **VMware Secrets Manager** to a Kubernetes cluster:

* You can use **helm charts**,
* Or you can use the `Makefile` targets

This page covers both approaches.

## Prerequisites

Before you start, make sure you have the following prerequisites:

* You have `helm` installed on your system.
* You have `kubectl` installed on your system.
* You have a Kubernetes cluster running and `kubectl` is configured to
  connect to it.
* You have `make` installed on your system.

## Installing Using `helm`

> **🚧 Work In Progress 🚧️**
>
> `helm` installation is a work in progress as we are migrating the
> Helm charts to their new home at `https://vsecm.com/helm-charts`.
> 
> We will update this section with instructions as soon as the charts
> are ready.
>
> Thank you for your patience and understanding 🙏.
{: .block-warning}

`helm` is the easiest way to install **VMware Secrets Manager** to 
your Kubernetes cluster.

Make sure you have the recent version of `helm` installed and
execute the following commands:

```bash 
TBD
```

## Installing Using `make`

Make sure you have `make` and `git` installed in your system.

First, clone the repository:

```bash
cd $WORKSPACE
git clone https://github.com/vmware-tanzu/secrets-manager.git
cd secrets-manager
```

Then, run the following command to install **VMware Secrets Manager** to your 
cluster:

```bash
make deploy
```

That’s it. You are all set 🤘.

## Verifying the Installation

To verify installation, check out the `vsecm-system` and `spire-system namespaces:

```bash
kubectl get po -n vsecm-system
```

You should see something similar to the following output:

```text
NAME                             READY   STATUS
vsecm-safe-85dd95949c-f4mhj      1/1     Running
vsecm-sentinel-6dc9b476f-djnq7   1/1     Running
```

Then, do the same for `spire-system` namespace:

```bash
kubectl get po -n spire-system
```

You should see something similar to the following output:

```text
NAME                           READY   STATUS
spire-agent-p9m27              3/3     Running
spire-server-6fb4f57c8-6s7ns   2/2     Running
```

> **SPIRE Agent and Server Might Restart**
> 
> It is okay if you see the SPIRE Agent and Server pods restarting once or twice.
> They will *eventually* stabilize within a few moments.
{: .block-tip}

## Uninstalling VMware Secrets Manager

Uninstallation can be done by running a script:

```bash 
cd $WORKSPACE/secrets-manager
./hack/uninstall.sh
```

Or, if you have installed **VMware Secrets Manager** using `helm`, you can 
use `helm uninstall`:

```bash
# Find the release name by running `helm list`; 
# then, uninstall it:
helm uninstall vsecm
```
