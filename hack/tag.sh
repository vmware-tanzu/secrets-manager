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

VERSION="$1"

echo ""
echo "--------"
echo "VSecM"
if git tag -s v"$VERSION"; then
  git push origin --tags
  gh release create
fi

echo ""
echo "--------"
echo "vsecm-safe"
docker trust sign vsecm/vsecm-ist-safe:"$VERSION"
docker trust sign vsecm/vsecm-ist-safe:latest
echo "vsecm-sentinel"
docker trust sign vsecm/vsecm-ist-sentinel:"$VERSION"
docker trust sign vsecm/vsecm-ist-sentinel:latest
echo "vsecm-sidecar"
docker trust sign vsecm/vsecm-ist-sidecar:"$VERSION"
docker trust sign vsecm/vsecm-ist-sidecar:latest
echo "vsecm-init-container"
docker trust sign vsecm/vsecm-ist-init-container:"$VERSION"
docker trust sign vsecm/vsecm-ist-init-container:latest
echo "example-using-sidecar"
docker trust sign vsecm/example-using_sidecar:"$VERSION"
docker trust sign vsecm/example-using_sidecar:latest
echo "example-using-sdk"
docker trust sign vsecm/example-using_sdk_go:"$VERSION"
docker trust sign vsecm/example-using_sdk_go:latest
echo "example-multiple-secrets"
docker trust sign vsecm/example-multiple_secrets:"$VERSION"
docker trust sign vsecm/example-multiple_secrets:latest
echo "example-using-init-container"
docker trust sign vsecm/example-using_init_container:"$VERSION"
docker trust sign vsecm/example-using_init_container:latest

echo "vsecm-photon-safe"
docker trust sign vsecm/vsecm-photon-safe:"$VERSION"
docker trust sign vsecm/vsecm-photon-safe:latest
echo "vsecm-photon-sentinel"
docker trust sign vsecm/vsecm-photon-sentinel:"$VERSION"
docker trust sign vsecm/vsecm-photon-sentinel:latest
echo "vsecm-photon-sidecar"
docker trust sign vsecm/vsecm-photon-sidecar:"$VERSION"
docker trust sign vsecm/vsecm-photon-sidecar:latest
echo "vsecm-photon-init-container"
docker trust sign vsecm/vsecm-photon-init-container:"$VERSION"
docker trust sign vsecm/vsecm-photon-init-container:latest

echo "Everything is awesome!"
