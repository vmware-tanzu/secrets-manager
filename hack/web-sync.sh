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

JEKYLL_ENV=production jekyll build

cd web || exit

rm -rf _site/versions
aws s3 sync _site/ s3://vsecm.com/

aws cloudfront create-invalidation --distribution-id EZFGMY32S3BBS --paths "/*"
