```text
|   Protect your secrets, protect your sensitive data.
:   Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/ keep your secrets... secret
```

## Use Case: Using an Init Container

This example demonstrates how to use **VSecM Init Container** to wait for 
secrets to be allocated to a workload before bootstrapping the workload.

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
images; or you can pull and use pre-deployed images from Docker Hub. The
next two sections describe both approaches respectively.

## Local Deployment

```bash
# Switch to the project folder:
cd $WORKSPACE/vmware-secrets-manager
# Build everything locally:
make build-local
# Deploy the use case:
make example-init-container-deploy-local
# Switch to this use case's folder:
cd $WORKSPACE/vmware-secrets-manager/examples/using_init_container
# Check and make sure that the workload pod is still initializing:
kubectl get po -n default
# Register a secret:
./register.sh
# Verify that the workload pod has initialized:
kubectl get po -n default
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
# Check and make sure that the workload pod is still initializing:
kubectl get po -n default
# Register a secret:
./register.sh
# Verify that the workload pod has initialized:
kubectl get po -n default
# Tail the workload's logs and verify that the secret is there:
./tail.sh
```
