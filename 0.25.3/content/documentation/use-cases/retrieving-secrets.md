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

title = "Retrieving Secrets"
weight = 193
+++

## Situation Analysis

**VMware Secrets Manager** does not allow direct access to the secrets stored in 
it. Only the workloads that are authorized to access the secrets can retrieve them.
Even an operator cannot read secrets in plain text because **VSecM Sentinel**
does not have an API to display plain text secrets.

This behavior is by design to prevent unauthorized access to secrets.

However, as an operator, you may need to retrieve secrets for debugging or
troubleshooting purposes. If you have the necessary permissions, it is possible
to export secrets in encrypted form, and then decrypt them using 
**VSecM Keygen**.

## High-Level Diagram

Open the image in a new tab to see the full-size version:

![High-Level Diagram](/assets/reveal-secrets.png "High-Level Diagram")

## Implementation

Let's start by exporting our secrets in encrypted form.

### Exporting Secrets

To export a secret in encrypted form, you can use the **VSecM CLI** as follows:

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe -l -e | tee secrets.json
```

### Getting Root Keys From the Cluster

If you are a cluster admin, you can get the **VSecM** root keys from the cluster
using the following command:

```bash
kubectl get secret vsecm-root-key -n vsecm-system \
  -o jsonpath="{.data.KEY_TXT}" | base64 --decode
```

Save the output in a file named `key.txt`.

### Decrypting Secrets

Since you have `secrets.json` and `key.txt`, you can now decrypt the secrets
using **VSecM Keygen**:

```bash
# Make sure that you have `./key.txt` and `./secrets.json` in the 
# current directory.
docker run --rm \
  -v "$(pwd)":/vsecm \
  -e VSECM_KEYGEN_EXPORTED_SECRET_PATH="/vsecm/secrets.json" \
  -e VSECM_KEYGEN_ROOT_KEY_PATH="/vsecm/key.txt" \
  -e VSECM_KEYGEN_DECRYPT="true" \
  vsecm/vsecm-keygen:latest
```

The output of this command will list the values of all the secrets 
in plain text.

## Conclusion

In conclusion, **VSecM** provides a robust framework for managing sensitive 
information securely by restricting direct access to secrets. This design ensures 
that only authorized workloads can retrieve the secrets and prevents operators 
from accessing them in plain text. 

For necessary operations such as debugging or troubleshooting, operators with 
the appropriate permissions can still access these secrets by exporting them in 
an encrypted form and subsequently decrypting them using the VSecM Keygen tool.

The process involves several steps:

This system's architecture not only enhances security by minimizing the risk of 
unauthorized secret access but also maintains flexibility for administrators, 
thus supporting essential maintenance and operational activities without 
compromising on the principles of secure information management.

## List of Use Cases in This Section

{{ use_cases() }}

{{ edit() }}
