#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

. ./env.sh

kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -s '{"username": "*root*", "password": "*Ca$#C0w*", "value": "!VSecMRocks!"}' \
  -e
