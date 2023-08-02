# VMware Secrets Manager Helm Chart Release

This [README.md](./README.md) file provides instructions on how to release a new version of the VMware Secrets Manager Helm Chart.

## Prerequisites

Before proceeding with the release process, please ensure that the following prerequisites are met:

- The necessary changes to the VMware Secrets Manager Helm Charts have been made in the `helm-charts/` directory.
- The VMware Secrets Manager Helm Charts have been tested with the latest changes using the command:

  ```bash
  helm install vsecm helm-charts/
  ```

  If everything looks good, you can proceed with the release process.

## Release Process

1. Package the Helm Charts:
   Run the following command to package the Helm Charts:

   ```bash
   git checkout main
   helm package helm-charts/ --version=<version>
   ```

   This command will produce an VSecM Helm Chart package. For example: `vsecm-0.1.0.tgz`

2. Checkout the `gh-pages` branch:

   ```bash
   git checkout gh-pages
   ```

3. Generate the Helm Repo Index:
   Run the following command to generate the Helm repository index file and merge it with the existing `index.yaml` file:

   ```bash
   helm repo index ./ --merge index.yaml
   ```

4. Checkout a topic branch:

   ```bash
   git checkout -b <topic_branch_name>
   ```

5. Add the Helm Chart package and index.yaml file changes and push to create a PR:

   ```bash
   git add vsecm-0.1.0.tgz index.yaml
   git commit -s -m "<commit-message>"
   git push origin <topic_branch_name>
   ```

6. Create a pull request and merge the above changes to the `gh-pages` branch.
