---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, tutorial, secrets, transform
title: Transforming Secrets
description: interpolate <strong>Aegis</strong> secrets onto Kubernetes <em>Secret</em>s
micro_nav: true
page_nav:
  prev:
    content: encrypting secrets
    url: '/docs/use-case-encrypt'
  next:
    content: <strong>Aegis</strong> showcase
    url: '/docs/use-case-showcase'
---

<p style="text-align:right;position:relative;top:-40px;"
><a href="https://github.com/ShieldWorks/aegis-web/blob/main/docs/use-case-transform.md"
style="border-bottom: none;background:#e0e0e0;padding:0.5em;display:inline-block;
border-radius:8px;">
edit this page on <strong>GitHub</strong> ✏️</a></p>

## Introduction

This tutorial will show various way you can interpolate and transform secrets.

Transforming secrets may come in handy when your workload expects the secret
in a different format than it has been initially provided, and you don’t want
to write custom code to do the transformation.

To help us explore these transformations, [we will use **Aegis Inspector**
from the previous tutorial](/docs/use-case-encrypt). If you haven’t installed
it, before you proceed, please [navigate to that lecture and install 
**Aegis Inspector**](/docs/use-case-encrypt)

## Preparation

Let us define a few aliases first, they will speed things up:

```bash 
SENTINEL=$(kubectl get po -n aegis-system \
  | grep "aegis-sentinel-" | awk '{print $1}')
SAFE=$(kubectl get po -n aegis-system \
  | grep "aegis-safe-" | awk '{print $1}')
WORKLOAD=$(kubectl get po -n default \
  | grep "example-" | awk '{print $1}')
INSPECTOR=$(kubectl get po -n default \
  | grep "aegis-inspector-" | awk '{print $1}')

# Delete secrets assigned to the workload:
alias delete-secret="kubectl exec $SENTINEL \
  -n aegis-system -- aegis \
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

**Aegis Sentinel** CLI accepts a format flag (`-f`), the possible values are

* `"json"`
* and `"yaml"`

If it is not given, it defaults to `"json"`; however, in the upcoming examples
we’ll be explicit and provide this argument at all times.

## Registering a JSON Secret

```bash
{% raw %}kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w example \
  -s '{"username": "admin", "password": "AegisRocks!"}' \
  -f json

inspect
# Output:
# {"username": "admin", "password": "AegisRocks!"}{% endraw %}
```

## Registering a YAML Secret

```bash 
{% raw %}kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w example \
  -s '{"username": "admin", "password": "AegisRocks!"}' \
  -f yaml

inspect
# Output:
# password: AegisRocks!
# username: admin{% endraw %}
```

# Registering a JSON String (with invalid JSON)

Now we’ll deliberately make an error in our JSON. Notice the missing `"`
in `username"`: That is not valid JSON.

```bash
{% raw %}kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w example \
  -s '{username": "admin", "password": "AegisRocks!"}' \
  -f json

inspect
# Output:
# {username": "admin", "password": "AegisRocks!"}{% endraw %}
```

# Registering a YAML String (with invalid JSON)

Since the JSON cannot be parsed, the output will not be a YAML:

```bash 
{% raw %}kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w example \
  -s '{username": "admin", "password": "AegisRocks!"}' \
  -f yaml

inspect
# Output:
# {username": "admin", "password": "AegisRocks!"}{% endraw %}
```

## Transforming A JSON Secret

```bash
{% raw %}kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w example \
  -s '{"username": "admin", "password": "AegisRocks!"}' \
  -t '{"USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f json

inspect
# Output:
# {"USR":"admin", "PWD":"AegisRocks!"}{% endraw %}
```

## Transforming a YAML Secret

```bash
{% raw %}kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w example \
  -s '{"username": "admin", "password": "AegisRocks!"}' \
  -t '{"USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f yaml

inspect
# Output:
# PWD: AegisRocks!
# USR: admin{% endraw %}
```

## Transforming a JSON Secret (invalid JSON)

If our secret is not valid JSON, then the YAML transformation will not be
possible. **Aegis** will still try its best to provide something.

```bash
{% raw %}kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w example \
  -s '{username": "admin", "password": "AegisRocks!"}' \
  -t '{"USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f json

inspect
# Output:
# {username": "admin", "password": "AegisRocks!"}{% endraw %}
```

## Transforming a JSON Secret (invalid template)

Since template is not valid, the template transformation will not happen.

```bash
{% raw %}kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w example \
  -s '{"username": "admin", "password": "AegisRocks!"}' \
  -t '{USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f json

inspect
# Output:
# {"username": "admin", "password": "AegisRocks!"}{% endraw %}
```

# Transforming a JSON Secret (invalid template and JSON)

**Aegis** will still try its best:

```bash
{% raw %}kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w example \
  -s '{username": "admin", "password": "AegisRocks!"}' \
  -t '{USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f json

inspect
# Output:
# {username": "admin", "password": "AegisRocks!"}{% endraw %}
```

## Transforming YAML Secret (invalid JSON)

```bash
{% raw %}kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w example \
  -s '{username": "admin", "password": "AegisRocks!"}' \
  -t '{"USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f yaml

inspect
# Output
# {username": "admin", "password": "AegisRocks!"}{% endraw %}
```

## Transforming YAML Secret (invalid template)

```bash
{% raw %}kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w example \
  -s '{"username": "admin", "password": "AegisRocks!"}' \
  -t '{USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f yaml

inspect
# Output:
# {USR":"admin", "PWD":"AegisRocks!"}{% endraw %}
```

## Transforming YAML Secret (invalid JSON and template)

```bash
{% raw %}kubectl exec $SENTINEL -n aegis-system -- aegis \
  -w example \
  -s '{username": "admin", "password": "AegisRocks!"}' \
  -t '{USR":"{{.username}}", "PWD":"{{.password}}"}' \
  -f yaml

inspect
# Output:
# {username": "admin", "password": "AegisRocks!"}{% endraw %}
```

## Conclusion

This tutorial demonstrated various ways to transform and interpolate secret
values into `JSON` and `YAML`. We also observed how the output is affected
when there is a formatting issue with the secret, or the template to 
transform the secret, or both of them.

The next section introduces a video tutorial that covers everything that has
been mentioned so far and some more.


