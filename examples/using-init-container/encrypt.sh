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

# Replace 5ADAE85E0C301567 with the recipient’s GPG public key fingerprint.
gpg --encrypt --recipient 5ADAE85E0C301567 register.sh
