
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go build -o bin/webserver ./cmd/webserver

FROM ubuntu:22.04

# Install certificates and any other needed dependencies in one RUN layer
RUN apt-get update && \
    apt-get install -y curl ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/bin/webserver /webserver
COPY --from=builder /app/web/static/index.html /index.html

ENTRYPOINT ["/webserver"]