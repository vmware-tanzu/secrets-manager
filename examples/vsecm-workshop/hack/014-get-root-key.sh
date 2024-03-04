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

echo ""
kubectl get secret vsecm-root-key -n vsecm-system \
  -o jsonpath="{.data.KEY_TXT}" | base64 --decode
echo ""
