# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages "VSecM Scout" into a container image.
scout-bundle-ist:
	./hack/bundle.sh "vsecm-ist-scout" \
		$(VERSION) "dockerfiles/vsecm-ist/scout.Dockerfile"

# Packages "VSecM Scout" into a container image for FIPS.
scout-bundle-ist-fips:
	./hack/bundle.sh "vsecm-ist-fips-scout" \
		$(VERSION) "dockerfiles/vsecm-ist-fips/scout.Dockerfile"

# Pushes "VSecM Scout" container to the public registry.
scout-push-ist:
	./hack/push.sh "vsecm-ist-scout" $(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-ist-scout"

# Pushes "VSecM Scout" (FIPS) container to the public registry.
scout-push-ist-fips:
	./hack/push.sh "vsecm-ist-fips-scout" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-ist-fips-scout"

# Pushes the "VSecM Scout" container image to the local registry.
scout-push-ist-local:
	./hack/push.sh "vsecm-ist-scout" $(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-scout"

# Pushes the "VSecM Scout" (FIPS) container image to the local registry.
scout-push-ist-fips-local:
	./hack/push.sh "vsecm-ist-fips-scout" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-fips-scout"
