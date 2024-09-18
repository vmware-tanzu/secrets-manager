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

title = "Mounting Kubernetes Secrets"
weight = 244
+++

## Situation Analysis

This use case will mount **VSecM**-minted secrets as environment variables or files 
through Kubernetes Secrets. This approach is useful when deploying workloads using 
Helm charts that only allow creating secrets as Kubernetes secrets.

When you deploy a workload using a Helm chart and the Helm chart only allows you 
to create secrets as Kubernetes secrets, approaches like 
[Fetching Secrets Using VSecM Sidecar][usecase-sidecar], 
[Fetching Secrets Using VSecM SDK][usecase-sdk], or even 
[Mutating an In-Memory Template File][use-case-template] could be unfeasible.

[usecase-sidecar]: @/documentation/use-cases/sidecar.md
[usecase-sdk]: @/documentation/use-cases/vsecm-sdk.md
[use-case-template]: @/documentation/use-cases/template-files.md

Luckily, **VMware Secrets Manager** can create and manage native Kubernetes 
Secrets in any namespace.

## Screencast

Here is a screencast that demonstrates this use case:

<script 
  src="https://asciinema.org/a/676194.js" 
  id="asciicast-676194" 
  async="true"></script>

## High-Level Diagram

Open the image in a new tab to see the full-size version:

![High-Level Diagram](/assets/mount-secrets.png "High-Level Diagram")

## Implementation

Using a **VSecM**-generated Kubernetes `Secret` is no different than how you would
use a regular Kubernetes `Secret`. This section will show you how to do it.

> **Environment Variables or Files**
> 
> You can use the generated Kubernetes `Secret` in the usual ways a Kubernetes
> `Secret` is used. For example, you can bind environment variables from the
> `Secret` to your container or create a volume from the generated Kubernetes
> `Secret`. **VSecM** will create a native Kubernetes `Secret` for you;
> how you use it will be up to you.


You can bind environment variables from the secret to your container or create 
a volume from the generated Kubernetes Secret.

### Waiting for the Secret to be Ready

Once the workload is initialized, the secret must be ready for the workload to 
consume it. There are several strategies that you can follow for this.

* Using a custom [Kubernetes Operator][k8s-operator], you can configure your 
  orchestrator to deploy the workload only after the secret is created.
* You can use a toolkit like [Carvel][carvel] to create a dependency relationship 
  to guarantee the Secrets are created before the workloads.
* You can use the [**VSecM Init Container**][vsecm-init-container] similarly to how 
  it's described in the [Mounting Secrets as Volumes][use-case-mount-volumes] use 
  case. For example, in an `initCommand:` stanza, you can first create a Kubernetes 
  `Secret`, wait for the cluster to reconcile, then register a dummy secret to the 
  workload to trigger a VSecM Init Container that the workload has to receive and 
  initialize the main container (*see the sample in the following section*).

> **Don't Want the Extra Wait?**
> 
> **VSecM Init Container** could also wait for the existence of a Kubernetes 
> `Secret`--if you help us implement the feature. This way, you won't have to add 
> a `wait:` command that may fail depending on how busy the cluster is. 
> 
> Once this feature is implemented, the **VSecM Init Container** will only initialize 
> the workload after ensuring the Kubernetes Secret is there and populated.
> 
> If you need this feature, your contributions are welcome:
> 
> * [VSecM Init Container shall be able to wait on a Kubernetes secret `#763`][ticket-763]

In this walkthrough we will assume that the VSecM-generated Kubernetes `Secret`
will be created before the workload begins its lifecycle.

[ticket-763]: https://github.com/vmware-tanzu/secrets-manager/issues/763
[k8s-operator]: https://kubernetes.io/docs/concepts/extend-kubernetes/operator/
[carvel]: https://carvel.dev
[vsecm-init-container]: @/documentation/use-cases/init-container.md
[use-case-mount-volumes]: @/documentation/use-cases/mounting-secrets.md

### Creating a Kubernetes Secret

[**VSecM** Helm Charts][helm-charts] can be configured to provide secure 
pattern-based random secrets for workloads. For example, we can modify the 
`initCommand` stanza of sentinel Helm chart's `values.yaml` to register a 
random username and password as secrets for our example workload.

[helm-charts]: https://github.com/vmware-tanzu/secrets-manager/tree/gh-pages

To create a Kubernetes `Secret`, prefix the name of the workload with `k8s:` 
similar to the following example:

```yaml
initCommand:
  enabled: true
  command: |
    --
    w:k8s:example-secret
    n:example-apps
    s:gen:{"username":"admin-[a-z0-9]{6}","password":"[a-zA-Z0-9]{12}"}
    t:{"ADMIN_USER":"{{.username}}","ADMIN_PASS":"{{.password}}"}
    --
    wait:5000
    --
    w:example
    n:example-apps
    s:init
    --
```

Note the `k8s:` prefix in the workload name. This prefix tells the **VSecM**
o create a Kubernetes `Secret` instead of a VSecM "secret".

The above stanza is equivalent to this [**VSecM Sentinel** command][vsecm-cli]:

```yaml
kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w "k8s:example-secret" \
  -s 'gen:{"username":"admin-[a-z0-9]{6}","password":"[a-zA-Z0-9]{12}"}' \
  -t '{"ADMIN_USER":"{{.username}}","ADMIN_PASS":"{{.password}}"}'

sleep(5)

kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w "example" \
  -s 'init'
```

Aside from the regular flags such as `w:` for `-w`, `n:` for `-n`, and `s:` for `-s`,
the `initCommand` accepts two special commands:

* `wait:$timInMilliseconds` to wait for a specified time before executing the next command.
* `exit:true` break execution without processing any further commands.

[vsecm-cli]: @/documentation/usage/cli.md

### Don't Forget to Create a `ClusterSPIFFEID`

For the **VSecM Init Container** to talk to **VSecM Safe**, your example workload will 
need a `ClusterSPIFFEID` similar to the following:

```yaml
apiVersion: spire.spiffe.io/v1alpha1
kind: ClusterSPIFFEID
metadata:
  name: example
spec:
  className: vsecm
  spiffeIDTemplate: "spiffe://vsecm.com\
    /workload/example\
    /ns/{{ .PodMeta.Namespace }}\
    /sa/{{ .PodSpec.ServiceAccountName }}\
    /n/{{ .PodMeta.Name }}"
  podSelector:
    matchLabels:
      app: example-app
  workloadSelectorTemplates:
    - "k8s:ns:example-apps"
    - "k8s:sa:example-sa"
```

You can check out [Mounting Secrets as Volumes][use-case-mount-volumes] for a 
separate example of how this `ClusterSPIFFEID` is created and evaluated.

### That's It

If you followed these steps, your workload would happily and securely consume 
**VSecM**-minted secure random secrets.

## Conclusion

In conclusion, integrating **VMware Secrets Manager** with Kubernetes `Secret`s 
offers a robust and flexible solution for managing and consuming secrets within 
your Kubernetes workloads. 

By leveraging VSecM to generate and manage Kubernetes `Secret`s, you can ensure 
that your applications are provisioned with the necessary credentials, configurations, 
and other sensitive data in a secure, automated fashion. The approach outlined 
provides a clear path for securely injecting secrets into workloads, whether 
through environment variables or files.

Additionally, implementing strategies such as waiting for secrets to be ready 
before workload deployment, using **VSecM Init Container**s, and configuring 
`ClusterSPIFFEID`s further enhances the security and reliability of the secret 
management process.

Contributions to further enhancing these capabilities, such as implementing new 
features in VSecM Init Containers, are encouraged and welcome.

## List of Use Cases in This Section

{{ use_cases() }}

{{ edit() }}
