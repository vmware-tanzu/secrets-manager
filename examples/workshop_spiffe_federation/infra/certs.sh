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

# Default values
HOST="10.211.55.111"
PORT="8443"

# Function to display usage information
usage() {
    echo "Usage: $0 [-h host] [-p port] [-a]"
    echo "  -h host   : Specify the host (default: $HOST)"
    echo "  -p port   : Specify the port (default: $PORT)"
    echo "  -a        : Show all certificate information"
    echo "  -s        : Show subject"
    echo "  -i        : Show issuer"
    echo "  -d        : Show validity dates"
    echo "  -f        : Show fingerprint"
    exit 1
}

# Parse command line options
while getopts "h:p:asidf" opt; do
    case $opt in
        h) HOST="$OPTARG" ;;
        p) PORT="$OPTARG" ;;
        a) ALL=true ;;
        s) SUBJECT=true ;;
        i) ISSUER=true ;;
        d) DATES=true ;;
        f) FINGERPRINT=true ;;
        *) usage ;;
    esac
done

# Base OpenSSL command
BASE_CMD="openssl s_client -connect $HOST:$PORT -servername $HOST < /dev/null 2>/dev/null"

# Function to execute OpenSSL command and filter output
execute_openssl() {
    local filter="$1"
    eval "$BASE_CMD | openssl x509 -noout $filter"
}

# Display requested information
if [ "$ALL" = true ]; then
    execute_openssl "-text"
else
    if [ "$SUBJECT" = true ]; then
        echo "Subject:"
        execute_openssl "-subject"
    fi
    if [ "$ISSUER" = true ]; then
        echo "Issuer:"
        execute_openssl "-issuer"
    fi
    if [ "$DATES" = true ]; then
        echo "Validity:"
        execute_openssl "-dates"
    fi
    if [ "$FINGERPRINT" = true ]; then
        echo "Fingerprint:"
        execute_openssl "-fingerprint"
    fi
fi

# If no options were selected, show usage
if [ "$ALL" != true ] && [ "$SUBJECT" != true ] && [ "$ISSUER" != true ] && [ "$DATES" != true ] && [ "$FINGERPRINT" != true ]; then
    usage
fi
