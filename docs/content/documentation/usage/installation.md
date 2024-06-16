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

title = "Installation"
weight = 2
+++

## Introduction

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

## Installing Using <img src="/assets/helm-icon-color.png" alt="helm" width="20"/>`helm`

`helm` is the easiest way to install **VMware Secrets Manager** to
your Kubernetes cluster.

Make sure you have `helm` v3 installed and execute the following commands:

```bash
helm repo add vsecm https://vmware-tanzu.github.io/secrets-manager/
helm repo update
helm install vsecm vsecm/vsecm
```

For detailed instruction on **VMware Secrets Manager** installation
through Helm Charts please refer to VSecM Helm Charts [README.md][README.md]

[README.md]: https://github.com/vmware-tanzu/secrets-manager/blob/main/helm-charts/README.md

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

That's it. You are all set 🤘.

## Verifying the Installation

To verify installation, check out the `vsecm-system` and `spire-system namespaces:

```bash
kubectl get po -n vsecm-system
```

You should see something similar to the following output:

```txt
NAME                             READY   STATUS
vsecm-safe-85dd95949c-f4mhj      1/1     Running
vsecm-sentinel-6dc9b476f-djnq7   1/1     Running
```

Then, do the same for `spire-system` namespace:

```bash
kubectl get po -n spire-system
```

You should see something similar to the following output:

```txt
NAME                           READY   STATUS
spire-agent-p9m27              3/3     Running
spire-server-6fb4f57c8-6s7ns   2/2     Running
```

> **SPIRE Agent and Server Might Restart**
>
> It is okay if you see the SPIRE Agent and Server pods restarting once or twice.
> They will *eventually* stabilize within a few moments.


## Uninstalling VMware Secrets Manager

Uninstallation can be done by running a script:

```bash
cd $WORKSPACE/secrets-manager
./hack/uninstall.sh
```

Or, if you have installed **VMware Secrets Manager** using `helm`, you can
use `make helm-delete`  command:

```bash
# note that using `helm uninstall vsecm` is not recommended as it may
# leave some resources behind in the cluster.
# You are encouraged to use `make helm-uninstall` instead.
make helm-uninstall
```

{{ edit() }}
