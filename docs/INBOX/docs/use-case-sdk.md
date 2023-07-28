---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, use case, sdk
title: Using Aegis SDK
description: a finer control over secrets
micro_nav: true
page_nav:
  prev:
    content: using <strong>Aegis</strong> sidecar
    url: '/docs/use-case-sidecar'
  next:
    content: using <strong>Aegis</strong> init container
    url: '/docs/use-case-init-container'
---

<p style="text-align:right;position:relative;top:-40px;"
><a href="https://github.com/ShieldWorks/aegis-web/blob/main/docs/use-case-sdk.md" 
style="border-bottom: none;background:#e0e0e0;padding:0.5em;display:inline-block;
border-radius:8px;">
edit this page on <strong>GitHub</strong> ✏️</a></p>

## Using **Aegis SDK**

Now, let’s programmatically consume the **Aegis Secrets** API from our
workload. That way, you will have more control over how you consume and cache
your secrets, and you will not need to add a sidecar to your pod.

## Cleanup

Let’s remove the workload first:

```bash 
{% raw %}kubectl delete deployment example{% endraw %}
```

That will get rid of the workload; but, assuming you have completed the tutorial
before this one, you’ll still have a secret registered. Let’s see it:

```bash
{% raw %}# Find the sentinel pod’s name:
kubectl get po -n aegis-system

# List secrets:
kubectl exec aegis-sentinel-778b7fdc78-86v6d -n \
  aegis-system -- aegis -l

{"secrets":[{"name":"example",
"created":"Sat May 13 20:42:20 +0000 2023",
"updated":"Sat May 13 20:42:20 +0000 2023"}]}{% endraw %}
```

Let’s delete it first:

```bash 
kubectl exec aegis-sentinel-778b7fdc78-86v6d -n \
  aegis-system -- aegis -w example -d

OK
```

And make sure that it is gone:

```bash
{% raw %}kubectl exec aegis-sentinel-778b7fdc78-86v6d -n \
  aegis-system -- aegis -l

{"secrets":[]}{% endraw %}
```

All right, our cluster is as clean as a baby’s butt; let’s move on.

## Read the Source

Make sure [you examine the manifests][workload-yaml] to gain an understanding
of what kinds of entities you’ve deployed to your cluster.

[workload-yaml]: https://github.com/shieldworks/aegis/tree/main/examples/workload-using-sdk/k8s

## The Benefit of Using **Aegis SDK**

**Aegis SDK** gives direct control of **Aegis Safe** to your workload.

The advantage of this approach is: you are in charge.
The downside of it is: Well, you are in charge 🙂.

But, jokes aside, your application will have to be
more tightly bound to **Aegis** without a sidecar.

However, when you use a sidecar, your application does not have any idea of
**Aegis**’s existence. From its perspective, it is merely reading from a file
that something magically updates every once in a while. This
“*separation of concerns*” can make your application architecture more
adaptable to changes.

As in anything, there is no one true way to do it. Your approach will depend
on your project’s requirements.

## Deploying the Demo Workload

That part taken care of; let’s deploy a workload that uses **Aegis SDK** 
instead of *Aegis Sidecar*.

```bash 
# Switch to the Aegis repo:
cd $WORKSPACE/aegis
# Install the workload:
make example-sdk-deploy
# If you are building from the source, 
# use `make example-sdk-deploy-local` instead.
```

And that’s it. You have your demo workload up and running.

## The Demo App

[Here is the source code of the demo container’s app][workload-src] for the
sake of completeness.

[workload-src]: https://github.com/shieldworks/aegis/blob/main/examples/workload-using-sdk/main.go

When you check the source code, you’ll see that our demo app tries to get the 
secret by querying the SDK via `sentry.Fetch()`, displays the secret if it finds
and repeats this every 5 seconds in an infinite loop.

```go 
for {
    log.Println("fetch")
    d, err := sentry.Fetch()
    
    // … (error handling) truncated …

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

NAME                                  READY   STATUS    RESTARTS   AGE
example-85bdbc4cf4-6n2ng  1/1     Running   0          9s{% endraw %}
```

Let’s check the logs of our pod:

```bash 
{% raw %}kubectl logs example-85bdbc4cf4-6n2ng -f

2023/05/13 21:27:41 fetch
2023/05/13 21:27:48 [TRACE] P4jKsYvr Sentry:Fetch 
https://aegis-safe.aegis-system.svc.cluster.local:8443/workload/v1/secrets
Failed to read the secrets file. Will retry in 5 seconds…
Secret does not exist
2023/05/13 21:27:53 fetch
2023/05/13 21:27:53 [TRACE] ovJCKy02 Sentry:Fetch 
https://aegis-safe.aegis-system.svc.cluster.local:8443/workload/v1/secrets

…{% endraw %}
```

We don’t have any secrets registered to our workload as expected. So, let’s
add some.

## Registering a Secret

Let’s register a secret and see how the logs change:

```bash 
{% raw %}# Find the name of the Aegis Sentinel pod:
kubectl get po -n aegis-system

# register a secret to our workload using Aegis Sentinel:
kubectl exec aegis-sentinel-778b7fdc78-86v6d -n aegis-system -- aegis \
  -w "example" \
  -s "AegisRocks!"
  
# Response: 
# OK{% endraw %}
```

Now let’s check the logs again:

```bash 
{% raw %}kubectl logs example-85bdbc4cf4-6n2ng -f

secret: updated: "2023-05-13T21:37:00Z", created: "2023-05-13T21:37:00Z", 
value: AegisRocks!
2023/05/13 21:37:11 fetch
2023/05/13 21:37:11 [TRACE] t0eK8Ecg Sentry:Fetch 
https://aegis-safe.aegis-system.svc.cluster.local:8443/workload/v1/secrets
secret: updated: "2023-05-13T21:37:00Z", created: "2023-05-13T21:37:00Z", 
value: AegisRocks!

…{% endraw %}
```

[demo-sidecar]: /examples/use-case-sidecar

## Conclusion

So, yay 🎉.  We got our secret; and since we use **Aegis SDK**, we also were able
to get additional metadata such as the creation and modification timestamps of
our secret, which was not possible to retrieve if we used the
[**Aegis Sidecar**][demo-sidecar] approach that we have seen in the
[former tutorial][demo-sidecar].

Next up, you’ll learn about **Aegis Init Container**.
