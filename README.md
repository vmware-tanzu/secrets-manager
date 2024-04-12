# VMware Secrets Manager for Cloud-Native Apps

[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/7793/badge)](https://www.bestpractices.dev/projects/7793)
[![Version](https://img.shields.io/github/v/release/vmware-tanzu/secrets-manager?color=blueviolet)](https://github.com/vmware-tanzu/secrets-manager/releases)
[![Contributors](https://img.shields.io/github/contributors/vmware-tanzu/secrets-manager.svg?color=orange)](https://github.com/vmware-tanzu/secrets-manager/graphs/contributors)
[![Slack](https://img.shields.io/badge/slack-vsecm-brightgreen.svg?logo=slack)](https://join.slack.com/t/a-101-103-105-s/shared_invite/zt-287dbddk7-GCX495NK~FwO3bh_DAMAtQ)
[![Twitch](https://img.shields.io/twitch/status/vadidekivolkan)](https://twitch.tv/vadidekivolkan)
[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/vsecm)](https://artifacthub.io/packages/helm/vsecm/vsecm)
[![License](https://img.shields.io/github/license/vmware-tanzu/secrets-manager)](https://github.com/vmware-tanzu/secrets-manager/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/vmware-tanzu/secrets-manager)](https://goreportcard.com/report/github.com/vmware-tanzu/secrets-manager)
[![Go Coverage](https://github.com/vmware-tanzu/secrets-manager/wiki/coverage.svg)](https://raw.githack.com/wiki/vmware-tanzu/secrets-manager/coverage.html)
[![Using Better Commits](https://img.shields.io/badge/better--commits-enabled?style=for-the-badge&logo=git&color=a6e3a1&logoColor=D9E0EE&labelColor=302D41)](https://github.com/Everduin94/better-commits)

```text
|   Protect your secrets, protect your sensitive data.
:   Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/ keep your secrets... secret
```

## The Elevator Pitch

[**VMware Secrets Manager**](https://vsecm.com/) is a delightfully-secure 
Kubernetes-native secrets store.

**VMware Secrets Manager** (*VSecM*) keeps your secrets secret.

With **VMware Secrets Manager**, you can rest assured that your sensitive data 
is always **secure** and **protected**.

**VMware Secrets Manager** is perfect for securely storing arbitrary configuration
information at a central location and securely dispatching it to workloads.

## Tell Me More

**VMware Secrets Manager** is a cloud-native secure store for secrets management.
It provides a minimal and intuitive API, ensuring practical security without
compromising user experience.

**VMware Secrets Manager** is resilient and **secure by default**, storing
sensitive data in memory and encrypting any data saved to disk.

[Endorsed by industry experts](https://vsecm.com/docs/endorsements/),
**VMware Secrets Manager** is a ground-up re-imagination of secrets management,
leveraging [**SPIFFE**](https://spiffe.io/) for authentication and providing a
cloud-native way to manage secrets end-to-end.

## Getting Your Hands Dirty

Before trying **VMware Secrets Manager**, you might want to learn about its
[architecture][architecture] and [design goals][design].

Once you are ready to start, [see the Quickstart guide][quickstart].

Or, if you are one of those who "*learn by doing*", you might want to dig into the
implementation details later. If that's the case, you can directly jump to the
fun part and [follow the steps here][installation] to install
**VMware Secrets Manager** to your Kubernetes cluster.

## Dive Into Example Use Cases

There are several examples demonstrating **VMware Secrets Manager** sample use
cases [inside the `./examples/` folder](./examples).

## Container Images

Pre-built container images of **VMware Secrets Manager** components can be found
at: <https://hub.docker.com/u/vsecm>.

## Build VMware Secrets Manager From the Source

[You can also build **VMware Secrets Manager** from the source][build].

## The Roadmap

[We publicly track all **VMware Secrets Manager** plans on this roadmap page][roadmap].

You can check it out to get a glimpse of the current planned features and how
the future of **VMware Secrets Manager** looks like.

## Status of This Software

**VMware Secrets Manager** is under dynamic and progressive development.

The code we've officially signed and released maintains a
high standard of stability and dependability. However, we do encourage
it to be used in a production environment (*at your own risk--[see LICENSE](LICENSE)*).

It's important to note that, technically speaking, **VMware Secrets Manager**
currently holds the status of an *alpha software*. This means that as we
journey towards our milestone of `v1.0.0`, it's possible for changes to
occur--both major and minor. While this might mean some aspects are not backward
compatible, it's a testament to our unwavering commitment to refining and
enhancing **VMware Secrets Manager**.

In a nutshell, we are ceaselessly pushing the boundaries of what's possible while
ensuring our software stays dependable and effective for production use.

## 🦆🦆🦆 (*Docs*)

* [Official documentation on **vsecm.com**](https://vsecm.com/).
* [Go Docs on **pkg.go.dev**](https://pkg.go.dev/github.com/vmware-tanzu/secrets-manager)

## A Note on Security

We take **VMware Secrets Manager**'s security seriously. If you believe you have
found a vulnerability, please [**follow this guideline**][vuln]
to responsibly disclose it.

[vuln]: https://github.com/vmware-tanzu/secrets-manager/blob/main/SECURITY.md

## A Tour Of VMware Secrets Manager

[Check out this quickstart guide][quickstart] for an overview of
**VMware Secrets Manager**.

[quickstart]: https://vsecm.com/docs/quickstart

## Community

Open Source is better together.

If you are a security enthusiast, join these communities
and let us change the world together 🤘:

* [Join **VMware Secrets Manager**'s Slack Workspace][slack-invite]
* [Join the **VMware Mecrets Manager** channel on Kampus' Discord Server][kampus]

## Links

### General Links

* **Homepage and Docs**: <https://vsecm.com/>
* **Changelog**: <https://vsecm.com/docs/changelog/>
* **Community**:
  * [Join **VMware Mecrets Manage**'s Slack Workspace][slack-invite]
  * [Join the **VMware Mecrets Manager** channel on Kampus' Discord Server][kampus]
* **Contact**: <https://vsecm.com/docs/community/>

### Guides and Tutorials

* **Installation and Quickstart**: <https://vsecm.com/docs/quickstart/>
* **Local Development Instructions**: <https://vsecm.com/docs/use-the-source/>
* **Developer SDK**: <https://vsecm.com/docs/sdk/>
* **CLI**: <https://vsecm.com/docs/cli/>
* **Architecture**: <https://vsecm.com/docs/architecture/>
* **Configuration**: <https://vsecm.com/docs/configuration/>
* **Design Philosophy**: <https://vsecm.com/docs/philosophy/>
* **Production Deployment Tips**: <https://vsecm.com/docs/production/>

## Installation

[Check out this quickstart guide][quickstart] for an overview of **VMware Secrets Manager**,
which also covers **installation** and **uninstallation** instructions.

[quickstart]: https://vsecm.com/docs/quickstart/

You need a **Kubernetes** cluster and sufficient admin rights on that cluster to
install **VMware Secrets Manager**.

## Usage

[Here is a list of step-by-step tutorials][register] covers
several usage scenarios that can show you where and how **VMware Secrets Manager**
could be used.

[register]: https://vsecm.com/docs/use-cases/

[use-cases]: https://vsecm.com/docs/use-cases

## Architecture Details

[Check out this **VMware Secrets Manager Deep Dive**][deep-dive] article for an overview
of **VMware Secrets Manager** system design and how each component fits together.

[deep-dive]: https://vsecm.com/docs/architecture/

## Folder Structure

> *VSecM* == "VMware Secrets Manager for Cloud-Native Apps"

Here are the important folders and files in this repository:

* `./app`: Contains core **VSecM** components' source code.
    * `./app/init_container`: Contains the source code for the **VSecM Init Container**.
    * `./app/inspector`: Contains the source code for the **VSecM Inspector**.
    * `./app/keygen`: Contains the source code for the **VSecM Keygen**.
    * `./app/keystone`: Contains the **VSecM KeyStone** source code.
    * `./app/safe`: Contains the **VSecM Safe** source code.
    * `./app/sentinel`: Contains the source code for the **VSecM Sentinel**.
    * `./app/sidecar`: Contains the source code for the **VSecM Sidecar**.
* `./ci`: Automation and CI/CD scripts.
* `./helm-charts`: Contains **VSecM** helm charts.
* `./core`: Contains core modules shared across **VSecM** components.
* `./dockerfiles`: Contains Dockerfiles for building **VSecM** container images.
* `./examples`: Contains the source code of example use cases.
* `./hack`: Contains scripts for building, publishing, development
  , and testing.
* `./k8s`: Contains Kubernetes manifests that are used to deploy **VSecM** and
  its use cases.
* `./sdk`: Contains the source code of the **VSecM Developer Go SDK**.
* `./sdk-cpp`: Contains the source code of the **VSecM Developer C++ SDK**.
* `./sdk-java`: Contains the source code of the **VSecM Developer Java SDK**.
* `./sdk-python`: Contains the source code of the **VSecM Developer Python SDK**.
* `./sdk-rust`: Contains the source code of the **VSecM Developer Rust SDK**.
* `./docs`: Contains the source code of the **VSecM Documentation** website (<https://vsecm.com>).
* `./CODE_OF_CONDUCT.md`: Contains **VSecM** Code of Conduct.
* `./CONTRIBUTING_DCO.md`: Contains **VSecM** Contributing Guidelines.
* `./SECURITY.md`: Contains **VSecM** Security Policy.
* `./LICENSE`: Contains **VSecM** License.
* `./Makefile`: The `Makefile` used for building,
  publishing, deploying, and testing the project.

## Branches

There are special long-living branches that the project maintains.

* `main`: This is the source code that is in active development. We try out best
  to keep it stable; however, there is no guarantees. We tag stable releases
  off of this branch during every release cut.
* `gh-pages`: This branch is where VSecM Helm charts are maintained.
  [ArtifactHub][artifacthub] references this branch.
* `docs`: This branch contains versioned documentation snapshots that we take  
   during releases.
* `tcx`: This is an internal "experimental" branch that is not meant for
  public consumption.

[artifacthub]: https://artifacthub.io/packages/helm/vsecm/vsecm

## Changelog

You can find the changelog and migration/upgrade instructions (*if any*)
on [**VMware Secrets Manager**'s Changelog Page](https://vsecm.com/docs/changelog/).

## What's Coming Up Next?

You can see the project's progress [in this **VMware Secrets Manager** roadmap][mdp].

[mdp]: https://vsecm.com/docs/roadmap/

## Code Of Conduct

[Be a nice citizen](CODE_OF_CONDUCT.md).

## Contributing

To contribute to **VMware Secrets Manager**, 
[follow the contributing guidelines](CONTRIBUTING.md) to get started.

Use GitHub issues to request features or file bugs.

## Communications

* [**Slack** is where the community hangs out][slack-invite].
* [Send comments and suggestions to **feedback@vsecm.com**](mailto:feedback@vsecm.com).

## Maintainers

Check out the [Maintainers Page](https://vsecm.com/docs/maintainers/) for a list 
of maintainers of **VMware Secrets Manager**.

Please send your feedback, suggestions, recommendations, and comments to
[feedback@vsecm.com](mailto:feedback@vsecm.com).

We'd love to have them.

## License

[BSD 2-Clause License](LICENSE).

[kampus]: https://discord.gg/kampus
[slack-invite]: https://join.slack.com/t/a-101-103-105-s/shared_invite/zt-287dbddk7-GCX495NK~FwO3bh_DAMAtQ "Join VSecM Slack"
[roadmap]: https://vsecm.com/docs/roadmap/  "The Roadmap"
[installation]: https://vsecm.com/docs/installation/ "Install VMware Secrets Manager"
[build]: https://vsecm.com/docs/use-the-source/ "Building, Deploying, and Testing"
[architecture]: https://vsecm.com/docs/architecture/ "VMware Secrets Manager Architecture"
[design]: https://vsecm.com/docs/philosophy/ "VMware Secrets Manager Design Philosophy"
[quickstart]: https://vsecm.com/docs/quickstart/ "Quickstart"
[spire]: https://spiffe.io/ "SPIFFE: Secure Production Identity Framework for Everyone"
