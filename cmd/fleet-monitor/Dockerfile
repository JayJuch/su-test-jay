FROM golang:1.25-alpine AS builder

WORKDIR /build
COPY go.mod ./
COPY cmd/fleet-monitor/main.go ./
COPY internal/ ./internal/
RUN go test ./...
RUN go build -o fleet-monitor .

FROM alpine:3.22.4

WORKDIR /app
COPY --from=builder /build/fleet-monitor ./
COPY devices.csv ./

EXPOSE 6733

ENTRYPOINT ["/app/fleet-monitor"]
