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

title = "Deploy an Example Workload"
weight = 30
+++

## Deploy an Example Workload

Now, let's deploy an example workload to the cluster to test
**VMware Secrets Manager** in action.

```bash
cd $WORKSPACE/secrets-manager
export VSECM_VERSION=latest
make example-sdk-deploy
```

This will take a few moments too.

When done you would ba able to list the pods in the `default` namespace:

```bash
kubectl get po -n default
```

```txt
# Output
NAME                       READY   STATUS    RESTARTS   AGE
example-6cbb96b768-dhm7c   1/1     Running   0          6m44s
```

Let's check the logs of our example workload:

```bash
kubectl logs example-6cbb96b768-dhm7c
```

The output will be something similar to this:

```txt
2024/03/25 17:30:32 fetch
2024/03/25 17:30:32 [TRACE] RpyyPfqx Sentry:Fetch 
https://vsecm-safe.vsecm-system.svc.cluster.local:8443/workload/v1/secrets
2024/03/25 17:30:32 [TRACE] RpyyPfqx Sentry:Fetch svid:id:  
spiffe://vsecm.com/workload/example/ns/default/sa/example/n
/example-6cbb96b768-dhm7c
Failed to read the secrets file. Will retry in 5 seconds...
Secret does not exist
2024/03/25 17:30:37 fetch
2024/03/25 17:30:37 [TRACE] kUWlDyo3 Sentry:Fetch 
https://vsecm-safe.vsecm-system.svc.cluster.local:8443/workload/v1/secrets
2024/03/25 17:30:37 [TRACE] kUWlDyo3 Sentry:Fetch svid:id:
spiffe://vsecm.com/workload/example/ns/default/sa/example/n/
example-6cbb96b768-dhm7c
Failed to read the secrets file. Will retry in 5 seconds...
Secret does not exist
2024/03/25 17:30:42 fetch
2024/03/25 17:30:42 [TRACE] dorrPbVN Sentry:Fetch 
https://vsecm-safe.vsecm-system.svc.cluster.local:8443/workload/v1/secrets
2024/03/25 17:30:42 [TRACE] dorrPbVN Sentry:Fetch svid:id:  
spiffe://vsecm.com/workload/example/ns/default/sa/example/n/
example-6cbb96b768-dhm7c
Failed to read the secrets file. Will retry in 5 seconds...
Secret does not exist
... truncated ...
```

Our sample workload is trying to fetch a secret, but it can't find it.

Here's the source code of our sample workload to provide some contxt:

```go
package main

// ... truncated headers ...

func main() {

	// ... truncated irrelevant code ...

	for {
		log.Println("fetch")
		d, err := sentry.Fetch()

		if err != nil {
			fmt.Println("Failed. Will retry in 5 seconds...")
			fmt.Println(err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		if d.Data == "" {
			fmt.Println("No secret yet... will check again later.")
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

`sentry.Fetch()` is a function provided by the **VMware Secrets Manager**;
it establishes a secure mTLS connection between the workload and
**VSecM Safe** to fetch the secret.

Since this workload does not have any secret registered, the request fails
and the workload retries every 5 seconds.

Since this is a quickstart example, we won't dive into the details of
how the workload establishes a secure mTLS connection with the
**VSecM Safe**. We'll cover this in the following sections.

For the sake of this quickstart, we can assume that secure communication
between the workload and the **VSecM Safe** is already taken
care of for us.

{{ edit() }}
