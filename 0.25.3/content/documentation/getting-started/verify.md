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

title = "Verify Installation"
weight = 15
+++

## Verify Installation

Let's check that the installation was successful by listing the pods int
the `spire-system` , `vsecm-system` and `keystone-system` namespaces:

```bash
kubectl get po -n spire-system
```

```txt
# Output:
NAME                          READY   STATUS    RESTARTS      
spire-agent-wdhdh             3/3     Running   0             
spire-server-b594bdfc-szssx   2/2     Running   0             
```

```bash
kubectl get po -n vsecm-system
```

```txt
# Output:
NAME                              READY   STATUS    RESTARTS
vsecm-keystone-c54d99d7b-c4jk4    1/1     Running   0          
vsecm-safe-6cc477f58f-x6wc9       1/1     Running   0          
vsecm-sentinel-74d648675b-8zdn2   1/1     Running   0                   
```

All the pods look up and running, so we can move on to the next step.

## List Available Commands

`make help` lists all the available make targets in a cheat sheet format:

```bash
make help

# The output will vary as we add more commands to the Makefile.
# It will contain useful information about the available commands.
```

{{ edit() }}
