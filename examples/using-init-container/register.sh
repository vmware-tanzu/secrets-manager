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

SENTINEL=$(kubectl get po -n vsecm-system \
  | grep "vsecm-sentinel-" | awk '{print $1}')

# To make the init container exit successfully and initialize the main
# container of the Pod, execute the following script.
#
# This will create a Kubernetes `Secret` that the `main` container is
# injecting as an environment variable and let the container consume
# that `Secret`.
#
# -n : identifies the namespace of the Kubernetes `Secret`.
# -k : means VSecM will update an associated Kubernetes Secret.
# -t : will be used to transform the fields of the payload.
# -s : is the actual value of the secret.

kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -w "example" \
  -n "default" \
  -s '{"username": "root", "password": "SuperSecret", "value": "VSecMRocks"}' \
  -t '{"USERNAME":"{{.username}}", "PASSWORD":"{{.password}}", "VALUE": "{{.value}}"}' \
  -k
