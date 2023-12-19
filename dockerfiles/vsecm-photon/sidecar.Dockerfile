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
COPY sdk /build/sdk
COPY vendor /build/vendor
COPY go.mod /build/go.mod
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -o vsecm-sidecar ./app/sidecar/cmd/main.go

# generate clean, final image for end users
FROM photon:5.0

LABEL "maintainers"="VSecM Maintainers <maintainers@vsecm.com>"
LABEL "version"="0.21.2"
LABEL "website"="https://vsecm.com/"
LABEL "repo"="https://github.com/vmware-tanzu/secrets-manager"
LABEL "documentation"="https://vsecm.com/"
LABEL "contact"="https://vsecm.com/contact/"
LABEL "community"="https://vsecm.com/community"
LABEL "changelog"="https://vsecm.com/changelog"

COPY --from=builder /build/vsecm-sidecar .

# executable
ENTRYPOINT [ "./vsecm-sidecar" ]
CMD [ "" ]
