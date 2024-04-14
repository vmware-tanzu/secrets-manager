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

title: Transforming Secrets
layout: post
prev_url: /docs/use-case-encryption/
permalink: /docs/use-case-transformation/
next_url: /docs/use-case-secrets-as-volumes/
---

## Situation Analysis

Sometimes your workload expects the secret in a different format than it has been
initially provided. You don't want to write custom code to do the transformation.

For this reason, **VMware Secrets Manager** provides a way to interpolate and
transform secrets. You can provide a template to transform the secret into the
desired format (e.g., `JSON`, `YAML`, or free-form text).

## Strategy

Use **VSecM Sentinel** to register a secret; use the `-t` flag to provide a
template to transform the secret into the desired format. The workload will
consume the transformed secret.

## High-Level Diagram

Open the image in a new tab to see the full-size version:

![High-Level Diagram](/assets/transform-secrets.jpg "High-Level Diagram")

## Implementation

This tutorial will show various way you can interpolate and transform secrets.

### Installing An Example Workload

You can use **VSecM Inspector** from **VSecM** Docker registry to test and
validate the transformations:

Here is a `Deployment` manifest to deploy **VSecM Inspector**:

```yaml
{% raw %}# ./Deployment.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
  name: vsecm-inspector
  namespace: default
  labels:
    app.kubernetes.io/name: vsecm-inspector
spec:
  replicas: 1
  selector:
    matchLabels:
      app.kubernetes.io/name: vsecm-inspector
  template:
    metadata:
      labels:
        app.kubernetes.io/name: vsecm-inspector
    spec:
      serviceAccountName: vsecm-inspector
      containers:
      - name: main
        image: vsecm/vsecm-inspector:latest
        volumeMounts:
        - name: spire-agent-socket
          mountPath: /spire-agent-socket
          readOnly: true
      volumes:
      - name: spire-agent-socket
        csi:
          driver: "csi.spiffe.io"
          readOnly: true{% endraw %}
```

Here is a `ServiceAccount` that you can use with the above `Deployment`:

```yaml
{% raw %}# ./ServiceAccount.yaml

apiVersion: v1
kind: ServiceAccount
metadata:
  name: vsecm-inspector
  namespace: default
automountServiceAccountToken: false{% endraw %}
```

And here is a `ClusterSPIFFEID` that you can use with these:

```yaml
{% raw %}# ./Identity.yaml

apiVersion: spire.spiffe.io/v1alpha1
kind: ClusterSPIFFEID
metadata:
  name: vsecm-inspector
spec:
  spiffeIDTemplate: "spiffe://vsecm.com\
    /workload/example\
    /ns/default\
    /sa/example\
    /n/{{ .PodMeta.Name }}"
  podSelector:
    matchLabels:
      app.kubernetes.io/name: vsecm-inspector
  workloadSelectorTemplates:
  - "k8s:ns:default"
  - "k8s:sa:vsecm-inspector"{% endraw %}
```

You can deploy these manifests with the following commands:

```bash
kubectl apply -f Deployment.yaml
kubectl apply -f ServiceAccount.yaml
kubectl apply -f Identity.yaml
```

### Preparation

Let us define a few aliases first, they will speed things up:

```bash 
SENTINEL=$(kubectl get po -n vsecm-system \
  | grep "vsecm-sentinel-" | awk '{print $1}')
SAFE=$(kubectl get po -n vsecm-system \
  | grep "vsecm-safe-" | awk '{print $1}')
WORKLOAD=$(kubectl get po -n default \
  | grep "vsecm-inspector-" | awk '{print $1}')

# Delete secrets assigned to the workload:
alias delete-secret="kubectl exec $SENTINEL \
  -n vsecm-system -- safe \
  -w example -s x -d"

alias inspect="kubectl exec $INSPECTOR -- ./env"
```

Now, we can start experimenting.

### Cleanup

Let's start with a blank slate again:

```bash 
delete-secret
# Output: OK

inspect
# Output:
# Failed to fetch the secrets. Try again later.
# Secret does not exist
```

### The Format (`-f`) Argument

**VSecM Sentinel** CLI accepts a format flag (`-f`), the possible values are

* `"json"`
* and `"yaml"`

If it is not given, it defaults to `"json"`; however, in the upcoming examples
we'll be explicit and provide this argument at all times.

### Registering a JSON Secret

```bash
{% raw %}kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s '{"username": "admin", "password": "VSecMRocks!"}' \
  -f json

inspect
# Output:
# {"username": "admin", "password": "VSecMRocks!"}{% endraw %}
```

### Registering a YAML Secret

```bash 
{% raw %}kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s '{"username": "admin", "password": "VSecMRocks!"}' \
  -f yaml

inspect
# Output:
# password: VSecMRocks!
# username: admin{% endraw %}
```

### Registering a JSON String (with invalid JSON)

Now we'll deliberately make an error in our JSON. Notice the missing `"`
in `username"`: That is not valid JSON.

```bash
{% raw %}kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s '{username": "admin", "password": "VSecMRocks!"}' \
  -f json

inspect
# Output:
# {username": "admin", "password": "VSecMRocks!"}{% endraw %}
```

### Registering a YAML String (with invalid JSON)

Since the JSON cannot be parsed, the output will not be a YAML:

```bash 
{% raw %}kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s '{username": "admin", "password": "VSecMRocks!"}' \
  -f yaml

inspect
# Output:
# {username": "admin", "password": "VSecMRocks!"}{% endraw %}
```

### Transforming A JSON Secret

```bash
{% raw %}kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s '{"username": "admin", "password": "VSecMRocks!"}' \
  -t '{"USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f json

inspect
# Output:
# {"USR":"admin", "PWD":"VSecMRocks!"}{% endraw %}
```

### Transforming a YAML Secret

```bash
{% raw %}kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s '{"username": "admin", "password": "VSecMRocks!"}' \
  -t '{"USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f yaml

inspect
# Output:
# PWD: VSecMRocks!
# USR: admin{% endraw %}
```

### Transforming a JSON Secret (invalid JSON)

If our secret is not valid JSON, then the YAML transformation will not be
possible. **VMware Secrets Manager** will still try its best to provide something.

```bash
{% raw %}kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s '{username": "admin", "password": "VSecMRocks!"}' \
  -t '{"USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f json

inspect
# Output:
# {username": "admin", "password": "VSecMRocks!"}{% endraw %}
```

### Transforming a JSON Secret (invalid template)

Since template is not valid, the template transformation will not happen.

```bash
{% raw %}kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s '{"username": "admin", "password": "VSecMRocks!"}' \
  -t '{USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f json

inspect
# Output:
# {"username": "admin", "password": "VSecMRocks!"}{% endraw %}
```

### Transforming a JSON Secret (invalid template and JSON)

**VMware Secrets Manager** will still try its best:

```bash
{% raw %}kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s '{username": "admin", "password": "VSecMRocks!"}' \
  -t '{USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f json

inspect
# Output:
# {username": "admin", "password": "VSecMRocks!"}{% endraw %}
```

### Transforming YAML Secret (invalid JSON)

```bash
{% raw %}kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s '{username": "admin", "password": "VSecMRocks!"}' \
  -t '{"USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f yaml

inspect
# Output
# {username": "admin", "password": "VSecMRocks!"}{% endraw %}
```

### Transforming YAML Secret (invalid template)

```bash
{% raw %}kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s '{"username": "admin", "password": "VSecMRocks!"}' \
  -t '{USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f yaml

inspect
# Output:
# {USR":"admin", "PWD":"VSecMRocks!"}{% endraw %}
```

### Transforming YAML Secret (invalid JSON and template)

```bash
{% raw %}kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s '{username": "admin", "password": "VSecMRocks!"}' \
  -t '{USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f yaml

inspect
# Output:
# {username": "admin", "password": "VSecMRocks!"}{% endraw %}
```

## Conclusion

**VMware Secrets Manager**, demonstrates a robust capability for the transformation 
and management of secrets in different formats, accommodating various 
organizational needs. By employing templates, users can easily convert secrets 
into the required formats such as JSON or YAML, ensuring seamless integration 
with various workloads and systems.

This flexibility not only simplifies secret management but also enhances security 
protocols without necessitating bespoke coding solutions, thereby streamlining 
operations and reducing potential errors. 

This use case showcases **VSecM**'s adaptability and strength, making it an 
invaluable asset for secure and efficient secret management in modern Cloud 
Native environments.

<p class="github-button">
  <a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0240-transform.md">
     Suggest edits ✏️
  </a>
</p>