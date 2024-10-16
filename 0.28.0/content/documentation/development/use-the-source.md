+++
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title = "Use the Source"
weight = 10
+++

## Introduction

This section describes how to build **VMware Secrets Manager** from source.

For a more detailed walkthrough about how to contribute to **VMware Secrets
Manager**, see the [**Contributing**](@/documentation/development/contributing.md) section.

## Prerequisites

Make you have the following installed on your system:

* [`make`](https://www.gnu.org/software/make/)
* [`git`](https://git-scm.com/)
* [`docker`](https://www.docker.com/)
* [`protoc`](https://grpc.io/docs/protoc-installation/)

## Clone the Project

```bash
cd $WORKSPACE
git clone https://github.com/vmware-tanzu/secrets-manager.git
cd secrets-manager
```

## Build the Project

Make sure you have a running local Docker daemon and execute the following:

```bash
make build-local
```

## Generating Protocol Buffers

You might need to generate the protocol buffers if you are working on the 
**VSecM** API. To do so, execute the following:

```bash
make generate-proto-files
```

If this command fails, you might need to install the `protoc` compiler:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Updating Vendor Dependencies

If the project fails to build, you might need to update the vendor dependencies. 
To do so, execute the following on the project root:

```bash
cd $WORKSPACE/secrets-manager
go mod tidy
go mod vendor
```

## That's All

That's it 🎉. You now have images of **VMware Secrets Manager** and other
related components built locally on your Docker registry.

## Next Up

For a more detailed guide about how you can use these local container images
in your custer [check out the **Contributing** section](@/documentation/development/contributing.md).

{{ edit() }}
