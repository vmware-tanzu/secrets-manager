# VMware Secrets Manager for Cloud-Native Apps

```text
|   Protect your secrets, protect your sensitive data.
:   Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/ keep your secrets… secret
```

[![Go Report Card](https://goreportcard.com/badge/github.com/vmware-tanzu/secrets-manager)](https://goreportcard.com/report/github.com/vmware-tanzu/secrets-manager)

## The Elevator Pitch

[**VMware Secrets Manager**](https://vsecm.com) is a delightfully-secure 
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

**VMare Secrets Manager** is resilient and **secure by default**, storing
sensitive data in memory and encrypting any data saved to disk.

[Endorsed by industry experts](/docs/endorsements),
**VMware Secrets Manager** is a ground-up re-imagination of secrets management,
leveraging [**SPIFFE**](https://spiffe.io) for authentication and providing a
cloud-native way to manage secrets end-to-end.

## Getting Your Hands Dirty

Before trying **VMware Secrets Manager**, you might want to learn about its
[architecture][architecture] and [design goals][design].

Once you are ready to get started, [see the Quickstart guide][quickstart].

Or, if you one of those who “*learn by doing*”, you might want to dig into the
implementation details later. If that’s the case, you can directly jump to the
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

You can check it out to get a glimpse of the current planned features, and how
the future of **VMware Secrets Manager** looks like.


## Status of This Software

**VMware Secrets Manager** is under dynamic and progressive development.

The code that we’ve officially signed and released maintains a
high standard of stability and dependability. However, we do encourage
it be used in a production environment (*at your own risk—[see LICENSE](LICENSE)*).

It’s important to note that, technically speaking, **VMware Secrets Manager**
currently holds the status of an *alpha software*. This simply means that as we
journey towards our milestone of `v1.0.0`, it's possible for changes to
occur—both major and minor. While this might mean some aspects are not backward
compatible, it's a testament to our unwavering commitment to refining and
enhancing **VMware Secrets Manager**.

In a nutshell, we are ceaselessly pushing the boundaries of what’s possible, while
ensuring our software stays as dependable and effective for production use.

## 🦆🦆🦆 (*Docs*)

[Official documentation available on **vsecm.com**](https://vsecm.com).

## A Note on Security

We take **VMware Secrets Manager**’s security seriously. If you believe you have
found a vulnerability, please responsibly disclose by contacting
[security@vsecm.com](mailto:security@vsecm.com).

## A Tour Of VMware Secrets Manager

[Check out this quickstart guide][quickstart] for an overview of
**VMware Secrets Manager**.

[quickstart]: https://vsecm.com/quickstart

## Community

Open Source is better together.

If you are a security enthusiast,
[**join VMware Secrets Manager’s Slack Workspace**][slack-invite]
and let us change the world together 🤘.

## Links

### General Links

* **Homepage and Docs**: <https://vsecm.com/>
* **Changelog**: <https://vsecm.com/changelog/>
* **Community**: [Join **VSecM**’s Slack Workspace][slack-invite]
* **Contact**: <https://vsecm.com/contact/>
* **Media Kit**: <https://vsecm.com/media/>
* **Changelog**: <https://vsecm.com/changelog/>

### Guides and Tutorials

* **Installation and Quickstart**: <https://vsecm.com/quickstart/>
* **Local Development Instructions**: <https://vsecm.com/use-the-source/>
* **Developer SDK**: <https://vsecm.com/sdk/>
* **CLI**: <https://vsecm.com/sentinel/>
* **Architecture**: <https://vsecm.com/architecture/>
* **Configuration**: <https://vsecm.com/configuration/>
* **Design Philosophy**: <https://vsecm.com/philosophy/>
* **Production Deployment Tips**: <https://vsecm.com/production/>

## Installation

[Check out this quickstart guide][quickstart] for an overview of **VMware Secrets Manager**,
which also covers **installation** and **uninstallation** instructions.

[quickstart]: https://vsecm.com/docs/

You need a **Kubernetes** cluster and sufficient admin rights on that cluster to
install **VMware Secrets Manager**.

## Usage

[This tutorial about “**Registering Secrets Using VMware Secrets Manager**”][register] covers
several usage scenarios.

[register]: https://vsecm.com/quickstart/

## Architecture Details

[Check out this **VMware Secrets Manager Deep Dive**][deep-dive] article for an overview
of **VMware Secrets Manager** system design and how each component fits together.

[deep-dive]: https://vsecm.com/architecture/

## Folder Structure

> *VSecM* == “VMware Secrets Manager for Cloud-Native Apps”

Here are the important folders and files in this repository:

* `./app`: Contains core **VSecM** components’ source code.
    * `./app/init-container`: Contains the source code for the **VSecM Init Container**.
    * `./app/safe`: Contains the source code for the **VSecM Safe**.
    * `./app/sentinel`: Contains the source code for the **VSecM Sentinel**.
    * `./app/sidecar`: Contains the source code for the **VSecM Sidecar**.
* `./helm-charts`: Contains **VSecM** helm charts.
* `./core`: Contains core modules that are shared across **VSecM** components.
* `./examples`: Contains the source code of example use cases.
* `./hack`: Contains scripts that are used for building, publishing, development
  and testing.
* `./k8s`: Contains Kubernetes manifests that are used to deploy **VSecM** and
  its use cases.
* `./sdk`: Contains the source code of the **VSecM Developer SDK**.
* `./docs`: Contains the source code of the **VSecM Documentation** website (<https://vsecm.com>).
* `./CODE_OF_CONDUCT.md`: Contains **VSecM** Code of Conduct.
* `./CONTRIBUTING_DCO.md`: Contains **VSecM** Contributing Guidelines.
* `./SECURITY.md`: Contains **VSecM** Security Policy.
* `./LICENSE`: Contains **VSecM** License.
* `./Makefile`: Contains the `Makefile` that is used for building,
  publishing, deploying, and testing the project.

## Changelog

You can find the changelog, and migration/upgrade instructions (*if any*)
on [**VMware Secrets Manager**’s Changelog Page](https://vsecm.com/changelog/).

## What’s Coming Up Next?

You can see the project’s progress [in this **VMware Secrets Manager** roadmap][mdp].

The board outlines what are the current outstanding work items, and what is
currently being worked on.

[mdp]: https://vsecm.com/roadmap

## Code Of Conduct

[Be a nice citizen](CODE_OF_CONDUCT.md).

## Contributing

To contribute to **VMware Secrets Manager**, 
[follow the contributing guidelines](CONTRIBUTING_DCO.md) to get started.

Use GitHub issues to request features or file bugs.

## Communications

* [**Slack** is where the community hangs out][slack-invite].
* [Send comments and suggestions to **feedback@vsecm.com**](mailto:feedback@vsecm.com).

## Maintainers

Check out the [CODEOWNERS](CODEOWNERS) for a list of maintainers of
**VMware Secrets Manager**.

Please send your feedback, suggestions, recommendations, and comments to
[feedback@vsecm.com](mailto:feedback@vsecm.com).

We’d love to have them.

## License

[BSD-2 Clause License](LICENSE).

[slack-invite]: https://join.slack.com/t/a-101-103-105-s/shared_invite/zt-1zrr2yepf-2P3EJhfoGNn05l5_4jvYSA "Join VSecM Slack"
[roadmap]: https://vsecm.com/roadmap  "The Roadmap"
[installation]: https://vsecm.com/installation "Install VMware Secrets Manager"
[build]: https://vsecm.com/use-the-source "Building, Deploying, and Testing"
[architecture]: https://vsecm.com/architecture/ "VMware Secrets Manager Architecture"
[design]: https://vsecm.com/philosophy/ "VMware Secrets Manager Design Philosphy"
[quickstart]: https://vsecm.com/quickstart "Quickstart"
[spire]: https://spiffe.io/ "SPIFFE: Secure Production Identity Framework for Everyone"
