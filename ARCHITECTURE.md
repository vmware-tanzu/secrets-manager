# VSecM Architecture

![VSecM Logo](https://github.com/vmware-tanzu/secrets-manager/assets/1041224/885c11ac-7269-4344-a376-0d0a0fb082a7)

The VMware Secrets Manager (*VSecM*) architecture consists of the following
main system components:

* [SPIRE][spire]: Acting as the identity control plane.
* **VSecM Safe**: The secure secrets store.
* **VSecM Sentinel**: Entry point to the system where secrets can be registered
  to the workloads.
* **VSecM Keystone**: A pod that is enabled only when the entire **VSecM**
  system reconciles.

![actors.jpg](assets/actors.jpg)

For more details, [you can view the full architecture documentation here][architecture].

[spire]: https://spiffe.io/downloads/ "SPIRE"
[architecture]: https://vsecm.com/docs/architecture "VSecM Architecture"
