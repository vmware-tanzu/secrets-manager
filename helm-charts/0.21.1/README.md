# VMware Secrets Manager (VSecM) Helm Chart
[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/vsecm)](https://artifacthub.io/packages/helm/vsecm/vsecm)

VMware Secrets Manager keeps your secrets secret. With VSecM, you can rest assured that your sensitive data is always secure and protected. VSecM is perfect for securely storing arbitrary configuration information at a central location and securely dispatching it to workloads.

## Installation

To use VMware Secrets Manager, follow the steps below:

1. Add VMware Secrets Manager Helm repository:

    ```bash
    helm repo add vsecm https://vmware-tanzu.github.io/secrets-manager/
    ```

2. Update helm repository:

    ```bash
    helm repo update
    ```

3. Install VMware Secrets Manager using Helm:

    ```bash
    helm install vsecm vsecm/vsecm --version 0.21.1
    ```

## Options

The following options can be passed to the `helm install` command to set global variables:

- `--set global.deploySpire=<true/false>`: This flag can be passed to install or skip Spire.
- `--set global.baseImage=<distroless/distroless-fips/photon/photos-fips>`: This flag can be passed to install VSecM with the given baseImage Docker image.

Default values are `true` and `distroless` for `global.deploySpire` and `global.baseImage` respectively.

Here's an example command with the above options:

```bash
helm install vsecm vsecm/helm-charts --version 0.21.1 --set global.deploySpire=true --set global.baseImage=distroless
```

Make sure to replace `<true/false>` and `<distroless/distroless-fips/photon/photos-fips>` with the desired values.

## License

This project is licensed under the [BSD 2-Clause License](https://github.com/vmware-tanzu/secrets-manager/blob/main/LICENSE).
