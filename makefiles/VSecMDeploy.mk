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
# ## Versioning ##
#

# tags a release
tag:
	./hack/tag.sh $(VERSION)
