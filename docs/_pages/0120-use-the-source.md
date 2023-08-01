---
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title: Building From Source
layout: post
next_url: /docs/contributing/
prev_url: /docs/configuration/
permalink: /docs/use-the-source/
---

This section describes how to build **VMware Secrets Manager** from source.

For a more detailed walkthrough about how to contribute to **VMware Secrets
Manager**, see the [**Contributing**](/docs/contributing/) section.

## Prerequisites

Make you have the following installed on your system:

* [`make`](https://www.gnu.org/software/make/)
* [`git`](https://git-scm.com/)
* [`docker`](https://www.docker.com/)

## Clone the Project

```bash
cd $WORKSPACE
git clone https://github.com/vmware-tanzu/secrets-manager.git
cd secrets-manager
```

## Build the Project

Make sure you have a running local Docker daemon and execute the following:

```bash
make build-local
```

That’s it 🎉. You now have images of **VMware Secrets Manager** and other 
related components built locally on your Docker registry.

## Next Up

For a more detailed guide about how you can use these local container images
in your custer [check out the **Contributing** section](/docs/contributing/).




















