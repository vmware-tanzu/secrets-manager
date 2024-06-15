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

# It is best to deploy EKS with helm. Otherwise there remains a bunch of
## resources that are not deleted when helm deletes the release, and
# the next deployment will fail.
#
# This script takes care of deleting dangling resources after a `helm delete`
# and then re-deploys vsecm.

helm install vsecm vsecm/vsecm

echo "verifying vsecm installation"
kubectl wait --timeout=120s --for=condition=Available deployment -n vsecm-system vsecm-sentinel
echo "vsecm-sentinel: deployment available"
kubectl wait --timeout=120s --for=condition=Available deployment -n vsecm-system vsecm-safe
echo "vsecm-safe: deployment available"
echo "vsecm installation successful"
