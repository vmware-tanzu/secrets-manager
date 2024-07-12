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
	sidecar-bundle-ist \
	sentinel-bundle-ist \
	init-container-bundle-ist

bundle-all: \
	inspector-bundle \
	keygen-bundle \
	example-sdk-bundle \
	example-sidecar-bundle \
	example-multiple-secrets-bundle \
	example-init-container-bundle \
	keystone-bundle-ist \
	keystone-bundle-ist-fips \
	safe-bundle-ist \
	safe-bundle-ist-fips \
	sidecar-bundle-ist \
	sidecar-bundle-ist-fips \
	sentinel-bundle-ist \
	sentinel-bundle-ist-fips \
	init-container-bundle-ist \
	init-container-bundle-ist-fips

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
	sidecar-push-ist \
	sidecar-push-ist-fips \
	sentinel-push-ist \
	sentinel-push-ist-fips \
	init-container-push-ist \
	init-container-push-ist-fips

# Builds everything and pushes to public DockerHub registries.
build:
	./hack/generate-proto-files.sh
	@$(MAKE) -j$(CPU) bundle-all
	@$(MAKE) -j$(CPU) push-all

# Login to the public EKS registry.
login-eks:
	$(eval VSECM_EKS_CONTEXT=$(shell kubectl config get-contexts -o name | grep "arn:aws:eks"))
	@if [ -z "$(VSECM_EKS_CONTEXT)" ]; then \
		echo "Warning: login-eks: No EKS context found."; \
	else \
		echo "Using EKS context: $(VSECM_EKS_CONTEXT)"; \
		kubectl config use-context $(VSECM_EKS_CONTEXT); \
	fi

	aws ecr-public get-login-password --region us-east-1 | \
		docker login --username AWS --password-stdin public.ecr.aws/h8y1n7y7

# Builds everything and pushes to the public EKS registry.
build-eks: \
	login-eks \
	inspector-bundle \
	inspector-push-eks \
	keygen-bundle \
	keygen-push-eks \
	example-sidecar-bundle \
	example-sidecar-push-eks \
	example-sdk-bundle \
	example-sdk-push-eks \
	example-multiple-secrets-bundle \
	example-multiple-secrets-push-eks \
	example-init-container-bundle \
	example-init-container-push-eks \
	keystone-bundle-ist \
	keystone-push-ist-eks \
	keystone-bundle-ist-fips \
	keystone-push-ist-fips-eks \
	safe-bundle-ist \
	safe-push-ist-eks \
	safe-bundle-ist-fips \
	safe-push-ist-fips-eks \
	sidecar-bundle-ist \
	sidecar-push-ist-eks \
	sidecar-bundle-ist-fips \
	sidecar-push-ist-fips-eks \
	sentinel-bundle-ist \
	sentinel-push-ist-eks \
	sentinel-bundle-ist-fips \
	sentinel-push-ist-fips-eks \
	init-container-bundle-ist \
	init-container-push-ist-eks \
	init-container-bundle-ist-fips \
	init-container-push-ist-fips-eks

push-all-local: \
	inspector-push-local \
	keygen-push-local \
	example-sidecar-push-local \
	example-sdk-push-local \
	example-multiple-secrets-push-local \
	example-init-container-push-local \
	keystone-push-ist-local \
	safe-push-ist-local \
	sidecar-push-ist-local \
	sentinel-push-ist-local \
	init-container-push-ist-local


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
