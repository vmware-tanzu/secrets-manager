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
	kubectl apply -f ./k8s/${VERSION}/crds
	kubectl apply -f ./k8s/${VERSION}/${VERSION}-remote-distroless.yaml
deploy-fips:
	kubectl apply -f ./k8s/${VERSION}/crds
	kubectl apply -f ./k8s/${VERSION}/${VERSION}-remote-distrolesss-fips.yaml
deploy-photon:
	kubectl apply -f ./k8s/${VERSION}/crds
	kubectl apply -f ./k8s/${VERSION}/${VERSION}-remote-photon.yaml
deploy-photon-fips:
	kubectl apply -f ./k8s/${VERSION}/crds
	kubectl apply -f ./k8s/${VERSION}/${VERSION}-remote-photon-fips.yaml
deploy-local:
	kubectl apply -f ./k8s/${VERSION}/crds
	kubectl apply -f ./k8s/${VERSION}/${VERSION}-local-distroless.yaml
deploy-fips-local:
	kubectl apply -f ./k8s/${VERSION}/crds
	kubectl apply -f ./k8s/${VERSION}/${VERSION}-local-distrolesss-fips.yaml
deploy-photon-local:
	kubectl apply -f ./k8s/${VERSION}/crds
	kubectl apply -f ./k8s/${VERSION}/${VERSION}-local-photon.yaml
deploy-photon-fips-local:
	kubectl apply -f ./k8s/${VERSION}/crds
	kubectl apply -f ./k8s/${VERSION}/${VERSION}-local-photon-fips.yaml

#
# ## Tests ##
#

# Integration tests.
test:
	./hack/test.sh "remote"
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
