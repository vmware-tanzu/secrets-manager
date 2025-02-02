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

# Check if minikube binary is present
if ! command -v minikube &> /dev/null
then
    echo "Command 'minikube' not found. Please install Minikube first."
    exit 1
fi

minikube delete