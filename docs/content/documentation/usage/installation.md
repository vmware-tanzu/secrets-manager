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

To verify installation, check out the `vsecm-system`, `spire-system`, and 
`spire-server` namespaces:

```bash
kubectl get po -n vsecm-system
# Example Output:
# NAME                              READY   STATUS    RESTARTS   AGE
# vsecm-keystone-59fc9568b6-hhnsj   1/1     Running   0          27s
# vsecm-safe-0                      1/1     Running   0          27s
# vsecm-sentinel-6998c5c5d7-lmdfh   1/1     Running   0          27s

kubectl get po -n spire-system
# Example Output:
# NAME                READY   STATUS    RESTARTS      AGE
# spire-agent-ts84q   3/3     Running   2 (56s ago)   58s

kubectl get po -n spire-server
# Example Output:
# NAME             READY   STATUS    RESTARTS   AGE
# spire-server-0   2/2     Running   0          62s
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
