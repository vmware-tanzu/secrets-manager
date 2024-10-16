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

title = "QuickStart"
weight = 11
+++

## VSecM In Action

This is a recording that demonstrates how to register secrets to 
**VSecM** using the VSecM CLI.

<script 
  src="https://asciinema.org/a/676191.js" 
  id="asciicast-676191"
  async="true"></script>

The following sections outline various highlights of the recording.

## Prerequisites

Make sure you have installed the following on your system:

* Minikube
* Make
* Docker
* Kubectl
* Helm

## Makefile Targets

* `make k8s-delete`: Deletes the Kubernetes cluster.
* `make k8s-start`: Starts the Kubernetes cluster.
* `make help`, and `make h`: Displays help about various `make` targets.

## Installing **VSecM** Using Helm

It's the easiest way to get started with **VSecM**:

```bash
helm repo add vsecm https://vmware-tanzu.github.io/secrets-manager/
helm repo update
helm install vsecm vsecm/vsecm --version 0.26.1
```

## Installing **VSecM** Using the `make` Targets

This is useful when you want to contribute to the source code, and you
want to build everything from the source:

```bash
git clone https://github.com/vmware-tanzu/secrets-manager.git
cd secrets-manager
make k8s-delete
make k8s-start
eval $(minikube -p minikube docker-env)
make build-local
make deploy-local
```

## Ensuring Everything is Running

```bash
kubectl get po -n spire-server
kubectl get po -n spire-system
kubectl get po -n vsecm-system
kubectl get clusterspiffeid
```

## Deploying a Demo Workload

```bash
cd ./examples/using_vsecm_inspector
kubectl apply -f .
```

## Registering Secrets to the Workload

```bash
# Find vsecm-sentinel:
kubectl get po -n vsecm-system
kubectl exec vsecm-sentinel-c6cf9f894-j9vfq -n vsecm-system \
-- safe \
-w example \
-s VSecMRocks \
-n default
```

* -n: Namespace
* -w: Name of the workload
* -s: The secret assigned to the workload

The name of the workload is provided by its `ClusterSPIFFEID`:

```bash
# cat ./examples/using_vsecm_inspector/Identity.yaml

apiVersion: spire.spiffe.io/v1alpha1
kind: ClusterSPIFFEID
metadata:
  name: vsecm-inspector
spec:
  className: "vsecm"
  spiffeIDTemplate: "spiffe://vsecm.com\
    /workload/example\
    /ns/{{ .PodMeta.Namespace }}\
    /sa/{{ .PodSpec.ServiceAccountName }}\
    /n/{{ .PodMeta.Name }}"
  podSelector:
    matchLabels:
      app.kubernetes.io/name: vsecm-inspector
  workloadSelectorTemplates:
    - "k8s:ns:default"
    - "k8s:sa:vsecm-inspector"
```

The name is `examples` in `/workload/example` in the `spiffeIDTemplate`.

## Verifying the Secret Has Been Registered

```bash
# Find the Workload:
kubectl get po 
kubectl exec vsecm-inspector-695d68875f-wxmfm -- ./env
# Output:
# VSecMRocks
```

## Conclusion

This was a quick overview of some basic operations with **VSecM**.

Check out rest of this documentation for more examples and use cases.

And keep your secrets... secret.

<p>&nbsp;</p>

[«« Back to Showcase](@/showcase/vsecm.md)
