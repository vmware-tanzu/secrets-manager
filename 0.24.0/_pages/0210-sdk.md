---
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title: VSecM Developer SDK
layout: post
prev_url: /docs/use-case-sidecar/
permalink: /docs/use-case-sdk/
next_url: /docs/use-case-init-container/
---

## Using **VSecM SDK**

Now, let's programmatically consume the **VSecM Safe** API from our
workload. That way, you will have more control over how you consume and cache
your secrets, and you will not need to add a sidecar to your pod.

## Cleanup

Let's remove the workload first:

```bash 
{% raw %}kubectl delete deployment example{% endraw %}
```

That will get rid of the workload; but, assuming you have completed the tutorial
before this one, you'll still have a secret registered. Let's see it:

```bash
{% raw %}# Find the sentinel pod's name:
kubectl get po -n vsecm-system

# List secrets:
kubectl exec vsecm-sentinel-778b7fdc78-86v6d -n \
  vsecm-system -- safe -l

{"secrets":[{"name":"example",
"created":"Sat May 13 20:42:20 +0000 2023",
"updated":"Sat May 13 20:42:20 +0000 2023"}]}{% endraw %}
```

Let's delete it first:

```bash 
kubectl exec vsecm-sentinel-778b7fdc78-86v6d -n \
  vsecm-system -- safe -w example -d

OK
```

And make sure that it is gone:

```bash
{% raw %}kubectl exec vsecm-sentinel-778b7fdc78-86v6d -n \
  vsecm-system -- safe -l
# Output:
# {"secrets":[]}{% endraw %}
```

All right, our cluster is as clean as a baby's butt; let's move on.

## Read the Source

Make sure [you examine the manifests][workload-yaml] to gain an understanding
of what kinds of entities you've deployed to your cluster.

[workload-yaml]: https://github.com/vmware-tanzu/secrets-manager/tree/main/examples/using_sdk/k8s

## The Benefit of Using **VSecM SDK**

**VSecM SDK** gives direct control of **VSecM Safe** to your workload.

The advantage of this approach is: you are in charge.
The downside of it is: Well, you are in charge 🙂.

But, jokes aside, your application will have to be
more tightly bound to **VMware Secrets Manager** without a sidecar.

However, when you use a sidecar, your application does not have any idea of
**VMware Secrets Manager**'s existence. From its perspective, it is merely reading from a file
that something magically updates every once in a while. This
"*separation of concerns*" can make your application architecture more
adaptable to changes.

As in anything, there is no one true way to do it. Your approach will depend
on your project's requirements.

## Deploying the Demo Workload

That part taken care of; let's deploy a workload that uses **VSecM SDK**
instead of *VSecM Sidecar*.

```bash 
# Switch to the VMware Secrets Manager repo:
cd $WORKSPACE/secrets-manager
# Install the workload:
make example-sdk-deploy
# If you are building from the source, 
# use `make example-sdk-deploy-local` instead.
```

And that's it. You have your demo workload up and running.

## The Demo App

[Here is the source code of the demo container's app][workload-src] for the
sake of completeness.

[workload-src]: https://github.com/vmware-tanzu/secrets-manager/blob/main/examples/using_sdk_go/main.go

When you check the source code, you'll see that our demo app tries to get the
secret by querying the SDK via `sentry.Fetch()`, displays the secret if it finds
and repeats this every 5 seconds in an infinite loop.

```go 
for {
	log.Println("fetch")
	d, err := sentry.Fetch()

	// ... (error handling) truncated ...

	fmt.Printf(
		"secret: updated: %s, created: %s, value: %s\n",
		d.Updated, d.Created, d.Data,
	)

	time.Sleep(5 * time.Second)
}
```

## Verifying the Deployment

If you have been following along so far, when you execute `kubectl get po` will
give you something like this:

```bash 
{% raw %}kubectl get po

NAME                              READY    AGE
example-85bdbc4cf4-6n2ng  1/1     Running  9s{% endraw %}
```

Let's check the logs of our pod:

```bash 
{% raw %}kubectl logs example-85bdbc4cf4-6n2ng -f{% endraw %}
```

The output should be something like this:

```text
{% raw %}2023/07/31 20:17:13 fetch
Failed to read the secrets file. Will retry in 5 seconds...
Secret does not exist
2023/07/31 20:17:19 fetch
Failed to read the secrets file. Will retry in 5 seconds...
Secret does not exist
...{% endraw %}
```

We don't have any secrets registered to our workload as expected. So, let's
add some.

## Registering a Secret

Let's register a secret and see how the logs change:

```bash 
{% raw %}# Find the name of the VSecM Sentinel pod:
kubectl get po -n vsecm-system

# register a secret to our workload using VSecM Sentinel:
kubectl exec vsecm-sentinel-778b7fdc78-86v6d -n vsecm-system \
  -- safe \
  -w "example" \
  -s "VSecMRocks!"
  
# Response: 
# OK{% endraw %}
```

Now let's check the logs again:

```bash 
{% raw %}kubectl logs example-85bdbc4cf4-6n2ng -f

2023/07/31 20:21:06 fetch
secret: updated: "2023-07-31T20:21:03Z", 
created: "2023-07-31T20:21:03Z", value: VSecMRocks!

...{% endraw %}
```

[demo-sidecar]: /docs/use-case-sidecar

## Conclusion

So, yay 🎉.  We got our secret; and since we use **VSecM SDK**, 
we also were able to get additional metadata such as the creation and modification 
timestamps of our secret, which was not possible to retrieve if we used the
[**VSecM Sidecar**][demo-sidecar] approach that we have seen in the
[former tutorial][demo-sidecar].

Next up, you'll learn about **VSecM Init Container**.

<p class="github-button">
	<a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0210-sdk.md">
		Suggest edits ✏️
	</a>
</p>