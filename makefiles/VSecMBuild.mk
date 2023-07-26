# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Builds everything and pushes to public registries.
build: \
	example-sidecar-bundle \
	example-sidecar-push \
	example-sdk-bundle \
	example-sdk-push \
	example-multiple-secrets-bundle \
	example-multiple-secrets-push \
	example-init-container-bundle \
	example-init-container-push \
	safe-bundle-ist \
	safe-push-ist \
	safe-bundle-ist-fips \
	safe-push-ist-fips \
	safe-bundle-photon \
	safe-push-photon \
	safe-bundle-photon-fips \
	safe-push-photon-fips \
	sidecar-bundle-ist \
	sidecar-push-ist \
	sidecar-bundle-ist-fips \
	sidecar-push-ist-fips \
	sidecar-bundle-photon \
	sidecar-push-photon \
	sidecar-bundle-photon-fips \
	sidecar-push-photon-fips \
	sentinel-bundle-ist \
	sentinel-push-ist \
	sentinel-bundle-ist-fips \
	sentinel-push-ist-fips \
	sentinel-bundle-photon \
	sentinel-push-photon \
	sentinel-bundle-photon-fips \
	sentinel-push-photon-fips \
	init-container-bundle-ist \
	init-container-push-ist \
	init-container-bundle-ist-fips \
	init-container-push-ist-fips \
	init-container-bundle-photon \
	init-container-push-photon \
	init-container-bundle-photon-fips \
	init-container-push-photon-fips

# Builds everything and pushes to the local registry.
build-local: \
	example-sidecar-bundle \
	example-sidecar-push-local \
	example-sdk-bundle \
	example-sdk-push-local \
	example-multiple-secrets-bundle \
	example-multiple-secrets-push-local \
	example-init-container-bundle \
	example-init-container-push-local \
	safe-bundle-ist \
	safe-push-ist-local \
	safe-bundle-ist-fips \
	safe-push-ist-fips-local \
	safe-bundle-photon \
	safe-push-photon-local \
	safe-bundle-photon-fips \
	safe-push-photon-fips-local \
	sidecar-bundle-ist \
	sidecar-push-ist-local \
	sidecar-bundle-ist-fips \
	sidecar-push-ist-fips-local \
	sidecar-bundle-photon \
	sidecar-push-photon-local \
	sidecar-bundle-photon-fips \
	sidecar-push-photon-fips-local \
	sentinel-bundle-ist \
	sentinel-push-ist-local \
	sentinel-bundle-ist-fips \
	sentinel-push-ist-fips-local \
	sentinel-bundle-photon \
	sentinel-push-photon-local \
	sentinel-bundle-photon-fips \
	sentinel-push-photon-fips-local \
	init-container-bundle-ist \
	init-container-push-ist-local \
	init-container-bundle-ist-fips \
	init-container-push-ist-fips-local \
	init-container-bundle-photon \
	init-container-push-photon-local \
	init-container-bundle-photon-fips \
	init-container-push-photon-fips-local

build-essentials-local: \
	safe-bundle-ist \
	safe-push-ist-local \
	sidecar-bundle-ist \
	sidecar-push-ist-local \
	sentinel-bundle-ist \
	sentinel-push-ist-local \
	init-container-bundle-ist \
	init-container-push-ist-local \
