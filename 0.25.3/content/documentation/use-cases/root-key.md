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

title = "Setting VSecM Root Key"
weight = 246
+++

## Situation Analysis

**VMware Secrets Manager** creates and stores its root cryptographic material
automatically during initialization. This is helpful because when **VSecM Safe**
crashes, or it is evicted, the root cryptographic material will be retained
and you won't have to manually unlock the system.

However, there are situations where you might want to set the root cryptographic
material externally. For example, you might want to set the root cryptographic
material to a value that is known to you and your organization. Or you might want
to set the root cryptographic material to a value that is known to a third-party
service that you trust. Or, you might want to control the root cryptographic
material due to regulatory requirements.

Whatever the reason, **VSecM Safe** allows you to set the root cryptographic
material externally. If you set the `VSECM_ROOT_KEY_INPUT_MODE_MANUAL` 
environment variable to `"true"`, **VSecM Safe** will seal itself and it will
wait for you to provide the root cryptographic material to begin operation.

When sealed, almost all API calls to **VSecM Safe** will fail. The only API
calls that will succeed are the ones that allow you to unseal **VSecM Safe**
by providing the root cryptographic material.

## High-Level Diagram

Open the image in a new tab to see the full-size version:

![High-Level Diagram](/assets/unseal.png "High-Level Diagram")

## Implementation

We will first make sure that **VSecM Safe** is switched to manual operation 
mode. For this, we define `VSECM_ROOT_KEY_INPUT_MODE_MANUAL` in its
`Deployment` manifest:

### Updating VSecM Safe Deployment

```yaml
# ./safe/Deployment.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
  name: vsecm-safe
  namespace: vsecm-system
  labels:
    app.kubernetes.io/name: vsecm-safe
    app.kubernetes.io/instance: vsecm
    app.kubernetes.io/part-of: vsecm-system
    app.kubernetes.io/version: "latest"
spec:
  replicas: 1
  selector:
    matchLabels:
      app.kubernetes.io/name: vsecm-safe
      app.kubernetes.io/instance: vsecm
      app.kubernetes.io/part-of: vsecm-system
  template:
    metadata:
      labels:
        app.kubernetes.io/name: vsecm-safe
        app.kubernetes.io/instance: vsecm
        app.kubernetes.io/part-of: vsecm-system
    spec:
      serviceAccountName: vsecm-safe
      containers:
        - name: main
          image: "vsecm/vsecm-ist-safe:latest"
          imagePullPolicy: IfNotPresent
          ports:
            - containerPort: 8443
              name: http
              protocol: TCP
          volumeMounts:
            - name: spire-agent-socket
              mountPath: /spire-agent-socket
              readOnly: true
            - name: vsecm-data
              mountPath: /data
            - name: vsecm-root-key
              mountPath: /key
              readOnly: true
# <--- BEGIN CHANGE
          env:
            - name: VSECM_ROOT_KEY_INPUT_MODE_MANUAL
              value: "true"
# ## <--- END CHANGE
          livenessProbe:
            httpGet:
              path: /
              port: 8081
            initialDelaySeconds: 1
            periodSeconds: 10
          readinessProbe:
            httpGet:
              path: /
              port: 8082
            initialDelaySeconds: 1
            periodSeconds: 10
          resources:
            requests:
              memory: 20Mi
              cpu: 5m
      volumes:
        - name: spire-agent-socket
          csi:
            driver: "csi.spiffe.io"
            readOnly: true
        - name: vsecm-data
          hostPath:
            path: /var/local/vsecm/data
            type: DirectoryOrCreate
        - name: vsecm-root-key
          secret:
            secretName: vsecm-root-key
            items:
              - key: KEY_TXT
                path: key.txt
```

Now, when we deploy **VSecM**, it will be in manual operation mode; we will
have to provide the root cryptographic material to unseal it to begin
registering secrets.

### Creating the Root Key

You can use **VSecM Keygen** via the following command to generate a new root key:

```bash
docker run --rm vsecm-keygen:latest

# Output
# AGE-SECRET-KEY-14DVY8Y0J4JQA45Z...truncated
# age1ghxkaqg5kkt8rl98x...truncated
# bc95a5e9e81fdaf40fe0exxx...truncated
```

Safe this output in a file named `key.txt`.

### Unsealing VSecM Safe

Here is how you provide the root key manually and unseal **VSecM Safe**:

```bash
kubectl exec "$SENTINEL" -n vsecm-system -- safe \
  -i "AGE-SECRET-KEY-1RZU...\nage1...\na6...ceec"
  
# Output:
#
# OK
```

### VSecM Safe Is Ready

After unsealing, **VSecM Safe** is ready for operation. 

By default, **VSecM Safe** uses a Kubernetes `Secret` to back up its root key,
so if **VSecM Safe** restarts, it will automatically fetch the root key from the
Kubernetes `Secret` and unseal itself, so you can "*sleep more*".

> **Note**
> 
> We are planning to add an "in-memory-only" root key mode in the future. This
> mode will not store the root key on disk or in Kubernetes. This mode will
> be useful for organizations that want to keep the root key in memory only.
> The downside of this mode is that if **VSecM Safe** restarts, it will have to
> be unsealed manually. 


If you re-key **VSecM Safe**, your secrets will be lost if you have
not backed them up.

> **Backup Your Secrets Before a Re-Key**
> 
> When **VSecM Safe**'s root key is re-keyed, the old root key is destroyed.
> If you have secrets encrypted with the old root key, you will not be able to
> decrypt them. Make sure you have backups of your secrets before re-keying
> **VSecM Safe**.
> 
> Follow the [**Revealing Secrets**][reveal-secrets] use case for more
> information on how to do this.

[reveal-secrets]: @/documentation/use-cases/retrieving-secrets.md

## Conclusion

**VMware Secrets Manager** offers flexible and secure management of root 
cryptographic materials. By default, it generates and stores these materials 
automatically, ensuring that they are preserved across system crashes or evictions. 

However, for organizations needing stricter control due to specific requirements or 
regulations, **VSecM Safe** supports manual setup through the 
`VSECM_ROOT_KEY_INPUT_MODE_MANUAL` environment variable. This mode enhances 
security by allowing only pre-approved root cryptographic materials to be used, 
although it necessitates manual intervention to unseal the system before it can 
operate. Implementing this manual setup involves updating the deployment configuration, 
generating a new root key, and manually providing this key to unseal VSecM Safe. 
Moreover, to streamline the re-unsealing process after restarts, a Kubernetes 
secret can be employed to store and automatically apply the root key. 

While the manual root key input approach provides robust control and complies 
with stringent security policies, it also introduces the need for careful 
management and potential operational overhead, especially in ensuring that 
secrets remain accessible following re-keying events. Hence, it is crucial to 
back up secrets (*and your cluster*) regularly and understand the implications 
of key management strategies to maintain both security and operational efficiency.

## List of Use Cases in This Section

{{ use_cases() }}

{{ edit() }}