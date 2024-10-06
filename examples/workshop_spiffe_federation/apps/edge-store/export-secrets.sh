#!/usr/bin/env bash

# Find the pod
S=$(microk8s kubectl get po -n vsecm-system \
  | grep "vsecm-sentinel-" | awk '{print $1}')

# Execute the command and save output directly to secrets.json in the project root
microk8s kubectl exec "$S" -n vsecm-system -- safe -l -e > ./secrets.json

# Inform the user
echo "Secrets have been saved to secrets.json in the current working directory."