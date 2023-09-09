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

# This is a script to fix unsigned images.
# Normally, signing and pushing should be a single step
# and we should not need to pull the images and sign them again.
# So we’d rarely (if ever) need to use this script.

VERSION="0.21.0"

export DOCKER_CONTENT_TRUST=0

docker pull vsecm/vsecm-ist-safe:"$VERSION"
docker pull vsecm/vsecm-ist-safe:latest
docker pull vsecm/vsecm-ist-sentinel:"$VERSION"
docker pull vsecm/vsecm-ist-sentinel:latest
docker pull vsecm/vsecm-ist-sidecar:"$VERSION"
docker pull vsecm/vsecm-ist-sidecar:latest
docker pull vsecm/vsecm-ist-init-container:"$VERSION"
docker pull vsecm/vsecm-ist-init-container:latest
docker pull vsecm/example-using-sidecar:"$VERSION"
docker pull vsecm/example-using-sidecar:latest
docker pull vsecm/example-using-sdk:"$VERSION"
docker pull vsecm/example-using-sdk:latest
docker pull vsecm/example-multiple-secrets:"$VERSION"
docker pull vsecm/example-multiple-secrets:latest
docker pull vsecm/example-using-init-container:"$VERSION"
docker pull vsecm/example-using-init-container:latest

docker pull vsecm/vsecm-photon-safe:"$VERSION"
docker pull vsecm/vsecm-photon-safe:latest
docker pull vsecm/vsecm-photon-sentinel:"$VERSION"
docker pull vsecm/vsecm-photon-sentinel:latest
docker pull vsecm/vsecm-photon-sidecar:"$VERSION"
docker pull vsecm/vsecm-photon-sidecar:latest
docker pull vsecm/vsecm-photon-init-container:"$VERSION"
docker pull vsecm/vsecm-photon-init-container:latest

export DOCKER_CONTENT_TRUST=1

docker trust sign vsecm/vsecm-ist-safe:"$VERSION"
docker trust sign vsecm/vsecm-ist-safe:latest
docker trust sign vsecm/vsecm-ist-sentinel:"$VERSION"
docker trust sign vsecm/vsecm-ist-sentinel:latest
docker trust sign vsecm/vsecm-ist-sidecar:"$VERSION"
docker trust sign vsecm/vsecm-ist-sidecar:latest
docker trust sign vsecm/vsecm-ist-init-container:"$VERSION"
docker trust sign vsecm/vsecm-ist-init-container:latest
docker trust sign vsecm/example-using-sidecar:"$VERSION"
docker trust sign vsecm/example-using-sidecar:latest
docker trust sign vsecm/example-using-sdk:"$VERSION"
docker trust sign vsecm/example-using-sdk:latest
docker trust sign vsecm/example-multiple-secrets:"$VERSION"
docker trust sign vsecm/example-multiple-secrets:latest
docker trust sign vsecm/example-using-init-container:"$VERSION"
docker trust sign vsecm/example-using-init-container:latest

docker trust sign vsecm/vsecm-photon-safe:"$VERSION"
docker trust sign vsecm/vsecm-photon-safe:latest
docker trust sign vsecm/vsecm-photon-sentinel:"$VERSION"
docker trust sign vsecm/vsecm-photon-sentinel:latest
docker trust sign vsecm/vsecm-photon-sidecar:"$VERSION"
docker trust sign vsecm/vsecm-photon-sidecar:latest
docker trust sign vsecm/vsecm-photon-init-container:"$VERSION"
docker trust sign vsecm/vsecm-photon-init-container:latest

