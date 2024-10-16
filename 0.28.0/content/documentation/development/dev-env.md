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

title = "Development Environment"
weight = 12
+++

## Introduction

This section contains instructions to set up your development environment
to actively contribute the **VMware Secrets Manager** project.

## Overview

Here is a brief recording of the steps that you need to take to build and test
**VMware Secrets Manager** locally:

<script 
  src="https://asciinema.org/a/676190.js"
  id="asciicast-676190" 
  async="true"></script>

## Prerequisites

Other than the source code, you need the following set up for your development
environment to be able to locally develop **VMware Secrets Manager**:

* [Docker][docker] installed and running for the local user.
* [Minikube][minikube] installed on your system.
* [Make][make] installed on your system.
* [Git][git] installed on your system.
* [Go][go] installed and configured on your system.

[minikube]: https://minikube.sigs.k8s.io/docs/
[make]: https://www.gnu.org/software/make/
[git]: https://git-scm.com/
[go]: https://go.dev/

> **Can I Use Something Other than Minikube and Docker**?
>
> Of course, you can use any Kubernetes cluster to develop, deploy, and test
> **VMware Secrets Manager** for Cloud-Native Apps.
>
> Similarly, you can use any OCI-compliant container runtime. It does not
> have to be Docker.
>
> We are giving **Minikube** and **Docker** as an example because they are
> easier to set up; and when you stumble upon, it is easy to find supporting
> documentation about these to help you out.
>
> For **macOS**, for example, you can check the
> [alternate setup](#alternate-non-minikube-setup-on-macos)
> section below.

## Alternate Non-Minikube Setup on macOS

The rest of this document assumes that you are using **Minikube** and **Docker**;
however, if you are on a Mac, want to use **Docker for Mac's Kubernetes Distribution**,
you can install them as described below and skip the sections that are
specific to **Minikube**.

### Installing Docker for Mac

Download and install [Docker for Mac][docker-mac] on your system.

[docker-mac]: https://docs.docker.com/desktop/install/mac-install/

### Enable Kubernetes on Docker for Mac

Once you have Docker for Mac installed, you'll need to enable Kubernetes on it.
To do that, follow the steps below:

1. Open Docker for Mac's preferences.
2. Go to the "**Kubernetes**" tab.
3. Check the "**Enable Kubernetes**" checkbox.

### Disabling Airplay Receiver

Airplay Receiver uses port 5000 by default, which is the same port that dokcer
registry uses.

To fix the first issue, on your Mac's "**System Settings**" you'll need to go to
"**Setting » Airdrop & Handoff » Airplay Receiver**" and on that screen...

* **uncheck** "*allow handoff between this Mac and your iCloud devices*",
* make sure "**Airdrop**" is selected as "**no one**".
* then you might need to kill the process that's using port 5000 or restart your
  system.

Alternatively, you can bind a different port to the docker registry. But when
you do that, you'll need to make sure that you update other files in the
project too that reference `localhost:5000`.

### Install Docker Registry

Use the following command to install the docker registry:

```bash
docker run -d -p 5000:5000 --restart=always --name registry registry:2
```

### You Are All Set

And you should be all set.

You can run `make build-local` to build local images, and `make deploy-local` to
build and install **VMware Secrets Manager** locally.

## Alternate Non-Minikube Setup Using Kind

If you are using [Kind][kind] to set up your local Kubernetes cluster, you
don't need to do anything special. Just make sure that you have a local
container registry running on port `5000`:

```bash
docker run -d -p 5000:5000 --restart=always --name registry registry:2
```

You may also want to check out [`kind`'s instructions for setting up a local
registry][kind-local-registry].

[kind]: https://kind.sigs.k8s.io/
[kind-local-registry]: https://kind.sigs.k8s.io/docs/user/local-registry/

Then initialize `kind` with the following command:

```bash
kind create cluster
```

Then you can run `make build-local` to build local images, and
`make deploy-local` install **VMware Secrets Manager** locally.

[kind]: https://kind.sigs.k8s.io/

## Cloning VMware Secrets Manager

Create a workspace folder and clone **VMware Secrets Manager** into it.

```bash 
mkdir $HOME/Desktop/WORKSPACE
cd $HOME/Desktop/WORKSPACE
git clone "https://github.com/vmware-tanzu/secrets-manager.git"
cd secrets-manager 
```

> **Want to Create a Pull Request**?
>
> If you are contributing to the source code, make sure you read
> [the contributing guidelines][contributing], and [the code of conduct][coc].


[fork]: https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/about-forks
[contributing]: https://github.com/vmware-tanzu/secrets-manager/blob/main/CONTRIBUTING_DCO.md
[coc]: https://github.com/vmware-tanzu/secrets-manager/blob/main/CODE_OF_CONDUCT.md

## Getting Help

Running `make help` at the project root will provide you with a list of
logically grouped commands. This page will not include the output because
the content of the output can change depending on the version of **VMware Secrets
Manager**; however, the output will give a nice overview of what you can do
with the `Makefile` at the project root.

```bash
make help
```

Additionally, you can run `make h` at the root of each project to get
a more release-specific help output.

```bash
make h
```

Both of these commands gives a brief overview of what you can do with the
make targets. If you want to learn more about a specific target, you can
read the source code of the relevant file inside the `./makefiles` folder.

## Building, Deploying, and Testing

Now let's explain some of these steps (*and for the remainder, you can read
the `Makefile` at the project root):

* `make k8s-delete`: Deletes your minikube cluster.
* `make k8s-start`: Starts an existing cluster, or creates a brand new one.
* `make build-local`: Builds all the projects locally and pushes them to
  the local container registry.
* `make deploy-local`: Deploys **VMware Secrets Manager** locally with the
  artifacts generated at the `build-local` step.
* `make test-local`: Runs integration tests to make sure that the changes
  that were made doesn't break anything.

If you run these commands in the above order, you'll be able to **build**,
**deploy**, and **test** your work locally.

[docker]: https://www.docker.com/

## Kubernetes-Related Troubleshooting

If you delete the `spire-system` and `spire-server` namespaces before deleting
the `vsecm-system` namespace, then you might have issues deleting the 
`vsecm-system` namespace because the Pods in the `vsecm-system` namespace
will remain hanging waiting on the SPIFFE CSI Driver’s volume to mount.

This can also happen if you install **VSecM** using `helm` and then execute
`helm delete` to remove the **VSecM** installation. To uninstall **VSecM**
helm deployment, we recommend using the `make helm-delete` target, which takes
care of the order of deletion. 

Or, alternatively, you can first delete the `vsecm-system` namespace, and then
execute `helm delete` to remove the **VSecM** installation.

```bash
kubectl delete namespace vsecm-system
helm delete vsecm
```

If you still see resources hanging, you can use the following command to
detect the resources that are still hanging and force delete them:

```bash
kubectl api-resources --verbs=list --namespaced -o name \
  | xargs -n 1 kubectl get -n vsecm-system
```

For example if `vsecm-keystone-59fc9568b6-zx7zl` pod is `Pending` or is in 
a `FailedMount` state, you can first delete its owner deployment, and then
force delete the pod with the following command:

```bash
kubectl delete deployment vsecm-keystone -n vsecm-system
kubectl delete pod vsecm-keystone-59fc9568b6-zx7zl \
  -n vsecm-system --grace-period=0 --force
```

When you do this for all the hanging resources, you can then delete the 
`vsecm-system` namespace.

```bash
kubectl delete namespace vsecm-system
```

## Minikube Quirks

### Docker for Mac Troubleshooting

When you use **Minikube** with **Docker for Mac**, you'll likely get a warning
similar to the following:

```bash 
make k8s-start

# below is the response to the command above

./hack/minikube-start.sh

...
... truncated ...
...

Registry addon with docker driver uses port 50565
please use that instead of 
default port 5000 

...
```

The port `50656` is a randomly-selected port. Every time you run `make k8s-start`
it will pick a different port.

You can verify that the repository is there:

```bash 
curl localhost:50565/v2/_catalog
# response:
# {"repositories":[]} 

```

> **Do You Need Multiple Nodes**
>
> If you are using Minikube Docker driver, you can configure it to start
> with multiple nodes. This can be useful if you want to test an
> advanced scenario where more than one **SPIRE Agent** may be required.
> 
> For that, set up the `VSECM_MINIKUBE_NODE_COUNT` variable to the number
> of Kubernetes worker nodes you wish to use in your development environment.

There are two issues here:

* First, all the local development scripts assume port 5000 as the repository port;
  however, port `5000` on your Mac will likely be controlled by the
  **Airplay Receiver**.
* And secondly, you'll need to forward `localhost:5000` to whatever port the error
  message shows you.

To fix the first issue, on your Mac's "**System Settings**" you'll need to go to
"**Setting » Airdrop & Handoff » Airplay Receiver**" and on that screen...

* **uncheck** "*allow handoff between this Mac and your iCloud devices*",
* make sure "**Airdrop**" is selected as "**no one**"
* and finally, after updating your settings,  **restart your Mac**
  (*this step is important; without restart, your macOS will still hold onto
  that port*)

Note that where these settings are can be slightly different from one version
of macOS to another.

As for the second issue, to redirect your local `:5000` port to the docker engine's
designated port, you can use [`socat`][socat].

[socat]: http://www.dest-unreach.org/socat/ "socat: Multipurpose Relay"

```bash 

# Install `socat` if you don't have it on your system.
brew install socat

# Replace 49876 with whatever port the warning message 
# gave you during the initial cluster setup.
socat TCP-LISTEN:5000,fork,reuseaddr TCP:localhost:49876

```

Then execute the following on a separate tab:

```bash

curl localhost:5000/v2/_catalog

# Should return something similar to this:
# {"repositories":[]}

```

If you get a successful response to the above `curl`, then congratulations,
you have successfully set up your local docker registry for your
**VMware Secrets Manager** development needs.

> **Make a Big McTunnel**
>
> If you have `localhost:5000` unallocated, there is a `make mac-tunnel`
> target in the **VMware Secrets Manager**'s project **Makefile** that will
> automatically find the exposed docker registry port, and establish a
> tunnel for you.
>
> Execute this:
>
> ```bash
> make mac-tunnel
> ```
>
> And then on a separate terminal window, verify that you can access the
> registry from `localhost:5000`.
>
> ```bash
> curl localhost:5000/v2/_catalog
> # Should return something similar to this:
> # "repositories":[]
> ```


### Ubuntu Troubleshooting

If you are using **Ubuntu**, it would be helpful to know that **Minikube** and
**snap** version of **Docker** do not play well together. If you are having
registry-related issues, or if you are not able to execute a `docker images`
without being the root user, then one resolution can be to remove the snap
version of docker and install it from the source:

```bash 
sudo apt update
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | \ 
    sudo apt-key add -
sudo apt-key fingerprint 0EBFCD88
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
```

> **Restart Your System**
>
> After doing this, you might need to restart your system and execute
> `minikube delete` on your terminal too. Although you might feel that this
> step is optional, it is **not**; trust me 🙂.


After installing a non-snap version of **Docker** and restarting your system, if
you can use **Minikube** *Docker registry*, then, perfect. If not, there are
a few things that you might want to try. So if you are still having issues
keep on reading.

Before trying anything else, it might be worth giving [Docker Post Installation
Steps][post-installation] from the official Docker website a shot. Following
that guideline **may** solve Docker-related permission issues that you might
still be having.

> **Restart, Maybe?**
>
> If you still have permission issues after following the official Docker
> post-installation steps outlined above, try **restarting** your computer once
> more.
>
> Especially when it comes to Docker permissions, restarting can help,
> and worst case it's still worth giving a try.


[post-installation]: https://docs.docker.com/engine/install/linux-postinstall/

Still no luck? Keep on reading.

Depending on your operating system, and the Minikube version that you use
it might take a while to find a way to push images to Minikube's internal
Docker registry. [The relevant section about the Minikube handbook][minikube-push]
covers a lot of details, and can be helpful; however, it is also really easy
skip or overlook certain steps.

If you have `docker push` issues, or you have problem your Kubernetes Pods
acquiring images from the local registry, try these:

* Execute `eval $(minikube docker-env)` before pushing things to **Docker**. This
  is one of the first instructions [on the "*push*" section of the Minikube
  handbook][minikube-push], yet it is still very easy to inadvertently skip it.
* Make sure you have the registry addon enabled
    * (`minikube addons list`).
* You might have luck directly pushing the image:
    * first: `docker build --tag $(minikube ip):5000/test-img`;
    * followed by: `docker push $(minikube ip):5000/test-img`.
* There are also `minikube image load` and `minikube image build` commands
  that might be helpful.

[minikube-push]: https://minikube.sigs.k8s.io/docs/handbook/pushing/

## Checking Logs

It's always a good idea to check **SPIRE Server**'s and **VSecM Safe**'s logs
to ensure that they are running as expected.

To check **SPIRE Server**'s logs, execute the following:

```bash
kubectl logs -n spire-system $NAME_OF_SPIRE_SERVER_POD -f
```

To check **VSecM Safe**'s logs, execute the following:

```bash
kubectl logs -n vsecm-system $NAME_OF_VSECM_SAFE_POD -f 
```

## Inspecting SPIRE Server via CLI

**SPIRE Server** has a command lin interface that you can use to directly
interact with it. This can prove to be useful when you are debugging issues.

Here's an example:

```bash
# $SPIRE_SERVER is the name of the SPIRE Server pod.
kubectl exec -n spire-system $SPIRE_SERVER -- \
  /opt/spire/bin/spire-server

## Output:
# Usage: spire-server [--version] [--help] <command> [<args>]
#
# Available commands are:
#    agent
#    bundle
#    entry
#    federation
#    healthcheck    Determines server health status
#    jwt
#    run            Runs the server
#    token
#    validate       Validates a SPIRE server configuration file
#    x509
```

You can also check registration entries:

```bash 
kubectl exec spire-server-0 -n spire-server \
  -- /opt/spire/bin/spire-server entry show \
  -spiffeID spiffe://vsecm.com/ns/vsecm-system/sa/vsecm-safe

# Found 1 entry
# Entry ID     : ric.4746eca9-0853-468e-9b4e-6283a46912df
# SPIFFE ID    : spiffe://vsecm.com/ns/vsecm-system/sa/vsecm-safe
# Parent ID    : spiffe://vsecm.com/spire/agent/k8s_psat/o-ric/fe
# Revision     : 0
# X509-SVID TTL: default
# JWT-SVID TTL : default
# Selector     : k8s:pod-uid:c5e3728c-35a6-482c-8208-543fb2e6630e
```


## Enjoy 🎉

Just explore the [Makefile][makefile] and get a feeling of it.

[Feel free to touch base](@/community/hello.md) if you have any questions, comments,
recommendations.

[makefile]: https://github.com/vmware-tanzu/secrets-manager/blob/main/Makefile

{{ edit() }}
