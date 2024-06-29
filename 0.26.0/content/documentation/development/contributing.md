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

title = "Contribute to VSecM"
weight = 11
+++

## Introduction

This section contains instructions to test and develop **VMware Secrets Manager**
locally.

> **📚 Familiarize Yourself with the Contributing Guidelines**
>
> Please make sure you read the [Contributing Guidelines][contributing]
> and  the [Code of Conduct][coc] on the **VSecM** GitHub repo first.

## Good First Issues for New Contributors

If you are new to **VMware Secrets Manager** or looking for smaller tasks to
start contributing, we have a set of issues labeled as
["*good first issue*"](https://github.com/vmware-tanzu/secrets-manager/labels/good%20first%20issue)
on our GitHub repository. These issues are a great place to start if you are
looking to make your first contribution.

### How to Find Good First Issues

1. Navigate to the [Issues tab](https://github.com/vmware-tanzu/secrets-manager/issues)
   in the GitHub repository.
2. Use the label filter and select the ["*good first issue*"](https://github.com/vmware-tanzu/secrets-manager/labels/good%20first%20issue)
   label.
3. Browse through the list and pick an issue that interests you.

### Claiming an Issue

Before starting work on an issue, it's a good practice to comment on it,
stating that you intend to work on it. This prevents multiple contributors from
working on the same issue simultaneously.

### Need Help?

If you have questions or need further clarification on a "*good first issue,*"
feel free to ask in the issue comments or reach out to the maintainers.

## Code Review Requirements

While we value pragmatism over process, we do have some basic requirements for
code reviews to ensure the quality and consistency of the codebase.

### Conducting Code Reviews

1. **Pull Requests**: All code changes must be submitted through a pull request
   (*PR*) on GitHub.
2. **Minimum Reviews**: Each PR must be reviewed by *at least one other person*
   before it can be merged.
3. **Open for Feedback**: PRs are open for comments and suggestions from any
   team member, not just the designated reviewer.

### What Must Be Checked

These are the minimum set of items that must be checked during a code review.
More items may be checked depending on the nature of the change.

1. **Canonical Go**: The code should adhere to canonical Go practices.
2. **Formatting**: The code must pass `gofmt` without any issues.
3. **Consistency**: The code should look like the rest of the codebase,
   as if it were written by a single individual.

### Acceptance Criteria

1. **Approval**: At least one reviewer must approve the PR.
2. **Automated Checks**: All automated tests and checks must pass.
3. **No Conflicts**: Resolve any merge conflicts before merging.

### How to Conduct a Code Review

1. Navigate to the [Pull Requests tab](https://github.com/vmware-tanzu/secrets-manager/pulls)
   in the GitHub repository.
2. Choose a PR that is awaiting review.
3. Review the code changes and provide your feedback, keeping the above criteria
   in mind.
4. If the PR meets all the criteria, approve it; otherwise, request changes and
   provide **constructive** feedback.

## Better Commits

At **VMware Secrets Manager**, we follow the
[Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

You are **strongly encouraged** (*but not required*) to follow the Conventional
Commits specification when writing commit messages.

You are encouraged (*but not required*) to use the
[better-commits][better-commits] tool to help you
create commit messages.

> **`make commit`**
>
> If you have [better-commits][better-commits] installed, then you can execute
> `make commit` to create a commit message.

[better-commits]: https://github.com/Everduin94/better-commits "Better Commits"

## Signing Off Commits

We require a sign-off commit message in the following format on each commit in
pull request with your name and email address:

```txt
Sample commit message.

Signed-off-by: Jane Doe <jane.doe@example.com>
```

Git has a `-s | --sign-off` command line option to append the message
automatically. This will use your default value on your git configure in
`user.name`, `user.email`.

If you miss your sign-off on the last commit, you can amend the commit and then
push to GitHub:

```bash
git commit --amend --signoff
```

If you miss series of commits, you can use `git rebase` with `-i` or `--interactive`
flags to edit and append.

## Add Tests for New Features

As you contribute major new functionalities to the software, it's essential to
adhere to a policy that requires the addition of corresponding tests.

This **policy** underscores the importance of validating new features through
thorough testing. For every significant addition you make, ensure that robust
tests are included.

This practice is **crucial** for maintaining the **VSecM**'s quality, reliability,
and stability. By integrating tests for new features, you contribute not only to
the early detection of defects but also to the facilitation of smoother future
enhancements and modifications.

Your commitment to this testing policy is integral to our collective success in
software development.

## What Technologies Do I Need to Know?

You don't have to be an expert in all of these technologies to contribute to
**VMware Secrets Manager**. However, being familiar with the following
concepts and technologies will help you get started faster.

### Go

**VMware Secrets Manager** is written in [Go](https://golang.org/), so you
should be familiar with the language and its idioms. If you are new to Go,
we recommend going through the [Go Tour](https://tour.golang.org/welcome/1)
and the [Effective Go](https://golang.org/doc/effective_go.html) guide.

### Kubernetes

**VMware Secrets Manager** is a Kubernetes-native application, so you should
be familiar with the basics of Kubernetes. If you are new to Kubernetes, we
recommend going through the
[Kubernetes Basics](https://kubernetes.io/docs/tutorials/kubernetes-basics/).

### Helm

**VMware Secrets Manager** is packaged as a Helm chart, so you should be
familiar with the basics of Helm. If you are new to Helm, we recommend going
through the [Helm Quickstart Guide](https://helm.sh/docs/intro/quickstart/).

### SPIFFE and SPIRE

**VMware Secrets Manager** uses [SPIFFE](https://spiffe.io/) and
[SPIRE](https://spiffe.io/spire/) to establish an identity control plane, so
you should be familiar with the basics of SPIFFE and SPIRE. If you are new to
SPIFFE and SPIRE, we recommend going through the quickstart guides on the
[SPIFFE](https://spiffe.io/docs/latest/spiffe/getting-started/).

### `go-spiffe`

**VMware Secrets Manager** uses the [`go-spiffe`](https://github.com/spiffe/go-spiffe)
to interact with the SPIFFE and SPIRE APIs, so you should be familiar with the
basics of `go-spiffe`. If you are new to `go-spiffe`, we recommend going through
the [documentation](https://pkg.go.dev/github.com/spiffe/go-spiffe/v2).

### `ClusterSPIFFEID` and `ClusterFederatedTrustDomain`

**VMware Secrets Manager** uses the `ClusterSPIFFEID` and `ClusterFederatedTrustDomain`
to dispatch identities to workloads, and federate cluster, respectively. If you
are new to these concepts, we recommend you [check out the SPIRE Controller Manager
repository](https://github.com/spiffe/spire-controller-manager/tree/main).

### VSecM Architecture

**VMware Secrets Manager** has several components, each with its own
responsibilities. [Check out the architecture overview](@/documentation/architecture/overview.md)
to get a high-level understanding of the components and their interactions.

While you are at there, we strongly recommend going through the entire
[documentation](@/documentation/welcome/overview.md) to get a good understanding of the product.

### Docker

We have several [Dockerfiles](https://docs.docker.com/engine/reference/builder/)
in the repository, so you should be familiar with the basics of Docker. If you
are new to Docker, we recommend going through the [Docker Get
Started](https://docs.docker.com/get-started/) guide.

[fork]: https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/about-forks
[contributing]: https://github.com/vmware-tanzu/secrets-manager/blob/main/CONTRIBUTING_DCO.md
[coc]: https://github.com/vmware-tanzu/secrets-manager/blob/main/CODE_OF_CONDUCT.md

{{ edit() }}
