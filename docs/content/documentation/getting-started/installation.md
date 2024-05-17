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
weight = 10
+++

## Clone the Repository

Let's start by cloning the **VMware Secrets Manager** repository first:

```bash
cd $WORKSPACE
git clone https://github.com/vmware-tanzu/secrets-manager.git
cd secrets-manager
```

## Initialize Minikube

Next, let's initialize *Minikube*:

```bash
cd $WORKSPACE/secrets-manager

# If you have a previous minikube setup, you can delete it with:
# make k8s-delete

make k8s-start
```

## Configure Docker Environment Variables

Next, let's configure the Docker environment variables for Minikube.

We don't strictly need this for the quickstart, but it's a good idea to do
it anyway.

```bash
eval $(minikube -p minikube docker-env)
```

## Install SPIRE and VMware Secrets Manager

Next we'll install [SPIRE][spire] and **VMware Secrets Manager** on the cluster.

```bash
helm repo add vsecm https://vmware-tanzu.github.io/secrets-manager/
helm repo update
helm install vsecm vsecm/vsecm
```

This will take a few minutes to complete.

[spire]: https://spiffe.io/spire/ "SPIRE"

{{ edit() }}

