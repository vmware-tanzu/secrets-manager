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

SPIRE_SERVER_NS="spire-server"

SPIRE_SERVER=$(kubectl get po -n $SPIRE_SERVER_NS \
  | grep $SPIRE_SERVER_NS | awk '{print $1}')
export SPIRE_SERVER=SPIRE_SERVER

kubectl exec -n $SPIRE_SERVER_NS $SPIRE_SERVER -- \
  /opt/spire/bin/spire-server bundle show
