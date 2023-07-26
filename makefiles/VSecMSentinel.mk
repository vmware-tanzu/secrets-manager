# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the “VMware Secrets Manager Sentinel” binary into a container image.
sentinel-bundle-ist:
	./hack/bundle.sh "vsecm-ist-sentinel" \
		$(VERSION) "dockerfiles/vsecm-ist/sentinel.Dockerfile"

# Packages the “VMware Secrets Manager Sentinel” binary into a container image for FIPS.
sentinel-bundle-ist-fips:
	./hack/bundle.sh "vsecm-ist-fips-sentinel" \
		$(VERSION) "dockerfiles/vsecm-ist-fips/sentinel.Dockerfile"

# Packages the “VMware Secrets Manager Sentinel” binary into a container image for Photon OS.
sentinel-bundle-photon:
	./hack/bundle.sh "vsecm-photon-sentinel" \
		$(VERSION) "dockerfiles/vsecm-photon/sentinel.Dockerfile"

# Packages the “VMware Secrets Manager Sentinel” binary into a container image for Photon OS and FIPS.
sentinel-bundle-photon-fips:
	./hack/bundle.sh "vsecm-photon-fips-sentinel" \
		$(VERSION) "dockerfiles/vsecm-photon-fips/sentinel.Dockerfile"

# Pushes the “VMware Secrets Manager Sentinel” container image the the public registry.
sentinel-push-ist:
	./hack/push.sh "vsecm-ist-sentinel" \
		$(VERSION) "vsecm/vsecm-ist-sentinel"

# Pushes the “VMware Secrets Manager Sentinel” (Photon OS) container image to the public registry.
sentinel-push-ist-fips:
	./hack/push.sh "vsecm-ist-fips-sentinel" \
		$(VERSION) "vsecm/vsecm-ist-fips-sentinel"

# Pushes the “VMware Secrets Manager Sentinel” (Photon OS) container image to the public registry.
sentinel-push-photon:
	./hack/push.sh "vsecm-photon-sentinel" \
		$(VERSION) "vsecm/vsecm-photon-sentinel"

# Pushes the “VMware Secrets Manager Sentinel” (Photon OS) container image to the public registry.
sentinel-push-photon-fips:
	./hack/push.sh "vsecm-photon-fips-sentinel" \
		$(VERSION) "vsecm/vsecm-photon-fips-sentinel"

# Pushes the “VMware Secrets Manager Sentinel” container image to the local registry.
sentinel-push-ist-local:
	./hack/push.sh "vsecm-ist-sentinel" \
		$(VERSION) "localhost:5000/vsecm-ist-sentinel"

# Pushes the “VMware Secrets Manager Sentinel” (Photon OS) container image to the local registry.
sentinel-push-ist-fips-local:
	./hack/push.sh "vsecm-ist-fips-sentinel" \
		$(VERSION) "localhost:5000/vsecm-ist-fips-sentinel"

# Pushes the “VMware Secrets Manager Sentinel” (Photon OS) container image to the local registry.
sentinel-push-photon-local:
	./hack/push.sh "vsecm-photon-sentinel" \
		$(VERSION) "localhost:5000/vsecm-photon-sentinel"

# Pushes the “VMware Secrets Manager Sentinel” (Photon OS) container image to the local registry.
sentinel-push-photon-fips-local:
	./hack/push.sh "vsecm-photon-fips-sentinel" \
		$(VERSION) "localhost:5000/vsecm-photon-fips-sentinel"
