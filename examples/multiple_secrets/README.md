```text
|   Protect your secrets, protect your sensitive data.
:   Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/ keep your secrets... secret
```

## Use Case: Multiple Secrets

This example demonstrates how to use **VMware Secrets Manager** to register more 
than one secret to your workload.

This demo is a slight variation of the 
[Registering Secrets Using Init Container](../using_init_container)
example.

## A Video Is Worth a Lot of Words

[You can watch these tutorial][videos] as to learn how to use **VMware
Secrets Manager** hands-on.

Note that the video is about the older version of VMware Secrets Manager,
called 'Aegis', but the concepts are the same. Also note that in the meantime
we have added a few more features to **VMware Secrets Manager**, so you will
find differences between the video and the current version of **VMware Secrets
Manager**.

In the videos, we use all the scripts and resource manifests you see in this
folder to demonstrate various **VMware Secrets Manager** use cases.

[video]: https://vimeo.com/v0lkan/vsecm-use-cases "VSecM Use Cases"

## Deployment Options

To follow this use case, you can either locally build and deploy the container
images; or you can pull and use pre-deployed images from Docker Hub.

The following sections describe both of these approaches respectively.

## Local Deployment

```bash
# Switch to the project folder:
cd $WORKSPACE/vmware-secrets-manager
# Build everything locally:
make build-local
# Deploy the use case:
make example-multiple-secrets-deploy-local
# Switch to this use case's folder:
cd $WORKSPACE/vmware-secrets-manager/examples/multiple-secrets
# Register a secret:
./register.sh
# List the secrets.
./secrets.sh
```

## Using Pre-Deployed Images

If you don't want to build the container images locally, you can use
pre-deployed container images.

```bash 
# Switch to the project folder:
cd $WORKSPACE/vmware-secrets-manager
# Deploy the use case from the pre-built image.
make example-multiple-secrets-deploy
# Switch to this use case's folder:
cd $WORKSPACE/vmware-secrets-manager/examples/multiple_secrets
# Register a secret:
./register.sh
# List the secrets.
./secrets.sh
```
