---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, sentinel, cli, bastion
title: Aegis Sentinel CLI
description: <strong>Aegis Sentinel</strong> command line interface
micro_nav: true
page_nav:
  prev:
    content: <strong>Aegis</strong> go SDK
    url: '/docs/sdk'
  next:
    content: <strong>Aegis</strong> deep dive
    url: '/docs/architecture'
---

<p style="text-align:right;position:relative;top:-40px;"
><a href="https://github.com/ShieldWorks/aegis-web/blob/main/docs/sentinel.md"
style="border-bottom: none;background:#e0e0e0;padding:0.5em;display:inline-block;
border-radius:8px;">
edit this page on <strong>GitHub</strong> ✏️</a></p>

## Introduction

This section contains usage examples and documentation for [**Aegis Sentinel**][sentinel].

[sentinel]: https://github.com/ShieldWorks/aegis/tree/main/app/sentinel

## Finding **Aegis Sentinel**

First, find which pod belongs to **Aegis System**:

```bash 
kubetctl get po -n aeis-system
```

The response to the above command will be similar to the following:

```text 
NAME                              READY
aegis-safe-5f6948c84c-vkrdh       1/1
aegis-sentinel-5998b5dbfc-lvw44   1/1
```

There, `aegis-sentinel-5998b5dbfc-lvw44` is the name of the Pod you’d need.

You can also execute a script similar to the following to save the Pod’s name
into an environment variable:

```bash 
SENTINEL=$(kubectl get po -n aegis-system \
  | grep "aegis-sentinel-" | awk '{print $1}')
```

In the following examples, we’ll use `$SENTINEL` in lieu of the **Aegis Sentinel**’s
Pod name.

## Displaying Help Information

```bash 
kubectl exec $SENTINEL -n aegis-system -- aegis --help
```

Output:

```text 
usage: aegis [-h|--help] [-l|--list] [-k|--use-k8s] [-d|--delete] [-a|--append]
             [-n|--namespace "<value>"] [-i|--input-keys "<value>"] [-b|--store
             "<value>"] [-w|--workload "<value>"] [-s|--secret "<value>"]
             [-t|--template "<value>"] [-f|--format "<value>"] [-e|--encrypt]

             Assigns secrets to workloads.

Arguments:

  -h  --help        Print help information
  -l  --list        lists all registered workloads.. Default: false
  -k  --use-k8s     update an associated Kubernetes secret upon save. Overrides
                    AEGIS_SAFE_USE_KUBERNETES_SECRETS.. Default: false
  -d  --delete      delete the secret associated with the workload.. Default:
                    false
  -a  --append      append the secret to the existing secret collection
                    associated with the workload.. Default: false
  -n  --namespace   the namespace of the Kubernetes Secret to create.. Default:
                    default
  -i  --input-keys  A string containing the private and public Age keys and AES
                    seed, each separated by '\n'.
  -b  --store       backing store type (file|memory) (default: file). Overrides
                    AEGIS_SAFE_BACKING_STORE.
  -w  --workload    name of the workload (i.e. the '$name' segment of its
                    ClusterSPIFFEID
                    ('spiffe://trustDomain/workload/$name/…')).
  -s  --secret      the secret to store for the workload.
  -t  --template    the template used to transform the secret stored.
  -f  --format      the format to display the secrets in. Has effect only when
                    `-t` is provided. Valid values: yaml, json, and none.
                    Defaults to none.
  -e  --encrypt     returns an encrypted version of the secret if used with
                    `-s`; decrypts the secret before registering it to the
                    workload if used with `-s` and `-w`.. Default: false
```

Note that based on your **Aegis Sentinel** version the output of the above
command can be slightly different.

## Registering a Secret for a Workload

Given our workload has the SPIFFE ID 
`"spiffe://aegis.ist/workload/billing/: …[trunacted]"`

```bash
kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w billing \
  -s "very secret value"
```

will register the secret `"very secret value"` to `billing`.

## Registering Multiple Secrets

You can use the `-a` (*append*) argument to register more than one secret
to a workload.

```bash
kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w billing \
  -s "first part of the token" \
  -a

kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w billing \
  -s "second part of the token" \
  -a
```

## Encrypting a Secret

Use the `-e` flag to encrypt a secret.

```bash
kubectl exec $SENTINEL -n aegis-system -- aegis \
  -s "very secret value" \
  -e 

# The above command outputs an encrypted string that can be
# securely stored anywhere, including source control systems.
```

## Registering an Encrypted Secret

Again, `-e` flag can be used to register an encrypted secret to a workload:

```bash
kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w billing \
  -s $ENCRYPTED_SECRET \
  -e
```

## Deleting a Secret

To delete the secret associated to a workload, use the `-d` flag:

```bash
kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w billing \
  -d
```

## Choosing a Backing Store

The registered secrets will be encrypted and backed up to 
**Aegis Safe**’s Kubernetes volume by default. This behavior can be configured
by changing the `AEGIS_SAFE_BACKING_STORE` environment variable that
**Aegis Safe** sees. In addition, this behavior can be overridden on a per-secret
basis too.

The following commands stores the secret to the backing volume 
(*default behavior*):

```bash
kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w billing \
  -s "very secret value" \
  -b file
```

This one, will **not** store the secret on file; the secret will only be 
persisted in memory, and will be lost if **Aegis Sentinel** needs to restart:

```bash
kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w billing \
  -s "very secret value" \
  -b memory
```

The following will stored the secret on the cluster itself as a Kubernetes
`Secret`. The value of the secret will be **encrypted** with the public key
of **Aegis Safe** before storing it on the Secret. 

```bash
kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w billing \
  -s "very secret value" \
  -b cluster
```

## Template Transformations

You can transform how the stored secret is displayed to the consuming workloads:

```bash 
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  -w "example" \
  -s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
  -t '{"USER":"{{.username}}", "PASS":"{{.password}}", "VALUE": "{{.value}}"}'{% endraw %}
```

When the workload fetches the secret through the workload API, this is what
it sees as the value:

```text 
{"USER": "root", "PASS": "SuperSecret", "VALUE": "AegisRocks"}
```

Instead of this default transformation, you can output it as `yaml` too:

```bash 
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  -w "example" \
  -s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
  -t '{"USER":"{{.username}}", "PASS":"{{.password}}", "VALUE": "{{.value}}"}' \
  -f yaml{% endraw %}
```

The above command will result in the following secret value to the workload 
that receives it:

```text 
USER: root
PASS: SuperSecret
VALUE: AegisRocks
```

> **`"json"` Is the Default Value**
> 
> If you don’t specify the `-f` flag, it will default to `"json"`.

You can create a YAML secret value without the `-t` flag too. In that case
**Aegis Safe** will assume an identity transformation:

```bash 
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  -w "example" \
  -s '{"username": "root", "password": "SuperSecret", "value": "AegisRocks"}' \
  -f yaml{% endraw %}
```

The above command will result in the following secret value to the workload
that receives it:

```text 
username: root
password: SuperSecret
value: AegisRocks
```

If you provide `-f json` as the format argument, the secret will have to be
a strict JSON. Otherwise, **Aegis Sentinel** will try to come up with a 
reasonable value, and not raise an error; however the output will likely be
in a format that the workload is not expecting.

> **Gotcha**
> 
> If `-t` is given, the `-s` argument will have to be a valid JSON regardless
> of what is chosen for `-f`.

The following is also possible:

```bash 
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  -w "example" \
  -s 'USER»»»{{.username}}'{% endraw %}
```

and will result in the following as the secret value for the workload:

```text 
USER»»»root
```

Or, equivalently:

```bash 
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  -w "example" \
  -s 'USER»»»{{.username}}' \
  -f json{% endraw %}
```

Will provide the following to the workload:

```text 
USER»»»root
```

Similarly, the following will **not** raise an error:

```bash 
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  -w "example" \
  -s 'USER»»»{{.username}}' \
  -f yaml{% endraw %}
```

To transform the value to *YAML*, or *JSON*, `-s` has to be a **valid** *JSON*.

## Creating Kubernetes Secrets

**Aegis Safe**-managed secrets can be interpolated onto Kubernetes secrets if
a template transformation is given.

Let’s take the following as an example:

```bash
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  -w "billing" \
  -n "finance" \
  -s '{"username": "root", "password": "SuperSecret"}' \
  -t '{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}"' \
  -k{% endraw %}
```

The `-k` flag hints **Aegis Safe** that the secret will be synced with a 
Kubernetes `Secret`. `-n` tells that the namespace of the secret is `"finance"`.

Before running this command, a secret with the following manifest should exist
in the cluster:

```yaml 
apiVersion: v1
kind: Secret
metadata:
  name: aegis-secret-billing
  namespace: finance
type: Opaque
```

The `aegis-secret-` prefix is required for **Aegis Safe** to locate the 
`Secret`. Also `metadata.namespace` attribute should match the namespace 
that’s provided with the `-n` flag to **Aegis Sentinel**.

After executing the command the secret will contain the new values in its
`data:` section.

```bash 
kubectl describe secret aegis-secret-billing -n finance

# Output:
#   Name:         aegis-secret-billing
#   Namespace:    finance
#   Labels:       <none>
#   Annotations:  <none>
#
#   Type:  Opaque
#
#   Data
#   ====
#   USERNAME:  137 bytes
#   PASSWORD:  196 bytes
```

## Setting the Master Secret Manually

**Aegis Safe** uses a master secret to encrypt the secrets that are stored.
Typically, this master secret is stored as a Kubernetes secret for your 
convenience. However, if you want you can set `AEGIS_MANUAL_KEY_INPUT` to
`"true"` and provide the master secret manually.

Although this approach enhances security, it also means that you will have to
provide the master secret to **Aegis Safe** whenever the pod is evicted, or
crashes, or restarts for any reason. Since, this brings a mild operational
inconvenience, it is not enabled by default.

```bash 
{% raw %}kubectl exec "$SENTINEL" -n aegis-system -- aegis \
  --input-keys "AGE-SECRET-KEY-1RZU…\nage1…\na6…ceec"
  
# Output:
#
# OK{% endraw %}
```
