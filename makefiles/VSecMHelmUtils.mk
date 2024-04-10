# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# variables used with helm commands to set global values
export LOCAL_REGISTRY := --set global.registry=localhost:5000
export EKS_REGISTRY := --set global.registry=public.ecr.aws/h8y1n7y7
export DISTROLESSS_IMAGE := --set global.baseImage=distroless
export DISTROLESSS_FIPS_IMAGE := --set global.baseImage=distroless-fips
export HELM_CHART_PATH := "./helm-charts/${VERSION}"
export NAME_TEMPLATE := --name-template vsecm
export DEPLOY_SAFE_FALSE := --set global.deploySafe=false
export DEPLOY_SENTINEL_FALSE := --set global.deploySentinel=false
export DEPLOY_SPIRE_FALSE := --set global.deploySpire=false
export DEPLOY_KEYSTONE_FALSE := --set global.deployKeystone=false

# Render helm chart and save as kubernetes manifests
k8s-manifests-update:
	@if [ ! -d ${HELM_CHART_PATH} ]; then\
		echo "helm-chart does not exists for version(${VERSION})";\
		echo "Make sure helm-charts exists at ${HELM_CHART_PATH}/";\
		exit 1;\
	fi
	@echo "**************************************************************"
	@echo "Producing k8s manifest files for helm-chart version ${VERSION}, to change version pass VERSION variable with another value."
	@echo "**************************************************************"
	@echo "Ex. make k8s-manifests-update VERSION=0.22.4"
	./hack/update-k8s-manifests.sh ${VERSION}

# add an echo statement to publish to user default version is being installed
helm-install:
	@$(call validate_parameters, ${IMAGE})
	@echo "**************************************************************"
	@echo "Deploying VSecM using helm-chart version \"${VERSION}\", \"${IMAGE}\" image and with helm deployment \
	name ${DEPLOYMENT_NAME}."
	@echo "To change version, image, deployment name pass VERSION, IMAGE, DEPLOYMENT_NAME variable with another value."
	@echo "**************************************************************"
	@echo "Ex. make helm-install VERSION=0.22.4 IMAGE=distroless-fips DEPLOYMENT_NAME=vsecm"
	make helm-install-${IMAGE}

helm-install-distroless:
	helm install vsecm ${HELM_CHART_PATH} ${DISTROLESSS_IMAGE}

helm-install-distroless-fips:
	helm install vsecm ${HELM_CHART_PATH} ${DISTROLESSS_FIPS_IMAGE}

# Deletes the vsecm helm chart, while taking care of the SPIFFE CSI driver related
# resource deletion prioritization issues.
# Simply executing `helm uninstall` can result in dangling resources, which can
# cause the helm chart to not be deleted, and installation of a new chart to fail.
# This make target ensures that the SPIFFE CSI driver's dependent resources are
# deleted before the helm chart is deleted, hence avoiding the above issue.
helm-uninstall:
	@echo "**************************************************************"
	@echo "Uninstalling VSecM for helm deployment name \"${DEPLOYMENT_NAME}\", to change helm"
	@echo "deployment name value pass DEPLOYMENT_NAME variable with another value"
	@echo "**************************************************************"
	@echo "Ex. make helm-uninstall DEPLOYMENT_NAME=vsecm"
	make clean
	helm uninstall ${DEPLOYMENT_NAME}

# make target to release helm-chart
# usage: make helm-chart-release VSECM_VERSION=0.22.4
helm-chart-release:
	./hack/release-helm-chart.sh ${VERSION}

define validate_parameters
	@if [ ${IMAGE} != "distroless" ] && [ ${IMAGE} != "distroless-fips" ]; then\
		echo "Invalid IMAGE, valid options for IMAGE are: [distroless, distroless-fips] ";\
		exit 1;\
	fi
	@if [ ! -d ${HELM_CHART_PATH} ]; then\
		echo "helm-chart does not exists for version(${VERSION})";\
		echo "Make sure helm-charts exists at ${HELM_CHART_PATH}/";\
		exit 1;\
	fi
endef
