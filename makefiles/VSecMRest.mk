# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the “VSecM Sentinel” binary into a container image.
rest-bundle-ist:
	./hack/bundle.sh "vsecm-ist-rest" \
		$(VERSION) "dockerfiles/vsecm-ist/rest.Dockerfile"

# Packages the “VSecM Sentinel” binary into a container image for FIPS.
rest-bundle-ist-fips:
	./hack/bundle.sh "vsecm-ist-fips-rest" \
		$(VERSION) "dockerfiles/vsecm-ist-fips/rest.Dockerfile"

# Packages the “VSecM Sentinel” binary into a container image for Photon OS.
rest-bundle-photon:
	./hack/bundle.sh "vsecm-photon-rest" \
		$(VERSION) "dockerfiles/vsecm-photon/rest.Dockerfile"

# Packages the “VSecM Sentinel” binary into a container image for Photon OS and FIPS.
rest-bundle-photon-fips:
	./hack/bundle.sh "vsecm-photon-fips-rest" \
		$(VERSION) "dockerfiles/vsecm-photon-fips/rest.Dockerfile"

# Pushes the “VSecM Sentinel” container image the the public registry.
rest-push-ist:
	./hack/push.sh "vsecm-ist-rest" \
		$(VERSION) "vsecm/vsecm-ist-rest"

# Pushes the “VSecM Sentinel” (Photon OS) container image to the public registry.
rest-push-ist-fips:
	./hack/push.sh "vsecm-ist-fips-rest" \
		$(VERSION) "vsecm/vsecm-ist-fips-rest"

# Pushes the “VSecM Sentinel” (Photon OS) container image to the public registry.
rest-push-photon:
	./hack/push.sh "vsecm-photon-rest" \
		$(VERSION) "vsecm/vsecm-photon-rest"

# Pushes the “VSecM Sentinel” (Photon OS) container image to the public registry.
rest-push-photon-fips:
	./hack/push.sh "vsecm-photon-fips-rest" \
		$(VERSION) "vsecm/vsecm-photon-fips-rest"

# Pushes the “VSecM Sentinel” container image to the local registry.
rest-push-ist-local:
	./hack/push.sh "vsecm-ist-rest" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-rest"

# Pushes the “VSecM Sentinel” (Photon OS) container image to the local registry.
rest-push-ist-fips-local:
	./hack/push.sh "vsecm-ist-fips-rest" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-fips-rest"

# Pushes the “VSecM Sentinel” (Photon OS) container image to the local registry.
rest-push-photon-local:
	./hack/push.sh "vsecm-photon-rest" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-photon-rest"

# Pushes the “VSecM Sentinel” (Photon OS) container image to the local registry.
rest-push-photon-fips-local:
	./hack/push.sh "vsecm-photon-fips-rest" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/vsecm-photon-fips-rest"
