---
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title: Secret Transformation
layout: post
prev_url: /docs/use-case-encryption/
permalink: /docs/use-case-transformation/
next_url: /docs/showcase/
---

<p class="github-button"
><a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0240-transform.md"
>edit this page on <strong>GitHub</strong> ✏️</a></p>

## Introduction

This tutorial will show various way you can interpolate and transform secrets.

Transforming secrets may come in handy when your workload expects the secret
in a different format than it has been initially provided, and you don’t want
to write custom code to do the transformation.

To help us explore these transformations, [we will use **VSecM Inspector**
from the previous tutorial](/docs/use-case-encryption). If you haven’t installed
it, before you proceed, please [navigate to that lecture and install
**VSecM Inspector**](/docs/use-case-encryption)

## Preparation

Let us define a few aliases first, they will speed things up:

```bash 
SENTINEL=$(kubectl get po -n vsecm-system \
  | grep "vsecm-sentinel-" | awk '{print $1}')
SAFE=$(kubectl get po -n vsecm-system \
  | grep "vsecm-safe-" | awk '{print $1}')
WORKLOAD=$(kubectl get po -n default \
  | grep "example-" | awk '{print $1}')
INSPECTOR=$(kubectl get po -n default \
  | grep "vsecm-inspector-" | awk '{print $1}')

# Delete secrets assigned to the workload:
alias delete-secret="kubectl exec $SENTINEL \
  -n vsecm-system -- safe \
  -w example -s x -d"

alias inspect="kubectl exec $INSPECTOR -- ./env"
```

Now, we can start experimenting.

## Cleanup

Let’s start with a blank slate again:

```bash 
delete-secret
# Output: OK

inspect
# Output:
# Failed to fetch the secrets. Try again later.
# Secret does not exist
```

## The Format (`-f`) Argument

**VSecM Sentinel** CLI accepts a format flag (`-f`), the possible values are

* `"json"`
* and `"yaml"`

If it is not given, it defaults to `"json"`; however, in the upcoming examples
we’ll be explicit and provide this argument at all times.

## Registering a JSON Secret

```bash
{% raw %}kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s '{"username": "admin", "password": "VSecMRocks!"}' \
  -f json

inspect
# Output:
# {"username": "admin", "password": "VSecMRocks!"}{% endraw %}
```

## Registering a YAML Secret

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

## Registering a JSON String (with invalid JSON)

Now we’ll deliberately make an error in our JSON. Notice the missing `"`
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

## Registering a YAML String (with invalid JSON)

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

## Transforming A JSON Secret

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

## Transforming a YAML Secret

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

## Transforming a JSON Secret (invalid JSON)

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

## Transforming a JSON Secret (invalid template)

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

## Transforming a JSON Secret (invalid template and JSON)

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

## Transforming YAML Secret (invalid JSON)

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

## Transforming YAML Secret (invalid template)

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

## Transforming YAML Secret (invalid JSON and template)

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

This tutorial demonstrated various ways to transform and interpolate secret
values into `JSON` and `YAML`. We also observed how the output is affected
when there is a formatting issue with the secret, or the template to
transform the secret, or both of them.

The next section introduces a video tutorial that covers everything that has
been mentioned so far and some more.


