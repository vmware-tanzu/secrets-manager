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

cd ./ci/poll || exit
go build -o vsecm-ci
rm -rf /opt/vsecm/vsecm-ci
mv vsecm-ci /opt/vsecm
