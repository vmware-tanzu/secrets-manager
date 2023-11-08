# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# variables used with helm commands to set global values
export LOCAL_REGISTRY := --set global.registry=localhost:5000
export DISTROLESSS_IMAGE := --set global.baseImage=distroless
export DISTROLESSS_FIPS_IMAGE := --set global.baseImage=distroless-fips
export PHOTON_IMAGE := --set global.baseImage=photon
export PHOTON_FIPS_IMAGE := --set global.baseImage=photos-fips
export HELM_CHART_PATH := "./helm-charts/${VERSION}"
export NAME_TEMPLATE := --name-template vsecm

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
	@echo "Ex. make k8s-manifests-update VERSION=0.22.0"
	./hack/update-k8s-manifests.sh ${VERSION}

# add an echo statement to publish to user default version is being installed
helm-install:
	@$(call validate_parameters, ${IMAGE})
	@echo "**************************************************************"
	@echo "Deploying VSecM using helm-chart version \"${VERSION}\", \"${IMAGE}\" image and with helm deployment \
	name ${DEPLOYMENT_NAME}."
	@echo "To change version, image, deployment name pass VERSION, IMAGE, DEPLOYMENT_NAME variable with another value."
	@echo "**************************************************************"
	@echo "Ex. make helm-install VERSION=0.22.0 IMAGE=distroless-fips DEPLOYMENT_NAME=vsecm"
	make helm-install-${IMAGE}

helm-install-distroless:
	helm install vsecm ${HELM_CHART_PATH} ${DISTROLESSS_IMAGE}

helm-install-distroless-fips:
	helm install vsecm ${HELM_CHART_PATH} ${DISTROLESSS_FIPS_IMAGE}

helm-install-photon:
	helm install vsecm ${HELM_CHART_PATH} ${PHOTON_IMAGE}

helm-install-photos-fips:
	helm install vsecm ${HELM_CHART_PATH} ${PHOTON_FIPS_IMAGE}

helm-uninstall:
	@echo "**************************************************************"
	@echo "Uninstalling VSecM for helm deployment name \"${DEPLOYMENT_NAME}\", to change helm"
	@echo "deployment name value pass DEPLOYMENT_NAME variable with another value"
	@echo "**************************************************************"
	@echo "Ex. make helm-uninstall DEPLOYMENT_NAME=vsecm"
	helm uninstall ${DEPLOYMENT_NAME}

# make target to release helm-chart
# usage: make helm-chart-release VSECM_VERSION=0.21.1
helm-chart-release:
	./hack/release-helm-chart.sh ${VERSION}

define validate_parameters
	@if [ ${IMAGE} != "distroless" ] && [ ${IMAGE} != "distroless-fips" ] && [ ${IMAGE} != "photon" ] && [ ${IMAGE} != "photon-fips" ]; then\
		echo "Invalid IMAGE, valid options for IMAGE are: [distroless, distroless-fips, photon, photon-fips] ";\
		exit 1;\
	fi
	@if [ ! -d ${HELM_CHART_PATH} ]; then\
		echo "helm-chart does not exists for version(${VERSION})";\
		echo "Make sure helm-charts exists at ${HELM_CHART_PATH}/";\
		exit 1;\
	fi
endef
