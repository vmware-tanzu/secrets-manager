# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

FROM ghcr.io/getzola/zola:v0.17.1 AS zola

COPY ./docs/content /project/content
COPY ./docs/sass /project/sass
COPY ./docs/static /project/static
COPY ./docs/templates /project/templates
COPY ./docs/config.toml /project/config.toml
WORKDIR /project
RUN ["zola", "build"]

FROM ghcr.io/static-web-server/static-web-server:2
WORKDIR /
COPY --from=zola /project/public /public