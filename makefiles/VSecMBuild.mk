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
	./hack/serve-docs.sh

# Builds everything and pushes to public DockerHub registries.
build: \
	inspector-bundle \
	inspector-push \
	keygen-bundle \
	keygen-push \
	example-sidecar-bundle \
	example-sidecar-push \
	example-sdk-bundle \
	example-sdk-push \
	example-multiple-secrets-bundle \
	example-multiple-secrets-push \
	example-init-container-bundle \
	example-init-container-push \
	keystone-bundle-ist \
	keystone-push-ist \
	keystone-bundle-ist-fips \
	keystone-push-ist-fips \
	safe-bundle-ist \
	safe-push-ist \
	safe-bundle-ist-fips \
	safe-push-ist-fips \
	sidecar-bundle-ist \
	sidecar-push-ist \
	sidecar-bundle-ist-fips \
	sidecar-push-ist-fips \
	sentinel-bundle-ist \
	sentinel-push-ist \
	sentinel-bundle-ist-fips \
	sentinel-push-ist-fips \
	init-container-bundle-ist \
	init-container-push-ist \
	init-container-bundle-ist-fips \
	init-container-push-ist-fips

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


# Builds everything and pushes to the local registry.
build-local: \
	inspector-bundle \
	inspector-push-local \
	keygen-bundle \
	keygen-push-local \
	example-sidecar-bundle \
	example-sidecar-push-local \
	example-sdk-bundle \
	example-sdk-push-local \
	example-multiple-secrets-bundle \
	example-multiple-secrets-push-local \
	example-init-container-bundle \
	example-init-container-push-local \
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
