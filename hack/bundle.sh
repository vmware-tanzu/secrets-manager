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

PACKAGE="$1"
VERSION="$2"
DOCKERFILE="$3"

go mod vendor
docker build -f "${DOCKERFILE}" . -t "${PACKAGE}":"${VERSION}"

sleep 10