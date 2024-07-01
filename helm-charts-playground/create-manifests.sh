#!/usr/bin/env bash

set -e

helm template -n spire-server spire-crds spire-crds \
	 --repo https://spiffe.github.io/helm-charts-hardened/ \
	 -f values.yaml \
		--create-namespace > spire-crds-manifest.yaml
 
helm template -n spire-server spire spire \
	 --repo https://spiffe.github.io/helm-charts-hardened/ \
	 -f values.yaml \
	 --create-namespace > spire-manifest.yaml
