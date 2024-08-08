# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the "VSecM Inspector" binary into a container image.
inspector-bundle:
	./hack/bundle.sh "vsecm-inspector" \
		$(VERSION) "dockerfiles/util/inspector.Dockerfile"

# Pushes the "VSecM Inspector" container image to the public registry.
inspector-push:
	./hack/push.sh "vsecm-inspector" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-inspector"

# Pushes the "VSecM Inspector" container image to the public EKS registry.
inspector-push-eks:
	./hack/push.sh "vsecm-inspector" $(VERSION) \
		"$(VSECM_EKS_REGISTRY_URL)/vsecm-inspector"

# Pushes the "VSecM Inspector" container image to the local registry.
inspector-push-local:
	./hack/push.sh "vsecm-inspector" $(VERSION) \
		"$(VSECM_LOCAL_REGISTRY_URL)/vsecm-inspector"

# Deploys the "VSecM Inspector" app from the public registry into the cluster.
inspector-deploy:
	./hack/inspector-deploy.sh
