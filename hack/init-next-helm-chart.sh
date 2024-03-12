#!/bin/bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# initialize variables
BOLD="\033[1m"
gitRoot=$(git rev-parse --show-toplevel)
helmChartDirName="helm-charts"
filesToBeUpdated=("Chart.yaml" "charts/safe/Chart.yaml" "charts/sentinel/Chart.yaml" "charts/spire/Chart.yaml")

if [[ "$1" == "-h" || "$1" == "--help" ]]; then
    echo "This script is used to initialize new VSecM helm-chart from an existing helm-chart"
    echo "-----------------------------------------------------------------"
    echo "Usage: ${gitRoot}/hack/init-next-helm-chart.sh <base-version> <new-version>"
    echo "<base-version>: existing helm-chart version, this will be used as base helm-chart"
    echo "<new-version>: new helm-chart will be initialized for this version"
    echo "-----------------------------------------------------------------"
    echo "example: ${gitRoot}/hack/init-next-helm-chart.sh 0.22.2 0.22.4"
    exit 0
fi

if [ "$#" -eq 2 ]; then
    # base helm chart and new helm chart version is passed
    baseHelmChartVersion=$1
    newHelmChartVersion=$2
    srcBaseHelmChartPath=$gitRoot/$helmChartDirName/$baseHelmChartVersion
    newHelmChartPath=$gitRoot/$helmChartDirName/$newHelmChartVersion
else
    echo "Insufficient arguments, base helm chart and new helm chart versions have to be provided"
    exit 1
fi

# validate:
# 1. if helm-chart exists for source helm-chart
# 2. helm-chart has already be initialized for next release
if [ ! -d "$srcBaseHelmChartPath" ]; then
    # helm-chart doesn't exists for source version
    echo "helm-chart doesn't exist for base version($1) at $srcBaseHelmChartPath"
    exit 1
elif [ -d "$newHelmChartPath" ]; then
    # helm-chart already present for version
    echo "helm-chart already present for version($2) at $newHelmChartPath"
    exit 1
fi

echo "Base helm-chart version: $baseHelmChartVersion"
echo "New helm-chart version: $newHelmChartVersion"
echo "Base helm-chart path: $srcBaseHelmChartPath"

# Create the new directory
mkdir -p "$newHelmChartPath" || exit 1

# Copy all contents of the source helm-chart to the new helm-chart directory
cp -r "$srcBaseHelmChartPath"/* "$newHelmChartPath" || exit 1

# update required files in new initialized helm-chart
for file_name in "${filesToBeUpdated[@]}"
do
    # updating version, appVersion and dependencies version in chart.yaml
    sed -i -e "s/^version: ${baseHelmChartVersion}/version: ${newHelmChartVersion}/;\
    s/^appVersion: \"${baseHelmChartVersion}\"/appVersion: \"${newHelmChartVersion}\"/;\
    s/^    version: ${baseHelmChartVersion}/    version: ${newHelmChartVersion}/"\
    "${newHelmChartPath}/${file_name}" || exit 1
    # deletion backup file
    rm "${newHelmChartPath}/${file_name}-e"
done

# success
echo "\m/ helm-chart for next release($newHelmChartVersion) is successfully initialized at ${newHelmChartPath} and ready for development"

echo "creating initial commit for helm-chart ${newHelmChartVersion}"
localBranchName=initializing-helm-chart/"${newHelmChartVersion}"
git checkout -b "${localBranchName}"
git add "${newHelmChartPath}"
git commit -S -s -m "Introducing initial helm-chart for version ${newHelmChartVersion}"
git push origin "${localBranchName}"
printf '\n%s***********************************************************\n'"$BOLD"
echo -e "Click on below link to create pull-request and merge the pull-request"
echo -e "https://github.com/vmware-tanzu/secrets-manager/compare/main...$localBranchName"
printf '\n%s***********************************************************\n'"$BOLD"

exit 0
