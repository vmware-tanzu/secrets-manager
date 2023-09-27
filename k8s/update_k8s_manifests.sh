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
localRegistry="global.registry=localhost:5000"
distrolesssImage="global.baseImage=distroless"
distrolesssFipsImage="global.baseImage=distroless-fips"
photonImage="global.baseImage=photon"
photonFipsImage="global.baseImage=photos-fips"

if [ "$version" == '' ]; then
    echo "VERSION has to be provided to the make target"
    echo "usage: make k8s_manifests_update VERSION=<helm-chart version>"
    echo "example: make k8s_manifests_update VERSION=0.21.0"
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
helm template "$helmChartPath" --set $localRegistry --set $distrolesssImage > "$k8sManifestsDirectory/$version.local.distrolesss" || exit 1
helm template "$helmChartPath" --set $localRegistry --set $distrolesssFipsImage > "$k8sManifestsDirectory/$version.local.distrolesss-fips" || exit 1
helm template "$helmChartPath" --set $localRegistry --set $photonImage > "$k8sManifestsDirectory/$version.local.photon" || exit 1
helm template "$helmChartPath" --set $localRegistry --set $photonFipsImage > "$k8sManifestsDirectory/$version.local.photon-fips" || exit 1

echo "producing manifests for remote deployments"
helm template "$helmChartPath" --set $distrolesssImage > "$k8sManifestsDirectory/$version.remote.distrolesss" || exit 1
helm template "$helmChartPath" --set $distrolesssFipsImage > "$k8sManifestsDirectory/$version.remote.distrolesss-fips" || exit 1
helm template "$helmChartPath" --set $photonImage > "$k8sManifestsDirectory/$version.remote.photon" || exit 1
helm template "$helmChartPath" --set $photonFipsImage > "$k8sManifestsDirectory/$version.remote.photon-fips" || exit 1

echo "Successfully updated manifests, create PR with updated files and merge!!"
exit 0