#!/usr/bin/env bash

# Read certificate files
SERVER_CRT=$(cat hack/server.crt | base64 -w 0)
SERVER_KEY=$(cat hack/server.key | base64 -w 0)
CA_CRT=$(cat hack/ca.crt | base64 -w 0)

# Export variables
export SERVER_CRT
export SERVER_KEY
export CA_CRT

# Identify VSecM Pod
S=$(kubectl get po -n vsecm-system \
  | grep "vsecm-sentinel-" | awk '{print $1}')
export S=$S

# Add cert and key to vsecm
k exec -n vsecm-system $S -- safe \
  -w "raw:vsecm-eso-webhook-server-crt" \
	-s "$SERVER_CRT"

k exec -n vsecm-system $S -- safe \
  -w "raw:vsecm-eso-webhook-server-key" \
	-s "$SERVER_KEY"

k exec -n vsecm-system $S -- safe \
  -w "raw:vsecm-eso-webhook-ca-crt" \
	-s "$CA_CRT"


# This is the demo secret:
k exec -n vsecm-system $S -- safe \
  -w "raw:coca-cola.cluster-001" \
  -s '{"namespaces": {"coke-system": {"secrets":{"admin-credentials":{"type":"k8s","value":"super-secret-secret","metadata": {"labels": {"managed-by": "coke-system"},"annotations": {"inject-sidecar": "true"},"creationTimestamp": "2024-01-01","lastUpdated": "2024-01-01"},"expires": "2024-01-01","notBefore": "2024-01-01"}}}}}'

