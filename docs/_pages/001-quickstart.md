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

title: Quickstart
layout: post
---

## VMware Secrets Manager for Cloud Native Apps

This is a quickstart guide to get you up and running with **VMware Secrets 
Manager**.

## Prerequisites

* [**Minikube**][minikube]: You can install **VMware Secrets Manager** on any 
  Kubernetes cluster, but we’ll use *Minikube* in this quickstart example. 
  Minikube is a tool that makes it easy to run Kubernetes locally. 
* [**make**][make]: You’ll need `make` to run certain build tasks. You can 
  install `make` using your favorite package manager.
* [**Docker**][docker]: This quickstart guide assumes that **Minikube** uses 
  the *Docker* driver. If you use a different driver, things will still likely
  work, but you might need to tweak some of the commands and configuration.

[minikube]: https://minikube.sigs.k8s.io/docs/ "Minikube"
[make]: https://www.gnu.org/software/make/ "GNU Make"
[docker]: https://www.docker.com "Docker"

## A Video Is Worth A Lot of Words

Here’s a video that walks you through the steps in this quickstart guide:

<div style="padding:56.25% 0 0 0;position:relative;"><iframe 
src="https://player.vimeo.com/video/849328819?h=46caa595f7&amp;badge=0&amp;autopause=0&amp;player_id=0&amp;app_id=58479" 
frameborder="0" allow="autoplay; fullscreen; picture-in-picture" allowfullscreen 
style="position:absolute;top:0;left:0;width:100%;height:100%;" 
title="VMware Secrets Manager (for Cloud-Native Apps) Quickstart"></iframe></div>
<script src="https://player.vimeo.com/api/player.js"></script>


## Clone the Repository

Let’s start by cloning the **VMware Secrets Manager** repository first:

```bash
cd $WORKSPACE
git clone https://github.com/vmware-tanzu/secrets-manager.git
cd secrets-manager
```

## Initialize Minikube

Next, let’s initialize *Minikube*:

```bash
cd $WORKSPACE/secrets-manager

# If you have a previous minikube setup, you can delete it with:
# make k8s-delete

make k8s-start
```

## Configure Docker Environment Variables

Next, let’s configure the Docker environment variables for Minikube.

We don’t strictly need this for the quickstart, but it’s a good idea to do
it anyway.

```bash
eval $(minikube -p minikube docker-env)
```

## Install SPIRE and VMware Secrets Manager

Next we’ll install [SPIRE][spire] and **VMware Secrets Manager** on the cluster.

```bash 
cd $WORKSPACE/secrets-manager
make deploy
```

This will take a few minutes to complete.

[spire]: https://spiffe.io/spire/ "SPIRE"

## Verify Installation

Let’s check that the installation was successful by listing the pods int
the `spire-system` and `vsecm-system` namespaces:

```bash
kubectl get po -n spire-system
```

```text
# Output:
NAME                           READY   STATUS    RESTARTS      AGE
spire-agent-p9m27              3/3     Running   1 (23s ago)   29s
spire-server-6fb4f57c8-6s7ns   2/2     Running   0             29s
```

```bash
kubectl get po -n vsecm-system
```

```text
# Output:
NAME                             READY   STATUS    RESTARTS   AGE
vsecm-safe-85dd95949c-f4mhj      1/1     Running   0          25s
vsecm-sentinel-6dc9b476f-djnq7   1/1     Running   0          24s
```

All the pods look up and running, so we can move on to the next step.

## List Available Commands

`make help` lists all the available make targets in a cheat sheet format:

```bash
make help
```

Here’s the output of this command:
  
```text
--------------------------------------------------------------------
          🛡️ VMware Secrets Manager: Keep your secrets… secret.
          🛡️ https://vsecm.com/
--------------------------------------------------------------------
        ℹ️ This Makefile assumes you use Minikube and Docker
        ℹ️ for most operations.
--------------------------------------------------------------------

… truncated …
 
--------------------------------------------------------------------
  Example Use Cases:
    Using local images:
          ˃ make example-sidecar-deploy-local;
          ˃ make example-sdk-deploy-local;
          ˃ make example-multiple-secrets-deploy-local;
    Using remote images:
          ˃ make example-sidecar-deploy;
          ˃ make example-sdk-deploy;
          ˃ make example-multiple-secrets-deploy;

… truncated …

--------------------------------------------------------------------
```

## Deploy an Example Workload

Now, let’s deploy an example workload to the cluster to test 
**VMware Secrets Manager** in action.

```bash
cd $WORKSPACE/secrets-manager
make example-sdk-deploy
```

This will take a few moments too.

When done you would ba able to list the pods in the `default` namespace:

```bash
kubectl get po -n default
```

```text
# Output
NAME                       READY   STATUS    RESTARTS   AGE
example-68997489c6-8j8kj   1/1     Running   0          1m51s
```

Let’s check the logs of our example workload:

```bash
kubectl get logs example-68997489c6-8j8kj
```

The output will be something similar to this:

```text
2023/07/28 01:26:51 fetch
Failed to read the secrets file. Will retry in 5 seconds…
Secret does not exist
2023/07/28 01:27:03 fetch
Failed to read the secrets file. Will retry in 5 seconds…
Secret does not exist
2023/07/28 01:27:08 fetch
Failed to read the secrets file. Will retry in 5 seconds…
… truncated …
```

Our sample workload is trying to fetch a secret, but it can’t find it.

Here’s the source code of our sample workload to provide some context:

```go
package main

// … truncated headers … 

func main() {
	
	// … truncated irrelevant code …
	
	for {
		log.Println("fetch")
		d, err := sentry.Fetch()

		if err != nil {
			fmt.Println("Failed. Will retry in 5 seconds…")
			fmt.Println(err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		if d.Data == "" {
			fmt.Println("No secret yet… will check again later.")
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Printf(
			"secret: updated: %s, created: %s, value: %s\n",
			d.Updated, d.Created, d.Data,
		)
		time.Sleep(5 * time.Second)
	}
}
```

What the demo workload does is to try to fetch a secret every 5 seconds
using the `sentry.Fetch()` function.

`sentery.Fetch()` is a function provided by the **VMware Secrets Manager**;
it establishes a secure mTLS connection between the workload and
**VMware Secrets Manager Safe** to fetch the secret.

Since this workload does not have any secret registered, the request fails
and the workload retries every 5 seconds.

Since this is a quickstart example, we won’t dive into the details of 
how the workload establishes a secure mTLS connection with the 
**VMware Secrets Manager Safe**. We’ll cover this in the following sections.

For the sake of this quickstart, we can assume that secure communication
between the workload and the **VMware Secrets Manager Safe** is already taken
care of for us.

## Register a Secret

Now, let’s register a secret and see what happens.

To register a secret we’ll need to find the `vsecm-sentinel` pod in the 
`vsecm-system` namespace and execute a command inside the pod.

Let’s get the pod first:

```bash
kubectl get po -n vsecm-system
```

Here’s a sample output:

```text
NAME                             READY   STATUS    RESTARTS   AGE
vsecm-safe-85dd95949c-f4mhj      1/1     Running   0          4h29m
vsecm-sentinel-6dc9b476f-djnq7   1/1     Running   0          4h29m
```

`vsecm-sentinel-6dc9b476f-djnq7` is what we need here.

Let’s use it and register a secret to our example workload:

```bash
kubectl exec vsecm-sentinel-6dc9b476f-djnq7 -n vsecm-system -- \
safe -w "example" -n "default" -s "VSecMRocks"
```

> **Sentinel Command Line Help**
>
> **VMware Secrets Manager Sentinel** comes with a command line tool 
> called`safe`. `safe` allows you to register secrets to
> **VMware Secrets Manager Safe**, delete secrets, or list existing secrets.
>
> You can execute `safe -h` or `safe --help` to get a list of available
> commands and options.
{: .block-tip }

You’ll get an `OK` as a response:

```text
OK
```

For the command `safe -w "example" -n "default" -s "VSecMRocks"`

* `-w` is the workload name
* `-n` is the namespace
* `-s` is the secret value

But how do you know what the workload name is?

That’s where **ClusterSPIFFEID** comes in:

```bash 
kubectl get ClusterSPIFFEID
```

And here’s the output:

```text
NAME             AGE
example          4h33m
vsecm-safe       4h35m
vsecm-sentinel   4h35m
```

Let’s see the details of this `example` SPIFFE ID:

```bash
kubectl describe ClusterSPIFFEID example
```

And the output:

```text 
{% raw %}Name:         example
Namespace:
Labels:       <none>
Annotations:  <none>
API Version:  spire.spiffe.io/v1alpha1
Kind:         ClusterSPIFFEID
Metadata:
  Creation Timestamp:  2023-07-28T01:26:48Z
  Generation:          1
  Resource Version:    832
  UID:                 b254294e-eed0-4116-8b38-2fb1d101e387
Spec:
  Pod Selector:
    Match Labels:
      app.kubernetes.io/name:  example
  Spiffe ID Template:          spiffe://vsecm.com/workload/example/
  ns/{{ .PodMeta.Namespace }}/
  sa/{{ .PodSpec.ServiceAccountName }}/n/{{ .PodMeta.Name }}
  Workload Selector Templates:
    k8s:ns:default
    k8s:sa:example
Status:
  Stats:
    Entries Masked:             0
    Entries To Set:             1
    Entry Failures:             0
    Namespaces Ignored:         4
    Namespaces Selected:        6
    Pod Entry Render Failures:  0
    Pods Selected:              1{% endraw %}
```

For the sake of keeping things simple because this is a quickstart, we can 
assume that someone has created this `example` SPIFFE ID for us, and using
this SPIFFE ID, our example workload can securely communicate with the
**VMware Secrets Manager Safe**.

## Verifying Secret Registration

Since we’ve registered a secret, let’s see if our example workload can fetch
the secret now and display it in its logs.

```bash
kubectl get logs example-68997489c6-8j8kj
```

And the output would be something like this:

```text 
2023/07/28 06:06:39 fetch
secret: updated: "2023-07-28T01:34:30Z", 
created: "2023-07-28T01:34:30Z", value: VSecMRocks
2023/07/28 06:06:44 fetch
secret: updated: "2023-07-28T01:34:30Z", 
created: "2023-07-28T01:34:30Z", value: VSecMRocks
2023/07/28 06:06:49 fetch
secret: updated: "2023-07-28T01:34:30Z", 
created: "2023-07-28T01:34:30Z", value: VSecMRocks
```

As you can see, the secret is now fetched and displayed in the logs.

The beauty of this approach is when we change the secret using
**VMware Secrets Manager Sentinel**, the workload will automatically fetch the
new value, without having to restart itself.

## Where to Go From Here

This quickstart is meant to give you a quick overview of how you can use
**VMware Secrets Manager** to securely manage secrets in your Kubernetes
clusters.

After successfully completing this quickstart, you can try the following:

* [Join the **VMware Secrets Manager** Community on **Slack**][slack-invite]
* [Learn more about **VMware Secrets Manager** Architecture](/docs/architecture)
* [Learn more about **VMware Secrets Manager** Design Philosophy](/docs/sentinel)
* [Follow a more detailed tutorial that contains multiple use cases](/docs/tutorial)
* Navigate this website to learn more about **VMware Secrets Manager**

[slack-invite]: https://join.slack.com/t/a-101-103-105-s/shared_invite/zt-1zrr2yepf-2P3EJhfoGNn05l5_4jvYSA "Join VSecM Slack"
