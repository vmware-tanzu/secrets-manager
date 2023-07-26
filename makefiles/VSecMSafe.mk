# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the “VMware Secrets Manager Safe” into a container image.
safe-bundle-ist:
	./hack/bundle.sh "vsecm-ist-safe" \
		$(VERSION) "dockerfiles/vsecm-ist/safe.Dockerfile"

# Packages the “VMware Secrets Manager Safe” into a container image for FIPS.
safe-bundle-ist-fips:
	./hack/bundle.sh "vsecm-ist-fips-safe" \
		$(VERSION) "dockerfiles/vsecm-ist-fips/safe.Dockerfile"

# Packages the “VMware Secrets Manager Safe” into a container image for Photon OS.
safe-bundle-photon:
	./hack/bundle.sh "vsecm-photon-safe" \
		$(VERSION) "dockerfiles/vsecm-photon/safe.Dockerfile"

# Packages the “VMware Secrets Manager Safe” into a container image for Photon OS and FIPS.
safe-bundle-photon-fips:
	./hack/bundle.sh "vsecm-photon-fips-safe" \
		$(VERSION) "dockerfiles/vsecm-photon-fips/safe.Dockerfile"

# Pushes the “VMware Secrets Manager Safe” container to the public registry.
safe-push-ist:
	./hack/push.sh "vsecm-ist-safe" $(VERSION) "vsecm/vsecm-ist-safe"

# Pushes the “VMware Secrets Manager Safe” container to the public registry.
safe-push-ist-fips:
	./hack/push.sh "vsecm-ist-fips-safe" \
		$(VERSION) "vsecm/vsecm-ist-fips-safe"

# Pushes the “VMware Secrets Manager Safe” (Photon OS) container to the public registry.
safe-push-photon:
	./hack/push.sh "vsecm-photon-safe" \
		$(VERSION) "vsecm/vsecm-photon-safe"

# Pushes the “VMware Secrets Manager Safe” (Photon OS) container to the public registry.
safe-push-photon-fips:
	./hack/push.sh "vsecm-photon-fips-safe" \
		$(VERSION) "vsecm/vsecm-photon-fips-safe"

# Pushes the “VMware Secrets Manager Safe” container image to the local registry.
safe-push-ist-local:
	./hack/push.sh "vsecm-ist-safe" $(VERSION) "localhost:5000/vsecm-ist-safe"

# Pushes the “VMware Secrets Manager Safe” container image to the local registry.
safe-push-ist-fips-local:
	./hack/push.sh "vsecm-ist-fips-safe" \
		$(VERSION) "localhost:5000/vsecm-ist-fips-safe"

# Pushes the “VMware Secrets Manager Safe” (Photon OS) container image to the local registry.
safe-push-photon-local:
	./hack/push.sh "vsecm-photon-safe" \
		$(VERSION) "localhost:5000/vsecm-photon-safe"

# Pushes the “VMware Secrets Manager Safe” (Photon OS) container image to the local registry.
safe-push-photon-fips-local:
	./hack/push.sh "vsecm-photon-fips-safe" \
		$(VERSION) "localhost:5000/vsecm-photon-fips-safe"
