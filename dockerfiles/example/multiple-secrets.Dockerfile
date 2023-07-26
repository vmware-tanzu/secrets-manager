# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# builder image
FROM golang:1.20.1-alpine3.17 as builder
COPY app /build/app
COPY core /build/core
COPY sdk /build/sdk
COPY examples /build/examples
COPY vendor /build/vendor
COPY go.mod /build/go.mod
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -o env \
  ./examples/multiple-secrets/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -o sloth \
  ./examples/multiple-secrets/busywait/main.go

# generate clean, final image for end users
FROM gcr.io/distroless/static-debian11

LABEL "maintainers"="Volkan Özçelik <volkan@vsecm.com>"
LABEL "version"="0.20.0"
LABEL "website"="https://vsecm.com/"
LABEL "repo"="https://github.com/vmware-tanzu/secrets-manager"
LABEL "documentation"="https://vsecm.com/"
LABEL "contact"="https://vsecm.com/contact/"
LABEL "community"="https://vsecm.com/community"
LABEL "changelog"="https://vsecm.com/changelog"

COPY --from=builder /build/env .
COPY --from=builder /build/sloth .

# executable
ENTRYPOINT [ "./sloth" ]
CMD [ "" ]
