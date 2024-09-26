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

docker build -t control-plane-server .
docker tag control-plane-server localhost:32000/control-plane-server:latest
docker push localhost:32000/control-plane-server:latest
