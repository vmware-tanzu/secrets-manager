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
prev_url: /
permalink: /docs/release-management/
next_url: /docs/governance/
---

<p class="github-button"
><a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0980-release-management.md"
>edit this page on <strong>GitHub</strong> ✏️</a></p>

## Introduction

This page discusses the release management process for **VMware Secrets Manager**.

If you are responsible for cutting a release, please follow the steps outlined
here.

## VMware Secrets Manager Build Server

> **The VSecM Build Server Contains Trust Material**
>
> The **VSecM** build server is a hardened and trusted environment
> with limited access. It contains trust material such as the
> Docker Content Trust root key, and the private key for signing
> the **VSecM** images.

We (*still*) have a manual build process, so you will need access to the
**VSecM** build server to be able to cut a release.

You can of course build **VSecM** locally, but without the build server, you
won’t be able to push the images to the registry and tag the release.

## Make Sure We Are Ready for a Release Cut

Check out [this internal link][release] to see if there is any outstanding
issues for the release. If they can be closed, close them. If they cannot
be closed, move them to the next version.

## Make Sure You Update the Release Notes

* Add any publicly-known vulnerabilities that are fixed in this release.
* Add any significant changes completed to the release notes.

[release]: https://github.com/orgs/vmware-tanzu/projects/70/views/1

## Configuring Minikube Local Registry

Switch to the `$WORKSPACE/secrets-manager` project folder
Then, delete any existing minikube cluster.

```bash
cd $WORKSPACE/secrets-manager
make k8s-delete
```

Then start the **Minikube** cluster.

```bash
cd $WORKSPACE/secrets-manager
make k8s-start
```

This will also start the local registry. However, you will need to
eval some environment variables to be able to use Minikube’s registry instead
of the local Docker registry.

```bash
cd $WORKSPACE/secrets-manager
eval $(minikube docker-env)

echo $DOCKER_HOST
# example: tcp://192.168.49.2:2376
#
# Any non-empty value to `echo $DOCKER_HOST` means that 
# the environment has been set up correctly.
```

## Creating a Local Deployment

Follow these steps to build **VSecM** from scratch and deploy it to your
local **Minikube** cluster, to experiment it with your workloads.

```bash
# Temporarily disable Docker Content Trust 
# to deploy Minikube:
export DOCKER_CONTENT_TRUST=0

make k8s-delete
make k8s-start

# The environment has changed; re-evaluate 
# the environment variables:
eval $(minikube docker-env)

make build-local
make deploy-local
```

When everything completes, you should be able to see **VMware Secrets Manager**
pods in the `vsecm-system` namespace.

```bash
kubectl get po -n vsecm-system

# Output should list `vsecm-safe` and `vsecm-sentinel`.
```

## Cutting a Release

Before every release cut, follow the steps outlined below.

### 0. Are you on a release branch?

Make sure you are on a release branch, forked off of the most recent `main` branch.

Also ensure that all changes have been merged to `main`.

### 1. Check Docker and Minikube

Also make sure your `docker` and `Minikube` are up and running.

Additionally, execute `eval $(minikube -p minikube docker-env)` once more to
update your environment.

### 2. `make help`

Check the `make help` command first, as it includes important information.

### 3. Test VSecM Distroless Images

**VMware Secrets Manager** Distroless series use lightweight and secure
distroless images.

```bash
make k8s-delete
make k8s-start
eval $(minikube -p minikube docker-env)

# For macOS, you might need to run `make mac-tunnel` 
# on a separate terminal.
# For other Linuxes, you might not need it.
#
# make mac-tunnel

make build-local
make deploy-local
make test-local
```

If the tests pass, go to the next step.

### 4. Test VSecM Photon (i.e. VMware Photon) Images

**VMware Secrets Manager** Photon series use [**VMware Photon OS**][photon] as
their base images.

[photon]: https://vmware.github.io/photon/

```bash
make k8s-delete
make k8s-start
eval $(minikube -p minikube docker-env)

# For macOS, you might need to run `make mac-tunnel` 
# on a separate terminal.
# For other Linuxes, you might not need it.
#
# make mac-tunnel

make build-local
make deploy-photon-local
make test-local
```

### 5. Test VSecM Distroless FIPS Images

```bash
make k8s-delete
make k8s-start
eval $(minikube -p minikube docker-env)

# For macOS, you might need to run `make mac-tunnel` 
# on a separate terminal.
# For other Linuxes, you might not need it.
#
# make mac-tunnel

make build-local
make deploy-fips-local
make test-local
```

### 6. Test VSecM Photon FIPS Images

```bash
make k8s-delete
make k8s-start
eval $(minikube -p minikube docker-env)

# For macOS, you might need to run `make mac-tunnel` 
# on a separate terminal.
# For other Linuxes, you might not need it.
#
# make mac-tunnel

make build-local
make deploy-photon-fips-local
make test-local
```

### 7. Tagging

Tagging needs to be done **on the build server**.

There is no automation for this yet.

> **Don’t forget to Bump the Version**
>
> If you are cutting a new release, do not forget to bump the version,
> before running the tagging script below.
{: .block-tip }

```bash
git checkout main
git stash
git pull
export DOCKER_CONTENT_TRUST=1
make build
make tag
```

### 8. Create a Release PR

Since we successfully published a release to DockerHub, and ArtifactHub, we
now can merge our changes to the `main` branch.

Create a PR from the release branch you are on, and follow the regular merge
approval process.

### 9. Release helm-charts

We offer the [release_helm_chart.sh][release_script] script for your use.
To execute the script, provide the version of the helm-charts that you want
to release as an argument.

Use the following format: `./helm-charts/release_helm_chart.sh <version>`
For example, to release version 0.22.0, run:
`./helm-charts/release_helm_chart.sh 0.22.0`

Follow the instructions provided by the script for successful execution.

Upon completion, the script will display a link on the console.
Use this link to create a pull request (PR) and merge it into
the `gh-pages` branch.

[release_script]: https://github.com/vmware-tanzu/secrets-manager/blob/main/helm-charts/release_helm_chart.sh

### 10. Initializing helm-charts

To start the release cycle, we initialize helm-charts for each official
release of VSecM. Helm-charts are continuously developed and updated
during the release development process.

At the beginning of a VSecM release, the [init_next_helm_chart.sh][init_script]
script is used to initialize the helm-charts.

To initialize a new helm-chart, run the following command using the init script:
`./helm-charts/init_next_helm_chart.sh <base-version> <new-version>`
base-version: the existing helm-charts version to be used as the base helm-chart.
new-version: the version helm-charts to be initialized.

For example: `./helm-charts/init_next_helm_chart.sh 0.21.0 0.22.0`

After execution, the script will display a link on the console.
Use this link to create a pull request (PR) and merge it into the main branch.
This will make the new helm-charts available for the VSecM release
development cycle.

[init_script]: https://github.com/vmware-tanzu/secrets-manager/blob/main/helm-charts/init_next_helm_chart.sh

### 11. All Set 🎉

You’re all set.

Happy releasing.
