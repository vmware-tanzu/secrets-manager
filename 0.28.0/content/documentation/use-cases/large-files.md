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

title = "Encrypting Large Files"
weight = 245
+++

## Situation Analysis

This use case is about encrypting large files and securely storing them in a 
database, on the file system, or in source control.

In certain use cases, you might need to securely store a large and sensitive file 
in a database, on the file system, or in source control. Since storing a massive 
file as a secret is impractical, we'd need an alternate approach.

One approach is to store a symmetric encryption key in **VMware Secrets Manager**, 
use that key to encrypt the file that we need to secure and store the file in 
encrypted form wherever we want to store it.

Then, when a workload needs a decrypted file, a sidecar can fetch the symmetric 
key, decrypt the file, and provide the file to the workload.

## Screencast

Here is a screencast that demonstrates this use case:

```text
WORK IN PROGRESS
```

## High-Level Diagram

Here's one of several ways the interaction mentioned in the former section can happen. 

This sequence diagram follows a GitOps-like workflow, yet it can be adapted to 
various other scenarios, too.

Open the image in a new tab to see the full-size version:

![High-Level Diagram](/assets/large-files.png "High-Level Diagram")

## Implementation

This is a relatively high-level design, so how it's implemented will widely vary 
on your specific use case. In this implementation section, we'll follow the above 
high-level sequence diagram.

### A Human Operator Creates an AES Key Off-Cycle

We'll create an AES key and store it in VSecM through [**VSecM CLI**][vsecm-cli]:

```bash
# Create an 256-bit AES Key:
AES_KEY=$(openssl rand -hex 32)
# Save the key in VSECM for our example workload:
kubectl exec $SENTINEL -n vsecm-system -- safe \
  -w example \
  -s $AES_KEY
```

[vsecm-cli]: @/documentation/usage/cli.md "VSecM CLI"

> **Do You Want the AES Key Created Automatically?**
> 
> It's possible not to involve a human operator and let VSecM create the AES key 
> needed internally and dispatch it to the workloads that need it--If you help 
> implement an outstanding issue that will make this happen:
> 
> * [VSecM shall be able to create an AES key `#764`][ticket-764]
>
> Your contributions are welcome.
{: .block-warning }

[ticket-764]: https://github.com/vmware-tanzu/secrets-manager/issues/764

### The Operator Encrypts the File Using the Key

Now we encrypt the sensitive file using the AES key off-cycle and store it in a 
Git repository:

```bash 
# Encrypt the file:
openssl enc -aes-256-cbc -salt -in myfile.txt -out myfile.enc -k $AES_KEY
# Store it to a repo.
# This part will likely be automated:
git add myfile.enc
git commit -m "added encrypted file"
git push origin main
```

### The Workload Uses The AES Key to Decrypt the File

This step will depend on how you mount the encrypted file to your workload and how 
you fetch the AES key.

You can fetch the AES key by…

* [Using VSecM Sidecar][use-case-sidecar]
* [Using VSecM SDK][use-case-sdk]
* Mounting it from a Kubernetes `Secret` as described in 
  [Mounting Secrets as Volumes[use-case-mount] use case
* or Using a custom-built Sidecar or [Init Container][init-container] that 
  uses [**VSecM SDK**][vsecm-sdk].

[use-case-sidecar]: @/documentation/use-cases/sidecar.md
[use-case-sdk]: @/documentation/usage/sdk.md
[init-container]: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/

We intentionally leave this section open-ended since it will depend on your 
architecture, GitOps tooling, and build pipelines.

The decryption of the file can also be done in a custom init container so that 
it will be ready once the workload starts its lifecycle, or [you can help with 
this issue][ticket-765] and add this capability to the **VSecM Init Container**.

[ticket-765]: https://github.com/vmware-tanzu/secrets-manager/issues/765

## Conclusion

In conclusion, the proposed solution provides a robust and secure method for 
managing the encryption and decryption of large and sensitive files within 
various IT environments. 

By leveraging **VMware Secrets Manager** for storing and managing symmetric 
encryption keys and combining it with a flexible approach to encryption and 
decryption workflows, organizations can ensure the confidentiality and integrity 
of their sensitive data.

Using a sidecar or custom-built container for key retrieval and file decryption
integrates seamlessly with modern DevOps practices, particularly within GitOps 
workflows, offering an adaptable solution to security challenges.

This system enhances security and maintains operational efficiency by automating 
critical aspects of the encryption process.

The potential for **VSecM** to autonomously generate and manage AES keys further 
streamlines the process, reducing the need for manual intervention and minimizing 
human error. 

Future contributions to the **VSecM** project, as highlighted by the open issues, 
will undoubtedly enhance its capabilities and ease of integration.

By adopting this approach, organizations can strike an effective balance between 
security requirements and operational demands, paving the way for a more secure 
and efficient management of sensitive data.

## List of Use Cases in This Section

{{ use_cases() }}

{{ edit() }}
