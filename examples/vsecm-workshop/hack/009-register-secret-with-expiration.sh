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

. ./env.sh

kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "example" \
  -n "default" \
  -s "VSecMRocks!" \
  -N "2019-10-12T07:20:50.52Z" \
  -E "2042-10-12T08:21:51.53Z"
