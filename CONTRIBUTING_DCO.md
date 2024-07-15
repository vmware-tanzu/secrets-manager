# Contributing

![VSecM Logo](https://github.com/vmware-tanzu/secrets-manager/assets/1041224/885c11ac-7269-4344-a376-0d0a0fb082a7)

## Contributing to VMware Secrets Manager

We welcome contributions from the community and first want to thank you for 
taking the time to contribute!

Please familiarize yourself with the 
[Code of Conduct](https://github.com/vmware/.github/blob/main/CODE_OF_CONDUCT.md) 
before contributing.

Before you start working with secrets-manager, please read our 
[Developer Certificate of Origin](https://cla.vmware.com/dco). All contributions 
to this repository must be signed as described on that page. Your signature 
certifies that you wrote the patch or have the right to pass it on as an 
open-source patch.

We appreciate any help, be it in the form of code, documentation, design,
or even bug reports and feature requests.

When contributing to this repository, please first discuss the change you wish
to make via an issue, email, or any other method before making a change.
This way, we can avoid misunderstandings and wasted effort.

Please note that [we have a code of conduct](CODE_OF_CONDUCT.md). We expect all
contributors to adhere to it in all interactions with the project.

## Ways to contribute

We welcome many different types of contributions and not all of them need a 
Pull request. Contributions may include:

* New features and proposals
* Documentation
* Bug fixes
* Issue Triage
* Answering questions and giving feedback
* Helping to onboard new contributors
* Other related activities

## Getting started

Please [read the developer guide](https://vsecm.com/use-the-source/) to 
learn how to build, deploy, and test **VMware Secrets Manager** from the
source. 

The developer guide also includes common errors that you might find when
building, deploying, and testing **VMware Secrets Manager**. 

The guide and other documentation is maintained at the `./docs` folder in this 
repository. You are more than welcome to contribute to it if you find anything
that's missing or needs improvement.

In addition, please [follow the quickstart manual](https://vsecm.com/quickstart/)
that is specifically designed for you to experiment with **VMware Secrets Manager**
and get familiar with its components.

The developer guide includes how to run tests too. Please make sure you are
able to build the project locally and execute `make test-local` without 
any errors before creating a pull request.

## Contribution Flow

This is a rough outline of what a contributor's workflow looks like:

* Make a fork of the repository within your GitHub account
* Create a topic branch in your fork from where you want to base your work
* Make commits of logical units
* Make sure your commit messages are with the proper format, 
  quality and descriptiveness (*see below*)
* Adhere to the code standards described below.
* Push your changes to the topic branch in your fork 
* Ensure all components build and function properly on a local
  Kubernetes cluster (*such as Minikube*).
* Update necessary `README.md` and other documents to reflect your changes. 
* Keep pull requests as granular as possible. Reviewing large amounts of code
  can be error-prone and time-consuming for the reviewers.
* Create a pull request containing that commit

* Engage in the discussion under the pull request and proceed accordingly.

## Pull Request Checklist

Before submitting your pull request, we advise you to use the following:

1. Check if your code changes will pass local tests 
   (*i.e., `make test-local` should exit with a `0` success status code*).
2. Ensure your commit messages are descriptive. We follow the conventions 
   on [How to Write a Git Commit Message](http://chris.beams.io/posts/git-commit/).
   Be sure to include any related GitHub issue references in the commit message. 
   See [GFM syntax](https://guides.github.com/features/mastering-markdown/#GitHub-flavored-markdown) 
   for referencing issues and commits.
3. Check the commits and commits messages and ensure they are free from typos.

## Reporting Bugs and Creating Issues

For specifics on what to include in your report, please follow the guidelines 
in the issue and pull request templates when available.


## Ask for Help

The best way to reach us with a question when contributing is to ask on:

* The original GitHub issue
* [**Our Slack Workspace**][slack-invite]

### Code Standards

In **VMware Secrets Manager**, we aim for a unified and clean codebase.
When contributing, please try to match the style of the code that you see in
the file you're working on. The file should look as if it was authored by a
single person after your changes.

For Go files, we require that you run `gofmt` before submitting your pull
request to ensure consistent formatting.

### Testing

Before submitting your pull request, make sure your changes pass all the
existing tests, and add new ones if necessary.

## Building VMware Secrets Manager for Development

To build **VMware Secrets Manager** from source code and develop locally,
[follow the contributing guidelines here][contributing].

If you are a maintainer, and you are preparing a release,
[follow the release guidelines here][release].

[contributing]: https://vsecm.com/use-the-source/
[release]: https://vsecm.com/release/
[slack-invite]: https://join.slack.com/t/a-101-103-105-s/shared_invite/zt-287dbddk7-GCX495NK~FwO3bh_DAMAtQ "Join VSecM Slack"
