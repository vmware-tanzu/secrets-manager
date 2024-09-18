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

title = "Release Management"
weight = 13
+++

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
won't be able to push the images to the registry and tag the release.

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
eval some environment variables to be able to use Minikube's registry instead
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

## Publishing the Current Documentation

Before cutting a release, make sure to publish the current documentation.

Make sure you have a local Docker server running before
publishing the documentation.

```bash
cd $WORKSPACE/secrets-manager

# Creates a local server to preview the documentation.
# Navigate to https://localhost:8000/ and make sure
# everything looks good.
./hack/web-serve.sh

# Run this while the local web server is still up
# to publish the documentation:
./hack/web-synch.sh
```

## Cutting a Release

Before every release cut, follow the steps outlined below.

### 0. Are you on a release branch?

Make sure you are on a release branch, forked off of the most recent `main` 
branch where all the changes to be included in the release are merged.

### 1. Pre-Release Checks

Before cutting a new release, perform the following security and code quality 
checks:

#### `go vet`

* Run `go vet ./...` and review results.

#### `govulncheck`

* Install: `go install golang.org/x/vuln/cmd/govulncheck@latest`
* Run: `govulncheck ./...` and review results.

#### Snyk:
    
* Install Snyk CLI: `npm install -g snyk`
* Navigate to the project directory: `cd $WORKSPACE/secrets-manager`
* Authenticate (if not already done): `snyk auth`
* Run: `snyk test` and review results.
* Monitor projects on https://app.snyk.io/org/$username/projects
  (*replace $username with your actual username*)

#### `golangci-lint`:

* Run: `golangci-lint run` and review results.

#### Manual review:

* Address any new vulnerabilities or issues found before proceeding with the 
  release.
* Document the results of these checks in the release notes.

### 2. Check Docker and Minikube

Also make sure your `docker` and `Minikube` are up and running.

Additionally, execute `eval $(minikube -p minikube docker-env)` once more to
update your environment.

### 3. `make help`

Check the `make help` command first, as it includes important information.

You can also check `make h` command that included release-related commands.

### 4. Test VSecM Distroless Images

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

### 5. Merge the Release Branch to `main`

If all tests pass, merge the release branch to `main`.

### 6. Tagging

Tagging needs to be done **on the build server**.

There is no automation for this yet.

> **Don't forget to Bump the Version**
>
> If you are cutting a new release, do not forget to bump the version,
> before running the tagging script below.


```bash
git checkout main
git stash
git pull
export DOCKER_CONTENT_TRUST=1
make build
make tag
```

### 7. Initializing Helm Charts

To start the release cycle, we initialize helm-charts for each official
release of VSecM. Helm charts are continuously developed and updated
during the release development process.

At the beginning of a VSecM release, the [./hack/init-next-helm-chart.sh][init_script]
script is used to initialize the helm-charts.

To initialize a new helm-chart, run the following command using the init script:
`./hack/init-next-helm-chart.sh <base-version> <new-version>`
base-version: the existing helm-charts version to be used as the base helm-chart.
new-version: the version helm-charts to be initialized.

For example: `./hack/init-next-helm-chart.sh 0.22.2 0.22.4`

After execution, the script will display a link on the console.
Use this link to create a pull request (PR) and merge it into the main branch.
This will make the new helm-charts available for the VSecM release
development cycle.

### 8. Update Kubernetes Manifests

Based on the generated helm charts run `make k8s-manifests-update VERSION=<version>` target
to update the Kubernetes manifests for the new release.

These manifests are used by people who want to install VSecM without using
Helm. To generate the manifests you need to have generated the helm charts
first.

For example `make k8s-manifests-update VERSION=0.22.4`

### 9. Update Helm Documentation

If you have updated inline documentation in helm charts, make sure to reflect
the changes by running `./hack/helm-docs.sh`.

### 10. Release Helm Charts

> **Pull Recent `gh-pages` Changes**
>
> Before you proceed, make sure that you have your `gh-pages` local branc
> is up-to-date:
>
> ```bash
> cd $WORKSPACE/secrets-manager
> git checkout gh-pages
> git pull
> git checkout main
> ```

We offer the [./hack/release-helm-chart.sh][release_script] script for your use.
To execute the script, provide the version of the helm-charts that you want
to release as an argument.

Use the following format: `./hack/release-helm-chart.sh <version>`
For example, to release version 0.22.4, run:
`./hack/release-helm-chart.sh 0.22.4`

Follow the instructions provided by the script for successful execution.

Upon completion, the script will display a link on the console.
Use this link to create a pull request (PR) and merge it into
the `gh-pages` branch.

> **Keep The Most Recent Version of the Helm Charts**
>
> Make sure you keep only the most recent version of the Helm Charts in the
> `main` branch. Older versions should be snapshotted in the `gh-pages` branch
> using the workflow described above.

### 11. Add a Snapshot of the Current Documentation

The `docs` branch contains a snapshot of each documentation in versioned
folders.

To add a snapshot of the current documentation:

1. Copy the `docs` folder into a temporary place like `/tmp/docs`.
2. Checkout the `docs` branch.
3. Copy the `docs` folder from `/tmp/docs` to the `docs` branch:
   `cp -r /tmp/docs $WORKSPACE/secrets-manager/docs/<version>`.
4. Update the `secrets-manager/docs/<version>/_includes/notification.html` file
   to include a link to the new documentation. You can copy the message from
   one of the existing versioned `notification.html` files.
5. Edit `./hack/publish-docs.sh` to include the new version.
6. Execute `./hack/publish-docs.sh` to publish the archived documentation.
7. Create a PR and merge it into the `docs` branch.
8. Checkout the `main` branch.
9. Update `... 0031-documentation-snapshots.md` to include a link to the new
   documentation snapshot.
10. Create a PR and merge it into the `main` branch.

[release_script]: https://github.com/vmware-tanzu/secrets-manager/blob/main/hack/release-helm-chart.sh

[init_script]: https://github.com/vmware-tanzu/secrets-manager/blob/main/hack/init-next-helm-chart.sh

### 12. All Set 🎉

You're all set.

Happy releasing.

{{ edit() }}
