# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# The common version tag is assigned to all the things.
ifdef VSECM_VERSION
	VERSION := $(VSECM_VERSION)
else
	VERSION := 0.27.2
endif

# Set deploySpire to false, if you want to use existing spire deployment
ifdef deploySpire
	DEPLOY_SPIRE := $(deploySpire)
else
	DEPLOY_SPIRE := "true"
endif

IMAGE=distroless
DEPLOYMENT_NAME=vsecm
VSECM_DOCKERHUB_REGISTRY_URL ?= "vsecm"
VSECM_LOCAL_REGISTRY_URL ?= "localhost:5000"
VSECM_EKS_REGISTRY_URL ?= "public.ecr.aws/h8y1n7y7"

VSECM_NAMESPACE_SYSTEM ?= "vsecm-system"
VSECM_NAMESPACE_SPIRE ?= "spire-system"
VSECM_NAMESPACE_SPIRE_SERVER ?= "spire-server"
# VSECM_NAMESPACE_SYSTEM ?= "vsecm-system-custom"
# VSECM_NAMESPACE_SPIRE ?= "spire-system-custom"
# VSECM_NAMESPACE_SPIRE_SERVER ?= "spire-server-custom"

# Utils
include ./makefiles/VSecMMacOs.mk
include ./makefiles/VSecMDeploy.mk

## Inspector
include ./makefiles/VSecMInspector.mk

## Keygen
include ./makefiles/VSecMKeyGen.mk

## VMware Secrets Manager
include ./makefiles/VSecMSafe.mk
include ./makefiles/VSecMSentinel.mk
include ./makefiles/VSecMKeystone.mk
include ./makefiles/VSecMInitContainer.mk
include ./makefiles/VSecMSidecar.mk

## Examples
include ./makefiles/VSecMExampleSidecar.mk
include ./makefiles/VSecMExampleSdk.mk
include ./makefiles/VSecMExampleMultipleSecrets.mk
include ./makefiles/VSecMExampleInitContainer.mk

## Build
include ./makefiles/VSecMBuild.mk

## Help
include ./makefiles/VSecMHelp.mk

## Helm-chart, k8s-manifests utils
include makefiles/VSecMHelmUtils.mk

## Git Helper
include makefiles/Git.mk

## Coverage
include makefiles/Test.mk

## Generate Proto Files
include makefiles/VSecMGenerateProtoFiles.mk
