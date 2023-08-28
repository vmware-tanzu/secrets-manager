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

title: Contribute to VSecM
layout: post
prev_url: /docs/use-the-source/
permalink: /docs/contributing/
next_url: /docs/production/
---

<p class="github-button"
><a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0130-contributing.md"
>edit this page on <strong>GitHub</strong> ✏️</a></p>

## Introduction

This section contains instructions to test and develop **VMware Secrets Manager**
locally.

> **📚 Familiarize Yourself with the Contributing Guidelines**
> 
> Please make sure you read the [Contributing Guidelines][contributing]
> and  the [Code of Conduct][coc] on the **VSecM** GitHub repo first.
{: .block-warning}

## Good First Issues for New Contributors

If you are new to **VMware Secrets Manager** or looking for smaller tasks to 
start contributing, we have a set of issues labeled as 
[“*good first issue*”](https://github.com/vmware-tanzu/secrets-manager/labels/good%20first%20issue) 
on our GitHub repository. These issues are a great place to start if you are 
looking to make your first contribution.

### How to Find Good First Issues

1. Navigate to the [Issues tab](https://github.com/vmware-tanzu/secrets-manager/issues) 
   in the GitHub repository.
2. Use the label filter and select the [“*good first issue*”](https://github.com/vmware-tanzu/secrets-manager/labels/good%20first%20issue) 
   label.
3. Browse through the list and pick an issue that interests you.

### Claiming an Issue

Before starting work on an issue, it’s a good practice to comment on it, 
stating that you intend to work on it. This prevents multiple contributors from 
working on the same issue simultaneously.

### Need Help?

If you have questions or need further clarification on a “*good first issue,*” 
feel free to ask in the issue comments or reach out to the maintainers.

## Code Review Requirements

While we value pragmatism over process, we do have some basic requirements for 
code reviews to ensure the quality and consistency of the codebase.

### Conducting Code Reviews

1. **Pull Requests**: All code changes must be submitted through a pull request 
   (*PR*) on GitHub.
2. **Minimum Reviews**: Each PR must be reviewed by *at least one other person* 
   before it can be merged.
3. **Open for Feedback**: PRs are open for comments and suggestions from any 
   team member, not just the designated reviewer.

### What Must Be Checked

These are the minimum set of items that must be checked during a code review.
More items may be checked depending on the nature of the change.

1. **Canonical Go**: The code should adhere to canonical Go practices.
2. **Formatting**: The code must pass `gofmt` without any issues.
3. **Consistency**: The code should look like the rest of the codebase, 
   as if it were written by a single individual.

### Acceptance Criteria

1. **Approval**: At least one reviewer must approve the PR.
2. **Automated Checks**: All automated tests and checks must pass.
3. **No Conflicts**: Resolve any merge conflicts before merging.

### How to Conduct a Code Review

1. Navigate to the [Pull Requests tab](https://github.com/vmware-tanzu/secrets-manager/pulls) 
   in the GitHub repository.
2. Choose a PR that is awaiting review.
3. Review the code changes and provide your feedback, keeping the above criteria 
   in mind.
4. If the PR meets all the criteria, approve it; otherwise, request changes and 
   provide **constructive** feedback.

## Prerequisites

Other than the source code, you need the following set up for your development
environment to be able to locally develop **VMware Secrets Manager**:

* [Docker][docker] installed and running for the local user.
* [Minikube][minikube] installed on your system.
* [Make][make] installed on your system.
* [Git][git] installed on your system.

[minikube]: https://minikube.sigs.k8s.io/docs/
[make]: https://www.gnu.org/software/make/
[git]: https://git-scm.com/

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
{: .block-tip}

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
{: .block-tip}

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

## Building, Deploying, and Testing

Now let’s explain some of these steps (*and for the remainder, you can read
the `Makefile` at the project root):

* `make k8s-delete`: Deletes your minikube cluster.
* `make k8s-start`: Starts an existing cluster, or creates a brand new one.
* `make build-local`: Builds all the projects locally and pushes them to
  the local container registry.
* `make deploy-local`: Deploys **VMware Secrets Manager** locally with the 
  artifacts generated at the `build-local` step.
* `make test-local`: Runs integration tests to make sure that the changes
  that were made doesn’t break anything.

If you run these commands in the above order, you’ll be able to **build**,
**deploy**, and **test** your work locally.

[docker]: https://www.docker.com/

## Minikube Quirks

### Docker for Mac Troubleshooting

When you use **Minikube** with **Docker for Mac**, you’ll likely get a warning
similar to the following:

```bash {% raw %}
make k8s-start

# below is the response to the command above

./hack/minikube-start.sh

…
… truncated …
…

Registry addon with docker driver uses port 50565
please use that instead of 
default port 5000 

…{% endraw %}
```

The port `50656` is a randomly-selected port. Every time you run `make k8s-start`
it will pick a different port.

You can verify that the repository is there:

```bash 
{% raw %}curl localhost:50565/v2/_catalog
# response:
# {"repositories":[]} 
{% endraw %}
```

There are two issues here:

* First, all the local development scripts assume port 5000 as the repository port;
  however, port `5000` on your Mac will likely be controlled by the 
  **Airplay Receiver**.
* And secondly, you’ll need to forward `localhost:5000` to whatever port the error
  message shows you.

To fix the first issue, on your Mac’s “**System Settings**” you’ll need to go to
“**Setting » Airdrop & Handoff » Airplay Receiver**” and on that screen…

* **uncheck** “*allow handoff between this Mac and your iCloud devices*”,
* make sure “**Airdrop**” is selected as “**no one**”
* and finally, after updating your settings,  **restart your Mac**
  (*this step is important; without restart, your macOS will still hold onto
  that port*)

Note that where these settings are can be slightly different from one version
of macOS to another.

As for the second issue, to redirect your local `:5000` port to the docker engine’s
designated port, you can use [`socat`][socat].

[socat]: http://www.dest-unreach.org/socat/ "socat: Multipurpose Relay"

```bash 
{% raw %}
# Install `socat` if you don’t have it on your system.
brew install socat

# Replace 49876 with whatever port the warning message 
# gave you during the initial cluster setup.
socat TCP-LISTEN:5000,fork,reuseaddr TCP:localhost:49876
{% endraw %}
```

Then execute the following on a separate tab:

```bash
{% raw %}
curl localhost:5000/v2/_catalog

# Should return something similar to this:
# {"repositories":[]}
{% endraw %}
```

If you get a successful response to the above `curl`, then congratulations,
you have successfully set up your local docker registry for your 
**VMware Secrets Manager** development needs.

> **Make a Big McTunnel**
>
> If you have `localhost:5000` unallocated, there is a `make mac-tunnel`
> target in the **VMware Secrets Manager**’ project **Makefile** that will 
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
{: .block-tip}

### Ubuntu Troubleshooting

If you are using **Ubuntu**, it would be helpful to know that **Minikube** and
**snap** version of **Docker** do not play well together. If you are having
registry-related issues, or if you are not able to execute a `docker images`
without being the root user, then one resolution can be to remove the snap
version of docker and install it from the source:

```bash 
{% raw %}sudo apt update
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
sudo apt-get install docker-ce docker-ce-cli containerd.io{% endraw %}
```

> **Restart Your System**
>
> After doing this, you might need to restart your system and execute
> `minikube delete` on your terminal too. Although you might feel that this
> step is optional, it is **not**; trust me 🙂.
{: .block-tip}

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
> and worst case it’s still worth giving a try.
{: .block-tip}

[post-installation]: https://docs.docker.com/engine/install/linux-postinstall/

Still no luck? Keep on reading.

Depending on your operating system, and the Minikube version that you use
it might take a while to find a way to push images to Minikube’s internal
Docker registry. [The relevant section about the Minikube handbook][minikube-push]
covers a lot of details, and can be helpful; however, it is also really easy
skip or overlook certain steps.

If you have `docker push` issues, or you have problem your Kubernetes Pods
acquiring images from the local registry, try these:

* Execute `eval $(minikube docker-env)` before pushing things to **Docker**. This
  is one of the first instructions [on the “*push*” section of the Minikube
  handbook][minikube-push], yet it is still very easy to inadvertently skip it.
* Make sure you have the registry addon enabled
  * (`minikube addons list`).
* You might have luck directly pushing the image:
  * first: `docker build --tag $(minikube ip):5000/test-img`;
  * followed by: `docker push $(minikube ip):5000/test-img`.
* There are also `minikube image load` and `minikube image build` commands
  that might be helpful.

[minikube-push]: https://minikube.sigs.k8s.io/docs/handbook/pushing/

## Enjoy 🎉

Just explore the [Makefile][makefile] and get a feeling of it.

[Feel free to touch base](/docs/community/) if you have any questions, comments,
recommendations.

[makefile]: https://github.com/vmware-tanzu/secrets-manager/blob/main/Makefile
