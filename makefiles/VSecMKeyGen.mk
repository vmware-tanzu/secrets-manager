# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the "Keygen" binary into a container image.
keygen-bundle:
	./hack/bundle.sh "vsecm-keygen" \
		$(VERSION) "dockerfiles/util/keygen.Dockerfile"

# Pushes the "Keygen" container image to the public registry.
keygen-push:
	./hack/push.sh "vsecm-keygen" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-keygen"

# Pushes the "Keygen" container image to the public EKS registry.
keygen-push-eks:
	./hack/push.sh "vsecm-keygen" $(VERSION) \
		"$(VSECM_EKS_REGISTRY_URL)/vsecm-keygen"

# Pushes the "Keygen" container image to the local registry.
keygen-push-local:
	./hack/push.sh "vsecm-keygen" $(VERSION) \
		"$(VSECM_LOCAL_REGISTRY_URL)/vsecm-keygen"
