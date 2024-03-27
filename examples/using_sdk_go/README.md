```text
|   Protect your secrets, protect your sensitive data.
:   Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/ keep your secrets... secret
```

## Use Case: Leveraging VMware Secrets Manager (*VSecM*) SDK

This example demonstrates how to use **VSecM SDK** along with your workload.

When you use **VSecM SDK**, you can communicate with **VSecM Safe** directly
and fetch the secrets to your workload whenever you need them.

This approach provides a great deal of flexibility, enabling you to make 
customizations to your code as needed. While adding the **VSecM SDK** as a 
dependency may require some extra effort, it also opens up a range of
features and capabilities that will benefit your project in the long run.

## A Video Is Worth a Lot of Words

[You can watch these tutorial][videos] as to learn how to use **VMware
Secrets Manager** hands-on.

Note that the videos are about the older version of VMware Secrets Manager,
called 'Aegis', but the concepts are the same. Also note that in the meantime
we have added a few more features to **VMware Secrets Manager**, so you will
find differences between the video and the current version of **VMware Secrets
Manager**.

In the videos, we use all the scripts and resource manifests you see in this
folder to demonstrate various **VMware Secrets Manager** use cases.

[videos]: https://vimeo.com/showcase/10074951 "VSecM Showcase"

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
make example-sdk-deploy-local
# Switch to this use case's folder:
cd $WORKSPACE/vmware-secrets-manager/examples/using_sdk_go
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
make example-sdk-deploy
# Switch to this use case's folder:
cd $WORKSPACE/vmware-secrets-manager/examples/using_sdk_go
# Register a secret:
./register.sh
# Tail the workload's logs and verify that the secret is there:
./tail.sh
```
