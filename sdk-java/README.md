# How to implement to your project

```
    <dependency>
        <groupId>org.wsecm</groupId>
        <artifactId>sdk-java</artifactId>
        <version>1.0.0</version>
    </dependency>
```

```
    public void start() {
        LOGGER.info("START....");
        String fetch = fetch("https://vsecm-safe.vsecm-system.svc.cluster.local:8443/sentinel/v1/secrets?reveal=true");
        LOGGER.info("RESULT .... " + fetch);
        TaskFactory taskFactory = new FetchSecretsTaskFactory();
        Future<String> watchedData = watch(taskFactory, "https://vsecm-safe.vsecm-system.svc.cluster.local:8443/sentinel/v1/secrets?reveal=true");
        if (watchedData.isDone()) {
            try {
                saveData(watchedData.get(), "/opt/vsecm/secrets.json");
            } catch (InterruptedException | ExecutionException e) {
                LOGGER.info("ERROR .... " + e.getMessage());
            }
        }
    }
```

# TEST INSTRUCTIONS
- If you want to test properly you should revert changes to this **[commit](https://github.com/vmware-tanzu/secrets-manager/pull/732/commits/caa16943b6e2aeac714486b5e73b2aa13f1515c1)**
- Check README