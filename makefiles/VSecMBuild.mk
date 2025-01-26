# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

serve-docs:
	./hack/web-serve.sh

sync-docs:
	./hack/web-sync.sh

bundle-all-local: \
	inspector-bundle \
	keygen-bundle \
	example-sdk-bundle \
	example-sidecar-bundle \
	example-multiple-secrets-bundle \
	example-init-container-bundle \
	keystone-bundle-ist \
	safe-bundle-ist \
	scout-bundle-ist \
	sidecar-bundle-ist \
	sentinel-bundle-ist \
	init-container-bundle-ist
	# relay-client-bundle-ist \
	# relay-server-bundle-ist

bundle-all: \
	inspector-bundle \
	keygen-bundle \
	scout-bundle-ist \
	example-sdk-bundle \
	example-sidecar-bundle \
	example-multiple-secrets-bundle \
	example-init-container-bundle \
	keystone-bundle-ist \
	keystone-bundle-ist-fips \
	safe-bundle-ist \
	safe-bundle-ist-fips \
	scout-bundle-ist \
	scout-bundle-ist-fips \
	sidecar-bundle-ist \
	sidecar-bundle-ist-fips \
	sentinel-bundle-ist \
	sentinel-bundle-ist-fips \
	init-container-bundle-ist \
	init-container-bundle-ist-fips
#	relay-client-bundle-ist \
#	relay-client-bundle-ist-fips \
#	relay-server-bundle-ist \
#	relay-server-bundle-ist-fips

push-all: \
	inspector-push \
	keygen-push \
	example-sidecar-push \
	example-sdk-push \
	example-multiple-secrets-push \
	example-init-container-push \
	keystone-push-ist \
	keystone-push-ist-fips \
	safe-push-ist \
	safe-push-ist-fips \
	scout-push-ist \
	scout-push-ist-fips \
	sidecar-push-ist \
	sidecar-push-ist-fips \
	sentinel-push-ist \
	sentinel-push-ist-fips \
	init-container-push-ist \
	init-container-push-ist-fips
#	relay-client-push-ist \
#	relay-client-push-ist-fips \
#	rely-server-push-ist \
#	relay-server-push-ist-fips

# Builds everything and pushes to public DockerHub registries.
build:
	./hack/generate-proto-files.sh
	@$(MAKE) -j$(CPU) bundle-all
	@$(MAKE) -j$(CPU) push-all

push-all-local: \
	inspector-push-local \
	keygen-push-local \
	example-sidecar-push-local \
	example-sdk-push-local \
	example-multiple-secrets-push-local \
	example-init-container-push-local \
	keystone-push-ist-local \
	safe-push-ist-local \
	scout-push-ist-local \
	sidecar-push-ist-local \
	sentinel-push-ist-local \
	init-container-push-ist-local
#	relay-client-push-ist-local \
#	relay-server-push-ist-local

# Builds everything and pushes to the local registry.
build-local:
	./hack/generate-proto-files.sh
	@$(MAKE) -j$(CPU) bundle-all-local
	@$(MAKE) -j$(CPU) push-all-local

build-essentials-local: \
	keygen-bundle \
	inspector-bundle \
	keystone-bundle-ist \
	keystone-push-ist-local \
	safe-bundle-ist \
	safe-push-ist-local \
	sidecar-bundle-ist \
	sidecar-push-ist-local \
	sentinel-bundle-ist \
	sentinel-push-ist-local \
	init-container-bundle-ist \
	init-container-push-ist-local
