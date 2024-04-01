```text
|   Protect your secrets, protect your sensitive data.
:   Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/ keep your secrets... secret
```

## VMware Secrets Manager (*VSecM*) Inspector

**VSecM Inspector** is a utility workload that you can use to inspect secrets
assigned to workloads without shelling into the workloads.

You'll need cluster administrator privileges to use it because you'll require
to create a `ClusterSPIFFEID` that makes **VSecM Inspector** act on
behalf of the workload that you are inspecting.
