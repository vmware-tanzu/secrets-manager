# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the "Sidecar" use case binary into a container image.
example-sidecar-bundle:
	./hack/bundle.sh "example-using-sidecar" \
		$(VERSION) "dockerfiles/example/sidecar.Dockerfile"

# Pushes the "Sidecar" use case container image to the public registry.
example-sidecar-push:
	./hack/push.sh "example-using-sidecar" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/example-using-sidecar"


# Pushes the "Sidecar" use case container image to the public EKS registry.
example-sidecar-push-eks:
	./hack/push.sh "example-using-sidecar" \
		$(VERSION) "$(VSECM_EKS_REGISTRY_URL)/example-using-sidecar"

# Pushes the "Sidecar" use case container image to the local registry.
example-sidecar-push-local:
	./hack/push.sh "example-using-sidecar" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/example-using-sidecar"

# Deploys the "Sidecar" use case app from the public registry into the cluster.
example-sidecar-deploy:
	./hack/example-sidecar-deploy.sh

# Deploys the "Sidecar" use case app from the local registry into the cluster.
example-sidecar-deploy-local:
	./hack/example-sidecar-deploy-local.sh

# Deploys the "Sidecar" use case app from the public EKS registry into the cluster.
example-sidecar-deploy-eks:
	./hack/example-sidecar-deploy-eks.sh