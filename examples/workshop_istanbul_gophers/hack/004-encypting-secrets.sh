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

source ./env.sh

secret=$(kubectl exec "$SENTINEL" -n vsecm-system -- safe \
	-s '{"username": "*root*", "password": "*Ca$#C0w*", "value": "🎸 İstanbul Gophers 🤘"}' \
	-e 2>&1)

trimmed_secret=$(echo "$secret" | tr -d '\n')

echo "secret='$trimmed_secret'"

echo "export SECRET='$trimmed_secret'" > secret.sh
