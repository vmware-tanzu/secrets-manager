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

cd ./examples/using_sdk_go || exit

kubectl apply -f ./k8s-eks/ServiceAccount.yaml
kubectl apply -k ./k8s-eks
kubectl apply -f ./k8s-eks/Identity.yaml
kubectl apply -f ./k8s-eks/Secret.yaml
