---
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets… secret
# >/
# <>/' Copyright 2023–present VMware, Inc.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title: Contribute to VSecM
layout: post
prev_url: /docs/use-the-source/
permalink: /docs/contributing/
next_url: /docs/dev-env/
---

<p class="github-button"
><a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0130-contributing.md"
>edit this page on <strong>GitHub</strong> ✏️</a></p>

## Introduction

This section contains instructions to test and develop **VMware Secrets Manager**
locally.

> **📚 Familiarize Yourself with the Contributing Guidelines**
> 
> Please make sure you read the [Contributing Guidelines][contributing]
> and  the [Code of Conduct][coc] on the **VSecM** GitHub repo first.
{: .block-warning}

## Good First Issues for New Contributors

If you are new to **VMware Secrets Manager** or looking for smaller tasks to 
start contributing, we have a set of issues labeled as 
[“*good first issue*”](https://github.com/vmware-tanzu/secrets-manager/labels/good%20first%20issue) 
on our GitHub repository. These issues are a great place to start if you are 
looking to make your first contribution.

### How to Find Good First Issues

1. Navigate to the [Issues tab](https://github.com/vmware-tanzu/secrets-manager/issues) 
   in the GitHub repository.
2. Use the label filter and select the [“*good first issue*”](https://github.com/vmware-tanzu/secrets-manager/labels/good%20first%20issue) 
   label.
3. Browse through the list and pick an issue that interests you.

### Claiming an Issue

Before starting work on an issue, it’s a good practice to comment on it, 
stating that you intend to work on it. This prevents multiple contributors from 
working on the same issue simultaneously.

### Need Help?

If you have questions or need further clarification on a “*good first issue,*” 
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

[fork]: https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/working-with-forks/about-forks
[contributing]: https://github.com/vmware-tanzu/secrets-manager/blob/main/CONTRIBUTING_DCO.md
[coc]: https://github.com/vmware-tanzu/secrets-manager/blob/main/CODE_OF_CONDUCT.md