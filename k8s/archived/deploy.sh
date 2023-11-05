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

kubectl apply -k ./spire

echo "waiting for SPIRE server to be ready."
kubectl wait --for=condition=Ready pod -n spire-system --selector=app=spire-server
echo "waiting for SPIRE agent to be ready."
kubectl wait --for=condition=Ready pod -n spire-system --selector=app=spire-agent
echo "next"

cd safe || exit
kubectl apply -f ./Namespace.yaml
kubectl apply -f ./Role.yaml
if kubectl get secret -n vsecm-system | grep vsecm-safe-age-key; then
  echo "!!! The secret 'vsecm-safe-age-key' already exists; not going to override it."
  echo "!!! If you want to modify it, make sure you back it up first."
else
  kubectl apply -f ./Secret.yaml
fi
kubectl apply -f ./ServiceAccount.yaml
kubectl apply -f ./Identity.yaml
kubectl apply -f ./Service.yaml
kubectl apply -k ./kustomizations/remote/istanbul

cd ..
cd sentinel || exit
kubectl apply -f Namespace.yaml
kubectl apply -f Identity.yaml
kubectl apply -f ServiceAccount.yaml
kubectl apply -k ./kustomizations/remote/istanbul

echo "Everything is awesome!"
