# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the “VMware Secrets Manager Sidecar” binary into a container image.
sidecar-bundle-ist:
	./hack/bundle.sh "vsecm-ist-sidecar" \
		$(VERSION) "dockerfiles/vsecm-ist/sidecar.Dockerfile"

# Packages the “VMware Secrets Manager Sidecar” binary into a container image for FIPS.
sidecar-bundle-ist-fips:
	./hack/bundle.sh "vsecm-ist-fips-sidecar" \
		$(VERSION) "dockerfiles/vsecm-ist-fips/sidecar.Dockerfile"

# Packages the “VMware Secrets Manager Sidecar” binary into a container image for Photon OS.
sidecar-bundle-photon:
	./hack/bundle.sh "vsecm-photon-sidecar" \
		$(VERSION) "dockerfiles/vsecm-photon/sidecar.Dockerfile"

# Packages the “VMware Secrets Manager Sidecar” binary into a container image for Photon OS and FIPS.
sidecar-bundle-photon-fips:
	./hack/bundle.sh "vsecm-photon-fips-sidecar" \
		$(VERSION) "dockerfiles/vsecm-photon-fips/sidecar.Dockerfile"

# Pushes the “VMware Secrets Manager Sidecar” container image to the public registry.
sidecar-push-ist:
	./hack/push.sh "vsecm-ist-sidecar" \
		$(VERSION) "vsecm/vsecm-ist-sidecar"

# Pushes the “VMware Secrets Manager Sidecar” (FIPS) container image to the public registry.
sidecar-push-ist-fips:
	./hack/push.sh "vsecm-ist-fips-sidecar" \
		$(VERSION) "vsecm/vsecm-ist-fips-sidecar"

# Pushes the “VMware Secrets Manager Sidecar” (Photon OS) container image to the public registry.
sidecar-push-photon:
	./hack/push.sh "vsecm-photon-sidecar" \
		$(VERSION) "vsecm/vsecm-photon-sidecar"

# Pushes the “VMware Secrets Manager Sidecar” (Photon OS and FIPS) container image to the public registry.
sidecar-push-photon-fips:
	./hack/push.sh "vsecm-photon-fips-sidecar" \
		$(VERSION) "vsecm/vsecm-photon-fips-sidecar"

# Pushes the “VMware Secrets Manager Sidecar” container image to the local registry.
sidecar-push-ist-local:
	./hack/push.sh "vsecm-ist-sidecar" \
		$(VERSION) "localhost:5000/vsecm-ist-sidecar"

# Pushes the “VMware Secrets Manager Sidecar” (FIPS) container image to the local registry.
sidecar-push-ist-fips-local:
	./hack/push.sh "vsecm-ist-fips-sidecar" \
		$(VERSION) "localhost:5000/vsecm-ist-fips-sidecar"

# Pushes the “VMware Secrets Manager Sidecar” (Photon OS) container image to the local registry.
sidecar-push-photon-local:
	./hack/push.sh "vsecm-photon-sidecar" \
		$(VERSION) "localhost:5000/vsecm-photon-sidecar"

# Pushes the “VMware Secrets Manager Sidecar” (Photon OS and FIPS) container image to the local registry.
sidecar-push-photon-fips-local:
	./hack/push.sh "vsecm-photon-fips-sidecar" \
		$(VERSION) "localhost:5000/vsecm-photon-fips-sidecar"
