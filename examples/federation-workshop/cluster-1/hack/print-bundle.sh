#!/usr/bin/env bash

SPIRE_SERVER=$(microk8s kubectl get po -n spire-system \
  | grep "spire-server-" | awk '{print $1}')

microk8s kubectl exec -c spire-server -n spire-system $SPIRE_SERVER -- \
  /opt/spire/bin/spire-server bundle show -format spiffe
