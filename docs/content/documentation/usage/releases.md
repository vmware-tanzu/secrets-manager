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

title = "VSecM Releases"
weight = 12
+++

**VMware Secrets Manager** signs all of its releases using GitHub's built-in
signing process. We also sign our container images using
[*Docker Content Trust*][docker-content-trust].

The following sections outline how you can verify the authenticity of our
releases.

## Getting the Releases

You can download the latest release from the [GitHub Releases page][releases].

The related container images can be found on [Docker Hub][docker-hub].

[releases]: https://github.com/vmware-tanzu/secrets-manager/releases.
[docker-hub]: https://hub.docker.com/u/vsecm.

## Verifying Code Releases

Our code releases are signed using GitHub's built-in signing process.
To verify a release:

**Clone the repository and navigate to it**:

```bash
git clone https://github.com/vmware-tanzu/secrets-manager.git
cd secrets-manager
```

**Fetch the tags**:

```bash
git fetch --tags
```

**Verify the tag**:

```bash
git tag -v <tag-name>
```

If the signature is valid, you will see a message confirming the signature
check passed.

## Verifying Container Images

We use [Docker Content Trust][docker-content-trust] to sign our Docker images.
To verify the signature of an image, you can enable Docker Content Trust by
setting the `DOCKER_CONTENT_TRUST` environment variable to `1`.

```bash
export DOCKER_CONTENT_TRUST=1
```

After enabling Docker Content Trust, any docker pull command will automatically
verify the image signature before pulling it.

```bash
docker pull vsecm/$yourImage
# For, e.g.: docker pull vsecm/vsecm-ist-safe
```

If the image signature is valid, the image will be pulled; otherwise, you will
receive an error message.

[docker-content-trust]: https://docs.docker.com/engine/security/trust/content_trust/ "Docker Content Trust"

{{ edit() }}
