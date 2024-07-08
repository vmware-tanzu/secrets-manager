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

title = "Encrypting Secrets Using VSecM"
weight = 230
+++

## Situation Analysis

Sometimes you might want to store your secrets in a safe place, and you might
want to share them with others. However, you don't want to share them in plain
text. This is where **VSecM Sentinel** comes in.

Using **VSecM Sentinel**, you can encrypt your secrets and store them in a safe
place. When you're ready to use the secret, you can decrypt it using **VSecM
Safe** and distribute it to the workload that needs it.

Since the secret will be encrypted, you can freely share it, and store in
source control systems. When you're ready to submit a secret to the workload, 
rather than providing the secret in plain text, you can deliver its encrypted 
version to **VSecM Safe**.

## Strategy

Use **VSecM Sentinel** to encrypt your secrets. Safely store the secrets in
a source control system, or share them with others. When you're ready to use
the secrets, provide it to **VSecM Safe** through **VSecM Sentinel** using the
`-e` flag to indicate that the secret is encrypted.

## High-Level Diagram

Open the image in a new tab to see the full-size version:

![High-Level Diagram](/assets/encrypt.jpg "High-Level Diagram")

## Implementation

This use case will introduce how you can use **VSecM Sentinel** encrypt secrets
for safe keeping outside your cluster.

### About the Encryption Process

Please note that the encryption process and its inner workings remain mostly
hidden to the end-user, ensuring a user-friendly experience.

The process employs asymmetric encryption, where the secret is encrypted with a
public key and decrypted using a private key by **VSecM Safe**. However,
this is an implementation detail which can be subject to change.

### Cleanup

Let's remove the workload as usual:

```bash 
kubectl delete deployment example
```

Next, delete the secret associated with this workload:

```bash
# Find the sentinel pod's name:
kubectl get po -n vsecm-system

# Delete secrets:
kubectl exec vsecm-sentinel-778b7fdc78-86v6d -n \
  vsecm-system -- safe -w example -d

OK
```

That should be enough cleanup for the next steps.

### Introducing **VSecM Inspector**

We will use **VSecM Inspector** like a debugger, to diagnose the
state of our system.

You can find sample deployment manifests for **VSecM Inspector** 
[in the **Transforming Secrets Using VSecM**][transforming-secrets] use case.

[transforming-secrets]: @/documentation/use-cases/transform.md "Transforming Secrets Using VSecM"

After following the linked guide, and installing **VSecM Inspector**,  let's
test it:

```bash
INSPECTOR=$(kubectl get po -n default \
  | grep "vsecm-inspector-" | awk '{print $1}')
  
kubectl exec $INSPECTOR -- ./env

# Output:
# Failed to fetch the secrets. Try again later.
# Secret does not exist
```

### Encrypting a Secret

Now, let's encrypt a secret using **VSecM Sentinel**:

```bash
export SENTINEL=$(kubectl get po -n vsecm-system \
  | grep "vsecm-sentinel-" | awk '{print $1}')
  
kubectl exec $SENTINEL -n vsecm-system -- safe \
  -s "VSecMRocks" \
  -e

# The output of the above command will be 
# similar to something like this:
#
#   YWdlLWVuY ... Truncated ... VZ2SDFiMjEY+V7JMg
#
# ☝️ This is a long random encrypted string. 
# We will use the variable $ENCRYPTED_SECRET in lieu of
# this value in the sections below for simplicity.
```

Here `-s` is for the secret we would like to encrypt, and `-e` indicates
that we are not going to store the secret (*yet*), instead we want **VSecM Sentinel**
to output the encrypted value of the secret to us.

### Registering the Encrypted Secret

To register an encrypted secret, we use the `-e` flag to indicate that the
secret is not plain text, and it is encrypted.

```bash
kubect exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s "$ENCRYPTED_SECRET" \
  -e 
```

And finally let's inspect and see if the secret is registered properly:

```bash
kubectl exec $INSPECTOR -- ./env

# Will return "VSecMRocks"
```

And yes, it did.

### Be Aware of the `vsecm-root-key` Kubernetes `Secret`

One thing to note is, if you lose access to the Kubernetes `Secret` named
`vsecm-root-key` in the `vsecm-system` namespace, then you will lose the
ability to register your encrypted secrets (*since, during bootstrapping
when VSecM Safe cannot find the secret, it will create a brand new one,
invalidating all encrypted values*).

As a rule of thumb, **always backup your cluster** regularly, so that if
such an incident occurs, you can recover the `vsecm-root-key` secret
from the backups.

## Conclusion

This use case demonstrated how you can encrypt a secret value and register the
encrypted value to **VSecM Safe** instead of the plain text secret. This
technique provides and added layer of protection, and also allows you to
safe the secret anywhere you like including source control systems.

**VMware Secrets Manager** offers a robust solution for the secure encryption, 
storage, and distribution of secrets within Cloud Native environments. By 
leveraging **VSecM Sentinel** and **VSecM Safe**, organizations can maintain the 
confidentiality of their sensitive information while ensuring easy access when 
necessary. 

The encrypted secrets can be safely stored in source control systems or shared 
with authorized personnel without risk, given the assurance of strong encryption 
measures. 

Regular backups and diligent management of encryption keys form the cornerstone 
of a robust security strategy, ensuring that even in the face of system failures 
or security breaches, data integrity and access can be quickly restored. 
This comprehensive approach to secret management showcases the efficacy of 
**VSecM** in addressing the complex security challenges faced by enterprises 
today, promoting a secure, efficient, and compliant operational environment.

## List of Use Cases in This Section

{{ use_cases() }}

{{ edit() }}
