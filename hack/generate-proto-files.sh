#!/usr/bin/env bash

# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

if ! command -v go &> /dev/null
then
    echo "Go binary could not be found. Please install go first."
    exit 1
fi

# Install or update the Google protocol buffers compiler plugin for Go.
go get -u google.golang.org/protobuf/cmd/protoc-gen-go

# Change directory to the logger package within the Sentinel application.
cd "$(dirname "$0")/../core/log/rpc/" || exit

# Set the environment variable GO_PATH to the Go workspace directory.
export GO_PATH=~/go

# Add the Go bin directory to the system PATH.
export PATH=$PATH:/$GO_PATH/bin

# Compile the log.proto file into Go source code using protocol buffers.
# Generate both standard Go code and gRPC service code.
protoc --proto_path=. \
       --go_out=./generated \
       --go-grpc_out=./generated \
       --go_opt=paths=source_relative \
       --go-grpc_opt=paths=source_relative \
       log.proto

# Download the required dependencies specified in go.mod and go.sum files to the local vendor directory.
go mod vendor