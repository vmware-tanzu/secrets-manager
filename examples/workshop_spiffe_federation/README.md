```text
# /*
# |    Protect your secrets, protect your sensitive data.
# :    Explore VMware Secrets Manager docs at https://vsecm.com/
# </
# <>/  keep your secrets... secret
# >/
# <>/' Copyright 2023-present VMware Secrets Manager contributors.
# >/'  SPDX-License-Identifier: BSD-2-Clause
# */
```

## Preparation

Assuming `microk8s` on ubuntu.

```bash
./infra/enable-k8s.sh
# ^ this will ask for IP ranges:
# Diablo:   10.211.55.110-10.211.55.119
# Mephisto: 10.211.55.120-10.211.55.129
# Baal:     10.211.55.130-10.211.55.139
# Azmodan:  10:211.55.140-10.211.55.149

# Next run this:
./infra/install-cluster-prerequisites.sh
```

## Bringing Up the Clusters

In each cluster folder (i.e., `./clusters/diablo`, `./clusters/mephisto`, etc)
execute the following.

```bash
cd <clusters/diablo|mephisto|baal|azmodan>
./hack/install-spire.sh
```

## Federating Clusters

After ensuring that SPIRE is up and running in all the clusters, execute
the following in each cluster.

```bash
cd ./hack 
./federate.sh
```

## Deploy the Workloads

```bash
cd clusters/<diablo|mephsito|baal|azmodan>
cd k8s/<control-plane-server|edge-store>
microk8s kubectl apply -f .
```
Install `control-plane-server` to `diablo`; `edge-store` to 
`mephisto`, `baal`, and `azmodan` (i.e., all edge stores).

Then check the logs of the edge stores. If everything went well, you should see
a new log line every ~10 seconds with an incremented sequence number.

## Other Helper Scripts

The `./infra` folder has the following scripts:

* `./infra/diablo.sh`: Displays cert information for `diablo` bundle endpoint.
* `./infra/mephisto.sh`: Displays cert information for `mephisto` bundle endpoint
* `./infra/baal.sh`: Displays cert information for `baal` bundle endpoint.
* `./infra/azmodan.sh`: Displays cert information for `azmodan` bundle endpoint.
* `./infra/reset.sh`: Resets the cluster and deletes the WORKSPACE folder (make
   sure to back up any important data before running this script; the script will
   NOT ask for confirmation!).
