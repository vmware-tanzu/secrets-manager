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

docker build -t edge-store .
docker tag edge-store localhost:32000/edge-store:latest
docker push localhost:32000/edge-store:latest
