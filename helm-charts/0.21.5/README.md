# VMware Secrets Manager (VSecM) Helm Chart
[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/vsecm)](https://artifacthub.io/packages/helm/vsecm/vsecm)

VMware Secrets Manager keeps your secrets secret. With VSecM, you can rest assured 
that your sensitive data is always secure and protected. VSecM is perfect for 
securely storing arbitrary configuration information at a central location and 
securely dispatching it to workloads.

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
    helm install vsecm vsecm/vsecm --version 0.21.4
    ```

## Options

The following options can be passed to the `helm install` command to set global 
variables:

*`--set global.deploySpire=<true/false>`: 
  This flag can be passed to install or skip SPIRE.
*`--set global.baseImage=<distroless/distroless-fips/photon/photos-fips>`: 
  This flag can be passed to install VSecM with the given baseImage Docker image.

Default values are `true` and `distroless` for `global.deploySpire` 
and `global.baseImage` respectively.

Here's an example command with the above options:

```bash
helm install vsecm vsecm/helm-charts --version 0.21.4 \
  --set global.deploySpire=true --set global.baseImage=distroless
```

Make sure to replace `<true/false>` and 
`<distroless/distroless-fips/photon/photos-fips>` with the desired values.

## Environment Configuration

**VMware Secrets Manager** can be tweaked further using environment variables.

[Check out **Configuring VSecM** on the official documentation][configuring-vsecm] 
for details.

These environment variable configurations are expose through subcharts. 
You can modify them as follows:

```bash
helm install vsecm vsecm/helm-charts --version 0.21.4 \
--set safe.environments.VSECM_LOG_LEVEL="7"
--set sentinel.environments.VSECM_LOGL_LEVEL="5"
# You can update other environment variables too.
# Most of the time VSecM assumes sane defaults if you don’t set them.
```

[configuring-vsecm]: https://vsecm.com/docs/configuration/

## License

This project is licensed under the [BSD 2-Clause License](https://github.com/vmware-tanzu/secrets-manager/blob/main/LICENSE).
