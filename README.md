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

