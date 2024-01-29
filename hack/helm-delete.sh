#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# *

set -e

check_namespace_deleted() {
  local namespace=$1
  local max_attempts=30
  local attempt=1
  local sleep_time=10

  echo "Checking for deletion of namespace: $namespace"
  while [[ $(kubectl get namespace "$namespace" --ignore-not-found) ]]; do
    if (( attempt > max_attempts )); then
      echo "Namespace $namespace still exists after $max_attempts attempts. Exiting with error."
      exit 1
    fi

    echo "Waiting for namespace $namespace to be deleted... Attempt $attempt/$max_attempts"
    sleep $sleep_time
    ((attempt++))
  done
  echo "Namespace $namespace deleted successfully."
}

if kubectl get deployment vsecm-sentinel -n vsecm-system; then
  kubectl delete deployment vsecm-sentinel -n vsecm-system || \
    { echo "Failed to delete vsecm-sentinel deployment"; exit 1; }
  kubectl wait --for=delete pod -l app=vsecm-sentinel -n vsecm-system --timeout=60s || \
    { echo "Timeout or error while waiting for vsecm-sentinel pods to delete"; exit 1; }
else
  echo "vsecm-sentinel deployment does not exist. Skipping delete."
fi

if kubectl get deployment vsecm-safe -n vsecm-system; then
  kubectl delete deployment vsecm-safe -n vsecm-system || \
    { echo "Failed to delete vsecm-safe deployment"; exit 1; }
  kubectl wait --for=delete pod -l app=vsecm-safe -n vsecm-system --timeout=60s || \
    { echo "Timeout or error while waiting for vsecm-safe pods to delete"; exit 1; }
else
  echo "vsecm-safe deployment does not exist. Skipping delete."
fi

if helm list --namespace vsecm-system | grep -q 'vsecm'; then
  helm delete vsecm || { echo "Failed to delete Helm release vsecm"; exit 1; }
else
  echo "Helm release vsecm does not exist. Exiting script with success status."
  exit 0
fi

check_namespace_deleted vsecm-system
check_namespace_deleted spire-system

# Just to be safe…
echo "Will wait for 30 seconds to allow k8s to drain any remaining resources"
sleep 30
