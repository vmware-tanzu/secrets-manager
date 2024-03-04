#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

SPIRE_SERVER=$(microk8s kubectl get po -n spire-system \
  | grep "spire-server-" | awk '{print $1}')

microk8s kubectl exec -c spire-server -n spire-system $SPIRE_SERVER -- \
  /opt/spire/bin/spire-server bundle show -format spiffe
