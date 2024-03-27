```text
|   Protect your secrets, protect your sensitive data.
:   Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/ keep your secrets... secret
```

## Use Case: Leveraging VSecM Sidecar

This example demonstrates how to use **VSecM Sidecar** along with your workload.

When you use **VSecM Sidecar**, you don't need to modify your workload. 
**VSecM Sidecar** can fetch and provide the secrets that your workload needs
automatically.

## Video Tutorials Anyone?

[Watch this showcase to learn how to use **VMware Secrets Manager** hands-on][videos].

If any of the instructions provided here are unclear, the video tutorials will
help explain them in much greater detail. Each video is designed around a
particular topic to keep it concise and to-the-point.

The container image is also used as the **inspector** workload to debug secret
registration during showcasing various scenarios [in the workshop](../vsecm-workshop).

[videos]: https://vimeo.com/showcase/10074951 "VSecM Showcase"

## Deployment Options

To follow this use case, you can either locally build and deploy the container
images; or you can pull and use pre-deployed images from Docker Hub. The
next two sections describe both approaches respectively.

## Local Deployment

```bash
# Switch to the project folder:
cd $WORKSPACE/vmware-secrets-manager 
# Build everything locally:
make build-local
# Deploy the use case:
make example-sidecar-deploy-local
# Switch to this use case's folder:
cd $WORKSPACE/vmware-secrets-manager/examples/using_sidecar
# Register a secret:
./register.sh
# Tail the workload's logs and verify that the secret is there:
./tail.sh
```

## Using Pre-Deployed Images

If you don't want to build the container images locally, you can use
pre-deployed container images.

```bash 
# Switch to the project folder:
cd $WORKSPACE/vmware-secrets-manager 
# Deploy the use case from the pre-built image.
make example-sidecar-deploy
# Switch to this use case's folder:
cd $WORKSPACE/vmware-secrets-manager/examples/using_sidecar
# Register a secret:
./register.sh
# Tail the workload's logs and verify that the secret is there:
./tail.sh
```
