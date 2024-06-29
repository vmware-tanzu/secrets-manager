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

title = "Security"
weight = 11
+++

> **Secure All the Things 🛡️**
>
> This section contains how we handle security in **VSecM**.

## Vulnerability Management And Remediation Policy

Ensuring the security of **VMware Secrets Manager** is a top priority and
a continuous process that involves various proactive and reactive measures.
On such measure is promptly addressing the security vulnerabilities that
are found in the codebase.

All medium and higher severity exploitable vulnerabilities discovered with
dynamic code analysis **MUST** be fixed in a timely way after they are confirmed.
This is mandatory to maintain the highest security standards for the project.

For additional context, here is how we categorize the severity of the
vulnerabilities:

* **Medium Severity**: A vulnerability that allows unauthorized disclosure of
  information or unauthorized modification of a system.
* **High Severity**: A vulnerability that allows unauthorized administrative  
  access to a system.

## Target Response Times

Our target time frame is **60 days** for fixing vulnerabilities of **medium or higher
severity**. That is to say, we do our best so that our codebase does not contain
any public vulnerabilities of medium or higher severity that have been
known for more than 60 days.

In addition, our target initial response time for **any** vulnerability report
received in the **last 6 months** is **less than or equal to 14 days**.

## About Keys and Nonce Values

Cryptography plays a vital role in securing both data at rest and in transit.
Ensuring that cryptographic keys and nonces are generated securely is fundamental
to maintaining a robust security posture.

**WMware Secrets Manager** generates all cryptographic keys and nonces that
it needs using secure random generators. In the event that this practice is not
followed, a bug **MUST** be filed to address this issue immediately.

## About Tests

High-quality software is not just about features and performance; it's also
about reliability and predictability. Automated testing is a cornerstone in
achieving these goals.

Here is an informal guideline about how we approach testing in
**VMware Secrets Manager**:

As major new functionality is added to the software produced by this project,
corresponding tests **MUST** be added to an automated test suite.

In this context, "**Major New Functionality**" means features or changes that
*substantially* alter the behavior, capabilities, or user experience of the
software.

## Report a Security Vulnerability

We take **VMware Secrets Manager**'s security seriously. If you believe you have
found a vulnerability, please [**follow this guideline**][vuln]
to responsibly disclose it.

[vuln]: https://github.com/vmware-tanzu/secrets-manager/blob/main/SECURITY.md

{{ edit() }}
