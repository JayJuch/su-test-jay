#!/bin/bash
set -e

# Stop and remove existing containers if running
docker rm -f fleet-monitor device-simulator 2>/dev/null || true

# Build and run fleet-monitor (detached so simulator can start after)
docker build -f ./cmd/fleet-monitor/Dockerfile -t fleet-monitor .
docker run --network host -d -p 6733:6733 --name fleet-monitor fleet-monitor

sleep 2

# Build and run simulator
docker build --platform=linux/arm64 -f ./simulator/Dockerfile-sim-arm -t device-simulator ./simulator
docker run --network host --platform=linux/arm64 -it --name device-simulator device-simulator:latest
