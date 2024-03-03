/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

SPIRE_SERVER=$(kubectl get po -n spire-system \
  | grep "spire-server-" | awk '{print $1}')
export SPIRE_SERVER=SPIRE_SERVER

kubectl exec -n spire-system $SPIRE_SERVER -- \
  /opt/spire/bin/spire-server bundle show
