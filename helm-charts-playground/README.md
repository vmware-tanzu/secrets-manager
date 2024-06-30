```text
|   Protect your secrets, protect your sensitive data.
:   Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/ keep your secrets... secret
```

## About

This is a temporary folder aimed as a playground to align the behavior of 
the official SPIFFE `helm-charts-hardened` SPIRE helm charts with the
VMware Secrets Manager-managed SPIRE helm charts.

Once we establish the alignment, we will delete this folder.

## To Do

- Create a script to parse `k8s/$version/spire.yaml` and `k8s/$version/crds/*` 
  to create `vsecm-manifests/spire/*` files automatically. This will make
  diffing easier and reduce human errors.
- Do the same for the generated `spire-manifests.openshift.yaml` and
  `spire-manifefest-no-openshift.yaml` files.
- Create issues based on the delta you find between the two sets of manifests.
- We may decide to keep this folder, since it's a good way to keep track of
  the changes we make to the official SPIFFE helm charts.

