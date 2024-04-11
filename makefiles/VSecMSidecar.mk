# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the "VSecM Sidecar" binary into a container image.
sidecar-bundle-ist:
	./hack/bundle.sh "vsecm-ist-sidecar" \
		$(VERSION) "dockerfiles/vsecm-ist/sidecar.Dockerfile"

# Packages the "VSecM Sidecar" binary into a container image for FIPS.
sidecar-bundle-ist-fips:
	./hack/bundle.sh "vsecm-ist-fips-sidecar" \
		$(VERSION) "dockerfiles/vsecm-ist-fips/sidecar.Dockerfile"

# Pushes the "VSecM Sidecar" container image to the public registry.
sidecar-push-ist:
	./hack/push.sh "vsecm-ist-sidecar" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-ist-sidecar"

# Pushes the "VSecM Sidecar" container image to the public EKS registry.
sidecar-push-ist-eks:
	./hack/push.sh "vsecm-ist-sidecar" \
		$(VERSION) "$(VSECM_EKS_REGISTRY_URL)/vsecm-ist-sidecar"

# Pushes the "VSecM Sidecar" (FIPS) container image to the public registry.
sidecar-push-ist-fips:
	./hack/push.sh "vsecm-ist-fips-sidecar" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-ist-fips-sidecar"

# Pushes the "VSecM Sidecar" (FIPS) container image to the public EKS registry.
sidecar-push-ist-fips-eks:
	./hack/push.sh "vsecm-ist-fips-sidecar" \
		$(VERSION) "$(VSECM_EKS_REGISTRY_URL)/vsecm-ist-fips-sidecar"

# Pushes the "VSecM Sidecar" container image to the local registry.
sidecar-push-ist-local:
	./hack/push.sh "vsecm-ist-sidecar" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-sidecar"

# Pushes the "VSecM Sidecar" (FIPS) container image to the local registry.
sidecar-push-ist-fips-local:
	./hack/push.sh "vsecm-ist-fips-sidecar" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-fips-sidecar"
