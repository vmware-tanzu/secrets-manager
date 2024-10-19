#!/usr/bin/env bash

# Generate a private key
openssl genrsa -out server.key 2048

# Create a self-signed certificate
openssl req -new -x509 -sha256 -key server.key -out server.crt \
  -days 3650 -subj "/CN=eso-webhook.default.svc.cluster.local"

# Create a CA bundle (in this case, it's just the server certificate)
cp server.crt ca.crt

# Base64 encode the CA bundle
# shellcheck disable=SC2002
CA_BUNDLE=$(cat ca.crt | base64 -w 0)

# Update the cluster-secret-store.yaml file
cat << EOF > cluster-secret-store.yaml
apiVersion: external-secrets.io/v1beta1
kind: ClusterSecretStore
metadata:
  name: webhook-backend
spec:
  provider:
    webhook:
      url: "https://eso-webhook.default.svc.cluster.local:8443/webhook?key={{ .remoteRef.key }}"
      method: GET
      result:
        jsonPath: "$"
      caBundle: ${CA_BUNDLE}
EOF

echo "Self-signed certificate generated and ClusterSecretStore YAML updated."
echo "Please review the cluster-secret-store.yaml file."

mv server.crt ./hack
mv server.key ./hack
mv ca.crt ./hack
