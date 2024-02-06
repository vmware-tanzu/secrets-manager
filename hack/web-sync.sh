#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

cd docs || exit

JEKYLL_ENV=production bundle exec jekyll build

if [[ -z "$VSECM_S3_BUCKET" || -z "$VSECM_DISTRIBUTION_ID" ]]; then
  echo "Error: VSECM_S3_BUCKET and VSECM_DISTRIBUTION_ID must be set."
  exit 1
fi

rm -rf _site/versions
aws s3 sync _site/ "$VSECM_S3_BUCKET"

aws cloudfront create-invalidation --distribution-id "$VSECM_DISTRIBUTION_ID" --paths "/*"
