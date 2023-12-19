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

cd ./examples/using-sdk || exit

kubectl apply -f ./k8s/ServiceAccount.yaml
kubectl apply -k ./k8s
kubectl apply -f ./k8s/Identity.yaml
kubectl apply -f ./k8s/Secret.yaml
