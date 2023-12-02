# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# The common version tag is assigned to all the things.
ifdef VSECM_VERSION
	VERSION := $(VSECM_VERSION)
else
	VERSION := 0.21.4
endif

IMAGE=distroless
DEPLOYMENT_NAME=vsecm
VSECM_LOCAL_REGISTRY_URL ?= "localhost:5000"

# Utils
include ./makefiles/VSecMMacOs.mk
include ./makefiles/VSecMDeploy.mk

## Keygen
include ./makefiles/VSecMKeyGen.mk

## VMware Secrets Manager
include ./makefiles/VSecMSafe.mk
include ./makefiles/VSecMSentinel.mk
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
include makefiles/helmUtils.mk
