# Aegis Helm Chart

Aegis keeps your secrets secret. With Aegis, you can rest assured that your sensitive data is always secure and protected. Aegis is perfect for securely storing arbitrary configuration information at a central location and securely dispatching it to workloads.

## Installation

To use Aegis, follow the steps below:

1. Add Aegis Helm repository:

    ```bash
    helm repo add aegis https://shieldworks.github.io/aegis/
    ```

2. Update helm repository:

    ```bash
    helm repo update
    ```

3. Install Aegis using Helm:

    ```bash
    helm install aegis aegis/aegis --version 0.1.0
    ```

## Options

The following options can be passed to the `helm install` command to set global variables:

- `--set global.deploySpire=<true/false>`: This flag can be passed to install or skip Spire.
- `--set global.baseImage=<distroless/distroless-fips/photon/photos-fips>`: This flag can be passed to install Aegis with the given baseImage Docker image.

Default values are `true` and `distroless` for `global.deploySpire` and `global.baseImage` respectively.

Here's an example command with the above options:

```bash
helm install aegis aegis/helm-charts --version 0.1.0 --set global.deploySpire=true --set global.baseImage=distroless
```

Make sure to replace `<true/false>` and `<distroless/distroless-fips/photon/photos-fips>` with the desired values.

## License

This project is licensed under the [MIT License](LICENSE).