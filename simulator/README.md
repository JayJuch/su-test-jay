# Device Simulator

Pre-built binary for testing fleet-monitor service.

Run it inside 
* For arm macos

## Build container

```
docker build --platform=linux/arm64 -f ./simulator/Dockerfile-sim-arm -t device-simulator ./simulator
```

## Run

* runs on the host's network

```
docker run --network host --platform=linux/arm64 -it --name device-simulator device-simulator:latest
```
