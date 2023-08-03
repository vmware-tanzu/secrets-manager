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

cd docs || exit

JEKYLL_ENV=production jekyll build

echo $VSECM_S3_BUCKET
echo $VSECM_DISTRIBUTION_ID

rm -rf _site/versions
aws s3 sync _site/ "$VSECM_S3_BUCKET"

aws cloudfront create-invalidation --distribution-id "$VSECM_DISTRIBUTION_ID" --paths "/*"
