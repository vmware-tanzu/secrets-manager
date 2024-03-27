# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

export SECRET="ComputeMe!"

SENTINEL=$(kubectl get po -n vsecm-system \
  | grep "vsecm-sentinel-" | awk '{print $1}')
export SENTINEL=$SENTINEL

SAFE=$(kubectl get po -n vsecm-system \
  | grep "vsecm-safe-" | awk '{print $1}')
export SAFE=$SAFE

WORKLOAD=$(kubectl get po -n default \
  | grep "example-" | awk '{print $1}')
export WORKLOAD=$WORKLOAD

INSPECTOR=$(kubectl get po -n default \
  | grep "vsecm-inspector-" | awk '{print $1}')
export INSPECTOR=$INSPECTOR

export DEPLOYMENT="example"
