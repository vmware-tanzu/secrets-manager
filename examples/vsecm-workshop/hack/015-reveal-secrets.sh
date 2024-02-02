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

eval $(minikube docker-env -u)

docker run --rm \
  -v "$(pwd)":/vsecm \
  -e VSECM_KEYGEN_EXPORTED_SECRET_PATH="/vsecm/secrets.json" \
  -e VSECM_KEYGEN_ROOT_KEY_PATH="/vsecm/key.txt" \
  -e VSECM_KEYGEN_DECRYPT="true" \
  vsecm/vsecm-keygen:0.22.2
