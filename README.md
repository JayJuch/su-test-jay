# Safelyyou test - fleet monitor service

Service built based on the openapi.json specification.
It receives device telemetry data and provides stats from received data.

## Build and run test
Use this shell script to build, start service and run tests.

```
./run-test.sh
```

## Sequence of tasks completed

Most of the tasks were done using Claude Code.

1. New golang project which serves RESTAPI requests based on the openapi.json - this implemented almost a full working version.
2. Make use devices.csv config for device validation
3. Run simulator and app run as docker containers - for host security and consistent exec env.
4. Refactor generated code into standard golang project structure.
5. 1st git push
6. Optimize
    a. Each upload time was tracked in memory. To save memory and average calculation, keep a running average and keep it updated.
    b. Each heart beat time was tracked in memory. To save memory and uptime calculation, track just start and end timestamp of the heart beats range plus a count of all the heart beats. Uptime percent is ~ counts received / expected counts given time range
7. Refactor server.go in to heartbeat and uploadtime stats handlers.



## Future thoughts and enhancements

* Make a generic data processing handler similar to how standard metrics collectors work. Note that this would require some changes from the client (api contract). Note also that the optimization would need to be reworked as the generic solution would want to track each data points. Also, this could be solved with off the shelf products - statsd/otel/Prometheus/Grafana.

* Production issues - because these stats are tracked in a single instance's memory, making any update to the service could result in loss of stats - manifesting as either false alarm or worse yet, loss of failure events. Moreover, because these stats are not shared, there no way to scale the service.

* If we were to try and solve the data sharing issue, I would suggest using Redis cache as stats data is transient. Although analytics folks may want to store as much historical data as possible. If that was the case, we could fork the stream, and send it to some data lake.

