# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the "Init Container" binary into a container image.
example-init-container-bundle:
	./hack/bundle.sh "example-using-init-container" \
		$(VERSION) "dockerfiles/example/init-container.Dockerfile"

# Pushes the "Init Container" container image to the public registry.
example-init-container-push:
	./hack/push.sh "example-using-init-container" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/example-using-init-container"

# Pushes the "Init Container" container image to the local registry.
example-init-container-push-local:
	./hack/push.sh "example-using-init-container" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/example-using-init-container"

# Pushes the "Init Container" container image to the public EKS registry.
example-init-container-push-eks:
	./hack/push.sh "example-using-init-container" \
		$(VERSION) "$(VSECM_EKS_REGISTRY_URL)/example-using-init-container"

# Deploys the "Init Container" app from the public registry into the cluster.
example-init-container-deploy:
	./hack/example-init-container-deploy.sh

# Deploys the "Init Container" app from the local registry into the cluster.
example-init-container-deploy-local:
	./hack/example-init-container-deploy-local.sh

# Deploys the "Init Container" app from the public EKS registry into the cluster.
example-init-container-deploy-eks:
	./hack/example-init-container-deploy-eks.sh
