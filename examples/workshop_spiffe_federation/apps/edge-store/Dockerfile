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

COPY client.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o client client.go

FROM alpine:latest AS client

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/client .

CMD ["./client"]
