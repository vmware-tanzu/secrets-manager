# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY server.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o server server.go

FROM ubuntu:22.04 AS server

RUN apt-get update && apt-get install -y \
    ca-certificates \
    curl \
    net-tools \
    netcat \
    vim \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 443

CMD ["./server"]
