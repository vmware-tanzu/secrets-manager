#!/usr/bin/env bash

SCOUT_DNS="vsecm-scout.vsecm-system.svc.cluster.local"

# Generate a private key
openssl genrsa -out server.key 2048

# Create a self-signed certificate
openssl req -new -x509 -sha256 -key server.key -out server.crt \
  -days 3650 -subj "/CN=${SCOUT_DNS}"

# Create a CA bundle (in this case, it's just the server certificate)
cp server.crt ca.crt

# Base64 encode the CA bundle
# shellcheck disable=SC2002
CA_BUNDLE=$(cat ca.crt | base64 -w 0)

# Update the cluster-secret-store.yaml file
# There is no sensitive info in it, so it can
# be sent to a remote cluster via GitOps.
cat << EOF > cluster-secret-store.yaml
apiVersion: external-secrets.io/v1beta1
kind: ClusterSecretStore
metadata:
  name: vsecm-scout
spec:
  provider:
    webhook:
      url: "https://${SCOUT_DNS}/webhook?key={{ .remoteRef.key }}"
      method: GET
      result:
        jsonPath: "$"
      caBundle: ${CA_BUNDLE}
EOF

echo "Certificates generated and ClusterSecretStore YAML updated."
echo "Please review the cluster-secret-store.yaml file."

# These will be stored in VSecM Safe and dynamically provided to VSecM Scout.
mv server.crt ./hack
mv server.key ./hack
mv ca.crt ./hack
