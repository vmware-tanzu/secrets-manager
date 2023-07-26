#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

VERSION="$1"

echo ""
echo "--------"
echo "VSecM"
if git tag -s v"$VERSION"; then
  git push origin --tags
  gh release create
fi

echo "vsecm-web"
cd ../vsecm-web || exit
if git tag -s v"$VERSION"; then
  git push origin --tags
  gh release create
fi

echo "Everything is awesome!"
