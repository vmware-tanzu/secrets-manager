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

VSECM_NS="$1"
SPIRE_NS="$2"
SPIRE_SERVER_NS="$3"

if kubectl get ns | grep vsecm-system; then
  # Order is important for SPIFFE SCI Driver to properly unmount volumes.
  # ref: https://github.com/spiffe/spiffe-csi#failure-to-terminate-pods-when-driver-is-unhealthy-or-removed
  kubectl delete ns $VSECM_NS
  kubectl delete ns $SPIRE_SERVER_NS
  kubectl delete ns $SPIRE_NS

  kubectl delete ClusterSPIFFEID example
  kubectl delete ClusterSPIFFEID vsecm-sentinel
  kubectl delete ClusterSPIFFEID vsecm-safe
  kubectl delete CSIDriver csi.spiffe.io
  kubectl delete ValidatingWebhookConfiguration spire-controller-manager-webhook
  kubectl delete clusterrolebinding vsecm-secret-readwriter-binding spire-server-spire-controller-manager spire-agent spire-server-spire-server
  kubectl delete clusterrole spire-agent spire-server-spire-server vsecm-secret-readwriter spire-server-spire-controller-manager

else
  echo "Nothing to clean."
fi

if kubectl delete deployment example -n default; then
  echo "Deleted demo workload too.";
else
  echo "No demo workload to delete?... No worries: That's fine.";
fi

echo "Everything is awesome!"
