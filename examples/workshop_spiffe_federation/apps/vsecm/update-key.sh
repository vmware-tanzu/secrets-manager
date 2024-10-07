#!/usr/bin/env bash

set -e

# Default values
NAMESPACE="vsecm-system"
SECRET_NAME="vsecm-root-key"
KEY_FILE="key.txt"
STATEFULSET_NAME="vsecm-safe"

# Function to display usage information
usage() {
    echo "Usage: $0 [-n <namespace>] [-s <secret_name>] [-k <key_file>]"
    echo "  -n: Kubernetes namespace (default: $NAMESPACE)"
    echo "  -s: Name of the secret (default: $SECRET_NAME)"
    echo "  -k: Path to the file containing the new key (default: $KEY_FILE)"
    exit 1
}

# Parse command line arguments
while getopts ":n:s:k:" opt; do
    case $opt in
        n) NAMESPACE="$OPTARG" ;;
        s) SECRET_NAME="$OPTARG" ;;
        k) KEY_FILE="$OPTARG" ;;
        \?) echo "Invalid option -$OPTARG" >&2; usage ;;
    esac
done

# Check if key file exists
[ ! -f "$KEY_FILE" ] && echo "Error: Key file not found: $KEY_FILE" && exit 1

# Read the new key and trim whitespace
new_key=$(tr -d '[:space:]' < "$KEY_FILE")

# Base64 encode the new key
# encoded_key=$(echo -n "$new_key" | base64 | tr -d '\n')
encoded_key=$new_key

# Get the current secret
current_secret=$(microk8s kubectl get secret "$SECRET_NAME" -n "$NAMESPACE" -o yaml)

# Update the secret
updated_secret=$(echo "$current_secret" | awk -v key="$encoded_key" '
    /KEY_TXT:/ {
        print "  KEY_TXT: " key
        next
    }
    {print}
')

# Apply the updated secret
echo "$updated_secret" | microk8s kubectl apply -f -

echo "Secret $SECRET_NAME in namespace $NAMESPACE updated successfully."

# Check if StatefulSet exists
if ! microk8s kubectl get statefulset "$STATEFULSET_NAME" -n "$NAMESPACE" &> /dev/null; then
    echo "Error: StatefulSet $STATEFULSET_NAME not found in namespace $NAMESPACE"
    exit 1
fi

# Rotate the StatefulSet
echo "Rotating StatefulSet $STATEFULSET_NAME..."
if ! microk8s kubectl rollout restart statefulset "$STATEFULSET_NAME" -n "$NAMESPACE"; then
    echo "Error: Failed to restart StatefulSet $STATEFULSET_NAME"
    exit 1
fi

# Wait for the rollout to complete
echo "Waiting for StatefulSet rollout to complete..."
if ! microk8s kubectl rollout status statefulset "$STATEFULSET_NAME" -n "$NAMESPACE" --timeout=5m; then
    echo "Error: StatefulSet rollout did not complete within the timeout period"
    exit 1
fi

echo "StatefulSet $STATEFULSET_NAME rotated successfully."
echo "Secret update and StatefulSet rotation completed."
