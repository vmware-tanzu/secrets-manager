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

# Enable strict error checking.
set -euo pipefail

# TODO: temporarily disabled origin checks; will implement in a follow-up PR.
#ORIGIN=${1:-"local"}
#if [[ "$ORIGIN" != "remote" && "$ORIGIN" != "eks" ]]; then
#  ORIGIN="local"
#fi
#
#CI="$2"

go run ./ci/test/main.go ./ci/test/run.go
