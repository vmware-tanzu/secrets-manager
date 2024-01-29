# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

#
# ## Lifecycle ##
#

MANIFESTS_BASE_PATH="./k8s/${VERSION}"
MANIFESTS_LOCAL_PATH="${MANIFESTS_BASE_PATH}/local"
MANIFESTS_REMOTE_PATH="${MANIFESTS_BASE_PATH}/remote"

# Removes the former VSecM deployment without entirely destroying the cluster.
clean:
	./hack/uninstall.sh

# Completely removes the Minikube cluster.
k8s-delete:
	./hack/minikube-delete.sh
# Brings up a fresh Minikube cluster.
k8s-start:
	./hack/minikube-start.sh

deploy-spire:
	@if [ "${DEPLOY_SPIRE}" = "true" ]; then \
		kubectl apply -f ${MANIFESTS_BASE_PATH}/crds; \
		kubectl apply -f ${MANIFESTS_BASE_PATH}/spire.yaml; \
		echo "verifying spire installation"; \
		kubectl wait --for=condition=Available deployment -n spire-system spire-server; \
		echo "spire-server: deployment available"; \
		echo "spire installation successful"; \
	fi

# Deploys VSecM to the cluster.
deploy: deploy-spire
	kubectl apply -f ${MANIFESTS_REMOTE_PATH}/vsecm-distroless.yaml
	$(MAKE) post-deploy
deploy-fips: deploy-spire
	kubectl apply -f ${MANIFESTS_REMOTE_PATH}/vsecm-distroless-fips.yaml
	$(MAKE) post-deploy
deploy-photon: deploy-spire
	kubectl apply -f ${MANIFESTS_REMOTE_PATH}/vsecm-photon.yaml
	$(MAKE) post-deploy
deploy-photon-fips: deploy-spire
	kubectl apply -f ${MANIFESTS_REMOTE_PATH}/vsecm-photon-fips.yaml
	$(MAKE) post-deploy
deploy-local: deploy-spire
	kubectl apply -f ${MANIFESTS_LOCAL_PATH}/vsecm-distroless.yaml
	$(MAKE) post-deploy
deploy-fips-local: deploy-spire
	kubectl apply -f ${MANIFESTS_LOCAL_PATH}/vsecm-distroless-fips.yaml
	$(MAKE) post-deploy
deploy-photon-local: deploy-spire
	kubectl apply -f ${MANIFESTS_LOCAL_PATH}/vsecm-photon.yaml
	$(MAKE) post-deploy
deploy-photon-fips-local: deploy-spire
	kubectl apply -f ${MANIFESTS_LOCAL_PATH}/vsecm-photon-fips.yaml
	$(MAKE) post-deploy
.SILENT:
.PHONY: post-deploy
post-deploy:
	echo "verifying vsecm installation"
	kubectl wait --for=condition=Available deployment -n vsecm-system vsecm-sentinel
	echo "vsecm-sentinel: deployment available"
	kubectl wait --for=condition=Available deployment -n vsecm-system vsecm-safe
	echo "vsecm-safe: deployment available"
	echo "vsecm installation successful"

#
# ## Tests ##
#

# Integration tests.
test:
	./hack/test.sh "remote" ""
test-remote:
	./hack/test.sh "remote" ""
test-local:
	./hack/test.sh "local" ""
test-eks:
	$(eval VSECM_EKS_CONTEXT=$(shell kubectl config get-contexts -o name | grep "arn:aws:eks"))
	@if [ -z "$(VSECM_EKS_CONTEXT)" ]; then \
	echo "Error: No EKS context found."; \
		exit 1; \
	fi
	@echo "Using EKS context: $VSECM_EKS_CONTEXT)"
	kubectl config use-context $(VSECM_EKS_CONTEXT)

	$(eval VSECM_EKS_VERSION=$(shell helm search repo vsecm/vsecm -o json | jq -r '.[0].version'))
	@if [ -z "$(VSECM_EKS_VERSION)" ]; then \
		echo "Error: Unable to determine VSECM_EKS_VERSION."; \
		exit 1; \
	fi
	@echo "Using VERSION: $$VSECM_EKS_VERSION"

	./hack/install-eks.sh

	(VERSION=$$VSECM_EKS_VERSION; ./hack/test.sh "remote" "eks")
	kubectl config use-context minikube
test-local-ci:
	./hack/test.sh "local" "ci"

#
# ## Versioning ##
#

# tags a release
tag:
	./hack/tag.sh $(VERSION)
