#!/usr/bin/env bash

docker run --rm \
  -v "$(pwd)":/vsecm \
  -e VSECM_KEYGEN_EXPORTED_SECRET_PATH="/vsecm/secrets.json" \
  -e VSECM_KEYGEN_ROOT_KEY_PATH="/vsecm/key.txt" \
  -e VSECM_KEYGEN_DECRYPT="true" \
  vsecm-keygen:0.22.3
