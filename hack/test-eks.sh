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

# TODO:
# 0. switch context to eks
# 0.5. copy former version of k8s files to ./k8s/$formerVersion
# 1. wipe out vsecm-system and spire-system namespaces
# 2. deploy
# 3. test
# 4. run this at the end of `make test` if VSECM_EKS_CTX is set
# 5. cleanup (remove ./k8s/$formerVersion)
