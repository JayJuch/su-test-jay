# Safelyyou test - fleet monitor service

Builds the service in a container.

## Build container

```
docker build -f ./cmd/fleet-monitor/Dockerfile -t fleet-monitor .
```

## Run

* runs on the host's network

```
docker run --network host -d -p 6733:6733 --name fleet-monitor fleet-monitor
```
