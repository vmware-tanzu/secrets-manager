#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# *

#!/bin/bash

if command -v helm-docs >/dev/null 2>&1; then
    echo "Running helm-docs..."
    helm-docs
else
    echo "Warning: helm-docs is not installed. Please install it to generate documentation."
fi
