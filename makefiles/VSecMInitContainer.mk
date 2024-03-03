# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the "VSecM Init Container" binary into a container image.
init-container-bundle-ist:
	./hack/bundle.sh "vsecm-ist-init-container" \
		$(VERSION) "dockerfiles/vsecm-ist/init-container.Dockerfile"

# Packages the "VSecM Init Container" binary into a container image for FIPS.
init-container-bundle-ist-fips:
	./hack/bundle.sh "vsecm-ist-fips-init-container" \
		$(VERSION) "dockerfiles/vsecm-ist-fips/init-container.Dockerfile"

# Packages the "VSecM Init Container" binary into a container image for Photon OS.
init-container-bundle-photon:
	./hack/bundle.sh "vsecm-photon-init-container" \
		$(VERSION) "dockerfiles/vsecm-photon/init-container.Dockerfile"

# Packages the "VSecM Init Container" binary into a container image for Photon OS and FIPS.
init-container-bundle-photon-fips:
	./hack/bundle.sh "vsecm-photon-fips-init-container" \
		$(VERSION) "dockerfiles/vsecm-photon-fips/init-container.Dockerfile"

# Pushes the "VSecM Init Container" container image to the public registry.
init-container-push-ist:
	./hack/push.sh "vsecm-ist-init-container" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-ist-init-container"

# Pushes the "VSecM Init Container" (container image to the public EKS registry.
init-container-push-ist-eks:
	./hack/push.sh "vsecm-ist-init-container" \
		$(VERSION) "$(VSECM_EKS_REGISTRY_URL)/vsecm-ist-init-container"

# Pushes the "VSecM Init Container" (FIPS) container image to the public registry.
init-container-push-ist-fips:
	./hack/push.sh "vsecm-ist-fips-init-container" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-ist-fips-init-container"

# Pushes the "VSecM Init Container" (FIPS) container image to the public EKS registry.
init-container-push-ist-fips-eks:
	./hack/push.sh "vsecm-ist-fips-init-container" \
		$(VERSION) "$(VSECM_EKS_REGISTRY_URL)/vsecm-ist-fips-init-container"

# Pushes the "VSecM Init Container" (Photon OS) container image to the public registry.
init-container-push-photon:
	./hack/push.sh "vsecm-photon-init-container" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-photon-init-container"

# Pushes the "VSecM Init Container" (Photon OS) container image to the public EKS registry.
init-container-push-photon-eks:
	./hack/push.sh "vsecm-photon-init-container" \
		$(VERSION) "$(VSECM_EKS_REGISTRY_URL)/vsecm-photon-init-container"

# Pushes the "VSecM Init Container" (Photon OS and FIPS) container image to the public registry.
init-container-push-photon-fips:
	./hack/push.sh "vsecm-photon-fips-init-container" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-photon-fips-init-container"

# Pushes the "VSecM Init Container" (Photon OS and FIPS) container image to the public EKS registry.
init-container-push-photon-fips-eks:
	./hack/push.sh "vsecm-photon-fips-init-container" \
		$(VERSION) "$(VSECM_EKS_REGISTRY_URL)/vsecm-photon-fips-init-container"

# Pushes the "VSecM Init Container" container image to the local registry.
init-container-push-ist-local:
	./hack/push.sh "vsecm-ist-init-container" $(VERSION) \
		"$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-init-container"

init-container-push-ist-fips-local:
	./hack/push.sh "vsecm-ist-fips-init-container" $(VERSION) \
		"$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-fips-init-container"

# Pushes the "VSecM Init Container" (Photon OS) container image to the local registry.
init-container-push-photon-local:
	./hack/push.sh "vsecm-photon-init-container" $(VERSION) \
		"$(VSECM_LOCAL_REGISTRY_URL)/vsecm-photon-init-container"

# Pushes the "VSecM Init Container" (Photon OS and FIPS) container image to the local registry.
init-container-push-photon-fips-local:
	./hack/push.sh "vsecm-photon-fips-init-container" $(VERSION) \
		"$(VSECM_LOCAL_REGISTRY_URL)/vsecm-photon-fips-init-container"
