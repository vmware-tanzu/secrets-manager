#!/bin/bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

version=$1
gitRoot=$(git rev-parse --show-toplevel)
helmChartDirName="helm-charts"
helmChartDirectory=$gitRoot/$helmChartDirName
k8sManifestsDirectory=$gitRoot/k8s/$version

if [ "$version" == '' ]; then
    echo "VERSION has to be provided to the make target"
    echo "usage: make k8s-manifests-update VERSION=<helm-chart version>"
    echo "example: make k8s-manifests-update VERSION=0.21.2"
    exit 1
else
    helmChartPath=$helmChartDirectory/$version
fi

if [ ! -d "$helmChartPath" ]; then
    # helm-chart doesn't exists for source version
    echo -e "helm-chart does not exists for version($1) at $helmChartPath,\n"\
        "make sure helm-chart is placed in: $helmChartPath directory"
    exit 1
fi

# Producing k8s manifests, using helm-template command
mkdir -p "$k8sManifestsDirectory" || exit 1

echo "producing manifests for local deployments"
helm template "$helmChartPath" $NAME_TEMPLATE $LOCAL_REGISTRY $DISTROLESSS_IMAGE > "$k8sManifestsDirectory/$version-local-distrolesss.yaml" || exit 1
helm template "$helmChartPath" $NAME_TEMPLATE $LOCAL_REGISTRY $DISTROLESSS_FIPS_IMAGE > "$k8sManifestsDirectory/$version-local-distrolesss-fips.yaml" || exit 1
helm template "$helmChartPath" $NAME_TEMPLATE $LOCAL_REGISTRY $PHOTON_IMAGE > "$k8sManifestsDirectory/$version-local-photon.yaml" || exit 1
helm template "$helmChartPath" $NAME_TEMPLATE $LOCAL_REGISTRY $PHOTON_FIPS_IMAGE > "$k8sManifestsDirectory/$version-local-photon-fips.yaml" || exit 1

echo "producing manifests for remote deployments"
helm template "$helmChartPath" $NAME_TEMPLATE $DISTROLESSS_IMAGE > "$k8sManifestsDirectory/$version-remote-distrolesss.yaml" || exit 1
helm template "$helmChartPath" $NAME_TEMPLATE $DISTROLESSS_FIPS_IMAGE > "$k8sManifestsDirectory/$version-remote-distrolesss-fips.yaml" || exit 1
helm template "$helmChartPath" $NAME_TEMPLATE $PHOTON_IMAGE > "$k8sManifestsDirectory/$version-remote-photon.yaml" || exit 1
helm template "$helmChartPath" $NAME_TEMPLATE $PHOTON_FIPS_IMAGE > "$k8sManifestsDirectory/$version-remote-photon-fips.yaml" || exit 1

echo "Successfully updated manifests, create PR with updated files and merge!!"
exit 0
