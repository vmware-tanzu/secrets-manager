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
  -w "example" \
  -n "default" \
  -s '{"username": "root", "password": "SuperSecret", "value": "VSecMRocks"}' \
  -t '{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}", "VALUE": "{{.value}}"}' \
  -k
