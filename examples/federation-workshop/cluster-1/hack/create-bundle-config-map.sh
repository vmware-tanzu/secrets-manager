#!/usr/bin/env bash

microk8s kubectl create configmap trust-bundle \
  --from-file=./vsecm-002-bundle.json -n spire-system