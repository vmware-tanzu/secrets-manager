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

title = "Verify Installation"
weight = 15
+++

## Verify Installation

Let's check that the installation was successful by listing the pods int
the `spire-system` , `vsecm-system` and `keystone-system` namespaces:

```bash
kubectl get po -n spire-system
```

```txt
# Output:
NAME                          READY   STATUS    RESTARTS      
spire-agent-wdhdh             3/3     Running   0             
spire-server-b594bdfc-szssx   2/2     Running   0             
```

```bash
kubectl get po -n vsecm-system
```

```txt
# Output:
NAME                              READY   STATUS    RESTARTS
vsecm-keystone-c54d99d7b-c4jk4    1/1     Running   0          
vsecm-safe-6cc477f58f-x6wc9       1/1     Running   0          
vsecm-sentinel-74d648675b-8zdn2   1/1     Running   0                   
```

All the pods look up and running, so we can move on to the next step.

## List Available Commands

`make help` lists all the available make targets in a cheat sheet format:

```bash
make help

# The output will vary as we add more commands to the Makefile.
# It will contain useful information about the available commands.
```

## Troubleshooting

Here are some common issues and how to resolve them:

### My Workload Cannot Get a Secret Registered for It

First make sure that the workload has a `ClusterSPIFFEID` created for it:

```bash
kubectl get clusterspiffeid

# Output:
#
# NAME                           AGE
# example                        20h
# spire-server-spire-test-keys   20h
# vsecm-keystone                 20h
# vsecm-safe                     20h
# vsecm-sentinel                 20h
```

Let's say the workload is `example`. Now let's describe its `ClusterSPIFFEID`
to see if it selects the Pod.

```bash
kubectl describe clusterspiffeid example

# Output:
#
# Name:         example
# Namespace:    
# Labels:       <none>
# Annotations:  <none>
# API Version:  spire.spiffe.io/v1alpha1
# Kind:         ClusterSPIFFEID
# Metadata:
#   Creation Timestamp:  2024-07-14T05:12:10Z
#   Generation:          1
#   Resource Version:    1901
#   UID:                 7ef51eff-4afc-421c-9199-a5d1934c4b63
# Spec:
#   Class Name:  vsecm
#   Pod Selector:
#     Match Labels:
#       app.kubernetes.io/name:  example
#   Spiffe ID Template:          spiffe://vsecm.com/workload/example\
#   /ns/{{ .PodMeta.Namespace }}/\
#   sa/{{ .PodSpec.ServiceAccountName }}/n/{{ .PodMeta.Name }}
#   Workload Selector Templates:
#     k8s:ns:default
#     k8s:sa:example
# Status:
#   Stats:
#     Entries Masked:             0
#     Entries To Set:             0
#     Entry Failures:             0
#     Namespaces Ignored:         2
#     Namespaces Selected:        7
#     Pod Entry Render Failures:  0
#     Pods Selected:              1
# Events:                         <none>
```

`Pods Selected: 1` indicates that the `ClusterSPIFFEID` is selecting the Pod.

If the workload still cannot get a secret, check the following then check the
name of the mounted socket for SPIFFE. By default, **VSecM**-managed **SPIRE**
uses `unix:///spire-agent-socket/spire-agent.sock` for the agent socket. If the
socket is named differently, the workload will not be able to communicate with
the **SPIRE** agent.

For example, your pd might have a socket named `spire-agent.sock` but the
workload is seeking `agent.sock`. In this case, the workload will not be able to
communicate with the **SPIRE** agent. Without this communication it won't get
its SVID, and without an SVID it won't be able to get a secret.

Here is a sample manifest that attaches the socket to the workload:

```yaml
 apiVersion: v1
 kind: Pod
 metadata:
   name: my-app
 spec:
   containers:
   - name: my-app
     image: "my-app:latest"
     imagePullPolicy: Always

     volumeMounts:
     
     # (2) Mount the volume to the path where the socket is expected.
     # The path of the socket will be:
     # `/spiffe-workload-api/spire-agent.sock`.
     - name: spiffe-workload-api
       mountPath: /spiffe-workload-api
       readOnly: true

     resources:
       requests:
         cpu: 200m
         memory: 32Mi
       limits:
         cpu: 500m
         memory: 64Mi

   volumes:

    # (1) Define a volume that mounts the socket
    # using the CSI driver.
    - name: spiffe-workload-api
      csi:
        driver: "csi.spiffe.io"
        readOnly: true
```

If you browse this website, you will find a lot of examples that shows
how to mount the socket to the workload. In addition, you can check out 
the [Examples folder of the **VSecM** repository][examples] for more examples.

[examples]: https://github.com/vmware-tanzu/secrets-manager/tree/main/examples

### I Have an OpenShift Cluster and I Get All Kinds of Errors

If you are using an OpenShift cluster, you will need to defined
OpenShift-specific security policies. The easiest way to do this is to 
set the `global.enableOpenShift` to `true` and use [**VSecM** Helm 
Charts][helm-charts] to install **VSecM**.

[helm-charts]: https://artifacthub.io/packages/helm/vsecm/vsecm

## Still Facing Issues?

If you are still having issues, it would be useful to check the logs of
**VSecM Safe**, **VSecM Sentinel**, **VSecM Keystone** pods in the `vsecm-system`
namespace; **SPIRE Server** in the `spire-server` namespace; and **SPIRE Agent**
in the `spire-agent` namespace.

By default, all logs are set up at the highest verbosity level, so you should
be able to get clues from the logs about what is going wrong.

If you are still facing issues, please reach out to us at
[the official **VSecM** Slack channel][slack].

[slack]: https://join.slack.com/t/a-101-103-105-s/shared_invite/zt-287dbddk7-GCX495NK~FwO3bh_DAMAtQ

{{ edit() }}
