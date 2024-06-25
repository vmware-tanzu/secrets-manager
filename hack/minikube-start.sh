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

# Minikube might need additional flags for SPIRE to work properly.
# A bare-metal or cloud Kubernetes cluster will not need these extra configs.
minikube start \
    --extra-config=apiserver.service-account-signing-key-file=/var/lib/minikube/certs/sa.key \
    --extra-config=apiserver.service-account-key-file=/var/lib/minikube/certs/sa.pub \
    --extra-config=apiserver.service-account-issuer=api \
    --extra-config=apiserver.api-audiences=api,spire-server,spire-server-custom \
    --extra-config=apiserver.authorization-mode=Node,RBAC \
    --memory="$MEMORY" \
    --cpus="$CPU" \
    --nodes="$NODES" \
    --insecure-registry "10.0.0.0/24"

echo "waiting 10 secs before enabling registry"
sleep 10
minikube addons enable registry
kubectl get node
