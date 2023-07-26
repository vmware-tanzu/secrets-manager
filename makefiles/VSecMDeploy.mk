# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

#
# ## Lifecycle ##
#

# Removes the former VSecM deployment without entirely destroying the cluster.
clean:
	./hack/uninstall.sh

# Completely removes the Minikube cluster.
k8s-delete:
	./hack/minikube-delete.sh
# Brings up a fresh Minikube cluster.
k8s-start:
	./hack/minikube-start.sh

# Deploys VSecM to the cluster.
deploy:
	./hack/deploy.sh
deploy-fips:
	./hack/deploy-fips.sh
deploy-photon:
	./hack/deploy-photon.sh
deploy-photon-fips:
	./hack/deploy-photon-fips.sh
deploy-local:
	./hack/deploy-local.sh
deploy-fips-local:
	./hack/deploy-fips-local.sh
deploy-photon-local:
	./hack/deploy-photon-local.sh
deploy-photon-fips-local:
	./hack/deploy-photon-fips-local.sh

#
# ## Tests ##
#

# Integration tests.
test-remote:
	./hack/test.sh "remote"
test-local:
	./hack/test.sh

#
# ## Versioning ##
#

# tags a release
tag:
	./hack/tag.sh $(VERSION)
