#!/bin/bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# NOTE:
#
# If, for whatever reason, `ClusterFederatedTrustDomain` fails to share the
# trust bundle with the SPIRE server, we can use this script to manually set
# the trust bundle on the SPIRE server.
#
# Since this is a demo setup, I am not investigating the root cause of the
# issue, but rather providing a workaround to set the trust bundle manually.

# Function to extract trust bundle from ClusterFederatedTrustDomain
get_trust_bundle() {
    local name=$1
    microk8s kubectl get clusterfederatedtrustdomain $name -o jsonpath='{.spec.trustDomainBundle}'
}

# Function to set the trust bundle
set_trust_bundle() {
    local name=$1
    local trust_domain=$2
    local bundle_data=$3

    # Create a temporary file to store the bundle data
    tmp_file=$(mktemp)
    echo "$bundle_data" > $tmp_file

    # Set the bundle using spire-server command, passing data via stdin
    echo "$bundle_data" | microk8s kubectl exec -i spire-server-0 -n spire-server -- /opt/spire/bin/spire-server bundle set \
        -id "spiffe://$trust_domain" \
        -format spiffe

    # Clean up the temporary file
    #rm $tmp_file
}

# Main script
main() {
    # List of trust domains to process
    trust_domains=("diablo.vsecm.com")

    for domain in "${trust_domains[@]}"; do
        echo "Processing $domain..."
        
        # Extract the name from the trust domain
        name=$(echo $domain | cut -d. -f1)
        
        # Get the trust bundle
        bundle_data=$(get_trust_bundle $name)
        
        if [ -z "$bundle_data" ]; then
            echo "Failed to retrieve trust bundle for $domain"
            continue
        fi

        # Set the trust bundle
        set_trust_bundle $name $domain "$bundle_data"
        
        echo "Completed processing for $domain"
    done
}

# Run the main function
main
