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

title: Release Management
layout: post
next_url: /
prev_url: /
permalink: /docs/release-management/
---

## Configuring Minikube Local Registry for Linux and Mac

Switch to the **Aegis** project folder
Then, delete any existing minikube cluster.

```bash
cd $WORKSPACE/aegis
make k8s-delete
```

Then start the **Minikube** cluster.

```bash 
cd $WORKSPACE/aegis
make k8s-start
```

This will also start the local registry. However, you will need to
eval some environment variables to be able to use Minikube’s registry instead
of the local Docker registry.

```bash 
cd $WORKSPACE/aegis
eval $(minikube docker-env)

echo $DOCKER_HOST
# example: tcp://192.168.49.2:2376
#
# Any non-empty value to `echo $DOCKER_HOST` means that the environment
# has been set up correctly.
```

## Creating a Local Deployment

Follow these steps to build **Aegis** from scratch and deploy it to your
local **Minikube** cluster, to experiment it with your workloads.

```bash
make k8s-delete
make k8s-start
make build-local
make deploy-local
```

When everything completes, you should be able to see **Aegis** pods in
the `aegis-system` namespace.

```bash
kubectl get po -n aegis-system

# Output should list `aegis-safe` and `aegis-sentinel`.
```

## Cutting a Release

Before every release cut, follow the steps outlined below.

### 1. Double Check

Ensure that all the changes have been merge to `main`.

Also make sure your `docker` and `Minikube` are up and running.

Additionally, execute `eval $(minikube -p minikube docker-env)` once more to
update your environment.

Finally, ensure all changes that need to go to a release in all
repositories have been merged to `main`.

### 2. `make help`

Check the `make help` command first, as it includes important information.

``` 
make help

--------------------------------------------------------------------
          🛡️ Aegis: Keep your secrets… secret.
          🛡️ https://aegis.ist

… Truncated …

  Prep/Cleanup:
        ˃ make k8s-delete;make k8s-start;
        ˃ make clean;
--------------------------------------------------------------------
  Testing:
    ⦿ Istanbul images:
        ˃ make build-local;make deploy-local;make test-local;
    ⦿ Photon images:
        ˃ make build-local;make deploy-photon-local;make test-local;
    ⦿ Istanbul (remote) images:
        ˃ make build;make deploy;make test-remote;
    ⦿ Photon (remote) images:
        ˃ make build;make deploy-photon;make test-remote
--------------------------------------------------------------------
  Tagging:
        ˃ make tag;
--------------------------------------------------------------------
  Example Use Cases:
        ˃ make example-sidecar-deploy(-local);
        ˃ make example-sdk-deploy(-local);
        ˃ make example-multiple-secrets-deploy(-local);
--------------------------------------------------------------------
```

### 3. Test Aegis Istanbul Images

**Aegis** Istanbul series use lightweight and secure distroless images.

```bash 
make k8s-delete
make k8s-start
eval $(minikube -p minikube docker-env)

# for macOS, you might need to run this on a separate terminal:
# make mac-tunnel

make build-local
make deploy-local
make test-local
```

If the tests pass, go to the next step.

### 4. Test Aegis Photon (i.e. VMware Photon) Images

**Aegis** Photon series use [**VMware Photon OS**][photon] as their base images.

[photon]: https://vmware.github.io/photon/

```bash 
make k8s-delete
make k8s-start
eval $(minikube -p minikube docker-env)

# for macOS, you might need to run this on a separate terminal:
# make mac-tunnel

make build-local
make deploy-photon-local
make test-local
```

### 5. Tagging

Tagging needs to be done **on the build server**.

There is no automation for this yet.

```bash 
make tag
```

Follow the instructions, and you should be good to go.

### 6. All Set 🎉

You’re all set.

Happy releasing.
