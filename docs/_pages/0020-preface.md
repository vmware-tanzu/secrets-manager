---
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */

title: Preface
layout: post
prev_url: /docs/blog/
permalink: /docs/navigation/
next_url: /docs/about/
---

## Hey, This Website Looks Like a GitBook 📖

Yes, this website is intentionally created like a book. We wanted to make
sure that you have a great experience reading our documentation, and we
believe that following the documentation cover-to-cover as if it were a book
is the best way to learn about **VMware Secrets Manager** (*VSecM*) for
Cloud-Native Apps.

We use a heavily customized version of [the Jekyll GitBook Theme][gitbook-theme]
to achieve this. [You can check the source code for this website on
GitHub][github] to see how that is done.

[gitbook-theme]: https://github.com/sighingnow/jekyll-gitbook "Jekyll GitBook Theme"
[github]: https://github.com/vmware-tanzu/secrets-manager/tree/main/docs "VMware Secrets Manager Documentation on GitHub"

## Terminology: A Tale of Two Secrets

There are two kinds of "*secret*"s mentioned throughout this documentation:

* Secrets that are stored in **VSecM Safe**: When discussing these, they will
  be used like a regular word "secret" or, emphasized "**secret**"; however,
  you will never see them in `monotype text`.
* The other kind of secret is Kubernetes `Secret` objects. Those types
  will be explicitly mentioned as "*Kubernetes `Secret`s*" in the documentation.

We hope this will clarify any confusion going forward.

<p class="github-button">
  <a href="https://github.com/vmware-tanzu/secrets-manager/blob/main/docs/_pages/0020-preface.md">
    Suggest edits ✏️ 
  </a>
</p>
