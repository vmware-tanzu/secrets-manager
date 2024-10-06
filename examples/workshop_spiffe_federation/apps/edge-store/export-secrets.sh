#!/usr/bin/env bash

S=$(microk8s kubectl get po -n vsecm-system \
  | grep "vsecm-sentinel-" | awk '{print $1}')
export S=$S

microk8s kubectl exec "$S" -n vsecm-system -- safe -l -e
