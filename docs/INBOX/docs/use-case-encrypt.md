---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, tutorial, secrets, encryption
title: Encrypting Secrets
description: encrypt secrets for external storage
micro_nav: true
page_nav:
  prev:
    content: using <strong>Aegis</strong> init container
    url: '/docs/use-case-init-container'
  next:
    content: transforming secrets
    url: '/docs/use-case-transform'
---

<p style="text-align:right;position:relative;top:-40px;"
><a href="https://github.com/ShieldWorks/aegis-web/blob/main/docs/use-case-encrypt.md"
style="border-bottom: none;background:#e0e0e0;padding:0.5em;display:inline-block;
border-radius:8px;">
edit this page on <strong>GitHub</strong> ✏️</a></p>

## Introduction

This tutorial will introduce how you can use **Aegis Sentinel** encrypt secrets
for safe keeping outside your cluster.

## What Is the Benefit?

Since the secret will be encrypted, you can freely share it, and store in 
source control systems. 

When you’re ready to submit a secret to the workload, rather than providing the 
secret in plain text, you can deliver its encrypted version to **Aegis Safe**. 

This method offers a couple of distinct benefits:

Firstly, it increases your overall security.

Secondly, it allows for role differentiation: The individual (*or process*) who 
submits the secret doesn’t have to know its actual content; instead, they work 
with the encrypted version. 

Consequently, even if an impostor tries to mimic this individual, they wouldn’t 
be able to decipher the secret’s true value, drastically reducing potential 
avenues for attack.

## About the Encryption Process

Please note that the encryption process and its inner workings remain mostly 
hidden to the end-user, ensuring a user-friendly experience. 

The process employs asymmetric encryption, where the secret is encrypted with a 
public key and decrypted using a private key by **Aegis Safe**. However,
this is an implementation detail which can be subject to change.

## Cleanup

Let’s remove the workload as usual:

```bash 
{% raw %}kubectl delete deployment example{% endraw %}
```

Next, delete the secret associated with this workload:

```bash
{% raw %}# Find the sentinel pod’s name:
kubectl get po -n aegis-system

# Delete secrets:
kubectl exec aegis-sentinel-778b7fdc78-86v6d -n \
  aegis-system -- aegis -w example -d

OK{% endraw %}
```

That should be enough cleanup for the next steps.

## Introducing **Aegis Inspector**

We will use **Aegis Inspector** like a debugger, to diagnose the 
state of our system.

By the time of this writing **Aegis Inspector** is not an official **Aegis**
component, so we’ll piggyback on a `Deployment` manifest that was used in 
a former workshop. When we have an `aegis-inspector` pod that we can officially
use for diagnostic purposes, this paragraph will be edited to reflect that too.

Yet, for now, let’s deploy the workshop version of it.

```bash 
# Switch to the Aegis repo:
cd $WORKSPACE/aegis
# Install Aegis Inspector:
cd examples/aegis-workshop/inspector
kubectl apply -f ServiceAccount.yaml 
kubectl apply -k .
# Register Aegis Inspector’s ClusterSPIFFEID
cd ../ids
kubectl apply -f Inspector.yaml
```

Now let’s test it:

```bash
INSPECTOR=$(kubectl get po -n default \
  | grep "aegis-inspector-" | awk '{print $1}')
  
kubectl exec $INSPECTOR -- ./env

# Output:
# Failed to fetch the secrets. Try again later.
# Secret does not exist
```

## Encrypting a Secret

Now, let’s encrypt a secret using **Aegis Sentinel**:

```bash
export SENTINEL=$(kubectl get po -n aegis-system \
  | grep "aegis-sentinel-" | awk '{print $1}')
  
kubectl exec $SENTINEL -n aegis-system -- aegis \
  -s "AegisRocks" \
  -e

# The output of the above command will be similar to something like this:
#
#   YWdlLWVuY … Truncated … VZ2SDFiMjEY+V7JMg
#
# ☝️ This is a long random encrypted string. 
# We will use the variable $ENCRYPTED_SECRET in lieu of this value 
# in the sections below for simplicity.
```

Here `-s` is for the secret we would like to encrypt, and `-e` indicates 
that we are not going to store the secret (*yet*), instead we want **Aegis Sentinel**
to output the encrypted value of the secret to us.

## Registering the Encrypted Secret

To register an encrypted secret, we use the `-e` flag to indicate that the 
secret is not plain text, and it is encrypted.

```bash
kubect exec $SENTINEL -n aegis-system -- aegis \
  -w example \
  -s "$ENCRYPTED_SECRET" \
  -e 
```

And finally let’s inspect and see if the secret is registered properly:

```bash
kubectl exec $INSPECTOR -- ./env

# Will return "AegisRocks"
```

And yes, it did.

## Be Aware of the `aegis-safe-age-key` Kubernetes `Secret`

One thing to note is, if you lose access to the Kubernetes `Secret` named
`aegis-safe-age-key` in `aegis-system` namespace, then you will lose the
ability to register your encrypted secrets (*since, during bootstrapping
when Aegis Safe cannot find the secret, it will create a brand new one,
invalidating all encrypted values*).

As a rule of thumb, **always backup your cluster** regularly, so that if
such an incident occurs, you can recover the `aegis-safe-age-key` secret
from the backups.

## Conclusion

This tutorial demonstrated how you can encrypt a secret value and register the
encrypted value to **Aegis Safe** instead of the plain text secret. This
technique provides and added layer of protection, and also allows you to 
safe the secret anywhere you like including source control systems.

Next up, you’ll learn about secret transformations.



