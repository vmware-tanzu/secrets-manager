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

readonly k8s="k8s"
readonly local="local"
readonly remote="remote"
readonly crds="crds"
readonly helmChartDirName="helm-charts"

version=$1
gitRoot=$(git rev-parse --show-toplevel)
helmChartDirectory=$gitRoot/$helmChartDirName
k8sManifestsDirectory=$gitRoot/$k8s/$version
localManifests="$k8sManifestsDirectory/$local/"
remoteManifests="$k8sManifestsDirectory/$remote/"

# producing k8s manifests, using helm-template command
function produceK8sManifests() {
    # initializing directory structure if doesn't exists
    mkdir -p "$k8sManifestsDirectory/$crds" || exit 1
    mkdir -p $localManifests || exit 1
    mkdir -p $remoteManifests || exit 1

    echo "copying spire CRDs"
    cp "$helmChartPath/$crds/"* "$k8sManifestsDirectory/$crds/"

    echo "producing manifests for spire deployments"
    helm template "$helmChartPath" $NAME_TEMPLATE $DEPLOY_SAFE_FALSE $DEPLOY_SENTINEL_FALSE > $k8sManifestsDirectory/spire.yaml || exit 1

    echo "producing manifests for vsecm local deployments"
    helm template "$helmChartPath" $NAME_TEMPLATE $LOCAL_REGISTRY $DISTROLESSS_IMAGE $DEPLOY_SPIRE_FALSE > $localManifests/vsecm-distroless.yaml || exit 1
    helm template "$helmChartPath" $NAME_TEMPLATE $LOCAL_REGISTRY $DISTROLESSS_FIPS_IMAGE $DEPLOY_SPIRE_FALSE > $localManifests/vsecm-distroless-fips.yaml || exit 1
    helm template "$helmChartPath" $NAME_TEMPLATE $LOCAL_REGISTRY $PHOTON_IMAGE $DEPLOY_SPIRE_FALSE > $localManifests/vsecm-photon.yaml || exit 1
    helm template "$helmChartPath" $NAME_TEMPLATE $LOCAL_REGISTRY $PHOTON_FIPS_IMAGE $DEPLOY_SPIRE_FALSE > $localManifests/vsecm-photon-fips.yaml || exit 1

    echo "producing manifests for vsecm remote deployments"
    helm template "$helmChartPath" $NAME_TEMPLATE $DISTROLESSS_IMAGE $DEPLOY_SPIRE_FALSE > $remoteManifests/vsecm-distroless.yaml || exit 1
    helm template "$helmChartPath" $NAME_TEMPLATE $DISTROLESSS_FIPS_IMAGE $DEPLOY_SPIRE_FALSE > $remoteManifests/vsecm-distroless-fips.yaml || exit 1
    helm template "$helmChartPath" $NAME_TEMPLATE $PHOTON_IMAGE $DEPLOY_SPIRE_FALSE > $remoteManifests/vsecm-photon.yaml || exit 1
    helm template "$helmChartPath" $NAME_TEMPLATE $PHOTON_FIPS_IMAGE $DEPLOY_SPIRE_FALSE > $remoteManifests/vsecm-photon-fips.yaml || exit 1
}

# validating version
if [ "$version" == '' ]; then
    echo "VERSION has to be provided to the make target"
    echo "usage: make k8s-manifests-update VERSION=<helm-chart version>"
    echo "example: make k8s-manifests-update VERSION=0.21.2"
    exit 1
else
    helmChartPath=$helmChartDirectory/$version
fi

# checking if helm-chart exists for requested version
if [ ! -d "$helmChartPath" ]; then
    # helm-chart doesn't exists for source version
    echo -e "helm-chart does not exists for version($1) at $helmChartPath,\n"\
        "make sure helm-chart is placed in: $helmChartPath directory"
    exit 1
fi

# producing k8s manifests for vsecm deployments
produceK8sManifests

echo "Successfully updated manifests, create PR with updated files and merge!!"
exit 0
