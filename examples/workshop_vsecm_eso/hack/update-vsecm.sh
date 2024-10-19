#!/usr/bin/env bash

# Read certificate files
SERVER_CRT=$(cat hack/server.crt | base64 -w 0)
SERVER_KEY=$(cat hack/server.key | base64 -w 0)
CA_CRT=$(cat hack/ca.crt | base64 -w 0)

# Generate JWT secret
JWT_SECRET=$(openssl rand -base64 32)

# Export variables
export SERVER_CRT
export SERVER_KEY
export CA_CRT
export JWT_SECRET

# Identify VSecM Pod
S=$(kubectl get po -n vsecm-system \
  | grep "vsecm-sentinel-" | awk '{print $1}')
export S=$S

# Add cert and key to vsecm
kubectl exec -n vsecm-system $S -- safe \
  -w "raw:vsecm-scout-crt" \
	-s "$SERVER_CRT"

kubectl exec -n vsecm-system $S -- safe \
  -w "raw:vsecm-scout-key" \
	-s "$SERVER_KEY"

kubectl exec -n vsecm-system $S -- safe \
  -w "raw:vsecm-scout-ca-crt" \
	-s "$CA_CRT"

# Add JWT secret
kubectl exec -n vsecm-system $S -- safe \
  -w "raw:vsecm-scout-jwt-secret" \
  -s "$JWT_SECRET"

# This is the demo secret:
kubectl exec -n vsecm-system $S -- safe \
  -w "raw:coca-cola.cluster-001" \
  -s '{"namespaces": {"cokeSystem": {"secrets":{"adminCredentials":{"type":"k8s","value":"super-secret-secret","metadata": {"labels": {"managedBy": "coke-system"},"annotations": {"injectSidecar": "true"},"creationTimestamp": "2024-01-01","lastUpdated": "2024-01-01"},"expires": "2024-01-01","notBefore": "2024-01-01"}}}}}'
