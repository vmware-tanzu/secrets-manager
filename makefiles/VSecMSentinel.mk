# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the "VSecM Sentinel" binary into a container image.
sentinel-bundle-ist:
	./hack/bundle.sh "vsecm-ist-sentinel" \
		$(VERSION) "dockerfiles/vsecm-ist/sentinel.Dockerfile"

# Packages the "VSecM Sentinel" binary into a container image for FIPS.
sentinel-bundle-ist-fips:
	./hack/bundle.sh "vsecm-ist-fips-sentinel" \
		$(VERSION) "dockerfiles/vsecm-ist-fips/sentinel.Dockerfile"

# Pushes the "VSecM Sentinel" container image the the public registry.
sentinel-push-ist:
	./hack/push.sh "vsecm-ist-sentinel" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-ist-sentinel"

# Pushes the "VSecM Sentinel" container image to the public EKS registry.
sentinel-push-ist-eks:
	./hack/push.sh "vsecm-ist-sentinel" \
		$(VERSION) "$(VSECM_EKS_REGISTRY_URL)/vsecm-ist-sentinel"

# Pushes the "VSecM Sentinel" (FIPS) container image to the public registry.
sentinel-push-ist-fips:
	./hack/push.sh "vsecm-ist-fips-sentinel" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-ist-fips-sentinel"

# Pushes the "VSecM Sentinel" (FIPS) container image to the public EKS registry.
sentinel-push-ist-fips-eks:
	./hack/push.sh "vsecm-ist-fips-sentinel" \
		$(VERSION) "$(VSECM_EKS_REGISTRY_URL)/vsecm-ist-fips-sentinel"

# Pushes the "VSecM Sentinel" container image to the local registry.
sentinel-push-ist-local:
	./hack/push.sh "vsecm-ist-sentinel" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-sentinel"

# Pushes the "VSecM Sentinel" (FIPS) container image to the local registry.
sentinel-push-ist-fips-local:
	./hack/push.sh "vsecm-ist-fips-sentinel" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-fips-sentinel"
