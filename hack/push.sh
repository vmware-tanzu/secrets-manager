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
REPO="$3"

# Push version tag
docker tag "${PACKAGE}":"${VERSION}" "${REPO}":"${VERSION}"
docker push "${REPO}":"${VERSION}"

# Push latest tag
docker tag "${PACKAGE}":"${VERSION}" "${REPO}":latest
docker push "${REPO}":latest

sleep 10
