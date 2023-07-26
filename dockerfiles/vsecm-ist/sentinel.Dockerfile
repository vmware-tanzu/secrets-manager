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
RUN mkdir /build
COPY app /build/app
COPY core /build/core
COPY vendor /build/vendor
COPY go.mod /build/go.mod
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -o safe ./app/sentinel/cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -o sloth ./app/sentinel/busywait/main.go

# generate clean, final image for end users
FROM gcr.io/distroless/static-debian11

LABEL "maintainers"="Volkan Özçelik <volkan@vsecm.com>"
LABEL "version"="0.20.0"
LABEL "website"="https://vsecm.com/"
LABEL "repo"="https://github.com/vmware-tanzu/secrets-manager-sentinel"
LABEL "documentation"="https://vsecm.com/"
LABEL "contact"="https://vsecm.com/contact/"
LABEL "community"="https://vsecm.com/community"
LABEL "changelog"="https://vsecm.com/changelog"

# Copy the required binaries
COPY --from=builder /build/safe /bin/safe
COPY --from=builder /build/sloth /bin/sloth

ENV HOSTNAME sentinel

# Prevent root access.
ENV USER nobody
USER nobody

# Keep the container alive.
ENTRYPOINT ["/bin/sloth"]
CMD [""]
