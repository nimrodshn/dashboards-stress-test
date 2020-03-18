# dashboards-stress-test
Stress testing for `cloud.redhat.com/openshift/dashboards` . Used for internal testing.

## How to run?
Use `go test` to run the benchmarks:
```
go test -bench=. -args --token=$NSHNEOR_TEST_TOKEN --url=https://clusters-service.apps-crc.testing --insecure=true --concurrent-requests=10
```

The above command creates ten concurrent requests for the benchmark test to run.

