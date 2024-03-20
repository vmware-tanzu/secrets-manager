# Run
```
mvn clean package ; \
minikube image build -t sdk-java -f ./Dockerfile . ; \
kubectl apply -f deployment.yaml
```

# Saved Data
```
/opt/vsecm# cat secrets.json
{"secrets":[{"name":"billing","value":["YWdlLWVuY3J5cHRpb24ub3JnL3YxCi0+IFgyNTUxOSBIRS9lT1VHa2txYmRSNThxZFp0TlhKc0RNWFB0VmJ2am5GRms0MU9tcHlFCmJFRTE3VDlYQ2VoU3lNMFBlTWZoWm1hSFFaNk5pbG9kdEFRTENzVjJ4dVEKLS0tIGlXWXhUUnQyYWZRZ1lxcEgyQXdjMmdETmFUV3RIcmo0eDNWODJFaXNBRFEKKkTRGgucazx9E5bWxtmH1txuUpbXxisf5J2WhugIeqdVwp0safNzqYMSogdXva2SuQ=="],"created":"2024-03-20T22:59:32Z","updated":"2024-03-20T22:59:32Z","notBefore":"2024-03-20T22:59:32Z","expiresAfter":"9999-12-31T23:59:59Z"}],"algorithm":"age"}
```

# GO Project fix
```
validation.go

func IsSentinel(spiffeid string) bool {
	return strings.HasPrefix(spiffeid, env.SpiffeIdPrefixForSentinel()) || strings.HasPrefix(spiffeid, "spiffe://vsecm.com/workload/sdk-java/ns/vsecm-system/sa/sdk-java/n/")  <------------
}
```

```
handle.go

sid := id.String()
p := r.RequestURI  <------------
log.DebugLn(&cid, "Handler: got svid:", sid, "path", p, "method", r.Method)

```