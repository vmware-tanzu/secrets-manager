#!/usr/bin/env bash

cd "$HOME"/WORKSPACE/secrets-manager || exit
export VSECM_LOCAL_REGISTRY_URL=localhost:32000
make build-local
