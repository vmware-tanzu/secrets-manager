# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# Packages the “Sidecar” use case binary into a container image.
example-sidecar-bundle:
	./hack/bundle.sh "example-using-sidecar" \
		$(VERSION) "dockerfiles/example/sidecar.Dockerfile"

# Pushes the “Sidecar” use case container image to the public registry.
example-sidecar-push:
	./hack/push.sh "example-using-sidecar" \
		$(VERSION) "vsecm/example-using-sidecar"

# Pushes the “Sidecar” use case container image to the local registry.
example-sidecar-push-local:
	./hack/push.sh "example-using-sidecar" \
		$(VERSION) "localhost:5000/example-using-sidecar"

# Deploys the “Sidecar” use case app from the public registry into the cluster.
example-sidecar-deploy:
	./hack/example-sidecar-deploy.sh

# Deploys the “Sidecar” use case app from the local registry into the cluster.
example-sidecar-deploy-local:
	./hack/example-sidecar-deploy-local.sh
