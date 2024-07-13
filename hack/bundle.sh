#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

PACKAGE="$1"
VERSION="$2"
DOCKERFILE="$3"
gitRoot=$(git rev-parse --show-toplevel)

# Check if docker binary is present
if ! command -v docker &> /dev/null
then
    echo "Docker binary could not be found. Please install Docker first."
    exit 1
fi

# Change directory to the root of the git repository.
cd "$gitRoot" || exit 1

if [ ! -d "./vendor" ]; then
    # vendor directory doesn't exist.
    echo "vendor directory doesn't exist. Please run 'go mod vendor' to create vendor directory."
    exit 1
fi

docker build -f "${DOCKERFILE}" . -t "${PACKAGE}":"${VERSION}"
