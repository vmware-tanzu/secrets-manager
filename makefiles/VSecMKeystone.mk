# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the "VSecM Keystone" into a container image.
keystone-bundle-ist:
	./hack/bundle.sh "vsecm-ist-keystone" \
		$(VERSION) "dockerfiles/vsecm-ist/keystone.Dockerfile"

# Packages the "VSecM Keystone" into a container image for FIPS.
keystone-bundle-ist-fips:
	./hack/bundle.sh "vsecm-ist-fips-keystone" \
		$(VERSION) "dockerfiles/vsecm-ist-fips/keystone.Dockerfile"

# Pushes the "VSecM Keystone" container to the public registry.
keystone-push-ist:
	./hack/push.sh "vsecm-ist-keystone" $(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-ist-keystone"

# Pushes the "VSecM Keystone" container to the public EKS registry.
keystone-push-ist-eks:
	./hack/push.sh "vsecm-ist-keystone" $(VERSION) "$(VSECM_EKS_REGISTRY_URL)/vsecm-ist-keystone"

# Pushes the "VSecM Keystone" (FIPS) container to the public registry.
keystone-push-ist-fips:
	./hack/push.sh "vsecm-ist-fips-keystone" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-ist-fips-keystone"

# Pushes the "VSecM Keystone" (FIPS) container to the public EKS registry.
keystone-push-ist-fips-eks:
	./hack/push.sh "vsecm-ist-fips-keystone" \
		$(VERSION) "$(VSECM_EKS_REGISTRY_URL)/vsecm-ist-fips-keystone"

# Pushes the "VSecM Safe" container image to the local registry.
keystone-push-ist-local:
	./hack/push.sh "vsecm-ist-keystone" $(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-keystone"

# Pushes the "VSecM Safe" container image to the local registry.
keystone-push-ist-fips-local:
	./hack/push.sh "vsecm-ist-fips-keystone" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-fips-keystone"
