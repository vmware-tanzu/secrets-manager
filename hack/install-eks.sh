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

# It is best to deploy EKS with helm. Otherwise there remains a bunch of
## resources that are not deleted when helm deletes the release, and
# the next deployment will fail.
#
# This script takes care of deleting dangling resources after a `helm delete`
# and then re-deploys vsecm.

helm delete vsecm

NAMESPACE="vsecm-system"

# Due to SPIFFE CSI Driver, pods in Terminating state may not be deleted.
# We’ll need to forcefully terminate them.
PODS=$(kubectl get pods -n "$NAMESPACE" \
  --field-selector=status.phase=Terminating -o name)
if [ -z "$PODS" ]; then
  echo "No pods in Terminating state found in namespace $NAMESPACE."
  echo "This is perfectly fine."
else
  for pod in $PODS; do
    echo "Force deleting $pod"
    kubectl delete $pod -n "$NAMESPACE" --grace-period=0 --force
  done
fi
kubectl delete namespace "$NAMESPACE" --grace-period=0 --froce

helm install vsecm vsecm/vsecm

echo "verifying vsecm installation"
kubectl wait --for=condition=Available deployment -n vsecm-system vsecm-sentinel
echo "vsecm-sentinel: deployment available"
kubectl wait --for=condition=Available deployment -n vsecm-system vsecm-safe
echo "vsecm-safe: deployment available"
echo "vsecm installation successful"
