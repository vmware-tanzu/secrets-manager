---
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title: Mutating a Template File
layout: post
prev_url: /docs/use-case-secrets-as-volumes/
permalink: /docs/use-case-in-memory-template/
next_url: /docs/use-case-sdk-retrieve-secrets/
---

## Situation Analysis

Certain apps may require initialization scripts, which may include secrets. 
Storing these scripts with hard-coded secrets is a security gap. Storing
these scripts in source control is a security incident waiting to happen.

## Solution

A solution would be to create a template file with a placeholder to interpolate 
the secrets at deployment time.

As long as this template file is in an ephemeral "*in-memory*" volume and direct 
access to the workload is prevented by strict RBAC rules, we can consider the 
script and the secrets within it secure because data in an in-memory file system 
will be protected by the operating system's built-in memory barriers: 
Only an app that can shell into the Pod can access the in-memory volume.

## Strategy

Follow the [Mounting Secrets as Volumes][secrets-as-volumes] use case 
and configure the sidecar to mutate the file you need accordingly.

[secrets-as-volumes]:  /docs/secrets-as-volumes/
