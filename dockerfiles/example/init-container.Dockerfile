# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

# builder image
FROM golang:1.22.0-alpine3.19 as builder
COPY app /build/app
COPY core /build/core
COPY examples /build/examples
COPY vendor /build/vendor
COPY go.mod /build/go.mod
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -o example \
  ./examples/using-init-container/main.go

# generate clean, final image for end users
FROM gcr.io/distroless/static-debian11

ENV APP_VERSION="0.23.1"

LABEL "maintainers"="VSecM Maintainers <maintainers@vsecm.com>"
LABEL "version"=$APP_VERSION
LABEL "website"="https://vsecm.com/"
LABEL "repo"="https://github.com/vmware-tanzu/secrets-manager"
LABEL "documentation"="https://vsecm.com/"
LABEL "contact"="https://vsecm.com/contact/"
LABEL "community"="https://vsecm.com/community"
LABEL "changelog"="https://vsecm.com/changelog"

COPY --from=builder /build/example .

# executable
ENTRYPOINT [ "./example" ]
CMD [ "" ]
