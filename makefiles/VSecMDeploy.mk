# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

#
# ## Lifecycle ##
#

MANIFESTS_BASE_PATH="./k8s/${VERSION}"
MANIFESTS_LOCAL_PATH="${MANIFESTS_BASE_PATH}/local"
MANIFESTS_EKS_PATH="${MANIFESTS_BASE_PATH}/eks"
MANIFESTS_REMOTE_PATH="${MANIFESTS_BASE_PATH}/remote"
CPU ?= $(or $(VSECM_MINIKUBE_CPU_COUNT),2)
NODES ?= $(or $(VSECM_MINIKUBE_NODE_COUNT),1)
MEMORY ?= $(or $(VSECM_MINIKUBE_MEMORY),4096)

# Removes the former VSecM deployment without entirely destroying the cluster.
clean:
	./hack/uninstall.sh $(VSECM_NAMESPACE_SYSTEM) $(VSECM_NAMESPACE_SPIRE) $(VSECM_NAMESPACE_SPIRE_SERVER)

# Completely removes the Minikube cluster.
k8s-delete:
	./hack/minikube-delete.sh

# Brings up a fresh Minikube cluster.
k8s-start:
	@NODES=$(NODES) CPU=$(CPU) MEMORY=$(MEMORY) ./hack/minikube-start.sh

deploy-spire-crds:
	kubectl apply -f ${MANIFESTS_BASE_PATH}/crds

deploy-spire:
	@if [ "${DEPLOY_SPIRE}" = "true" ]; then \
		kubectl apply -f ${MANIFESTS_BASE_PATH}/crds; \
		kubectl apply -f ${MANIFESTS_BASE_PATH}/spire.yaml; \
		echo "verifying SPIRE installation"; \
		kubectl wait --for=condition=ready pod spire-server-0 --timeout=120s -n $(VSECM_NAMESPACE_SPIRE_SERVER); \
		echo "spire-server: deployment available"; \
		echo "spire installation successful"; \
		echo "sleeping for 15 seconds for webhooks to become responsive"; \
		sleep 15; \
	fi

# Deploys VSecM to the cluster.
deploy: deploy-spire
	kubectl apply -f ${MANIFESTS_REMOTE_PATH}/vsecm-distroless.yaml || { \
		echo "Command failed, retrying in 15 seconds..." \
		sleep 15; \
		kubectl apply -f ${MANIFESTS_REMOTE_PATH}/vsecm-distroless.yaml; \
	}
	$(MAKE) post-deploy
deploy-fips: deploy-spire
	kubectl apply -f ${MANIFESTS_REMOTE_PATH}/vsecm-distroless-fips.yaml || { \
		echo "Command failed, retrying in 15 seconds..." \
		sleep 15; \
		kubectl apply -f ${MANIFESTS_REMOTE_PATH}/vsecm-distroless-fips.yaml; \
	}
	$(MAKE) post-deploy
deploy-local: deploy-spire
	kubectl apply -f ${MANIFESTS_LOCAL_PATH}/vsecm-distroless.yaml || { \
		echo "Command failed, retrying in 15 seconds..." \
		sleep 15; \
		kubectl apply -f ${MANIFESTS_LOCAL_PATH}/vsecm-distroless.yaml; \
	}
	$(MAKE) post-deploy
deploy-fips-local: deploy-spire
	kubectl apply -f ${MANIFESTS_LOCAL_PATH}/vsecm-distroless-fips.yaml || { \
		echo "Command failed, retrying in 15 seconds..." \
		sleep 15; \
		kubectl apply -f ${MANIFESTS_LOCAL_PATH}/vsecm-distroless-fips.yaml; \
	}

	$(MAKE) post-deploy
deploy-eks: deploy-spire
	kubectl apply -f ${MANIFESTS_EKS_PATH}/vsecm-distroless.yaml || { \
		echo "Command failed, retrying in 15 seconds..." \
		sleep 15; \
		kubectl apply -f ${MANIFESTS_EKS_PATH}/vsecm-distroless.yaml; \
	}
	$(MAKE) post-deploy
deploy-fips-eks: deploy-spire
	kubectl apply -f ${MANIFESTS_EKS_PATH}/vsecm-distroless-fips.yaml || { \
		echo "Command failed, retrying in 15 seconds..." \
		sleep 15; \
		kubectl apply -f ${MANIFESTS_EKS_PATH}/vsecm-distroless-fips.yaml; \
	}
	$(MAKE) post-deploy

.SILENT:
.PHONY: post-deploy
post-deploy:
	echo "verifying vsecm installation"
	kubectl wait --timeout=120s --for=condition=Available deployment -n $(VSECM_NAMESPACE_SYSTEM) vsecm-sentinel
	echo "vsecm-sentinel: deployment available"
	kubectl wait --for=condition=ready pod vsecm-safe-0 --timeout=120s -n $(VSECM_NAMESPACE_SYSTEM)
	echo "vsecm-safe: deployment available"
	echo "vsecm installation successful"

#
# ## Versioning ##
#

# tags a release
tag:
	./hack/tag.sh $(VERSION)
