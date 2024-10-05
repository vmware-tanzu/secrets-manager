#!/usr/bin/env bash

set -e

# Default values
NAMESPACE="vsecm-system"
SECRET_NAME="vsecm-root-key"
KEY_FILE="key.txt"

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
current_secret=$(kubectl get secret "$SECRET_NAME" -n "$NAMESPACE" -o yaml)

# Update the secret
updated_secret=$(echo "$current_secret" | awk -v key="$encoded_key" '
    /KEY_TXT:/ {
        print "  KEY_TXT: " key
        next
    }
    {print}
')

# Apply the updated secret
echo "$updated_secret" | kubectl apply -f -

echo "Secret $SECRET_NAME in namespace $NAMESPACE updated successfully."

# Rotate the StatefulSet
echo "Rotating StatefulSet $STATEFULSET_NAME..."
kubectl rollout restart statefulset "$STATEFULSET_NAME" -n "$NAMESPACE"

# Wait for the rollout to complete
kubectl rollout status statefulset "$STATEFULSET_NAME" -n "$NAMESPACE" --timeout=5m

echo "StatefulSet $STATEFULSET_NAME rotated successfully."
echo "Secret update and StatefulSet rotation completed."
