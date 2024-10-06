#!/usr/bin/env bash

S=$(microk8s kubectl get po -n vsecm-system \
  | grep "vsecm-sentinel-" | awk '{print $1}')
export S=$S

microk8s kubectl exec "$S" -n vsecm-system -- safe \
  -w "vsecm-relay:mephisto.vsecm.com" \
  -n "vsecm-restricted" \
  -s 'gen:{"vsecm:for":"mephisto", "vsecm:workload": "k8s:admin-password", "vsecm:namespace": "default", "username": "admin-[a-z0-9]{6}", "password": "[a-zA-Z0-9]{12}"}'

microk8s kubectl exec "$S" -n vsecm-system -- safe \
  -w "vsecm-relay:baal.vsecm.com" \
  -n "vsecm-restricted" \
  -s 'gen:{"vsecm:for":"baal", "vsecm:workload": "k8s:admin-password", "vsecm:namespace": "default", "username": "admin-[a-z0-9]{6}", "password": "[a-zA-Z0-9]{12}"}'

microk8s kubectl exec "$S" -n vsecm-system -- safe \
  -w "vsecm-relay:azmodan.vsecm.com" \
  -n "vsecm-restricted" \
  -s 'gen:{"vsecm:for":"azmodan", "vsecm:workload": "k8s:admin-password", "vsecm:namespace": "default", "username": "admin-[a-z0-9]{6}", "password": "[a-zA-Z0-9]{12}"}'
