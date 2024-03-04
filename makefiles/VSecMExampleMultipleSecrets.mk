# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the "multiple secrets" use case binary into a container image.
example-multiple-secrets-bundle:
	./hack/bundle.sh "example-multiple-secrets" \
		$(VERSION) "dockerfiles/example/multiple-secrets.Dockerfile"

# Pushes the "multiple secrets" use case container image to the public registry.
example-multiple-secrets-push:
	./hack/push.sh "example-multiple-secrets" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/example-multiple-secrets"

# Pushes the "multiple secrets" use case container image to the local registry.
example-multiple-secrets-push-local:
	./hack/push.sh "example-multiple-secrets" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/example-multiple-secrets"

# Pushes the "multiple secrets" use case container image to the public EKS registry.
example-multiple-secrets-push-eks:
	./hack/push.sh "example-multiple-secrets" \
		$(VERSION) "$(VSECM_EKS_REGISTRY_URL)/example-multiple-secrets"

# Deploys the "multiple secrets" use case app from the public registry into the cluster.
example-multiple-secrets-deploy:
	./hack/example-multiple-secrets-deploy.sh

# Deploys the "multiple secrets" use case app from the local registry into the cluster.
example-multiple-secrets-deploy-local:
	./hack/example-multiple-secrets-deploy-local.sh
