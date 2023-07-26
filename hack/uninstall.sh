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

if kubectl get ns | grep vsecm-system; then
  # Order is important for SPIFFE SCI Driver to properly unmount volumes.
  # ref: https://github.com/spiffe/spiffe-csi#failure-to-terminate-pods-when-driver-is-unhealthy-or-removed
  kubectl delete ns vsecm-system
  kubectl delete ns spire-system

  kubectl delete ClusterSPIFFEID example
  kubectl delete ClusterSPIFFEID vsecm-sentinel
  kubectl delete ClusterSPIFFEID vsecm-safe
else
  echo "Nothing to clean."
fi

if kubectl delete deployment example -n default; then
  echo "Deleted demo workload too.";
else
  echo "No demo workload to delete?… No worries: That’s fine.";
fi

echo "Everything is awesome!"
