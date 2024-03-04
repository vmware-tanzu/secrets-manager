/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

export GIT_SSH_COMMAND="ssh -i /home/vsecm/.ssh/id_ed25519"

cd /home/aegis/WORKSPACE/VSecM/ || exit

git stash
git checkout main
git pull
