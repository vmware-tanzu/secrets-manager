# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the “SDK” use case binary into a container image.
example-sdk-bundle:
	./hack/bundle.sh "example-using-sdk" \
		$(VERSION) "dockerfiles/example/sdk.Dockerfile"

# Pushes the “SDK” use case container image to the public registry.
example-sdk-push:
	./hack/push.sh "example-using-sdk" \
		$(VERSION) "vsecm/example-using-sdk"

# Pushes the “SDK” use case container image to the local registry.
example-sdk-push-local:
	./hack/push.sh "example-using-sdk" \
		$(VERSION) "$(VSECM_LOCAL_REGISTRY_URL)/example-using-sdk"

# Deploys the “SDK” use case app from the public registry into the cluster.
example-sdk-deploy:
	./hack/example-sdk-deploy.sh

# Deploys the “SDK” use case app from the local registry into the cluster.
example-sdk-deploy-local:
	./hack/example-sdk-deploy-local.sh
