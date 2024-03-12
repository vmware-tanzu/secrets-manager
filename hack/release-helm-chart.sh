#!/bin/bash

# initialise variables
BOLD="\033[1m"
gitRoot=$(git rev-parse --show-toplevel)
helmChartDirName="helm-charts"
baseHelmChartDirectory=$gitRoot/$helmChartDirName

if [[ "$1" == "-h" || "$1" == "--help" ]]; then
    echo "This script is used to release helm-chart"
    echo "-----------------------------------------------------------------"
    echo "Usage: ${gitRoot}/hack/release-helm-chart.sh <version>"
    echo "<version>: helm-chart version to be releases"
    echo "-----------------------------------------------------------------"
    echo "example: ${gitRoot}/hack/init-next-helm-chart.sh 0.22.2 0.22.4"
    exit 0
fi

if [ "$#" -eq 1 ]; then
    # helm-chart version to be release is passed as argument
    releaseHelmChartVersion=$1
    releaseHelmChartPath=$baseHelmChartDirectory/$releaseHelmChartVersion
    localBranchName="releasing-helm-chart/$releaseHelmChartVersion"
else
    echo "Insufficient arguments, helm-chart release version has to be provided as argument"
    echo "execute: release-helm-chart.sh script with -h or --help for more help"
    exit 1
fi

# At this point, we have helm charts in the $releaseHelmChartPath directory.

# Print warning message
printf '\n%s***********************************************************\n'"$BOLD"
printf 'WARNING: You are about to release helm-chart version %s'"$releaseHelmChartVersion\n"
printf 'Before proceeding with the release process, please ensure that the following prerequisites are met:\n\n'
echo -e "1. The necessary changes to the Helm Charts have been made in the helm-charts/$releaseHelmChartVersion directory"
echo -e "   Do not forget to update image tags in helm-charts/$releaseHelmChartVersion/values.yaml\n"
printf '2. The VMware Secrets Manager Helm Charts deployment have been tested using below command:\n'
printf "\t-------------------------------------\n"
printf '\thelm install vsecm helm-charts/%s\n' "$releaseHelmChartVersion"
printf "\t-------------------------------------\n\n"
printf "And if everything looks good, you can proceed with the release process.\n"
printf '%s**********************************************************\n\n'"$BOLD"

# Ask for confirmation
read -rp "Are you sure you want to continue? (y/n): " choice

# Check user input
case "$choice" in
  y|Y )
    echo "Continuing with release process..."

    git checkout main
    git pull origin main

    cd "$gitRoot" || exit 1

    helm package "$releaseHelmChartPath/" --version="$releaseHelmChartVersion"

    git checkout gh-pages

    echo "generate the Helm Repo Index"
    helm repo index ./ --merge index.yaml

    git checkout -b "$localBranchName"
    git add vsecm-"$releaseHelmChartVersion".tgz index.yaml

    echo "creating commit"
    git commit -S -s -m "Releasing helm-chart for version $releaseHelmChartVersion"
    git push origin "$localBranchName"

    printf '\n%s***********************************************************\n'"$BOLD"

    echo -e "Click on below link to create pull-request and merge the pull-request"
    echo -e "https://github.com/vmware-tanzu/secrets-manager/compare/gh-pages...$localBranchName"

    printf '%s***********************************************************\n'"$BOLD"

    exit 0
    ;;
  n|N )
    echo "Script terminated."
    exit 0
    ;;
  * )
    echo "Invalid choice. Script terminated."
    exit 1
    ;;
esac
