#!/usr/bin/env bash

set -e

helm template -n spire-server spire-crds spire-crds \
	 --repo https://spiffe.github.io/helm-charts-hardened/ \
	 -f values-no-openshift.yaml \
		--create-namespace > spire-crds-manifest-no-openshift.yaml

helm template -n spire-server spire spire \
	 --repo https://spiffe.github.io/helm-charts-hardened/ \
	 -f values-no-openshift.yaml \
	 --create-namespace > spire-manifest-no-openshift.yaml
