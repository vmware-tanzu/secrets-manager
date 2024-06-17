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

title = "VSecM CLI"
weight = 3
+++

## Introduction

This section contains usage examples and documentation for
[**VSecM Sentinel**][sentinel]'s Command Line Interface (*CLI*).

[sentinel]: https://github.com/vmware-tanzu/secrets-manager/tree/main/app/sentinel

## Finding **VSecM Sentinel**

First, find which pod belongs to `vsecm-system`:

```bash
kubectl get po -n vsecm-system
```

The response to the above command will be similar to the following:

```txt
NAME                              READY
vsecm-safe-5f6948c84c-vkrdh       1/1
vsecm-sentinel-5998b5dbfc-lvw44   1/1
```

There, `vsecm-sentinel-5998b5dbfc-lvw44` is the name of the Pod you'd need.

You can also execute a script similar to the following to save the Pod's name
into an environment variable:

```bash
SENTINEL=$(kubectl get po -n vsecm-system \
  | grep "vsecm-sentinel-" | awk '{print $1}')
```

In the following examples, we'll use `$SENTINEL` in lieu of the **VSecM Sentinel**'s
Pod name.

## Displaying Help Information

**VSecM Sentinel** has a binary called `safe` that can be used to interact with
**VSecM Safe** API. You can use the following command to display help information
for `safe`:

```bash
kubectl exec $SENTINEL -n vsecm-system -- safe --help
```

> **About `--help`**
>
> The output of `safe --help` will depend on the version of `safe` you
> use; however, it will contain useful information about how to use the program.


## Registering a Secret for a Workload

Given our workload has the SPIFFE ID `"spiffe://vsecm.com/workload/billing/: ...[truncated]"`

```bash
kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w billing \
  -s "very secret value"
```

will register the secret `"very secret value"` to `billing`.

## Registering a Secrets as a Kubernetes `Secret`

For legacy systems, or for systems that cannot be modified to use **VSecM Sidecar**,
or **VSecM SDK**, of for workloads that you don't have direct access to the
source code, or when modifying the source code will take too much time,
introducing lots of upstream changes, you can use **VSecM Sentinel** to
register secrets as Kubernetes `Secret`s and let the workloads consume them
by binding those `Secret`s to the workloads with usual Kubernetes mechanisms
such as environment variables, or volume mounts.

For this, you'll need to prefix the workload name with
`VSECM_SAFE_STORE_WORKLOAD_SECRET_AS_K8S_SECRET_PREFIX` environment variable's value
(*which is `k8s:` by default*):

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "k8s:example-secret" \
  -n "default" \
  -s '{"username": "root", \
    "password": "SuperSecret", \
    "value": "VSecMRocks"}' \
  -t '{"USER":"{{.username}}", \
    "PASS":"{{.password}}", \
    "VALUE": "{{.value}}"}'
# Will create a Kubernetes Secret with name `example-secret` in the "default 
# namespace, and will have the following data:
# USER: root
# PASS: SuperSecret
# VALUE: VSecMRocks

```   

## Registering Multiple Secrets

You can use the `-a` (*append*) argument to register more than one secret
to a workload.

```bash
kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w billing \
  -s "first part of the token" \
  -a

kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w billing \
  -s "second part of the token" \
  -a
```

## Encrypting a Secret

Use the `-e` flag to encrypt a secret.

```bash
kubectl exec $SENTINEL -n vsecm-system -- safe \
  -s "very secret value" \
  -e 

# The above command outputs an encrypted string that can be
# securely stored anywhere, including source control systems.
```

## Registering an Encrypted Secret

Again, `-e` flag can be used to register an encrypted secret to a workload:

```bash
kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w billing \
  -s $ENCRYPTED_SECRET \
  -e
```

## Deleting a Secret

To delete the secret associated to a workload, use the `-d` flag:

```bash
kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w billing \
  -d
```

## Listing Secrets Registered to a Workload

**VSecM Sentinel** does not show the secrets in plain text by design.
However, you can use the `-l` flag to get certain metadata about the secrets
registered to a workload.

```bash
kubectl exec $SENTINEL -n vsecm-system -- safe -l
# Output:
# {"secrets":[{"name":"example",
#  "created":"Tue Jan 02 02:06:30 +0000 2024",
#  "updated":"Tue Jan 02 02:06:30 +0000 2024"}]}
```

In addition, you can use `-l -e` to get encrypted versions of the secrets.
If you have adequate privileges to get the VSecM root key, you can use
the root key to decrypt the secrets.

```bash
kubectl exec $SENTINEL -n vsecm-system -- safe -l -e
# Output:
# {"secrets":[{"name":"example",
#  "value":["YWdlLWVuY3J5cHRpb24ub3JnL3YxCi0...truncated"],
#  "created":"Tue Jan 02 02:06:30 +0000 2024",
#  "updated":"Tue Jan 02 02:06:30 +0000 2024"}],
#  "algorithm":"age"}
```

## Choosing a Backing Store

The registered secrets will be encrypted and backed up to
**VSecM Safe**'s Kubernetes volume by default. This behavior can be configured
by changing the `VSECM_SAFE_BACKING_STORE` environment variable that
**VSecM Safe** sees. In addition, this behavior can be overridden on a per-secret
basis too.

The following commands stores the secret to the backing volume
(*default behavior*):

```bash
kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w billing \
  -s "very secret value" \
  -b file
```

This one, will **not** store the secret on file; the secret will only be
persisted in memory, and will be lost if **VSecM Sentinel** needs to restart:

```bash
kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w billing \
  -s "very secret value" \
  -b memory
```

## Template Transformations

You can transform how the stored secret is displayed to the consuming workloads:

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "example" \
  -s '{"username": "root", \
    "password": "SuperSecret", \
    "value": "VSecMRocks"}' \
  -t '{"USER":"{{.username}}", \
    "PASS":"{{.password}}", \
    "VALUE": "{{.value}}"}'
```

When the workload fetches the secret through the workload API, this is what
it sees as the value:

```txt
{"USER": "root", "PASS": "SuperSecret", "VALUE": "VSecMRocks"}
```

Instead of this default transformation, you can output it as `yaml` too:

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "example" \
  -s '{"username": "root", \
    "password": "SuperSecret", \
    "value": "VSecMRocks"}' \
  -t '{"USER":"{{.username}}", \
    "PASS":"{{.password}}", \
    "VALUE": "{{.value}}"}' \
  -f yaml
```

The above command will result in the following secret value to the workload
that receives it:

```txt
USER: root
PASS: SuperSecret
VALUE: VSecMRocks
```

> **`"json"` Is the Default Value**
>
> If you don't specify the `-f` flag, it will default to `"json"`.


You can create a YAML secret value without the `-t` flag too. In that case
**VSecM Safe** will assume an identity transformation:

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "example" \
  -s '{"username": "root", \
    "password": "SuperSecret", \
    "value": "VSecMRocks"}' \
  -f yaml
```

The above command will result in the following secret value to the workload
that receives it:

```txt
username: root
password: SuperSecret
value: VSecMRocks
```

If you provide `-f json` as the format argument, the secret will have to be
a strict JSON. Otherwise, **VSecM Sentinel** will try to come up with a
reasonable value, and not raise an error; however the output will likely be
in a format that the workload is not expecting.

> **Gotcha**
>
> If `-t` is given, the `-s` argument will have to be a valid JSON regardless
> of what is chosen for `-f`.


The following is also possible:

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "example" \
  -s 'USER»»»{{.username}}'
```

and will result in the following as the secret value for the workload:

```txt
USER»»»root
```

Or, equivalently:

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "example" \
  -s 'USER»»»{{.username}}' \
  -f json
```

Will provide the following to the workload:

```txt
USER»»»root
```

Similarly, the following will **not** raise an error:

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "example" \
  -s 'USER»»»{{.username}}' \
  -f yaml
```

To transform the value to *YAML*, or *JSON*, `-s` has to be a **valid** *JSON*.

## Creating Kubernetes Secrets

**VMware Secrets Manager** can create Kubernetes `Secret`s for you, too.
This is especially useful for use case where you cannot directly modify the
source code of the workloads, or where you don't have access to the source.

This way, you can use **VSecM Sentinel** to create Kubernetes `Secret`s and
bind them to the workloads with usual Kubernetes mechanisms such as environment
variables, or volume mounts.

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "k8s:example-kubernetes-secret" \
  -n "default" \
  -s '{"username": "root","password": "SuperSecret"}' \
  -t '{"USER":"{{.username}}", "PASS":"{{.password}}"}'

````

The `k8s:` prefix tells **VSecM Sentinel** to create a Kubernetes `Secret`
with name `example-kubernetes-secret`. The `-n` flag tells **VSecM Sentinel**
to create the `Secret` in the `default` namespace. The `-s` flag tells
**VSecM Sentinel** to store the secret as a JSON. The `-t` flag tells
**VSecM Sentinel** to transform the secret to a different format.

Once created, you can retrieve the `Secret` by usual `kubectl` commands:

```bash 
kubectl get secret example-kubernetes-secret -o yaml

# Sample Output:
# apiVersion: v1
# data:
#  PASS: U3VwZXJTZWNyZXQ=
#  USER: cm9vdA==
# kind: Secret
# metadata:
#   name: example-kubernetes-secret
#   namespace: default
# type: Opaque

```

> **Using VSecM Init Container**
>
> It could also be useful to use **VSecM Init Container** to let the workload
> wait until the `Secret` is created.
>
> This can be done by adding the **VSecM Init Container** to the workload's
> `Deployment` or `StatefulSet` manifest, then you can execute a command
> similar to the following to trigger the **VSecM Init Container** finish
> its job and let the workload start:
>
> ```bash 
> kubectl exec "$SENTINEL" -n vsecm-system -- safe -w "k8s:$WORKLOAD_NAME"
> 
> Check out the [examples folder][github-examples] for more information.


[github-examples]: https://github.com/vmware-tanzu/secrets-manager/tree/main/examples "VSecM Examples"


## Setting the Root Key Manually

**VSecM Safe** uses a **root key** to encrypt the secrets that are stored.
Typically, this **root key** is stored as a Kubernetes `Secret` for your
convenience. However, if you want you can set `VSECM_MANUAL_KEY_INPUT` to
`"true"` and provide the *root key* manually.

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -i "AGE-SECRET-KEY-1RZU...\nage1...\na6...ceec"
  
# Output:
#
# OK
```

> **The New Root Key Will Be Not Persist After Pod Restart**
>
> While this behavior will change in the future, and there will be an option
> to persist the root key via a command line flag, the current behavior is
> to save root key only to the memory of **VSecM Safe**
>
> Although this approach enhances security, it also means that you will have to
> provide the *root key* to **VSecM Safe** whenever the pod is evicted, or
> crashes, or restarts for any reason. Since, this brings a mild operational
> inconvenience, it is not enabled by default.


Also keep in mind that when you rotate the root key, the secrets that have
been encrypted with the old root key will no longer be accessible. Again,
this behavior may change in the future may become more convenient for the
end users.

## Retrieving the Current Root Key

If you have cluster-admin privileges, you can retrieve the current root key
with the following command:

```bash
./hack/get-root-key.sh

## Output:
# AGE-SECRET-KEY-1Q2UVA35Q7666TK7Y...truncated
# age1sy723jqnnff2aasfz...truncated
# 89febb5c78c484ef459b861f6a8c...truncated
```

## Generating a New Root Key

You can use **VSecM Keygen** via the following command to generate a new root key:

```bash
docker run --rm vsecm-keygen:latest

# Output
# AGE-SECRET-KEY-14DVY8Y0J4JQA45Z...truncated
# age1ghxkaqg5kkt8rl98x...truncated
# bc95a5e9e81fdaf40fe0exxx...truncated
```

## Decrypting Exported Encrypted Secrets

If you have access to the root key, you can decrypt the exported secrets
that have been exported using `safe -l -e`:

```bash
docker run --rm \
  -v "$(pwd)":/vsecm \
  -e VSECM_KEYGEN_EXPORTED_SECRET_PATH="/vsecm/secrets.json" \
  -e VSECM_KEYGEN_ROOT_KEY_PATH="/vsecm/key.txt" \
  -e VSECM_KEYGEN_DECRYPT="true" \
  vsecm-keygen:latest
```

Where `key.txt` is the new root key you want to use, and `secrets.json` is
the output of `kubectl exec $SENTINEL -n vsecm-system -- safe -l -e`.

## Registering Random Pattern-Based Secrets to a Workload

**VSecM Sentinel** can generate random secrets based on a pattern. Here are
some of the patterns that you can use:

* `footballt[\w]{8}bar`: will generate a random string that starts with `football`,
  ends with `bar`, and has 8 characters in between.
* `admin[a-z0-9]{3}`: will generate a random string that starts with `admin`,
  and has 3 characters in between, which can be either lowercase letters or
  numbers.
* `admin[a-z0-9]{3}something[\w]{3}`: will generate a random string that starts
  with `admin`, has 3 characters in between, which can be either lowercase
  letters or numbers, and ends with `something` followed by 3 more characters.
* `pass[a-zA-Z0-9]{12}`: will generate a random string that starts with `pass`,
  and has 12 characters in between, which can be either lowercase letters,
  uppercase letters, or numbers.
* `football[\d]{8}bar`: will generate a random string that starts with `football`,
  ends with `bar`, and has 8 digits in between.

To use these patterns, simply prefix the `-v` flag with `gen:` (*or
`VSECM_SENTINEL_SECRET_GENERATION_PREFIX` if you are using the environment
to override it*) as follows:

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "example" \
  -s "gen:football[\w]{8}bartender"
# The secret will be randomized based on the pattern above.  

```

## Using "Init Commands" to Pre-Register Secrets

You can use an `initCommand` in your helm charts to pre-register secrets, too.

Here is an example:

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

This is a relatively advanced use case that is covered in the
[Registering Secrets][registering-secrets] page.

[registering-secrets]: @/documentation/use-cases/registering-secrets.md "Registering Secrets"

{{ edit() }}
