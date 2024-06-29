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

title = "Generating Randomized Secrets"
weight = 247
+++

## Situation Analysis

Oftentimes, you need to generate randomized secrets for your applications for
security reasons. Because if the secret is randomized, no one, including the
operator that created it will know the secret. 

This is important because if the operator knows the secret, they become part of 
your trust boundary, and you will have to ensure that they will not willingly or 
unwillingly disclose the secret. Randomizing the secret removes the operator 
from the trust boundary, enhancing security.

Luckily, **VMware Secrets Manager** allows you to generate pattern-based
randomized secrets.

## High-Level Diagram

Open the image in a new tab to see the full-size version:

![High-Level Diagram](/assets/generate.png "High-Level Diagram")


## Implementation

You will just need to define a regex-like pattern and prefix your secrets
with `gen:` (*for "generate"*) to let **VSecM Sentinel** know that you want to
generate a randomized secret.

Here is an example:

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w example \
  -n default \
  -s 'gen:{"username":"admin-[a-z0-9]{6}","password":"[a-zA-Z0-9]{12}"}'
  -t '{"ADMIN_USER":"{{.username}}","ADMIN_PASS":"{{.password}}"}'
```

Check out [**VSecM CLI reference**](@/documentation/usage/cli.md) for more details.

## Conclusion

In conclusion, **VMware Secrets Manager** provides a robust solution for 
enhancing security through the generation of pattern-based randomized secrets. 
This feature is particularly beneficial as it excludes the operator from the trust 
boundary by ensuring that no one, not even the creator of the secret, 
knows the actual values, thereby mitigating the risk of intentional or accidental 
disclosure. 

By defining a regex-like pattern and using the prefix gen: in the secret creation 
process, users can easily configure the system to automatically generate and 
manage secrets that comply with specified complexity requirements.

This capability is crucial for maintaining stringent security protocols in 
environments where the integrity and confidentiality of access credentials are 
paramount. 

The straightforward implementation process, which involves a few 
command-line inputs to specify the desired patterns for usernames and passwords, 
ensures that application security can be fortified with minimal operational 
overhead. Additionally, the ability to customize patterns allows organizations 
to adhere to their specific security policies effectively.

Overall, **VSecM**'s approach not only simplifies the management of secrets but 
also significantly enhances the security posture of applications by reducing the 
attack surface associated with human factors in the creation and handling of 
sensitive information.

## List of Use Cases in This Section

{{ use_cases() }}

{{ edit() }}
