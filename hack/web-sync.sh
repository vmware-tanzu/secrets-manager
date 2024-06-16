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

rm -rf public

container_id=$(docker ps --filter "ancestor=aegis-web:latest" --format "{{.ID}}")

if [ -n "$container_id" ]; then
  docker cp "$container_id":/public "$(pwd)"

  echo "Files copied from container ID $container_id to $(pwd)/public"
else
  echo "No running container found for image aegis-web:latest"
  exit 1
fi

if [[ -z "$VSECM_S3_BUCKET" || -z "$VSECM_DISTRIBUTION_ID" ]]; then
  echo "Error: VSECM_S3_BUCKET and VSECM_DISTRIBUTION_ID must be set."
  exit 1
fi

aws s3 sync public/ "$VSECM_S3_BUCKET"

aws cloudfront create-invalidation --distribution-id "$VSECM_DISTRIBUTION_ID" --paths "/*"
