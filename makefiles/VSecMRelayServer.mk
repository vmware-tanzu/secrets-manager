# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

relay-server-bundle-ist:
	./hack/bundle.sh "vsecm-ist-relay-server" \
		$(VERSION) "dockerfiles/vsecm-ist/relay-server.Dockerfile"

relay-server-bundle-ist-fips:
	./hack/bundle.sh "vsecm-ist-fips-relay-server" \
		$(VERSION) "dockerfiles/vsecm-ist-fips/relay-server.Dockerfile"

relay-server-push-ist:
	./hack/push.sh "vsecm-ist-relay-server" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-ist-relay-server"

relay-server-push-ist-fips:
	./hack/push.sh "vsecm-ist-fips-relay-server" \
		$(VERSION) "$(VSECM_DOCKERHUB_REGISTRY_URL)/vsecm-ist-fips-relay-server"

relay-server-push-ist-local:
	./hack/push.sh "vsecm-ist-relay-server" $(VERSION) \
		"$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-relay-server"

relay-server-push-ist-fips-local:
	./hack/push.sh "vsecm-ist-fips-relay-server" $(VERSION) \
		"$(VSECM_LOCAL_REGISTRY_URL)/vsecm-ist-fips-relay-server"
