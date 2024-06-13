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

SEARCH_DIR="."

env_vars=$(find "$SEARCH_DIR" -type f -name "*.go" \
  -not -path "*/Makefile" \
  -not -path "*/makefiles/*" \
  -not -path "*/vendor/*" \
  -not -path "*/examples/*" \
  -not -path "*/ci/test/*" \
  -exec grep -H 'os.Getenv' {} + | sed -n 's/.*os.Getenv("\([^"]*\)").*/\1/p')

sorted_env_vars=$(echo "$env_vars" | sort | uniq)

echo $sorted_env_vars | tr ' ' '\n'

